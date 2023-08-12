

package cloudaccount

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"k8s.io/klog"
)

// CloudAccountMgr 云凭证管理
type CloudAccountMgr struct {
	ctx    context.Context
	client clustermanager.ClusterManagerClient
}

// New 创建云凭证管理实例
func New(ctx context.Context) types.CloudAccountMgr {
	client, cliCtx, err := pkg.NewClientWithConfiguration(ctx)
	if err != nil {
		klog.Fatalf("init client failed: %v", err.Error())
	}

	return &CloudAccountMgr{
		ctx:    cliCtx,
		client: client,
	}
}
