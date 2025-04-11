package models

import (
	"gorm.io/gorm"
)

type BabySendAirDrop struct {
	gorm.Model
	Address    string `gorm:"index:idx_address_sn,unique;type:varchar(128);not null;default:''" json:"address"`
	Amount     string `gorm:"type:varchar(128);not null;default:''" json:"amount"`
	SendSerial string `gorm:"index:idx_address_sn,unique;type:varchar(128);not null;default:''" json:"sendSerial"`
	Hash       string `gorm:"type:varchar(128);not null;default:''" json:"hash"`
	Confirm    int    `gorm:"type:int;not null;default:0" json:"confirm"`
}
