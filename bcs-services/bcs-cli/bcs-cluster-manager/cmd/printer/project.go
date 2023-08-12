

package printer

import (
	"strconv"

	"k8s.io/klog"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// PrintClustersListInTable prints the response that list projects
func PrintClustersListInTable(flagOutput string, resp *clustermanager.ListClusterResp) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Fatalf("list projects output json to stdout failed: %s", err.Error())
		}
	}
	tw := defaultTableWriter()
	tw.SetHeader(func() []string {
		return []string{
			"CLUSTER", "NAME", "FEDERATION_ID", "REGION", "VPC", "PROJECT", "BUSINESS", "ENV",
			"EXCLUSIVE", "TYPE", "CREATOR", "CREATE",
		}
	}())
	tw.SetAutoMergeCells(true)
	for _, item := range resp.Data {
		tw.Append(func() []string {
			return []string{
				item.GetClusterID(), item.GetClusterName(), item.GetRegion(), item.GetVpcID(),
				item.GetProjectID(), item.GetBusinessID(), item.GetEnvironment(),
				strconv.FormatBool(item.GetIsExclusive()), item.GetClusterType(), item.GetCreator(),
				item.GetCreateTime(),
			}
		}())
	}
	tw.Render()
}
