

package get

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/util"
	clusterMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cluster"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	getClusterExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager get cluster --clusterID xxx`))
)

func newGetClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   "get cluster from bcs-cluster-manager",
		Example: getClusterExample,
		Run:     getCluster,
	}

	cmd.Flags().StringVarP(&clusterID, "clusterID", "c", "", "cluster ID (required)")
	cmd.MarkFlagRequired("clusterID")

	return cmd
}

func getCluster(cmd *cobra.Command, args []string) {
	resp, err := clusterMgr.New(context.Background()).Get(types.GetClusterReq{ClusterID: clusterID})
	if err != nil {
		klog.Fatalf("get cluster failed: %v", err)
	}

	util.Output2Json(resp.Data)
}
