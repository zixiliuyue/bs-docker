

package store

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/types"
)

// Store xxx
type Store interface {
	// StoreProcessInfo xxx
	// Store process info
	StoreProcessInfo(processInfo *types.ProcessInfo) error

	// GetAllProcessInfos xxx
	// Get all process infos
	GetAllProcessInfos() ([]*types.ProcessInfo, error)

	// DeleteProcessInfo xxx
	// delete process info
	DeleteProcessInfo(processInfo *types.ProcessInfo) error
}
