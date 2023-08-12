

package cloudaccount

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Update 更新云凭证
func (c *CloudAccountMgr) Update(req types.UpdateCloudAccountReq) error {
	resp, err := c.client.UpdateCloudAccount(c.ctx, &clustermanager.UpdateCloudAccountRequest{
		CloudID:     req.CloudID,
		AccountID:   req.AccountID,
		AccountName: req.AccountName,
		ProjectID:   req.ProjectID,
		Desc:        req.Desc,
		Enable:      wrapperspb.Bool(true),
		Updater:     "bcs",
	})
	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
