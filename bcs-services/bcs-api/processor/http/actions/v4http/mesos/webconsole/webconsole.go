

// Package webconsole xxx
package webconsole

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/config"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/metric"
	"github.com/gorilla/websocket"
)

// WebconsoleProxy xxx
type WebconsoleProxy struct {

	// Backend returns the backend URL which the proxy uses to reverse proxy
	Backend func(*http.Request) (*url.URL, error)

	CertConfig *config.CertConfig
}

// NewWebconsoleProxy xxx
func NewWebconsoleProxy(certConfig *config.CertConfig) *WebconsoleProxy {
	backend := func(req *http.Request) (*url.URL, error) {
		v := req.URL.Query()
		hostIp := v.Get("host_ip")
		if hostIp == "" {
			return nil, fmt.Errorf("param host_ip must not be empty")
		}
		consoleproxyPort := config.MesosWebconsoleProxyPort
		host := net.JoinHostPort(hostIp, strconv.FormatUint(uint64(consoleproxyPort), 10))

		u := &url.URL{}
		u.Host = host
		u.Fragment = req.URL.Fragment
		u.Path = strings.Replace(req.URL.Path, "webconsole", "consoleproxy", 1)
		u.RawQuery = req.URL.RawQuery
		blog.Infof(u.String())
		return u, nil
	}

	return &WebconsoleProxy{
		Backend:    backend,
		CertConfig: certConfig,
	}
}

// ServeHTTP 用于HTTP服务
func (w *WebconsoleProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	start := time.Now()

	backendURL, err := w.Backend(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	if websocket.IsWebSocketUpgrade(req) {
		websocketProxy := NewWebsocketProxy(w.CertConfig, backendURL)
		websocketProxy.ServeHTTP(rw, req)
		return
	}

	httpProxy, err := NewHttpReverseProxy(backendURL, w.CertConfig)
	if err != nil {
		metric.RequestErrorCount.WithLabelValues("mesos_webconsole", req.Method).Inc()
		metric.RequestErrorLatency.WithLabelValues("mesos_webconsole", req.Method).Observe(time.Since(start).Seconds())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	httpProxy.ServeHTTP(rw, req)
	metric.RequestCount.WithLabelValues("mesos_webconsole", req.Method).Inc()
	metric.RequestLatency.WithLabelValues("mesos_webconsole", req.Method).Observe(time.Since(start).Seconds())
	return
}
