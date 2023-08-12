

package cloudaccount

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Delete 删除云凭证
func (c *CloudAccountMgr) Delete(req types.DeleteCloudAccountReq) error {
	resp, err := c.client.DeleteCloudAccount(c.ctx, &clustermanager.DeleteCloudAccountRequest{
		CloudID:   req.CloudID,
		AccountID: req.AccountID,
	})
	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
