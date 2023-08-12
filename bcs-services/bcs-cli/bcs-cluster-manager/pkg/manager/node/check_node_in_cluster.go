

package node

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// CheckNodeInCluster 获取node信息是否存在bcs集群
func (c *NodeMgr) CheckNodeInCluster(req types.CheckNodeInClusterReq) (types.CheckNodeInClusterResp, error) {
	var (
		resp types.CheckNodeInClusterResp
		err  error
	)

	servResp, err := c.client.CheckNodeInCluster(c.ctx, &clustermanager.CheckNodesRequest{
		InnerIPs: req.InnerIPs,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make(map[string]types.NodeResult)

	for k, v := range servResp.Data {
		resp.Data[k] = types.NodeResult{
			IsExist:     v.IsExist,
			ClusterID:   v.ClusterID,
			ClusterName: v.ClusterName,
		}
	}

	return resp, nil
}
