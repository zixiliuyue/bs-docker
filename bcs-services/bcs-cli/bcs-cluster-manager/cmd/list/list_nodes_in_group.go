

package list

import (
	"context"
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/printer"
	nodegroup "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node_group"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listNodesInGroupExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager list nodesInGroup --nodeGroupID xxx`))
)

func newListNodesInGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodesInGroup",
		Short:   "list nodes in group from bcs-cluster-manager",
		Example: listNodesInGroupExample,
		Run:     listNodesInGroup,
	}

	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	cmd.MarkFlagRequired("nodeGroupID")

	return cmd
}

func listNodesInGroup(cmd *cobra.Command, args []string) {
	resp, err := nodegroup.New(context.Background()).ListNodes(types.ListNodesInGroupReq{
		NodeGroupID: nodeGroupID,
	})
	if err != nil {
		klog.Fatalf("list nodes in group failed: %v", err)
	}

	header := []string{"NODE_ID", "INNER_IP", "INSTANCE_TYPE", "CPU", "MEM", "GPU", "STATUS",
		"ZONE_ID", "NODE_GROUP_ID", "CLUSTER_ID", "VPC", "REGION", "DEVICE_ID", "INSTANCE_ROLE"}
	data := make([][]string, len(resp.Data))
	for key, item := range resp.Data {
		data[key] = []string{
			item.NodeID,
			item.InnerIP,
			item.InstanceType,
			fmt.Sprintf("%d", item.CPU),
			fmt.Sprintf("%d", item.Mem),
			fmt.Sprintf("%d", item.GPU),
			item.Status,
			item.ZoneID,
			item.NodeGroupID,
			item.ClusterID,
			item.VPC,
			item.Region,
			item.DeviceID,
			item.InstanceRole,
		}
	}

	err = printer.PrintList(flagOutput, resp.Data, header, data)
	if err != nil {
		klog.Fatalf("list cloud account to perm failed: %v", err)
	}
}
