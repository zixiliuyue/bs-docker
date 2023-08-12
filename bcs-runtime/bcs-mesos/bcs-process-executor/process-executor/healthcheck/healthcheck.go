

// Package healthcheck xxx
package healthcheck

import "time"

// TimeMechanism item for checker
type TimeMechanism struct {
	DelaySeconds        int // Amount of time to wait until starting the health checks.
	IntervalSeconds     int // Interval between health checks.
	TimeoutSeconds      int // Amount of time to wait for the health check to complete.
	ConsecutiveFailures int // Number of consecutive failures until signaling kill task.
	GracePeriodSeconds  int // Amount of time to allow failed health checks since launch.
}

// Checker health check interface
type Checker interface {
	IsStarting() bool           // checker is running
	SetHost(host string)        // setting ip/host for checker
	IsHealthy() bool            // get if check is healthy
	TotalTicks() int            // total check ticks
	ConsecutiveFailure() int    // get consecutive failure times
	Failure() int               // total failures
	LastFailureTime() time.Time // last Failure
	LastCheckTime() time.Time   // last check tick
	Start()                     // start checker
	Stop()                      // stop checker
	ReCheck() bool              // ask checker to check
	Pause() error               // pause check
	Resume() error              // arouse checker
	Name() string               // checker name
	Relation() string           // checker relative to container
}
