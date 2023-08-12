

package nginx

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Describe implements collector interface
func (m *Manager) Describe(ch chan<- *prometheus.Desc) {

}

// Collect implements collector interface
func (m *Manager) Collect(ch chan<- prometheus.Metric) {

}
