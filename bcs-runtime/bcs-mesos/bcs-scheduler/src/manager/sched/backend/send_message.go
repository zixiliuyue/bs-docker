

package backend

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
)

// SendToApplication send message to all tasks of the specified application
func (b *backend) SendToApplication(runAs, appID string, msg *types.BcsMessage) ([]*types.TaskGroupOpResult,
	[]*types.TaskGroupOpResult, error) {
	// var succ, fail []*types.TaskGroup
	var succ, fail []*types.TaskGroupOpResult
	b.store.LockApplication(runAs + "." + appID)
	defer b.store.UnLockApplication(runAs + "." + appID)
	taskGroups, err := b.store.ListTaskGroups(runAs, appID)
	if err != nil {
		blog.Error("send message to application(%s.%s) failed: %s", runAs, appID, err.Error())
		return nil, nil, err
	}

	if len(taskGroups) == 0 {
		blog.Error("send message to application(%s.%s), no taskgroup exist", runAs, appID)
		return nil, nil, fmt.Errorf("empty taskgroups")
	}

	for _, taskGroup := range taskGroups {
		_, err := b.sched.SendBcsMessage(taskGroup, msg)
		if err != nil {
			taskGroupStatus := new(types.TaskGroupOpResult)
			taskGroupStatus.ID = taskGroup.ID
			taskGroupStatus.Status = taskGroup.Status
			taskGroupStatus.Err = err.Error()
			fail = append(fail, taskGroupStatus)
		} else {
			taskGroupStatus := new(types.TaskGroupOpResult)
			taskGroupStatus.ID = taskGroup.ID
			taskGroupStatus.Status = taskGroup.Status
			taskGroupStatus.Err = "no error"
			succ = append(succ, taskGroupStatus)
		}
	}

	return succ, fail, nil
}

// SendToApplicationTaskGroup send message to tasks of the specified taskgroup with taskgroup id
func (b *backend) SendToApplicationTaskGroup(runAs, appID string, taskGroupID string, msg *types.BcsMessage) error {

	b.store.LockApplication(runAs + "." + appID)
	defer b.store.UnLockApplication(runAs + "." + appID)

	taskGroups, err := b.store.ListTaskGroups(runAs, appID)
	if err != nil {
		blog.Error("send message to taskgroup(%s) failed: %s", taskGroupID, err.Error())
		return err
	}

	for _, taskGroup := range taskGroups {
		if taskGroup.ID == taskGroupID {
			_, err := b.sched.SendBcsMessage(taskGroup, msg)
			return err
		}
	}

	return fmt.Errorf("no taskgroup(%s) in application(%s.%s)", taskGroupID, runAs, appID)
}
