

package enable

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
	enableNodeGroupAutoScaleExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager enable nodeGroupAutoScale --nodeGroupID xxx`))
)

func newEnableNodeGroupAutoScaleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodeGroupAutoScale",
		Short:   "enable group auto scale from bcs-cluster-manager",
		Example: enableNodeGroupAutoScaleExample,
		Run:     enableNodeGroupAutoScale,
	}

	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	cmd.MarkFlagRequired("nodeGroupID")

	return cmd
}

func enableNodeGroupAutoScale(cmd *cobra.Command, args []string) {
	err := nodegroup.New(context.Background()).EnableAutoScale(types.EnableNodeGroupAutoScaleReq{
		NodeGroupID: nodeGroupID,
	})
	if err != nil {
		klog.Fatalf("enable group auto scale failed: %v", err)
	}

	fmt.Println("enable group auto scale succeed")
}
