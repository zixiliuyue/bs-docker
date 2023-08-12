

package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// GenerateTransactionID generate transaction id
func GenerateTransactionID(kind string) string {
	return fmt.Sprintf("%s-%d", kind, time.Now().UnixNano())
}

// GetTaskgroupIP get ip of a taskgroup, if status!=running, return ""
func GetTaskgroupIP(t *TaskGroup) string {
	if t.Status != TASKGROUP_STATUS_RUNNING {
		return ""
	}
	bcsInfo := new(BcsContainerInfo)
	for _, oneTask := range t.Taskgroup {
		// process task do not have the statusData upload by executor, because process executor
		// do not have the hostIP and port information. So we make NodeIP, ContainerIP, HostIP directly with AgentIPAddress
		// which is got from offer
		// current running taskgroup kind maybe empty, regard them as APP.
		switch oneTask.Kind {
		case commtypes.BcsDataType_PROCESS:
			return oneTask.AgentIPAddress
		case commtypes.BcsDataType_APP, "":
			if len(oneTask.StatusData) == 0 {
				continue
			}
			if err := json.Unmarshal([]byte(oneTask.StatusData), &bcsInfo); err != nil {
				blog.Warn("task %s StatusData unmarshal err: %s, cannot add to backend", oneTask.ID, err.Error())
				continue
			}

			if bcsInfo.IPAddress != "" {
				return bcsInfo.IPAddress
			}
		}
	}

	return ""
}
