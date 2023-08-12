

package check

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
	checkNodeInClusterExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager check nodeInCluster --innerIPs xxx`))
)

func newCheckNodeInClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodeInCluster",
		Short:   "check node in cluster from bcs-cluster-manager",
		Example: checkNodeInClusterExample,
		Run:     checkNodeInCluster,
	}

	cmd.Flags().StringSliceVarP(&innerIPs, "innerIPs", "i", []string{}, "node inner ip, for example: -i xxx.xxx.xxx.xxx -i xxx.xxx.xxx.xxx")
	cmd.MarkFlagRequired("innerIPs")

	return cmd
}

func checkNodeInCluster(cmd *cobra.Command, args []string) {
	resp, err := nodeMgr.New(context.Background()).CheckNodeInCluster(types.CheckNodeInClusterReq{
		InnerIPs: innerIPs,
	})
	if err != nil {
		klog.Fatalf("check node in cluster failed: %v", err)
	}

	util.Output2Json(resp.Data)
}
