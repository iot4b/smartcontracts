package db

import "github.com/prometheus/client_golang/prometheus"

var (
	WritesCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "db_writes",
		Help: "Writes to db devices count",
	})
	DevicesNum = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_devices_num",
		Help: "Number of registered devices",
	})
	IpPortChange = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "db_alive_change_ipport",
		Help: "Count changing ip or port",
	}, []string{
		"type",
	})
)
