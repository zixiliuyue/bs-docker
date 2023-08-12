

package cluster

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"k8s.io/klog"
)

// ClusterMgr 集群管理实例
type ClusterMgr struct {
	ctx    context.Context
	client clustermanager.ClusterManagerClient
}

// New 返回集群管理实例
func New(ctx context.Context) types.ClusterMgr {
	client, cliCtx, err := pkg.NewClientWithConfiguration(ctx)
	if err != nil {
		klog.Fatalf("init client failed: %v", err.Error())
	}

	return &ClusterMgr{
		ctx:    cliCtx,
		client: client,
	}
}

// NewHttpClient 返回http client实例
func NewHttpClient(ctx context.Context) *pkg.HttpClient {
	return pkg.NewHttpClientWithConfiguration(ctx)
}
