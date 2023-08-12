

package update

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
	updateGroupDesiredNodeExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager update groupDesiredNode --nodeGroupID xxx --desiredNode 1`))
)

func newUpdateGroupDesiredNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "groupDesiredNode",
		Short:   "update group desired node from bcs-cluster-manager",
		Example: updateGroupDesiredNodeExample,
		Run:     updateGroupDesiredNode,
	}

	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	cmd.MarkFlagRequired("nodeGroupID")
	cmd.Flags().Uint32VarP(&desiredNode, "desiredNode", "d", 0, "desired node")
	cmd.MarkFlagRequired("desiredNode")

	return cmd
}

func updateGroupDesiredNode(cmd *cobra.Command, args []string) {
	resp, err := nodegroup.New(context.Background()).UpdateDesiredNode(types.UpdateGroupDesiredNodeReq{
		NodeGroupID: nodeGroupID,
		DesiredNode: desiredNode,
	})
	if err != nil {
		klog.Fatalf("update group desired node failed: %v", err)
	}

	fmt.Printf("update group desired node succeed: taskID: %v\n", resp.TaskID)
}
