package device

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"smartcontracts/cmd"
	"smartcontracts/everscale"
	"smartcontracts/shared/config"
	log "smartcontracts/shared/golog"
	"smartcontracts/utils"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new",
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
		//if len(args) < 2 {
		//	return errors.New("not enough arguments")
		//}
		//
		//public := config.Get("signer.public")
		//secret := config.Get("signer.secret")
		//
		//var data initialData
		//if err := json.Unmarshal([]byte(args[2]), &data); err != nil {
		//	return err
		//}
		//
		//err := data.validate()
		//if err != nil {
		//	return err
		//}
		//log.Debug("data.toMap()", data.toMap())

		//log.Debugw("Deploy new device",
		//	"public", public,
		//	"secret", secret)
		// если аргументов больше 3, то значит передали giver и balance для зачисления начального баланса
		// для тестирования или ручного запуска ок. в реальной системе вендор сам должен пополнять
		// баланс каждого активированного девайса. если giver передан с ключами, значит девайс сам
		// может пополнить свой баланс при инициализации, но не более чем передан в поле balance

		public, secret := everscale.KeysFromFile()

		// giver - это такой кошелек, который по
		abi, tvc, err := everscale.ReadContract("./contracts", "device")
		if err != nil {
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
		err = giver.SendTokens("./contracts/giverv3.abi.json", walletAddress, 1_500_000_000)
		if err != nil {
			log.Errorf("giver.SendTokens()", err)
			return
		}

		// после всех сборок деплоим контракт
		err = device.Deploy()
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
	Node    everscale.EverAddress   `json:"node"`
	Elector everscale.EverAddress   `json:"elector"`
	Vendor  everscale.EverAddress   `json:"vendor"`
	Owners  []everscale.EverAddress `json:"owners"`

	Active bool `json:"active"`
	Lock   bool `json:"lock"`
	Stat   bool `json:"stat"`

	VendorName string      `json:"vendorName"`
	VendorData interface{} `json:"vendorData"`

	Type    string `json:"type"`
	Version string `json:"version"`
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
