

package add

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	addLong = templates.LongDesc(i18n.T(`add a resource from stdin.`))

	addExample = templates.Examples(i18n.T(`
	# add a node to cluster
	kubectl-bcs-cluster-manager add`))

	clusterID    string
	initPassword string
	nodes        []string
)

// NewAddCmd 创建add子命令实例
func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add",
		Short:   i18n.T("add a resource from stdin"),
		Long:    addLong,
		Example: addExample,
	}

	// add subcommands
	cmd.AddCommand(newAddNodesToClusterCmd())

	return cmd
}
