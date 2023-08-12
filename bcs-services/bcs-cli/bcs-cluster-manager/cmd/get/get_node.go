

package get

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/util"
	nodeMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	getNodeExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager get node --innerIP xxx`))
)

func newGetNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "node",
		Short:   "get node from bcs-cluster-manager",
		Example: getNodeExample,
		Run:     getNode,
	}

	cmd.Flags().StringVarP(&innerIP, "innerIP", "i", "", "")
	cmd.MarkFlagRequired("innerIP")

	return cmd
}

func getNode(cmd *cobra.Command, args []string) {
	resp, err := nodeMgr.New(context.Background()).Get(types.GetNodeReq{
		InnerIP: innerIP,
	})
	if err != nil {
		klog.Fatalf("get node failed: %v", err)
	}

	util.Output2Json(resp.Data)
}
