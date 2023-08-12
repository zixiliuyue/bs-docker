

package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type PromServer struct{}

var (
	hrCreateDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "bkbcs",
		Subsystem: "gameworkload",
		Name:      "hookrun_create_duration_seconds",
		Help:      "create duration(seconds) of hookrun",
		Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1, 5, 10, 20, 30, 60, 120},
	}, []string{"namespace", "name", "status", "action", "objectKind"})
)

func init() {
	prometheus.MustRegister(hrCreateDuration)
}

func (p *PromServer) CollectHRCreateDurations(namespace, name, status, action, objectKind string, d time.Duration) {
	hrCreateDuration.WithLabelValues(namespace, name, status, action, objectKind).Observe(d.Seconds())
}
