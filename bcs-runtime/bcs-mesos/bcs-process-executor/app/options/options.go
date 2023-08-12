

// Package options xxx
package options

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// HealthCheckOption is option in flags
type HealthCheckOption struct {
	conf.FileConfig
	conf.ServiceConfig
	conf.MetricConfig
	conf.ZkConfig
	conf.CertConfig
	conf.LicenseServerConfig
	conf.LogConfig
	conf.ProcessConfig
}

// NewHealthCheckOption create HealthCheckOption object
func NewHealthCheckOption() *HealthCheckOption {
	return &HealthCheckOption{}
}
