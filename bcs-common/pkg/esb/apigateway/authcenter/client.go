

package authcenter

import (
	"crypto/tls"
	"fmt"

	paasclient "github.com/Tencent/bk-bcs/bcs-common/pkg/esb/client"
)

// Config item for BlueKing Auth Center
// AuthCenter requires AccessToken for authorization
type Config struct {
	// Hosts AuthCenter hosts, default https
	Hosts []string
	// config for https, when setting tls, use https instead
	TLSConfig   *tls.Config
	AccessToken string
}

// Client BlueKing Auth Center interface difinition
type Client interface {
	// QueryProjectUsers list all project users
	QueryProjectUsers(projectID string) ([]string, error)
}

// NewAuthClient create authClient instance
func NewAuthClient(cfg *Config) (Client, error) {
	// validate config
	if len(cfg.Hosts) == 0 {
		return nil, fmt.Errorf("Lost hosts config item(required)")
	}
	if len(cfg.AccessToken) == 0 {
		return nil, fmt.Errorf("lost AccessToken(required)")
	}
	var c *authClient
	if cfg.TLSConfig != nil {
		c = &authClient{
			config: cfg,
			client: paasclient.NewRESTClientWithTLS(cfg.TLSConfig),
		}
	} else {
		c = &authClient{
			config: cfg,
			client: paasclient.NewRESTClient(),
		}
	}
	return c, nil
}

// authClient auth center sdk implementation
type authClient struct {
	config *Config
	client *paasclient.RESTClient
}

// QueryProjectUsers list all project users
func (c *authClient) QueryProjectUsers(projectID string) ([]string, error) {
	response := &UsersQueryResponse{}
	err := c.client.Get().
		WithEndpoints(c.config.Hosts).
		WithBasePath("/").
		SubPathf("/projects/%s/users", projectID).
		WithParam("access_token", c.config.AccessToken).
		Do().
		Into(response)
	if err != nil {
		return nil, err
	}
	if response.Code != 0 {
		return nil, fmt.Errorf("query project user list failed, %s", response.Message)
	}
	return response.Data, nil
}
