

package main

import (
	"runtime"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/app"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/app/options"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	op := options.NewConsoleProxyOption()
	conf.Parse(op)

	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()

	blog.Info("init logs success")
	blog.Info("init config success")

	app := app.NewConsoleProxy(op)
	app.Run()

	blog.Infof("console proxy is running")
	ch := make(chan bool)
	<-ch
}
