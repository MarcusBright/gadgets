package airdrop

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/babylon/app"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/shopspring/decimal"
)

func TestMultiSend(t *testing.T) {
	mnemonic := ""
	rpcEndPoint := ""
	recipient0 := "bbn137umamrx66a590lh8neep05wlwqkhuw6u6959d"
	recipient0Amount := "0.001"
	recipient1 := "bbn1y6037muljpczznwg0q7jju9jh2hvv8q0jzyrm6"
	recipient1Amount := "0.002"
	recipient2 := "bbn1j2t9f3vencxyvt6knyu8e7srfzfeesdtedssle"
	recipient2Amount := "0.003"

	hdPath := sdk.GetConfig().GetFullBIP44Path()
	encodingConfig := app.GetEncodingConfig()

	kr := keyring.NewInMemory(encodingConfig.Codec)
	keyInfo, err := kr.NewAccount("sender", mnemonic, keyring.DefaultBIP39Passphrase, hdPath, hd.Secp256k1)
	if err != nil {
		fmt.Println("NewAccount error:", err)
		return
	}
	senderAddr, err := keyInfo.GetAddress()
	fmt.Println(senderAddr.String())
	rpcClient, err := rpchttp.New(rpcEndPoint, "/websocket")
	if err != nil {
		fmt.Println("NewWithTimeout error:", err)
		return
	}
	clientCtx := client.Context{}.WithClient(rpcClient).
		WithChainID("bbn-test-5").
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithNodeURI(rpcEndPoint).
		WithClient(rpcClient).
		WithBroadcastMode(flags.BroadcastSync).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithKeyring(kr)

	var outputs []banktypes.Output
	var totalSent sdk.Coins

	parsedAmount0 := sdk.NewCoin("ubbn", math.NewIntFromBigInt(decimal.RequireFromString(recipient0Amount).Mul(decimal.New(1, 6)).BigInt()))
	parsedAmount1 := sdk.NewCoin("ubbn", math.NewIntFromBigInt(decimal.RequireFromString(recipient1Amount).Mul(decimal.New(1, 6)).BigInt()))
	parsedAmount2 := sdk.NewCoin("ubbn", math.NewIntFromBigInt(decimal.RequireFromString(recipient2Amount).Mul(decimal.New(1, 6)).BigInt()))

	recipientAddr0, err := sdk.AccAddressFromBech32(strings.TrimSpace(recipient0))
	if err != nil {
		fmt.Println("AccBech32 err:", err)
		return
	}
	recipientAddr1, err := sdk.AccAddressFromBech32(strings.TrimSpace(recipient1))
	if err != nil {
		fmt.Println("AccBech32 err:", err)
		return
	}
	recipientAddr2, err := sdk.AccAddressFromBech32(strings.TrimSpace(recipient2))
	if err != nil {
		fmt.Println("AccBech32 err:", err)
		return
	}
	output0 := banktypes.NewOutput(recipientAddr0, sdk.NewCoins(parsedAmount0))
	output1 := banktypes.NewOutput(recipientAddr1, sdk.NewCoins(parsedAmount1))
	output2 := banktypes.NewOutput(recipientAddr2, sdk.NewCoins(parsedAmount2))
	outputs = append(outputs, output0, output1, output2)
	totalSent = totalSent.Add(parsedAmount0, parsedAmount1, parsedAmount2)
	fmt.Println("totalSend:", totalSent)
	input := banktypes.NewInput(senderAddr, totalSent)
	msg := banktypes.NewMsgMultiSend(input, outputs)

	txFactory := tx.Factory{}.
		WithChainID(clientCtx.ChainID).
		WithTxConfig(clientCtx.TxConfig).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
		WithGasAdjustment(1.5)

	// txFactory = txFactory.WithGasPrices("0.025ubbn")

	num, seq, err := clientCtx.AccountRetriever.GetAccountNumberSequence(clientCtx, senderAddr)
	if err != nil {
		fmt.Println("getAcccountNumberSe:", err)
		return
	}
	txFactory = txFactory.WithAccountNumber(num).WithSequence(seq)
	sim, adjusted, err := tx.CalculateGas(clientCtx, txFactory, msg)
	if err != nil {
		fmt.Println("calculate:", err)
		return
	}
	txFactory = txFactory.WithGas(adjusted).WithKeybase(clientCtx.Keyring).WithFees("300ubbn")
	fmt.Println("gasLimit:", txFactory.Gas())
	fmt.Println("sim used:", sim.GasInfo.GasUsed)
	fmt.Println("sim wanted:", sim.GasInfo.GasWanted)
	fmt.Println("gasPrice:", txFactory.GasPrices())
	fmt.Println("fee:", txFactory.Fees())
	txBuilder, err := txFactory.BuildUnsignedTx(msg)
	if err != nil {
		fmt.Println("Build UnsigedTx:", err)
		return
	}
	err = tx.Sign(context.Background(), txFactory, keyInfo.Name, txBuilder, true)
	if err != nil {
		fmt.Println("sign err:", err)
		return
	}
	txBytes, err := clientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
	broadCastTx, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		fmt.Println("BroadCastTx error:", err)
		return
	}
	fmt.Println("txHash:", broadCastTx.TxHash)
	if broadCastTx.Code != 0 {
		fmt.Println("tx error:", broadCastTx.Code)
		fmt.Println("name:", broadCastTx.Codespace)
		fmt.Println("gaswanted:", broadCastTx.GasWanted)
		fmt.Println("gasUsed:", broadCastTx.GasUsed)
		fmt.Println("logs:", broadCastTx.Logs)
		// fmt.Printf("gaswanted:%s.logs:%s", broadCastTx.GasWanted, broadCastTx.Logs)
		return
	}
	fmt.Println("tx success:", broadCastTx.Code)
}

