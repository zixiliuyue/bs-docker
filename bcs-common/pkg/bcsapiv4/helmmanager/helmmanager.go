

package helmmanager

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

// NewHelmClient create HelmManager SDK implementation
func NewHelmClient(config *bcsapi.Config) (HelmManagerClient, func()) {
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
			blog.Errorf("Create helm manager grpc client with %s error: %s", addr, err.Error())
			continue
		}
		if conn != nil {
			break
		}
	}
	if conn == nil {
		blog.Errorf("create no helm manager client after all instance tries")
		return nil, nil
	}
	return NewHelmManagerClient(conn), func() { conn.Close() }
}

// NewHelmAddonsClient create HelmManager addons SDK implementation
func NewHelmAddonsClient(config *bcsapi.Config) (ClusterAddonsClient, func()) {
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
			blog.Errorf("Create helm manager grpc client with %s error: %s", addr, err.Error())
			continue
		}
		if conn != nil {
			break
		}
	}
	if conn == nil {
		blog.Errorf("create no helm manager client after all instance tries")
		return nil, nil
	}
	return NewClusterAddonsClient(conn), func() { conn.Close() }
}
