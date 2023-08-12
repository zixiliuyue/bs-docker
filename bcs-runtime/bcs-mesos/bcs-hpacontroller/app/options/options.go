

// Package options xxx
package options

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/hpacontroller/config"
)

// HpaControllerOption is option in flags
type HpaControllerOption struct {
	conf.FileConfig
	conf.ServiceConfig
	conf.MetricConfig
	conf.ZkConfig
	conf.CertConfig
	conf.LicenseServerConfig
	conf.LogConfig
	conf.ProcessConfig

	ClusterZkAddr string `json:"cluster_zookeeper" value:"" usage:"bcs mesos cluster zk address"`
	KubeConfig    string `json:"kubeconfig" value:"" usage:"kubeconfig for kube-apiserver"`
	CadvisorPort  int    `json:"cadvisor_port" value:"" usage:"container cadvisor port"`
	ClusterID     string `json:"clusterid" value:"" usage:"bcs mesos cluster id"`

	Conf *config.Config
}

// NewHpaControllerOption create HpaControllerOption object
func NewHpaControllerOption() *HpaControllerOption {
	return &HpaControllerOption{
		Conf: config.NewConfig(),
	}
}