func TestParsCoin(t *testing.T) {
	// amount := "0.00230013"
	// amount := "123243253.00230013"
	// amount := "123243253.0023"
	// amount := "0.665579063"
	amount := "2040.06264"
	a := decimal.RequireFromString(amount)
	t.Log(a)
	adenom := a.Mul(decimal.New(1, 6))
	coin := sdk.NewCoin("ubbn", math.NewIntFromBigInt(adenom.BigInt()))
	t.Log(coin)
}

func TestBalance(t *testing.T) {
	rpcEndPoint := "http://51.89.40.26:26657"
	rpcClient, err := rpchttp.New(rpcEndPoint, "/websocket")
	if err != nil {
		fmt.Println("NewWithTimeout error:", err)
		return
	}
	encodingConfig := app.GetEncodingConfig()
	clientCtx := client.Context{}.WithClient(rpcClient).
		WithChainID("bbn-test-5").
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithNodeURI(rpcEndPoint).
		WithClient(rpcClient).
		WithBroadcastMode(flags.BroadcastSync).
		WithAccountRetriever(authtypes.AccountRetriever{})

	recipientAddr2, err := sdk.AccAddressFromBech32(strings.TrimSpace("bbn1ppxykaduglgrgmt6jmg64kfxf9mawqxjf0h2pr"))
	if err != nil {
		fmt.Println("AccBech32 err:", err)
		return
	}
	bankQueryClient := banktypes.NewQueryClient(clientCtx)

	// 4. Abfrage vorbereiten: Alle Salden für die Adresse abrufen
	queryReq := &banktypes.QueryAllBalancesRequest{
		Address: recipientAddr2.String(),
		// Pagination kann hier hinzugefügt werden, wenn die Adresse viele verschiedene Token hält
		Pagination: nil,
	}
	queryResp, err := bankQueryClient.AllBalances(context.Background(), queryReq)
	if err != nil {
		t.Log("err:", err)
		return
	}
	t.Log(queryResp)
}

func TestGetHash(t *testing.T) {
	// hash := "CE31A60759069BA4264B5E35D5720FC039DE73F5334B26953501007BD62250A1"
	// hash := "049E22A52FF1F82C5AADF89FC355F0EDF23FB650002D5E3F100210385517228B"
	hash := "E31A60759069BA4264B5E35D5720FC039DE73F5334B26953501007BD62250A1A"
	rpcEndPoint := "http://51.89.40.26:26657"
	rpcClient, err := rpchttp.New(rpcEndPoint, "/websocket")
	if err != nil {
		fmt.Println("NewWithTimeout error:", err)
		return
	}
	encodingConfig := app.GetEncodingConfig()
	clientCtx := client.Context{}.WithClient(rpcClient).
		WithChainID("bbn-test-5").
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithNodeURI(rpcEndPoint).
		WithClient(rpcClient).
		WithBroadcastMode(flags.BroadcastSync).
		WithAccountRetriever(authtypes.AccountRetriever{})
	txResp, err := authtx.QueryTx(clientCtx, hash)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			fmt.Println("not found")
		}
		return
	}
	t.Log(txResp)
}
