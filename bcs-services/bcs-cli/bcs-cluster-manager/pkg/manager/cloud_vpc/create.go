

package cloudvpc

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Create 创建云私有网络
func (c *CloudVPCMgr) Create(req types.CreateCloudVPCReq) error {
	resp, err := c.client.CreateCloudVPC(c.ctx, &clustermanager.CreateCloudVPCRequest{
		CloudID:     req.CloudID,
		NetworkType: req.NetworkType,
		Region:      req.Region,
		VpcName:     req.VPCName,
		VpcID:       req.VPCID,
		Available:   req.Available,
		Creator:     "bcs",
	})
	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
