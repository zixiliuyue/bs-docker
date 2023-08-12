

package prometheus

import (
	"log"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/export/metric/aggregation"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

// MemoryOption xxx
type MemoryOption bool

// NewMeterProvider xxx
func NewMeterProvider(mo MemoryOption, opts ...controller.Option) (metric.MeterProvider, *prometheus.Exporter, error) {
	config := prometheus.Config{}
	c := controller.New(processor.NewFactory(
		selector.NewWithHistogramDistribution(
			histogram.WithExplicitBoundaries(config.DefaultHistogramBoundaries),
		),
		aggregation.CumulativeTemporalitySelector(),
		processor.WithMemory(bool(mo)),
	),
		opts...,
	)
	exporter, err := prometheus.New(config, c)
	if err != nil {
		log.Panicf("failed to initialize prometheus exporter %v", err)
	}
	global.SetMeterProvider(exporter.MeterProvider())
	return exporter.MeterProvider(), exporter, err
}
