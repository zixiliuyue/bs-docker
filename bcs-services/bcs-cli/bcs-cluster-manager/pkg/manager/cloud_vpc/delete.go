

package cloudvpc

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Delete 删除云私有网络
func (c *CloudVPCMgr) Delete(req types.DeleteCloudVPCReq) error {
	resp, err := c.client.DeleteCloudVPC(c.ctx, &clustermanager.DeleteCloudVPCRequest{
		CloudID: req.CloudID,
		VpcID:   req.VPCID,
	})
	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
