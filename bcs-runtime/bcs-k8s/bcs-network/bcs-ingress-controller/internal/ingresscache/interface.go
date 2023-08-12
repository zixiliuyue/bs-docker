

package ingresscache

import (
	networkextensionv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
)

// IngressCache store relationship about ingress
type IngressCache interface {
	Add(ingress *networkextensionv1.Ingress)
	Remove(ingress *networkextensionv1.Ingress)
	GetRelatedIngressOfService(serviceNamespace, serviceName string) []IngressMeta
	GetRelatedIngressOfWorkload(workloadKind, workloadNamespace, workloadName string) []IngressMeta
}

// IngressMeta meta info of ingress
type IngressMeta struct {
	Namespace string
	Name      string
}
