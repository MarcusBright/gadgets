package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	BtcTranStatusInit          = "init"
	BtcTranStatusBinded        = "binded"
	BtcTranStatusRecievedInEvm = "receivedInEvm"
	BtcTranStatusApprovedInEvm = "approvedInEvm"
	BtcTranStatusRejectedInEvm = "rejectedInEvm"
)

type BtcTran struct {
	gorm.Model
	TransactionHash  string         `gorm:"size:255;default:'';index:hash,unique"`
	TreasuryAddress  datatypes.JSON `gorm:"not null"` //[]string, address
	AmountSatoshi    string         `gorm:"default:'0'"`
	FeeSatoshi       string         `gorm:"default:'0'"`
	InputUtxo        datatypes.JSON `gorm:"not null"` //[]string, address
	Status           string         `gorm:"default:'init'"`
	BlockNumber      uint64         `gorm:"default:0"`
	BlockTime        uint64         `gorm:"default:0"`
	ConfirmNumber    uint64         `gorm:"default:0"`
	ConfirmThreshold uint64         `gorm:"default:0"`
	// Evm info
	BindedEvmAddress  string `gorm:"size:255;default:''"`
	ChainId           uint   `gorm:"default:0"`
	RecievedEvmTxHash string `gorm:"size:255;default:''"`
	AcceptedEvmTxHash string `gorm:"size:255;default:''"`
	RejectedEvmTxHash string `gorm:"size:255;default:''"`
	ProcessIdx        uint64 `gorm:"default:0"` //when recievedEvm
	Signature         string `gorm:"size:255;default:''"`
}

type BindEvmSign struct {
	gorm.Model
	Message          string `gorm:"not null;"`
	Signature        string `gorm:"size:255;default:''"`
	Signer           string `gorm:"size:255;default:''"`
	BtcAddress       string `gorm:"size:255;default:''"`
	ChainId          uint   `gorm:"default:0"`
	BindedEvmAddress string `gorm:"size:255;default:''"`
	BtcTranHash      string `gorm:"size:255;default:'';index:hash,unique"`
}

const (
	CursorEvmLogScan = "evmLogScan"
)

type Cursor struct {
	gorm.Model
	IsBtc bool `gorm:"default:false;index:t_chainid_address,unique"`
	// ChainId used for evm chain
	ChainId uint   `gorm:"default:0;index:t_chainid_address"`
	Address string `gorm:"size:255;index:t_chainid_address"`
	Txhash  string `gorm:"size:255"`
	// BlockNumber when we scan from_blockNUmber to to_blockNumber, has processed
	BlockNumber uint64 `gorm:"default:0"`
}

func (c *Cursor) GetAllBtcAddress(db *gorm.DB) ([]Cursor, error) {
	var cursors []Cursor
	err := db.Where("is_btc = ?", true).Find(&cursors).Error
	if err != nil {
		return nil, err
	}
	return cursors, nil
}
