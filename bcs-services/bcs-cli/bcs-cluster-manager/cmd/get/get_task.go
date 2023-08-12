

package get

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/util"
	taskMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/task"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	getTaskExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager get task --taskID xxx`))
)

func newGetTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "task",
		Short:   "get task from bcs-cluster-manager",
		Example: getTaskExample,
		Run:     getTask,
	}

	cmd.Flags().StringVarP(&taskID, "taskID", "t", "", `task ID`)
	cmd.MarkFlagRequired("taskID")

	return cmd
}

func getTask(cmd *cobra.Command, args []string) {
	resp, err := taskMgr.New(context.Background()).Get(types.GetTaskReq{
		TaskID: taskID,
	})
	if err != nil {
		klog.Fatalf("get task failed: %v", err)
	}

	util.Output2Json(resp.Data)
}
