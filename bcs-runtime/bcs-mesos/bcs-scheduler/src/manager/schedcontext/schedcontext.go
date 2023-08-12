

// Package schedcontext xxx
package schedcontext

import (
	// "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-scheduler/src/manager/apiserver"
	"github.com/Tencent/bk-bcs/bcs-common/common/http/httpserver"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-scheduler/src/manager/remote/alertmanager"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-scheduler/src/manager/store"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-scheduler/src/util"
)

// SchedContext context for scheduler
type SchedContext struct {
	Store store.Store

	// change 0414
	// ApiServer *apiserver.ApiServer
	// for HTTPS
	ApiServer2 *httpserver.HttpServer

	// Config for scheduler config
	Config util.SchedConfig
	// AlertManager for alert interface
	AlertManager alertmanager.AlertManageInterface
}
