

// Package options xxx
package options

import (
	"fmt"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// ServerOption is option in flags
type ServerOption struct {
	conf.FileConfig
	conf.MetricConfig
	conf.LogConfig
	conf.ProcessConfig

	Address        string `json:"address" short:"a" value:"0.0.0.0" usage:"IP address to listen on for this service"`
	Port           uint   `json:"port" short:"p" value:"443" usage:"Port to listen on for this service"`
	ServerCertFile string `json:"server_cert_file" value:"" usage:"Server public key file(*.crt). If both server_cert_file and server_key_file are set, it will set up an HTTPS server"`
	ServerKeyFile  string `json:"server_key_file" value:"" usage:"Server private key file(*.key). If both server_cert_file and server_key_file are set, it will set up an HTTPS server"`
	EngineType     string `json:"engine_type" value:"kubernetes" usage:"the platform that bcs-webhook-server runs in, kubernetes or mesos"`
	PluginDir      string `json:"plugin_dir" value:"./plugins" usage:"directory for bcs webhook plugins"`
	Plugins        string `json:"plugins" value:"" usage:"plugin names, call plugin Handle in this order"`
}

const (
	// EngineTypeKubernetes kubernetes engine type
	EngineTypeKubernetes = "kubernetes"
	// EngineTypeMesos mesos engine type
	EngineTypeMesos = "mesos"
)

// NewServerOption create a ServerOption object
func NewServerOption() *ServerOption {
	s := ServerOption{}
	return &s
}

// Parse parse server options
func Parse(ops *ServerOption) error {
	conf.Parse(ops)
	if ops.EngineType != EngineTypeKubernetes && ops.EngineType != EngineTypeMesos {
		return fmt.Errorf("unsupported engine type %s", ops.EngineType)
	}
	strings.Replace(ops.Plugins, ";", ",", -1)
	return nil
}
