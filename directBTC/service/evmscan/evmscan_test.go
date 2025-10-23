package evmscan

import (
	"directBTC/service/evmscan/config"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func TestLogScan(t *testing.T) {
	chainInfo := config.ChainInfo{
		Client: config.EvmClient{
			ChainId:   1,
			Host:      os.Getenv("MAIN_INFURA"),
			Request:   1,
			PeriodSec: 1,
		},
		Miner:            "0x91fD8C7a5FDA7d52Ab41Bbe423eEdd3A65d64500",
		ProcessIdxOffset: 10,
	}
	config := config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
		ChainInfo:  []config.ChainInfo{chainInfo},
	}
	scanner := NewScanner(&config)
	scanner.LogScan()
}
