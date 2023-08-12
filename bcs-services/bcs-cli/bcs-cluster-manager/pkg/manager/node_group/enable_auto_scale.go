

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// EnableAutoScale 开启节点池自动扩缩容
func (c *NodeGroupMgr) EnableAutoScale(req types.EnableNodeGroupAutoScaleReq) error {
	resp, err := c.client.EnableNodeGroupAutoScale(c.ctx, &clustermanager.EnableNodeGroupAutoScaleRequest{
		NodeGroupID: req.NodeGroupID,
	})

	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
