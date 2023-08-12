

package clean

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	cleanLong = templates.LongDesc(i18n.T(`
	clean some resource from stdin.`))

	cleanExample = templates.Examples(i18n.T(`
	# clean nodes in node group
	kubectl-bcs-cluster-manager clean`))

	clusterID   string
	nodeGroupID string
	nodes       []string
)

// NewCleanCmd 创建clean子命令实例
func NewCleanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clean",
		Short:   i18n.T("clean some resource from stdin"),
		Long:    cleanLong,
		Example: cleanExample,
	}

	// clean subcommands
	cmd.AddCommand(newCleanNodesInGroupCmd())
	cmd.AddCommand(newCleanNodesInGroupV2Cmd())

	return cmd
}
