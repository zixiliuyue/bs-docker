

package check

import (
	"flag"
	"sync"
	"time"
)

type healthWatcher struct {
	stableTime time.Duration
	err        chan error
	lock       sync.Mutex
}

var watcher healthWatcher

func (hc *healthWatcher) change(err error) {
	hc.lock.Lock()
	defer hc.lock.Unlock()

	hc.err <- err
}

func (hc *healthWatcher) run() {
	go func() {
		select {
		case err := <-hc.err:
			if err == nil {
				Succeed()
			} else {
				Fail(err.Error())
			}
		case <-time.After(hc.stableTime):
			Succeed()
		}
	}()
}

var isStarted = false
var startLock sync.Mutex

// StartHealthCheck start to check module status.
// After stableTime, it will print success-message to log file
func StartHealthCheck() {
	startLock.Lock()
	defer startLock.Unlock()
	if isStarted {
		return
	}
	isStarted = true

	stableTime := *flag.Int("stableTime", 5,
		"if there are not problems occur in stableTime seconds, means start successful. default 5 seconds.")
	watcher = healthWatcher{err: make(chan error, 1), stableTime: time.Duration(stableTime) * time.Second}
	watcher.run()
}

// Occur can be called in stableTime and print failure-message to log file
// and it will end the check loop
func Occur(err error) {
	if !isStarted {
		return
	}
	if err != nil {
		watcher.change(err)
	}
}

// Complete can be called in stableTime to end the check loop with success-message
func Complete() {
	if !isStarted {
		return
	}
	watcher.change(nil)
}
