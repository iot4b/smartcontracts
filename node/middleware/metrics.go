package middleware

import "github.com/prometheus/client_golang/prometheus"

var (
	Stat = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "requests",
		Help:        "Статистика запрошенных ресурсов",
		ConstLabels: nil,
	}, []string{"proto", "method", "path", "status"},
	)
)
