

package custom

import (
	"fmt"

	"github.com/gofrs/uuid"

	"github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// ServiceNode is node info for bcs services.
type ServiceNode types.ServerInfo

// NewServiceNode create kubedriver service node
func NewServiceNode(info types.ServerInfo) ServiceNode {
	return ServiceNode{
		IP:           info.IP,
		Port:         info.Port,
		ExternalIp:   info.ExternalIp,
		ExternalPort: info.ExternalPort,
		MetricPort:   info.MetricPort,
		HostName:     info.HostName,
		Scheme:       info.Scheme,
		Version:      info.Version,
		Cluster:      info.Cluster,
		Pid:          info.Pid,
	}
}

// PrimaryKey key for indexer
func (n *ServiceNode) PrimaryKey() string {
	u, _ := uuid.NewV4()
	return fmt.Sprintf("%s", u.String())
}

// Payload content length
func (n *ServiceNode) Payload() []byte {
	result, _ := json.Marshal(n)
	return result
}

// OwnsPayload xxxx
func (n *ServiceNode) OwnsPayload(payload []byte) bool {
	return true
}
