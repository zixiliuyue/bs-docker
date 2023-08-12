

package list

import (
	"context"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/printer"
	clusterMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cluster"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listClusterExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager list commonCluster --offset 0 --limit 10`))
)

func newListClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   "list cluster from bcs-cluster-manager",
		Example: listClusterExample,
		Run:     listCluster,
	}

	cmd.Flags().Uint32VarP(&offset, "offset", "s", 0, "offset number of queries")
	cmd.Flags().Uint32VarP(&limit, "limit", "l", 10, "limit number of queries")

	return cmd
}

func listCluster(cmd *cobra.Command, args []string) {
	resp, err := clusterMgr.New(context.Background()).List(types.ListClusterReq{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		klog.Fatalf("list cluster failed: %v", err)
	}

	header := []string{"CLUSTER_ID", "CLUSTER_NAME", "PROJECT_ID", "BUSINESS_ID", "ENGINE_TYPE", "CLUSTER_TYPE",
		"MANAGE_TYPE", "PROVIDER", "NETWORK_TYPE", "REGION", "VPC_ID", "MASTER", "CREATOR", "UPDATER", "DESCRIPTION"}
	data := make([][]string, len(resp.Data))
	for key, item := range resp.Data {
		data[key] = []string{
			item.ClusterID,
			item.ClusterName,
			item.ProjectID,
			item.BusinessID,
			item.EngineType,
			item.ClusterType,
			item.ManageType,
			item.Provider,
			item.NetworkType,
			item.Region,
			item.VpcID,
			strings.Join(item.Master, "\n"),
			item.Creator,
			item.Updater,
			item.Description,
		}
	}

	err = printer.PrintList(flagOutput, resp.Data, header, data)
	if err != nil {
		klog.Fatalf("list cloud account to perm failed: %v", err)
	}
}
