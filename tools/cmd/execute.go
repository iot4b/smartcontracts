package cmd

import (
	"encoding/json"
	"smartcontracts/everscale"
	"smartcontracts/shared/golog"

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

		var input map[string]interface{}
		if len(args) > 3 {
			if err := json.Unmarshal([]byte(args[3]), &input); err != nil {
				return err
			}
		}
		log.Debug("input", input)
		_, err := everscale.Execute(name, address, method, input)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(executeCmd)
}
