

package create

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	createLong = templates.LongDesc(i18n.T(`
	create a resource from a file or from stdin.`))

	createExample = templates.Examples(i18n.T(`
	# create a cluster variable
	kubectl-bcs-cluster-manager create`))

	virtual  bool
	filename string
)

// NewCreateCmd 创建create子命令实例
func NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create -f FILENAME",
		Short:   i18n.T("Create a resource from a file or from stdin"),
		Long:    createLong,
		Example: createExample,
	}

	cmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "File address Support json file")

	// create subcommands
	cmd.AddCommand(newCreateCloudAccountCmd())
	cmd.AddCommand(newCreateCloudVPCCmd())
	cmd.AddCommand(newCreateClusterCmd())
	cmd.AddCommand(newCreateNodeGroupCmd())
	cmd.AddCommand(newCreateTaskCmd())

	return cmd
}
