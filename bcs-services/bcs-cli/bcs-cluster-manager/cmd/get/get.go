

package get

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	getLong = templates.LongDesc(i18n.T(`
	display one resources.
	Prints a table of the most important information about the specified resources.`))

	getExample = templates.Examples(i18n.T(`
	# get a cluster variable
	kubectl-bcs-cluster-manager get`))

	vpcID       string
	clusterID   string
	innerIP     string
	nodeGroupID string
	taskID      string
)

// NewGetCmd 创建get子命令实例
func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   i18n.T("display one resources."),
		Long:    getLong,
		Example: getExample,
	}

	// get subcommands
	cmd.AddCommand(newGetVPCCidrCmd())
	cmd.AddCommand(newGetClusterCmd())
	cmd.AddCommand(newGetNodeCmd())
	cmd.AddCommand(newGetNodeGroupCmd())
	cmd.AddCommand(newGetTaskCmd())

	return cmd
}
