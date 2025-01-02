package main

import (
	"flag"
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
	var env string
	flag.StringVar(&env, "env", "dev", "set environment")
	flag.Parse()

	config.Init(env)

	log.Init(true, true, log.Console)
}
