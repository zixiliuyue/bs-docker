

package node

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Drain 节点pod迁移,将节点上的 Pod 驱逐
func (c *NodeMgr) Drain(req types.DrainNodeReq) (types.DrainNodeResp, error) {
	var (
		resp types.DrainNodeResp
		err  error
	)

	servResp, err := c.client.DrainNode(c.ctx, &clustermanager.DrainNodeRequest{
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
