

package zipkin

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/otel/trace/resource"
)

// NewTracerProvider xxx
func NewTracerProvider(exporterURL, serviceName string) (*sdktrace.TracerProvider, error) {
	// Create Zipkin Exporter and install it as a global tracer.
	//
	// For demoing purposes, always sample. In a production application, you should
	// configure the sampler to a trace.ParentBased(trace.TraceIDRatioBased) set at the desired
	// ratio.
	exporter, err := zipkin.New(exporterURL)
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.New(serviceName)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}
