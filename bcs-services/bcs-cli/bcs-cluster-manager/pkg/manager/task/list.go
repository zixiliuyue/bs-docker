

package task

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// List 查询任务列表
func (c *TaskMgr) List(req types.ListTaskReq) (types.ListTaskResp, error) {
	var (
		resp types.ListTaskResp
		err  error
	)

	servResp, err := c.client.ListTask(c.ctx, &clustermanager.ListTaskRequest{
		ClusterID: req.ClusterID,
		ProjectID: req.ProjectID,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.Task, 0)

	for _, v := range servResp.Data {
		steps := make(map[string]types.Step)
		for x, y := range v.Steps {
			steps[x] = types.Step{
				Name:          y.Name,
				System:        y.System,
				Link:          y.Link,
				Params:        y.Params,
				Retry:         y.Retry,
				Start:         y.Start,
				End:           y.End,
				ExecutionTime: y.ExecutionTime,
				Status:        y.Status,
				Message:       y.Message,
				LastUpdate:    y.LastUpdate,
				TaskMethod:    y.TaskMethod,
				TaskName:      y.TaskName,
			}
		}

		resp.Data = append(resp.Data, &types.Task{
			TaskID:         v.TaskID,
			TaskType:       v.TaskType,
			Status:         v.Status,
			Message:        v.Message,
			Start:          v.Start,
			End:            v.End,
			ExecutionTime:  v.ExecutionTime,
			CurrentStep:    v.CurrentStep,
			StepSequence:   v.StepSequence,
			Steps:          steps,
			ClusterID:      v.ClusterID,
			ProjectID:      v.ProjectID,
			Creator:        v.Creator,
			LastUpdate:     v.LastUpdate,
			Updater:        v.Updater,
			ForceTerminate: v.ForceTerminate,
			CommonParams:   v.CommonParams,
			TaskName:       v.TaskName,
			NodeIPList:     v.NodeIPList,
			NodeGroupID:    v.NodeGroupID,
		})
	}

	return resp, nil
}
