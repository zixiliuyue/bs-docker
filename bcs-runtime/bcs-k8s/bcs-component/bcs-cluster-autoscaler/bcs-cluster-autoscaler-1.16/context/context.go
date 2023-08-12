

// Package context xxx
package context

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/context"
	"k8s.io/autoscaler/cluster-autoscaler/expander"
	processor_callbacks "k8s.io/autoscaler/cluster-autoscaler/processors/callbacks"
	"k8s.io/autoscaler/cluster-autoscaler/simulator"

	estimatorinternal "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/estimator"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/scalingconfig"
)

// Context contains user-configurable constant and configuration-related objects passed to
// scale up/scale down functions.
type Context struct {
	*context.AutoscalingContext
	ExtendedEstimatorBuilder estimatorinternal.ExtendedEstimatorBuilder
}

// NewAutoscalingContext returns an autoscaling context from all the necessary parameters passed via arguments
func NewAutoscalingContext(options scalingconfig.Options, predicateChecker *simulator.PredicateChecker,
	autoscalingKubeClients *context.AutoscalingKubeClients, cloudProvider cloudprovider.CloudProvider,
	expanderStrategy expander.Strategy, estimatorBuilder estimatorinternal.ExtendedEstimatorBuilder,
	processorCallbacks processor_callbacks.ProcessorCallbacks) *Context {
	return &Context{
		AutoscalingContext: &context.AutoscalingContext{
			AutoscalingOptions:     options.AutoscalingOptions,
			CloudProvider:          cloudProvider,
			AutoscalingKubeClients: *autoscalingKubeClients,
			PredicateChecker:       predicateChecker,
			ExpanderStrategy:       expanderStrategy,
			ProcessorCallbacks:     processorCallbacks,
		},
		ExtendedEstimatorBuilder: estimatorBuilder,
	}
}
