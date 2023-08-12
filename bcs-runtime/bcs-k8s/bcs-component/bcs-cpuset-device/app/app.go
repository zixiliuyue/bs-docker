

// Package app xxx
package app

import (
	"os"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/static"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cpuset-device/app/options"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cpuset-device/config"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cpuset-device/cpusetdevice"
)

// Run run the server
func Run(op *options.Option) error {
	conf := config.NewConfig()
	setConfig(conf, op)

	controller := cpusetdevice.NewCpusetDevicePlugin(conf)
	err := controller.Start()
	if err != nil {
		blog.Errorf("CpusetDevicePlugin Start failed: %s", err.Error())
		os.Exit(1)
	}

	blog.Info("CpusetDevicePlugin server ... ")
	return nil
}

func setConfig(conf *config.Config, op *options.Option) {
	conf.DockerSocket = op.DockerSock
	conf.PluginSocketDir = op.PluginSocketDir
	conf.CgroupCpusetRoot = op.CgroupCpusetRoot
	conf.BcsZk = op.BCSZk
	conf.Engine = op.Engine
	conf.ClusterID = op.ClusterID
	conf.NodeIP = op.Address

	// client cert directory
	if op.CertConfig.ClientCertFile != "" && op.CertConfig.CAFile != "" &&
		op.CertConfig.ClientKeyFile != "" {

		conf.ClientCert.CertFile = op.CertConfig.ClientCertFile
		conf.ClientCert.KeyFile = op.CertConfig.ClientKeyFile
		conf.ClientCert.CAFile = op.CertConfig.CAFile
		conf.ClientCert.IsSSL = true
		conf.ClientCert.CertPasswd = static.ClientCertPwd
	}

	conf.ReservedCPUSet = make(map[string]struct{})
	// parse reserved cpuset list
	if len(op.ReservedCPUSetList) != 0 {
		cpuSetStrList := strings.Split(op.ReservedCPUSetList, ",")
		for _, cpuSetStr := range cpuSetStrList {
			conf.ReservedCPUSet[strings.TrimSpace(cpuSetStr)] = struct{}{}
		}
	}
}
