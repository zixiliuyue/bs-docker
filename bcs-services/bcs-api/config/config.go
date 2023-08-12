

// Package config xxx
package config

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/static"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/options"
)

// CertConfig is configuration of Cert
type CertConfig struct {
	CAFile     string
	CertFile   string
	KeyFile    string
	CertPasswd string
	IsSSL      bool
}

// ApiServConfig is a configuration of apiserver
type ApiServConfig struct {
	Address         string
	Port            uint
	InsecureAddress string
	InsecurePort    uint
	Sock            string
	CIHost          string
	CCHost          string
	CCPhpHost       string
	BcsRoute        string
	BcsDataHost     string
	HubHost         string
	TDHost          string
	AuthHost        string
	RegDiscvSrv     string
	JfrogAccount    map[string]string
	Filter          bool
	LocalIp         string
	MetricPort      uint
	VerifyClientTLS bool

	BKIamAuth options.AuthOption

	ServCert   *CertConfig
	ClientCert *CertConfig

	BKE                      options.BKEOptions
	TKE                      options.TKEOptions
	Edition                  string
	MesosWebconsoleProxyPort uint
	PeerToken                string
}

var (
	// Edition xxx
	Edition = ""
	// TurnOnRBAC xxx
	TurnOnRBAC = false
	// BKIamAuth xxx
	BKIamAuth options.AuthOption
	// ClusterCredentialsFixtures xxx
	ClusterCredentialsFixtures options.CredentialsFixturesOptions
	// MesosWebconsoleProxyPort xxx
	MesosWebconsoleProxyPort uint
	// TkeConf xxx
	TkeConf options.TKEOptions
)

// NewApiServConfig create a config object
func NewApiServConfig() *ApiServConfig {
	return &ApiServConfig{
		Address: "127.0.0.1",
		Port:    50001,
		ServCert: &CertConfig{
			CertPasswd: static.ServerCertPwd,
			IsSSL:      false,
		},
		ClientCert: &CertConfig{
			CertPasswd: static.ClientCertPwd,
			IsSSL:      false,
		},
		Filter: false,
	}
}
