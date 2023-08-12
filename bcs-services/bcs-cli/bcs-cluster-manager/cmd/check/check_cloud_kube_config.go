

package check

import (
	"context"
	"fmt"
	"os"

	clusterMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cluster"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	checkCloudKubeConfigExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager check cloudKubeConfig --filename xxx`))
)

func newCheckCloudKubeConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudKubeConfig",
		Short:   "check cloud kube config from bcs-cluster-manager",
		Example: checkCloudKubeConfigExample,
		Run:     checkCloudKubeconfig,
	}

	return cmd
}

func checkCloudKubeconfig(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		klog.Fatalf("read file failed: %v", err)
	}

	err = clusterMgr.New(context.Background()).CheckCloudKubeConfig(types.CheckCloudKubeConfigReq{
		Kubeconfig: string(data),
	})
	if err != nil {
		klog.Fatalf("check cloud kube config failed: %v", err)
	}

	fmt.Println("check cloud kube config succeed")
}
