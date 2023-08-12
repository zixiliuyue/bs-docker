

// Package app xxx
package app

import (
	"fmt"
	exec "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-container-executor/executor"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-container-executor/logs"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Run is entry point for container executor
func Run(cmd *CommandFlags) error {
	cmd.NetworkMode = strings.ToLower(cmd.NetworkMode)

	// create ExecutorDriver
	e := NewBcsExecutor(cmd)
	if e == nil {
		return fmt.Errorf("Create Executor failed")
	}
	driver := exec.NewExecutorDriver(e)
	if driver == nil {
		return fmt.Errorf("Create ExeuctorDriver failed")
	}
	e.SetDriver(driver)
	_, err := driver.Start()
	if err != nil {
		return nil
	}
	interupt := make(chan os.Signal, 10)
	signal.Notify(interupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	go handleSysSignal(interupt, e, driver)
	_, waitErr := driver.Join()
	return waitErr
}

func handleSysSignal(signalChan <-chan os.Signal, executor exec.Executor, driver exec.ExecutorDriver) {
	select {
	case <-signalChan:
		logs.Infoln("Get signal from system. Executor was killed, ready to Stop")
		executor.Shutdown(driver)
		driver.Stop()
		return
	}
}
