

package cache

import (
	"fmt"
)

// Store is storage interface
type Store interface {
	Add(obj interface{}) error
	Update(obj interface{}) error
	Delete(obj interface{}) error
	List() []interface{}
	ListKeys() []string
	Get(obj interface{}) (item interface{}, exists bool, err error)
	GetByKey(key string) (item interface{}, exists bool, err error)
	Num() int
	Clear()
	Replace([]interface{}) error
}

// ObjectKeyFunc define make object to a uniq key
type ObjectKeyFunc func(obj interface{}) (string, error)

// KeyError wrapper error return from ObjectKeyFunc
type KeyError struct {
	Obj interface{}
	Err error
}

// Error return string info for KeyError
func (k KeyError) Error() string {
	return fmt.Sprintf("create key for object %+v failed: %v", k.Obj, k.Err)
}

// DataNoExist return when No data in Store
type DataNoExist struct {
	Obj interface{}
}

// Error return string info
func (k DataNoExist) Error() string {
	return fmt.Sprintf("no data object %+v in Store", k.Obj)
}
