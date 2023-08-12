

// Package manager xxx
package manager

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/config"
	"github.com/fsouza/go-dockerclient"
	"sync"
)

type manager struct {
	sync.RWMutex
	conf                *config.ConsoleProxyConfig
	dockerClient        *docker.Client
	connectedContainers map[string]bool
}

// NewManager create a Manager object
func NewManager(conf *config.ConsoleProxyConfig) Manager {
	return &manager{
		conf:                conf,
		connectedContainers: make(map[string]bool),
	}
}

// Start create docker client
func (m *manager) Start() error {
	var err error
	m.dockerClient, err = docker.NewClient(m.conf.DockerEndpoint)
	return err
}
