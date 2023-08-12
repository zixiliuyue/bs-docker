

package delete

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	deleteLong = templates.LongDesc(i18n.T(`
	delete resources by stdin, resources and names, or by resources and label selector.`))

	deleteExample = templates.Examples(i18n.T(`
	# delete one or many resources
	kubectl-bcs-cluster-manager delete`))

	virtual     bool
	cloudID     string
	accountID   string
	vpcID       string
	clusterID   string
	nodeGroupID string
	taskID      string
	nodes       []string
)

// NewDeleteCmd 创建delete子命令实例
func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   i18n.T("delete resources by file names, stdin, resources and names, or by resources and label selector"),
		Long:    deleteLong,
		Example: deleteExample,
	}

	// delete subcommands
	cmd.AddCommand(newDeleteCloudAccoundCmd())
	cmd.AddCommand(newDeleteCloudVPCCmd())
	cmd.AddCommand(newDeleteClusterCmd())
	cmd.AddCommand(newDeleteNodesFromClusterCmd())
	cmd.AddCommand(newDeleteNodeGroupCmd())
	cmd.AddCommand(newDeleteTaskCmd())

	return cmd
}
