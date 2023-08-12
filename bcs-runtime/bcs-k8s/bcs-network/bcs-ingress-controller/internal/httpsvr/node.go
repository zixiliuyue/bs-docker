

package httpsvr

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/emicklei/go-restful"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-network/bcs-ingress-controller/internal/metrics"
)

func (h *HttpServerClient) listNode(request *restful.Request, response *restful.Response) {
	startTime := time.Now()
	mf := func(status string) {
		metrics.ReportAPIRequestMetric("list_node", "GET", status, startTime)
	}
	nodeName := request.QueryParameter("node_name")
	nodeInternalIP := request.QueryParameter("node_ip")
	if nodeName == "" && nodeInternalIP == "" {
		_, _ = response.Write(CreateResponseData(errors.New("empty parameter: node_name and node_ip"), "", nil))
		mf(strconv.Itoa(http.StatusInternalServerError))
		return
	}

	var nodeIPs []string
	var err error
	if nodeName != "" {
		nodeIPs, err = h.NodeCache.GetNodeExternalIPsByName(nodeName)
		if err != nil {
			_, _ = response.Write(CreateResponseData(err, "", nil))
			mf(strconv.Itoa(http.StatusInternalServerError))
			return
		}
	}
	if nodeInternalIP != "" {
		nodeIPs, err = h.NodeCache.GetNodeExternalIPsByIP(nodeInternalIP)
		if err != nil {
			_, _ = response.Write(CreateResponseData(err, "", nil))
			mf(strconv.Itoa(http.StatusInternalServerError))
			return
		}
	}
	mf(strconv.Itoa(http.StatusOK))
	_, _ = response.Write(CreateResponseData(nil, "", nodeIPs))
}
