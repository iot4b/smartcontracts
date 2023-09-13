package everscale

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/markgenuine/ever-client-go/domain"
	"github.com/pkg/errors"
	"smartcontracts/config"
)

// Version of ever client
func Version() {
	res, err := ever.Client.Version()
	if err != nil {
		fmt.Println("ever.Client.Version():", err)
		return
	}
	fmt.Println(res.Version)
}

// Destroy client when finished
func Destroy() {
	ever.Client.Destroy()
}

// Deploy a contract from directory "smartcontracts/contrat-[name]" with initial [balance]
// returns deployed contract address
func Deploy(name string, balance int) (string, error) {
	fmt.Println("deploying contract:", name)

	abiFile := fmt.Sprintf("contract-%s/%s.abi.json", name, name)
	tvcFile := fmt.Sprintf("contract-%s/%s.tvc", name, name)

	signerKeys := &domain.KeyPair{
		Public: config.Get("signer.public"),
		Secret: config.Get("signer.secret"),
	}

	abi, err := getAbiFromFile(abiFile)
	if err != nil {
		return "", errors.Wrapf(err, "getAbiFromFile(%s)", abiFile)
	}

	tvc, err := readFile(tvcFile)
	if err != nil {
		return "", errors.Wrapf(err, "readFile(%s)", tvcFile)
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
		},
	}

	msg, err := ever.Abi.EncodeMessage(deployParams)
	if err != nil {
		return "", errors.Wrap(err, "ever.Abi.EncodeMessage")
	}

	if err = getTokensFromGiver(msg.Address, balance); err != nil {
		return "", errors.Wrapf(err, "getTokensFromGiver(%s, %v)", msg.Address, balance)
	}

	params := &domain.ParamsOfProcessMessage{
		MessageEncodeParams: deployParams,
		SendEvents:          false,
	}
	_, err = ever.Processing.ProcessMessage(params, nil)
	if err != nil {
		return "", errors.Wrap(err, "ProcessMessage")
	}

	fmt.Println(name, "contract is deployed, new address:")
	fmt.Println(msg.Address)

	return msg.Address, nil
}

// Execute a [method] on a contract [name] deployed to [address]
func Execute(name, address, method string, input interface{}) ([]byte, error) {
	fmt.Println("executing", method, "on", name, "contract at address", address)

	abiFile := fmt.Sprintf("contract-%s/%s.abi.json", name, name)
	abi, err := getAbiFromFile(abiFile)
	if err != nil {
		return nil, errors.Wrapf(err, "getAbiFromFile(%s)", abiFile)
	}

	result, err := processMessage(abi, address, method, input, domain.SignerNone{})
	if err != nil {
		return nil, errors.Wrap(err, "processMessage")
	}

	fmt.Println(string(result.Decoded.Output))
	return result.Decoded.Output, nil
}

// GetAccountInfo get balance and boc of the [address] account
func GetAccountInfo(address string) (AccountInfo, error) {
	res, err := ever.Net.Query(&domain.ParamsOfQuery{
		Query: fmt.Sprintf(`
		query {
		  blockchain {
			account(
			  address: "%s"
			) {
			   info {
				balance(format: DEC)
				boc
			  }
			}
		  }
		}`, address),
	})
	if err != nil {
		return AccountInfo{}, errors.Wrap(err, "ever.Net.Query")
	}

	result := &queryResult{}
	err = json.Unmarshal(res.Result, result)
	if err != nil {
		return AccountInfo{}, errors.Wrap(err, "json.Unmarshal")
	}

	info := result.Data.Blockchain.Account.Info
	info.Address = address

	return info, nil
}
