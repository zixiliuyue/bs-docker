

package network

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-container-executor/container"
)

// NetManager manager for NetworkPlugin
type NetManager interface {
	Init() error                                       // manager init
	Stop()                                             // manager stop if necessary
	GetPlugin(name string) NetworkPlugin               // get plugin by name
	AddPlugin(name string, plguin NetworkPlugin) error // Add plugin to manager dynamic if necessary
	SetUpPod(podInfo container.Pod) error              // for setting Pod network interface
	TearDownPod(podInfo container.Pod) error           // for release pod network resource
}
