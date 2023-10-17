package main

import (
	"fmt"
	"github.com/ever-iot/node/api"
	"github.com/ever-iot/node/cryptoKeys"
	"github.com/ever-iot/node/everscale"
	"github.com/ever-iot/node/node"
	"github.com/ever-iot/node/nodeListChecker"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/ever-iot/node/db"
	"github.com/ever-iot/node/metrics"
	"github.com/ever-iot/node/system/Metrics"
	"github.com/ever-iot/node/system/config"
	log "github.com/ndmsystems/golog"
)

func main() {
	log.Info("Run server")

	everscale.Init()
	defer everscale.Destroy()

	cryptoKeys.Init()
	node.Init(config.Get("localFiles.nodeContractData"), config.Get("coala.port"))

	//TODO при старте гурзить (если нет, создать)  из файла  node.key  публичныйй
	// и приватный ключи и исползовать его в коале

	db.Init(config.Get("localFiles.boltDB"), config.Time("device.aliveTimeout"))
	defer db.Close()

	metrics.Init()

	//COALA
	go api.Start(config.Info(), config.Get("coala.port"))

	go Metrics.Start(config.Get("prometheus.port"))

	go nodeListChecker.Run(config.Time("nodes.checkInterval"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

// инитим конфиги и logger
func init() {
	if len(os.Args) < 1 { //TODO убрать аргументы вообще, все в конфиге
		fmt.Println(`Usage: server [env]`)
		fmt.Println("Not enough arguments. Use defaults : dev")
		os.Exit(0)
	}
	config.Init(os.Args[1])
	log.Init(config.Bool("debug"))
}
