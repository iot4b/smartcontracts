package gen

import (
	"bufio"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"smartcontracts/everscale"
	"smartcontracts/utils"
)

// keysCmd represents the node command
var keysCmd = &cobra.Command{
	Use:   "keys (filePath) (ymlKey)",
	Short: "Generate ever sign keyPair",
	RunE: func(cmd *cobra.Command, args []string) error {
		var filePath string
		if len(args) > 0 {
			filePath = args[0]
		}
		keys, err := everscale.GenerateKeyPair()
		if err != nil {
			return err
		}
		data, err := json.Marshal(keys)
		if err != nil {
			return err
		}
		// значит нужно сохранять по указанному пути
		if len(filePath) > 0 {
			// если указан ymlKey, значит надо записать в файл с таким yml ключем
			if len(args) > 1 {
				// todo дописать в файл с указанным ключом
				return nil
			}
			return utils.SaveFile(filePath, data)
		}

		// иначе выводим в Stdout
		writer := bufio.NewWriter(os.Stdout)
		_, err = writer.WriteString(string(data))
		if err != nil {
			return err
		}
		writer.Flush()

		return nil
	},
}

func init() {
	generatorCmd.AddCommand(keysCmd)
}
