

package options

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// GitopsWebhookOptions defines the option of gitops webhook
type GitopsWebhookOptions struct {
	conf.FileConfig
	conf.LogConfig
	conf.CertConfig

	Address     string `json:"address,omitempty"`
	IPv6Address string `json:"ipv6address,omitempty"`
	GRPCPort    uint   `json:"grpcPort,omitempty"`
	HTTPPort    uint   `json:"httpPort,omitempty"`
	MetricPort  uint   `json:"metricPort,omitempty"`

	Registry Registry `json:"registry,omitempty"`

	GitOpsWebhook string `json:"gitOpsWebhook"`
	GitOpsToken   string `json:"gitOpsToken"`
}

// Registry defines the registry of gitops webhook
type Registry struct {
	Endpoints string `json:"endpoints,omitempty"`
	CA        string `json:"ca,omitempty"`
	Key       string `json:"key,omitempty"`
	Cert      string `json:"cert,omitempty"`
}
