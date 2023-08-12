

package builder

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/cloudprovider/bcs"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/scalingconfig"
)

// AvailableCloudProviders supported by the cloud provider builder.
var AvailableCloudProviders = []string{
	bcs.ProviderName,
}

// DefaultCloudProvider is bcs.
const DefaultCloudProvider = bcs.ProviderName

func buildCloudProvider(opts scalingconfig.Options, do cloudprovider.NodeGroupDiscoveryOptions,
	rl *cloudprovider.ResourceLimiter) cloudprovider.CloudProvider {
	switch opts.CloudProviderName {
	case bcs.ProviderName:
		return bcs.BuildCloudProvider(opts, do, rl)
	}
	return nil
}
