

package kubedriver

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/websocketDialer"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/kubedriver/options"
)

const (
	// Module model name
	Module = "BCS-API-Tunnel-Module"
	// RegisterToken token header
	RegisterToken = "BCS-API-Tunnel-Token" // nolint
	// Params tunnel params
	Params = "BCS-API-Tunnel-Params"
	// Cluster cluster params
	Cluster = "BCS-API-Tunnel-ClusterId"
)

func buildWebsocketToApi(o *options.KubeDriverServerOptions) error {
	if o.RegisterUrl == "" {
		return errors.New("register url is empty")
	}
	bcsApiUrl, err := url.Parse(o.RegisterUrl)
	if err != nil {
		return err
	}

	if o.RegisterToken == "" {
		return errors.New("register token is empty")
	}
	if o.CustomClusterID == "" {
		return errors.New("custom clusterid is empty")
	}

	var serverAddress string
	if o.SecureServerConfigured() {
		serverAddress = fmt.Sprintf("https://%s:%d", o.BindAddress.String(), o.SecurePort)
	} else {
		serverAddress = fmt.Sprintf("http://%s:%d", o.BindAddress.String(), o.InsecurePort)
	}

	params := map[string]interface{}{
		"address": serverAddress,
	}
	bytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	headers := map[string][]string{
		Module:        {ModuleName},
		Cluster:       {o.CustomClusterID},
		RegisterToken: {o.RegisterToken},
		Params:        {base64.StdEncoding.EncodeToString(bytes)},
	}

	var tlsConfig *tls.Config
	if o.InsecureSkipVerify {
		tlsConfig = &tls.Config{InsecureSkipVerify: true} // nolint
	} else {
		// use bcs cacert
		pool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(o.ServerTLS.CAFile)
		if err != nil {
			return err
		}
		if ok := pool.AppendCertsFromPEM(ca); ok != true {
			return fmt.Errorf("append ca cert failed")
		}
		tlsConfig = &tls.Config{RootCAs: pool}
	}

	go func() {
		for {
			wsURL := fmt.Sprintf("wss://%s/bcsapi/v1/websocket/connect", bcsApiUrl.Host)
			blog.Infof("Connecting to %s with token %s", wsURL, o.RegisterToken)

			websocketDialer.ClientConnect(context.Background(), wsURL, headers, tlsConfig, nil, func(proto,
				address string) bool {
				switch proto {
				case "tcp":
					return true
				case "unix":
					return address == "/var/run/docker.sock"
				}
				return false
			})
			time.Sleep(5 * time.Second)
		}
	}()

	return nil
}
