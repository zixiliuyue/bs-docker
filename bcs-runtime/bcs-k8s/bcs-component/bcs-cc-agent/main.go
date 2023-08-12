

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cc-agent/app"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cc-agent/options"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// create option object
	op := options.NewServerOption()
	if err := options.Parse(op); err != nil {
		fmt.Printf("parse options failed: %v\n", err)
		os.Exit(1)
	}

	// init bcs log config
	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()

	// run bcs-cc-agent app
	app.Run(op)
}
