

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// RemoveNodes 指定IP地址从节点组中移除,但是该节点仍然在集群中
func (c *NodeGroupMgr) RemoveNodes(req types.RemoveNodesFromGroupReq) (types.RemoveNodesFromGroupResp, error) {
	var (
		resp types.RemoveNodesFromGroupResp
		err  error
	)

	servResp, err := c.client.RemoveNodesFromGroup(c.ctx, &clustermanager.RemoveNodesFromGroupRequest{
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
