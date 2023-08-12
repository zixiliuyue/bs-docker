

package cnm

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-container-executor/container"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-container-executor/network"
)

// NewNetManager interface for return DockerNetManager
func NewNetManager() network.NetManager {
	manager := &DockerNetManager{}
	return manager
}

// DockerNetManager docker network manager for using docker network
type DockerNetManager struct {
}

// Init loading all configuration in directory
func (manager *DockerNetManager) Init() error {
	return nil
}

// Stop manager stop if necessary
func (manager *DockerNetManager) Stop() {
	// empty
}

// GetPlugin get plugin by name
func (manager *DockerNetManager) GetPlugin(name string) network.NetworkPlugin {
	return nil
}

// AddPlugin Add plugin to manager dynamic if necessary
func (manager *DockerNetManager) AddPlugin(name string, plguin network.NetworkPlugin) error {
	return nil
}

// SetUpPod for setting Pod network interface
func (manager *DockerNetManager) SetUpPod(podInfo container.Pod) error {
	return nil
}

// TearDownPod for release pod network resource
func (manager *DockerNetManager) TearDownPod(podInfo container.Pod) error {
	return nil
}
