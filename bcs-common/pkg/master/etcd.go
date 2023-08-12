

package master

import (
	bcstypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// EtcdMaster xxx
// Empty for test
type EtcdMaster struct{}

// Init init stage, like create connection
func (e *EtcdMaster) Init() error {
	return nil
}

// Finit init stage, like create connection
func (e *EtcdMaster) Finit() {
}

// Register registry information to storage
func (e *EtcdMaster) Register() error {
	return nil

}

// Clean clean self node
func (e *EtcdMaster) Clean() error {
	return nil
}

// IsMaster check if self is master or not
func (e *EtcdMaster) IsMaster() bool {
	return false
}

// CheckSelfNode check self node exist, and data correct
func (e *EtcdMaster) CheckSelfNode() (bool, error) {
	return false, nil
}

// GetAllNodes get all server nodes
func (e *EtcdMaster) GetAllNodes() ([]*bcstypes.ServerInfo, error) {
	return nil, nil
}

// GetPath setting self info, now is ip address & port
func (e *EtcdMaster) GetPath() string {
	return ""
}
