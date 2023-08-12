

package native

import (
	"errors"

	corev1 "k8s.io/api/core/v1"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// NativeNodeClient default node exporter
type NativeNodeClient struct {
}

// NewNativeNodeClient return new native node client
func NewNativeNodeClient() *NativeNodeClient {
	return &NativeNodeClient{}
}

// GetNodeExternalIpList 获取node.Status.Addresses下的ExternalIP
func (n *NativeNodeClient) GetNodeExternalIpList(node *corev1.Node) ([]string, error) {
	externalIpList := make([]string, 0)

	for _, addr := range node.Status.Addresses {
		if addr.Type == corev1.NodeExternalIP {
			externalIpList = append(externalIpList, addr.Address)
		}
	}

	if len(externalIpList) == 0 {
		return nil, errors.New("empty node external ip list")
	}

	blog.Infof("get node %s ip list: %v", node.Name, externalIpList)
	return externalIpList, nil
}
