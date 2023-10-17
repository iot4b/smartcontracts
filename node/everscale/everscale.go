package everscale

import (
	"github.com/ever-iot/node/system/config"
	"github.com/markgenuine/ever-client-go"
	log "github.com/ndmsystems/golog"
)

var ever *goever.Ever

func Init() {
	var err error
	ever, err = goever.NewEver("", config.List("everscale.endpoints"), "")
	if err != nil {
		log.Fatal(err)
	}
}
