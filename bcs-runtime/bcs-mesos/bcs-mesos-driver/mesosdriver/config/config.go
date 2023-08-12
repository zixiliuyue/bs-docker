

// Package config xxx
package config

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/registry"
)

// CertConfig is configuration of Cert
type CertConfig struct {
	CAFile     string
	CertFile   string
	KeyFile    string
	CertPasswd string
	IsSSL      bool
}

// MesosDriverConfig is a configuration of mesos driver
type MesosDriverConfig struct {
	Address      string
	Port         uint
	ExternalIp   string
	ExternalPort uint

	MetricPort uint

	RegDiscvSvr   string
	SchedDiscvSvr string
	Cluster       string

	ServCert   *CertConfig
	ClientCert *CertConfig

	AdmissionWebhook bool
	// KubeConfig kubeconfig for CustomResource
	KubeConfig string
	// MesosWebconsoleProxyPort
	MesosWebconsoleProxyPort uint

	// websocket register
	RegisterWithWebsocket bool
	RegisterToken         string
	RegisterURL           string
	InsecureSkipVerify    bool

	Etcd *registry.CMDOptions
}
