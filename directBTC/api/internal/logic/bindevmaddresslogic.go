// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/model"
	"directBTC/pkg/slack"

	"gorm.io/gorm"

	"directBTC/service/contract/directbtcminter"
	evmscanconfig "directBTC/service/evmscan/config"

	verifier "github.com/bitonicnl/verify-signed-message/pkg"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeromicro/go-zero/core/logx"
)

type BindEvmAddressLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	trialLogic *GetBtcAddressIsTrialLogic
}

func NewBindEvmAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindEvmAddressLogic {
	return &BindEvmAddressLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		trialLogic: NewGetBtcAddressIsTrialLogic(ctx, svcCtx),
	}
}

func (l *BindEvmAddressLogic) BindEvmAddress(req *types.BindEvmAddressReq) (resp *types.BindEvmAddressResp, err error) {
	var message types.Message
	err = json.Unmarshal([]byte(req.Message), &message)
	if err != nil {
		return nil, fmt.Errorf("message unmarshal:%v", err)
	}

	btcTran, err := l.canBind(&message)
	if err != nil {
		return nil, err
	}

	valid, signer, signType, err := l.signVerify(req.Message, message.SignAddress, req.Signature, btcTran)
	if err != nil || !valid {
		return nil, err
	}

	if signType == "btcSign" { //btc
		if err := l.validateBtcTrialAndWhitelist(message, btcTran); err != nil {
			return nil, err
		}
	} else { //evm sys
		//check evm whitelist in contract
		if in := l.checkEvmInContract(message.EvmAddress, uint(message.EvmChainId)); !in {
			return nil, fmt.Errorf("not in whitelist")
		}
	}

	if err := l.persistBindData(message, signer, req); err != nil {
		return nil, err
	}

	resp = &types.BindEvmAddressResp{
		Message: message,
	}
	go slack.SendTo(l.svcCtx.Config.NotifySlack, "new bind, need check and recieve")
	return resp, nil
}

func (l *BindEvmAddressLogic) canBind(req *types.Message) (*model.BtcTran, error) {
	//check param
	if req.EvmAddress == "" || req.EvmChainId == 0 || req.SignAddress == "" || req.TransactionHash == "" {
		return nil, fmt.Errorf("message param error")
	}
	if !slices.ContainsFunc(l.svcCtx.Config.EvmScan.ChainInfo, func(c evmscanconfig.ChainInfo) bool {
		return req.EvmChainId == uint64(c.Client.ChainId)
	}) {
		return nil, fmt.Errorf("chainid not allowed")
	}

	var btcTran model.BtcTran
	if err := l.svcCtx.DB.WithContext(l.ctx).Model(&model.BtcTran{}).
		Where("transaction_hash = ?", req.TransactionHash).First(&btcTran).Error; err != nil {
		l.Errorf("hash:%v, error: %v", req.TransactionHash, err)
		return nil, err
	}
	if btcTran.Status != model.BtcTranStatusInit || btcTran.BindedEvmAddress != "" || btcTran.ChainId != 0 {
		return nil, fmt.Errorf("hash has binded:%v", btcTran.Status)
	}
	return &btcTran, nil
}

func (l *BindEvmAddressLogic) signVerify(message, signAddress, signature string, btcTran *model.BtcTran) (bool, string, string, error) {
	btcValid, err := l.btcSignVerify(message, signAddress, signature, btcTran)
	if err == nil && btcValid {
		return true, signAddress, "btcSign", nil
	}
	evmSiner, err := l.evmSignVerify(message, signature)
	if err == nil && evmSiner != "" {
		return true, evmSiner, "evmSign", nil
	}
	return false, "", "", fmt.Errorf("sign error")
}

func (l *BindEvmAddressLogic) btcSignVerify(message, signAddress, signature string, btcTran *model.BtcTran) (bool, error) {
	inputAddress := func() []string {
		var addrs []string
		_ = json.Unmarshal(btcTran.InputUtxo, &addrs)
		return addrs
	}()

	if inputAddress[0] != signAddress {
		l.Errorf("signer address not equal input[0]")
		return false, errors.New("signer address not equal input[0]")
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

func (l *BindEvmAddressLogic) evmSignVerify(message, sig string) (string, error) {
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	messageHash := crypto.Keccak256Hash([]byte(prefixedMessage))
	signature, err := hexutil.Decode(sig)
	if err != nil {
		return "", err
	}
	// Adjust the recovery ID (v) if needed (e.g., if v is 27/28)
	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}
	// 3. Recover the public key
	pubKeyBytes, err := crypto.Ecrecover(messageHash.Bytes(), signature)
	if err != nil {
		return "", err
	}
	publicKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return "", err
	}
	// 4. Derive the address and verify
	recoveredAddress := crypto.PubkeyToAddress(*publicKey)
	l.Infof("recoverAddress:%v", recoveredAddress)
	// check if is system
	if slices.Contains(l.svcCtx.Config.SysEvmAddress, recoveredAddress.String()) {
		return recoveredAddress.String(), nil
	}
	l.Errorf("not in system signer:%v", l.svcCtx.Config.SysEvmAddress)
	return "", fmt.Errorf("not system signer")
}

