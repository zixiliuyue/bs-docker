

package create

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	nodegroup "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node_group"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	createNodeGroupExample = templates.Examples(i18n.T(`create node group from json file.file template:
	{"name":"test001","autoScaling":{"maxSize":10,"minSize":0,"scalingMode":"CLASSIC_SCALING",
	"multiZoneSubnetPolicy":"PRIORITY","retryPolicy":"IMMEDIATE_RETRY","subnetIDs":["subnet-xxxxxxx"]},
	"enableAutoscale":true,"nodeTemplate":{"unSchedulable":0,"labels":{},"taints":[],"dataDisks":[],
	"dockerGraphPath":"/var/lib/docker","runtime":{"containerRuntime":"docker","runtimeVersion":"19.x"}},
	"launchTemplate":{"imageInfo":{"imageID":"img-xxxxx"},"CPU":2,"Mem":2,"instanceType":"S4.MEDIUM2",
	"systemDisk":{"diskType":"CLOUD_PREMIUM","diskSize":"50"},"internetAccess":{"internetChargeType":
	"TRAFFIC_POSTPAID_BY_HOUR","internetMaxBandwidth":"0","publicIPAssigned":false},"initLoginPassword":"123456",
	"securityGroupIDs":["sg-xxx"],"dataDisks":[],"isSecurityService":true,"isMonitorService":true},
	"clusterID":"BCS-K8S-xxxxx","region":"ap-shanghai","creator":"bcs"}`))
)

func newCreateNodeGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodeGroup",
		Short:   "create node group from bcs-cluster-manager",
		Example: createNodeGroupExample,
		Run:     createNodeGroup,
	}

	return cmd
}

func createNodeGroup(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		klog.Fatalf("read json file failed: %v", err)
	}

	req := types.CreateNodeGroupReq{}

	err = json.Unmarshal(data, &req)
	if err != nil {
		klog.Fatalf("unmarshal json file failed: %v", err)
	}

	resp, err := nodegroup.New(context.Background()).Create(req)
	if err != nil {
		klog.Fatalf("create node group failed: %v", err)
	}

	fmt.Printf("create node group succeed: nodeGroupID: %v, taskID: %v\n", resp.NodeGroupID, resp.TaskID)
}
