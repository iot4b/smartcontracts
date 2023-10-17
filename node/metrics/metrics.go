package metrics

import (
	"github.com/ever-iot/node/api/proxy"
	"github.com/ever-iot/node/db"
	"github.com/ever-iot/node/middleware"
	"github.com/ever-iot/node/system/Metrics"

	"github.com/prometheus/client_golang/prometheus"
)

func Init() {
	Metrics.ServiceMetricsCounters = []prometheus.Collector{
		db.WritesCount,
		db.DevicesNum,
		db.IpPortChange,
		proxy.Rate,
		middleware.Stat,
	}
}
