package everscale

import (
	"encoding/base64"
	"encoding/json"
	"github.com/ever-iot/node/dsm"
	"github.com/ever-iot/node/system/config"
	"github.com/markgenuine/ever-client-go/domain"
	log "github.com/ndmsystems/golog"
	"github.com/pkg/errors"
)

const (
	deviceAbiFile = "../smartcontracts/device/device.abi.json"
	deviceTvcFile = "../smartcontracts/device/device.tvc"
	giverAbiFile  = "../smartcontracts/giver/giver.abi.json"
)

type Device struct {
	Address dsm.EverAddress   `json:"address,omitempty"` //ever address of smart contract
	Node    dsm.EverAddress   `json:"node,omitempty"`    //address, изменяется
	Elector dsm.EverAddress   `json:"elector,omitempty"` //address
	Vendor  dsm.EverAddress   `json:"vendor,omitempty"`  //address
	Owners  []dsm.EverAddress `json:"owners,omitempty"`  //address

	//изменяемые параметры
	Active  bool   `json:"active,omitempty"`  //активен ли девайс в iot, изменяется
	Lock    bool   `json:"lock,omitempty"`    //статус залочен ли, изменяется
	Stat    bool   `json:"stat,omitempty"`    //слать ли статистику, изменяется
	Version string `json:"version,omitempty"` //версия прошивки, изменяется

	PublicKey  string `json:"publicKey,omitempty"`
	Type       string `json:"type,omitempty"`       //тип девайса
	VendorName string `json:"vendorName,omitempty"` //название вендора
	VendorData string `json:"vendorData,omitempty"` //происзолный блок данных вендора в зашифрованном виде
}

// Deploy device contract with initial [balance]
// returns keypair of the owner
func (d *Device) Deploy(balance int) (*domain.KeyPair, error) {
	log.Debug("deploying contract", *d)

	signerKeys, _ := GenerateKeyPair()

	abi, err := getAbiFromFile(deviceAbiFile)
	if err != nil {
		return nil, errors.Wrapf(err, "getAbiFromFile(%s)", deviceAbiFile)
	}

	tvc, err := readFile(deviceTvcFile)
	if err != nil {
		return nil, errors.Wrapf(err, "readFile(%s)", deviceTvcFile)
	}

	deployParams := &domain.ParamsOfEncodeMessage{
		Abi:    abi,
		Signer: domain.NewSigner(domain.SignerKeys{Keys: signerKeys}),
		DeploySet: &domain.DeploySet{
			Tvc:         base64.StdEncoding.EncodeToString(tvc),
			InitialData: json.RawMessage(`{}`),
		},
		CallSet: &domain.CallSet{
			FunctionName: "constructor",
			Input:        d,
		},
	}

	msg, err := ever.Abi.EncodeMessage(deployParams)
	if err != nil {
		return nil, errors.Wrap(err, "ever.Abi.EncodeMessage")
	}

	if err = d.getTokensFromGiver(balance); err != nil {
		return nil, errors.Wrapf(err, "getTokensFromGiver(%s, %v)", msg.Address, balance)
	}

	params := &domain.ParamsOfProcessMessage{
		MessageEncodeParams: deployParams,
		SendEvents:          false,
	}
	_, err = ever.Processing.ProcessMessage(params, nil)
	if err != nil {
		return nil, errors.Wrap(err, "ProcessMessage")
	}

	d.Address = dsm.EverAddress(msg.Address)
	d.PublicKey = signerKeys.Public

	log.Debug("new device contract deployed,", d.Address, *signerKeys)

	return signerKeys, nil
}

// SetNode to device contract
func (d *Device) SetNode(address dsm.EverAddress) error {
	_, err := Execute("device", d.Address, "setNode", setNodeReq{Address: address})
	return err
}

// GetNode from device contract
func (d *Device) GetNode() (address dsm.EverAddress, err error) {
	var data []byte
	data, err = Execute("device", d.Address, "getNode", nil)
	if err != nil {
		return
	}

	var res getNodeRes
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	address = res.Address
	d.Address = address

	return
}

// getTokensFromGiver transfer a [value] of test nanotokens from giverAddress to [account]
func (d *Device) getTokensFromGiver(value int) (err error) {
	giverAddress := dsm.EverAddress(config.Get("everscale.giver.address"))
	signerKeys := &domain.KeyPair{
		Public: config.Get("everscale.giver.public"),
		Secret: config.Get("everscale.giver.secret"),
	}

	abi, err := getAbiFromFile(giverAbiFile)
	if err != nil {
		return errors.Wrapf(err, "getAbiFromFile(%s)", giverAbiFile)
	}

	input := sendTransactionReq{
		Dest:   string(d.Address),
		Value:  value,
		Bounce: false,
	}
	_, err = processMessage(abi, giverAddress, "sendTransaction", input, domain.SignerKeys{Keys: signerKeys})
	return
}
