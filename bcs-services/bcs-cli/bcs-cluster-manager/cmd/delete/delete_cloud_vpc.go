

package delete

import (
	"context"
	"fmt"

	cloudvpcMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_vpc"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	deleteCloudVPCExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager delete cloudVPC --cloudID xxx --vpcID xxx`))
)

func newDeleteCloudVPCCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudVPC",
		Short:   "delete cloud vpc from bcs-cluster-manager",
		Example: deleteCloudVPCExample,
		Run:     deleteCloudVPC,
	}

	cmd.Flags().StringVarP(&cloudID, "cloudID", "c", "", `cloud ID (required)`)
	cmd.MarkFlagRequired("cloudID")
	cmd.Flags().StringVarP(&vpcID, "vpcID", "", "", `VPC ID (required)`)
	cmd.MarkFlagRequired("vpcID")

	return cmd
}

func deleteCloudVPC(cmd *cobra.Command, args []string) {
	err := cloudvpcMgr.New(context.Background()).Delete(types.DeleteCloudVPCReq{
		CloudID: cloudID,
		VPCID:   vpcID,
	})
	if err != nil {
		klog.Fatalf("delete cloud vpc failed: %v", err)
	}

	fmt.Println("delete cloud vpc succeed")
}
