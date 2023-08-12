

package move

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	moveLong = templates.LongDesc(i18n.T(`
	move some resource to node group.`))

	moveExample = templates.Examples(i18n.T(`
	# move a cluster variable
	kubectl-bcs-cluster-manager move`))

	clusterID   string
	nodeGroupID string
	nodes       []string
)

// NewMoveCmd 创建move子命令实例
func NewMoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "move",
		Short:   i18n.T("move some resource to node group."),
		Long:    moveLong,
		Example: moveExample,
	}

	// move subcommands
	cmd.AddCommand(newMoveNodesToGroupCmd())

	return cmd
}
