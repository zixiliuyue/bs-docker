

package proxier

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// UpstreamServer xxx
type UpstreamServer struct {
	clusterId string

	servers         []string
	availableServer string
	mu              sync.Mutex

	CheckPeriod     time.Duration
	TcpCheckTimeOut time.Duration

	// Callback functions
	// OnAvailableChanged will be called when upstreamServer detects a failure
	OnAvailabilityChanged func()
}

// NewUpstreamServer xxx
func NewUpstreamServer(clusterId string, serverAddresses []string, onAvailabilityChanged func()) *UpstreamServer {
	return &UpstreamServer{
		clusterId:       clusterId,
		servers:         serverAddresses,
		CheckPeriod:     time.Second * 5,
		TcpCheckTimeOut: time.Second * 2,

		OnAvailabilityChanged: onAvailabilityChanged,
	}
}

// Initialize xxx
func (s *UpstreamServer) Initialize() error {
	if len(s.servers) > 0 {
		s.availableServer = s.servers[0]
		s.startPeriodicChecker()
		return nil
	}
	return errors.New("no servers")
}

func (s *UpstreamServer) startPeriodicChecker() {
	go func() {
		stop := make(chan struct{})
		t := time.NewTicker(s.CheckPeriod)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				if len(s.servers) == 0 {
					stop <- struct{}{}
				} else {
					if CheckTcpConn(s.availableServer) != nil {
						s.reload()
					}
				}
			case <-stop:
				blog.Infof("stop PeriodicChecker of cluster %s", s.clusterId)
				return
			}
		}
	}()
}

// GetAvailableServer xxx
func (s *UpstreamServer) GetAvailableServer() string {
	return s.availableServer
}

// UpdateServerAddresses update the server addresses in upstreamServer object
func (s *UpstreamServer) UpdateServerAddresses(addresses []string) {
	s.servers = addresses
	s.reload()
}

func (s *UpstreamServer) reload() {
	oldServer := s.GetAvailableServer()
	s.updateAvailableSrv()

	if oldServer != s.GetAvailableServer() {
		// Call callback function
		blog.Infof("available server for cluster %s changed, current: %s", s.clusterId, s.availableServer)
		s.OnAvailabilityChanged()
	}
}

func (s *UpstreamServer) updateAvailableSrv() {
	defer s.mu.Unlock()
	s.mu.Lock()

	for _, serverAddr := range s.servers {
		if CheckTcpConn(serverAddr) == nil {
			s.availableServer = serverAddr
			return
		}
	}

	if len(s.servers) > 0 {
		blog.Debug(fmt.Sprintf("no available server for cluster %s, so only choose first one: %s", s.clusterId, s.servers[0]))
		s.availableServer = s.servers[0]
	} else {
		s.availableServer = ""
	}
}

// Stop stops the current UpstreamServer checker
func (s *UpstreamServer) Stop() {
	s.servers = []string{}
}
