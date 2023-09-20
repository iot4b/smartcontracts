package device

import (
	"github.com/spf13/cobra"
	"smartcontracts/cmd"
)

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Commands for device",
}

func init() {
	cmd.RootCmd.AddCommand(deviceCmd)
}
