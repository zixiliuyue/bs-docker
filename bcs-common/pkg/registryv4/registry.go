

package registryv4

import (
	"crypto/tls"
	"time"

	//microRegistry "github.com/micro/go-micro/v2/registry"
	"go-micro.dev/v4/registry"
)

//Options registry options
type Options struct {
	//UUID for registry
	id string
	//Registry address, formation like ip:port
	RegistryAddr []string
	//register name, like $module.bkbcs.tencent.com
	Name string
	//bkbcs version information
	Version string
	//Meta info for go-micro registry
	Meta map[string]string
	//Address information for other module discovery
	// format likes ip:port
	RegAddr  string
	Config   *tls.Config
	TTL      time.Duration
	Interval time.Duration
	//EventHandler & modules that registry watchs
	EvtHandler EventHandler
	Modules    []string
}

//EventHandler handler for module update notification
type EventHandler func(name string)

// Registry interface for go-micro etcd discovery
type Registry interface {
	//Register service information to registry
	// register do not block, if module want to
	// clean registe information, call Deregister
	Register() error
	//Deregister clean service information from registry
	// it means that registry is ready to exit
	Deregister() error
	//Get get specified service by name in local cache
	//Get(name string) (*microRegistry.Service, error)
	Get(name string) (*registry.Service, error)
}
