

package get

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
	getClustersNamespaceLong = templates.LongDesc(i18n.T(`
		Display one or many project clusters namespaces.
		Prints a table of the most important information about the specified resources.`))

	getClustersNamespaceExample = templates.Examples(i18n.T(`
		# List project all clusters namespaces in ps output format
		kubectl-bcs-project-manager list namespace --cluster-id=clusterID`))
)

func listClustersNamespace() *cobra.Command {
	request := new(pkg.ListNamespacesRequest)
	cmd := &cobra.Command{
		Use:     "namespace --cluster-id=clusterID",
		Aliases: []string{"n"},
		Short:   i18n.T("Display one or many project clusters namespaces."),
		Long:    getClustersNamespaceLong,
		Example: getClustersNamespaceExample,
		Run: func(cmd *cobra.Command, args []string) {
			projectCode := viper.GetString("bcs.project_code")
			if len(projectCode) == 0 {
				klog.Infoln("Project code (English abbreviation), global unique, the length cannot exceed 64 characters")
				return
			}
			request.ProjectCode = projectCode
			resp, err := pkg.NewClientWithConfiguration(context.Background()).ListNamespaces(request)
			if err != nil {
				klog.Infoln("list clusters namespaces failed: %v", err)
				return
			}
			printer.PrintClusterNamespaceInTable(flagOutput, resp)
		},
	}

	cmd.Flags().StringVarP(&request.ClusterID, "cluster-id", "", "",
		"cluster ID, required")

	return cmd
}
