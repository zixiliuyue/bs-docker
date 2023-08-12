

package imported

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	importLong = templates.LongDesc(i18n.T(`
	import a resource from a file or from stdin.`))

	importExample = templates.Examples(i18n.T(`
	# import a cluster variable
	kubectl-bcs-cluster-manager import`))

	filename string
)

// NewImportCmd 创建import子命令实例
func NewImportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "import -f FILENAME",
		Short:   i18n.T("import a resource from a file or from stdin"),
		Long:    importLong,
		Example: importExample,
	}

	cmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "File address Support json file")

	// import subcommands
	cmd.AddCommand(newImportClusterCmd())

	return cmd
}
