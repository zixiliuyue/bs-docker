

// Package options xxx
package options

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// Option is option in flags
type Option struct {
	conf.FileConfig
	conf.LicenseServerConfig
	conf.LogConfig
	conf.ZkConfig
	conf.CertConfig
	conf.ServiceConfig

	DockerSock         string `json:"docker_sock" value:"unix:///var/run/docker.sock" usage:"docker socket file"`
	PluginSocketDir    string `json:"plugin_socket_dir" value:"/var/lib/kubelet/device-plugins" usage:"device-plugin socket directory"`
	CgroupCpusetRoot   string `json:"cgroup_cpuset_root" value:"/sys/fs/cgroup/cpuset/docker" usage:"root path of cgroup cpuset"`
	ClusterID          string `json:"clusterid" value:"" usage:"mesos cluster id"`
	Engine             string `json:"engine" value:"k8s" usage:"enum: k8s、mesos; default: k8s"`
	ReservedCPUSetList string `json:"reserved_cpuset_list" value:"" usage:"cpuset number list to be reserved, e.g. 6,7,8"`
}

// NewOption create Option object
func NewOption() *Option {
	return &Option{}
}
