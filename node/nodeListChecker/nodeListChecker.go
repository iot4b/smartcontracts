package nodeListChecker

import (
	"github.com/coalalib/coalago"
	"github.com/ever-iot/node/smartmoc"
	log "github.com/ndmsystems/golog"
	"time"
)

// Run status monitoring of all current nodes
func Run(interval time.Duration) {
	log.Debug("Run nodeListChecker")

	coala := coalago.NewClient()
	elector := smartmoc.Elector{}
	nodeList := elector.GetCurrentPeriodNodes()
	for {
		time.Sleep(interval)
		for _, node := range nodeList {
			res, err := coala.GET("coap://" + node.IpPort + "/info")
			if err != nil {
				log.Errorf("ping %s: offline, kick node! %s", node, err)
				elector.ReportNodeFail(node.IpPort)
				continue
			}
			log.Debugf("ping %s: online, %s", node, res.Body)
		}
	}
}
