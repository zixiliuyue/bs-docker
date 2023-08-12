

package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	userManagerModels "github.com/Tencent/bk-bcs/bcs-services/bcs-user-manager/app/user-manager/models"
)

const (
	createClusterUrl       = "/v1/clusters"
	createRegisterTokenUrl = "/v1/clusters/%s/register_tokens" // nolint
	getRegisterTokenUrl    = "/v1/clusters/%s/register_tokens"
	updateCredentialsUrl   = "/v1/clusters/%s/credentials"
	getCredentialsUrl      = "/v1/clusters/%s/credentials"
	listCredentialsUrl     = "/v1/clusters/credentials"
)

// CreateClusterForm request form for create cluster
type CreateClusterForm struct {
	ClusterID        string `json:"cluster_id" validate:"required"`
	ClusterType      string `json:"cluster_type" validate:"required"`
	TkeClusterID     string `json:"tke_cluster_id"`
	TkeClusterRegion string `json:"tke_cluster_region"`
}

// CreateClusterResponse defines the response of create cluster
type CreateClusterResponse struct {
	Result  bool                          `json:"result"`
	Code    int                           `json:"code"`
	Message string                        `json:"message"`
	Data    *userManagerModels.BcsCluster `json:"data"`
}

// CreateCluster request create cluster from bcs-user-manager
func (c *UserManagerClient) CreateCluster(clusterBody string) (*CreateClusterResponse, error) {
	reqForm := new(CreateClusterForm)
	if err := json.Unmarshal([]byte(clusterBody), reqForm); err != nil {
		return nil, errors.Wrapf(err, "form json unmarshal failed with '%s'", reqForm)
	}
	bs, err := c.do(createClusterUrl, http.MethodPost, nil, reqForm)
	if err != nil {
		return nil, errors.Wrapf(err, "create cluster with '%s' failed", reqForm)
	}
	resp := new(CreateClusterResponse)
	if err := json.Unmarshal(bs, resp); err != nil {
		return nil, errors.Wrapf(err, "create cluster unmarshal failed with response '%s'", string(bs))
	}
	return resp, nil
}

// CreateRegisterTokenResponse defines the response of create register token
type CreateRegisterTokenResponse struct {
	Result  bool                                `json:"result"`
	Code    int                                 `json:"code"`
	Message string                              `json:"message"`
	Data    *userManagerModels.BcsRegisterToken `json:"data"`
}

// CreateRegisterToken request create register token  from bcs-user-manager
func (c *UserManagerClient) CreateRegisterToken(clusterId string) (*CreateRegisterTokenResponse, error) {
	bs, err := c.do(fmt.Sprintf(createRegisterTokenUrl, clusterId), http.MethodPost, nil, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "create register token with '%s' failed", clusterId)
	}
	resp := new(CreateRegisterTokenResponse)
	if err := json.Unmarshal(bs, resp); err != nil {
		return nil, errors.Wrapf(err, "create register token unmarshal failed with response '%s'", string(bs))
	}
	return resp, nil
}

// GetRegisterTokenResponse defines the response of get register token
type GetRegisterTokenResponse struct {
	Result  bool                                `json:"result"`
	Code    int                                 `json:"code"`
	Message string                              `json:"message"`
	Data    *userManagerModels.BcsRegisterToken `json:"data"`
}

// GetRegisterToken request get register token  from bcs-user-manager
func (c *UserManagerClient) GetRegisterToken(clusterId string) (*GetRegisterTokenResponse, error) {
	bs, err := c.do(fmt.Sprintf(getRegisterTokenUrl, clusterId), http.MethodGet, nil, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "get register token with '%s' failed", clusterId)
	}
	resp := new(GetRegisterTokenResponse)
	if err := json.Unmarshal(bs, resp); err != nil {
		return nil, errors.Wrapf(err, "get register token unmarshal failed with response '%s'", string(bs))
	}
	return resp, nil
}

// UpdateCredentialsForm update form for credential
type UpdateCredentialsForm struct {
	RegisterToken   string `json:"register_token" validate:"required"`
	ServerAddresses string `json:"server_addresses" validate:"required,apiserver_addresses"`
	CaCertData      string `json:"cacert_data" validate:"required"`
	UserToken       string `json:"user_token" validate:"required"`
}

// UpdateCredentialsResponse defines the response of update cluster credentials
type UpdateCredentialsResponse struct {
	Result  bool                                    `json:"result"`
	Code    int                                     `json:"code"`
	Message string                                  `json:"message"`
	Data    *userManagerModels.BcsClusterCredential `json:"data"`
}

// UpdateCredentials request update cluster credentials  from bcs-user-manager
func (c *UserManagerClient) UpdateCredentials(clusterId, credentialsForm string) (*UpdateCredentialsResponse, error) {
	reqForm := new(UpdateCredentialsForm)
	if err := json.Unmarshal([]byte(credentialsForm), reqForm); err != nil {
		return nil, errors.Wrapf(err, "credentials form json unmarshal failed with '%s'", reqForm)
	}
	bs, err := c.do(fmt.Sprintf(updateCredentialsUrl, clusterId), http.MethodPut, nil, reqForm)
	if err != nil {
		return nil, errors.Wrapf(err, "update cluster credentials with '%s' failed", reqForm)
	}
	resp := new(UpdateCredentialsResponse)
	if err := json.Unmarshal(bs, resp); err != nil {
		return nil, errors.Wrapf(err, "update cluster credentials unmarshal failed with response '%s'", string(bs))
	}
	return resp, nil
}

// GetCredentialsResponse defines the response of get cluster credentials
type GetCredentialsResponse struct {
	Result  bool                                    `json:"result"`
	Code    int                                     `json:"code"`
	Message string                                  `json:"message"`
	Data    *userManagerModels.BcsClusterCredential `json:"data"`
}

// GetCredentials request get cluster credentials  from bcs-user-manager
func (c *UserManagerClient) GetCredentials(clusterId string) (*GetCredentialsResponse, error) {
	bs, err := c.do(fmt.Sprintf(getCredentialsUrl, clusterId), http.MethodGet, nil, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "get cluster credentials with '%s' failed", clusterId)
	}
	resp := new(GetCredentialsResponse)
	if err := json.Unmarshal(bs, resp); err != nil {
		return nil, errors.Wrapf(err, "get cluster credentials unmarshal failed with response '%s'", string(bs))
	}
	return resp, nil
}

// CredentialResp response
type CredentialResp struct {
	ServerAddresses string `json:"server_addresses"`
	CaCertData      string `json:"ca_cert_data"`
	UserToken       string `json:"user_token"`
	ClusterDomain   string `json:"cluster_domain"`
}

// ListCredentialsResponse defines the response of list cluster credentials
type ListCredentialsResponse struct {
	Result  bool                      `json:"result"`
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Data    map[string]CredentialResp `json:"data"`
}

// ListCredentials request list cluster credentials  from bcs-user-manager
func (c *UserManagerClient) ListCredentials() (*ListCredentialsResponse, error) {
	bs, err := c.do(listCredentialsUrl, http.MethodGet, nil, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "list cluster credentials failed")
	}
	resp := new(ListCredentialsResponse)
	if err := json.Unmarshal(bs, resp); err != nil {
		return nil, errors.Wrapf(err, "list cluster credentials unmarshal failed with response '%s'", string(bs))
	}
	return resp, nil
}
