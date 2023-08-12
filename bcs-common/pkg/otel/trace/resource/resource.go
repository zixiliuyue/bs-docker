

// Package resource xxx
package resource

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// New returns a resource describing this application.
func New(serviceName string, attrs ...attribute.KeyValue) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName)),
	)
	if attrs != nil {
		for _, a := range attrs {
			r, _ = resource.Merge(r, resource.NewSchemaless(a))
		}
	}
	return r
}
