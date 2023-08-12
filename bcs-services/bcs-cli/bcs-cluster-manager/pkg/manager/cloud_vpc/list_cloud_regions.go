

package cloudvpc

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// ListCloudRegions 查询云region列表
func (c *CloudVPCMgr) ListCloudRegions(req types.ListCloudRegionsReq) (types.ListCloudRegionsResp, error) {
	var (
		resp types.ListCloudRegionsResp
		err  error
	)

	servResp, err := c.client.ListCloudRegions(c.ctx, &clustermanager.ListCloudRegionsRequest{
		CloudID: req.CloudID,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.CloudRegion, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, &types.CloudRegion{
			CloudID:    v.CloudID,
			Region:     v.Region,
			RegionName: v.RegionName,
		})
	}

	return resp, nil
}
