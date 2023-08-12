

package procdaemon

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/types"
)

// ProcDaemon xxx
type ProcDaemon interface {
	// CreateProcess xxx
	// create process object
	CreateProcess(*types.ProcessInfo) error

	// InspectProcessStatus xxx
	// inspect process status
	InspectProcessStatus(procId string) (*types.ProcessStatusInfo, error)

	// StopProcess xxx
	// stop process
	StopProcess(procId string, timeout int) error

	// DeleteProcess xxx
	// Delete process
	DeleteProcess(procId string) error

	// set process envs
	// types.BcsKV: key = env.key, value = env.value
	// SetProcessEnvs([]types.BcsKV)error

	// ReloadProcess xxx
	// reload process, exec reloadCmd
	ReloadProcess(procId string) error

	// RestartProcess xxx
	// restart process, exec restartCmd
	RestartProcess(procId string) error
}
