

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// UpdateDesiredNode 更新节点池中DesiredNode信息
func (c *NodeGroupMgr) UpdateDesiredNode(req types.UpdateGroupDesiredNodeReq) (types.UpdateGroupDesiredNodeResp, error) {
	var (
		resp types.UpdateGroupDesiredNodeResp
		err  error
	)

	servResp, err := c.client.UpdateGroupDesiredNode(c.ctx, &clustermanager.UpdateGroupDesiredNodeRequest{
		NodeGroupID: req.NodeGroupID,
		DesiredNode: req.DesiredNode,
		Operator:    "bcs",
	})

	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.TaskID = servResp.Data.TaskID

	return resp, nil
}
