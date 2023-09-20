package everscale

import (
	"log"
	"smartcontracts/shared/config"

	"github.com/markgenuine/ever-client-go"
)

var ever *goever.Ever

func Init() {
	var err error
	ever, err = goever.NewEver("", config.List("everscale.endpoints"), "")
	if err != nil {
		log.Fatal(err)
	}
}
