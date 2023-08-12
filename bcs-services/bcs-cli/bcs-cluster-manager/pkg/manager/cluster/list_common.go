

package cluster

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// ListCommon 查询公共集群及公共集群所属权限
func (c *ClusterMgr) ListCommon() (types.ListCommonClusterResp, error) {
	var (
		resp types.ListCommonClusterResp
		err  error
	)

	servResp, err := c.client.ListCommonCluster(c.ctx, &clustermanager.ListCommonClusterReq{})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.Cluster, 0)

	for _, v := range servResp.Data {
		master := make([]string, 0)
		for _, y := range v.Master {
			master = append(master, y.InnerIP)
		}

		resp.Data = append(resp.Data, &types.Cluster{
			ClusterID:   v.ClusterID,
			ProjectID:   v.ProjectID,
			BusinessID:  v.BusinessID,
			EngineType:  v.BusinessID,
			IsExclusive: v.IsExclusive,
			ClusterType: v.ClusterType,
			Creator:     v.Creator,
			Updater:     v.Updater,
			ManageType:  v.ManageType,
			ClusterName: v.ClusterName,
			Environment: v.Environment,
			Provider:    v.Provider,
			Description: v.Description,
			ClusterBasicSettings: types.ClusterBasicSettings{
				Version: v.ClusterBasicSettings.Version,
			},
			NetworkType: v.NetworkType,
			Region:      v.Region,
			VpcID:       v.VpcID,
			NetworkSettings: types.NetworkSettings{
				CidrStep:      v.NetworkSettings.CidrStep,
				MaxNodePodNum: v.NetworkSettings.MaxNodePodNum,
				MaxServiceNum: v.NetworkSettings.MaxServiceNum,
			},
			Master: master,
		})
	}

	return resp, nil
}
