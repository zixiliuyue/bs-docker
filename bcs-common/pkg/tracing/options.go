

package tracing

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

// TracerType set factory tracing type
func TracerType(t TraceType) Option {
	return func(o *Options) {
		o.TracingType = string(t)
	}
}

// TracerSwitch set factory tracing switch: on or off
func TracerSwitch(s string) Option {
	return func(o *Options) {
		o.TracingSwitch = s
	}
}

// ServiceName set service name for tracing system
func ServiceName(sn string) Option {
	return func(o *Options) {
		o.ServiceName = sn
	}
}

// RPCMetrics set tracer rpcMetrics
func RPCMetrics(rm bool) Option {
	return func(opts *Options) {
		opts.RPCMetrics = rm
	}
}

// ReportMetrics set on prometheus report metrics
func ReportMetrics(rm bool) Option {
	return func(opts *Options) {
		opts.ReportMetrics = rm
	}
}

// ReportLog set report tracer/span info
func ReportLog(rl bool) Option {
	return func(opts *Options) {
		opts.ReportLog = rl
	}
}

// AgentFromEnv set report agent from env
func AgentFromEnv(af bool) Option {
	return func(opts *Options) {
		opts.AgentFromEnv = af
	}
}

// AgentHostPort set reporter agent server
func AgentHostPort(ah string) Option {
	return func(opts *Options) {
		opts.AgentHostPort = ah
	}
}

// SampleType set the jaeger samplerType
func SampleType(st string) Option {
	return func(opts *Options) {
		opts.SampleType = st
	}
}

// SampleParameter for sampler parameter
func SampleParameter(sp float64) Option {
	return func(opts *Options) {
		opts.SampleParameter = sp
	}
}

// SampleFromEnv set the jaeger samplerType
func SampleFromEnv(sf bool) Option {
	return func(opts *Options) {
		opts.SampleFromEnv = sf
	}
}

// SamplingServerURL for sampler parameter
func SamplingServerURL(su string) Option {
	return func(opts *Options) {
		opts.SamplingServerURL = su
	}
}
