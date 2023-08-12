

// Package webconsole xxx
package webconsole

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-driver/mesosdriver/config"

	"github.com/gorilla/websocket"
)

// WebconsoleProxy mesos web console proxy implementation for mesos-driver
type WebconsoleProxy struct {
	// Backend returns the backend URL which the proxy uses to reverse proxy
	Backend func(*http.Request) (*url.URL, error)
	// Certificatio configuration for backend console proxy
	CertConfig *config.CertConfig
}

// NewWebconsoleProxy create proxy instance for mesos-driver
func NewWebconsoleProxy(port uint, certConfig *config.CertConfig) *WebconsoleProxy {
	backend := func(req *http.Request) (*url.URL, error) {
		v := req.URL.Query()
		hostIp := v.Get("host_ip")
		if hostIp == "" {
			return nil, fmt.Errorf("param host_ip must not be empty")
		}
		host := net.JoinHostPort(hostIp, strconv.Itoa(int(port)))

		u := &url.URL{}
		u.Host = host
		u.Fragment = req.URL.Fragment
		// ! this is hard code in WebconsoleProxy from 1.17.x, it's not elegent.
		// todo(DeveloperJim): discussion about mechanism that we can avoid hard code
		u.Path = strings.Replace(req.URL.Path, "/mesosdriver/v4/webconsole", "/bcsapi/v1/consoleproxy", 1)
		u.RawQuery = req.URL.RawQuery
		blog.Infof(u.String())
		return u, nil
	}

	return &WebconsoleProxy{
		Backend:    backend,
		CertConfig: certConfig,
	}
}

// ServeHTTP original http interface implementation
func (w *WebconsoleProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	backendURL, err := w.Backend(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if websocket.IsWebSocketUpgrade(req) {
		websocketProxy := NewWebsocketProxy(w.CertConfig, backendURL)
		websocketProxy.ServeHTTP(rw, req)
		return
	}

	httpProxy, err := NewHttpReverseProxy(backendURL, w.CertConfig)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	httpProxy.ServeHTTP(rw, req)
	return
}
