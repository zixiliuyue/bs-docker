

// Package config xxx
package config

// Config xxx
type Config struct {
	DockerSock   string
	LogbeatDir   string
	TemplateFile string
	PrefixFile   string
	// kube-apiserver config file path
	Kubeconfig string
	// whether to enable remove symbol link in the log path
	// this should be false if deployed as in-cluster mode
	EvalSymlink bool
	// logbeat pid file path
	LogbeatPIDFilePath  string
	NeedReload          bool
	FileExtension       string
	LogbeatOutputFormat string
	NodeName            string
}

// NewConfig create a config object
func NewConfig() *Config {
	return &Config{}
}
