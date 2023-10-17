package everscale

import (
	"encoding/json"
	"fmt"
	"github.com/ever-iot/node/dsm"
	"github.com/markgenuine/ever-client-go/domain"
	log "github.com/ndmsystems/golog"
	"github.com/pkg/errors"
)

// Version of ever client
func Version() string {
	res, err := ever.Client.Version()
	if err != nil {
		log.Errorf("ever.Client.Version(): %s", err)
		return ""
	}
	RequestsCount.Inc()
	return res.Version
}

// Destroy client when finished
func Destroy() {
	ever.Client.Destroy()
}

// Execute a method on smart-contract deployed to address
func Execute(cType string, address dsm.EverAddress, method string, input interface{}) ([]byte, error) {
	log.Debug("Execute:", address, method, input)
	abi, err := getAbiFromFile(fmt.Sprintf("../smartcontracts/%s/%s.abi.json", cType, cType))
	if err != nil {
		return nil, errors.Wrapf(err, "getAbiFromFile(%s)", deviceAbiFile)
	}

	result, err := processMessage(abi, address, method, input, domain.SignerNone{})
	if err != nil {
		return nil, errors.Wrap(err, "processMessage")
	}

	log.Debug(string(result.Decoded.Output))
	return result.Decoded.Output, nil
}

// GetAccountInfo get balance and boc of the account
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
		return AccountInfo{}, err
	}

	RequestsCount.Inc()

	result := &queryResult{}
	err = json.Unmarshal(res.Result, result)
	if err != nil {
		return AccountInfo{}, errors.Wrap(err, "json.Unmarshal")
	}

	info := result.Data.Blockchain.Account.Info
	info.Address = address

	return info, nil
}

func GenerateKeyPair() (kp *domain.KeyPair, err error) {
	kp, err = ever.Crypto.GenerateRandomSignKeys()
	if err != nil {
		return
	}
	RequestsCount.Inc()
	return
}
