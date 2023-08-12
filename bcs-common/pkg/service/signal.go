

package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// SetupSignalContext catch exit signal for process graceful exit
func SetupSignalContext() context.Context {
	closeHandler := make(chan os.Signal, 2)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(closeHandler, syscall.SIGINT, syscall.SIGTERM)
	// waiting for signal
	go func() {
		<-closeHandler
		blog.Infof("service catch exit signal, exit in 3 seconds...")
		fmt.Printf("service catch exit signal, exit in 3 seconds...\n")
		cancel()
		time.Sleep(time.Second * 3)
		<-closeHandler
		fmt.Printf("service catch second exit signal, exit immediately...\n")
		os.Exit(-2)
	}()
	return ctx
}
