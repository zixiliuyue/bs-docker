

package cmd

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/version"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of clb contoller",
	Long:  `Print the version number of clb contoller`,
	Run: func(cmd *cobra.Command, args []string) {
		version.ShowVersion()
	},
}
