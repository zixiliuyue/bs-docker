

package utils

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"go-micro.dev/v4/metadata"
	"go.opentelemetry.io/otel/trace"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/otel/trace/constants"
)

// ContextWithRequestID returns a copy of parent with requestID set as the current Span.
func ContextWithRequestID(parent context.Context, requestID string) context.Context {
	requestID = strings.Replace(requestID, "-", "", -1)
	tid, _ := trace.TraceIDFromHex(requestID)
	sid, _ := trace.SpanIDFromHex(requestID)
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    tid,
		SpanID:     sid,
		TraceFlags: trace.FlagsSampled,
		Remote:     true,
	})
	return trace.ContextWithSpanContext(parent, sc)
}

// GetOrCreateReqID 尝试读取 X-Request-Id，若不存在则随机生成
func GetOrCreateReqID(ctx context.Context) string {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return uuid.New().String()
	}
	if reqID, ok := md.Get(constants.RequestIDHeaderKey); ok {
		return reqID
	}
	return uuid.New().String()
}
