package btcscan

import (
	"directBTC/model"
	"directBTC/service/btcscan/config"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var HOST string

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	HOST = os.Getenv("MEMPOOL_HOST")
	fmt.Printf("HOST: %s\n", HOST)
}

func TestNewTrans(t *testing.T) {
	config := &config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
		MempoolClient: config.MempoolClient{
			Host:      HOST,
			Request:   3,
			PeriodSec: 1,
		},
		ConfirmThreshold: 2,
	}
	scanner := NewScanner(config)
	scanner.NewTrans()
	//get all btc transactions from memoryDB
	var btcTran []model.BtcTran
	// logx.Must(scanner.memoryDB.Find(&btcTran).Error)
	t.Log(func() string {
		jsonStr, _ := json.Marshal(btcTran)
		return string(jsonStr)
	}())
	// address: bc1qnxuxe6leghrw3pezr37wd6dzufrsnpfrvrhawy
}

func TestUpdateConfirmNumber(t *testing.T) {
	config := &config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
		MempoolClient: config.MempoolClient{
			Host:      HOST,
			Request:   3,
			PeriodSec: 1,
		},
		ConfirmThreshold: 2,
	}
	scanner := NewScanner(config)
	scanner.UpdateConfirmNumber()
}

func TestMultiMemDB(t *testing.T) {
	memoryDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		// Logger: gormz.NewGormLogger(),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "direct_btc_",
		},
	})
	logx.Must(err)
	logx.Must(memoryDB.AutoMigrate(&model.BtcTran{}))
	go func() {
		secondDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
			// Logger: gormz.NewGormLogger(),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "direct_btc_",
			},
		})
		logx.Must(err)
		logx.Must(secondDB.AutoMigrate(&model.BtcTran{}))
		for {
			//read
			var btcTran []model.BtcTran
			logx.Must(secondDB.Find(&btcTran).Error)
			t.Logf("len: %d", len(btcTran))
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		//write
		logx.Must(memoryDB.Create(&model.BtcTran{
			TransactionHash: fmt.Sprintf("%d", i),
			TreasuryAddress: datatypes.JSON([]byte("[]")),
			InputUtxo:       datatypes.JSON([]byte("[]")),
		}).Error)
	}
}
