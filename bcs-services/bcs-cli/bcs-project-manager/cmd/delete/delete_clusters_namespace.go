

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
	deleteClustersNamespaceLong = templates.LongDesc(i18n.T(`
		Delete project namespace by project id cluster id and name.`))

	deleteClustersNamespaceExample = templates.Examples(i18n.T(`
		# Delete a variable with project code cluster id and name
		kubectl-bcs-project-manager delete namespace --cluster-id=clusterID --name=name`))
)

func deleteClustersNamespace() *cobra.Command {
	request := new(pkg.DeleteNamespaceRequest)
	cmd := &cobra.Command{
		Use:                   "namespace --cluster-id=clusterID --name=name",
		DisableFlagsInUseLine: true,
		Aliases:               []string{"n"},
		Short:                 i18n.T("Delete project namespace by project id cluster id and name."),
		Long:                  deleteClustersNamespaceLong,
		Example:               deleteClustersNamespaceExample,
		Run: func(cmd *cobra.Command, args []string) {
			projectCode := viper.GetString("bcs.project_code")
			if len(projectCode) == 0 {
				klog.Infoln("Project code (English abbreviation), global unique, the length cannot exceed 64 characters")
				return
			}
			request.ProjectCode = projectCode
			resp, err := pkg.NewClientWithConfiguration(context.Background()).DeleteNamespace(request)
			if err != nil {
				klog.Infoln("delete project namespace failed: %v", err)
				return
			}
			printer.PrintInJSON(resp)
		},
	}

	cmd.Flags().StringVarP(&request.ClusterID, "cluster-id", "", "",
		"Cluster ID, required")
	cmd.Flags().StringVarP(&request.Name, "name", "", "",
		"namespace name, required")

	return cmd
}
