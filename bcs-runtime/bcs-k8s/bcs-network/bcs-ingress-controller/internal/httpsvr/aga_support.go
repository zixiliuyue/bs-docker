

package httpsvr

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-network/bcs-ingress-controller/internal/metrics"
)

func (h *HttpServerClient) getPodRelatedAgaEntrance(request *restful.Request, response *restful.Response) {
	startTime := time.Now()
	mf := func(status string) {
		metrics.ReportAPIRequestMetric("aga_entrance", "GET", status, startTime)
	}
	if h.AgaSupporter == nil {
		_, _ = response.Write(CreateResponseData(errors.New("Not init, check if select aws as cloud provider"), "",
			nil))
		mf(strconv.Itoa(http.StatusInternalServerError))
		return
	}
	podName := request.QueryParameter("pod_name")
	podNamespace := request.QueryParameter("pod_namespace")
	if podName == "" || podNamespace == "" {
		_, _ = response.Write(CreateResponseData(errors.New("empty parameter: pod_name or pod_namespace"), "", nil))
		mf(strconv.Itoa(http.StatusInternalServerError))
		return
	}

	region := request.QueryParameter("region")
	if region == "" {
		blog.V(4).Infof("empty region parameter, use default region ")
		region = h.Ops.Region
	}

	blog.V(3).Infof("getPodRelatedAgaEntrance req[pod_name='%s', pod_namespace='%s', region='%s']", podName,
		podNamespace, region)

	pod := &v1.Pod{}
	if err := h.Mgr.GetClient().Get(context.TODO(), types.NamespacedName{
		Namespace: podNamespace,
		Name:      podName,
	}, pod); err != nil {
		err = errors.Wrapf(err, "get pods '%s/%s' failed", podNamespace, podName)
		blog.Errorf(err.Error())
		_, _ = response.Write(CreateResponseData(err, "", nil))
		mf(strconv.Itoa(http.StatusInternalServerError))
		return
	}

	// get node from pod's spec
	node := &v1.Node{}
	if err := h.Mgr.GetClient().Get(context.TODO(), types.NamespacedName{
		Name: pod.Spec.NodeName,
	}, node); err != nil {
		err = errors.Wrapf(err, "get node '%s' failed", pod.Spec.NodeName)
		blog.Errorf(err.Error())
		_, _ = response.Write(CreateResponseData(err, "", nil))
		mf(strconv.Itoa(http.StatusInternalServerError))
		return
	}

	resp, err := h.AgaSupporter.ListCustomRoutingByDefinition(region, getNodeInstanceID(node), pod.Status.PodIP)
	if err != nil {
		_, _ = response.Write(CreateResponseData(err, "", nil))
		mf(strconv.Itoa(http.StatusInternalServerError))
		return
	}

	_, _ = response.Write(CreateResponseData(nil, "", resp))
	mf(strconv.Itoa(http.StatusOK))
}

// aws中， node.Spec.ProviderID的格式为aws:///<availability_zone>/<instance_id>
// 如， aws:///us-west-2a/i-0abcdef1234567890
func getNodeInstanceID(node *v1.Node) string {
	providerID := node.Spec.ProviderID
	parts := strings.Split(providerID, "/")
	instanceID := parts[len(parts)-1]
	return instanceID
}
