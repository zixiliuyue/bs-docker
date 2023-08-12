

// Package cluster xxx
package cluster

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-watch/types"

	"golang.org/x/net/context"
)

// DataExister checker interface for data exist
type DataExister interface {
	IsExist(data interface{}) bool
}

// Reporter for report data
type Reporter interface {
	ReportData(data *types.BcsSyncData) error
	GetClusterID() string
}

// Cluster is interface for reading Cluster info
type Cluster interface {
	Run(cxt context.Context)  // start cluster
	Sync(tp string) error     // ready to sync data, type like services, pods and etc.
	Stop()                    // stop cluster
	GetClusterStatus() string // get curr status
}

// EventHandler hold event interface for All Watch
type EventHandler interface {
	AddEvent(obj interface{})
	DeleteEvent(obj interface{})
	UpdateEvent(old, cur interface{})
}

// ReportFunc report function for handle detail report data
type ReportFunc func(data *types.BcsSyncData) error
