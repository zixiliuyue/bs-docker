

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Delete 删除节点池
func (c *NodeGroupMgr) Delete(req types.DeleteNodeGroupReq) (types.DeleteNodeGroupResp, error) {
	var (
		resp types.DeleteNodeGroupResp
		err  error
	)

	servResp, err := c.client.DeleteNodeGroup(c.ctx, &clustermanager.DeleteNodeGroupRequest{
		NodeGroupID:           req.NodeGroupID,
		IsForce:               false,
		ReserveNodesInCluster: false,
		KeepNodesInstance:     false,
		Operator:              "bcs",
	})

	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.TaskID = servResp.GetData().GetTask().GetTaskID()

	return resp, nil
}
