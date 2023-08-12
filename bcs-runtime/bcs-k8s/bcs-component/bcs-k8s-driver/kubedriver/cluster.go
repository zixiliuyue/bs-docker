

package kubedriver

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/kubedriver/options"

	"github.com/parnurzeal/gorequest"
)

// AccessTokenRequestBody xxx
type AccessTokenRequestBody struct {
	EnvName    string `json:"env_name"`
	AppCode    string `json:"app_code"`
	AppSecret  string `json:"app_secret"`
	IDProvider string `json:"id_provider"`
	GrantType  string `json:"grant_type"`
}

// APIGWAuthHeaders xxx
type APIGWAuthHeaders struct {
	AccessToken string `json:"access_token"`
}

// GetClusterID xxx
func GetClusterID(o *options.KubeDriverServerOptions) (clusterID string, err error) {
	if o.Environment == "develop" {
		return "driver-debug-clusterID", nil
	}
	if o.CustomClusterID != "" {
		return o.CustomClusterID, nil
	}
	err = o.GetClusterKeeperAddr()
	if nil != err {
		return "", fmt.Errorf("get cluster keeper  api server addr failed. reason: %v", err)
	}

	goReq := gorequest.New()
	if o.NeedClusterTLSConfig() {
		clusterTLSConfig, err := o.ClusterClientTLS.ToConfigObj()
		if err != nil {
			return "", fmt.Errorf("config cluster keeper tls failed, reason: %v", err)
		}
		goReq = goReq.TLSClientConfig(clusterTLSConfig)
	}

	generatedUrl := fmt.Sprintf("%s%s", o.ClusterKeeperUrl, "/bcsclusterkeeper/v1/cluster/id/byip")
	resp, respBody, errs := goReq.Get(generatedUrl).Query(map[string]string{"ip": o.HostIP}).End()
	if len(errs) != 0 {
		return "", errs[0]
	}

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("cluster keeper request error, status code: %d, url: %s", resp.StatusCode, o.ClusterKeeperUrl)
	}

	clusterID = json.Get([]byte(respBody), "data").Get("clusterID").ToString()
	return clusterID, nil
}
