

package enable

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	enableLong = templates.LongDesc(i18n.T(`
	enable auto-scale a deployment, replica set, stateful set, or replication controller.`))

	enableExample = templates.Examples(i18n.T(`
	# enable a node group auto scale
	kubectl-bcs-cluster-manager enable`))

	nodeGroupID string
)

// NewEnableCmd 创建enable子命令实例
func NewEnableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "enable",
		Short:   i18n.T("enable auto-scale a deployment, replica set, stateful set, or replication controller"),
		Long:    enableLong,
		Example: enableExample,
	}

	// enable subcommands
	cmd.AddCommand(newEnableNodeGroupAutoScaleCmd())

	return cmd
}
