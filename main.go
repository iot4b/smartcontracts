package main

import (
	"os"
	"smartcontracts/cmd"
	_ "smartcontracts/cmd/device"
	_ "smartcontracts/cmd/node"
	"smartcontracts/config"
	"smartcontracts/everscale"
	log "smartcontracts/golog"
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
