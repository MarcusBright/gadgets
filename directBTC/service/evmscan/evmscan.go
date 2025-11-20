package evmscan

import (
	"context"
	"directBTC/clientrpc/goeth"
	"directBTC/model"
	"directBTC/pkg/gormz"
	"directBTC/pkg/slack"
	"directBTC/service/contract/directbtcminter"
	"directBTC/service/evmscan/config"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Scanner struct {
	database           *gorm.DB
	config             *config.Config
	evmClients         []*ethclient.Client
	directBTCMinter    *directbtcminter.Directbtcminter
	directBTCMinterAbi *abi.ABI
}

func NewScanner(c *config.Config) *Scanner {
	db, err := gorm.Open(mysql.Open(c.DataSource))
	if c.SqlLog {
		db.Logger = gormz.NewGormLogger()
	}
	logx.Must(err)
	evmClients := lo.Map(c.ChainInfo, func(item config.ChainInfo, index int) *ethclient.Client {
		return goeth.NewClient(item.Client.Host, item.Client.Request, item.Client.PeriodSec)
	})
	directBTCMinter, _ := directbtcminter.NewDirectbtcminter(common.HexToAddress(c.ChainInfo[0].Miner), evmClients[0])
	directBTCMinterAbi, _ := directbtcminter.DirectbtcminterMetaData.GetAbi()
	return &Scanner{
		database:           db,
		config:             c,
		evmClients:         evmClients,
		directBTCMinter:    directBTCMinter,
		directBTCMinterAbi: directBTCMinterAbi,
	}
}

func (s *Scanner) LogScan() {
	for k, chain := range s.config.ChainInfo {
		logx.Infof("chain: %v, logscan", chain.Client.ChainId)

		cursor, err := s.getCursor(uint64(chain.Client.ChainId))
		logx.Must(err)

		start, end, err := s.getScanRange(s.evmClients[k], cursor.BlockNumber)
		if err != nil {
			logx.Errorf("get chain: %v, latest block number failed, err: %v", chain.Client.ChainId, err)
			continue
		}
		if start == 0 || end == 0 {
			logx.Infof("chain: %v, no new block", chain.Client.ChainId)
			continue
		}
		logx.Infof("chain: %v, need scan blocks start:%v, end:%v", chain.Client.ChainId, start, end)

		logs, err := s.fetchLogs(s.evmClients[k], start, end, chain.Miner)
		if err != nil {
			logx.Errorf("get chain: %v, filter logs failed, err: %v", chain.Client.ChainId, err)
			continue
		}

		count, err := s.getPendingCount()
		if err != nil {
			logx.Errorf("get chain: %v, count error:%v", chain.Client.ChainId, err)
			continue
		}

		events, err := s.processLogs(logs, count+int64(chain.ProcessIdxOffset), s.evmClients[k])
		if err != nil {
			logx.Errorf("process logs failed, err: %v", err)
			continue
		}

		if err := s.saveScanResult(events, chain, cursor, end); err != nil {
			logx.Errorf("save events failed, err: %v", err)
			continue
		}
	}
}

type MinterEvent struct {
	RecievedEvmTxHash string
	AcceptedEvmTxHash string
	RejectedEvmTxHash string
	ProcessIdx        int64
	BtcTransaction    string
	Recipient         string
	Amount            string
	EventHash         string
	EventBlockNumber  uint64
	EventBlockTime    uint64
}

func (s *Scanner) processLogs(logs []types.Log, index int64, client *ethclient.Client) ([]*MinterEvent, error) {
	events := make([]*MinterEvent, 0)
	for _, log := range logs {
		if log.Removed {
			logx.Errorf("log removed, hash:%v, blockNumber:%v, blockHash:%v", log.TxHash, log.BlockNumber, log.BlockHash)
			return nil, fmt.Errorf("log removed")
		}
		transactionRecipient, err := client.TransactionReceipt(context.Background(), log.TxHash)
		if err != nil {
			logx.Errorf("get transaction receipt failed, err: %v", err)
			return nil, err
		}
		if transactionRecipient.Status != types.ReceiptStatusSuccessful {
			logx.Infof("transaction status not successful, hash: %v, status: %v", log.TxHash.Hex(), transactionRecipient.Status)
			continue
		}
		eventName, err := s.directBTCMinterAbi.EventByID(log.Topics[0])
		if err != nil {
			logx.Errorf("get event name failed, err: %v", err)
			return nil, err
		}
		switch eventName.Name {
		case "Received":
			receivedEvent, err := s.directBTCMinter.ParseReceived(log)
			if err != nil {
				logx.Errorf("parse received event failed, err: %v", err)
				return nil, err
			}
			events = append(events, &MinterEvent{
				BtcTransaction:    common.BytesToHash(receivedEvent.TxHash[:]).Hex(),
				RecievedEvmTxHash: log.TxHash.Hex(),
				ProcessIdx:        index,
				Recipient:         receivedEvent.Recipient.Hex(),
				Amount:            receivedEvent.Amount.String(),
				EventHash:         log.TxHash.Hex(),
				EventBlockNumber:  log.BlockNumber,
				EventBlockTime:    log.BlockTimestamp,
			})
			index++
		case "Accepted":
			acceptedEvent, err := s.directBTCMinter.ParseAccepted(log)
			if err != nil {
				logx.Errorf("parse accepted event failed, err: %v", err)
				return nil, err
			}
			events = append(events, &MinterEvent{
				BtcTransaction:    common.BytesToHash(acceptedEvent.TxHash[:]).Hex(),
				AcceptedEvmTxHash: log.TxHash.Hex(),
				ProcessIdx:        -1,
				Recipient:         acceptedEvent.Recipient.Hex(),
				Amount:            acceptedEvent.Amount.String(),
				EventHash:         log.TxHash.Hex(),
				EventBlockNumber:  log.BlockNumber,
				EventBlockTime:    log.BlockTimestamp,
			})
		case "Rejected":
			rejectedEvent, err := s.directBTCMinter.ParseRejected(log)
			if err != nil {
				logx.Errorf("parse rejected event failed, err: %v", err)
				return nil, err
			}
			events = append(events, &MinterEvent{
				BtcTransaction:    common.BytesToHash(rejectedEvent.TxHash[:]).Hex(),
				RejectedEvmTxHash: log.TxHash.Hex(),
				ProcessIdx:        -1,
				Recipient:         rejectedEvent.Recipient.Hex(),
				Amount:            rejectedEvent.Amount.String(),
				EventHash:         log.TxHash.Hex(),
				EventBlockNumber:  log.BlockNumber,
				EventBlockTime:    log.BlockTimestamp,
			})
		}
	}
	return events, nil
}

func (s *Scanner) getCursor(chainID uint64) (*model.Cursor, error) {
	var cursor model.Cursor
	err := s.database.Model(&model.Cursor{}).Where("chain_id = ?", chainID).
		Where("is_btc = ?", false).Where("address = ?", model.CursorEvmLogScan).First(&cursor).Error
	return &cursor, err
}

func (s *Scanner) getScanRange(client *ethclient.Client, cursorBlock uint64) (int64, int64, error) {
	latestBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, 0, err
	}
	latestBlockNumber = latestBlockNumber - 1 //delay block
	needScanBlocksNumber := int64(latestBlockNumber) - int64(cursorBlock)
	if needScanBlocksNumber <= 0 {
		return 0, 0, nil
	}
	if needScanBlocksNumber > 1000 {
		needScanBlocksNumber = 1000
	}
	start := int64(cursorBlock + 1)
	end := int64(cursorBlock) + needScanBlocksNumber
	return start, end, nil
}

