/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"babylon/airdrop"
	"babylon/config"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

// getbalanceCmd represents the getbalance command
var getbalanceCmd = &cobra.Command{
	Use:   "getbalance",
	Short: "Query account balance on Babylon blockchain",
	Long: `Check the token balance of an address on the Babylon blockchain.

This command will:
- Connect to the configured Babylon network
- Query the balance of the specified address
- Display the balance in both raw and decimal format

Example:
  babylon getbalance --config ./config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getbalance called")
		configPath := cmd.Flag("config").Value.String()
		config := config.LoadConfig(configPath)
		airdrop := airdrop.New(config, nil)
		balance, err := airdrop.GetBalance(airdrop.SenderAddr.String())
		if err != nil {
			fmt.Println("GetBalance error:", err)
			return
		}
		fmt.Println("Address:", airdrop.SenderAddr.String())
		fmt.Println("GetBalance:", balance)
		fmt.Println("GetBalance:", decimal.New(balance.Amount.Int64(), -6))
	},
}

func init() {
	rootCmd.AddCommand(getbalanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getbalanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getbalanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
