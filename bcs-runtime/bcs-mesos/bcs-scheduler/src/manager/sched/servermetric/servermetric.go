

// Package servermetric xxx
package servermetric

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/metric"
	"sync"
)

var (
	// Metrics xxx
	Metrics    ServerMetric
	metricLock sync.Mutex
)

// ServerMetric xxx
type ServerMetric struct {
	role        metric.RoleType
	mesosMaster string
}

// SetRole xxx
func SetRole(role metric.RoleType) {
	metricLock.Lock()
	defer metricLock.Unlock()

	Metrics.role = role
}

// SetMesosMaster xxx
func SetMesosMaster(master string) {
	metricLock.Lock()
	defer metricLock.Unlock()

	Metrics.mesosMaster = master
}

// GetRole xxx
func GetRole() metric.RoleType {
	metricLock.Lock()
	defer metricLock.Unlock()

	return Metrics.role
}

// IsHealthy xxx
func IsHealthy() (bool, string) {
	metricLock.Lock()
	defer metricLock.Unlock()

	if Metrics.role != metric.MasterRole && Metrics.role != metric.SlaveRole {
		return false, "scheduler unknown role type:" + string(Metrics.role)
	}

	if Metrics.mesosMaster == "" {
		return false, "no mesos master"
	}

	return true, "run ok"
}
