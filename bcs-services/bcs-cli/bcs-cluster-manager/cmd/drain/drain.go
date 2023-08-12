

package drain

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	drainLong = templates.LongDesc(i18n.T(`
	drain node in preparation for maintenance.
	The given node will be marked unschedulable to prevent new pods from arriving.`))

	drainExample = templates.Examples(i18n.T(`
	# mark node "foo" as unschedulable
	kubectl-bcs-cluster-manager drain`))

	clusterID string
	innerIPs  []string
)

// NewDrainCmd 创建drain子命令实例
func NewDrainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "drain",
		Short:   i18n.T("drain node in preparation for maintenance."),
		Long:    drainLong,
		Example: drainExample,
	}

	// drain subcommands
	cmd.AddCommand(newDrainNodeCmd())

	return cmd
}
