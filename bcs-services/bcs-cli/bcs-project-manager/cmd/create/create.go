

package create

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	createLong = templates.LongDesc(i18n.T(`
		Create a resource from a file or from stdin.`))

	createExample = templates.Examples(i18n.T(`
		# Create a project variable
		kubectl-bcs-project-manager create`))

	filename string
)

// NewCmdCreate 新建命令
func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create -f FILENAME",
		Short:   i18n.T("Create a resource from a file or from stdin"),
		Long:    createLong,
		Example: createExample,
	}

	cmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "File address Support yaml or json file")

	// create subcommands
	cmd.AddCommand(createVariable())
	cmd.AddCommand(createClustersNamespace())

	return cmd
}
