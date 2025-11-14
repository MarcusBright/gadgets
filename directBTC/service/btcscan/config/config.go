package config

type Config struct {
	ConfirmThreshold uint64
	DataSource       string `json:",inherit"`
	SqlLog           bool   `json:",optional,default=false,inherit"`
	FeeSatoshi       int64  `json:",default=0"`
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
