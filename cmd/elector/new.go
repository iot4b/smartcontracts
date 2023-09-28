package elector

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"smartcontracts/everscale"
	"smartcontracts/shared/config"
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
		// todo проверять количество аргументов, иначе брать из stdin
		var (
			input          map[string]interface{}
			stdin          = cmd.InOrStdin()
			buf            []byte
			err            error
			public, secret string
		)

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

		// giver - это такой кошелек, который по
		abi, tvc, err := everscale.ReadContract("./contracts", "elector")
		if err != nil {
			log.Fatal(err)
			return
		}

		// init ContractBuilder
		elector := &everscale.ContractBuilder{Public: public, Secret: secret, Abi: abi, Tvc: tvc}
		elector.InitDeployOptions()

		// вычислив адрес, нужно на него завести средства, чтобы вы
		walletAddress := elector.CalcWalletAddress()

		// пополняем баланс wallet'a нового девайса
		giver := &everscale.Giver{
			Address: config.Get("giver.address"),
			Public:  config.Get("giver.public"),
			Secret:  config.Get("giver.secret"),
		}
		amount := 1_500_000_000
		log.Debugf("Giver: %s", giver.Address)
		log.Debug("Send Tokens from giver", "amount", amount, "from", giver.Address, "to", walletAddress, "amount", amount)
		err = giver.SendTokens("./contracts/giverv3.abi.json", walletAddress, amount)
		if err != nil {
			log.Fatalf("giver.SendTokens()", err)
			return
		}

		wait := 15 * time.Second
		log.Debugf("Wait %d seconds ...", wait.Seconds())
		time.Sleep(wait)

		// после всех сборок деплоим контракт
		log.Debug("Deploy ...")
		err = elector.Deploy(data)
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
		out["account"] = walletAddress
		out["public"] = public
		out["secret"] = secret

		// формируем json на выход
		result, err := json.Marshal(out)
		if err != nil {
			log.Fatal(err)
			return
		}
		// на выход адрес контракта отдаем
		err = utils.WriteToStdout(result)
		if err != nil {
			log.Fatal(err)
		}
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
