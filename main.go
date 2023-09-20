package main

import (
	"os"
	"smartcontracts/everscale"
	"smartcontracts/shared/config"
	"smartcontracts/shared/golog"

	"smartcontracts/cmd"
	_ "smartcontracts/cmd/device"
	_ "smartcontracts/cmd/gen"
	_ "smartcontracts/cmd/node"
)

func main() {

	everscale.Init()
	defer everscale.Destroy()

	cmd.Execute()
}

func init() {
	if _, err := os.Stat("config.yml"); err == nil {
		config.Init("config") // read config from ./config.yml
	} else {
		config.Init("config.sample") // read config from ./config.sample.yml
	}

	log.Init(true, true, log.Json)
}
