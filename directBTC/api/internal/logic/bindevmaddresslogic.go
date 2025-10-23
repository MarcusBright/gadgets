// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"

	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/model"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	verifier "github.com/bitonicnl/verify-signed-message/pkg"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/zeromicro/go-zero/core/logx"
)

type BindEvmAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindEvmAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindEvmAddressLogic {
	return &BindEvmAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindEvmAddressLogic) BindEvmAddress(req *types.BindEvmAddressReq) (resp *types.BindEvmAddressResp, err error) {
	var message types.Message
	err = json.Unmarshal([]byte(req.Message), &message)
	if err != nil {
		return nil, fmt.Errorf("message unmarshal:%v", err)
	}
	btcTran, err := l.canBind(&message)
	if btcTran == nil || err != nil {
		l.Errorf("canBind,hash:%v, error: %v", message.TransactionHash, err)
		return nil, fmt.Errorf("not found or error:%v", err)
	}

	valid, err := l.btcSignVerify(btcTran, req.Message, message.SignAddress, req.Signature)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("sign not valid")
	}

	if err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(l.ctx).Model(&model.BtcTran{}).Where("transaction_hash = ?", message.TransactionHash).
			Updates(map[string]interface{}{
				"binded_evm_address": message.EvmAddress,
				"chain_id":           message.EvmChainId,
				"status":             model.BtcTranStatusBinded,
			}).Error; err != nil {
			l.Errorf("hash:%v, error: %v", message.TransactionHash, err)
			return err
		}
		signData := model.SignData{
			Message:     datatypes.JSON(req.Message),
			Signature:   req.Signature,
			SignType:    model.BTCSignTypeBindAddress,
			Signer:      message.SignAddress,
			BtcTranHash: message.TransactionHash,
		}
		if err := tx.WithContext(l.ctx).Model(&model.SignData{}).Create(&signData).Error; err != nil {
			l.Errorf("hash:%v, error: %v", message.TransactionHash, err)
			return err
		}
		return nil
	}); err != nil {
		l.Errorf("hash:%v, error: %v", message.TransactionHash, err)
		return nil, err
	}

	resp = &types.BindEvmAddressResp{
		Message: message,
	}

	return resp, nil
}

func (l *BindEvmAddressLogic) canBind(req *types.Message) (*model.BtcTran, error) {
	var btcTran model.BtcTran
	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BtcTran{}).Where("transaction_hash = ?", req.TransactionHash).First(&btcTran).Error; err != nil {
		l.Errorf("hash:%v, error: %v", req.TransactionHash, err)
		return nil, err
	}
	if slices.Contains([]string{model.BtcTranStatusApprovedInEvm, model.BtcTranStatusRecievedInEvm,
		model.BtcTranStatusRejectedInEvm}, btcTran.Status) {
		return nil, fmt.Errorf("transaction status not init/bind, status: %v", btcTran.Status)
	}
	return &btcTran, nil
}

func (l *BindEvmAddressLogic) btcSignVerify(btcTran *model.BtcTran, message, signAddress, signature string) (bool, error) {
	inputAddress := func() []string {
		var addrs []string
		_ = json.Unmarshal(btcTran.InputUtxo, &addrs)
		return addrs
	}()
	if !slices.Contains(inputAddress, signAddress) {
		return false, errors.New("signer address not in input utxo")
	}
	net := func() *chaincfg.Params {
		if l.svcCtx.Config.BTCScan.Network == "signet" {
			return &chaincfg.SigNetParams
		}
		return &chaincfg.MainNetParams
	}()
	return verifier.VerifyWithChain(verifier.SignedMessage{
		Address:   signAddress,
		Message:   message,
		Signature: signature,
	}, net)
}
