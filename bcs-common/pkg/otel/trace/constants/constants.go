

package constants

import "go.opentelemetry.io/otel/attribute"

const (
	// RequestIDKey xxx
	RequestIDKey       = "requestID"
	TracerKey          = "otel-go-contrib-tracer"
	TracerName         = "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	RequestIDHeaderKey = "X-Request-Id"
	GRPCStatusCodeKey  = attribute.Key("rpc.grpc.status_code")
)
