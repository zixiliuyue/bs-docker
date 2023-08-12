

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
	deleteClusterNodesExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager delete clusterNodes --clusterID xxx --node xxx.xxx.xxx.xxx --node xxx.xxx.xxx.xxx`))
)

func newDeleteNodesFromClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodesFromCluster",
		Short:   "delete nodes to cluster from bcs-cluster-manager",
		Example: deleteClusterNodesExample,
		Run:     deleteNodesFromCluster,
	}

	cmd.Flags().StringVarP(&clusterID, "clusterID", "c", "", "cluster ID (required)")
	cmd.MarkFlagRequired("clusterID")

	cmd.Flags().StringSliceVarP(&nodes, "node", "n", []string{}, "node ip, for example: -n xxx.xxx.xxx.xxx -n xxx.xxx.xxx.xxx")
	cmd.MarkFlagRequired("node")

	return cmd
}

func deleteNodesFromCluster(cmd *cobra.Command, args []string) {
	resp, err := clusterMgr.New(context.Background()).DeleteNodes(types.DeleteNodesClusterReq{
		ClusterID: clusterID,
		Nodes:     nodes,
	})
	if err != nil {
		klog.Fatalf("delete nodes to cluster failed: %v", err)
	}

	fmt.Printf("delete nodes to cluster succeed: taskID: %v\n", resp.TaskID)
}
