

package resthdrs

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/metric"
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/server/resthdrs/filters"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/storages/sqlstore"

	"github.com/emicklei/go-restful"
	"time"
)

// ListRegisterTokens xxx
func ListRegisterTokens(request *restful.Request, response *restful.Response) {

	start := time.Now()

	cluster := filters.GetCluster(request)

	token := sqlstore.GetRegisterToken(cluster.ID)
	if token == nil {
		metric.RequestErrorCount.WithLabelValues("k8s_rest", request.Request.Method).Inc()
		metric.RequestErrorLatency.WithLabelValues("k8s_rest", request.Request.Method).Observe(time.Since(start).Seconds())
		message := fmt.Sprintf("errcode: %d, register token not found", common.BcsErrApiBadRequest)
		WriteClientError(response, "RTOKEN_NOT_FOUND", message)
		return
	}
	response.WriteEntity([]*m.RegisterToken{token})

	metric.RequestCount.WithLabelValues("k8s_rest", request.Request.Method).Inc()
	metric.RequestLatency.WithLabelValues("k8s_rest", request.Request.Method).Observe(time.Since(start).Seconds())
}

// CreateRegisterToken xxx
func CreateRegisterToken(request *restful.Request, response *restful.Response) {

	start := time.Now()

	cluster := filters.GetCluster(request)
	clusterId := cluster.ID

	err := sqlstore.CreateRegisterToken(clusterId)
	if err != nil {
		metric.RequestErrorCount.WithLabelValues("k8s_rest", request.Request.Method).Inc()
		metric.RequestErrorLatency.WithLabelValues("k8s_rest", request.Request.Method).Observe(time.Since(start).Seconds())
		message := fmt.Sprintf("errcode: %d, can not create register token: %s", common.BcsErrApiBadRequest, err.Error())
		WriteServerError(response, "CANNOT_CREATE_RTOKEN", message)
		return
	}
	response.WriteEntity([]*m.RegisterToken{
		sqlstore.GetRegisterToken(clusterId),
	})

	metric.RequestCount.WithLabelValues("k8s_rest", request.Request.Method).Inc()
	metric.RequestLatency.WithLabelValues("k8s_rest", request.Request.Method).Observe(time.Since(start).Seconds())
}
