

package list

import (
	"context"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/printer"
	cloudAccountMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_account"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	listCloudAccountExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager list cloudAccount`))
)

func newListCloudAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudAccount",
		Short:   "list cloud account from bcs-cluster-manager",
		Example: listCloudAccountExample,
		Run:     listCloudAccount,
	}

	return cmd
}

func listCloudAccount(cmd *cobra.Command, args []string) {
	resp, err := cloudAccountMgr.New(context.Background()).List(types.ListCloudAccountReq{})
	if err != nil {
		klog.Fatalf("list cloud account failed: %v", err)
	}

	header := []string{"ACCOUNT_ID", "ACCOUNT_NAME", "PROJECT_ID", "DESC", "Clusters"}
	data := make([][]string, len(resp.Data))
	for key, item := range resp.Data {
		data[key] = []string{
			item.AccountID,
			item.AccountName,
			item.ProjectID,
			item.Desc,
			strings.Join(item.Clusters, "\n"),
		}
	}

	err = printer.PrintList(flagOutput, resp.Data, header, data)
	if err != nil {
		klog.Fatalf("list cloud account to perm failed: %v", err)
	}
}
