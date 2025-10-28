// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"

	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/model"

	"gorm.io/gorm"

	evmscanconfig "directBTC/service/evmscan/config"

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
	//check param
	if message.EvmAddress == "" || message.EvmChainId == 0 || message.SignAddress == "" {
		return nil, fmt.Errorf("message param error")
	}
	if !slices.ContainsFunc(l.svcCtx.Config.EvmScan.ChainInfo, func(c evmscanconfig.ChainInfo) bool {
		return message.EvmChainId == uint64(c.Client.ChainId)
	}) {
		return nil, fmt.Errorf("chainid not allowed")
	}

	valid, err := l.btcSignVerify(req.Message, message.SignAddress, req.Signature)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("sign not valid")
	}

	if err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(l.ctx).Model(&model.BtcTran{}).
			//5.7+
			Where("JSON_EXTRACT(input_utxo, '$[0]') = ?", message.SignAddress).
			Where("status = ?", model.BtcTranStatusInit).
			Updates(map[string]interface{}{
				"binded_evm_address": message.EvmAddress,
				"chain_id":           message.EvmChainId,
				"status":             model.BtcTranStatusBinded,
			}).Error; err != nil {
			l.Errorf("message:%v, error: %v", message, err)
			return err
		} //only one evmChain, multi evm chain need bind
		signData := model.BindEvmSign{
			Message:          req.Message,
			Signature:        req.Signature,
			Signer:           message.SignAddress,
			BindedEvmAddress: message.EvmAddress,
			ChainId:          uint(message.EvmChainId),
			BtcAddress:       message.SignAddress,
		}
		if err := tx.WithContext(l.ctx).Model(&model.BindEvmSign{}).Create(&signData).Error; err != nil {
			l.Infof("binded:%v", signData.BtcAddress)
			return err
		}
		return nil
	}); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("Duplicate entry")) {
			return nil, errors.New("already binded")
		}
		return nil, err
	}

	resp = &types.BindEvmAddressResp{
		Message: message,
	}

	return resp, nil
}

func (l *BindEvmAddressLogic) btcSignVerify(message, signAddress, signature string) (bool, error) {
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
