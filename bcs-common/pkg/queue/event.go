

package queue

import "github.com/Tencent/bk-bcs/bcs-common/pkg/watch"

// Event holding event info for data object
type Event struct {
	Type watch.EventType
	Data interface{}
}
