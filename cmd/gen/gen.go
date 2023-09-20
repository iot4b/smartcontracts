package gen

import (
	"github.com/spf13/cobra"
	"smartcontracts/cmd"
)

// generatorCmd represents the device command
var generatorCmd = &cobra.Command{
	Use:   "gen",
	Short: "Commands for crypto keys, additional data for smart chains",
}

func init() {
	cmd.RootCmd.AddCommand(generatorCmd)
}
