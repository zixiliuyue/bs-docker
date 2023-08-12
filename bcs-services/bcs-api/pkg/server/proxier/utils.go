

package proxier

import (
	"crypto/tls"
	"net/url"
	"strings"

	restclient "k8s.io/client-go/rest"

	"fmt"

	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
)

const (
	bcsK8sClusterDomain = "kubernetes"
)

// ExtractIpAddress xxx
func ExtractIpAddress(serverAddress string) (*url.URL, error) {
	if !strings.HasSuffix(serverAddress, "/") {
		serverAddress = serverAddress + "/"
	}
	ipAddress, err := url.Parse(serverAddress)
	if err != nil {
		return nil, err
	}
	return ipAddress, nil
}

// TurnCredentialsIntoConfig xxx
func TurnCredentialsIntoConfig(clusterCredentials *m.ClusterCredentials) (*restclient.Config, error) {

	tlsClientConfig := restclient.TLSClientConfig{
		ServerName: bcsK8sClusterDomain,
		CAData:     []byte(clusterCredentials.CaCertData),
	}
	return &restclient.Config{
		Host:            clusterCredentials.ServerAddresses,
		BearerToken:     clusterCredentials.UserToken,
		TLSClientConfig: tlsClientConfig,
	}, nil
}

// CheckTcpConn xxx
// check tcp connection to addr
func CheckTcpConn(addr string) error {
	checkUrl, err := url.Parse(addr)
	if err != nil {
		return err
	}
	err = dialTls(checkUrl.Host)
	if err != nil {
		return fmt.Errorf("connection to "+addr+" failed: %s", err.Error())
	}
	return nil
}

func dialTls(host string) error {
	conf := &tls.Config{
		InsecureSkipVerify: true, // nolint
	}
	conn, err := tls.Dial("tcp", host, conf)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
