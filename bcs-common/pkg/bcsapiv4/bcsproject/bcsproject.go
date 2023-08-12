

package bcsproject

import (
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	blog "k8s.io/klog/v2"

	bcsapi "github.com/Tencent/bk-bcs/bcs-common/pkg/bcsapiv4"
)

// ProjectClient xxx
type ProjectClient struct {
	Healthz   HealthzClient
	Project   BCSProjectClient
	Namespace NamespaceClient
	Variable  VariableClient
}

// NewProjectManagerClient create ProjectManager SDK implementation
func NewProjectManagerClient(config *bcsapi.Config) (*ProjectClient, func()) {
	rand.Seed(time.Now().UnixNano())
	if len(config.Hosts) == 0 {
		// ! pay more attention for nil return
		return nil, nil
	}
	// create grpc connection
	header := map[string]string{
		"x-content-type": "application/grpc+proto",
		"Content-Type":   "application/grpc",
	}
	if len(config.AuthToken) != 0 {
		header["Authorization"] = fmt.Sprintf("Bearer %s", config.AuthToken)
	}
	for k, v := range config.Header {
		header[k] = v
	}
	md := metadata.New(header)
	auth := &bcsapi.Authentication{InnerClientName: config.InnerClientName}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.Header(&md)))
	if config.TLSConfig != nil {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config.TLSConfig)))
	} else {
		opts = append(opts, grpc.WithInsecure())
		auth.Insecure = true
	}
	opts = append(opts, grpc.WithPerRPCCredentials(auth))
	if config.AuthToken != "" {
		opts = append(opts, grpc.WithPerRPCCredentials(bcsapi.NewTokenAuth(config.AuthToken)))
	}

	var conn *grpc.ClientConn
	var err error
	maxTries := 3
	for i := 0; i < maxTries; i++ {
		selected := rand.Intn(1024) % len(config.Hosts)
		addr := config.Hosts[selected]
		conn, err = grpc.Dial(addr, opts...)
		if err != nil {
			blog.Errorf("Create project manager grpc client with %s error: %s", addr, err.Error())
			continue
		}
		if conn != nil {
			break
		}
	}
	if conn == nil {
		blog.Errorf("create no project manager client after all instance tries")
		return nil, nil
	}
	// init project manager client
	c := &ProjectClient{
		Healthz:   NewHealthzClient(conn),
		Project:   NewBCSProjectClient(conn),
		Namespace: NewNamespaceClient(conn),
		Variable:  NewVariableClient(conn),
	}
	return c, func() { conn.Close() }
}
