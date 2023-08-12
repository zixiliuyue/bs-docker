

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"k8s.io/klog"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-user-manager/cmd/printer"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-user-manager/pkg"
)

// newApplyCmd create new apply command
func newApplyCmd() *cobra.Command {
	applyCmd := &cobra.Command{
		Use:   "apply",
		Short: "apply tkecidrs",
		Long:  "apply tkecidrs from bcs-user-manager",
	}
	applyCmd.AddCommand(applyTkeCidrCmd())
	return applyCmd
}

func applyTkeCidrCmd() *cobra.Command {
	var tkeCidrForm string
	subCmd := &cobra.Command{
		Use:     "tkecidrs",
		Aliases: []string{"tkecidrs"},
		Short:   "apply tke cidrs",
		Long:    "apply tke cidrs from user manager",
		Example: "kubectl-bcs-user-manager apply tkecidrs --tkecidr_form '{\"vpc\":\"\",\"cluster\":\"\", \"ip_number\":}' ",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.OnInitialize(ensureConfig)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			client := pkg.NewClientWithConfiguration(ctx)
			resp, err := client.ApplyTkeCidr(tkeCidrForm)
			if err != nil {
				klog.Fatalf("apply tke cidrs failed: %v", err)
			}
			if resp != nil && resp.Code != 0 {
				klog.Fatalf("apply tke cidrs response code not 0 but %d: %s", resp.Code, resp.Message)
			}
			printer.PrintApplyTkeCidrCmdResult(flagOutput, resp)
		},
	}

	subCmd.Flags().StringVarP(&tkeCidrForm, "tkecidr_form", "f", "",
		"the form json used to apply tke cidrs")
	return subCmd
}
