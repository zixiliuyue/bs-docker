

package pkgs

import (
	"sync"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/cmd/config"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/handler/eventhandler"
)

const (
	// EventHandleConcurrencyNum xxx
	// alert-system interface concurrency
	EventHandleConcurrencyNum = 100
	// EventHandleAlertEventNum xxx
	// alert-system handle batch EventNum
	EventHandleAlertEventNum = 100
	// EventHandleChanQueueLen xxx
	// alert-system chan QueueLen
	EventHandleChanQueueLen = 1024
	// EventHandleBatchAggregation xxx
	// alert-system batchAggregation switch
	EventHandleBatchAggregation = false
)

var (
	eventHandlerOnce sync.Once
	eventHandler     *eventhandler.SyncEventHandler
)

// GetEventSyncHandler get eventSyncHandler consumer
func GetEventSyncHandler(options *config.AlertManagerOptions) *eventhandler.SyncEventHandler {
	eventHandlerOnce.Do(func() {
		eventHandler = eventhandler.NewSyncEventHandler(eventhandler.Options{
			AlertEventBatchNum: func() int {
				if options.HandlerConfig.AlertEventNum <= 0 {
					return EventHandleAlertEventNum
				}
				return options.HandlerConfig.AlertEventNum
			}(),
			ConcurrencyNum: func() int {
				if options.HandlerConfig.ConcurrencyNum <= 0 {
					return EventHandleConcurrencyNum
				}
				return options.HandlerConfig.ConcurrencyNum
			}(),
			ChanQueueNum: func() int {
				if options.HandlerConfig.ChanQueueNum <= 0 {
					return EventHandleChanQueueLen
				}
				return options.HandlerConfig.ChanQueueNum
			}(),
			IsBatchAggregation: func() bool {
				if options.HandlerConfig.IsBatchAggregation {
					return options.HandlerConfig.IsBatchAggregation
				}
				return EventHandleBatchAggregation
			}(),
			Client: GetAlertClient(options),
		})
		if eventHandler == nil {
			panic("init NewSyncEventHandler failed")
		}
		blog.Infof("init EventSyncHandler successful")
	})

	return eventHandler
}
