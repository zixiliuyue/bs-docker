

package cloudaccount

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Create 创建云凭证
func (c *CloudAccountMgr) Create(req types.CreateCloudAccountReq) error {
	resp, err := c.client.CreateCloudAccount(c.ctx, &clustermanager.CreateCloudAccountRequest{
		CloudID:     req.CloudID,
		AccountName: req.AccountName,
		ProjectID:   req.ProjectID,
		Desc:        req.Desc,
		Account: &clustermanager.Account{
			SecretID:          req.Account.SecretID,
			SecretKey:         req.Account.SecretKey,
			SubscriptionID:    req.Account.SubscriptionID,
			TenantID:          req.Account.TenantID,
			ResourceGroupName: req.Account.ResourceGroupName,
			ClientID:          req.Account.ClientID,
			ClientSecret:      req.Account.ClientSecret,
		},
		Enable:  wrapperspb.Bool(true),
		Creator: "bcs",
	})
	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
