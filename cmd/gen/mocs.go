package gen

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// mocsCmd - Deploy smart contract for new Device
var mocsCmd = &cobra.Command{
	Use:   "mocs {name} ",
	Short: "Use {public} and {secret} keys for Sign with {initialData}",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("invalid arguments count")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}
