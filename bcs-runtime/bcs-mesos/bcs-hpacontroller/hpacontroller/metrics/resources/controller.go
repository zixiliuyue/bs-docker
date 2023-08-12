

package resources

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"sync"

	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/hpacontroller/config"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/hpacontroller/metrics"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/hpacontroller/reflector"
)

type resourceMetrics struct {
	sync.RWMutex
	// hpa controller config
	config *config.Config

	// Reflector watches a specified resource and causes all changes to be reflected in the given store
	store reflector.Reflector

	// hpa autoscaler work queue, key = BcsAutoscaler.GetUuid()
	workQueue map[string]*resourcesCollector
}

// NewResourceMetrics xxx
func NewResourceMetrics(conf *config.Config, store reflector.Reflector) metrics.MetricsController {
	resources := &resourceMetrics{
		config:    conf,
		store:     store,
		workQueue: make(map[string]*resourcesCollector),
	}

	return resources
}

// StartScalerMetrics xxx
// start to collect scaler metrics
func (resources *resourceMetrics) StartScalerMetrics(scaler *commtypes.BcsAutoscaler) {
	resources.Lock()
	defer resources.Unlock()

	_, ok := resources.workQueue[scaler.GetUuid()]
	if ok {
		return
	}

	// start collector scaler target ref resources metrics
	resources.workQueue[scaler.GetUuid()] = newResourcesCollector(resources, scaler)
	resources.workQueue[scaler.GetUuid()].start()
	blog.Infof("start collector scaler %s resources metrics", scaler.GetUuid())
}

// StopScalerMetrics xxx
// stop to collect scaler metrics
func (resources *resourceMetrics) StopScalerMetrics(scaler *commtypes.BcsAutoscaler) {
	resources.Lock()
	defer resources.Unlock()

	collector, ok := resources.workQueue[scaler.GetUuid()]
	if !ok {
		return
	}
	collector.stop()
	blog.Infof("stop collector scaler %s resources metrics", scaler.GetUuid())
}

// GetResourceMetric gets the given resource metric (and an associated oldest timestamp)
// for all taskgroup matching the specified scaler uuid
func (resources *resourceMetrics) GetResourceMetric(resourceName, uuid string) (metrics.TaskgroupMetricsInfo, error) {
	resources.RLock()
	defer resources.RUnlock()

	collector, ok := resources.workQueue[uuid]
	if !ok {
		return nil, fmt.Errorf("sacler %s not found", uuid)
	}

	switch resourceName {
	case metrics.TaskgroupResourcesCpuMetricsName:
		return collector.getCpuMetricsInfo(), nil

	case metrics.TaskgroupResourcesMemoryMetricsName:
		return collector.getMemoryMetricsInfo(), nil

	}

	return nil, fmt.Errorf("resource name %s is invalid", resourceName)
}
