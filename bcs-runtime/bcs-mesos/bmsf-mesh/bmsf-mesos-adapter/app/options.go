

package app

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// Config detail configuration item
type Config struct {
	conf.FileConfig
	conf.ServiceConfig
	conf.MetricConfig
	conf.CertConfig
	conf.ZkConfig
	conf.LogConfig
	conf.ProcessConfig
	Scheme     string `json:"metric_scheme" value:"http" usage:"scheme for metric api"`
	Zookeeper  string `json:"zookeeper" value:"127.0.0.1:3181" usage:"data source for taskgroups and services"`
	Cluster    string `json:"cluster" value:"" usage:"cluster id or name"`
	KubeConfig string `json:"kubeconfig" value:"kubeconfig" usage:"configuration file for kube-apiserver"`
}

// Validate validate command line parameter
func (c *Config) Validate() error {
	if len(c.Cluster) == 0 {
		return fmt.Errorf("cluster cannot be empty")
	}
	return nil
}
