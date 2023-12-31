

// Package healthcheck xxx
package healthcheck

import "time"

// TimeMechanism item for checker
type TimeMechanism struct {
	// DelaySeconds        int // Amount of time to wait until starting the health checks.
	IntervalSeconds     int // Interval between health checks.
	TimeoutSeconds      int // Amount of time to wait for the health check to complete.
	ConsecutiveFailures int // Number of consecutive failures until signaling kill task.
	GracePeriodSeconds  int // Amount of time to allow failed health checks since launch.
}

// FailureNotify func handler
type FailureNotify func(Checker)

// CheckerStat checker statistic
type CheckerStat struct {
	Started             bool      // status for started
	StartPoint          time.Time // starting time
	LastFailure         time.Time // time record
	LastCheck           time.Time // time record
	Failures            int       // amount of failure
	Ticks               int       // total checks
	ConsecutiveFailures int       // failure
	Healthy             bool      // healthy status
}

// TotalTicks total check ticks
func (check *CheckerStat) TotalTicks() int {
	return check.Ticks
}

// ConsecutiveFailure get consecutive failure times
func (check *CheckerStat) ConsecutiveFailure() int {
	return check.ConsecutiveFailures
}

// IsTicks xxx
func (check *CheckerStat) IsTicks() bool {
	if check.Ticks > 0 {
		return true
	}

	return false
}

// Failure total failures
func (check *CheckerStat) Failure() int {
	return check.Failures
}

// LastFailureTime last Failure
func (check *CheckerStat) LastFailureTime() time.Time {
	return check.LastFailure
}

// LastCheckTime last check tick
func (check *CheckerStat) LastCheckTime() time.Time {
	return check.LastCheck
}

// IsHealthy get if check is healthy
func (check *CheckerStat) IsHealthy() bool {
	return check.Healthy
}

const (
	// HttpHealthcheck xxx
	HttpHealthcheck = "http"
	// TcpHealthcheck xxx
	TcpHealthcheck = "tcp"
	// CommandHealthcheck xxx
	CommandHealthcheck = "command"
)

// Checker health check interface
type Checker interface {
	IsStarting() bool           // checker is running
	SetHost(host string)        // setting ip/host for checker
	IsHealthy() bool            // get if check is healthy
	TotalTicks() int            // total check ticks
	IsTicks() bool              // is health check
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
