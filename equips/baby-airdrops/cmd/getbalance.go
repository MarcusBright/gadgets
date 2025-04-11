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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
