

package nodegroup

import (
	"errors"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// CleanNodesV2 从节点池中清理回收指定IP节点
func (c *NodeGroupMgr) CleanNodesV2(req types.CleanNodeGroupV2Req) (types.CleanNodeGroupV2Resp, error) {
	var (
		resp types.CleanNodeGroupV2Resp
		err  error
	)
	servResp, err := c.client.CleanNodesInGroupV2(c.ctx, &clustermanager.CleanNodesInGroupV2Request{
		ClusterID:   req.ClusterID,
		Nodes:       strings.Join(req.Nodes, ","),
		NodeGroupID: req.NodeGroupID,
		Operator:    "bcs",
	})

	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.TaskID = servResp.GetData().GetTaskID()

	return resp, nil
}
