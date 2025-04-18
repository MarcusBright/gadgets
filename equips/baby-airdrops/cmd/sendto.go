/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"babylon/airdrop"
	"babylon/config"
	"babylon/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sendtoCmd represents the sendto command
var sendtoCmd = &cobra.Command{
	Use:   "sendto",
	Short: "Send tokens to a specific address",
	Long: `Transfer tokens to a specified address on the Babylon blockchain.

This command allows you to:
- Send tokens to a single address
- Specify the amount to send
- Confirm the transaction before execution

Example:
  babylon sendto -t <to_address> -a <amount> --config ./config.yaml
	amount must be in decimal format,etc 12.345678

The command will display the transaction hash upon successful completion.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sendto called")
		configPath := cmd.Flag("config").Value.String()
		config := config.LoadConfig(configPath)

		toaddress, err := cmd.Flags().GetString("toaddress")
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, err := cmd.Flags().GetString("amount")
		if err != nil {
			fmt.Println(err)
			return
		}
		//confirm
		fmt.Println("toaddress:", toaddress)
		fmt.Println("amount:", amount)
		fmt.Println("press YES to send")
		confirm := ""
		fmt.Scanln(&confirm)
		if confirm != "YES" {
			return
		}

		airDrop := airdrop.New(config, nil)
		txHash, err := airDrop.MultiSend([]models.BabySendAirDrop{
			{
				Address: toaddress,
				Amount:  amount,
			},
		})
		if err != nil {
			fmt.Println("sendto err:", err)
			return
		}
		err = airDrop.WaitTxMined(txHash)
		if err != nil {
			logrus.Errorf("WaitTxMined error:%v", err)
			return
		}
		fmt.Println("send OK, hash:", txHash)
	},
}

func init() {
	rootCmd.AddCommand(sendtoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendtoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendtoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sendtoCmd.Flags().StringP("toaddress", "t", "", "toaddress")
	sendtoCmd.Flags().StringP("amount", "a", "", "amount")
}
