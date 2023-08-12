

package delete

import (
	"context"
	"fmt"

	cloudAccountMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_account"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	deleteCloudAccountExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager delete cloudAccount --cloudID xxx --accountID xxx`))
)

func newDeleteCloudAccoundCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudAccount",
		Short:   "delete cloud account from bcs-cluster-manager",
		Example: deleteCloudAccountExample,
		Run:     deleteCloudAccount,
	}

	cmd.Flags().StringVarP(&cloudID, "cloudID", "c", "", `cloud ID`)
	cmd.MarkFlagRequired("cloudID")
	cmd.Flags().StringVarP(&accountID, "accountID", "a", "", `account ID`)
	cmd.MarkFlagRequired("accountID")

	return cmd
}

func deleteCloudAccount(cmd *cobra.Command, args []string) {
	err := cloudAccountMgr.New(context.Background()).Delete(types.DeleteCloudAccountReq{
		CloudID:   cloudID,
		AccountID: accountID,
	})
	if err != nil {
		klog.Fatalf("delete cloud account failed: %v", err)
	}

	fmt.Println("delete cloud account succeed")
}
