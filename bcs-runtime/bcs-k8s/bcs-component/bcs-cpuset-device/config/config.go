

// Package config xxx
package config

import "github.com/Tencent/bk-bcs/bcs-common/common/types"

// Config parsed config for bcs-cpuset-device
type Config struct {
	// PluginSocketDir device plugin socket dir, examples: /var/lib/kubelet/device-plugins
	PluginSocketDir string
	// DockerSocket docker socket
	DockerSocket string
	// CgroupCpusetRoot root path of cpuset cgroup
	CgroupCpusetRoot string
	// ClientCert client https certs
	ClientCert *types.CertConfig `json:"-"`
	// BcsZk cluster zk address
	BcsZk string
	// ClusterID clusterid
	ClusterID string
	// Engine engine type enum: k8s、mesos
	Engine string
	// NodeIP node IP
	NodeIP string
	// ReservedCPUSet set of reserved cpu set
	ReservedCPUSet map[string]struct{}
}

// NewConfig create a config object
func NewConfig() *Config {
	return &Config{
		ClientCert: &types.CertConfig{},
	}
}
