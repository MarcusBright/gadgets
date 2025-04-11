package airdrop

import (
	"babylon/config"
	"babylon/models"
	"context"
	"fmt"
	"strings"
	"time"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/babylon/app"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AirDrop struct {
	Config        *config.Config
	DBEngine      *gorm.DB
	ClientContext *client.Context
	SenderAddr    types.AccAddress
	Quit          chan struct{}
}

func New(config *config.Config, quit chan struct{}) *AirDrop {
	airDrop := &AirDrop{Config: config, Quit: quit}

	var err error
	airDrop.DBEngine, err = gorm.Open(postgres.Open(config.Dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(fmt.Sprintf("Open Mysql err:%v", err))
	}
	err = airDrop.DBEngine.AutoMigrate(&models.BabySendAirDrop{})
	if err != nil {
		panic(fmt.Sprintf("AutoMigrate err:%v", err))
	}

	hdPath := sdk.GetConfig().GetFullBIP44Path()
	encodingConfig := app.GetEncodingConfig()

	kr := keyring.NewInMemory(encodingConfig.Codec)
	keyInfo, err := kr.NewAccount("sender", config.Mnemonic, keyring.DefaultBIP39Passphrase, hdPath, hd.Secp256k1)
	if err != nil {
		panic(err)
	}
	senderAddr, err := keyInfo.GetAddress()
	if err != nil {
		panic(err)
	}
	if senderAddr.String() != config.Address {
		panic("mnemonic, address not match")
	}
	rpcClient, err := rpchttp.New(config.RpcEndPoint, "/websocket")
	if err != nil {
		panic(err)
	}
	airDrop.SenderAddr = senderAddr
	clientCtx := client.Context{}.WithClient(rpcClient).
		WithChainID(config.ChainId).
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithNodeURI(config.RpcEndPoint).
		WithClient(rpcClient).
		WithBroadcastMode(flags.BroadcastSync).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithKeyring(kr)
	airDrop.ClientContext = &clientCtx
	return airDrop
}

func (airDrop *AirDrop) Run() {
	for {
		select {
		case _, b := <-airDrop.Quit:
			if !b {
				logrus.Infoln("Get close chain Quit scan")
				return
			}
		default:
		}

		var dropItems []models.BabySendAirDrop
		rest := airDrop.DBEngine.Where("hash = ?", "").Where("confirm = ?", 0).
			Where("send_serial = ?", airDrop.Config.SendSerial).Order("id asc").
			Limit(int(airDrop.Config.BatchLimit)).Find(&dropItems)
		if rest.Error != nil {
			panic(rest.Error)
		}
		if len(dropItems) == 0 {
			logrus.Infoln("No need to send")
			return
		}
		logrus.Infof("MultiSend len(%d)", len(dropItems))
		txHash, err := airDrop.MultiSend(dropItems)
		if err != nil {
			logrus.Errorf("MultiSend error:%v", err)
			return
		}
		logrus.Infof("MultiSend success:%s", txHash)
		var addresses []string
		for _, v := range dropItems {
			addresses = append(addresses, v.Address)
		}
		// update hash
		rest = airDrop.DBEngine.Model(&models.BabySendAirDrop{}).Where("address in (?)", addresses).
			Where("send_serial = ?", airDrop.Config.SendSerial).
			Updates(map[string]interface{}{"hash": txHash})
		if rest.Error != nil {
			panic(rest.Error)
		}
		//wait
		err = airDrop.WaitTxMined(txHash)
		if err != nil {
			logrus.Errorf("WaitTxMined error:%v", err)
			return
		}
		// update confirm
		rest = airDrop.DBEngine.Model(&models.BabySendAirDrop{}).Where("address in (?)", addresses).
			Where("send_serial = ?", airDrop.Config.SendSerial).
			Updates(map[string]interface{}{"confirm": 1})
		if rest.Error != nil {
			panic(rest.Error)
		}
		logrus.Infof("MultiSend success:%s, len(%d)", txHash, len(addresses))
		if airDrop.Config.NoOncePerTime {
			logrus.Info("NoOncePerTime set continue")
		} else {
			logrus.Info("NoOncePerTime false")
			return
		}
	}
}

func (airDrop *AirDrop) WaitTxMined(hash string) error {
	for i := 0; i < 10; i++ {
		time.Sleep(7 * time.Second)
		txResp, err := authtx.QueryTx(*airDrop.ClientContext, hash)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				logrus.Warnf("tx not mined:%s, wait...", hash)
				continue
			} else {
				return err
			}
		}
		if txResp.Code != 0 {
			logrus.Errorf("tx error:%d", txResp.Code)
			return fmt.Errorf("tx error:%d", txResp.Code)
		} else {
			logrus.Infof("tx mined:%s", hash)
			return nil
		}
	}
	return fmt.Errorf("tx not mined")
}

