

// Package monitor xxx
package monitor

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
)

// Resource for a kind of clb metric
type Resource interface {
	Register(container *restful.Container)
}

// Monitor monitoring data for bcs-loadbalance
type Monitor struct {
	container *restful.Container
	server    *http.Server
}

// NewMonitor create clb monitor
func NewMonitor(addr string, port int) *Monitor {
	address := fmt.Sprintf("%s:%d", addr, port)
	container := restful.NewContainer()
	server := &http.Server{
		Addr:    address,
		Handler: container,
	}
	return &Monitor{
		container: container,
		server:    server,
	}
}

// Run start monitor server
func (cm *Monitor) Run() error {
	return cm.server.ListenAndServe()
}

// Close close monitor
func (cm *Monitor) Close() error {
	return cm.server.Close()
}

// RegisterResource register resource to monitor
func (cm *Monitor) RegisterResource(mr Resource) {
	mr.Register(cm.container)
}
