

package cluster

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// ListNodes 查询集群下所有节点列表
func (c *ClusterMgr) ListNodes(req types.ListClusterNodesReq) (types.ListClusterNodesResp, error) {
	var (
		resp types.ListClusterNodesResp
		err  error
	)

	servResp, err := c.client.ListNodesInCluster(c.ctx, &clustermanager.ListNodesInClusterRequest{
		ClusterID: req.ClusterID,
		Offset:    req.Offset,
		Limit:     req.Limit,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.ClusterNode, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, &types.ClusterNode{
			NodeID:  v.NodeID,
			InnerIP: v.InnerIP,
		})
	}

	return resp, nil
}
