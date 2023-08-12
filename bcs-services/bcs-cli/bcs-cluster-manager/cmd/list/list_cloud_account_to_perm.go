

package list

import (
	"context"
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/printer"
	cloudAccountMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_account"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listCloudAccountToPermExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager list cloudAccountToPerm`))
)

func newListCloudAccountToPermCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudAccountToPerm",
		Short:   "list cloud account to perm from bcs-cluster-manager",
		Example: listCloudAccountToPermExample,
		Run:     listCloudAccountToPerm,
	}

	return cmd
}

func listCloudAccountToPerm(cmd *cobra.Command, args []string) {
	resp, err := cloudAccountMgr.New(context.Background()).ListToPerm(types.ListCloudAccountToPermReq{})
	if err != nil {
		klog.Fatalf("list cloud account to perm failed: %v", err)
	}

	header := []string{"COUND_ID", "PROJECT_ID", "ACCOUNT_ID", "ACCOUNT_NAME", "ENABLE", "CREATOR", "CREAT_TIME"}
	data := make([][]string, len(resp.Data))
	for key, item := range resp.Data {
		data[key] = []string{
			item.CloudID,
			item.ProjectID,
			item.AccountID,
			item.AccountName,
			fmt.Sprintf("%t", item.Enable),
			item.Creator,
			item.CreatTime,
		}
	}

	err = printer.PrintList(flagOutput, resp.Data, header, data)
	if err != nil {
		klog.Fatalf("list cloud account to perm failed: %v", err)
	}
}
