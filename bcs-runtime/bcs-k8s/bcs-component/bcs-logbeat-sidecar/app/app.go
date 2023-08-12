

// Package app xxx
package app

import (
	"os"
	"strconv"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-logbeat-sidecar/app/options"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-logbeat-sidecar/config"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-logbeat-sidecar/metric"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-logbeat-sidecar/sidecar"
)

// Run xxx
func Run(op *options.SidecarOption) error {

	conf := &config.Config{}
	setConfig(conf, op)
	// pid
	if err := common.SavePid(op.ProcessConfig); err != nil {
		blog.Error("fail to save pid: err:%s", err.Error())
	}

	metric.NewMetricClient(op.Address, strconv.Itoa(int(op.MetricPort))).Start()
	blog.Info("app start Metric client on %s:%s", op.Address, strconv.Itoa(int(op.MetricPort)))

	controller, err := sidecar.NewSidecarController(conf)
	if err != nil {
		blog.Errorf("NewSidecarController error %s", err.Error())
		os.Exit(1)
	}

	controller.Start()
	blog.Info("app start SidecarController server ... ")
	return nil
}

func setConfig(conf *config.Config, op *options.SidecarOption) {
	conf.DockerSock = op.DockerSock
	conf.PrefixFile = op.PrefixFile
	conf.TemplateFile = op.TemplateFile
	conf.LogbeatDir = op.LogbeatDir
	conf.Kubeconfig = op.Kubeconfig
	conf.EvalSymlink = op.EvalSymlink
	conf.LogbeatPIDFilePath = op.LogbeatPIDFilePath
	conf.NeedReload = op.NeedReload
	conf.LogbeatOutputFormat = op.LogbeatOutputFormat
	conf.NodeName = op.NodeName
	if op.FileExtension == "" {
		conf.FileExtension = "yaml"
	} else {
		conf.FileExtension = op.FileExtension
	}
}
