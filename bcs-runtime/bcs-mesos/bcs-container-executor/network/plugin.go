

package network

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-container-executor/container"
	"net"
)

// NetworkPlugin defination for all network
type NetworkPlugin interface {
	Name() string // Get plugin name
	// Init TODO
	// Type() string                            //Get plugin type, executable binary name
	Init(host string) error // init Plugin
	// SetUpPod TODO
	// Status() *NetStatus                      //get network status
	// Version() string                         //Plugin version
	SetUpPod(podInfo container.Pod) error    // Setup Network info for pod
	TearDownPod(podInfo container.Pod) error // Teardown pod network info
}

// NetStatus hold pod network info
type NetStatus struct {
	IfName  string    `json:"ifname"`  // device name
	IP      net.IP    `json:"ip"`      // ip address for pod
	Net     net.IPNet `json:"net"`     // net for ip address, including network mask
	Gateway net.IP    `json:"gateway"` // network gateway
}
