

package backend

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
)

// CommitImage xxx
func (b *backend) CommitImage(taskgroup, image, url string) (*types.BcsMessage, error) {

	taskGroup, err := b.store.FetchTaskGroup(taskgroup)
	if err != nil {
		blog.Errorf("CommitImage: fetch taskgroup(%s) error %s", taskgroup, err.Error())
		return nil, err
	}

	if taskGroup.Status != types.TASKGROUP_STATUS_RUNNING {
		blog.Warnf("taskgroup(%s) cannot commit image under status(%s)", taskgroup, taskGroup.Status)
		return nil, fmt.Errorf("taskgroup(%s) cannot commit image under status(%s)", taskgroup, taskGroup.Status)
	}

	msg := &types.Msg_CommitTask{}

	exist := false
	for _, task := range taskGroup.Taskgroup {

		if task.Image == image {
			exist = true
			msg.Tasks = append(msg.Tasks, &types.CommitTask{
				TaskId: &task.ID,
				Image:  &url,
			})
		}
	}

	if exist == false {
		blog.Errorf("image(%s) not found in taskgroup(%s)", image, taskgroup)
		return nil, fmt.Errorf("image(%s) not found in taskgroup(%s)", image, taskgroup)
	}

	bcsMsg := &types.BcsMessage{
		Type:       types.Msg_COMMIT_TASK.Enum(),
		CommitTask: msg,
	}

	return b.sched.SendBcsMessage(taskGroup, bcsMsg)
}
