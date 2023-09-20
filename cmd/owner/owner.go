package owner

import (
	"github.com/spf13/cobra"
	"smartcontracts/cmd"
)

// ownerCmd represents the device command
var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Commands for owner",
}

func init() {
	cmd.RootCmd.AddCommand(ownerCmd)
}
