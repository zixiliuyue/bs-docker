

package controller

// Controller xxx
type Controller interface {
	RunController(stopCh <-chan struct{}) error
	StopController()
}
