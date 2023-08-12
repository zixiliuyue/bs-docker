

package task

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Get 查询任务
func (c *TaskMgr) Get(req types.GetTaskReq) (types.GetTaskResp, error) {
	var (
		resp types.GetTaskResp
		err  error
	)

	servResp, err := c.client.GetTask(c.ctx, &clustermanager.GetTaskRequest{
		TaskID: req.TaskID,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	steps := make(map[string]types.Step)
	for k, v := range servResp.Data.Steps {
		steps[k] = types.Step{
			Name:          v.Name,
			System:        v.System,
			Link:          v.Link,
			Params:        v.Params,
			Retry:         v.Retry,
			Start:         v.Start,
			End:           v.End,
			ExecutionTime: v.ExecutionTime,
			Status:        v.Status,
			Message:       v.Message,
			LastUpdate:    v.LastUpdate,
			TaskMethod:    v.TaskMethod,
			TaskName:      v.TaskName,
		}
	}

	resp = types.GetTaskResp{
		Data: types.Task{
			TaskID:         servResp.Data.TaskID,
			TaskType:       servResp.Data.TaskType,
			Status:         servResp.Data.Status,
			Message:        servResp.Data.Message,
			Start:          servResp.Data.Start,
			End:            servResp.Data.End,
			ExecutionTime:  servResp.Data.ExecutionTime,
			CurrentStep:    servResp.Data.CurrentStep,
			StepSequence:   servResp.Data.StepSequence,
			Steps:          steps,
			ClusterID:      servResp.Data.ClusterID,
			ProjectID:      servResp.Data.ProjectID,
			Creator:        servResp.Data.Creator,
			LastUpdate:     servResp.Data.LastUpdate,
			Updater:        servResp.Data.Updater,
			ForceTerminate: servResp.Data.ForceTerminate,
			CommonParams:   servResp.Data.CommonParams,
			TaskName:       servResp.Data.TaskName,
			NodeIPList:     servResp.Data.NodeIPList,
			NodeGroupID:    servResp.Data.NodeGroupID,
		},
	}

	return resp, nil
}
