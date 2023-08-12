

package util

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// ErrStatus for call api failure
	ErrStatus = "failure"
	// SucStatus for success status
	SucStatus = "success"
)

const (
	// BkBcsScheduler for bcs-scheduler moduler metrics prefix
	BkBcsScheduler = "bkbcs_scheduler"
)

var (
	// bcs-alert-manager api http request metrics
	requestTotalAPI = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: BkBcsScheduler,
		Name:      "lib_request_total_num",
		Help:      "The total num of requests for lib(bcs-alertmanager) api",
	}, []string{"handler", "method", "status"})

	requestLatencyAPI = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: BkBcsScheduler,
		Name:      "lib_request_latency_time",
		Help:      "api request latency statistic for lib(bcs-alertmanager) api",
		Buckets:   []float64{0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.0, 3.0},
	}, []string{"handler", "method", "status"})
)

func init() {
	// register api metrics
	prometheus.MustRegister(requestTotalAPI)
	prometheus.MustRegister(requestLatencyAPI)
}

// ReportLibAlertManagerAPIMetrics report all api action metrics
func ReportLibAlertManagerAPIMetrics(handler, method, status string, started time.Time) {
	requestTotalAPI.WithLabelValues(handler, method, status).Inc()
	requestLatencyAPI.WithLabelValues(handler, method, status).Observe(time.Since(started).Seconds())
}
