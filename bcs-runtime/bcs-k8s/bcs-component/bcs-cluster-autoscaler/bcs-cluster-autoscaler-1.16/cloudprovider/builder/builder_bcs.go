//go:build clusterapi
// +build clusterapi



package builder

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/config"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/cloudprovider/bcs"
)

// AvailableCloudProviders supported by the cloud provider builder.
var AvailableCloudProviders = []string{
	bcs.ProviderName,
}

// DefaultCloudProvider is tenc.
const DefaultCloudProvider = bcs.ProviderName

func buildCloudProvider(opts config.AutoscalingOptions, do cloudprovider.NodeGroupDiscoveryOptions,
	rl *cloudprovider.ResourceLimiter) cloudprovider.CloudProvider {
	switch opts.CloudProviderName {
	case bcs.ProviderName:
		return bcs.BuildCloudProvider(opts, do, rl)
	}
	return nil
}
