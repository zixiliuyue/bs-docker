

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/otel/metric"

	"go.opentelemetry.io/otel/attribute"
	otelmetric "go.opentelemetry.io/otel/metric"
)

var (
	lemonsKey = attribute.Key("ex.com/lemons")
)

func main() {
	opts := &metric.Options{
		MetricSwitchStatus:  "on",
		MetricType:          "prometheus",
		ProcessorWithMemory: true,
	}

	op := []metric.Option{}
	if opts.MetricSwitchStatus != "" {
		op = append(op, metric.SwitchStatus(opts.MetricSwitchStatus))
	}
	if opts.MetricType != "" {
		op = append(op, metric.TypeMetric(opts.MetricType))
	}
	if !opts.ProcessorWithMemory {
		op = append(op, metric.ProcessorWithMemory(opts.ProcessorWithMemory))
	}

	mp, exp, err := metric.InitMeterProvider(op...)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/metrics", exp)
	go func() {
		log.Fatal(http.ListenAndServe(":30080", nil))
	}()

	meter := mp.Meter("ex.com/basic")
	observerLock := new(sync.RWMutex)
	observerValueToReport := new(float64)
	observerLabelsToReport := new([]attribute.KeyValue)
	cb := func(_ context.Context, result otelmetric.Float64ObserverResult) {
		(*observerLock).RLock()
		value := *observerValueToReport
		labels := *observerLabelsToReport
		(*observerLock).RUnlock()
		result.Observe(value, labels...)
	}
	_ = otelmetric.Must(meter).NewFloat64GaugeObserver("ex.com.one", cb,
		otelmetric.WithDescription("A GaugeObserver set to 1.0"),
	)

	histogram := otelmetric.Must(meter).NewFloat64Histogram("ex.com.two",
		otelmetric.WithDescription("A Histogram set to 1.0"))
	counter := otelmetric.Must(meter).NewFloat64Counter("ex.com.three",
		otelmetric.WithDescription("A Counter set to 1.0"))

	commonLabels := []attribute.KeyValue{
		lemonsKey.Int(10),
		attribute.String("A", "1"),
		attribute.String("B", "2"),
		attribute.String("C", "3")}
	notSoCommonLabels := []attribute.KeyValue{lemonsKey.Int(13)}

	ctx := context.Background()

	(*observerLock).Lock()
	*observerValueToReport = 1.0
	*observerLabelsToReport = commonLabels
	(*observerLock).Unlock()
	meter.RecordBatch(
		ctx,
		commonLabels,
		histogram.Measurement(2.0),
		counter.Measurement(12.0),
	)

	time.Sleep(2 * time.Second)

	(*observerLock).Lock()
	*observerValueToReport = 1.0
	*observerLabelsToReport = notSoCommonLabels
	(*observerLock).Unlock()
	meter.RecordBatch(
		ctx,
		notSoCommonLabels,
		histogram.Measurement(2.0),
		counter.Measurement(22.0),
	)

	time.Sleep(2 * time.Second)

	(*observerLock).Lock()
	*observerValueToReport = 13.0
	*observerLabelsToReport = commonLabels
	(*observerLock).Unlock()
	meter.RecordBatch(
		ctx,
		commonLabels,
		histogram.Measurement(12.0),
		counter.Measurement(13.0),
	)

	fmt.Println("Example finished updating, please visit :30080/metrics")

	select {}
}
