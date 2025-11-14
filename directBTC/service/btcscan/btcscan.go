package btcscan

import (
	mempoolspace "directBTC/clientrpc/mempool.space"
	"directBTC/model"
	"directBTC/pkg/gormz"
	"directBTC/service/btcscan/config"
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"

	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Scanner struct {
	client   *mempoolspace.Client
	database *gorm.DB
	// memoryDB *gorm.DB
	config *config.Config
}

func NewScanner(c *config.Config) *Scanner {
	db, err := gorm.Open(mysql.Open(c.DataSource))
	if c.SqlLog {
		db.Logger = gormz.NewGormLogger()
	}
	logx.Must(err)

	// memoryDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
	// 	Logger: gormz.NewGormLogger(),
	// })
	// logx.Must(err)
	// logx.Must(memoryDB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").
	// 	AutoMigrate(&model.BtcTran{}))
	// logx.Must(memoryDB.AutoMigrate(&model.BtcTran{}))

	return &Scanner{
		client:   mempoolspace.NewClient(c.MempoolClient.Host, c.MempoolClient.Request, c.MempoolClient.PeriodSec),
		database: db,
		// memoryDB: memoryDB,
		config: c,
	}
}

func (s *Scanner) NewTrans() {
	btcAddress, err := new(model.Cursor).GetAllBtcAddress(s.database)
	if err != nil {
		logx.Errorf("GetAllBtcAddress failed: %v", err)
		return
	}
	logx.Infof("GetAllBtcAddress: %v", len(btcAddress))

	for _, addr := range btcAddress {
		logx.Infof("GetAddressNewTransactions address: %s", addr.Address)
		txs, err := s.client.GetAddressNewTransactions(addr.Address, addr.Txhash)
		if err != nil {
			logx.Errorf("GetAddressNewTransactions failed: %v", err)
			continue
		}
		group := lo.GroupBy(txs, func(item mempoolspace.AddressTransaction) string {
			if item.Status.Confirmed {
				return "mined"
			}
			return "mempool"
		})
		var minedGroup, mempoolGroup = group["mined"], group["mempool"]
		minedAddrTrans := s.filterAndMapTrans(minedGroup, lo.Map(btcAddress, func(item model.Cursor, _ int) string {
			return item.Address
		}))
		mempoolAddrTrans := s.filterAndMapTrans(mempoolGroup, lo.Map(btcAddress, func(item model.Cursor, _ int) string {
			return item.Address
		}))
		logx.Infof("GetAddressNewTransactions mempoolGroup: %d, minedGroup: %d, minedAddrTrans: %d, mempoolAddrTrans: %d",
			len(mempoolGroup), len(minedGroup), len(minedAddrTrans), len(mempoolAddrTrans))
		newTrans := append(mempoolAddrTrans, minedAddrTrans...)
		newTrans = lo.UniqBy(newTrans, func(item model.BtcTran) string {
			return item.TransactionHash
		})
		if len(newTrans) == 0 {
			continue
		}
		if err = s.database.Transaction(func(tx *gorm.DB) error {
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).
				CreateInBatches(&newTrans, 50).Error; err != nil {
				return err
			}
			if len(minedGroup) != 0 {
				if err := tx.Model(&model.Cursor{}).Where("address = ?", addr.Address).
					Update("txhash", minedGroup[0].Txid).Error; err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			logx.Errorf("CreateInBatches mined failed: %v", err)
		}
	}
	/*if err := s.memoryDB.Transaction(func(tx *gorm.DB) error {
		// trucate memoryDB
		// DELETE FROM Users;
		//UPDATE SQLITE_SEQUENCE SET seq = 0 WHERE name = 'Users';
		//VACUUM;
		if err := tx.Exec("DELETE FROM btc_trans;UPDATE SQLITE_SEQUENCE SET seq = 0 WHERE name = 'btc_trans'").Error; err != nil {
			return err
		}
		if len(mempoolTrans) != 0 {
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).
				CreateInBatches(&mempoolTrans, 50).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		logx.Errorf("Truncate mempoolTrans failed: %v", err)
	}*/
}

func (s *Scanner) filterAndMapTrans(trans []mempoolspace.AddressTransaction, treasuryAddress []string) []model.BtcTran {
	return lo.FilterMap(trans, func(item mempoolspace.AddressTransaction, _ int) (model.BtcTran, bool) {
		if _, _, in := lo.FindIndexOf(item.Vin, func(v mempoolspace.Vin) bool {
			return lo.Contains(treasuryAddress, v.Prevout.ScriptpubkeyAddress)
		}); in {
			logx.Infof("treasuryAddress is input,address:%s, txid: %s", treasuryAddress, item.Txid)
			return model.BtcTran{}, false
		}

		var amount int64
		var treasuryAddressIn []string
		for _, v := range item.Vout {
			if lo.Contains(treasuryAddress, v.ScriptpubkeyAddress) {
				amount += v.Value
				treasuryAddressIn = append(treasuryAddressIn, v.ScriptpubkeyAddress)
			}
		}
		amount = amount - s.config.FeeSatoshi
		if amount <= 0 || len(treasuryAddressIn) == 0 {
			logx.Errorf("treasuryAddress amount is 0 or lower,address:%s, txid: %s", treasuryAddress, item.Txid)
			return model.BtcTran{}, false
		}
		return model.BtcTran{
			TransactionHash: item.Txid,
			TreasuryAddress: func() datatypes.JSON {
				b, _ := json.Marshal(lo.Uniq(treasuryAddressIn))
				return datatypes.JSON(b)
			}(),
			AmountSatoshi: fmt.Sprintf("%d", amount),
			FeeSatoshi:    fmt.Sprintf("%d", s.config.FeeSatoshi),
			InputUtxo: func() datatypes.JSON {
				utxoAddress := lo.Map(item.Vin, func(item mempoolspace.Vin, _ int) string {
					return item.Prevout.ScriptpubkeyAddress
				})
				b, _ := json.Marshal(utxoAddress)
				return datatypes.JSON(b)
			}(),
			Status:           model.BtcTranStatusInit,
			BlockNumber:      item.Status.BlockHeight,
			BlockTime:        item.Status.BlockTime,
			ConfirmThreshold: s.config.ConfirmThreshold,
		}, true
	})
}

func (s *Scanner) UpdateConfirmNumber() {
	var btcTran []model.BtcTran
	if err := s.database.
		// Where("status in ?", []string{model.BtcTranStatusInit, model.BtcTranStatusBinded}).
		Where("confirm_number < confirm_threshold").Find(&btcTran).Error; err != nil {
		logx.Errorf("Find database btcTran failed: %v", err)
		return
	}
	if len(btcTran) == 0 {
		logx.Infof("no btcTran need update")
		return
	}
	latestHeight, err := s.client.GetLatestBlockNumber()
	if err != nil {
		logx.Errorf("GetLatestBlockNumber failed: %v", err)
		return
	}
	for k, tran := range btcTran {
		tx, err := s.client.GetTx(tran.TransactionHash)
		if err != nil {
			logx.Errorf("GetTx %s failed: %v", tran.TransactionHash, err)
			continue
		}
		if !tx.Status.Confirmed || tx.Status.BlockHeight == 0 {
			logx.Errorf("tx not confirmed, txid: %s", tran.TransactionHash)
			continue
		}
		confirmedNumber := int64(latestHeight) - int64(tx.Status.BlockHeight)
		if confirmedNumber < 0 {
			logx.Errorf("confrimedNumber[%d], latestHeight:%d, tx.BlockHeight:%d", confirmedNumber,
				latestHeight, tx.Status.BlockHeight)
			continue
		}
		confirmedNumber++
		btcTran[k].ConfirmNumber = uint64(confirmedNumber)
		btcTran[k].BlockNumber = tx.Status.BlockHeight
		btcTran[k].BlockTime = tx.Status.BlockTime
	}
	if err := s.database.Save(&btcTran).Error; err != nil {
		logx.Errorf("Save btcTran failed: %v", err)
		return
	}
}
