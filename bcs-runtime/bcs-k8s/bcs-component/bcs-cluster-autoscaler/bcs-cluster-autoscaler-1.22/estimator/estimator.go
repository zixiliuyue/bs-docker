

// Package estimator xxx
package estimator

import (
	"fmt"

	"k8s.io/autoscaler/cluster-autoscaler/estimator"
	"k8s.io/autoscaler/cluster-autoscaler/simulator"
)

const (
	// BinpackingEstimatorName is the name of binpacking estimator.
	BinpackingEstimatorName = "binpacking"
	// ResourceEstimatorName is the name of the resource estimator.
	ResourceEstimatorName = "clusterresource"
)

// AvailableEstimators is a list of available estimators.
var AvailableEstimators = []string{BinpackingEstimatorName, ResourceEstimatorName}

// ExtendedEstimatorBuilder creates a new estimator object.
type ExtendedEstimatorBuilder func(simulator.PredicateChecker, simulator.ClusterSnapshot) estimator.Estimator

// NewEstimatorBuilder creates a new estimator object from flag.
func NewEstimatorBuilder(name string,
	cpuRatio, memRatio, resourceRatio float64) (ExtendedEstimatorBuilder, error) {
	switch name {
	case BinpackingEstimatorName:
		return func(
			predicateChecker simulator.PredicateChecker,
			clusterSnapshot simulator.ClusterSnapshot) estimator.Estimator {
			return estimator.NewBinpackingNodeEstimator(predicateChecker, clusterSnapshot)
		}, nil
	case ResourceEstimatorName:
		return func(predicateChecker simulator.PredicateChecker,
			clusterSnapshot simulator.ClusterSnapshot,
		) estimator.Estimator {
			return NewClusterResourceEstimator(predicateChecker, clusterSnapshot,
				cpuRatio, memRatio, resourceRatio)
		}, nil
	}
	return nil, fmt.Errorf("unknown estimator: %s", name)
}
