

package proxier

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"
	"strings"
	"time"

	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/storages/sqlstore"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/tunnel"
	"k8s.io/apimachinery/pkg/util/proxy"
	"k8s.io/client-go/transport"
)

// WsTunnel xxx
type WsTunnel struct {
	httpTransport *http.Transport
	serverAddress string
	userToken     string
	caCertData    string
}

// lookupWsHandler will lookup websocket dialer in cache
func (f *ReverseProxyDispatcher) lookupWsHandler(clusterId string, req *http.Request) (*proxy.UpgradeAwareHandler, bool,
	error) {
	credentials := sqlstore.GetWsCredentials(clusterId)
	if credentials == nil {
		return nil, false, nil
	}

	serverAddress := credentials.ServerAddress
	if !strings.HasSuffix(serverAddress, "/") {
		serverAddress = serverAddress + "/"
	}
	u, err := url.Parse(serverAddress)
	if err != nil {
		return nil, false, err
	}

	transport := f.getTransport(clusterId, credentials)
	if transport == nil {
		return nil, false, nil
	}

	responder := &responder{}
	proxyHandler := proxy.NewUpgradeAwareHandler(u, transport, true, false, responder)
	proxyHandler.UseRequestLocation = true

	return proxyHandler, true, nil
}

func (f *ReverseProxyDispatcher) wsTunnelChanged(clusterId string, credentials *m.WsClusterCredentials) bool {
	wsTunnel := f.wsTunnelStore[clusterId]
	return wsTunnel.serverAddress != credentials.ServerAddress || wsTunnel.caCertData != credentials.CaCertData ||
		wsTunnel.userToken != credentials.UserToken
}

// getTransport generate transport with dialer from tunnel
func (f *ReverseProxyDispatcher) getTransport(clusterId string, credentials *m.WsClusterCredentials) http.RoundTripper {
	tunnelServer := tunnel.DefaultTunnelServer
	if tunnelServer.HasSession(clusterId) {
		f.wsTunnelMutateLock.Lock()
		defer f.wsTunnelMutateLock.Unlock()

		if f.wsTunnelStore[clusterId] != nil && !f.wsTunnelChanged(clusterId, credentials) {
			return f.wsTunnelStore[clusterId].httpTransport
		}

		tp := &http.Transport{
			MaxIdleConnsPerHost: 10,
		}
		if credentials.CaCertData != "" {
			certs := x509.NewCertPool()
			caCrt := []byte(credentials.CaCertData)
			certs.AppendCertsFromPEM(caCrt)
			tp.TLSClientConfig = &tls.Config{
				RootCAs: certs,
			}
		}
		cd := tunnelServer.Dialer(clusterId, 15*time.Second)
		tp.Dial = cd

		if f.wsTunnelStore[clusterId] != nil {
			f.wsTunnelStore[clusterId].httpTransport.CloseIdleConnections()
		}
		f.wsTunnelStore[clusterId] = &WsTunnel{
			httpTransport: tp,
			serverAddress: credentials.ServerAddress,
			userToken:     credentials.UserToken,
			caCertData:    credentials.CaCertData,
		}

		bearerToken := credentials.UserToken
		bearerAuthRoundTripper := transport.NewBearerAuthRoundTripper(bearerToken, tp)

		return bearerAuthRoundTripper
	}

	return nil
}
