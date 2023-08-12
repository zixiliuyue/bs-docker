

// Package config xxx
package config

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/static"
)

// CertConfig is configuration of Cert
type CertConfig struct {
	CAFile     string
	CertFile   string
	KeyFile    string
	CertPasswd string
	IsSSL      bool
}

// CustomSchedulerConfig is a configuration of CustomScheduler
type CustomSchedulerConfig struct {
	Address         string
	Port            uint
	InsecureAddress string
	InsecurePort    uint
	Sock            string
	MetricPort      uint
	ZkHosts         string
	ServCert        *CertConfig
	ClientCert      *CertConfig

	Cluster                  string
	KubeConfig               string
	KubeMaster               string
	UpdatePeriod             uint
	CustomSchedulerType      string
	VerifyClientTLS          bool
	CniAnnotationKey         string
	CniAnnotationValue       string
	FixedIpAnnotationKey     string
	FixedIpAnnotationValue   string
	CloudNetserviceEndpoints []string
	CloudNetserviceCert      *CertConfig
}

// NewCustomSchedulerConfig create a config object
func NewCustomSchedulerConfig() *CustomSchedulerConfig {
	return &CustomSchedulerConfig{
		Address: "127.0.0.1",
		Port:    80,
		ServCert: &CertConfig{
			CertPasswd: static.ServerCertPwd,
			IsSSL:      false,
		},
		ClientCert: &CertConfig{
			CertPasswd: static.ClientCertPwd,
			IsSSL:      false,
		},
	}
}
