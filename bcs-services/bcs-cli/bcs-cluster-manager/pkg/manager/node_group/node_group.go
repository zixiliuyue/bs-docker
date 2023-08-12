

package nodegroup

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"k8s.io/klog"
)

// NodeGroupMgr 节点池管理
type NodeGroupMgr struct {
	ctx    context.Context
	client clustermanager.ClusterManagerClient
}

// New 创建节点池管理实例
func New(ctx context.Context) types.NodeGroupMgr {
	client, cliCtx, err := pkg.NewClientWithConfiguration(ctx)
	if err != nil {
		klog.Fatalf("init client failed: %v", err.Error())
	}

	return &NodeGroupMgr{
		ctx:    cliCtx,
		client: client,
	}
}
