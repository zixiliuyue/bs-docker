

package main

import (
	"encoding/json"
	"runtime"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/app"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/app/options"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	op := options.NewHealthCheckOption()
	conf.Parse(op)
	op.LogConfig.ToStdErr = true

	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()
	blog.Info("init logs success")

	by, _ := json.Marshal(op)
	blog.Infof("options %s", string(by))

	blog.Info("init config success")

	app.Run()

	ch := make(chan bool)
	<-ch
}
