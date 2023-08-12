

// Package consumer xxx
package consumer

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/msgqueue"
)

// Consumer for subscribe handler interface
type Consumer interface {
	Consume(ctx context.Context, queue msgqueue.MessageQueue) error
	Stop() error
}

// Consumers manager all consumer for subscribe handler
type Consumers struct {
	ctx       context.Context
	cancel    context.CancelFunc
	que       msgqueue.MessageQueue
	consumers []Consumer
}

// NewConsumers init consumers
func NewConsumers(consumers []Consumer, queue msgqueue.MessageQueue) *Consumers {
	c := &Consumers{
		que:       queue,
		consumers: consumers,
	}

	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c
}

// Run run all consumer
func (c *Consumers) Run() {
	if c == nil {
		return
	}

	for idx := range c.consumers {
		go func(ctx context.Context, consumer Consumer) {
			defer func() {
				if r := recover(); r != nil {
					blog.Errorf("[monitor][panic] consumer panic: %v\n", r)
				}
			}()

			consumer.Consume(ctx, c.que)
		}(c.ctx, c.consumers[idx])
	}
}

// Stop stop subscribe & close queue
func (c *Consumers) Stop() {
	blog.Info("recv term signal")
	for idx := range c.consumers {
		c.consumers[idx].Stop()
	}
	c.cancel()
	c.que.Stop()
}
