

package cloudaccount

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// List 查询云凭证列表
func (c *CloudAccountMgr) List(req types.ListCloudAccountReq) (types.ListCloudAccountResp, error) {
	var (
		resp types.ListCloudAccountResp
		err  error
	)

	servResp, err := c.client.ListCloudAccount(c.ctx, &clustermanager.ListCloudAccountRequest{})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.CloudAccountInfo, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, &types.CloudAccountInfo{
			AccountID:   v.Account.AccountID,
			AccountName: v.Account.AccountName,
			ProjectID:   v.Account.ProjectID,
			Desc:        v.Account.Desc,
			Account: types.Account{
				SecretID:          v.Account.Account.SecretID,
				SecretKey:         v.Account.Account.SecretKey,
				SubscriptionID:    v.Account.Account.SubscriptionID,
				TenantID:          v.Account.Account.TenantID,
				ResourceGroupName: v.Account.Account.ResourceGroupName,
				ClientID:          v.Account.Account.ClientID,
				ClientSecret:      v.Account.Account.ClientSecret,
			},
			Clusters: v.Clusters,
		})
	}

	return resp, nil
}
