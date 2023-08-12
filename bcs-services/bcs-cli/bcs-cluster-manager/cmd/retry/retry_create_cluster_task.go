

package retry

import (
	"context"
	"fmt"

	clusterMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cluster"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	retryCreateClusterTaskExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager retry createClusterTask --clusterID xxx`))
)

func newRetryCreateClusterTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "createClusterTask",
		Short:   "retry create cluster from bcs-cluster-manager",
		Example: retryCreateClusterTaskExample,
		Run:     retryCreateClusterTask,
	}

	cmd.Flags().StringVarP(&clusterID, "clusterID", "c", "", "cluster ID (required)")
	cmd.MarkFlagRequired("clusterID")

	return cmd
}

func retryCreateClusterTask(cmd *cobra.Command, args []string) {
	resp, err := clusterMgr.New(context.Background()).RetryCreate(types.RetryCreateClusterReq{
		ClusterID: clusterID,
	})
	if err != nil {
		klog.Fatalf("retry create cluster failed: %v", err)
	}

	fmt.Printf("retry create cluster succeed: clusterID: %v, taskID: %v\n", resp.ClusterID, resp.TaskID)
}
