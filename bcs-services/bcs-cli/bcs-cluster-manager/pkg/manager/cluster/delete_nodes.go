

package cluster

import (
	"errors"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// DeleteNodes 从集群中删除节点
func (c *ClusterMgr) DeleteNodes(req types.DeleteNodesClusterReq) (types.DeleteNodesClusterResp, error) {
	var (
		resp types.DeleteNodesClusterResp
		err  error
	)

	servResp, err := c.client.DeleteNodesFromCluster(c.ctx, &clustermanager.DeleteNodesRequest{
		ClusterID:      req.ClusterID,
		Nodes:          strings.Join(req.Nodes, ","),
		Operator:       "bcs",
		IsForce:        true,
		OnlyDeleteInfo: true,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.TaskID = servResp.Data.GetTaskID()

	return resp, nil
}
