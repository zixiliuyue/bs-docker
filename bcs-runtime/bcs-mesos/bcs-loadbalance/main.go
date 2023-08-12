

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-loadbalance/app"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-loadbalance/option"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

func main() {
	var err error
	config := option.NewConfig()
	if err = config.Parse(); err != nil {
		fmt.Printf("parse config failed, err %s", err.Error())
		blog.Errorf("parse config failed, err %s", err.Error())
		os.Exit(1)
	}
	app.InitLogger(config)
	defer app.CloseLogger()
	processor := app.NewEventProcessor(config)
	if err = processor.Start(); err != nil {
		processor.Stop()
		fmt.Printf("bcs-loadbalance starting error %s\n", err.Error())
		blog.Errorf("bcs-loadbalance starting error %s", err.Error())
		os.Exit(1)
	}
	interrupt := make(chan os.Signal, 10)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGUSR1, syscall.SIGUSR2)
	processor.HandleSignal(interrupt)
}
