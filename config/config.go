package config

import (
	"os"

	"github.com/babylonlabs-io/babylon/app"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/yaml.v3"
)

type Config struct {
	RpcEndPoint   string `yaml:"rpcEndPoint"`
	Dsn           string `yaml:"dsn"`
	ChainId       string `yaml:"chainId"`
	Mnemonic      string `yaml:"mnemonic"`
	Address       string `yaml:"address"`
	BatchLimit    uint   `yaml:"batchLimit"`
	Fee           string `yaml:"fee"`
	NoOncePerTime bool   `yaml:"noOncePerTime"`
}

func LoadConfig(path string) (config *Config) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
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
	return
}
