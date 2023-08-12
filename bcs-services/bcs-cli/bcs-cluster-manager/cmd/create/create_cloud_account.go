

package create

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	cloudAccountMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_account"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	createCloudAccountExample = templates.Examples(i18n.T(`create cloud account from json file. file template: 
	{"cloudID":"tencentCloud","accountName":"test001","desc":"腾讯云测试账号","account":{"secretID":"xxxxxxxxxx",
	"secretKey":"xxxxxxxxxxxx"},"enable":true,"creator":"bcs","projectID":"b363e23b1b354928xxxxxxxxxxxxxxx"}`))
)

func newCreateCloudAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudAccount",
		Short:   i18n.T("create cloud account from bcs-cluster-manager"),
		Example: createCloudAccountExample,
		Run:     createCloudAccount,
	}

	return cmd
}

func createCloudAccount(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		klog.Fatalf("read json file failed: %v", err)
	}

	req := types.CreateCloudAccountReq{}
	err = json.Unmarshal(data, &req)
	if err != nil {
		klog.Fatalf("unmarshal json file failed: %v", err)
	}

	err = cloudAccountMgr.New(context.Background()).Create(req)
	if err != nil {
		klog.Fatalf("create cloud account failed: %v", err)
	}

	fmt.Println("create cloud account succeed")
}
