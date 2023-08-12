

package trace

import (
	"go.opentelemetry.io/otel/attribute"
)

// TraceType tracing type
type TraceType string

const (
	// Jaeger show jaeger system
	Jaeger TraceType = "jaeger"
	// Zipkin show zipkin system
	Zipkin TraceType = "zipkin"
	// NullTrace show opentracing NoopTracer
	NullTrace TraceType = "null"
)

// TracerSwitch sets a factory tracing switch: on or off
func TracerSwitch(s string) Option {
	return func(o *Options) {
		o.TracingSwitch = s
	}
}

// TracerType sets a factory tracing type
func TracerType(t string) Option {
	return func(o *Options) {
		o.TracingType = string(t)
	}
}

// ServiceName sets a service name for a tracing system
func ServiceName(sn string) Option {
	return func(o *Options) {
		o.ServiceName = sn
	}
}

// ExporterURL sets a exporter url for tracing system
func ExporterURL(eu string) Option {
	return func(o *Options) {
		o.ExporterURL = eu
	}
}

// ResourceAttrs sets resource attributes
func ResourceAttrs(ra []attribute.KeyValue) Option {
	return func(o *Options) {
		o.ResourceAttrs = append(o.ResourceAttrs, ra...)
	}
}

// OTLPEndpoint sets OpenTelemetry Collector service endpoint
func OTLPEndpoint(en string) Option {
	return func(o *Options) {
		o.OTLPEndpoint = en
	}
}
