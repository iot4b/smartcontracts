package elector

import (
	"github.com/spf13/cobra"
	"smartcontracts/cmd"
)

// electorCmd represents the device command
var electorCmd = &cobra.Command{
	Use:   "elector",
	Short: "Commands for elector",
}

func init() {
	cmd.RootCmd.AddCommand(electorCmd)
}
