

package main

import (
	"os"
	"runtime"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cpuset-device/app"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cpuset-device/app/options"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	op := options.NewOption()
	conf.Parse(op)

	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()
	blog.Info("init logs success")

	err := app.Run(op)
	if err != nil {
		blog.Errorf(err.Error())
		os.Exit(1)
	}

	ch := make(chan bool)
	<-ch
}
