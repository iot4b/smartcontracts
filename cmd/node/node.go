/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package node

import (
	"fmt"
	"smartcontracts/cmd"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(nodeCmd)
}
