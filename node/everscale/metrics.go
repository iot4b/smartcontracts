package everscale

import "github.com/prometheus/client_golang/prometheus"

var RequestsCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "everscale_requests",
	Help: "everscale requests count",
})
