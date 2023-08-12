

package create

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-project-manager/cmd/printer"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-project-manager/pkg"
)

var (
	createVariableLong = templates.LongDesc(i18n.T(`
		Create a project variable with the specified project code.`))

	createVariableExample = templates.Examples(i18n.T(`
		# Create a project variable with a command
		kubectl-bcs-project-manager create variable`))
)

func createVariable() *cobra.Command {
	request := new(pkg.CreateVariableRequest)
	cmd := &cobra.Command{
		Use:                   "variable --name=name",
		DisableFlagsInUseLine: true,
		Short:                 i18n.T("Create a project variable with the specified project code"),
		Long:                  createVariableLong,
		Example:               createVariableExample,
		Run: func(cmd *cobra.Command, args []string) {
			projectCode := viper.GetString("bcs.project_code")
			if len(projectCode) == 0 {
				klog.Infoln("Project code (English abbreviation), global unique, the length cannot exceed 64 characters")
				return
			}
			request.ProjectCode = projectCode
			resp, err := pkg.NewClientWithConfiguration(context.Background()).CreateVariable(request)
			if err != nil {
				klog.Infoln(err)
				return
			}
			printer.PrintInJSON(resp)
		},
	}
	cmd.Flags().StringVarP(&request.Name, "name", "n", "",
		"Variable name, length cannot exceed 32 characters, required")
	cmd.Flags().StringVarP(&request.Key, "key", "k", "",
		"The variable key, unique within the project, cannot exceed 64 characters in length, required")
	cmd.Flags().StringVarP(&request.Scope, "scope", "s", "",
		"Variable scope, value range: global, cluster, namespace")
	cmd.Flags().StringVarP(&request.Desc, "desc", "d", "",
		"Variable description and description, limited to 100 characters")
	cmd.Flags().StringVarP(&request.Default, "default", "", "",
		"Variable default value")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("key")

	return cmd
}
