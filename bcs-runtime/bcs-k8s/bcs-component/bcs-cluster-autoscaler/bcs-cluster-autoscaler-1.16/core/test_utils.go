

package core

import (
	"k8s.io/autoscaler/cluster-autoscaler/processors"
	"k8s.io/autoscaler/cluster-autoscaler/processors/nodegroups"
	"k8s.io/autoscaler/cluster-autoscaler/processors/nodegroupset"
	"k8s.io/autoscaler/cluster-autoscaler/processors/status"
)

// NewTestProcessors returns a set of simple processors for use in tests.
func NewTestProcessors() *processors.AutoscalingProcessors {
	return &processors.AutoscalingProcessors{
		PodListProcessor:       NewFilterOutSchedulablePodListProcessor(),
		NodeGroupListProcessor: &nodegroups.NoOpNodeGroupListProcessor{},
		NodeGroupSetProcessor:  &nodegroupset.BalancingNodeGroupSetProcessor{},
		// DOTO(bskiba): change scale up test so that this can be a NoOpProcessor
		ScaleUpStatusProcessor:     &status.EventingScaleUpStatusProcessor{},
		ScaleDownStatusProcessor:   &status.NoOpScaleDownStatusProcessor{},
		AutoscalingStatusProcessor: &status.NoOpAutoscalingStatusProcessor{},
		NodeGroupManager:           nodegroups.NewDefaultNodeGroupManager(),
	}
}
