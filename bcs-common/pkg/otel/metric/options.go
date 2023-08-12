

package metric

import "github.com/Tencent/bk-bcs/bcs-common/pkg/otel/metric/prometheus"

// SwitchStatus sets a factory metric switch: on or off
func SwitchStatus(s string) Option {
	return func(o *Options) {
		o.MetricSwitchStatus = s
	}
}

// TypeMetric sets a factory metric type
func TypeMetric(mt string) Option {
	return func(o *Options) {
		o.MetricType = mt
	}
}

// ProcessorWithMemory xxx
func ProcessorWithMemory(p prometheus.MemoryOption) Option {
	return func(o *Options) {
		o.ProcessorWithMemory = p
	}
}
