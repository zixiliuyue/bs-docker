

package jaeger

// ServiceName set tracer serviceName
func ServiceName(name string) JaeOption {
	return func(opts *JaeOptions) {
		opts.ServiceName = name
	}
}

// RPCMetrics set tracer rpcMetrics
func RPCMetrics(rm bool) JaeOption {
	return func(opts *JaeOptions) {
		opts.RPCMetrics = rm
	}
}

// ReportMetrics set on prometheus report metrics
func ReportMetrics(rm bool) JaeOption {
	return func(opts *JaeOptions) {
		opts.ReportMetrics = rm
	}
}

// FromEnv set jaeger-agent hostPort from env by container deploy
func FromEnv(fe bool) JaeOption {
	return func(opts *JaeOptions) {
		opts.FromEnv = fe
	}
}

// AgentHostPort set jaeger-agent hostPort by idc deploy
func AgentHostPort(hp string) JaeOption {
	return func(opts *JaeOptions) {
		opts.AgentHostPort = hp
	}
}

// ReportLog set report tracer/span info
func ReportLog(rl bool) JaeOption {
	return func(opts *JaeOptions) {
		opts.ReportLog = rl
	}
}

// SamplerConfigInit set the jaeger sampler
func SamplerConfigInit(sc SamplerConfig) JaeOption {
	return func(opts *JaeOptions) {
		opts.Sampler = sc
	}
}
