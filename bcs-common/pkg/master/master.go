

// Package master xxx
package master

import bcstypes "github.com/Tencent/bk-bcs/bcs-common/common/types"

// Master register server node in event storage, like
// zookeeper, etcd, check if local node is master
type Master interface {
	Init() error                                  // init stage, like create connection
	Finit()                                       // finit, release resource
	Register() error                              // registry information to storage
	Clean() error                                 // clean self node
	IsMaster() bool                               // check if self is master or not
	CheckSelfNode() (bool, error)                 // check self node exist, and data correct
	GetAllNodes() ([]*bcstypes.ServerInfo, error) // get all server nodes
	GetPath() string                              // get parent path
}

// Empty for test
type Empty struct{}

// Init init stage, like create connection
func (e *Empty) Init() error {
	return nil
}

// Finit init stage, like create connection
func (e *Empty) Finit() {
}

// Register registry information to storage
func (e *Empty) Register() error {
	return nil

}

// Clean clean self node
func (e *Empty) Clean() error {
	return nil
}

// IsMaster check if self is master or not
func (e *Empty) IsMaster() bool {
	return false
}

// CheckSelfNode check self node exist, and data correct
func (e *Empty) CheckSelfNode() (bool, error) {
	return false, nil
}

// GetAllNodes get all server nodes
func (e *Empty) GetAllNodes() ([]*bcstypes.ServerInfo, error) {
	return nil, nil
}

// GetPath setting self info, now is ip address & port
func (e *Empty) GetPath() string {
	return ""
}
