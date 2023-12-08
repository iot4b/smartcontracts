package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"smartcontracts/everscale"
	"smartcontracts/utils"
)

// upgradeCmd runs upgrade method on a contract to upgrade its code
var upgradeCmd = &cobra.Command{
	Use:   "upgrade {name} {address}",
	Short: "Upgrade {name} smartcontract code deployed to {address} with the new code from file {name}.code",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("not enough arguments")
		}

		name := args[0]
		address := args[1]

		code, err := utils.ReadFile(fmt.Sprintf("../_%s/%s.tvc", name, name))
		if err != nil {
			return err
		}

		_, err = everscale.Execute(name, address, "upgrade", map[string][]byte{"code": code})
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(upgradeCmd)
}
