

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// DisableAutoScale 关闭节点池自动扩缩容
func (c *NodeGroupMgr) DisableAutoScale(req types.DisableNodeGroupAutoScaleReq) error {
	resp, err := c.client.DisableNodeGroupAutoScale(c.ctx, &clustermanager.DisableNodeGroupAutoScaleRequest{
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
