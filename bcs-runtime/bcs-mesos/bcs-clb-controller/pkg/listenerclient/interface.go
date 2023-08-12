

package listenerclient

import (
	cloudListenerType "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/network/v1"
)

// Interface xxx
type Interface interface {
	ListListeners() ([]*cloudListenerType.CloudListener, error)
	Create(listener *cloudListenerType.CloudListener) error
	Update(listener *cloudListenerType.CloudListener) error
	Delete(listener *cloudListenerType.CloudListener) error
}
