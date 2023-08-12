

package modulediscovery

// ModuleDiscovery bkbcs module discovery interface difinition
type ModuleDiscovery interface {
	// GetModuleServers xx
	// module: types.BCS_MODULE_SCHEDULER...
	// list all servers
	// if mesos-apiserver/k8s-apiserver module={module}/clusterid, for examples: mesosdriver/BCS-TESTBCSTEST01-10001
	GetModuleServers(module string) ([]interface{}, error)

	// GetRandModuleServer xx
	// get random one server
	// if mesos-apiserver/k8s-apiserver module={module}/clusterid, for examples: mesosdriver/BCS-TESTBCSTEST01-10001
	GetRandModuleServer(moduleName string) (interface{}, error)

	// RegisterEventFunc xx
	// register event handle function
	RegisterEventFunc(handleFunc EventHandleFunc)

	// Stop close discovery
	Stop()
}

// EventHandleFunc module: types.BCS_MODULE_SCHEDULER...
// if mesos-apiserver/k8s-apiserver module={module}/clusterid, for examples: mesosdriver/BCS-TESTBCSTEST01-10001
type EventHandleFunc func(module string)
