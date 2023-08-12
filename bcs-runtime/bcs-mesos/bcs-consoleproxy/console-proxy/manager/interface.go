

package manager

import (
	"net/http"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/types"
)

// Manager is an interface
type Manager interface {
	Start() error

	StartExec(http.ResponseWriter, *http.Request, *types.WebSocketConfig)
	CreateExec(http.ResponseWriter, *http.Request, *types.WebSocketConfig)
	ResizeExec(http.ResponseWriter, *http.Request, *types.WebSocketConfig)
}
