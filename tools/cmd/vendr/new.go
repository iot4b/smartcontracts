package vendor

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"smartcontracts/cmd"
	"smartcontracts/everscale"
	log "smartcontracts/shared/golog"
	"smartcontracts/utils"
	"time"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new [initialData]",
	Short: "Use {public} and {secret} keys for Sign with {initialData}",
	Long: `Deploy smart contract for new vendor.
	на выходе получаем адресс нового контракта  
{initialData}
{
	elector			address - адрес электора
	vendorName		string	- название производителя девайсов
	contactInfo		string	- контактная информация
	profitShare		int		- доля прибыли в процентах относительно ноды   
}
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

		abi, tvc, err := everscale.ReadContract("Vendor")
		if err != nil {
			log.Fatal(err)
			return
		}

		// init ContractBuilder
		vendor := &everscale.ContractBuilder{Public: public, Secret: secret, Abi: abi, Tvc: tvc}
		vendor.InitDeployOptions(data)

		// пополняем баланс wallet'a нового девайса
		amount := 1_500_000_000
		log.Debugf("Giver: %s", everscale.Giver.Address)
		log.Debug("Send Tokens from giver", "amount", amount, "from", everscale.Giver.Address, "to", vendor.Address, "amount", amount)
		err = everscale.Giver.SendTo(vendor.Address, amount)
		if err != nil {
			log.Fatal("giver.SendTokens()", err)
			return
		}

		wait := 15 * time.Second
		log.Debugf("Wait %v seconds ...", wait.Seconds())
		time.Sleep(wait)

		// после всех сборок деплоим контракт
		log.Debug("Deploy ...")
		err = vendor.Deploy()
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
		out["account"] = vendor.Address
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
	vendorCmd.AddCommand(newCmd)
}

type initialData struct {
	Elector     everscale.EverAddress `json:"elector"`
	VendorName  string                `json:"vendorName"`
	ContactInfo string                `json:"contactInfo"`
	ProfitShare int                   `json:"profitShare"`
}

func (d initialData) validate() error {
	if len(d.Elector) == 0 {
		return errors.Wrap(cmd.ErrIsRequired, "elector")
	}
	if len(d.VendorName) == 0 {
		return errors.Wrap(cmd.ErrIsRequired, "vendorName")
	}
	if len(d.ContactInfo) == 0 {
		return errors.Wrap(cmd.ErrNotSpecified, "contactInfo")
	}
	if d.ProfitShare == 0 {
		return errors.Wrap(cmd.ErrNotSpecified, "profitShare")
	}
	return nil
}

func (d initialData) toMap() (result map[string]interface{}) {
	data, _ := json.Marshal(d)
	json.Unmarshal(data, &result)
	return
}
