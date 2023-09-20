/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package node

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new {public} {secret}",
	Short: "Deploy new smart contract for Node",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

func init() {
	nodeCmd.AddCommand(newCmd)
}
