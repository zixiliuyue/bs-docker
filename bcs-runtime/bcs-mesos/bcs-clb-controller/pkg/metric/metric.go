

// Package metric xxx
package metric

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
)

// Resource for a kind of clb metric
type Resource interface {
	Register(container *restful.Container)
}

// ClbMetric clb metrics
type ClbMetric struct {
	container *restful.Container
	server    *http.Server
}

// NewClbMetric create clb metrics
func NewClbMetric(port int) *ClbMetric {
	address := fmt.Sprintf(":%d", port)
	container := restful.NewContainer()
	server := &http.Server{
		Addr:    address,
		Handler: container,
	}
	return &ClbMetric{
		container: container,
		server:    server,
	}
}

// Run start metric server
func (cm *ClbMetric) Run() error {
	return cm.server.ListenAndServe()
}

// Close close clb metric
func (cm *ClbMetric) Close() error {
	return cm.server.Close()
}

// RegisterResource register resource to clb metric
func (cm *ClbMetric) RegisterResource(mr Resource) {
	mr.Register(cm.container)
}
