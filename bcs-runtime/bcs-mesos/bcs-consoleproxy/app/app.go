

// Package app xxx
package app

import (
	"os"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	comconf "github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/app/options"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/api"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/config"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/manager"
)

// ConsoleProxy is an console proxy struct
type ConsoleProxy struct {
	backend manager.Manager
	route   *api.Router
	conf    *config.ConsoleProxyConfig
}

// NewConsoleProxy create an ConsoleProxy object
func NewConsoleProxy(op *options.ConsoleProxyOption) *ConsoleProxy {
	setConfig(op)

	c := &ConsoleProxy{
		conf:    &op.Conf,
		backend: manager.NewManager(&op.Conf),
	}

	err := c.backend.Start()
	if err != nil {
		blog.Errorf("start manager error %s", err.Error())
		os.Exit(1)
	}

	// http server
	c.route = api.NewRouter(c.backend, c.conf)
	return c
}

// Run create a pid
func (c *ConsoleProxy) Run() {
	// pid
	if err := common.SavePid(comconf.ProcessConfig{}); err != nil {
		blog.Error("fail to save pid: err:%s", err.Error())
	}
}

func setConfig(op *options.ConsoleProxyOption) {
	op.Conf.Address = op.Address
	op.Conf.Port = int(op.Port)
	op.Conf.DockerEndpoint = op.DockerEndpoint
	op.Conf.Tty = op.Tty
	op.Conf.Privilege = op.Privilege
	op.Conf.Cmd = op.Cmd
	op.Conf.Ips = op.Ips
	op.Conf.IsAuth = op.IsAuth
	op.Conf.IsOneSeesion = op.IsOneSession

	// server cert directory
	if op.CertConfig.ServerCertFile != "" && op.CertConfig.CAFile != "" &&
		op.CertConfig.ServerKeyFile != "" {

		op.Conf.ServCert.CertFile = op.CertConfig.ServerCertFile
		op.Conf.ServCert.KeyFile = op.CertConfig.ServerKeyFile
		op.Conf.ServCert.CAFile = op.CertConfig.CAFile
		op.Conf.ServCert.IsSSL = true
	}
}
