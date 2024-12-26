package main

import (
	"smartcontracts/everscale"
	"smartcontracts/shared/config"
	"smartcontracts/shared/golog"

	"smartcontracts/cmd"
	_ "smartcontracts/cmd/device"
	_ "smartcontracts/cmd/elector"
	_ "smartcontracts/cmd/gen"
	_ "smartcontracts/cmd/node"
	_ "smartcontracts/cmd/vendr"
)

func main() {

	everscale.Init()
	defer everscale.Destroy()

	cmd.Execute()
}

func init() {
	config.Init("config") // read config from ./config.yml

	log.Init(true, true, log.Console)
}
