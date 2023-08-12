

// Package options xxx
package options

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// SidecarOption is option in flags
type SidecarOption struct {
	conf.FileConfig
	conf.ServiceConfig
	conf.MetricConfig
	conf.ZkConfig
	conf.CertConfig
	conf.LicenseServerConfig
	conf.LogConfig
	conf.ProcessConfig

	DockerSock          string `json:"docker_sock" value:"unix:///var/run/docker.sock" usage:"docker socket file"`
	LogbeatDir          string `json:"logbeat_dir" value:"" usage:"logbeat config directory"`
	LogbeatOutputFormat string `json:"output_format" value:"" usage:"logbeat output format, \"v1\" for unifytlogc adapted mode, empty string for logbeat mode"`
	NeedReload          bool   `json:"need_reload" value:"" usage:"whether need reload logbeat when updating log collection config"`
	TemplateFile        string `json:"template_file" value:"./unifytlogc-template.conf" usage:"logbeat template config file"`
	PrefixFile          string `json:"prefix_file" value:"" usage:"logbeat config file prefix name"`
	FileExtension       string `json:"file_extension" value:"" usage:"logbeat config file extension"`
	Kubeconfig          string `json:"kubeconfig" value:"" usage:"kubeconfig"`
	EvalSymlink         bool   `json:"eval_symlink" value:"false" usage:"whether to enable remove symbol link in the log path"`
	LogbeatPIDFilePath  string `json:"logbeat_pid_file_path" value:"" usage:"logbeat pid file path, which is used to reload logbeat"`
	NodeName            string `json:"node_name" value:"" usage:"logbeat list pods filtered by node name"`
}

// NewSidecarOption create SidecarOption object
func NewSidecarOption() *SidecarOption {
	return &SidecarOption{}
}
