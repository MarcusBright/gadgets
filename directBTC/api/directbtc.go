// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"

	"directBTC/api/internal/config"
	"directBTC/api/internal/handler"
	"directBTC/api/internal/svc"
	"directBTC/pkg/slack"
	"directBTC/service/btcscan"
	"directBTC/service/crons"
	"directBTC/service/evmscan"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/directbtc-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.PrintRoutes()
	//log
	if c.LogSlack != "" {
		logx.AddWriter(logx.NewWriter(slack.NewSlackWriter(c.LogSlack)))
		logx.AddGlobalFields(logx.Field("server", c.Name))
	}
	//cron
	crontab := crons.New()
	btcScan := btcscan.NewScanner(&c.BTCScan)
	evmScan := evmscan.NewScanner(&c.EvmScan)
	if c.BTCScan.MemScanSpec != "" {
		_, err := crontab.AddFunc(c.BTCScan.MemScanSpec, btcScan.NewTrans)
		logx.Must(err)
		logx.Infof("add cron btc mem scan spec: %v", c.BTCScan.MemScanSpec)
	}
	if c.BTCScan.ConfirmScanSpec != "" {
		_, err := crontab.AddFunc(c.BTCScan.ConfirmScanSpec, btcScan.UpdateConfirmNumber)
		logx.Must(err)
		logx.Infof("add cron btc confirm scan spec: %v", c.BTCScan.ConfirmScanSpec)
	}
	if c.EvmScan.LogsScanSpec != "" {
		_, err := crontab.AddFunc(c.EvmScan.LogsScanSpec, evmScan.LogScan)
		logx.Must(err)
		logx.Infof("add cron evm logs scan spec: %v", c.EvmScan.LogsScanSpec)
	}
	serverGroup := service.NewServiceGroup()
	serverGroup.Add(server)
	serverGroup.Add(crontab)

	serverGroup.Start()
}
