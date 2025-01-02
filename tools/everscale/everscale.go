package everscale

import (
	"github.com/markgenuine/ever-client-go/domain"
	"github.com/markgenuine/ever-client-go/gateway/client"
	"github.com/markgenuine/ever-client-go/usecase/boc"
	"log"
	"smartcontracts/shared/config"

	"github.com/markgenuine/ever-client-go"
)

var (
	Boc  domain.BocUseCase
	Ever *goever.Ever
)

func Init() {
	address := ""
	endPoints := config.List("endpoints")
	accessKey := ""

	var err error
	Ever, err = goever.NewEver(address, endPoints, accessKey)
	if err != nil {
		log.Fatal(err)
	}

	cfg := domain.NewDefaultConfig(address, endPoints, accessKey)
	cl, err := client.NewClientGateway(cfg)
	if err != nil {
		log.Fatal(err)
	}
	Boc = boc.NewBoc(cfg, cl)

	Giver = giver{
		Address: EverAddress(config.Get("giver.address")),
		Public:  config.Get("giver.public"),
		Secret:  config.Get("giver.secret"),
	}
}
