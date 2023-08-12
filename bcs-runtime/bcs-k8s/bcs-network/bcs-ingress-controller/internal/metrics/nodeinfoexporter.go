

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	nodeNotFound = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "bkbcs_ingressctrl",
		Subsystem: "node",
		Name:      "not_found_count",
		Help:      "node not found count for node exporter",
	}, []string{"node_name"})
)

func init() {
	metrics.Registry.MustRegister(nodeNotFound)
}

// IncreaseNodeNotFoundCounter increase node not found counter
func IncreaseNodeNotFoundCounter(nodeName string) {
	nodeNotFound.WithLabelValues(nodeName).Inc()
}
