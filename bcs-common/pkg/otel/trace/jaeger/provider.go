

package jaeger

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/otel/trace/resource"
)

// NewTracerProvider xxx
func NewTracerProvider(exporterURL, serviceName string) (*sdktrace.TracerProvider, error) {
	// Create the Jaeger exporter
	traceExp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(exporterURL)))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(traceExp),
		// Record information about this application in an Resource.
		sdktrace.WithResource(resource.New(serviceName)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}
