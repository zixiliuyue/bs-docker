

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog"

	"github.com/Tencent/bk-bcs/bcs-common/common/version"
)

const (
	defaultCfgFile = "/etc/bcs/bcs-user-manager.yaml"
)

var (
	cfgFile    string
	flagOutput string
)

func ensureConfig() {
	if cfgFile == "" {
		cfgFile = defaultCfgFile
	}

	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		klog.Fatalf("read config from '%s' failed,err: %s", cfgFile, err.Error())
	}
}

func init() {
	log.SetFlags(0)
	cobra.OnInitialize(ensureConfig)
}

// NewRootCommand returns the rootCmd instance
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "kubectl-bcs-user-manager",
		Short: "kubectl-bcs-user-manager used to operate bcs-user-manager service",
		Long: `
kubectl-bcs-user-manager allows operators to get project info from bcs-user-manager.
`,
	}
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "print the version detail info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.GetVersion())
		},
	}
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newCreateCmd())
	rootCmd.AddCommand(newUpdateCmd())
	rootCmd.AddCommand(newGetCmd())
	rootCmd.AddCommand(newDeleteCmd())
	rootCmd.AddCommand(newVerifyCmd())
	rootCmd.AddCommand(newReleaseCmd())
	rootCmd.AddCommand(newApplyCmd())
	rootCmd.AddCommand(newSyncCmd())
	rootCmd.AddCommand(newGrantCmd())
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile, "config", "c", defaultCfgFile, "config file")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&flagOutput, "output", "o", "wide",
		"optional parameter: json/wide, json will print the json string to stdout")
	return rootCmd
}
