

package bcscustom

import (
	"github.com/coredns/coredns/plugin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Success TODO
	Success = "success"
	// Failure TODO
	Failure = "failure"
)

// Metrics the bcscustom plugin exports.
var (
	RequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcscustom",
		Name:      "request_count_total",
		Help:      "Counter of requests to plugin bcscustom.",
	}, []string{"status"})

	RequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcscustom",
		Name:      "request_latency_seconds",
		Buckets:   plugin.TimeBuckets,
		Help:      "Histogram of the time (in seconds) each request took.",
	}, []string{"status"})

	DnsTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcscustom",
		Name:      "dns_total",
		Help:      "Total counter of dns in plugin bcscustom.",
	})
)
