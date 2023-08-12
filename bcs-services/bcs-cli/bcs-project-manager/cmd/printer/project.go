

package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/andreazorzetto/yh/highlight"
	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/pretty"
	"k8s.io/klog/v2"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/proto/bcsproject"
)

// PrintProjectsListInTable prints the response that list projects
func PrintProjectsListInTable(flagOutput string, resp *bcsproject.ListProjectsResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Infoln("list projects output json to stdout failed: %s", err.Error())
			return
		}
		return
	}

	if flagOutput == outputTypeYaml {
		marshal, err := json.Marshal(resp.Data.Results)
		if err != nil {
			klog.Infoln("[outputTypeYaml] deserialize failed: %v", err)
			return
		}
		original, err := yaml.JSONToYAML(marshal)
		if err != nil {
			klog.Infoln("json to yaml failed: %v", err)
			return
		}
		PrintInYaml(string(original))
		return
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"PROJECT_ID", "PROJECT_CODE", "NAME", "BUSINESS_ID", "CREATOR", "UPDATER", "CREATE", "UPDATE",
		}
	}())
	// 添加页脚
	tw.SetFooter([]string{"", "", "", "", "", "", "Total", strconv.Itoa(int(resp.Data.Total))})
	// 合并相同值的列
	//tw.SetAutoMergeCells(true)
	for _, item := range resp.Data.Results {
		tw.Append(func() []string {
			return []string{
				item.GetProjectID(),
				item.GetProjectCode(),
				item.GetName(),
				item.GetBusinessID(),
				item.GetCreator(),
				item.GetUpdater(),
				item.GetCreateTime(),
				item.GetUpdateTime(),
			}
		}())
	}
	tw.Render()
}

// PrintProjectVariablesListInTable prints the response that list projects variables definitions
func PrintProjectVariablesListInTable(flagOutput string, resp *bcsproject.ListVariableDefinitionsResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Infoln("list projects output json to stdout failed: %s", err.Error())
		}
		return
	}
	if flagOutput == outputTypeYaml {
		marshal, err := json.Marshal(resp.Data.Results)
		if err != nil {
			klog.Infoln("[outputTypeYaml] deserialize failed: %v", err)
			return
		}
		original, err := yaml.JSONToYAML(marshal)
		if err != nil {
			klog.Infoln("json to yaml failed: %v", err)
			return
		}
		PrintInYaml(string(original))
		return
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"ID", "KEY", "NAME", "DEFAULT", "DEFAULT_VALUE", "SCOPE", "SCOPE_NAME", "CATEGORY", "CATEGORY_NAME", "CREATOR", "UPDATER", "CREATE", "UPDATE",
		}
	}())
	// 添加页脚
	tw.SetFooter([]string{"", "", "", "", "", "", "", "", "", "", "", "Total", strconv.Itoa(int(resp.Data.Total))})
	// 合并相同值的列
	//tw.SetAutoMergeCells(true)
	for _, item := range resp.Data.Results {
		tw.Append(func() []string {
			return []string{
				item.GetId(),
				item.GetKey(),
				item.GetName(),
				item.GetDefault(),
				item.GetDefaultValue(),
				item.GetScope(),
				item.GetScopeName(),
				item.GetCategory(),
				item.GetCategoryName(),
				item.GetCreator(),
				item.GetUpdater(),
				item.GetCreated(),
				item.GetUpdated(),
			}
		}())
	}
	tw.Render()
}

// PrintClusterNamespaceInTable prints the response that list projects clusters namespace
func PrintClusterNamespaceInTable(flagOutput string, resp *bcsproject.ListNamespacesResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Infoln("list projects output json to stdout failed: %s", err.Error())
		}
		return
	}
	if flagOutput == outputTypeYaml {
		marshal, err := json.Marshal(resp.Data)
		if err != nil {
			klog.Infoln("[outputTypeYaml] deserialize failed: %v", err)
			return
		}
		original, err := yaml.JSONToYAML(marshal)
		if err != nil {
			klog.Infoln("json to yaml failed: %v", err)
			return
		}
		PrintInYaml(string(original))
		return
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"NAME", "ITSM_TICKET_URL", "ITSM_TICKET_TYPE", "CREATE",
		}
	}())

	// 合并相同值的列
	//tw.SetAutoMergeCells(true)
	for _, item := range resp.Data {
		tw.Append(func() []string {
			return []string{
				item.GetName(),
				item.GetItsmTicketURL(),
				item.GetItsmTicketType(),
				item.GetCreateTime(),
			}
		}())
	}
	tw.Render()
}

// PrinListAuthorizedProjectsInTable  prints the response that list authorized projects
func PrinListAuthorizedProjectsInTable(flagOutput string, resp *bcsproject.ListAuthorizedProjResp) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Infoln("list projects output json to stdout failed: %s", err.Error())
			return
		}
		return
	}

	if flagOutput == outputTypeYaml {
		marshal, err := json.Marshal(resp.Data.Results)
		if err != nil {
			klog.Infoln("[outputTypeYaml] deserialize failed: %v", err)
			return
		}
		original, err := yaml.JSONToYAML(marshal)
		if err != nil {
			klog.Infoln("json to yaml failed: %v", err)
			return
		}
		PrintInYaml(string(original))
		return
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"PROJECT_ID", "PROJECT_CODE", "NAME", "BUSINESS_ID", "CREATOR", "UPDATER", "CREATE", "UPDATE",
		}
	}())
	// 添加页脚
	tw.SetFooter([]string{"", "", "", "", "", "", "Total", strconv.Itoa(int(resp.Data.Total))})
	// 合并相同值的列
	//tw.SetAutoMergeCells(true)
	for _, item := range resp.Data.Results {
		tw.Append(func() []string {
			return []string{
				item.GetProjectID(),
				item.GetProjectCode(),
				item.GetName(),
				item.GetBusinessID(),
				item.GetCreator(),
				item.GetUpdater(),
				item.GetCreateTime(),
				item.GetUpdateTime(),
			}
		}())
	}
	tw.Render()
}

// PrintInJSON prints the response
func PrintInJSON(response interface{}) {
	if response == nil {
		return
	}
	var data []byte
	_ = encodeJSONWithIndent(4, response, &data)
	fmt.Println(string(pretty.Color(pretty.Pretty(data), nil)))
}

// PrintInYaml Print yaml
func PrintInYaml(response string) {
	var r io.Reader
	r = strings.NewReader(response)
	h, err := highlight.Highlight(r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(h)
}
