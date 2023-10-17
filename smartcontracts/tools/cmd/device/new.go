package device

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"smartcontracts/cmd"
	"smartcontracts/everscale"
	"smartcontracts/shared/config"
	log "smartcontracts/shared/golog"
	"smartcontracts/utils"
	"time"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new [initialData]",
	Short: "Use {public} and {secret} keys for Sign with {initialData}",
	Long: `Deploy smart contract for new Device.

{initialData} - данные в json, которые требуются для сохранения в смарт контракте
{
Device: адреса контрактов связей
	node		address - адрес текущей ноды. при регистрации девайса нода установит свой адрес
	elector		address - адрес контракта Elector
	vendor		address - адрес контракта Vendor (производителя)
	owners		[]address - список аккаунтов, которые считаются владельцами устройства

Owner: по-умолчанию false, эти поля может изменять только owner. не передаются в конструктор
	active		bool - взаимодействует с системой или нет. если false, то устройство не обслуживается
	lock		bool - добровольная блокировка устройства. оно будет видно в системе, ему можно слать команды,
						но прямое взаимодействии с ним будет заблокировано. только owner может изменить его 
	stat		bool - вкл/выкл транслирование метрик и статистики через ноду в драйвчейн

Vendor:
	dtype		string - тип/модель устройства
	version		string - версия прошивки
	vendorName 	string - название вендора
	vendorData 	string - зашифрованные данные производителя устройства

Keys: пара ключей, если не указать, то генеряться новые
	private 	string - приватный ключ
	secret		string - секретный ключ
}

вспомогательные аргументы
giver		{} - адрес и ключи контракта, с которого будет зачислен начальны баланс
			если не укзаны ключи, то это должен сделать автоматически указанный адрес
balance		string - начальный баланс, который должен зачислить giver

на выходе получаем адресс контракта нового девайса 
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
		abi, tvc, err := everscale.ReadContract("../smartcontracts/device", "device")
		if err != nil {
			log.Fatal(err)
			return
		}

		// init ContractBuilder
		device := &everscale.ContractBuilder{Public: public, Secret: secret, Abi: abi, Tvc: tvc}
		device.InitDeployOptions()

		// вычислив адрес, нужно на него завести средства, чтобы вы
		walletAddress := device.CalcWalletAddress()

		// пополняем баланс wallet'a нового девайса
		giver := &everscale.Giver{
			Address: config.Get("giver.address"),
			Public:  config.Get("giver.public"),
			Secret:  config.Get("giver.secret"),
		}
		amount := 1_500_000_000
		log.Debugf("Giver: %s", giver.Address)
		log.Debug("Send Tokens from giver", "amount", amount, "from", giver.Address, "to", walletAddress, "amount", amount)
		err = giver.SendTokens("../smartcontracts/giver.abi.json", walletAddress, amount)
		if err != nil {
			log.Fatalf("giver.SendTokens()", err)
			return
		}

		wait := 15 * time.Second
		log.Debugf("Wait %d seconds ...", wait.Seconds())
		time.Sleep(wait)

		// после всех сборок деплоим контракт
		log.Debug("Deploy ...")
		err = device.Deploy(data)
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
	deviceCmd.AddCommand(newCmd)
}

type initialData struct {
	Node    everscale.EverAddress   `json:"node"`
	Elector everscale.EverAddress   `json:"elector"`
	Vendor  everscale.EverAddress   `json:"vendor"`
	Owners  []everscale.EverAddress `json:"owners"`

	Type       string `json:"dtype"`
	Version    string `json:"version"`
	VendorName string `json:"vendorName"`
	VendorData string `json:"vendorData"`
}

func (d initialData) validate() error {
	if len(d.Node) == 0 {
		return errors.Wrap(cmd.ErrIsRequired, "node")
	}
	if len(d.Elector) == 0 {
		return errors.Wrap(cmd.ErrIsRequired, "elector")
	}
	if len(d.Vendor) == 0 {
		return errors.Wrap(cmd.ErrIsRequired, "vendor")
	}
	if len(d.Owners) == 0 {
		return errors.Wrap(cmd.ErrIsEmpty, "owners")
	}
	if len(d.Owners) > 0 {
		// todo сделать валидатор ever адресов
		for i, owner := range d.Owners {
			if len(owner) == 0 {
				return errors.Wrapf(cmd.ErrInvalidValue, "owners[%d]", i)
			}
		}
	}
	if len(d.VendorName) == 0 {
		return errors.Wrap(cmd.ErrNotSpecified, "vendorName")
	}
	if len(d.Type) == 0 {
		return errors.Wrap(cmd.ErrNotSpecified, "type")
	}
	if len(d.Version) == 0 {
		return errors.Wrap(cmd.ErrNotSpecified, "version")
	}
	return nil
}

func (d initialData) toMap() (result map[string]interface{}) {
	data, _ := json.Marshal(d)
	json.Unmarshal(data, &result)
	return
}
