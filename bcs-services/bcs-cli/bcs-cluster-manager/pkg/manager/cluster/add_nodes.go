

package cluster

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// AddNodes 添加节点到集群
func (c *ClusterMgr) AddNodes(req types.AddNodesClusterReq) (types.AddNodesClusterResp, error) {
	var (
		resp types.AddNodesClusterResp
		err  error
	)

	servResp, err := c.client.AddNodesToCluster(c.ctx, &clustermanager.AddNodesRequest{
		ClusterID:         req.ClusterID,
		Nodes:             req.Nodes,
		InitLoginPassword: req.InitPassword,
		OnlyCreateInfo:    true,
		Operator:          "bcs",
	})
	if err != nil {
		return servResp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.TaskID = servResp.Data.GetTaskID()

	return resp, nil
}
