

package list

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/printer"
	cloudvpcMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_vpc"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listCloudVPCExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager list cloudVPC --networkType overlay`))
)

func newListCloudVPCCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudVPC",
		Short:   "list cloud vpc from bcs-cluster-manager",
		Example: listCloudVPCExample,
		Run:     listCloudVPC,
	}

	cmd.Flags().StringVarP(&networkType, "networkType", "n", "overlay", `cloud VPC network type (required) overlay/underlay`)

	return cmd
}

func listCloudVPC(cmd *cobra.Command, args []string) {
	resp, err := cloudvpcMgr.New(context.Background()).List(types.ListCloudVPCReq{
		NetworkType: networkType,
	})
	if err != nil {
		klog.Fatalf("list cloud vpc failed: %v", err)
	}

	header := []string{"CLOUD_ID", "REGION", "REGION_NAME", "NETWORK_TYPE", "VPC_ID", "VPC_NAME",
		"AVAILABLE", "EXTRA", "CREATOR", "UPDATER", "CREAT_TIME", "UPDATE_TIME"}
	data := make([][]string, len(resp.Data))
	for key, item := range resp.Data {
		data[key] = []string{
			item.CloudID,
			item.Region,
			item.RegionName,
			item.NetworkType,
			item.VPCID,
			item.VPCName,
			item.Available,
			item.Extra,
			item.Creator,
			item.Updater,
			item.CreatTime,
			item.UpdateTime,
		}
	}

	err = printer.PrintList(flagOutput, resp.Data, header, data)
	if err != nil {
		klog.Fatalf("list cloud account to perm failed: %v", err)
	}
}
