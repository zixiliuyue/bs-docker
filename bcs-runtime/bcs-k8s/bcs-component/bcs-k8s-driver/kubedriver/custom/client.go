

package custom

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	urllib "net/url"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/kubedriver/options"

	"github.com/emicklei/go-restful"
)

// BcsClientAPIHandler xxx
type BcsClientAPIHandler struct {
	KubeMasterUrl string
	TLSConfig     options.TLSConfig
}

// Handler xxx
func (h *BcsClientAPIHandler) Handler(request *restful.Request, response *restful.Response) {
	subPath := request.PathParameter("subpath")
	targetPath := strings.Split(subPath, "bcsclient/")[1]

	rawRequest := request.Request
	body, err := ioutil.ReadAll(rawRequest.Body)
	if err != nil {
		CustomServerErrorResponse(response, "Reading raw request body failed")
		return
	}

	url := fmt.Sprintf("%s/%s",
		strings.TrimSuffix(h.KubeMasterUrl, "/"),
		strings.TrimPrefix(targetPath, "/"))

	// url param
	if rawRequest.URL.RawQuery != "" {
		url = fmt.Sprintf("%s?%s", url, rawRequest.URL.RawQuery)
	}
	proxyReq, _ := http.NewRequest(rawRequest.Method, url, bytes.NewReader(body))

	proxyReq.Header = make(http.Header)
	for key, value := range rawRequest.Header {
		proxyReq.Header[key] = value
	}

	// Respect client TLS config
	var httpClient *http.Client
	if h.IfKubeNeedTls() {
		tlsConfig, err := h.TLSConfig.ToConfigObj()
		if err != nil {
			CustomServerErrorResponse(response, err.Error())
			return
		}
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		httpClient = &http.Client{Transport: transport}
	} else {
		httpClient = &http.Client{}
	}

	blog.V(3).Infof("forwarding request to %s, method=%s", url, rawRequest.Method)
	resp, err := httpClient.Do(proxyReq)
	if err != nil {
		message := fmt.Sprintf("error request kube server: %s", err)
		blog.Warn(message)
		CustomServerErrorResponse(response, message)
		return
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	for key := range resp.Header {
		response.AddHeader(key, resp.Header.Get(key))
	}
	if resp.Header.Get("Content-Type") == "" {
		response.AddHeader("Content-Type", "application/json")
	}
	response.Write(respBody)
}

// Config xxx
func (h *BcsClientAPIHandler) Config(k8sMasterUrl string, tlsCfg options.TLSConfig) error {
	h.KubeMasterUrl = k8sMasterUrl
	h.TLSConfig = tlsCfg
	return nil
}

// IfKubeNeedTls xxx
func (h *BcsClientAPIHandler) IfKubeNeedTls() bool {
	kubeURL, _ := urllib.Parse(h.KubeMasterUrl)
	return kubeURL.Scheme == options.HTTPS
}
