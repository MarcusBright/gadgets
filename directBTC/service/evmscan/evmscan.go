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
		var cursor model.Cursor
		rest := s.database.Model(&model.Cursor{}).Where("chain_id = ?", chain.Client.ChainId).
			Where("is_btc = ?", false).Where("address = ?", model.CursorEvmLogScan).First(&cursor)
		logx.Must(rest.Error)

		latestBlockNumber, err := s.evmClients[k].BlockNumber(context.Background())
		if err != nil {
			logx.Errorf("get chain: %v, latest block number failed, err: %v", chain.Client.ChainId, err)
			continue
		}
		latestBlockNumber = latestBlockNumber - 2 //delay block
		needScanBlocksNumber := int64(latestBlockNumber) - int64(cursor.BlockNumber)
		if needScanBlocksNumber <= 0 {
			logx.Infof("chain: %v, no new block", chain.Client.ChainId)
			continue
		}
		if needScanBlocksNumber > 1000 {
			needScanBlocksNumber = 1000
		}
		start, end := int64(cursor.BlockNumber+1), int64(cursor.BlockNumber)+1+needScanBlocksNumber
		logx.Infof("chain: %v, need scan blocks start:%v, end:%v", chain.Client.ChainId, start, end)
		logs, err := s.evmClients[k].FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: big.NewInt(start),
			ToBlock:   big.NewInt(end),
			Addresses: []common.Address{common.HexToAddress(chain.Miner)},
		})
		if err != nil {
			logx.Errorf("get chain: %v, filter logs failed, err: %v", chain.Client.ChainId, err)
			continue
		}
		var count int64
		if err := s.database.Model(&model.BtcTran{}).Where("status in (?)",
			[]string{model.BtcTranStatusRecievedInEvm, model.BtcTranStatusApprovedInEvm, model.BtcTranStatusRejectedInEvm}).
			Count(&count).Error; err != nil {
			logx.Errorf("get chain: %v, count error:%v", chain.Client.ChainId, err)
			continue
		}
		events, err := s.processLogs(logs, count+int64(chain.ProcessIdxOffset), s.evmClients[k])
		if err != nil {
			logx.Errorf("process logs failed, err: %v", err)
			continue
		}
		//save events
		if err := s.database.Transaction(func(tx *gorm.DB) error {
			for _, event := range events {
				columns := make(map[string]any)
				//get compare data
				var btcTran model.BtcTran
				if err := tx.Model(&model.BtcTran{}).Where("transaction_hash = ?", strings.TrimPrefix(event.BtcTransaction, "0x")).
					First(&btcTran).Error; err != nil {
					logx.Errorf("get btc tran failed, hash: %v, err: %v", event.BtcTransaction, err)
					return err
				}
				if btcTran.BindedEvmAddress != event.Recipient {
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
				slack.SendTo(s.config.NotifySlack, fmt.Sprintf("btcHash[%v] status to %v, evm hash[%v]", event.BtcTransaction, status, event.EventHash))
			}
			//save cursor
			cursor.BlockNumber = uint64(end)
			if err := tx.Save(&cursor).Error; err != nil {
				logx.Errorf("save cursor failed, err: %v", err)
				return err
			}
			return nil
		}); err != nil {
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
			})
		}
	}
	return events, nil
}
