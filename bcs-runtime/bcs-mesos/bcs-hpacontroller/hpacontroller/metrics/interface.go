

package metrics

import (
	"time"

	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// TaskgroupMetric contains pod metric value (the metric values are expected to be the metric as a milli-value)
type TaskgroupMetric struct {
	Timestamp time.Time
	Window    int // seconds
	Value     float32
}

const (
	// TaskgroupResourcesCpuMetricsName TODO
	TaskgroupResourcesCpuMetricsName = "cpu"
	// TaskgroupResourcesMemoryMetricsName TODO
	TaskgroupResourcesMemoryMetricsName = "memory"
)

// TaskgroupMetricsInfo contains taskgroup metrics as a map from pod names to TaskgroupMetricsInfo
type TaskgroupMetricsInfo map[string]TaskgroupMetric

// MetricsController collect external metrics or taskgroup resource metrics
type MetricsController interface {
	// StartScalerMetrics TODO
	// start to collect scaler metrics
	StartScalerMetrics(scaler *commtypes.BcsAutoscaler)

	// StopScalerMetrics TODO
	// stop to collect scaler metrics
	StopScalerMetrics(scaler *commtypes.BcsAutoscaler)

	// GetResourceMetric gets the given resource metric (and an associated oldest timestamp)
	// for all taskgroup matching the specified uuid
	GetResourceMetric(resourceName, uuid string) (TaskgroupMetricsInfo, error)
}
