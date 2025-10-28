package model

import (
	"directBTC/pkg/gormz"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/directbtc?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: gormz.NewGormLogger(),
	})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").
		AutoMigrate(&BtcTran{}, &BindEvmSign{}, &Cursor{})
	fmt.Printf("err: %v", err)
}
