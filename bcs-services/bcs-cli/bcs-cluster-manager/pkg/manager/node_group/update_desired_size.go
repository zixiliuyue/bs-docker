

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// UpdateDesiredSize 更新节点池期望的节点数,扩容前保证数据一致性
func (c *NodeGroupMgr) UpdateDesiredSize(req types.UpdateGroupDesiredSizeReq) error {
	resp, err := c.client.UpdateGroupDesiredSize(c.ctx, &clustermanager.UpdateGroupDesiredSizeRequest{
		NodeGroupID: req.NodeGroupID,
		DesiredSize: req.DesiredSize,
		Operator:    "bcs",
	})

	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
