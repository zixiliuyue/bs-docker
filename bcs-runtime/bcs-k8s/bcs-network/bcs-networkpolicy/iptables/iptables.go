

package iptables

import (
	original "github.com/coreos/go-iptables/iptables"
)

// Error wrapper for original iptables error
type Error struct {
	original.Error
}

// New create original iptables implementation
func New() (Interface, error) {
	return original.New()
}
