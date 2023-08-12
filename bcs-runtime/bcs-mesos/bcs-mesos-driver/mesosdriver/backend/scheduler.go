

package backend

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/http/httpclient"
	"github.com/Tencent/bk-bcs/bcs-common/common/http/httpserver"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-driver/mesosdriver/config"
)

// Scheduler xxx
type Scheduler interface {
	InitConfig(*config.MesosDriverConfig)
	Actions() []*httpserver.Action
	GetHost() string
	SetHost(hosts []string)
	GetHttpClient() *httpclient.HttpClient
}
