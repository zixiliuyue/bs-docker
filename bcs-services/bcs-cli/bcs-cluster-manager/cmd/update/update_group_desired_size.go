

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
	updateGroupDesiredSizeExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager update groupDesiredNode --nodeGroupID xxx --desiredSize 1`))
)

func newUpdateGroupDesiredSizeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "groupDesiredSize",
		Short:   "update group desired node from bcs-cluster-manager",
		Example: updateGroupDesiredSizeExample,
		Run:     updateGroupDesiredSize,
	}

	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	cmd.MarkFlagRequired("nodeGroupID")
	cmd.Flags().Uint32VarP(&desiredSize, "desiredSize", "d", 0, "desired size")
	cmd.MarkFlagRequired("desiredSize")

	return cmd
}

func updateGroupDesiredSize(cmd *cobra.Command, args []string) {
	err := nodegroup.New(context.Background()).UpdateDesiredSize(types.UpdateGroupDesiredSizeReq{
		NodeGroupID: nodeGroupID,
		DesiredSize: desiredSize,
	})
	if err != nil {
		klog.Fatalf("update group desired size failed: %v", err)
	}

	fmt.Println("update group desired size succeed")
}
