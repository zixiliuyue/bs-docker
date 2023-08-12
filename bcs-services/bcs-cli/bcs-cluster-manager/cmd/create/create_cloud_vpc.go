

package create

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	cloudvpcMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_vpc"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	createCloudVPCExample = templates.Examples(i18n.T(`create cloud vpc from json file.file template: 
	{"cloudID":"tencentCloud","networkType":"overlay","region":"ap-guangzhou","regionName":"广州",
	"vpcName":"vpc-xxxxxxx-1","vpcID":"vpc-123456789","available":"true","extra":"","creator":"bcs"}`))
)

func newCreateCloudVPCCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudVPC",
		Short:   "create cloud vpc from bcs-cluster-manager",
		Example: createCloudVPCExample,
		Run:     createCloudVPC,
	}

	return cmd
}

func createCloudVPC(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		klog.Fatalf("read json file failed: %v", err)
	}

	req := types.CreateCloudVPCReq{}
	err = json.Unmarshal(data, &req)
	if err != nil {
		klog.Fatalf("unmarshal json file failed: %v", err)
	}

	err = cloudvpcMgr.New(context.Background()).Create(req)
	if err != nil {
		klog.Fatalf("create cloud vpc failed: %v", err)
	}

	fmt.Println("create cloud vpc succeed")
}
