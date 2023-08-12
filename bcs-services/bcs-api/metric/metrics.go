

package metric

import (
	"net/http"
	"strconv"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TimeBuckets is based on Prometheus client_golang prometheus.DefBuckets
var timeBuckets = prometheus.ExponentialBuckets(0.00025, 2, 16) // from 0.25ms to 8 seconds

// Metrics the bcs-api exports.
var (
	RequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "bcs_api",
		Name:      "request_count_total",
		Help:      "Counter of requests to bcs-api.",
	}, []string{"type", "method"})

	RequestErrorCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "bcs_api",
		Name:      "request_err_count_total",
		Help:      "Counter of error requests to bcs-api.",
	}, []string{"type", "method"})

	RequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "bcs_api",
		Name:      "request_latency_seconds",
		Buckets:   timeBuckets,
		Help:      "Histogram of the time (in seconds) each request took.",
	}, []string{"type", "method"})

	RequestErrorLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "bcs_api",
		Name:      "request_error_latency_seconds",
		Buckets:   timeBuckets,
		Help:      "Histogram of the time (in seconds) each error request took.",
	}, []string{"type", "method"})
)

// RunMetric xxx
func RunMetric(conf *config.ApiServConfig, err error) {

	blog.Infof("run metric: port(%d)", conf.MetricPort)
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestErrorCount)
	prometheus.MustRegister(RequestLatency)
	prometheus.MustRegister(RequestErrorLatency)
	http.Handle("/metrics", promhttp.Handler())
	addr := conf.Address + ":" + strconv.Itoa(int(conf.MetricPort))
	go http.ListenAndServe(addr, nil)

	blog.Infof("run metric ok")
}
