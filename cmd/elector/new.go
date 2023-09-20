package elector

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	log "smartcontracts/shared/golog"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new {public} {secret} {defaultNodes}",
	Short: "Use {public} and {secret} keys for Sign",
	Long: `Deploy smart contract for new Elector.
{defaultNodes} - список нод, которые по-умолчанию прописываются в Elector
	
на выходе получаем адресс нового контракта  
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("not enough arguments")
		}

		public := args[0]
		secret := args[1]
		// todo провалидировать как массив
		defaultNodes := args[2]

		log.Debugw("Deploy new elector",
			"public", public,
			"secret", secret,
			"defaultNodes", defaultNodes)

		return nil
	},
}

func init() {
	electorCmd.AddCommand(newCmd)
}
