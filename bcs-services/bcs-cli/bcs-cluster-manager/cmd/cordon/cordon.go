

package cordon

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	cordonLong = templates.LongDesc(i18n.T(`
	mark node as unschedulable.`))

	cordonExample = templates.Examples(i18n.T(`
	# mark node "foo" as unschedulable
	kubectl-bcs-cluster-manager cordon`))

	clusterID string
	innerIPs  []string
)

// NewCordonCmd 创建cordon子命令实例
func NewCordonCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cordon",
		Short:   i18n.T("mark node as unschedulable."),
		Long:    cordonLong,
		Example: cordonExample,
	}

	// cordon subcommands
	cmd.AddCommand(newCordonNodeCmd())

	return cmd
}
