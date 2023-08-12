

package disable

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	disableLong = templates.LongDesc(i18n.T(`
	disable auto-scale a deployment, replica set, stateful set, or replication controller.`))

	disableExample = templates.Examples(i18n.T(`
	# disable a cluster variable
	kubectl-bcs-cluster-manager disable`))

	nodeGroupID string
)

// NewDisableCmd 创建disable子命令实例
func NewDisableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "disable",
		Short:   i18n.T("disable auto-scale a deployment, replica set, stateful set, or replication controller"),
		Long:    disableLong,
		Example: disableExample,
	}

	// disable subcommands
	cmd.AddCommand(newDisableNodeGroupAutoScaleCmd())

	return cmd
}
