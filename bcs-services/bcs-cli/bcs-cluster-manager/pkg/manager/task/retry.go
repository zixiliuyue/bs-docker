

package task

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Retry 针对失败的任务, 进行重试操作
func (c *TaskMgr) Retry(req types.RetryTaskReq) (types.RetryTaskResp, error) {
	var (
		resp types.RetryTaskResp
		err  error
	)

	servResp, err := c.client.RetryTask(c.ctx, &clustermanager.RetryTaskRequest{
		TaskID:  req.TaskID,
		Updater: "bcs",
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.TaskID = servResp.Data.TaskID

	return resp, nil
}
