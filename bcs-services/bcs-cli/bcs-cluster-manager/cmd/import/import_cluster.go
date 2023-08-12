

package imported

import (
	"context"
	"encoding/json"
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
	importClusterExample = templates.Examples(i18n.T(`import cluster from json file. file template: 
	{"clusterID":"","projectID":"","businessID":"100001","engineType":"k8s","isExclusive":false,
	"clusterType":"single","clusterName":"ceshi","environment":"stag","provider":"tencentCloud"}`))
)

func newImportClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   "import cluster from bcs-cluster-manager",
		Example: importClusterExample,
		Run:     importCluster,
	}

	return cmd
}

func importCluster(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		klog.Fatalf("read file failed: %v", err)
	}

	req := types.ImportClusterReq{}
	err = json.Unmarshal(data, &req)
	if err != nil {
		klog.Fatalf("unmarshal json file failed: %v", err)
	}

	err = clusterMgr.New(context.Background()).Import(req)
	if err != nil {
		klog.Fatalf("import cluster failed: %v", err)
	}

	fmt.Println("import cluster succeed")
}
