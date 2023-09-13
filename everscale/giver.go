package everscale

import (
	"github.com/markgenuine/ever-client-go/domain"
	"github.com/pkg/errors"
	"smartcontracts/config"
)

// getTokensFromGiver transfer a [value] of test nanotokens from giverAddress to [account]
func getTokensFromGiver(account string, value int) (err error) {
	giverAddress := config.Get("giver.address")
	signerKeys := &domain.KeyPair{
		Public: config.Get("giver.public"),
		Secret: config.Get("giver.secret"),
	}

	giverAbiFile := "contract-giver/giver.abi.json"
	abi, err := getAbiFromFile(giverAbiFile)
	if err != nil {
		return errors.Wrapf(err, "getAbiFromFile(%s)", giverAbiFile)
	}

	input := sendTransaction{
		Dest:   account,
		Value:  value,
		Bounce: false,
	}
	_, err = processMessage(abi, giverAddress, "sendTransaction", input, domain.SignerKeys{Keys: signerKeys})
	return
}
