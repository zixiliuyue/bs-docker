
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/common/bcs-hook/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// HookRuns returns a HookRunInformer.
	HookRuns() HookRunInformer
	// HookTemplates returns a HookTemplateInformer.
	HookTemplates() HookTemplateInformer
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

// HookRuns returns a HookRunInformer.
func (v *version) HookRuns() HookRunInformer {
	return &hookRunInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HookTemplates returns a HookTemplateInformer.
func (v *version) HookTemplates() HookTemplateInformer {
	return &hookTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
