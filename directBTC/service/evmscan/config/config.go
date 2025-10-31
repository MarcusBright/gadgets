package config

type Config struct {
	DataSource   string
	ChainInfo    []ChainInfo
	LogsScanSpec string
	NotifySlack  string `json:",optional"`
}

type EvmClient struct {
	ChainId   uint
	Host      string
	Request   int
	PeriodSec int
}

type ChainInfo struct {
	Client           EvmClient
	Miner            string
	ProcessIdxOffset uint64
}
