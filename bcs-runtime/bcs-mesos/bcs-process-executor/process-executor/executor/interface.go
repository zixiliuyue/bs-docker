

package executor

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/mesosproto/mesos"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/types"
)

// Executor xxx
type Executor interface {
	// LaunchTaskgroup xxx
	// launch taskgroup
	LaunchTaskgroup(*mesos.TaskGroupInfo)

	// Shutdown xxx
	// shut down
	Shutdown()

	// RegisterCallbackFunc xxx
	// register callback function
	RegisterCallbackFunc(types.CallbackFuncType, interface{})

	// GetExecutorStatus xxx
	// Get Executor status
	GetExecutorStatus() types.ExecutorStatus

	// ReloadTasks xxx
	// reload tasks, exec reloadCmd
	ReloadTasks() error

	// RestartTasks xxx
	// restart tasks, exec restartCmd
	RestartTasks() error

	AckTaskStatusMessage(taskId string, uuid []byte)
}
