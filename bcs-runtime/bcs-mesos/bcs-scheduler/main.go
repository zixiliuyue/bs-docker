

/*
Package main is the entry of scheduler process.

Implements of init log, init flags and start scheduler manager.
*/
package main

import (
	"runtime"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-scheduler/src/manager"
	schedutil "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-scheduler/src/util"

	"golang.org/x/net/context"
)

// MesosSched mesos scheduler interface
type MesosSched struct {
	manager *manager.Manager
	config  schedutil.SchedConfig
	ctx     context.Context
}

// NewScheduler create scheduler instance
func NewScheduler(config schedutil.SchedConfig) (*MesosSched, error) {
	m, err := manager.New(config)
	if err != nil {
		return nil, err
	}

	sched := &MesosSched{
		config:  config,
		manager: m,
	}

	return sched, nil
}

// Start all mesos scheduler features
func (s *MesosSched) Start(ctx context.Context) error {

	if err := s.runManager(ctx); err != nil {
		return err
	}

	return nil
}

func (s *MesosSched) runManager(ctx context.Context) error {
	return s.manager.Start()
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	op := &schedutil.SchedulerOptions{}
	conf.Parse(op)

	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()

	if err := common.SavePid(op.ProcessConfig); err != nil {
		blog.Error("fail to save pid: err:%s", err.Error())
	}

	config := schedutil.NewSchedulerCfg()

	schedutil.SetSchedulerCfg(config, op)

	scheduler, _ := NewScheduler(*config)
	scheduler.Start(context.Background())
}
