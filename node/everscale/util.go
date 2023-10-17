package everscale

import (
	"encoding/json"
	"github.com/ever-iot/node/dsm"
	"github.com/markgenuine/ever-client-go/domain"
	log "github.com/ndmsystems/golog"
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
	ac := &domain.AbiContract{}
	err = json.Unmarshal(byteAbi, &ac)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal(byteAbi, &ac)")
	}
	return domain.NewAbiContract(ac), nil
}

func processMessage(abi *domain.Abi, address dsm.EverAddress, method string, input, signer interface{}) (res *domain.ResultOfProcessMessage, err error) {
	res, err = ever.Processing.ProcessMessage(&domain.ParamsOfProcessMessage{
		MessageEncodeParams: &domain.ParamsOfEncodeMessage{
			Address: string(address),
			Abi:     abi,
			CallSet: &domain.CallSet{
				FunctionName: method,
				Input:        input,
			},
			Signer: domain.NewSigner(signer),
		},
		SendEvents: false,
	}, nil)
	if err != nil {
		log.Error("ever.Processing.ProcessMessage:", err)
		return
	}
	RequestsCount.Inc()
	return
}
