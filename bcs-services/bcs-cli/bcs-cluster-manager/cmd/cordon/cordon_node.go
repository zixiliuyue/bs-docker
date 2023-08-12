

package cordon

import (
	"context"
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/util"
	nodeMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	cordonNodeExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager cordon node --clusterID xxx --innerIPs xxx`))
)

func newCordonNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "node",
		Short:   "cordon node from bcs-cluster-manager",
		Example: cordonNodeExample,
		Run:     cordonNode,
	}

	cmd.Flags().StringSliceVarP(&innerIPs, "innerIPs", "i", []string{}, "node inner ip, for example: -i xxx.xxx.xxx.xxx -i xxx.xxx.xxx.xxx")
	cmd.MarkFlagRequired("innerIPs")

	cmd.Flags().StringVarP(&clusterID, "clusterID", "c", "", "更新节点所属的clusterID")
	cmd.MarkFlagRequired("clusterID")

	return cmd
}

func cordonNode(cmd *cobra.Command, args []string) {
	resp, err := nodeMgr.New(context.Background()).Cordon(types.CordonNodeReq{
		InnerIPs:  innerIPs,
		ClusterID: clusterID,
	})
	if err != nil {
		klog.Fatalf("cordon node failed: %v", err)
	}

	if len(resp.Data) == 0 {
		fmt.Println("cordon node succeed")
		return
	}

	fmt.Println("cordon the following node failed:")
	util.Output2Json(resp.Data)
}
