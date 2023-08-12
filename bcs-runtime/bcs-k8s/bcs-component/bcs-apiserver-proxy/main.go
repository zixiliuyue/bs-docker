

package main

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-apiserver-proxy/cmd"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-apiserver-proxy/cmd/config"
	_ "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-apiserver-proxy/pkg/utils/metrics"
)

func main() {
	// init apiServer-proxy options config
	apiServerProxyOptions := config.NewProxyAPIServerOptions()
	conf.Parse(apiServerProxyOptions)

	// init log config
	blog.InitLogs(apiServerProxyOptions.LogConfig)
	defer blog.CloseLogs()

	// set proxyManager
	proxyManager, err := cmd.NewProxyManager(apiServerProxyOptions)
	if err != nil {
		blog.Fatalf("init NewProxyManager failed: %v", err)
		return
	}

	// init proxyManager
	err = proxyManager.Init(apiServerProxyOptions)
	if err != nil {
		blog.Fatalf("init proxyManager failed: %v", err)
		return
	}

	// run proxyManager
	if err := proxyManager.Run(); err != nil {
		blog.Infof("proxyManager run quit: %v", err)
	}

	return
}
