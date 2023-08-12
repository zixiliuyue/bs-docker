

package clbingress

import (
	ingressv1 "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/clb/v1"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-clb-controller/pkg/model"
)

// Registry xxx
// IngressRegistry interface for clb ingress rule discovery
type Registry interface {
	AddIngressHandler(handler model.EventHandler)
	ListIngresses() ([]*ingressv1.ClbIngress, error)
	GetIngress(name string) (*ingressv1.ClbIngress, error)
	SetIngress(*ingressv1.ClbIngress) error
}
