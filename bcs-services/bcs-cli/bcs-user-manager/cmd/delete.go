

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"k8s.io/klog"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-user-manager/cmd/printer"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-user-manager/pkg"
)

// newDeleteCmd create delete resource command
func newDeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete resource",
		Long:  "delete resource from bcs-user-manager",
	}
	deleteCmd.AddCommand(revokePermissionCmd())
	deleteCmd.AddCommand(deleteTokenCmd())
	return deleteCmd
}

func revokePermissionCmd() *cobra.Command {
	var reqBody string
	subCmd := &cobra.Command{
		Use:     "permission",
		Aliases: []string{"permission", "ps"},
		Short:   "revoke permission",
		Long:    "revoke permissions from user manager",
		Example: "kubectl-bcs-user-manager delete permission --permission_form '{\n  \"apiVersion\": \"\",\n  \"kind\": \"\",\n  " +
			"\"metadata\": {\n    \"name\": \"\",\n    \"namespace\": \"\",\n    \"creationTimestamp\": \"0001-01-01T00:00:00Z\",\n    " +
			"\"labels\": {\n      \"a\": \"a\"\n    },\n    \"annotations\": {\n      \"a\": \"a\"\n    },\n    \"clusterName\": \"\"\n  },\n  " +
			"\"spec\": {\n    \"permissions\": [\n      {\n        \"user_name\": \"\",\n        \"resource_type\": \"\",\n        \"resource\": \"\",\n        " +
			"\"role\": \"\"\n      }\n    ]\n  }\n}' ",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.OnInitialize(ensureConfig)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			client := pkg.NewClientWithConfiguration(ctx)
			resp, err := client.RevokePermission(reqBody)
			if err != nil {
				klog.Fatalf("revoke permissions failed: %v", err)
			}
			if resp != nil && resp.Code != 0 {
				klog.Fatalf("revoke permissions response code not 0 but %d: %s", resp.Code, resp.Message)
			}
			printer.PrintRevokePermissionCmdResult(flagOutput, resp)
		},
	}
	subCmd.Flags().StringVarP(&reqBody, "permission_form", "f", "",
		"the permissions which will be revoked")

	return subCmd
}

func deleteTokenCmd() *cobra.Command {
	var token string
	subCmd := &cobra.Command{
		Use:     "token",
		Example: "kubectl-bcs-manager delete token -t  [token]",
		Aliases: []string{},
		Short:   "delete token",
		Long:    "delete token from user manager",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.OnInitialize(ensureConfig)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			client := pkg.NewClientWithConfiguration(ctx)
			resp, err := client.DeleteToken(token)
			if err != nil {
				klog.Fatalf("delete token failed: %v", err)
			}
			if resp != nil && resp.Code != 0 {
				klog.Fatalf("delete token response code not 0 but %d: %s", resp.Code, resp.Message)
			}
			printer.PrintDeleteTokenCmdResult(flagOutput, resp)
		},
	}
	subCmd.Flags().StringVarP(&token, "token", "t", "",
		"the token which will be deleted")

	return subCmd
}
