package everscale

import (
	"encoding/json"
	"fmt"
	"github.com/markgenuine/ever-client-go/domain"
	"github.com/pkg/errors"
	"io"
	"os"
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "os.Open(%s)", path)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "io.ReadAll(%s)", path)
	}
	return bytes, nil
}

func getAbiFromFile(path string) (*domain.Abi, error) {
	byteAbi, err := readFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "readFile")
	}
	ac := &domain.AbiContract{}
	err = json.Unmarshal(byteAbi, ac)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal(byteAbi, ac)")
	}
	return domain.NewAbiContract(ac), nil
}

func processMessage(abi *domain.Abi, address, method string, input interface{}, signer *domain.Signer) (*domain.ResultOfProcessMessage, error) {
	return Ever.Processing.ProcessMessage(&domain.ParamsOfProcessMessage{
		MessageEncodeParams: &domain.ParamsOfEncodeMessage{
			Address: address,
			Abi:     abi,
			CallSet: &domain.CallSet{
				FunctionName: method,
				Input:        input,
			},
			Signer: signer,
		},
		SendEvents: false,
	}, nil)
}

func NewSigner(public, secret string) *domain.Signer {
	return domain.NewSigner(domain.SignerKeys{Keys: &domain.KeyPair{
		Public: public,
		Secret: secret,
	}})
}

func ReadContract(name string) (abi *domain.Abi, tvc []byte, err error) {
	abi, err = getAbiFromFile(fmt.Sprintf("../build/%s.abi.json", name))
	if err != nil {
		err = fmt.Errorf("getAbiFromFile(%s.abi.json): %w", name, err)
		return
	}

	tvc, err = readFile(fmt.Sprintf("../build/%s.tvc", name))
	if err != nil {
		err = fmt.Errorf("readFile(%s.tvc): %w", name, err)
	}
	return
}