func (s *Scanner) fetchLogs(client *ethclient.Client, start, end int64, minerAddress string) ([]types.Log, error) {
	return client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(start),
		ToBlock:   big.NewInt(end),
		Addresses: []common.Address{common.HexToAddress(minerAddress)},
	})
}

func (s *Scanner) getPendingCount() (int64, error) {
	var count int64
	err := s.database.Model(&model.BtcTran{}).Where("status in (?)",
		[]string{model.BtcTranStatusRecievedInEvm, model.BtcTranStatusApprovedInEvm, model.BtcTranStatusRejectedInEvm}).
		Count(&count).Error
	return count, err
}

func (s *Scanner) saveScanResult(events []*MinterEvent, chain config.ChainInfo, cursor *model.Cursor, end int64) error {
	return s.database.Transaction(func(tx *gorm.DB) error {
		for _, event := range events {
			columns := make(map[string]any)
			//get compare data
			var btcTran model.BtcTran
			if err := tx.Model(&model.BtcTran{}).Where("transaction_hash = ?", strings.TrimPrefix(event.BtcTransaction, "0x")).
				First(&btcTran).Error; err != nil {
				logx.Errorf("get btc tran failed, hash: %v, err: %v", event.BtcTransaction, err)
				return err
			}
			if !strings.EqualFold(btcTran.BindedEvmAddress, event.Recipient) {
				logx.Errorf("btc tran binded evm address not equal, hash: %v, binded: %v, event: %v",
					event.BtcTransaction, btcTran.BindedEvmAddress, event.Recipient)
				return fmt.Errorf("not match recipient")
			}
			if btcTran.AmountSatoshi != event.Amount {
				logx.Errorf("btc tran amount not equal, hash: %v, amount: %v, event: %v",
					event.BtcTransaction, btcTran.AmountSatoshi, event.Amount)
				return fmt.Errorf("not match amount")
			}
			if btcTran.ChainId != chain.Client.ChainId {
				logx.Errorf("btc tran chain id not equal, hash: %v, chain id: %v, event: %v",
					event.BtcTransaction, btcTran.ChainId, chain.Client.ChainId)
				return fmt.Errorf("not match chain id")
			}
			status := ""
			if event.RecievedEvmTxHash != "" {
				columns["recieved_evm_tx_hash"] = event.RecievedEvmTxHash
				status = model.BtcTranStatusRecievedInEvm
			}
			if event.AcceptedEvmTxHash != "" {
				columns["accepted_evm_tx_hash"] = event.AcceptedEvmTxHash
				status = model.BtcTranStatusApprovedInEvm
			}
			if event.RejectedEvmTxHash != "" {
				columns["rejected_evm_tx_hash"] = event.RejectedEvmTxHash
				status = model.BtcTranStatusRejectedInEvm
			}
			if event.ProcessIdx != -1 {
				columns["process_idx"] = event.ProcessIdx
			}
			columns["status"] = status
			if err := tx.Model(&model.BtcTran{}).Where("transaction_hash = ?", strings.TrimPrefix(event.BtcTransaction, "0x")).
				Updates(columns).Error; err != nil {
				logx.Errorf("update btc tran failed, hash: %v, err: %v", event.BtcTransaction, err)
				return err
			}
			//save evm transaction
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Model(&model.EvmHashInfo{}).Create(&model.EvmHashInfo{
				TransactionHash: event.EventHash,
				BlockNumber:     event.EventBlockNumber,
				BlockTime:       event.EventBlockTime,
			}).Error; err != nil {
				logx.Errorf("create evmHashInfo failed, hash: %v, err: %v", event.EventHash, err)
				return err
			}
			slack.SendTo(s.config.NotifySlack, fmt.Sprintf("btcHash[%v] status to %v, evm hash[%v]", event.BtcTransaction, status, event.EventHash))
		}
		//save cursor
		cursor.BlockNumber = uint64(end)
		if err := tx.Save(cursor).Error; err != nil {
			logx.Errorf("save cursor failed, err: %v", err)
			return err
		}
		return nil
	})
}
