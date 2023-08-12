

package edit

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	editLong = templates.LongDesc(i18n.T(`
		Edit a resource from the default editor.
		The edit command allows you to directly edit any API resource you can retrieve via the
		command-line tools. It will open the editor 'vi' for Linux or 'notepad' for Windows.`))

	editExample = templates.Examples(i18n.T(`
		# Edit a project or project variable
		kubectl-bcs-project-manager edit project/variable`))
)

// NewCmdEdit 新建命令编辑
func NewCmdEdit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "edit",
		Short:   i18n.T("Edit a resource on the server"),
		Long:    editLong,
		Example: editExample,
	}

	// edit subcommands
	cmd.AddCommand(editProject())
	cmd.AddCommand(editVariable())
	cmd.AddCommand(editClustersNamespace())

	return cmd
}
