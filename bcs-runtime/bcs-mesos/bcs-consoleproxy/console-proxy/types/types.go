

// Package types xxx
package types

// WebSocketConfig is config
type WebSocketConfig struct {
	Height      int
	Width       int
	Privilege   bool
	Cmd         []string
	Tty         bool
	ContainerID string
	Token       string
	Origin      string
	User        string
	ExecID      string
}
