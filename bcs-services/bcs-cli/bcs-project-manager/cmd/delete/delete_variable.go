

package delete

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
	deleteVariableLong = templates.LongDesc(i18n.T(`
		Delete project variable by stdin, project code and variable id.`))

	deleteVariableExample = templates.Examples(i18n.T(`
		# Delete a variable with project code
		kubectl-bcs-project-manager delete variable

		# Delete a variable with project code, and contains multiple variable id
		kubectl-bcs-project-manager delete variable --id=id1,id2...`))
)

func deleteVariable() *cobra.Command {
	request := new(pkg.DeleteVariableDefinitionsRequest)
	cmd := &cobra.Command{
		Use:                   "variable --id=id1,id2...",
		DisableFlagsInUseLine: true,
		Short:                 i18n.T("Delete project variable by stdin, project code and variable id."),
		Long:                  deleteVariableLong,
		Example:               deleteVariableExample,
		Run: func(cmd *cobra.Command, args []string) {
			projectCode := viper.GetString("bcs.project_code")
			if len(projectCode) == 0 {
				klog.Infoln("Project code (English abbreviation), global unique, the length cannot exceed 64 characters")
				return
			}
			resp, err := pkg.NewClientWithConfiguration(context.Background()).DeleteVariableDefinitions(request, projectCode)
			if err != nil {
				klog.Infoln("delete project variable failed: %v", err)
				return
			}
			printer.PrintInJSON(resp)
		},
	}

	cmd.Flags().StringVarP(&request.IdList, "id", "i", "",
		"A list of variable definition id, separated by commas, semicolons or spaces")

	return cmd
}
