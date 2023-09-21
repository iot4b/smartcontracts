package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		//if len(args) < 2 {
		//	return errors.New("not enough arguments")
		//}
		//
		//name := args[0]
		//balance, err := strconv.Atoi(args[1])
		//if err != nil {
		//	return errors.New("wrong {balance} value")
		//}
		//
		//_, err = everscale.DeployWithBalance(name,
		//	"./contract-"+name,
		//	config.Get("signer.public"),
		//	config.Get("signer.secret"),
		//	everscale.Giver{
		//		Address: config.Get("giver.address"),
		//		Public:  config.Get("giver.public"),
		//		Secret:  config.Get("giver.secret"),
		//	}, balance)
		//return err
	},
}

func init() {
	RootCmd.AddCommand(deployCmd)
}
