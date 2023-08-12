

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	ListenerTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "bkbcs_ingressctrl",
		Subsystem: "listener",
		Name:      "total",
		Help:      "The total listener managed by controller",
	}, []string{"status", "target_group_type"})

	PortBindingTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "bkbcs_ingressctrl",
		Subsystem: "portbinding",
		Name:      "total",
		Help:      "The total portBinding managed by controller",
	}, []string{"status"})
)

func init() {
	metrics.Registry.MustRegister(ListenerTotal)
	metrics.Registry.MustRegister(PortBindingTotal)
}
