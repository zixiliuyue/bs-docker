

package delete

import (
	"context"
	"fmt"

	taskMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/task"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	deleteTaskExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager delete task --taskID xxx`))
)

func newDeleteTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "task",
		Short:   "delete task from bcs-cluster-manager",
		Example: deleteTaskExample,
		Run:     deleteTask,
	}

	cmd.Flags().StringVarP(&taskID, "taskID", "t", "", `task ID`)
	cmd.MarkFlagRequired("taskID")

	return cmd
}

func deleteTask(cmd *cobra.Command, args []string) {
	err := taskMgr.New(context.Background()).Delete(types.DeleteTaskReq{
		TaskID: taskID,
	})
	if err != nil {
		klog.Fatalf("delete task failed: %v", err)
	}

	fmt.Println("delete task succeed")
}
