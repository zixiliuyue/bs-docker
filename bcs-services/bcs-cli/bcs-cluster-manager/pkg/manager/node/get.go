

package node

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Get 获取节点
func (c *NodeMgr) Get(req types.GetNodeReq) (types.GetNodeResp, error) {
	var (
		resp types.GetNodeResp
		err  error
	)

	servResp, err := c.client.GetNode(c.ctx, &clustermanager.GetNodeRequest{
		InnerIP: req.InnerIP,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.Node, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, &types.Node{
			NodeID:       v.NodeID,
			InnerIP:      v.InnerIP,
			InstanceType: v.InstanceType,
			CPU:          v.CPU,
			Mem:          v.Mem,
			GPU:          v.GPU,
			Status:       v.Status,
			ZoneID:       v.ZoneID,
			NodeGroupID:  v.NodeGroupID,
			ClusterID:    v.ClusterID,
			VPC:          v.VPC,
			Region:       v.Region,
			Passwd:       v.Passwd,
			Zone:         v.Zone,
			DeviceID:     v.DeviceID,
		})
	}

	return resp, nil
}
