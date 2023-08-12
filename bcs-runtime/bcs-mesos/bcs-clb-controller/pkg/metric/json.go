

package metric

import (
	"github.com/emicklei/go-restful"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// JSONStatus json status
type JSONStatus struct {
	hFunc restful.RouteFunction
}

// NewJSONStatus new json status object
func NewJSONStatus(hf restful.RouteFunction) Resource {
	return &JSONStatus{
		hFunc: hf,
	}
}

// Register implements metric.Resource interface
func (js *JSONStatus) Register(container *restful.Container) {
	blog.Infof("register json status resource to metric")
	ws := new(restful.WebService)
	ws.Path("/status").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(ws.GET("/").To(js.hFunc))
	container.Add(ws)
}
