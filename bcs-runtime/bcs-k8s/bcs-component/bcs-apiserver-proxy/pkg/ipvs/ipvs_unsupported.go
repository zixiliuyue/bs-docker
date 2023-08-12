//go:build !linux
// +build !linux



package ipvs

import (
	"fmt"
)

// New returns a dummy Interface for unsupported platform.
func New() Interface {
	return &runner{}
}

type runner struct {
}

// Flush xxx
func (runner *runner) Flush() error {
	return fmt.Errorf("IPVS not supported for this platform")
}

// AddVirtualServer xxx
func (runner *runner) AddVirtualServer(*VirtualServer) error {
	return fmt.Errorf("IPVS not supported for this platform")
}

// UpdateVirtualServer xxx
func (runner *runner) UpdateVirtualServer(*VirtualServer) error {
	return fmt.Errorf("IPVS not supported for this platform")
}

// DeleteVirtualServer xxx
func (runner *runner) DeleteVirtualServer(*VirtualServer) error {
	return fmt.Errorf("IPVS not supported for this platform")
}

// GetVirtualServer xxx
func (runner *runner) GetVirtualServer(*VirtualServer) (*VirtualServer, error) {
	return nil, fmt.Errorf("IPVS not supported for this platform")
}

// GetVirtualServers xxx
func (runner *runner) GetVirtualServers() ([]*VirtualServer, error) {
	return nil, fmt.Errorf("IPVS not supported for this platform")
}

// AddRealServer xxx
func (runner *runner) AddRealServer(*VirtualServer, *RealServer) error {
	return fmt.Errorf("IPVS not supported for this platform")
}

// GetRealServers xxx
func (runner *runner) GetRealServers(*VirtualServer) ([]*RealServer, error) {
	return nil, fmt.Errorf("IPVS not supported for this platform")
}

// DeleteRealServer xxx
func (runner *runner) DeleteRealServer(*VirtualServer, *RealServer) error {
	return fmt.Errorf("IPVS not supported for this platform")
}

// UpdateRealServer xxx
func (runner *runner) UpdateRealServer(*VirtualServer, *RealServer) error {
	return fmt.Errorf("IPVS not supported for this platform")
}

var _ = Interface(&runner{})
