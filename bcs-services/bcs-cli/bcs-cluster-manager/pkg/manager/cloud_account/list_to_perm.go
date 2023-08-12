

package cloudaccount

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// ListToPerm 查询云凭证列表, 主要用于权限资源查询
func (c *CloudAccountMgr) ListToPerm(req types.ListCloudAccountToPermReq) (types.ListCloudAccountToPermResp, error) {
	var (
		resp types.ListCloudAccountToPermResp
		err  error
	)

	servResp, err := c.client.ListCloudAccountToPerm(c.ctx, &clustermanager.ListCloudAccountPermRequest{})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.CloudAccount, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, &types.CloudAccount{
			CloudID:     v.CloudID,
			ProjectID:   v.ProjectID,
			AccountID:   v.AccountID,
			AccountName: v.AccountName,
			Account: types.Account{
				SecretID:          v.Account.SecretID,
				SecretKey:         v.Account.SecretKey,
				SubscriptionID:    v.Account.SubscriptionID,
				TenantID:          v.Account.TenantID,
				ResourceGroupName: v.Account.ResourceGroupName,
				ClientID:          v.Account.ClientID,
				ClientSecret:      v.Account.ClientSecret,
			},
			Enable:    v.Enable,
			Creator:   v.Creator,
			CreatTime: v.CreatTime,
		})
	}

	return resp, nil
}
