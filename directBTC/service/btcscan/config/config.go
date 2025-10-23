package config

type Config struct {
	ConfirmThreshold uint64
	DataSource       string
	MempoolClient    MempoolClient
	MemScanSpec      string
	ConfirmScanSpec  string
	//mainnet,signet
	Network string
}

type MempoolClient struct {
	Host      string
	Request   int
	PeriodSec int
}
