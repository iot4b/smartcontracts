package node

import (
	"encoding/json"
	"github.com/pkg/errors"
	"smartcontracts/cmd"
	"smartcontracts/everscale"
	"smartcontracts/shared/config"
	log "smartcontracts/shared/golog"
	"smartcontracts/utils"
	"time"

	"github.com/spf13/cobra"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new [ initialData ]",
	Short: "Use {public} and {secret} keys for Sign with {initialData}",
	Long: `Deploy smart contract for new node.
	на выходе получаем адресс нового контракта  
{initialData}
{	
	elector			address - адрес Elector'a
	location		string	- геолокация ноды
	ipPort			string	- фактический ip:port для подключения    
	contactInfo		string	- контактная информация
}
`,
	Run: func(cmd *cobra.Command, args []string) {
		public, secret := everscale.KeysFromFile()

		data := initialData{
			Elector:     "0:8e18edd847fdc6bdd95640b3ff76a90d1d12d757c92061d0bfb12a03440f759e",
			Location:    "59.984752,30.203979",
			IpPort:      "157.245.57.218:5683",
			ContactInfo: "Mister Jixer, +2284200248",
		}

		// giver - это такой кошелек, который по
		abi, tvc, err := everscale.ReadContract("./contracts", "node")
		if err != nil {
			log.Error(err)
			return
		}

		log.Debug("initialData", data)
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
		err = device.Deploy(nil)
		if err != nil {
			log.Error(err)
		}

		log.Debug("data.toMap()", data.toMap())
		log.Debugw("Deploy new device",
			"public", public,
			"secret", secret,
			"initialData", data)

		// на выход адрес контракта отдаем
		err = utils.WriteToStdout([]byte(walletAddress))
		if err != nil {
			log.Error(err)
		}
	},
}

func init() {
	nodeCmd.AddCommand(newCmd)
}

type initialData struct {
	Elector     everscale.EverAddress `json:"elector"`
	Location    string                `json:"location"`
	IpPort      string                `json:"ipPort"`
	ContactInfo string                `json:"contactInfo"`
}

func (d initialData) validate() error {
	if len(d.Elector) == 0 {
		return errors.Wrap(cmd.ErrIsRequired, "elector")
	}
	if len(d.Location) == 0 {
		return errors.Wrap(cmd.ErrNotSpecified, "location")
	}
	if len(d.IpPort) == 0 {
		return errors.Wrap(cmd.ErrNotSpecified, "ipPort")
	}
	if len(d.ContactInfo) == 0 {
		return errors.Wrap(cmd.ErrNotSpecified, "contactInfo")
	}
	return nil
}

func (d initialData) toMap() (result map[string]interface{}) {
	data, _ := json.Marshal(d)
	json.Unmarshal(data, &result)
	return
}
