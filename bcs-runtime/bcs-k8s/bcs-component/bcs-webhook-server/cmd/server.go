

package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-webhook-server/internal/server"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-webhook-server/options"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	// create option object
	op := options.NewServerOption()
	if err := options.Parse(op); err != nil {
		fmt.Printf("parse options failed, err %s\n", err.Error())
		os.Exit(1)
	}

	// init bcs log config
	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()

	// run bcs-webhook-server
	s, err := server.NewWebhookServer(op)
	if err != nil {
		blog.Fatalf("create webhook server failed, err %s", err.Error())
	}
	s.Run()
}
