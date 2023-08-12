

package bcsapi

import (
	"fmt"
	"net/http"

	restclient "github.com/Tencent/bk-bcs/bcs-common/pkg/esb/client"

	"github.com/parnurzeal/gorequest"
)

func configureRequest(r *gorequest.SuperAgent, config *Config) *gorequest.SuperAgent {
	// setting insecureSkipVerify
	if config.TLSConfig != nil {
		config.TLSConfig.InsecureSkipVerify = true
		r.TLSClientConfig(config.TLSConfig)
	}
	if config.AuthToken != "" {
		r.Set("Authorization", fmt.Sprintf("Bearer %s", config.AuthToken))
	}
	if config.ClusterID != "" {
		r.Set(clusterIDHeader, config.ClusterID)
	}
	return r
}

func newGet(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Get(address)
	return configureRequest(r, config)
}

func newPost(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Post(address)
	return configureRequest(r, config)
}

func newDelete(config *Config, address string) *gorequest.SuperAgent {
	r := gorequest.New().Delete(address)
	return configureRequest(r, config)
}

func bkbcsSetting(req *restclient.Request, config *Config) *restclient.Request {
	header := make(http.Header)
	if config.AuthToken != "" {
		header.Add("Authorization", fmt.Sprintf("Bearer %s", config.AuthToken))
	}
	if config.ClusterID != "" {
		header.Add(clusterIDHeader, config.ClusterID)
	}
	return req.WithHeaders(header)
}
