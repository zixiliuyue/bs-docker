

package appauth

import (
	"crypto/tls"
	"fmt"

	paasclient "github.com/Tencent/bk-bcs/bcs-common/pkg/esb/client"
)

// Config item for BlueKing Auth Center
type Config struct {
	// Hosts AuthCenter hosts, without http/https, default is http
	Hosts []string
	// config for https
	TLSConfig *tls.Config
	AppCode   string
	AppSecret string
}

// Client BlueKing Auth Center interface difinition
type Client interface {
	// GetAccessToken get specified environmental BK PaaS access token
	// by App identifier. Access token expires in 180 days
	GetAccessToken(env string) (string, error)
}

// NewAuthClient create authClient instance
func NewAuthClient(cfg *Config) (Client, error) {
	// validate config
	if len(cfg.Hosts) == 0 {
		return nil, fmt.Errorf("Lost hosts config item(required)")
	}
	if len(cfg.AppCode) == 0 || len(cfg.AppSecret) == 0 {
		return nil, fmt.Errorf("Lost Authorization information(required)")
	}
	var c *authClient
	if cfg.TLSConfig != nil {
		c = &authClient{
			config: cfg,
			client: paasclient.NewRESTClientWithTLS(cfg.TLSConfig).
				WithCredential(map[string]interface{}{
					"app_code":   cfg.AppCode,
					"app_secret": cfg.AppSecret,
				}),
		}
	} else {
		c = &authClient{
			config: cfg,
			client: paasclient.NewRESTClient().
				WithCredential(map[string]interface{}{
					"app_code":   cfg.AppCode,
					"app_secret": cfg.AppSecret,
				}),
		}
	}
	return c, nil
}

// authClient auth center sdk implementation
type authClient struct {
	config *Config
	client *paasclient.RESTClient
}

// GetAccessToken get specified environmental BK PaaS access token
func (c *authClient) GetAccessToken(env string) (string, error) {
	if !(env == PaasOAuthEnvProd || env == PaasOAuthEnvTest) {
		return "", fmt.Errorf("Error Environment")
	}
	request := map[string]interface{}{
		"env_name":   env,
		"app_code":   c.config.AppCode,
		"app_secret": c.config.AppSecret,
		"grant_type": PaasGrantTypeClient,
	}
	var response OAuthResponse
	err := c.client.Post().
		WithEndpoints(c.config.Hosts).
		WithBasePath("/").
		SubPathf("/auth_api/token/").
		Body(request).
		Do().
		Into(&response)
	if err != nil {
		return "", err
	}
	if response.Code != 0 {
		return "", fmt.Errorf("GetAccessToken failed, %s", response.Message)
	}
	return response.Data.AccessToken, nil
}
