

package app

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-driver/app/options"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-driver/mesosdriver"
)

// Run the mesos driver
func Run(op *options.MesosDriverOption) error {

	blog.Info("config: %+v", op)

	setConfig(op)

	driver, err := mesosdriver.NewMesosDriverServer(op.DriverConf)
	if err != nil {
		blog.Error("fail to create mesos driver. err:%s", err.Error())
		return err
	}

	blog.Info("app begin to start mesos driver ... ")
	driver.Start()
	blog.Info("mesos driver finish ")
	return nil
}

func setConfig(op *options.MesosDriverOption) {

	if op.DriverConf.ServCert.CertFile != "" && op.DriverConf.ServCert.KeyFile != "" {
		op.DriverConf.ServCert.IsSSL = true
	}

	if op.DriverConf.ClientCert.CertFile != "" && op.DriverConf.ClientCert.KeyFile != "" {
		op.DriverConf.ClientCert.IsSSL = true
	}
}
