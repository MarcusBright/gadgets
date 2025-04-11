/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	appparams "github.com/babylonlabs-io/babylon/app/params"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/cobra"
)

// newmnemonicCmd represents the newmnemonic command
var newmnemonicCmd = &cobra.Command{
	Use:   "newmnemonic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newmnemonic called")
		mnemonicEntropySize := 128
		entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
		if err != nil {
			fmt.Println("NewEntropy error:", err)
			return
		}

		mnemonic, err := bip39.NewMnemonic(entropySeed)
		if err != nil {
			fmt.Println("NewMnemonic error:", err)
			return
		}

		fmt.Println("New mnemonic:", mnemonic)

		hdPath := sdk.GetConfig().GetFullBIP44Path()
		encodingConfig := appparams.DefaultEncodingConfig()

		kr := keyring.NewInMemory(encodingConfig.Codec)
		keyInfo, err := kr.NewAccount("sender", mnemonic, keyring.DefaultBIP39Passphrase, hdPath, hd.Secp256k1)
		if err != nil {
			fmt.Println("NewAccount error:", err)
			return
		}
		senderAddr, err := keyInfo.GetAddress()
		if err != nil {
			fmt.Println("GetAddress error:", err)
		}
		fmt.Printf("Sender Adresse: %s\n", senderAddr.String())
	},
}

func init() {
	rootCmd.AddCommand(newmnemonicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newmnemonicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newmnemonicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
