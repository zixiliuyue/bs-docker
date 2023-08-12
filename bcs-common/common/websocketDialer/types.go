

package websocketDialer

import "time"

const (
	// PingWaitDuration duration for wait ping
	PingWaitDuration = 60 * time.Second
	// PingWriteInterval ping interval
	PingWriteInterval = 5 * time.Second
	// HandshakeTimeOut timeout for client handshake
	HandshakeTimeOut = 10 * time.Second
)
