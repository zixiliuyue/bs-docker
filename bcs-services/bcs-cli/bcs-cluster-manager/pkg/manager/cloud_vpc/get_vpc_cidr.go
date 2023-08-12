

package cloudvpc

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// GetVPCCidr 查询云私有网络cidr列表
func (c *CloudVPCMgr) GetVPCCidr(req types.GetVPCCidrReq) (types.GetVPCCidrResp, error) {
	var (
		resp types.GetVPCCidrResp
		err  error
	)

	servResp, err := c.client.GetVPCCidr(c.ctx, &clustermanager.GetVPCCidrRequest{
		VpcID: req.VPCID,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.VPCCidr, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, &types.VPCCidr{
			VPC:      v.Vpc,
			Cidr:     v.Cidr,
			IPNumber: v.IPNumber,
			Status:   v.Status,
		})
	}

	return resp, nil
}
