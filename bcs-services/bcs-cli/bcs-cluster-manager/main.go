

package main

import (
	"k8s.io/klog"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		klog.Fatalf(err.Error())
	}
}
