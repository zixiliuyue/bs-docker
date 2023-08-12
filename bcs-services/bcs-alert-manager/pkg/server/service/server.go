

package service

import (
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/remote/alert"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/server/actions"
)

// AlertManager server for alert manager
type AlertManager struct {
	consoleAction actions.Console
}

// NewAlertManager create alert handler for http server
func NewAlertManager(alertClient alert.BcsAlarmInterface) *AlertManager {
	return &AlertManager{
		consoleAction: actions.NewAlertAction(alertClient),
	}
}
