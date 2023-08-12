

package zk

import (
	"encoding/json"
	"sync"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

var cmdLocks map[string]*sync.Mutex
var cmdRWlock sync.RWMutex

// InitCmdLockPool xxx
func (store *managerStore) InitCmdLockPool() {
	if cmdLocks == nil {
		blog.Info("init command lock pool")
		cmdLocks = make(map[string]*sync.Mutex)
	}
}

// LockCommand xxx
func (store *managerStore) LockCommand(cmdId string) {

	cmdRWlock.RLock()
	myLock, ok := cmdLocks[cmdId]
	cmdRWlock.RUnlock()
	if ok {
		myLock.Lock()
		return
	}

	cmdRWlock.Lock()
	myLock, ok = cmdLocks[cmdId]
	if !ok {
		blog.Info("create command lock(%s), current locknum(%d)", cmdId, len(cmdLocks))
		cmdLocks[cmdId] = new(sync.Mutex)
		myLock, _ = cmdLocks[cmdId]
	}
	cmdRWlock.Unlock()

	myLock.Lock()
	return
}

// UnLockCommand xxx
func (store *managerStore) UnLockCommand(cmdId string) {
	cmdRWlock.RLock()
	myLock, ok := cmdLocks[cmdId]
	cmdRWlock.RUnlock()

	if !ok {
		blog.Error("command lock(%s) not exist when do unlock", cmdId)
		return
	}
	myLock.Unlock()
}

func getCommandRootPath() string {
	return "/" + bcsRootNode + "/" + commandNode
}

// SaveCommand xxx
func (store *managerStore) SaveCommand(command *commtypes.BcsCommandInfo) error {

	data, err := json.Marshal(command)
	if err != nil {
		return err
	}

	path := getCommandRootPath() + "/" + command.Id

	return store.Db.Insert(path, string(data))
}

// FetchCommand xxx
func (store *managerStore) FetchCommand(ID string) (*commtypes.BcsCommandInfo, error) {

	path := getCommandRootPath() + "/" + ID

	data, err := store.Db.Fetch(path)
	if err != nil {
		blog.Info("get path(%s) err:%s", path, err.Error())
		return nil, err
	}

	command := &commtypes.BcsCommandInfo{}
	if err := json.Unmarshal(data, command); err != nil {
		blog.Error("fail to unmarshal command(%s). err:%s", string(data), err.Error())
		return nil, err
	}

	return command, nil
}

// DeleteCommand xxx
func (store *managerStore) DeleteCommand(ID string) error {

	path := getCommandRootPath() + "/" + ID
	if err := store.Db.Delete(path); err != nil {
		blog.Error("fail to delete command(%s) err:%s", path, err.Error())
		return err
	}

	return nil
}
