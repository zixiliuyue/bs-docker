

package delete

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	deleteLong = templates.LongDesc(i18n.T(`
		Delete resources by stdin, resources and names, or by resources and label selector.`))

	deleteExample = templates.Examples(i18n.T(`
		# Delete one or many resources
		kubectl-bcs-project-manager delete`))
)

// NewCmdDelete 新建命令删除
func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete ( variable )",
		Short:   i18n.T("Delete resources by stdin, resources and names, or by resources and label selector"),
		Long:    deleteLong,
		Example: deleteExample,
	}

	// delete subcommands
	cmd.AddCommand(deleteVariable())
	cmd.AddCommand(deleteClustersNamespace())

	return cmd
}
