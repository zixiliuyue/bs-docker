

package websocketDialer

import (
	"net"
	"time"
)

// Dialer dialer for tunnel request
type Dialer func(network, address string) (net.Conn, error)

// HasSession to see if server has session for the client key
func (s *Server) HasSession(clientKey string) bool {
	_, err := s.sessions.getDialer(clientKey, 0)
	return err == nil
}

// Dial do dial
func (s *Server) Dial(clientKey string, deadline time.Duration, proto, address string) (net.Conn, error) {
	d, err := s.sessions.getDialer(clientKey, deadline)
	if err != nil {
		return nil, err
	}

	return d(proto, address)
}

// Dialer get dialer for client key
func (s *Server) Dialer(clientKey string, deadline time.Duration) Dialer {
	return func(proto, address string) (net.Conn, error) {
		return s.Dial(clientKey, deadline, proto, address)
	}
}
