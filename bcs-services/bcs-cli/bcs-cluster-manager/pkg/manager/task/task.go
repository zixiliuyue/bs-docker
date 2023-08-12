

package task

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"k8s.io/klog"
)

// TaskMgr 任务管理
type TaskMgr struct {
	ctx    context.Context
	client clustermanager.ClusterManagerClient
}

// New 创建任务管理实例
func New(ctx context.Context) types.TaskMgr {
	client, cliCtx, err := pkg.NewClientWithConfiguration(ctx)
	if err != nil {
		klog.Fatalf("init client failed: %v", err.Error())
	}

	return &TaskMgr{
		ctx:    cliCtx,
		client: client,
	}
}
