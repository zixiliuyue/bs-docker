

package core

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/clusterstate"
	"k8s.io/autoscaler/cluster-autoscaler/context"
	"k8s.io/autoscaler/cluster-autoscaler/core"
	"k8s.io/autoscaler/cluster-autoscaler/expander"
	"k8s.io/autoscaler/cluster-autoscaler/expander/factory"
	ca_processors "k8s.io/autoscaler/cluster-autoscaler/processors"
	"k8s.io/autoscaler/cluster-autoscaler/simulator"
	"k8s.io/autoscaler/cluster-autoscaler/utils/backoff"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	kube_client "k8s.io/client-go/kubernetes"

	cloudBuilder "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/cloudprovider/builder"
	estimatorinternal "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/estimator"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/scalingconfig"
)

// AutoscalerOptions is the whole set of options for configuring an autoscaler
type AutoscalerOptions struct {
	scalingconfig.Options
	KubeClient             kube_client.Interface
	EventsKubeClient       kube_client.Interface
	AutoscalingKubeClients *context.AutoscalingKubeClients
	CloudProvider          cloudprovider.CloudProvider
	PredicateChecker       simulator.PredicateChecker
	ClusterSnapshot        simulator.ClusterSnapshot
	ExpanderStrategy       expander.Strategy
	EstimatorBuilder       estimatorinternal.ExtendedEstimatorBuilder
	Processors             *ca_processors.AutoscalingProcessors
	Backoff                backoff.Backoff
}

// NewAutoscaler creates an autoscaler of an appropriate type according to the parameters
func NewAutoscaler(opts AutoscalerOptions) (core.Autoscaler, errors.AutoscalerError) {
	err := initializeDefaultOptions(&opts)
	if err != nil {
		return nil, errors.ToAutoscalerError(errors.InternalError, err)
	}
	return NewBufferedAutoscaler(
		opts.Options,
		opts.PredicateChecker,
		opts.ClusterSnapshot,
		opts.AutoscalingKubeClients,
		opts.Processors,
		opts.CloudProvider,
		opts.ExpanderStrategy,
		opts.EstimatorBuilder,
		opts.Backoff), nil
}

// initializeDefaultOptions initialize default options if not provided.
func initializeDefaultOptions(opts *AutoscalerOptions) error {
	if opts.Processors == nil {
		opts.Processors = ca_processors.DefaultProcessors()
	}
	if opts.AutoscalingKubeClients == nil {
		opts.AutoscalingKubeClients = context.NewAutoscalingKubeClients(opts.Options.AutoscalingOptions,
			opts.KubeClient, opts.EventsKubeClient)
	}
	if opts.PredicateChecker == nil {
		predicateCheckerStopChannel := make(chan struct{})
		predicateChecker, err := simulator.NewSchedulerBasedPredicateChecker(opts.KubeClient,
			predicateCheckerStopChannel)
		if err != nil {
			return err
		}
		opts.PredicateChecker = predicateChecker
	}
	if opts.ClusterSnapshot == nil {
		opts.ClusterSnapshot = simulator.NewBasicClusterSnapshot()
	}
	if opts.CloudProvider == nil {
		opts.CloudProvider = cloudBuilder.NewCloudProvider(opts.Options)
	}
	if opts.ExpanderStrategy == nil {
		expanderStrategy, err := factory.ExpanderStrategyFromString(opts.ExpanderName,
			opts.CloudProvider, opts.AutoscalingKubeClients, opts.KubeClient, opts.ConfigNamespace)
		if err != nil {
			return err
		}
		opts.ExpanderStrategy = expanderStrategy
	}
	if opts.EstimatorBuilder == nil {
		estimatorBuilder, err := estimatorinternal.NewEstimatorBuilder(opts.EstimatorName,
			opts.Options.BufferedCPURatio,
			opts.Options.BufferedMemRatio,
			opts.Options.BufferedResourceRatio)
		if err != nil {
			return err
		}
		opts.EstimatorBuilder = estimatorBuilder
	}
	if opts.Backoff == nil {
		opts.Backoff =
			backoff.NewIdBasedExponentialBackoff(clusterstate.InitialNodeGroupBackoffDuration,
				clusterstate.MaxNodeGroupBackoffDuration, clusterstate.NodeGroupBackoffResetTimeout)
	}

	return nil
}
