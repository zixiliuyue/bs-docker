

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// CleanNodes 从节点池中清理回收指定IP节点
func (c *NodeGroupMgr) CleanNodes(req types.CleanNodeGroupReq) (types.CleanNodeGroupResp, error) {
	var (
		resp types.CleanNodeGroupResp
		err  error
	)

	servResp, err := c.client.CleanNodesInGroup(c.ctx, &clustermanager.CleanNodesInGroupRequest{
		ClusterID:   req.ClusterID,
		Nodes:       req.Nodes,
		NodeGroupID: req.NodeGroupID,
		Operator:    "bcs",
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
