// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"directBTC/api/internal/config"
	"directBTC/api/internal/middleware"
	"directBTC/clientrpc/goeth"
	"directBTC/pkg/gormz"

	evmconfig "directBTC/service/evmscan/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config                config.Config
	DB                    *gorm.DB
	EvmClientsMap         map[uint]*ChainIdClient
	DefaultJsonMiddleware rest.Middleware
	// MemDB  *gorm.DB
}

type ChainIdClient struct {
	ChainId  uint
	Client   *ethclient.Client
	Contract string
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource))
	if c.SqlLog {
		db.Logger = gormz.NewGormLogger()
	}
	logx.Must(err)

	evmClients := lo.Map(c.EvmScan.ChainInfo, func(item evmconfig.ChainInfo, index int) *ChainIdClient {
		return &ChainIdClient{item.Client.ChainId, goeth.NewClient(item.Client.Host, item.Client.Request, item.Client.PeriodSec), item.Miner}
	})
	evmClientsMap := lo.SliceToMap(evmClients, func(item *ChainIdClient) (uint, *ChainIdClient) {
		return item.ChainId, item
	})
	// memoryDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
	// 	Logger: gormz.NewGormLogger(),
	// })
	// logx.Must(err)
	// logx.Must(memoryDB.AutoMigrate(&model.BtcTran{}))

	return &ServiceContext{
		Config: c,
		DB:     db,
		// MemDB:  memoryDB,
		EvmClientsMap:         evmClientsMap,
		DefaultJsonMiddleware: middleware.NewDefaultJsonMiddleware().Handle,
	}
}
