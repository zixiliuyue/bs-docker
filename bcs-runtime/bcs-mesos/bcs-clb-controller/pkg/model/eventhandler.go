

package model

// EventHandler Handler interface for object changes
type EventHandler interface {
	OnAdd(obj interface{})
	OnUpdate(oldObj interface{}, newObj interface{})
	OnDelete(obj interface{})
}
