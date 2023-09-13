package main

import (
	"os"
	"smartcontracts/cmd"
	"smartcontracts/config"
	"smartcontracts/everscale"
)

func main() {
	if _, err := os.Stat("config.yml"); err == nil {
		config.Init("config") // read config from ./config.yml
	} else {
		config.Init("config.sample") // read config from ./config.sample.yml
	}

	everscale.Init()
	defer everscale.Destroy()

	cmd.Execute()
}
