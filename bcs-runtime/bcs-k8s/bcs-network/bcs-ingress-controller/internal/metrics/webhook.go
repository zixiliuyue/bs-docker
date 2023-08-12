

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	podCreateCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "bkbcs_ingressctrl",
		Subsystem: "webhook",
		Name:      "pod_create",
		Help:      "The total number of pod create checked by webhook",
	}, []string{"allow"})
)

func init() {
	metrics.Registry.MustRegister(podCreateCounter)
}

// IncreasePodCreateCounter increase pod create reject counter
func IncreasePodCreateCounter(allow bool) {
	var allowStr string
	if allow {
		allowStr = "true"
	} else {
		allowStr = "false"
	}
	podCreateCounter.WithLabelValues(allowStr).Inc()
}
