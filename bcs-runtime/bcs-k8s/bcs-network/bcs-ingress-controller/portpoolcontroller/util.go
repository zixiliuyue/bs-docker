

package portpoolcontroller

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-network/bcs-ingress-controller/internal/common"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-network/bcs-ingress-controller/internal/constant"
	netextv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
)

// return true if listener's labels are correct
func checkListenerLabels(labels map[string]string, portPoolName, itemName string) bool {
	poolNameLabel := common.GetPortPoolListenerLabelKey(portPoolName, itemName)
	if v, ok := labels[poolNameLabel]; !ok || v != netextv1.LabelValueForPortPoolItemName {
		return false
	}

	if v, ok := labels[netextv1.LabelKeyForOwnerName]; !ok || v != portPoolName {
		return false
	}

	return true
}

func checkPortPoolStatus(pool *netextv1.PortPool) string {
	statusReady := true
	for _, ts := range pool.Status.PoolItemStatuses {
		if ts.Status != constant.PortPoolItemStatusReady {
			statusReady = false
			break
		}
	}
	if statusReady {
		return constant.PortPoolStatusReady
	}
	return constant.PortPoolStatusNotReady
}
