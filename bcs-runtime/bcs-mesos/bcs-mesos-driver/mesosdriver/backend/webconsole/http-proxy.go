

package webconsole

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/ssl"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-driver/mesosdriver/config"
)

// NewHttpReverseProxy create golang library reverse proxy for normal http request
func NewHttpReverseProxy(target *url.URL, certConfig *config.CertConfig) (*httputil.ReverseProxy, error) {
	if certConfig.IsSSL {
		target.Scheme = "https"
	}

	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		req.URL.RawQuery = target.RawQuery
	}
	reverseProxy := &httputil.ReverseProxy{Director: director}
	if certConfig.IsSSL {
		cliTls, err := ssl.ClientTslConfVerity(certConfig.CAFile, certConfig.CertFile, certConfig.KeyFile,
			certConfig.CertPasswd)
		if err != nil {
			blog.Errorf("set client tls config error %s", err.Error())
			return nil, fmt.Errorf("set client tls config error %s", err.Error())
		}
		reverseProxy.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 30 * time.Second,
			TLSClientConfig:     cliTls,
		}
	}
	return reverseProxy, nil
}
