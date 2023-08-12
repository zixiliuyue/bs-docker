

package metric

import (
	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	testDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "test_datas_seconds",
			Help:       "test data distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"service"},
	)
)

func init() {
	prometheus.MustRegister(testDurations)
}

// PromMetric metric for prometheus
type PromMetric struct {
}

// NewPromMetric create prometheus metric
func NewPromMetric() Resource {
	return &PromMetric{}
}

// Register implements Resource interface
func (p *PromMetric) Register(container *restful.Container) {
	container.Handle("/metrics", promhttp.Handler())
}
