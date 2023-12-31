

package nodeclient

import (
	"fmt"
	"time"

	"k8s.io/client-go/informers"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	k8scorev1 "k8s.io/api/core/v1"
	k8slabels "k8s.io/apimachinery/pkg/labels"
	k8sinfcorev1 "k8s.io/client-go/informers/core/v1"
	k8scorecliset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

// KubeClient xxx
type KubeClient struct {
	nodesInformer k8sinfcorev1.NodeInformer
	nodesCache    cache.Store
	stopCh        chan struct{}
}

// NewKubeClient xxx
func NewKubeClient(kubeconfig string, handler cache.ResourceEventHandler, syncPeriod time.Duration) (Client, error) {
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		blog.Errorf("create internal client with kubeconfig %s failed, %s", kubeconfig, err.Error())
		return nil, fmt.Errorf("create internal client with kubeconfig %s failed, %s", kubeconfig, err.Error())
	}
	cliset, err := k8scorecliset.NewForConfig(restConfig)
	if err != nil {
		blog.Errorf("create clientset failed, with rest config %v, err %s", restConfig, err.Error())
		return nil, fmt.Errorf("create clientset failed, with rest config %v, err %s", restConfig, err.Error())
	}
	blog.Infof("start create informer factory")
	factory := informers.NewSharedInformerFactory(cliset, syncPeriod)
	nodesInformer := factory.Core().V1().Nodes()
	nodesCache := nodesInformer.Informer().GetStore()

	nodesInformer.Informer().AddEventHandler(handler)

	client := &KubeClient{
		nodesInformer: nodesInformer,
		nodesCache:    nodesCache,
		stopCh:        make(chan struct{}),
	}

	go nodesInformer.Informer().Run(client.stopCh)

	return client, nil
}

// ListNodes xxx
func (kc *KubeClient) ListNodes() ([]*Node, error) {
	selector := k8slabels.Everything()
	nodes, err := kc.nodesInformer.Lister().List(selector)
	if err != nil {
		blog.Errorf("list k8s nodes failed, err %s", err.Error())
		return nil, fmt.Errorf("list k8s nodes failed, err %s", err.Error())
	}

	var retNodes []*Node
	for _, node := range nodes {
		if len(node.Status.Addresses) != 0 {
			var internalIPs []string
			for _, addrStruct := range node.Status.Addresses {
				if addrStruct.Type == k8scorev1.NodeInternalIP {
					internalIPs = append(internalIPs, addrStruct.Address)
				}
			}
			if len(internalIPs) != 0 {
				retNodes = append(retNodes, &Node{
					IPs:           internalIPs,
					Unschedulable: node.Spec.Unschedulable,
				})
			}
		}
	}
	return retNodes, nil
}

// Close xxx
func (kc *KubeClient) Close() {
	close(kc.stopCh)
}
