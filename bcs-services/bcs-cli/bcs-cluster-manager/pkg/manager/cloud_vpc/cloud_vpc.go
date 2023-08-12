

package cloudvpc

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"k8s.io/klog"
)

// CloudVPCMgr 云私有网络管理
type CloudVPCMgr struct {
	ctx    context.Context
	client clustermanager.ClusterManagerClient
}

// New 创建云私有网络管理实例
func New(ctx context.Context) types.CloudVPCMgr {
	client, cliCtx, err := pkg.NewClientWithConfiguration(ctx)
	if err != nil {
		klog.Fatalf("init client failed: %v", err.Error())
	}

	return &CloudVPCMgr{
		ctx:    cliCtx,
		client: client,
	}
}
