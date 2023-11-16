package vendor

import (
	"github.com/spf13/cobra"
	"smartcontracts/cmd"
)

// vendorCmd represents the device command
var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Commands for vendor",
}

func init() {
	cmd.RootCmd.AddCommand(vendorCmd)
}
