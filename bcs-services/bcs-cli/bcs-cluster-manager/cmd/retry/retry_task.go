

package retry

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
	retryTaskExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager retry task --taskID xxx`))
)

func newRetryTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "task",
		Short:   "retry task from bcs-cluster-manager",
		Example: retryTaskExample,
		Run:     retryTask,
	}

	cmd.Flags().StringVarP(&taskID, "taskID", "t", "", `task ID`)
	cmd.MarkFlagRequired("taskID")

	return cmd
}

func retryTask(cmd *cobra.Command, args []string) {
	resp, err := taskMgr.New(context.Background()).Retry(types.RetryTaskReq{
		TaskID: taskID,
	})
	if err != nil {
		klog.Fatalf("retry task failed: %v", err)
	}

	fmt.Printf("retry task succeed: taskID: %v\n", resp.TaskID)
}
