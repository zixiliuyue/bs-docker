

package main

import (
	"runtime"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/app"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-hpacontroller/app/options"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	op := options.NewHpaControllerOption()
	conf.Parse(op)

	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()
	blog.Info("init logs success")

	/*by, _ := json.Marshal(op)
	blog.Infof("options %s", string(by))*/

	blog.Info("init config success")

	app.Run(op)
	ch := make(chan bool)
	<-ch
}
