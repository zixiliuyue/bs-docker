

package cluster

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Delete 删除集群
func (c *ClusterMgr) Delete(req types.DeleteClusterReq) (types.DeleteClusterResp, error) {
	var (
		resp types.DeleteClusterResp
		err  error
	)

	servResp, err := c.client.DeleteCluster(c.ctx, &clustermanager.DeleteClusterReq{
		ClusterID:      req.ClusterID,
		OnlyDeleteInfo: true,
		Operator:       "bcs",
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
