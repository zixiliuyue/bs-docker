

package filter

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-driver/mesosdriver/config"

	"github.com/emicklei/go-restful"
)

// HeaderValidFilter for BCS header BCS-ClusterID
type HeaderValidFilter struct {
	conf *config.MesosDriverConfig
}

// NewHeaderValidFilter create header instance
func NewHeaderValidFilter(conf *config.MesosDriverConfig) RequestFilterFunction {
	return &HeaderValidFilter{
		conf: conf,
	}
}

// Execute check header BCS-ClusterID
func (h *HeaderValidFilter) Execute(req *restful.Request) (int, error) {
	clusterID := req.Request.Header.Get("BCS-ClusterID")
	if clusterID != h.conf.Cluster {
		blog.Errorf("request %s lost BCS-ClusterID header, detail: %+v", req.Request.URL.Path, req.Request.Header)
		return common.BcsErrMesosDriverHttpFilterFailed, fmt.Errorf("http header BCS-ClusterID %s don't exist", clusterID)
	}

	return 0, nil
}
