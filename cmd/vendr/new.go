package vendor

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"smartcontracts/cmd"
	"smartcontracts/everscale"
	log "smartcontracts/shared/golog"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new {public} {secret} {initialData}",
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
		if err := data.validate(); err != nil {
			return err
		}

		log.Debug("data.toMap()", data.toMap())

		log.Debugw("Deploy new vendor",
			"public", public,
			"secret", secret,
			"initialData", data)

		return nil
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
