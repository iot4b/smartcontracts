/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package device

import (
	"fmt"
	"smartcontracts/cmd"

	"github.com/spf13/cobra"
)

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Commands for device",
	Long: `Commands:
	new - deploy smart contract for new Device in Blockchain`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("device called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(deviceCmd)
}
