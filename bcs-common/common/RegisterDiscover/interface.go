

package RegisterDiscover

// RegDiscvServer define the register and discover function interface
type RegDiscvServer interface {
	// Start the register and discover service
	Start() error
	// Stop the register and discover service
	Stop() error
	// Register server info into registe-discover service platform
	Register(key string, data []byte) error
	// RegisterAndWatch register server info into registe-discover service platform, and watch the info, if not exist, then register again
	RegisterAndWatch(key string, data []byte) error
	// Discover server from the registe-discover service platform
	Discover(key string) (<-chan *DiscoverEvent, error)
	// DiscoverNodes xx
	// discover nodes from path
	DiscoverNodes(path string) (*DiscoverEvent, error)
	// DiscoverNodesV2 xx
	// discover nodes from path v2
	DiscoverNodesV2(path string) (*DiscoverEvent, error)
}
