package cmd

import (
	"errors"
	"smartcontracts/everscale"
	"strconv"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy {name} {balance}",
	Short: "Deploys a contract with provided name",
	Long: `Deploys a contract from directory "contract-{name}" with initial {balance}
Directory should contain files:
* {name}.abi.json
* {name}.tvc`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("not enough arguments")
		}

		name := args[0]
		balance, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.New("wrong {balance} value")
		}

		_, err = everscale.Deploy(name, balance)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(deployCmd)
}
