

// Package app xxx
package app

import (
	"os"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/app/options"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/hpacontroller/controller"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/hpacontroller/metrics/resources"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/hpacontroller/reflector"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/hpacontroller/scaler"
)

// Run the health check
func Run(op *options.HpaControllerOption) error {

	setConfig(op)
	// pid
	if err := common.SavePid(op.ProcessConfig); err != nil {
		blog.Error("fail to save pid: err:%s", err.Error())
	}

	// init store
	var store reflector.Reflector
	if op.Conf.ClusterZkAddr != "" {
		// init zk store
		store = reflector.NewZkReflector(op.Conf)
		blog.Infof("init cluster store zk %s success", op.Conf.ClusterZkAddr)
	} else if op.Conf.KubeConfig != "" {
		// init etcd store
		store = reflector.NewEtcdReflector(op.Conf)
		blog.Infof("init cluster store kubeconfig %s success", op.Conf.KubeConfig)
	} else {
		blog.Errorf("cluster zk addresses and kubeconfig not provided, exit")
		os.Exit(1)
	}
	// store := reflector.NewZkReflector(op.Conf)
	// blog.Infof("init cluster store zk %s success", op.Conf.ClusterZkAddr)

	// init bcs mesos driver
	scaleController := scaler.NewBcsMesosScalerController(op.Conf)
	blog.Infof("init cluster mesos driver controller success")

	// init cluster resource metrics
	resoucesCollector := resources.NewResourceMetrics(op.Conf, store)
	blog.Infof("init cluster resouces metrics collector success")

	hpaController := controller.NewAutoscaler(op.Conf, store, resoucesCollector, nil, scaleController)
	hpaController.Start()
	blog.Infof("hpa controller start work...")

	return nil
}

func setConfig(op *options.HpaControllerOption) {
	op.Conf.KubeConfig = op.KubeConfig
	op.Conf.ClusterZkAddr = op.ClusterZkAddr
	op.Conf.CadvisorPort = op.CadvisorPort
	op.Conf.BcsZkAddr = op.BCSZk
	op.Conf.ClusterID = op.ClusterID

	// client cert directoty
	if op.CertConfig.ClientCertFile != "" && op.CertConfig.CAFile != "" &&
		op.CertConfig.ClientKeyFile != "" {

		op.Conf.ClientCert.CertFile = op.CertConfig.ClientCertFile
		op.Conf.ClientCert.KeyFile = op.CertConfig.ClientKeyFile
		op.Conf.ClientCert.CAFile = op.CertConfig.CAFile
		op.Conf.ClientCert.IsSSL = true
	}
}
