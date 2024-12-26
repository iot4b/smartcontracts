package elector

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"smartcontracts/everscale"
	log "smartcontracts/shared/golog"
	"smartcontracts/utils"
	"time"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new [initialData]",
	Short: "Use {public} and {secret} keys for Sign",
	Long: `Deploy smart contract for new Elector.
{initialData} - список нод, которые по-умолчанию прописываются в Elector
{
	defaultNodes []string - список нод по-умолчанию
}
на выходе получаем адресс нового контракта  
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("invalid arguments count")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			input          map[string]interface{}
			stdin          = cmd.InOrStdin()
			buf            []byte
			err            error
			public, secret string
		)

		log.Debug("args", args, "stdin", stdin)

		// если передаем входные данные строкой
		if len(args) > 0 {
			err = json.Unmarshal([]byte(args[0]), &input)
			if err != nil {
				log.Fatal(err)
			}
		} else { // если передаем входные данные из stdin
			// парсим stdin c initial data. формат json
			buf, err = io.ReadAll(cmd.InOrStdin())
			if err != nil {
				log.Fatal(err)
				return
			}
			err = json.Unmarshal(buf, &input)
			if err != nil {
				log.Fatal(err)
			}
		}

		public, ok := input["public"].(string)
		secret, secOk := input["secret"].(string)
		if !(ok && len(public) > 0 && secOk && len(secret) > 0) {
			public, secret = everscale.GenKeys()
		}

		log.Debug("initial data", input)

		// оставляем толко те поля, которые нужны для инициализации контракта
		data := initialData{}
		err = utils.JsonMapToStruct(input, &data)
		if err != nil {
			log.Fatal(err)
			return
		}
		// валидируем
		err = data.validate()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Debug("validate initial data OK!")

		abi, tvc, err := everscale.ReadContract("Elector")
		if err != nil {
			log.Fatal(err)
			return
		}

		// init ContractBuilder
		elector := &everscale.ContractBuilder{Public: public, Secret: secret, Abi: abi, Tvc: tvc}
		elector.InitDeployOptions(data)

		// пополняем баланс wallet'a нового девайса
		amount := 1_500_000_000
		log.Debugf("Giver: %s", everscale.Giver.Address)
		log.Debug("Send Tokens from giver", "amount", amount, "from", everscale.Giver.Address, "to", elector.Address, "amount", amount)
		err = everscale.Giver.SendTo(elector.Address, amount)
		if err != nil {
			log.Fatal("giver.SendTokens()", err)
			return
		}

		wait := 15 * time.Second
		log.Debugf("Wait %v seconds ...", wait.Seconds())
		time.Sleep(wait)

		// после всех сборок деплоим контракт
		log.Debug("Deploy ...")
		err = elector.Deploy()
		if err != nil {
			log.Fatal(err)
			return
		}

		// формируем ответ в формате json
		out := make(map[string]interface{})
		err = utils.JsonMapToStruct(data, &out)
		if err != nil {
			log.Fatal(err)
			return
		}
		out["account"] = elector.Address
		out["public"] = public
		out["secret"] = secret

		// формируем json на выход
		result, err := json.Marshal(out)
		if err != nil {
			log.Fatal(err)
			return
		}
		// на выход адрес контракта отдаем
		utils.WriteToStdout(result)
	},
}

func init() {
	electorCmd.AddCommand(newCmd)
}

type initialData struct {
	DefaultNodes []everscale.EverAddress `json:"defaultNodes"`
}

func (d initialData) validate() error {
	return nil
}

func (d initialData) toMap() (result map[string]interface{}) {
	data, _ := json.Marshal(d)
	json.Unmarshal(data, &result)
	return
}
