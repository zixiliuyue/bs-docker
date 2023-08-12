

package nodegroup

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// ListNodes 查询节点池节点列表
func (c *NodeGroupMgr) ListNodes(req types.ListNodesInGroupReq) (types.ListNodesInGroupResp, error) {
	var (
		resp types.ListNodesInGroupResp
		err  error
	)

	servResp, err := c.client.ListNodesInGroup(c.ctx, &clustermanager.ListNodesInGroupRequest{
		NodeGroupID: req.NodeGroupID,
	})

	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]types.NodeGroupNode, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, types.NodeGroupNode{
			NodeID:        v.NodeID,
			InnerIP:       v.InnerIP,
			InstanceType:  v.InstanceType,
			CPU:           v.CPU,
			Mem:           v.Mem,
			GPU:           v.CPU,
			Status:        v.Status,
			ZoneID:        v.ZoneID,
			NodeGroupID:   v.NodeGroupID,
			ClusterID:     v.ClusterID,
			VPC:           v.VPC,
			Region:        v.Region,
			Passwd:        v.Passwd,
			Zone:          v.Zone,
			DeviceID:      v.DeviceID,
			InstanceRole:  v.InstanceType,
			UnSchedulable: v.UnSchedulable,
		})
	}

	return resp, nil
}
