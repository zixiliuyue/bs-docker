

package cloudvpc

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Update 更新云私有网络
func (c *CloudVPCMgr) Update(req types.UpdateCloudVPCReq) error {
	resp, err := c.client.UpdateCloudVPC(c.ctx, &clustermanager.UpdateCloudVPCRequest{
		CloudID:     req.CloudID,
		NetworkType: req.NetworkType,
		Region:      req.Region,
		RegionName:  req.RegionName,
		VpcName:     req.VPCName,
		VpcID:       req.VPCID,
		Available:   req.Available,
		Updater:     "bcs",
	})
	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
