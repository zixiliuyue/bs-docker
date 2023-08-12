

package uncordon

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	uncordonLong = templates.LongDesc(i18n.T(`
	mark node as schedulable.`))

	uncordonExample = templates.Examples(i18n.T(`
	# mark node "foo" as schedulable
	kubectl-bcs-cluster-manager uncordon`))

	clusterID string
	innerIPs  []string
)

// NewUncordonCmd 创建uncordon子命令实例
func NewUncordonCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "uncordon",
		Short:   i18n.T("mark node as schedulable."),
		Long:    uncordonLong,
		Example: uncordonExample,
	}

	// uncordon subcommands
	cmd.AddCommand(newUnCordonNodeCmd())

	return cmd
}
