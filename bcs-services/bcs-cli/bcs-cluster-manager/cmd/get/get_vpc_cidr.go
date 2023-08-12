

package get

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/util"
	cloudvpcMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_vpc"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	getVPCCidrExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager get VPCCidr --vpcID xxx`))
)

func newGetVPCCidrCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "VPCCidr",
		Short:   "get VPC Cidr from bcs-cluster-manager",
		Example: getVPCCidrExample,
		Run:     getVPCCidr,
	}

	cmd.Flags().StringVarP(&vpcID, "vpcID", "v", "", `VPC ID (required)`)
	cmd.MarkFlagRequired("vpcID")

	return cmd
}

func getVPCCidr(cmd *cobra.Command, args []string) {
	resp, err := cloudvpcMgr.New(context.Background()).GetVPCCidr(types.GetVPCCidrReq{
		VPCID: vpcID,
	})
	if err != nil {
		klog.Fatalf("get VPC Cidr failed: %v", err)
	}

	util.Output2Json(resp.Data)
}
