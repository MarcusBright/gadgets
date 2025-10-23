package config

type Config struct {
	DataSource   string
	ChainInfo    []ChainInfo
	LogsScanSpec string
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
