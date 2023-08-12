

// Package options xxx
package options

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// ServerOption is option in flags
type ServerOption struct {
	conf.FileConfig
	conf.MetricConfig
	conf.LogConfig
	conf.ProcessConfig

	EngineType string `json:"engine_type" value:"kubernetes" usage:"the platform that bcs-webhook-server runs in, kubernetes or mesos"`
	KubeConfig string `json:"kubeconfig" value:"" usage:"kubeconfig for kube-apiserver, Only required if out-of-cluster."`
	KubeMaster string `json:"kube-master" value:"" usage:"The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster." mapstructure:"kube-master"`
	EsbUrl     string `json:"esb_url" value:"" usage:"esb api url to privilege"`
	AppCode    string `json:"app_code" value:"" usage:"app_code to call esb cc api"`
	AppSecret  string `json:"app_secret" value:"" usage:"app_secret to call esb cc api"`
	BkUsername string `json:"bk_username" value:"" usage:"bk username to call esb cc api"`
}

// NewServerOption create a ServerOption object
func NewServerOption() *ServerOption {
	s := ServerOption{}
	return &s
}

// Parse xxx
func Parse(ops *ServerOption) error {
	conf.Parse(ops)
	return nil
}
