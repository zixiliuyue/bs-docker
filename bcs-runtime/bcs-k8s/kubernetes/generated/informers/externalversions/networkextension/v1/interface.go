
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Ingresses returns a IngressInformer.
	Ingresses() IngressInformer
	// Listeners returns a ListenerInformer.
	Listeners() ListenerInformer
	// PortBindings returns a PortBindingInformer.
	PortBindings() PortBindingInformer
	// PortPools returns a PortPoolInformer.
	PortPools() PortPoolInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Ingresses returns a IngressInformer.
func (v *version) Ingresses() IngressInformer {
	return &ingressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Listeners returns a ListenerInformer.
func (v *version) Listeners() ListenerInformer {
	return &listenerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PortBindings returns a PortBindingInformer.
func (v *version) PortBindings() PortBindingInformer {
	return &portBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PortPools returns a PortPoolInformer.
func (v *version) PortPools() PortPoolInformer {
	return &portPoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
