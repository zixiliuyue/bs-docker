

// Package types xxx
package types

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/registry"
)

const (
	// ActionAdd add event
	ActionAdd = "Add"
	// ActionDelete delete event
	ActionDelete = "Delete"
	// ActionUpdate update event
	ActionUpdate = "Update"
)

// BcsSyncData holder for sync data
type BcsSyncData struct {
	DataType string      // data type: reflect.TypeOf(Item).Name()
	Action   string      // operation, like Add, Delete, Update
	Item     interface{} // SyncData, data is Endpoint, Service, Pod
}

// CmdConfig hold all command line config item
type CmdConfig struct {
	ClusterID   string
	ClusterInfo string
	IsExternal  bool
	CAFile      string
	CertFile    string
	KeyFile     string
	PassWord    string

	RegDiscvSvr            string
	Address                string
	ApplicationThreadNum   int
	TaskgroupThreadNum     int
	ExportserviceThreadNum int
	DeploymentThreadNum    int

	MetricPort uint

	ServerCAFile   string
	ServerCertFile string
	ServerKeyFile  string
	ServerPassWord string
	ServerSchem    string

	KubeConfig  string
	StoreDriver string

	// NetServiceZK is zookeeper address config for netservice discovery,
	// reuse RegDiscvSvr by default.
	NetServiceZK string

	// Etcd etcd options for service registry and discovery
	Etcd registry.CMDOptions

	// StorageAddresses address for bcs-storage
	StorageAddresses []string
}

const (
	// ApplicationChannelPrefix prefix for event post channel
	ApplicationChannelPrefix = "Application_"
	// TaskgroupChannelPrefix prefix for event post channel
	TaskgroupChannelPrefix = "TaskGroup_"
	// ExportserviceChannelPrefix prefix for event post channel
	ExportserviceChannelPrefix = "Exportservice_"
	// DeploymentChannelPrefix deployment prefix for post channel
	DeploymentChannelPrefix = "Deployment_"
)
