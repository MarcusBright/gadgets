// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"directBTC/api/internal/config"
	"directBTC/api/internal/svc"
	"directBTC/api/internal/types"
	"directBTC/model"
	"directBTC/pkg/webpki.org/jsoncanonicalizer"
	"encoding/json"
	"testing"

	verifier "github.com/bitonicnl/verify-signed-message/pkg"
)

func TestVerifySignError(t *testing.T) {
	svc := svc.NewServiceContext(config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
	})
	l := NewBindEvmAddressLogic(context.Background(), svc)
	req := &types.BindEvmAddressReq{
		SignData: types.BindEvmAddressSignData{
			TransactionHash: "27bcb9009c68b677e45700f035e9f2b79b1d57cba1b35a45d07234839dae446a",
			EvmAddress:      "0xE343aB9eEfB1d3991F4CC3d10238225b56526EF4",
			EvmChainId:      1,
			SignAddress:     "1GYBGRrWBnAFWEPMGXyo1UqpE5SH2pb6bX",
			SignType:        "btc",
		},
		Signature: "HyiLDcQQ1p2bKmyqM0e5oIBQtKSZds4kJQ+VbZWpr0kYA6Qkam2MlUeTr+lm1teUGHuLapfa43JjyrRqdSA0pxs=",
	}
	resp, err := l.BindEvmAddress(req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestVerifySignNoHash(t *testing.T) {
	svc := svc.NewServiceContext(config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
	})
	l := NewBindEvmAddressLogic(context.Background(), svc)
	req := &types.BindEvmAddressReq{
		SignData: types.BindEvmAddressSignData{
			TransactionHash: "209c68b677e45700f035e9f2b79b1d57cba1b35a45d07234839dae446a",
			EvmAddress:      "0xE343aB9eEfB1d3991F4CC3d10238225b56526EF4",
			EvmChainId:      1,
			SignAddress:     "1GYBGRrWBnAFWEPMGXyo1UqpE5SH2pb6bX",
			SignType:        "btc",
		},
		Signature: "HyiLDcQQ1p2bKmyqM0e5oIBQtKSZds4kJQ+VbZWpr0kYA6Qkam2MlUeTr+lm1teUGHuLapfa43JjyrRqdSA0pxs=",
	}
	resp, err := l.BindEvmAddress(req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}
func TestVerifySignApproved(t *testing.T) {
	svc := svc.NewServiceContext(config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
	})
	l := NewBindEvmAddressLogic(context.Background(), svc)
	req := &types.BindEvmAddressReq{
		SignData: types.BindEvmAddressSignData{
			TransactionHash: "3f0f56ae7a206e388923fb568e251289c91b38274939c390c5e075f785bb77b8",
			EvmAddress:      "0xE343aB9eEfB1d3991F4CC3d10238225b56526EF4",
			EvmChainId:      1,
			SignAddress:     "1GYBGRrWBnAFWEPMGXyo1UqpE5SH2pb6bX",
			SignType:        "btc",
		},
		Signature: "HyiLDcQQ1p2bKmyqM0e5oIBQtKSZds4kJQ+VbZWpr0kYA6Qkam2MlUeTr+lm1teUGHuLapfa43JjyrRqdSA0pxs=",
	}
	resp, err := l.BindEvmAddress(req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}
func TestVerifySignNotUtxo(t *testing.T) {
	svc := svc.NewServiceContext(config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
	})
	l := NewBindEvmAddressLogic(context.Background(), svc)
	req := &types.BindEvmAddressReq{
		SignData: types.BindEvmAddressSignData{
			TransactionHash: "11425054932c008c2ab2b7c198e3bbf995ae2c995ce0398fca5955a17e413630",
			EvmAddress:      "0xE343aB9eEfB1d3991F4CC3d10238225b56526EF4",
			EvmChainId:      1,
			SignAddress:     "1GYBGRrWBnAFWEPMGXyo1UqpE5SH2pb6bX",
			SignType:        "btc",
		},
		Signature: "HyiLDcQQ1p2bKmyqM0e5oIBQtKSZds4kJQ+VbZWpr0kYA6Qkam2MlUeTr+lm1teUGHuLapfa43JjyrRqdSA0pxs=",
	}
	resp, err := l.BindEvmAddress(req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

/*
Private Key (HEX): 850cece14ffefdb864f6007718a5243dae9194841617c7d6d77b67482d40d856
Private Key (WIF): L1gLtHEKG4FbbxQDzth3ksCZ4jTSjRvcU7K2KDeDE368pG8MjkFg
Public Key (Raw): (X=691ab7d2b2e1b41a8df334a5471a3abd7a93c8822b2abf3de64c552147dc33b8, Y=b1eed621c6b9e790a901ca30eb55ee95d591c3e6dc2e6aa30f2b9f5c525e7e32)
Public Key (HEX Compressed): 02691ab7d2b2e1b41a8df334a5471a3abd7a93c8822b2abf3de64c552147dc33b8
Legacy Address: 1N3kZRUrEioGxXQbSyCWuBwmoFp4T62i93
Nested SegWit Address: 3KWsrxLMHPU1v8riptj33zCsWD8bf6jfLF
Native SegWit Address: bc1qum0at29ayuq2ndk39z4zwf4zdpxv5ker570ape
Taproot Address: bc1p5utaw0g77graev5yw575c3jnzh8j88ezzw39lgr250ghppwpyccsvjkvyp
----
bmt sign -p -a legacy -d -m '{"evmAddress":"0xE343aB9eEfB1d3991F4CC3d10238225b56526EF4","evmChainId":1,"signAddress":"1N3kZRUrEioGxXQbSyCWuBwmoFp4T62i93","signType":"btc","transactionHash":"1232456ae7a206e388923fb568e251289c91b38274939c390c5e075f785bb77b8"}'
*/
func TestVerifySignOK(t *testing.T) {
	svc := svc.NewServiceContext(config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
	})
	l := NewBindEvmAddressLogic(context.Background(), svc)
	req := &types.BindEvmAddressReq{
		SignData: types.BindEvmAddressSignData{
			TransactionHash: "1232456ae7a206e388923fb568e251289c91b38274939c390c5e075f785bb77b8",
			EvmAddress:      "0xE343aB9eEfB1d3991F4CC3d10238225b56526EF4",
			EvmChainId:      1,
			SignAddress:     "1N3kZRUrEioGxXQbSyCWuBwmoFp4T62i93",
			SignType:        "btc",
		},
		Signature: "IFoYpbJeWuaJewecxS6TEhA6FHeaRqqmDg8BSJWisEIvIhO4kCoZEmArymrGF8H8gareOZfUgVXzYuZdDdRzTiE=",
	}
	signDataJson, _ := json.Marshal(req.SignData)
	message, _ := jsoncanonicalizer.Transform(signDataJson)
	t.Logf("message:%s", string(message))
	resp, err := l.BindEvmAddress(req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestVerifyData(t *testing.T) {
	svc := svc.NewServiceContext(config.Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local",
	})
	var signData []model.SignData
	svc.DB.WithContext(context.Background()).Model(&model.SignData{}).Find(&signData)
	for _, item := range signData {
		signDataJson, _ := json.Marshal(item.Message)
		message, _ := jsoncanonicalizer.Transform(signDataJson)
		verified, err := verifier.Verify(verifier.SignedMessage{
			Address:   item.Signer,
			Message:   string(message),
			Signature: item.Signature,
		})
		if err != nil {
			t.Log(err)
			return
		}
		t.Log(verified)
	}
}
