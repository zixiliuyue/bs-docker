

package main

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/cmd"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/cmd/config"
)

func main() {
	// init alertmanager options config
	alertManagerOptions := config.NewAlertManagerOptions()
	conf.Parse(alertManagerOptions)

	// init log config
	blog.InitLogs(alertManagerOptions.LogConfig)
	defer blog.CloseLogs()

	// init alertmanager & run server/consumer
	alertManager := cmd.NewAlertManager(alertManagerOptions)
	if err := alertManager.Run(); err != nil {
		blog.Fatalf("alertmanager run failed: %v", err)
	}
}
