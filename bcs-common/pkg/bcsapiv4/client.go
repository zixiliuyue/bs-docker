

package bcsapiv4

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"

	registry "github.com/Tencent/bk-bcs/bcs-common/pkg/registryv4"
)

// ! v4 version binding~

const (
	gatewayPrefix   = "/bcsapi/v4/"
	clusterIDHeader = "BCS-ClusterID"
)

// Config for bcsapi
type Config struct {
	// bcsapi host, available like 127.0.0.1:8080
	Hosts []string
	// tls configuratio
	TLSConfig *tls.Config
	// AuthToken for permission verification
	AuthToken string
	// clusterID for Kubernetes/Mesos operation
	ClusterID string
	// proxy flag for go through bcs-api-gateway
	Gateway bool
	// etcd registry config for bcs modules
	Etcd registry.CMDOptions
	// Header for request header
	Header map[string]string
	// InnerClientName for bcs inner auth, like bcs-cluster-manager
	InnerClientName string
}

// BasicResponse basic http response for bkbcs
type BasicResponse struct {
	Code    int             `json:"code"`
	Result  bool            `json:"result"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// Authentication defines the common interface for the credentials which need to
// attach auth info to every RPC
type Authentication struct {
	InnerClientName string
	Insecure        bool
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{"X-Bcs-Client": a.InnerClientName}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (a *Authentication) RequireTransportSecurity() bool {
	return !a.Insecure
}

// NewTokenAuth implementations of grpc credentials interface
func NewTokenAuth(t string) *GrpcTokenAuth {
	return &GrpcTokenAuth{
		Token: t,
	}
}

// GrpcTokenAuth grpc token
type GrpcTokenAuth struct {
	Token string
}

// GetRequestMetadata convert http Authorization for grpc key
func (t GrpcTokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", t.Token),
	}, nil
}

// RequireTransportSecurity RequireTransportSecurity
func (t GrpcTokenAuth) RequireTransportSecurity() bool {
	return false
}

// NewClient create new bcsapi instance
func NewClient(config *Config) *Client {
	return &Client{
		config: config,
	}
}

// Client all module client api composition
type Client struct {
	config *Config
}

// UserManager client interface
//func (c *Client) UserManager() UserManager {
//	return NewUserManager(c.config)
//}

// MesosDriver client interface
// func (c *Client) MesosDriver() MesosDriver {
// 	return &MesosDriverClient{}
// }

//// Storage client interface
//func (c *Client) Storage() Storage {
//	return NewStorage(c.config)
//}
