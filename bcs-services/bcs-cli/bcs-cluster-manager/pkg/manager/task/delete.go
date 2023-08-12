

package task

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Delete 删除任务
func (c *TaskMgr) Delete(req types.DeleteTaskReq) error {
	resp, err := c.client.DeleteTask(c.ctx, &clustermanager.DeleteTaskRequest{
		TaskID: req.TaskID,
	})
	if err != nil {
		return err
	}

	if resp != nil && resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}
