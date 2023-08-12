

package check

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	checkLong = templates.LongDesc(i18n.T(`
	check whether a resource exists from a file or from stdin.`))

	checkExample = templates.Examples(i18n.T(`
	# check a cluster node variable
	kubectl-bcs-cluster-manager check`))

	filename string
	innerIPs []string
)

// NewCheckCmd 创建check子命令实例
func NewCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check",
		Short:   i18n.T("check whether a resource exists from a file or from stdin"),
		Long:    checkLong,
		Example: checkExample,
	}

	cmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "File address Support json file")

	// check subcommands
	cmd.AddCommand(newCheckCloudKubeConfigCmd())
	cmd.AddCommand(newCheckNodeInClusterCmd())

	return cmd
}
