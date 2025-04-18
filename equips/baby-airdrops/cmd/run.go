/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"babylon/airdrop"
	"babylon/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute batch airdrop operations",
	Long: `Run batch airdrop operations according to the configuration file.

This command reads the airdrop configuration file and executes token transfers
to multiple addresses in batch. It requires a config file that specifies:
- Sender's mnemonic
- Network configuration
- List of recipient addresses and amounts, which read from database.

Example:
  babylon run --config ./config.yaml

The command will ask for confirmation before executing the transfers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		configPath := cmd.Flag("config").Value.String()
		config := config.LoadConfig(configPath)

		sigs := make(chan os.Signal, 1)
		quit := make(chan struct{}, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
		go func() {
			for sig := range sigs {
				switch sig {
				case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP:
					logrus.Infoln("Get Signal", sig)
					close(quit)
					return
				default:
					logrus.Infoln("Get Signal ", sig)
				}
			}
		}()
		airdrop := airdrop.New(config, quit)
		fmt.Println("press YES to send")
		confirm := ""
		fmt.Scanln(&confirm)
		if confirm != "YES" {
			return
		}
		airdrop.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
