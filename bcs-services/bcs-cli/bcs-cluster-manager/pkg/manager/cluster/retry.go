

package cluster

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// RetryCreate 重试创建集群
func (c *ClusterMgr) RetryCreate(req types.RetryCreateClusterReq) (types.RetryCreateClusterResp, error) {
	var (
		resp types.RetryCreateClusterResp
		err  error
	)

	servResp, err := c.client.RetryCreateClusterTask(c.ctx, &clustermanager.RetryCreateClusterReq{
		ClusterID: req.ClusterID,
		Operator:  "bcs",
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.ClusterID = servResp.Data.GetClusterID()
	resp.TaskID = servResp.Task.GetTaskID()

	return resp, nil
}
