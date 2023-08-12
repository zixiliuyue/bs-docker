

package remove

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	removeLong = templates.LongDesc(i18n.T(`
	remove some resource to node group.`))

	removeExample = templates.Examples(i18n.T(`
	# remove a cluster variable
	kubectl-bcs-cluster-manager remove`))

	clusterID   string
	nodeGroupID string
	nodes       []string
)

// NewRemoveCmd 创建remove子命令实例
func NewRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Short:   i18n.T("remove some resource to node group."),
		Long:    removeLong,
		Example: removeExample,
	}

	// remove subcommands
	cmd.AddCommand(newRemoveNodesFromGroupCmd())

	return cmd
}
