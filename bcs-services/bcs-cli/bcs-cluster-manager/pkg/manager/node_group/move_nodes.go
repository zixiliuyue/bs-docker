

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// MoveNodes 将集群中节点移入节点池
func (c *NodeGroupMgr) MoveNodes(req types.MoveNodesToGroupReq) (types.MoveNodesToGroupResp, error) {
	var (
		resp types.MoveNodesToGroupResp
		err  error
	)

	servResp, err := c.client.MoveNodesToGroup(c.ctx, &clustermanager.MoveNodesToGroupRequest{
		ClusterID:   req.ClusterID,
		Nodes:       req.Nodes,
		NodeGroupID: req.NodeGroupID,
	})

	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.TaskID = servResp.GetData().GetTaskID()

	return resp, nil
}
