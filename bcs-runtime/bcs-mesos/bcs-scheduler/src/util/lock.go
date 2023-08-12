

package util

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// Lock xxx
var Lock *ObjectLock

func init() {
	Lock = &ObjectLock{
		locks: make(map[string]*sync.RWMutex),
	}
}

// ObjectLock xxx
type ObjectLock struct {
	rw    sync.RWMutex
	locks map[string]*sync.RWMutex
}

// Lock xxx
func (l *ObjectLock) Lock(obj interface{}, key string) {
	k := fmt.Sprintf("%s.%s", reflect.TypeOf(obj).Name(), key)
	l.rw.RLock()
	myLock, ok := l.locks[k]
	l.rw.RUnlock()
	if ok {
		myLock.Lock()
		return
	}

	l.rw.Lock()
	myLock, ok = l.locks[k]
	if !ok {
		blog.Info("create lock(%s), current locknum(%d)", k, len(l.locks))
		l.locks[k] = new(sync.RWMutex)
		myLock, _ = l.locks[k]
	}
	l.rw.Unlock()
	myLock.Lock()
	return
}

// UnLock xxx
func (l *ObjectLock) UnLock(obj interface{}, key string) {
	k := fmt.Sprintf("%s.%s", reflect.TypeOf(obj).Name(), key)
	l.rw.RLock()
	myLock, ok := l.locks[k]
	l.rw.RUnlock()

	if !ok {
		blog.Error("lock(%s) not exist when do unlock", k)
		return
	}
	myLock.Unlock()
}
