

// Package config xxx
package config

// BcsCcAgentConfig is a configuration of  bcs-cc-agent
type BcsCcAgentConfig struct {
	MetricPort uint
	EngineType string
	Kubeconfig string
	KubeMaster string
	EsbUrl     string
	AppCode    string
	AppSecret  string
	BkUsername string
}

// NewBcsCcAgentConfig create a config object
func NewBcsCcAgentConfig() *BcsCcAgentConfig {
	return &BcsCcAgentConfig{}
}
