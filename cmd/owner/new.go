package owner

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	log "smartcontracts/shared/golog"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new {public} {secret} {elector}",
	Short: "Use {public} and {secret} keys for Sign",
	Long: `Deploy smart contract for new Owner.
	на выходе получаем адресс нового контракта  
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("not enough arguments")
		}

		public := args[0]
		secret := args[1]
		elector := args[2]

		log.Debugw("Deploy new owner",
			"public", public,
			"secret", secret,
			"elector", elector)

		return nil
	},
}

func init() {
	ownerCmd.AddCommand(newCmd)
}
