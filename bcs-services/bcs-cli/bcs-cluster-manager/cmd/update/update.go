

package update

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	updateLong = templates.LongDesc(i18n.T(`
	update a resource from a file or from stdin.`))

	updateExample = templates.Examples(i18n.T(`
	# update a cluster variable
	kubectl-bcs-cluster-manager update`))

	desiredNode uint32
	desiredSize uint32
	filename    string
	status      string
	clusterID   string
	nodeGroupID string
	innerIPs    []string
)

// NewUpdateCmd 创建update子命令实例
func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update -f FILENAME",
		Short:   i18n.T("update a resource from a file or from stdin"),
		Long:    updateLong,
		Example: updateExample,
	}

	cmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "File address Support json file")

	// update subcommands
	cmd.AddCommand(newUpdateCloudAccountCmd())
	cmd.AddCommand(newUpdateCloudVPCCmd())
	cmd.AddCommand(newUpdateClusterCmd())
	cmd.AddCommand(newUpdateNodeCmd())
	cmd.AddCommand(newUpdateNodeGroupCmd())
	cmd.AddCommand(newUpdateGroupDesiredNodeCmd())
	cmd.AddCommand(newUpdateGroupDesiredSizeCmd())
	cmd.AddCommand(newUpdateTaskCmd())

	return cmd
}
