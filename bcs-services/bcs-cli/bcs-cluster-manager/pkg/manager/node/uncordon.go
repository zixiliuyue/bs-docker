

package node

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// UnCordon 节点设置可调度状态
func (c *NodeMgr) UnCordon(req types.UnCordonNodeReq) (types.UnCordonNodeResp, error) {
	var (
		resp types.UnCordonNodeResp
		err  error
	)

	servResp, err := c.client.UnCordonNode(c.ctx, &clustermanager.UnCordonNodeRequest{
		InnerIPs:  req.InnerIPs,
		ClusterID: req.ClusterID,
		Updater:   "bcs",
	})
	if err != nil {
		return resp, err
	}

	resp.Data = make([]string, 0)

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = servResp.Fail

	return resp, nil
}
