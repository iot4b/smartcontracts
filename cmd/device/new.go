package device

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"smartcontracts/cmd"
	"smartcontracts/shared/golog"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new {public} {secret} {initialData}",
	Short: "Use {public} and {secret} keys for Sign with {initialData}",
	Long: `Deploy smart contract for new Device.

{initialData} - данные в json, которые требуются для сохранения в смарт контракте
{
	vendor		address - адрес контракта Vendor (производителя)
	owners		[]address - список аккаунтов, которые считаются владельцами устройства

	vendorName 	string - название вендора
	vendorData 	any	- зашифрованные данные производителя устройства

	type		string - тип/модель устройства
	version		string - версия прошивки
}

	на выходе получаем адресс контракта нового девайса 
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("not enough arguments")
		}

		public := args[0]
		secret := args[1]
		var data initialData

		if err := json.Unmarshal([]byte(args[2]), &data); err != nil {
			return err
		}
		err := data.validate()
		if err != nil {
			return err
		}
		log.Debug("data.toMap()", data.toMap())

		log.Debugw("Deploy new device",
			"public", public,
			"secret", secret,
			"initialData", data)

		return nil
	},
}

func init() {
	deviceCmd.AddCommand(newCmd)
}

type initialData struct {
	Vendor cmd.EverAddress   `json:"vendor"`
	Owners []cmd.EverAddress `json:"owners"`

	VendorName string      `json:"vendorName"`
	VendorData interface{} `json:"vendorData"`

	Type    string `json:"type"`
	Version string `json:"version"`
}

func (d initialData) validate() error {
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
