

// Package prometheus xxx
package prometheus

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-loadbalance/monitor"

	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PromMetric metric instance for prometheus
type PromMetric struct{}

// NewPromMetric new a promethues instance
func NewPromMetric() monitor.Resource {
	return &PromMetric{}
}

// Register implements metrics.Resource interface
func (p *PromMetric) Register(container *restful.Container) {
	blog.Infof("register prometheus monitor resource")
	container.Handle("/metrics", promhttp.Handler())
}
