

package render

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	renderLong = templates.LongDesc(i18n.T(`
		render resource a specific cluster, namespace.`))

	renderExample = templates.Examples(i18n.T(`
		# render a resource
		kubectl-bcs-project-manager render`))
)

// NewCmdRender 新建Cmd渲染
func NewCmdRender() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "render",
		Short:   i18n.T("render resource a specific cluster, namespace."),
		Long:    renderLong,
		Example: renderExample,
	}

	// render subcommands
	cmd.AddCommand(renderVariable())

	return cmd
}
