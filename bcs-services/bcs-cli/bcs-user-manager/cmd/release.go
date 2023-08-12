

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"k8s.io/klog"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-user-manager/cmd/printer"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-user-manager/pkg"
)

// newReleaseCmd create the release tke cidrs command
func newReleaseCmd() *cobra.Command {
	releaseCmd := &cobra.Command{
		Use:   "release",
		Short: "release tkecidrs",
		Long:  "release tkecidrs from bcs-user-manager",
	}
	releaseCmd.AddCommand(releaseTkeCidrCmd())
	return releaseCmd
}

func releaseTkeCidrCmd() *cobra.Command {
	var tkeCidrForm string
	subCmd := &cobra.Command{
		Use:     "tkecidrs",
		Aliases: []string{"tkecidrs"},
		Short:   "release tke cidrs",
		Long:    "release tke cidrs from user manager",
		Example: "kubectl-bcs-user-manager release tkecidrs --tkecidr_form '{\"vpc\":\"\",\"cidr\":\"\",\"cluster\":\"\"}' ",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.OnInitialize(ensureConfig)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			client := pkg.NewClientWithConfiguration(ctx)
			resp, err := client.ReleaseTkeCidr(tkeCidrForm)
			if err != nil {
				klog.Fatalf("release tke cidrs failed: %v", err)
			}
			if resp != nil && resp.Code != 0 {
				klog.Fatalf("release tke cidrs response code not 0 but %d: %s", resp.Code, resp.Message)
			}
			printer.PrintReleaseTkeCidrCmdResult(flagOutput, resp)
		},
	}

	subCmd.Flags().StringVarP(&tkeCidrForm, "tkecidr_form", "f", "",
		"the form used to release tke cidrs")
	return subCmd
}
