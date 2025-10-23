// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"directBTC/api/internal/config"
	"directBTC/model"
	"directBTC/pkg/gormz"

	"github.com/glebarez/sqlite"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	MemDB  *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: gormz.NewGormLogger(),
	})
	logx.Must(err)

	memoryDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: gormz.NewGormLogger(),
	})
	logx.Must(err)
	logx.Must(memoryDB.AutoMigrate(&model.BtcTran{}))

	return &ServiceContext{
		Config: c,
		DB:     db,
		MemDB:  memoryDB,
	}
}
