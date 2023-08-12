

package cloudnode

import (
	corev1 "k8s.io/api/core/v1"
)

// NodeClient node exporter
type NodeClient interface {
	// GetNodeExternalIpList get node external IP list
	GetNodeExternalIpList(node *corev1.Node) ([]string, error)
}
