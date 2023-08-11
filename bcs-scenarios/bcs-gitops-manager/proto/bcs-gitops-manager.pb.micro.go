// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bcs-gitops-manager.proto

package bcsgitopsmanager

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for BcsGitopsManager service

func NewBcsGitopsManagerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "BcsGitopsManager.Ping",
			Path:    []string{"/gitopsmanager/v1/ping"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "BcsGitopsManager.StartupProject",
			Path:    []string{"/gitopsmanager/v1/project"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "BcsGitopsManager.SyncCluster",
			Path:    []string{"/gitopsmanager/v1/cluster/sync"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
	}
}

// Client API for BcsGitopsManager service

type BcsGitopsManagerService interface {
	Ping(ctx context.Context, in *GitOpsRequest, opts ...client.CallOption) (*GitOpsResponse, error)
	StartupProject(ctx context.Context, in *ProjectSyncRequest, opts ...client.CallOption) (*GitOpsResponse, error)
	SyncCluster(ctx context.Context, in *ClusterSyncRequest, opts ...client.CallOption) (*GitOpsResponse, error)
}

type bcsGitopsManagerService struct {
	c    client.Client
	name string
}

func NewBcsGitopsManagerService(name string, c client.Client) BcsGitopsManagerService {
	return &bcsGitopsManagerService{
		c:    c,
		name: name,
	}
}

func (c *bcsGitopsManagerService) Ping(ctx context.Context, in *GitOpsRequest, opts ...client.CallOption) (*GitOpsResponse, error) {
	req := c.c.NewRequest(c.name, "BcsGitopsManager.Ping", in)
	out := new(GitOpsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcsGitopsManagerService) StartupProject(ctx context.Context, in *ProjectSyncRequest, opts ...client.CallOption) (*GitOpsResponse, error) {
	req := c.c.NewRequest(c.name, "BcsGitopsManager.StartupProject", in)
	out := new(GitOpsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcsGitopsManagerService) SyncCluster(ctx context.Context, in *ClusterSyncRequest, opts ...client.CallOption) (*GitOpsResponse, error) {
	req := c.c.NewRequest(c.name, "BcsGitopsManager.SyncCluster", in)
	out := new(GitOpsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BcsGitopsManager service

type BcsGitopsManagerHandler interface {
	Ping(context.Context, *GitOpsRequest, *GitOpsResponse) error
	StartupProject(context.Context, *ProjectSyncRequest, *GitOpsResponse) error
	SyncCluster(context.Context, *ClusterSyncRequest, *GitOpsResponse) error
}

func RegisterBcsGitopsManagerHandler(s server.Server, hdlr BcsGitopsManagerHandler, opts ...server.HandlerOption) error {
	type bcsGitopsManager interface {
		Ping(ctx context.Context, in *GitOpsRequest, out *GitOpsResponse) error
		StartupProject(ctx context.Context, in *ProjectSyncRequest, out *GitOpsResponse) error
		SyncCluster(ctx context.Context, in *ClusterSyncRequest, out *GitOpsResponse) error
	}
	type BcsGitopsManager struct {
		bcsGitopsManager
	}
	h := &bcsGitopsManagerHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BcsGitopsManager.Ping",
		Path:    []string{"/gitopsmanager/v1/ping"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BcsGitopsManager.StartupProject",
		Path:    []string{"/gitopsmanager/v1/project"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BcsGitopsManager.SyncCluster",
		Path:    []string{"/gitopsmanager/v1/cluster/sync"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&BcsGitopsManager{h}, opts...))
}

type bcsGitopsManagerHandler struct {
	BcsGitopsManagerHandler
}

func (h *bcsGitopsManagerHandler) Ping(ctx context.Context, in *GitOpsRequest, out *GitOpsResponse) error {
	return h.BcsGitopsManagerHandler.Ping(ctx, in, out)
}

func (h *bcsGitopsManagerHandler) StartupProject(ctx context.Context, in *ProjectSyncRequest, out *GitOpsResponse) error {
	return h.BcsGitopsManagerHandler.StartupProject(ctx, in, out)
}

func (h *bcsGitopsManagerHandler) SyncCluster(ctx context.Context, in *ClusterSyncRequest, out *GitOpsResponse) error {
	return h.BcsGitopsManagerHandler.SyncCluster(ctx, in, out)
}