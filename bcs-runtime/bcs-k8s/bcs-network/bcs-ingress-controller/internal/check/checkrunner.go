

package check

import (
	"context"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// CheckRunner start check list
type CheckRunner struct {
	ctx       context.Context
	checkList []Checker
}

// NewCheckRunner return new check runner
func NewCheckRunner(ctx context.Context) *CheckRunner {
	return &CheckRunner{
		ctx:       ctx,
		checkList: make([]Checker, 0),
	}
}

// Register register checker
func (c *CheckRunner) Register(checker Checker) *CheckRunner {
	c.checkList = append(c.checkList, checker)
	return c
}

// Start 定时启动注册的所有checker
func (c *CheckRunner) Start() {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, item := range c.checkList {
					go item.Run()
				}
			case <-c.ctx.Done():
				blog.Infof("Stop run checker")
				return
			}
		}
	}()
}
