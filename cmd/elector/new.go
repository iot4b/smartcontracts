package elector

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io"
	log "smartcontracts/shared/golog"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new [initialData]",
	Short: "Use {public} and {secret} keys for Sign",
	Long: `Deploy smart contract for new Elector.
{defaultNodes} - список нод, которые по-умолчанию прописываются в Elector
	
на выходе получаем адресс нового контракта  
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// todo проверять количество аргументов, иначе брать из stdin
		var input map[string]interface{}
		var stdin = cmd.InOrStdin()
		var buf []byte
		var err error
		log.Debug("args", args, "stdin", stdin)

		// если передаем входные данные строкой
		if len(args) == 1 {
			err = json.Unmarshal([]byte(args[0]), &input)
			if err != nil {
				log.Fatal(err)
			}
		}
		// если передаем входные данные из stdin
		if len(args) < 1 {
			// парсим stdin c initial data. формат json
			buf, err = io.ReadAll(cmd.InOrStdin())
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(buf, &input)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Debug("initial data", input)

		return nil
	},
}

func init() {
	electorCmd.AddCommand(newCmd)
}
