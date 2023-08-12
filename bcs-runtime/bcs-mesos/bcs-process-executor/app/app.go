

// Package app xxx
package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/driver"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/executor"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/types"
)

// Run xxx
func Run() {
	cxt := context.Background()
	bcsExecutor := executor.NewExecutor(cxt)
	driver, err := driver.NewExecutorDriver(cxt, bcsExecutor)
	if err != nil {
		blog.Errorf("new xecutorDriver error %s, and exit", err.Error())
		os.Exit(1)
	}

	driver.Start()
	interrupt := make(chan os.Signal, 10)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	go handleSysSignal(interrupt, bcsExecutor)
	go handleExecutorFinish(bcsExecutor)
}

func handleSysSignal(signalChan <-chan os.Signal, executor executor.Executor) {
	select {
	case <-signalChan:
		blog.Infof("Get signal from system. Executor was killed, ready to Stop")
		executor.Shutdown()
		return
	}
}

func handleExecutorFinish(bcsExecutor executor.Executor) {
	for {
		time.Sleep(time.Second)

		if bcsExecutor.GetExecutorStatus() == types.ExecutorStatusFinish {
			blog.Infof("executor status %s, and exit", bcsExecutor.GetExecutorStatus())
			os.Exit(1)
		}
	}
}
