

package get

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/util"
	nodegroup "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node_group"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	getNodeGroupExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager get nodeGroup --nodeGroupID xxx`))
)

func newGetNodeGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodeGroup",
		Short:   "get node group from bcs-cluster-manager",
		Example: getNodeGroupExample,
		Run:     getNodeGroup,
	}

	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	cmd.MarkFlagRequired("nodeGroupID")

	return cmd
}

func getNodeGroup(cmd *cobra.Command, args []string) {
	resp, err := nodegroup.New(context.Background()).Get(types.GetNodeGroupReq{
		NodeGroupID: nodeGroupID,
	})
	if err != nil {
		klog.Fatalf("get node group failed: %v", err)
	}

	util.Output2Json(resp.Data)
}
