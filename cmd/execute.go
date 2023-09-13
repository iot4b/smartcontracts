package cmd

import (
	"smartcontracts/everscale"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:   "execute {name} {address} {method} [input]",
	Short: "Execute a {method} on a contract {name} deployed to {address}",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("not enough arguments")
		}

		name := args[0]
		address := args[1]
		method := args[2]

		_, err := everscale.Execute(name, address, method, nil)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
}
