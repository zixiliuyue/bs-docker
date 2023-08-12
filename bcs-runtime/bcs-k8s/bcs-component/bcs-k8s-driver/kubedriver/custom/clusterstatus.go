

package custom

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/client"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/kubedriver/options"

	"github.com/emicklei/go-restful"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ClusterStatusAPIHandler cluster node http api implementation
type ClusterStatusAPIHandler struct {
	clientSet *kubernetes.Clientset
}

// Handler http implementation
func (csh *ClusterStatusAPIHandler) Handler(request *restful.Request, response *restful.Response) {
	nodes, err := csh.clientSet.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		CustomServerErrorResponse(response, "Get node list failed")
		return
	}

OutLoop:
	for _, node := range nodes.Items {
		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" && condition.Status != "True" {
				continue OutLoop
			}
		}
	}
}

// Config config kube clientset
func (csh *ClusterStatusAPIHandler) Config(k8sMasterUrl string, tlsCfg options.TLSConfig) error {
	csh.clientSet = client.NewClientSet(k8sMasterUrl, tlsCfg)
	if csh.clientSet == nil {
		return fmt.Errorf("failed to get k8s clientSet")
	}
	return nil
}
