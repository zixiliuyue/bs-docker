

// Package status xxx
package status

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-loadbalance/monitor"

	"github.com/emicklei/go-restful"
)

// Status status for bcs-loadbalance
type Status struct {
	hFunc restful.RouteFunction
}

// NewStatus new a status
func NewStatus(hf restful.RouteFunction) monitor.Resource {
	return &Status{
		hFunc: hf,
	}
}

// Register implements metrics.Resource interface
func (s *Status) Register(container *restful.Container) {
	blog.Infof("register prometheus monitor resource")
	ws := new(restful.WebService)
	ws.Path("/status").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(ws.GET("/").To(s.hFunc))
	container.Add(ws)
}
