package proxy

import "github.com/prometheus/client_golang/prometheus"

var Rate = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "coala_proxied_messages_count",
	Help:        "Number of proxied messages sent",
	ConstLabels: nil,
})
