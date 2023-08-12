

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-driver/app"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-driver/app/options"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	op := &options.MesosDriverOptionsOut{}
	conf.Parse(op)

	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()

	blog.Infof("driver options: %+v", op)

	if err := common.SavePid(op.ProcessConfig); err != nil {
		blog.Error("fail to save pid: err:%s", err.Error())
	}

	opIn := options.NewMesosDriverOption(op)

	if opIn.DriverConf.Cluster == "" {
		blog.Error("cluster unknown, mesos driver run fail")
		fmt.Fprintf(os.Stderr, "cluster unknown, mesos driver run fail\n")
		os.Exit(1)
	}

	if err := app.Run(opIn); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	blog.Info("mesos driver exit")
}
