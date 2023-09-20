package node

import (
	"smartcontracts/cmd"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
}

func init() {
	cmd.RootCmd.AddCommand(nodeCmd)
}
