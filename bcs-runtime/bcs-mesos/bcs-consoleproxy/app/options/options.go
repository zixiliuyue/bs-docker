

// Package options xxx
package options

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/config"
)

// ConsoleProxyOption is option in flags
type ConsoleProxyOption struct {
	conf.FileConfig
	conf.ServiceConfig
	conf.CertConfig
	conf.LicenseServerConfig
	conf.LogConfig
	conf.ProcessConfig

	Privilege      bool     `json:"privilege" value:"" usage:"container exec privilege"`
	Cmd            []string `json:"cmd" value:"" usage:"cosntainer exec cmd"`
	Tty            bool     `json:"tty" value:"" usage:"tty"`
	DockerEndpoint string   `json:"docker-endpoint" value:"" usage:"docker endpoint"`
	Ips            []string `json:"ips" value:"" usage:"IP white list"`
	IsAuth         bool     `json:"is-auth" value:"" usage:"is auth"`
	IsOneSession   bool     `json:"is-one-session" value:"" usage:"support just one session for an container"`

	Conf config.ConsoleProxyConfig
}

// NewConsoleProxyOption create ConsoleProxyOption object
func NewConsoleProxyOption() *ConsoleProxyOption {
	return &ConsoleProxyOption{
		Conf: config.NewConsoleProxyConfig(),
	}
}
