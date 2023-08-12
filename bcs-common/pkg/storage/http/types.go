

package http

import (
	"encoding/json"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/watch"
)

// Response basic response from http api
type Response struct {
	Code    int             `json:"code"`           // operation code
	Message string          `json:"message"`        // response message
	Data    json.RawMessage `json:"data,omitempty"` // response data
}

// WatchResponse basic response from http api
type WatchResponse struct {
	Code    int    `json:"code"`           // operation code
	Message string `json:"message"`        // response message
	Data    *Event `json:"data,omitempty"` // response data
}

// Event for http watch event
type Event struct {
	Type watch.EventType `json:"type"`
	Data json.RawMessage `json:"data"`
}
