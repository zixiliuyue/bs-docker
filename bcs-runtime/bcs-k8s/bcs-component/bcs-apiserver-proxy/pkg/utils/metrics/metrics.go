

// Package metrics xxx
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"time"
)

const (
	// ErrStatus for call api failure
	ErrStatus = "failure"
	// SucStatus for call api success
	SucStatus = "success"
	// FailConnect for server connect failure
	FailConnect = "999"

	// BkBcsApiserverProxy for bcs-apiserver-proxy module metrics prefix
	BkBcsApiserverProxy = "bkbcs_apiserverproxy"
)

var (
	// bcs-apiserver-proxy request action metrics
	requestTotalAPI = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: BkBcsApiserverProxy,
		Name:      "apiserverproxy_request_total_num",
		Help:      "The total num of requests for external api",
	}, []string{"handler", "method", "code"})
	requestLatencyAPI = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: BkBcsApiserverProxy,
		Name:      "apiserverproxy_request_latency_time",
		Help:      "api request latency statistic for external api",
		Buckets:   []float64{0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.0, 3.0},
	}, []string{"handler", "method", "code"})
)

func init() {
	// bcs-apiserver-proxy call external api
	prometheus.MustRegister(requestTotalAPI)
	prometheus.MustRegister(requestLatencyAPI)
}

// reportBcsApiserverProxyAPIMetrics report all api action metrics
func reportBcsApiserverProxyAPIMetrics(handler, method, code string, started time.Time) {
	requestTotalAPI.WithLabelValues(handler, method, code).Inc()
	requestLatencyAPI.WithLabelValues(handler, method, code).Observe(time.Since(started).Seconds())
}
