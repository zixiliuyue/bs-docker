

package nodeclient

// Client xxx
type Client interface {
	ListNodes() ([]*Node, error)
	Close()
}
