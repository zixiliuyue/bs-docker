

package node

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"k8s.io/klog"
)

// NodeMgr 节点管理
type NodeMgr struct {
	ctx    context.Context
	client clustermanager.ClusterManagerClient
}

// New 创建节点管理实例
func New(ctx context.Context) types.NodeMgr {
	client, cliCtx, err := pkg.NewClientWithConfiguration(ctx)
	if err != nil {
		klog.Fatalf("init client failed: %v", err.Error())
	}

	return &NodeMgr{
		ctx:    cliCtx,
		client: client,
	}
}