func (l *BindEvmAddressLogic) persistBindData(message types.Message, signer string, req *types.BindEvmAddressReq) error {
	if err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(l.ctx).Model(&model.BtcTran{}).
			Where("transaction_hash = ?", message.TransactionHash).
			Where("status = ?", model.BtcTranStatusInit).
			Updates(map[string]any{
				"binded_evm_address": message.EvmAddress,
				"chain_id":           message.EvmChainId,
				"status":             model.BtcTranStatusBinded,
			}).Error; err != nil {
			l.Errorf("message:%v, error: %v", message, err)
			return err
		}
		signData := model.BindEvmSign{
			Message:          req.Message,
			Signature:        req.Signature,
			Signer:           signer,
			BindedEvmAddress: message.EvmAddress,
			ChainId:          uint(message.EvmChainId),
			BtcAddress:       message.SignAddress,
			BtcTranHash:      message.TransactionHash,
		}
		if err := tx.WithContext(l.ctx).Model(&model.BindEvmSign{}).Create(&signData).Error; err != nil {
			l.Errorf("binded:%v", signData.BtcAddress)
			return err
		}
		return nil
	}); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("Duplicate entry")) {
			return errors.New("already binded")
		}
		return err
	}
	return nil
}

func (l *BindEvmAddressLogic) validateBtcTrialAndWhitelist(message types.Message, btcTran *model.BtcTran) error {
	amountSatoshi, err := strconv.ParseUint(btcTran.AmountSatoshi, 10, 64)
	if err != nil {
		return fmt.Errorf("amountSatoshi parse uint:%v", err)
	}
	feeSatoshi, err := strconv.ParseUint(btcTran.FeeSatoshi, 10, 64)
	if err != nil {
		return fmt.Errorf("feeSatoshi parse uint:%v", err)
	}

	trialResp, err := l.trialLogic.GetBtcAddressIsTrial(&types.GetBtcAddressIsTrialReq{
		Address: message.SignAddress,
	})
	if err != nil {
		return err
	}

	if amountSatoshi+feeSatoshi != l.svcCtx.Config.TinyTry {
		if !trialResp.TrialComplete || trialResp.TrialInfo == nil || trialResp.TrialInfo.BindInfo == nil {
			return fmt.Errorf("trial not complete")
		}
		if trialResp.TrialInfo.BindInfo.BindedEvmAddress != message.EvmAddress {
			return fmt.Errorf("evmAddress not the trial address")
		}
		if in := l.checkEvmInContract(message.EvmAddress, uint(message.EvmChainId)); !in {
			return fmt.Errorf("not in whitelist")
		}
	} else {
		if trialResp.TrialInfo == nil || trialResp.TrialInfo.BindInfo == nil {
			return fmt.Errorf("sys wrong")
		}
		if trialResp.TrialInfo.Hash != message.TransactionHash {
			if !trialResp.TrialComplete {
				return fmt.Errorf("trial not complete")
			}
			if trialResp.TrialInfo.BindInfo.BindedEvmAddress != message.EvmAddress {
				return fmt.Errorf("evmAddress not the trial address")
			}
			if in := l.checkEvmInContract(message.EvmAddress, uint(message.EvmChainId)); !in {
				return fmt.Errorf("not in whitelist")
			}
		}
	}
	return nil
}

func (l *BindEvmAddressLogic) checkEvmInContract(address string, chainId uint) bool {
	directBTCMinter, _ := directbtcminter.NewDirectbtcminter(common.HexToAddress(l.svcCtx.EvmClientsMap[chainId].Contract),
		l.svcCtx.EvmClientsMap[chainId].Client)
	in, err := directBTCMinter.Recipients(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		l.Errorf("call contract error:%v", err)
		return false
	}
	return in
}
