package node

import (
	"encoding/json"
	"github.com/pkg/errors"
	"smartcontracts/cmd"
	"smartcontracts/everscale"
	log "smartcontracts/shared/golog"

	"github.com/spf13/cobra"
)

// newCmd - Deploy smart contract for new Device
var newCmd = &cobra.Command{
	Use:   "new {public} {secret} {initialData}",
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

		log.Debugw("Deploy new device",
			"public", public,
			"secret", secret,
			"initialData", data)

		return nil
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