func (airDrop *AirDrop) MultiSend(items []models.BabySendAirDrop) (string, error) {
	var outputs []banktypes.Output
	totalSent := sdk.NewCoin("ubbn", math.NewInt(0))
	for _, v := range items {
		amount := decimal.RequireFromString(v.Amount)
		adenom := amount.Mul(decimal.New(1, 6))
		coin := sdk.NewCoin("ubbn", math.NewIntFromBigInt(adenom.BigInt()))
		if coin.IsZero() {
			logrus.Errorf("coin is zero:%s", v.Address)
			return "", fmt.Errorf("coin is zero")
		}
		recipient, err := sdk.AccAddressFromBech32(strings.TrimSpace(v.Address))
		if err != nil {
			logrus.Errorf("AccAddressFromBech32 error:%s", v.Address)
			return "", err
		}
		output := banktypes.NewOutput(recipient, sdk.NewCoins(coin))
		outputs = append(outputs, output)
		totalSent = totalSent.Add(coin)
	}
	if totalSent.IsZero() {
		logrus.Errorf("totalSent is zero")
		return "", fmt.Errorf("totalSent is zero")
	}
	balance, err := airDrop.GetBalance(airDrop.SenderAddr.String())
	if err != nil {
		logrus.Errorf("Balance error:%v", err)
		return "", err
	}
	logrus.Infof("balance:%s, totalSent:%s", balance.String(), totalSent.String())
	if balance.IsLTE(totalSent) {
		return "", fmt.Errorf("balance is not enough")
	}

	input := banktypes.NewInput(airDrop.SenderAddr, sdk.NewCoins(totalSent))
	msg := banktypes.NewMsgMultiSend(input, outputs)

	txFactory := tx.Factory{}.
		WithChainID(airDrop.ClientContext.ChainID).
		WithTxConfig(airDrop.ClientContext.TxConfig).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
		WithGasAdjustment(1.5)
	num, seq, err := airDrop.ClientContext.AccountRetriever.GetAccountNumberSequence(*airDrop.ClientContext, airDrop.SenderAddr)
	if err != nil {
		logrus.Errorf("GetAccountNumberSequence error:%v", err)
		return "", err
	}
	txFactory = txFactory.WithAccountNumber(num).WithSequence(seq)
	sim, adjusted, err := tx.CalculateGas(airDrop.ClientContext, txFactory, msg)
	if err != nil {
		logrus.Errorf("CalculateGas error:%v", err)
		return "", err
	}
	txFactory = txFactory.WithGas(adjusted).WithKeybase(airDrop.ClientContext.Keyring).WithFees(airDrop.Config.Fee).WithMemo(airDrop.Config.Memo)
	logrus.Infof("fee:%v, gasLimit:%v, sim used:%v, sim wanted: %v, gasPrice:%v", txFactory.Fees(), txFactory.Gas(), sim.GasInfo.GasUsed, sim.GasInfo.GasWanted, txFactory.GasPrices())
	txBuilder, err := txFactory.BuildUnsignedTx(msg)
	if err != nil {
		logrus.Errorf("BuildUnsignedTx error:%v", err)
		return "", err
	}
	err = tx.Sign(context.Background(), txFactory, "sender", txBuilder, true)
	if err != nil {
		logrus.Errorf("Sign error:%v", err)
		return "", err
	}
	txBytes, err := airDrop.ClientContext.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		logrus.Errorf("TxEncoder error:%v", err)
		return "", err
	}
	broadCastTx, err := airDrop.ClientContext.BroadcastTx(txBytes)
	if err != nil {
		logrus.Errorf("BroadcastTx error:%v", err)
		return "", err
	}
	if broadCastTx.Code != 0 {
		logrus.Errorf("tx error:%d", broadCastTx.Code)
		return "", fmt.Errorf("tx error:%d", broadCastTx.Code)
	}
	return broadCastTx.TxHash, nil
}

func (airDrop *AirDrop) GetBalance(address string) (sdk.Coin, error) {
	bankQueryClient := banktypes.NewQueryClient(airDrop.ClientContext)
	coin := sdk.Coin{}
	queryReq := &banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   "ubbn",
	}
	queryResp, err := bankQueryClient.Balance(context.Background(), queryReq)
	if err != nil {
		return coin, err
	}
	coin = sdk.NewCoin("ubbn", queryResp.Balance.Amount)
	return coin, nil
}
