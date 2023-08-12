

package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	netUrl "net/url"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config describe the options Client need
type Config struct {
	// APIServer for bcs-api-gateway address
	APIServer string
	// AuthToken for bcs permission token
	AuthToken string
	// Operator for the bk-repo operations
	Operator string
}

// UserManagerClient defines the client for bcs-user-manager
type UserManagerClient struct {
	cfg *Config
	ctx context.Context
}

// NewClientWithConfiguration new client with config
func NewClientWithConfiguration(ctx context.Context) *UserManagerClient {
	return &UserManagerClient{
		ctx: ctx,
		cfg: &Config{
			APIServer: viper.GetString("config.apiserver"),
			AuthToken: viper.GetString("config.bcs_token"),
			Operator:  viper.GetString("config.operator"),
		},
	}
}

func (c *UserManagerClient) do(url string, httpType string, query netUrl.Values, body interface{}) ([]byte, error) {
	url = c.cfg.APIServer + url
	var req *http.Request
	var err error

	_, err = netUrl.Parse(c.cfg.APIServer)
	if err != nil {
		return nil, fmt.Errorf("url failed %v", err)
	}

	if body != nil {
		var bs []byte
		bs, err = json.Marshal(body)
		if err != nil {
			return nil, errors.Wrapf(err, "marshal body failed")
		}
		req, err = http.NewRequestWithContext(c.ctx, httpType, url, bytes.NewReader(bs))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequestWithContext(c.ctx, httpType, url, nil)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "create request failed")
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	// 添加鉴权
	if len(c.cfg.AuthToken) != 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.cfg.AuthToken))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "http do request failed")
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "read response body failed")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errors.Errorf(string(bs)), "http response status not 200 but %d",
			resp.StatusCode)
	}
	return bs, nil
}
