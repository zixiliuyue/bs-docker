

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/otel/trace"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/otel/trace/utils"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
)

func welcomePage(w http.ResponseWriter, r *http.Request) {
	_, span := utils.Tracer("demo-http-tracer").Start(r.Context(), "WelcomePage")
	log.Printf("traceID:%v, spanID:%v",
		span.SpanContext().TraceID().String(), span.SpanContext().SpanID().String())
	defer span.End()
	w.Write([]byte("Welcome to my website!\n"))
}

func main() {
	opts := trace.Options{
		TracingSwitch: "on",
		ExporterURL:   "http://localhost:14268/api/traces",
		ResourceAttrs: []attribute.KeyValue{
			attribute.String("endpoint", "http_server"),
		},
	}
	op := []trace.Option{}
	op = append(op, trace.TracerSwitch(opts.TracingSwitch))
	op = append(op, trace.ResourceAttrs(opts.ResourceAttrs))
	op = append(op, trace.ExporterURL(opts.ExporterURL))

	tp, err := trace.InitTracerProvider("demo-http-server", op...)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	wrappedHandler := otelhttp.NewHandler(http.HandlerFunc(welcomePage), "/")
	http.Handle("/", wrappedHandler)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}
