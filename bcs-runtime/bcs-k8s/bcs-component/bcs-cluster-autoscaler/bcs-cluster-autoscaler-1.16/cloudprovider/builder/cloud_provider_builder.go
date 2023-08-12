

package builder

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/context"
	"k8s.io/klog"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/scalingconfig"
)

// NewCloudProvider builds a cloud provider from provided parameters.
func NewCloudProvider(opts scalingconfig.Options) cloudprovider.CloudProvider {
	klog.V(1).Infof("Building %s cloud provider.", opts.CloudProviderName)

	do := cloudprovider.NodeGroupDiscoveryOptions{
		NodeGroupSpecs:              opts.NodeGroups,
		NodeGroupAutoDiscoverySpecs: opts.NodeGroupAutoDiscovery,
	}

	rl := context.NewResourceLimiterFromAutoscalingOptions(opts.AutoscalingOptions)

	if opts.CloudProviderName == "" {
		// Ideally this would be an error, but several unit tests of the
		// StaticAutoscaler depend on this behaviour.
		klog.Warning("Returning a nil cloud provider")
		return nil
	}

	provider := buildCloudProvider(opts, do, rl)
	if provider != nil {
		return provider
	}

	klog.Fatalf("Unknown cloud provider: %s", opts.CloudProviderName)
	return nil // This will never happen because the Fatalf will os.Exit
}
