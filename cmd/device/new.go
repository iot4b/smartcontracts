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
	node		address - адрес текущей ноды. при регистрации девайса нода установит свой адрес
	elector		address - адрес контракта Elector
	vendor		address - адрес контракта Vendor (производителя)
	owners		[]address - список аккаунтов, которые считаются владельцами устройства

	по-умолчанию false, эти поля может изменять только owner
	active		bool - взаимодействует с системой или нет. если false, то устройство не обслуживается
	lock		bool - добровольная блокировка устройства. оно будет видно в системе, ему можно слать команды,
						но прямое взаимодействии с ним будет заблокировано. только owner может изменить его 
	stat		bool - вкл/выкл транслирование метрик и статистики через ноду в драйвчейн

	vendorName 	string - название вендора
	vendorData 	any	- зашифрованные данные производителя устройства

	type		string - тип/модель устройства
	version		string - версия прошивки
}

вспомогательные аргументы
giver		{} - адрес и ключи контракта, с которого будет зачислен начальны баланс
			если не укзаны ключи, то это должен сделать автоматически указанный адрес
balance		string - начальный баланс, который должен зачислить giver

на выходе получаем адресс контракта нового девайса 
`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("args", args)

		// this does the trick
		buf, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			log.Fatal(err)
		}

		log.Debug("stdin", string(buf))

		var input = make(map[string]interface{})
		//input["node"] = "0:8e18edd847fdc6bdd95640b3ff76a90d1d12d757c92061d0bfb12a03440f759e"
		err = json.Unmarshal(buf, &input)
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("initial data", input)

		public, secret := everscale.KeysFromFile()
		log.Debugf("init keys from file. public: %s secret: %s", public, secret)

		// giver - это такой кошелек, который по
		abi, tvc, err := everscale.ReadContract("./contracts", "device")
		if err != nil {
			log.Error(err)
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
		err = giver.SendTokens("./contracts/giverv3.abi.json", walletAddress, amount)
		if err != nil {
			log.Errorf("giver.SendTokens()", err)
			return
		}

		wait := 15 * time.Second
		log.Debugf("Wait %d seconds ...", wait.Seconds())
		time.Sleep(wait)

		// после всех сборок деплоим контракт
		log.Debug("Deploy ...")
		err = device.Deploy(input)
		if err != nil {
			log.Error(err)
		}

		// на выход адрес контракта отдаем
		err = utils.WriteToStdout([]byte(walletAddress))
		if err != nil {
			log.Error(err)
		}
	},
}

func init() {
	deviceCmd.AddCommand(newCmd)
}

type initialData struct {
	Node    everscale.EverAddress   `json:"_node"`
	Elector everscale.EverAddress   `json:"_elector"`
	Vendor  everscale.EverAddress   `json:"_vendor"`
	Owners  []everscale.EverAddress `json:"_owners"`

	Type       string `json:"_type"`
	Version    string `json:"_version"`
	VendorName string `json:"_vendorName"`
	VendorData string `json:"_vendorData"`
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
