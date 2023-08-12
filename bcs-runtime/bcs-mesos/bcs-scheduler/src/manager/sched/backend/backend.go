

// Package backend xxx
package backend

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-scheduler/src/manager/sched/scheduler"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-scheduler/src/manager/store"
)

type backend struct {
	sched *scheduler.Scheduler
	store store.Store
}

// NewBackend xxx
func NewBackend(sched *scheduler.Scheduler, zkStore store.Store) Backend {
	return &backend{
		sched: sched,
		store: zkStore,
	}
}

// ClusterId xxx
func (b *backend) ClusterId() string {
	return b.sched.ClusterId
}

// GetRole xxx
func (b *backend) GetRole() string {
	return b.sched.Role
}
