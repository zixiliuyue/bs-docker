

package v1

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/bcsapi"
	types "github.com/Tencent/bk-bcs/bcs-common/pkg/bcsapi/netservice"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/config"

	"fmt"
	"strings"
)

// BcsConfig config item for ipscheduler
type BcsConfig struct {
	ZkHost   string         `json:"zkHost"`
	TLS      *types.SSLInfo `json:"tls,omitempty"`
	Interval int            `json:"interval,omitempty"`
}

func createNetSvcClient(conf *config.CustomSchedulerConfig) (bcsapi.Netservice, error) {
	bcsConf := newBcsConf(conf)

	client := bcsapi.NewNetserviceCli()
	if bcsConf.TLS != nil {
		if err := client.SetCerts(
			bcsConf.TLS.CACert, bcsConf.TLS.Key, bcsConf.TLS.PubKey, bcsConf.TLS.Passwd); err != nil {
			return nil, err
		}
	}

	// client get bcs-netservice info
	hosts := strings.Split(bcsConf.ZkHost, ";")
	if err := client.GetNetService(hosts); err != nil {
		return nil, fmt.Errorf("get netservice failed, %s", err.Error())
	}
	return client, nil
}

func newBcsConf(conf *config.CustomSchedulerConfig) BcsConfig {
	bcsConf := BcsConfig{
		ZkHost: conf.ZkHosts,
		TLS: &types.SSLInfo{
			CACert: conf.ClientCert.CAFile,
			Key:    conf.ClientCert.KeyFile,
			PubKey: conf.ClientCert.CertFile,
			Passwd: conf.ClientCert.CertPasswd,
		},
	}

	return bcsConf
}
