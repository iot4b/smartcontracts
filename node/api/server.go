package api

import (
	"github.com/coalalib/coalago"
	"github.com/ever-iot/node/api/proxy"
	"github.com/ever-iot/node/middleware"
	log "github.com/ndmsystems/golog"
)

var (
	client    = coalago.NewClient()
	coalaPort string
)

func Start(info map[string]interface{}, port string) {
	coalaPort = port

	log.Info("Start COALA server on :" + coalaPort)
	server := coalago.NewServer()

	server.GET("/info", handlerInfo(info))

	server.GET(middleware.CoalaMetrics("/a", alive))
	server.POST(middleware.CoalaMetrics("/cmd", sendCmd))
	server.GET(middleware.CoalaMetrics("/endpoints", getEndpoints))

	server.GET(middleware.CoalaMetrics("/contract", getDeviceContract))
	server.POST(middleware.CoalaMetrics("/register", registerDevice))

	server.GET(middleware.CoalaMetrics("/device/info", getAccountInfo))
	server.GET(middleware.CoalaMetrics("/device/locate", locateDevice))

	server.GET(middleware.CoalaMetrics("/blockchain/accountInfo", getAccountInfo))

	p := proxy.New()
	log.Fatal(p.RunWithServer(":"+coalaPort, server))
}
