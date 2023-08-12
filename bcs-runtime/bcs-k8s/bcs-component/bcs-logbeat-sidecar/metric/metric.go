

// Package metric xxx
package metric

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// MetricClient is interface for prometheus client
type MetricClient interface {
	Start()
}

// SidecarMetricClient is implementation of MetricClient
type SidecarMetricClient struct {
	Host string
	Port string
}

// NewMetricClient create a new MetricClient
func NewMetricClient(host, port string) MetricClient {
	return &SidecarMetricClient{
		Host: host,
		Port: port,
	}
}

// Start starts the metric client
func (c *SidecarMetricClient) Start() {
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%s", c.Host, c.Port), nil)
		if err != nil {
			blog.Errorf("Metric client stopped: %s", err.Error())
		}
	}()
	blog.Infof("Metric client start...")
}
