

package retry

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	retryLong = templates.LongDesc(i18n.T(`
	retry create a resource from stdin.`))

	retryExample = templates.Examples(i18n.T(`
	# retry a cluster variable
	kubectl-bcs-cluster-manager retry`))

	taskID    string
	clusterID string
)

// NewRetryCmd 创建retry子命令实例
func NewRetryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "retry",
		Short:   i18n.T("retry create a resource from stdin."),
		Long:    retryLong,
		Example: retryExample,
	}

	// retry subcommands
	cmd.AddCommand(newRetryTaskCmd())
	cmd.AddCommand(newRetryCreateClusterTaskCmd())

	return cmd
}
