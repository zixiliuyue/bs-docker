

package main

import (
	"fmt"
	"os"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bmsf-mesh/bmsf-mesos-adapter/app"
)

func main() {
	cfg := &app.Config{}
	conf.Parse(cfg)
	if err := cfg.Validate(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	blog.InitLogs(cfg.LogConfig)
	defer blog.CloseLogs()
	if err := app.Run(cfg); err != nil {
		fmt.Printf("bmsf-mesos-adaptor starting failed: %v", err)
		os.Exit(1)
	}
}
