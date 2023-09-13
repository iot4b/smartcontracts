package everscale

import (
	"encoding/json"
	"io"
	"os"

	"github.com/markgenuine/ever-client-go/domain"
	"github.com/pkg/errors"
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
	ac := &domain.AbiContract{}
	err = json.Unmarshal(byteAbi, &ac)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal(byteAbi, &ac)")
	}
	return domain.NewAbiContract(ac), nil
}

func processMessage(abi *domain.Abi, address, method string, input, signer interface{}) (*domain.ResultOfProcessMessage, error) {
	return ever.Processing.ProcessMessage(&domain.ParamsOfProcessMessage{
		MessageEncodeParams: &domain.ParamsOfEncodeMessage{
			Address: address,
			Abi:     abi,
			CallSet: &domain.CallSet{
				FunctionName: method,
				Input:        input,
			},
			Signer: domain.NewSigner(signer),
		},
		SendEvents: false,
	}, nil)
}
