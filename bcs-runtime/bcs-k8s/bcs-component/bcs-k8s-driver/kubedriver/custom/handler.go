

package custom

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/kubedriver/options"

	"github.com/emicklei/go-restful"
)

// APIHandler http API interface for info request
type APIHandler interface {
	Handler(request *restful.Request, response *restful.Response)
	Config(k8sMasterUrl string, tlsCfg options.TLSConfig) error
}
