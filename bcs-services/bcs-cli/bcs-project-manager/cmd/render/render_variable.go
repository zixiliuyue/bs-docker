

package render

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
	clusterID          string
	namespace          string
	renderVariableLong = templates.LongDesc(i18n.T(`
		render variable's value under a specific cluster, namespace.`))

	renderVariableExample = templates.Examples(i18n.T(`
		# render variable's value
		kubectl-bcs-project-manager render variable`))
)

func renderVariable() *cobra.Command {
	request := new(pkg.RenderVariablesRequest)
	cmd := &cobra.Command{
		Use:                   "variable --cluster-id=clusterID -- [COMMAND] [args...]",
		DisableFlagsInUseLine: true,
		Short:                 i18n.T("Render variable's value under a specific cluster, namespace."),
		Long:                  renderVariableLong,
		Example:               renderVariableExample,
		Run: func(cmd *cobra.Command, args []string) {
			projectCode := viper.GetString("bcs.project_code")
			if len(projectCode) == 0 {
				klog.Infoln("Project code (English abbreviation), global unique, the length cannot exceed 64 characters")
				return
			}
			resp, err := pkg.NewClientWithConfiguration(context.Background()).RenderVariables(request, projectCode, clusterID, namespace)
			if err != nil {
				klog.Infoln("render variable's value failed: %v", err)
				return
			}
			printer.PrintInJSON(resp)
		},
	}
	cmd.Flags().StringVarP(&clusterID, "cluster-id", "", "",
		"ClusterID, required")
	cmd.Flags().StringVarP(&namespace, "namespace", "", "",
		"Namespace name, required")
	cmd.Flags().StringVarP(&request.KeyList, "key", "", "",
		"key. A list of variable key, separated by commas, semicolons or spaces")

	return cmd
}
