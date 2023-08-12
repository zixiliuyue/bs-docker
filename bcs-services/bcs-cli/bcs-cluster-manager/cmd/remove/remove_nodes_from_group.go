

package remove

import (
	"context"
	"fmt"

	nodegroup "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node_group"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	removeNodesFromGroupExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager remove nodesFromGroup --clusterID xxx --nodeGroupID xxx --nodes xxx`))
)

func newRemoveNodesFromGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodesFromGroup",
		Short:   "remove nodes to group from bcs-cluster-manager",
		Example: removeNodesFromGroupExample,
		Run:     removeNodesFromGroup,
	}

	cmd.Flags().StringVarP(&clusterID, "clusterID", "c", "", "cluster ID")
	cmd.MarkFlagRequired("clusterID")
	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	cmd.MarkFlagRequired("nodeGroupID")
	cmd.Flags().StringSliceVarP(&nodes, "nodes", "i", []string{}, "node inner ip, for example: -i xxx.xxx.xxx.xxx -i xxx.xxx.xxx.xxx")
	cmd.MarkFlagRequired("nodes")

	return cmd
}

func removeNodesFromGroup(cmd *cobra.Command, args []string) {
	resp, err := nodegroup.New(context.Background()).RemoveNodes(types.RemoveNodesFromGroupReq{
		ClusterID:   clusterID,
		NodeGroupID: nodeGroupID,
		Nodes:       nodes,
	})
	if err != nil {
		klog.Fatalf("remove nodes to group failed: %v", err)
	}

	fmt.Printf("remove nodes to group succeed: taskID: %v\n", resp.TaskID)
}
