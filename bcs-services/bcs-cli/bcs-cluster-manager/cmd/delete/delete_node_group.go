

package delete

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
	deleteNodeGroupExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager delete nodeGroup --nodeGroupID xxx`))
)

func newDeleteNodeGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodeGroup",
		Short:   "delete node group from bcs-cluster-manager",
		Example: deleteNodeGroupExample,
		Run:     deleteNodeGroup,
	}

	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	cmd.MarkFlagRequired("nodeGroupID")

	return cmd
}

func deleteNodeGroup(cmd *cobra.Command, args []string) {
	resp, err := nodegroup.New(context.Background()).Delete(types.DeleteNodeGroupReq{
		NodeGroupID: nodeGroupID,
	})
	if err != nil {
		klog.Fatalf("delete node group failed: %v", err)
	}

	fmt.Printf("delete node group succeed: taskID: %v\n", resp.TaskID)
}
