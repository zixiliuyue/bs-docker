

package manager

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-daemon/process-daemon/config"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/types"
)

// Manager xxx
type Manager interface {
	// Init func
	Init() error

	// Start func
	Start()

	// GetConfig xxx
	// get config
	GetConfig() *config.Config

	// HeartBeat xxx
	// heartbeat
	HeartBeat(heartbeat *types.HeartBeat)

	// CreateProcess xxx
	// Create process
	CreateProcess(processInfo *types.ProcessInfo) error

	// InspectProcessStatus xxx
	// inspect process status info
	// processId = types.ProcessInfo.Id
	InspectProcessStatus(processId string) (*types.ProcessStatusInfo, error)

	// StopProcess xxx
	// Stop process
	// processId = types.ProcessInfo.Id
	// process will be killed when timeout seconds
	StopProcess(processId string, timeout int) error

	// DeleteProcess xxx
	// delete process
	// processId = types.ProcessInfo.Id
	DeleteProcess(processId string) error
}
