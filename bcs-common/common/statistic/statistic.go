

// Package statistic xxx
package statistic

import (
	"strings"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		stc = new(statistic)
		stc.pool = make(map[string][]error)
	})
}

var stc *statistic

// IncAccess xxx
func IncAccess() {
	stc.incAccess()
}

// GetTotalAccess xxx
func GetTotalAccess() int64 {
	return stc.getTotalAccess()
}

// Set xxx
func Set(id string, err ...error) {
	stc.set(id, err...)
}

// Reset xxx
func Reset(id string) {
	stc.reset(id)
}

// ResetAll xxx
func ResetAll() {
	stc.resetAll()
}

// Status xxx
func Status() (string, bool) {
	return stc.status()
}

type statistic struct {
	sync.RWMutex
	pool   map[string][]error
	access int64
}

func (s *statistic) incAccess() {
	s.Lock()
	defer s.Unlock()
	s.access = s.access + 1
}

func (s *statistic) getTotalAccess() int64 {
	s.RLock()
	defer s.RUnlock()
	return s.access
}

func (s *statistic) set(id string, err ...error) {
	s.Lock()
	defer s.Unlock()
	s.pool[id] = append(s.pool[id], err...)
}

func (s *statistic) reset(id string) {
	s.Lock()
	defer s.Unlock()
	delete(s.pool, id)
}

func (s *statistic) resetAll() {
	s.Lock()
	defer s.Unlock()
	s.pool = make(map[string][]error)
}

func (s *statistic) status() (string, bool) {
	s.RLock()
	defer s.RUnlock()
	msg := make([]string, 0)
	for _, e := range s.pool {
		for _, m := range e {
			msg = append(msg, m.Error())
		}
	}

	if len(msg) != 0 {
		return strings.Join(msg, ";"), true
	}

	return "", false
}
