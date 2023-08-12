

package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	portBindLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "bkbcs_ingressctrl",
		Subsystem: "portpool",
		Name:      "bind_latency_seconds",
		Help:      "port bind latency for bcs ingress controller port-pool",
		Buckets:   []float64{1, 5, 10, 30, 60, 120, 180, 300, 600, 1200},
	}, []string{})
)

func init() {
	metrics.Registry.MustRegister(portBindLatency)
}

// ReportPortBindMetric report port bind metrics
func ReportPortBindMetric(started time.Time) {
	portBindLatency.WithLabelValues().Observe(time.Since(started).Seconds())
}
