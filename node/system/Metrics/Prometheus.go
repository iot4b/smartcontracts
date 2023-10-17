package Metrics

import (
	"github.com/coalalib/coalago"
	"net/http"
	"os"
	"runtime"
	"time"

	log "github.com/ndmsystems/golog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	startTime    = time.Now()
	CoalaMetrics = Desc{
		{
			Desc: prometheus.NewDesc(
				"coala_received_messages_count",
				"Number of received messages",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricReceivedMessages.Val()) },
			ValType: prometheus.GaugeValue,
		},
		{
			Desc: prometheus.NewDesc(
				"coala_sent_messages_count",
				"Number of messages sent",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricSentMessages.Val()) },
			ValType: prometheus.GaugeValue,
		},
		{
			Desc: prometheus.NewDesc(
				"coala_sent_messages_error_count",
				"Number of errors that occur when sending messages",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricSentMessageErrors.Val()) },
			ValType: prometheus.GaugeValue,
		},
		{
			Desc: prometheus.NewDesc(
				"coala_expired_messages_count",
				"Number of messages that were not responsed",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricExpiredMessages.Val()) },
			ValType: prometheus.GaugeValue,
		},
		{
			Desc: prometheus.NewDesc(
				"coala_retransmited_messages_count",
				"Number of messages will be retransmited",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricRetransmitMessages.Val()) },
			ValType: prometheus.GaugeValue,
		},
		{
			Desc: prometheus.NewDesc(
				"coala_successful_handshakes_count",
				"Number of successful handshakes",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricSuccessfulHandhshakes.Val()) },
			ValType: prometheus.GaugeValue,
		},
		{
			Desc: prometheus.NewDesc(
				"coala_sessions_rate_count",
				"Count of added security sessions",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricSessionsRate.Val()) },
			ValType: prometheus.GaugeValue,
		},
		{
			Desc: prometheus.NewDesc(
				"coala_sessions_count",
				"Number of active security sessions",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricSessionsCount.Val()) },
			ValType: prometheus.GaugeValue,
		},
		{
			Desc: prometheus.NewDesc(
				"coala_breaked_msessage",
				"Number of breaked packages",
				nil, nil,
			),
			Eval:    func() float64 { return float64(coalago.MetricBreakedMessages.Val()) },
			ValType: prometheus.GaugeValue,
		},
	}
	ServiceMetricsDescs    Desc
	ServiceMetricsCounters []prometheus.Collector
)

type Desc []struct {
	Desc    *prometheus.Desc
	Eval    func() float64
	ValType prometheus.ValueType
}

type Collector struct {
	goroutines *prometheus.Desc
	mems       *prometheus.Desc
	uptime     *prometheus.Desc
}

func NewCollector() prometheus.Collector {
	return &Collector{
		goroutines: prometheus.NewDesc(
			"runtime_num_goroutines",
			"Number of goroutines that currently exist.",
			nil, nil),

		mems: prometheus.NewDesc(
			"runtime_memstats_sys",
			"total bytes of memory obtained from the OS.",
			nil, nil),

		uptime: prometheus.NewDesc(
			"service_uptime",
			"Service uptime",
			nil, nil),
	}
}

// Describe returns all descriptions of the collector.
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.goroutines
	ch <- c.mems
	ch <- c.uptime

	for _, i := range CoalaMetrics {
		ch <- i.Desc
	}

	for _, i := range ServiceMetricsDescs {
		ch <- i.Desc
	}
}

// Collect returns the current state of all metrics of the collector.
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.goroutines, prometheus.GaugeValue, float64(runtime.NumGoroutine()))

	ms := &runtime.MemStats{}
	runtime.ReadMemStats(ms)
	ch <- prometheus.MustNewConstMetric(c.mems, prometheus.GaugeValue, float64(ms.Sys))

	ch <- prometheus.MustNewConstMetric(c.uptime, prometheus.GaugeValue, float64(time.Since(startTime).Seconds()))

	for _, i := range CoalaMetrics {
		ch <- prometheus.MustNewConstMetric(i.Desc, i.ValType, i.Eval())
	}
	for _, i := range ServiceMetricsDescs {
		ch <- prometheus.MustNewConstMetric(i.Desc, i.ValType, i.Eval())
	}
}

// Init prometheus monitoring
func Start(port string) {
	opt := prometheus.ProcessCollectorOpts{
		PidFn: func() (int, error) {
			return os.Getpid(), nil
		},
	}
	prometheus.Unregister(prometheus.NewProcessCollector(opt))
	prometheus.Unregister(prometheus.NewGoCollector())

	prometheus.MustRegister(NewCollector())
	prometheus.MustRegister(ServiceMetricsCounters...)

	// start server
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Info("Prometheus init on port:", port)
	srv := http.Server{
		Addr:     "0.0.0.0:" + port,
		ErrorLog: zap.NewStdLog(zap.NewNop()),
		Handler:  mux,
	}
	log.Fatal(srv.ListenAndServe())
}
