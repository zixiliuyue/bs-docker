

package pkgs

import (
	"sync"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/cmd/config"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/remote/alert"
)

var (
	alertClientOnce sync.Once
	alertClient     alert.BcsAlarmInterface
)

// GetAlertClient for init alert system client
func GetAlertClient(options *config.AlertManagerOptions) alert.BcsAlarmInterface {
	alertClientOnce.Do(func() {
		alertClient = alert.NewAlertServer(&options.AlertServerOptions)
		if alertClient == nil {
			panic("init alertClient failed")
		}
	})

	return alertClient
}
