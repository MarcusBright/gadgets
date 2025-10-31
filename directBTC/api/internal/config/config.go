// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	btcscanconfig "directBTC/service/btcscan/config"
	evmscanconfig "directBTC/service/evmscan/config"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource    string
	LogSlack      string   `json:",optional"`
	NotifySlack   string   `json:",optional"`
	SysEvmAddress []string `json:",optional"`
	TinyTry       uint64   `json:",default=0"`
	BTCScan       btcscanconfig.Config
	EvmScan       evmscanconfig.Config
}
