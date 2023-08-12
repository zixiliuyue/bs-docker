

package delete

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
	deleteClusterExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager delete cluster --clusterID xxx`))
)

func newDeleteClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   "delete cluster from bcs-cluster-manager",
		Example: deleteClusterExample,
		Run:     deleteCluster,
	}

	cmd.Flags().BoolVarP(&virtual, "virtual", "v", false, `delete virtual cluster`)
	cmd.Flags().StringVarP(&clusterID, "clusterID", "c", "", "cluster ID (required)")
	cmd.MarkFlagRequired("clusterID")

	return cmd
}

func deleteCluster(cmd *cobra.Command, args []string) {
	if virtual {
		resp := &types.CreateVirtualClusterResp{}
		url := fmt.Sprintf("/bcsapi/v4/clustermanager/v1/vcluster/%s?operator=bcs&onlyDeleteInfo=false", clusterID)

		err := clusterMgr.NewHttpClient(context.Background()).Delete(url, resp)
		if err != nil {
			klog.Fatalf("delete virtual cluster failed: %v", err)
		}

		if resp != nil && resp.Code != 0 {
			klog.Fatalf("delete virtual cluster failed: %s", resp.Message)
		}

		fmt.Printf("delete virtual cluster succeed: clusterID: %v, taskID: %v\n", resp.Data.ClusterID, resp.Task.TaskID)

		return
	}

	resp, err := clusterMgr.New(context.Background()).Delete(types.DeleteClusterReq{ClusterID: clusterID})
	if err != nil {
		klog.Fatalf("delete cluster failed: %v", err)
	}

	fmt.Printf("create cluster succeed: clusterID: %v, taskID: %v\n", resp.ClusterID, resp.TaskID)
}
