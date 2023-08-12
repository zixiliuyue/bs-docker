

package printer

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"k8s.io/klog/v2"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-user-manager/pkg"
)

// PrintCreateTokenCmdResult prints the response that create token
func PrintCreateTokenCmdResult(flagOutput string, resp *pkg.CreateTokenResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Fatalf("create token output json to stdout failed: %s", err.Error())
		}
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"TOKEN", "JWT", "STATUS", "EXPIRED_AT",
		}
	}())
	tw.SetAutoMergeCells(true)
	data := resp.Data
	tw.Append(func() []string {
		return []string{
			data.Token,
			data.JWT,
			data.Status.Val(),
			data.ExpiredAt.Format(timeFormatter),
		}
	}())
	tw.Render()
}

// PrintCreateTempTokenCmdResult prints the response that create temp token
func PrintCreateTempTokenCmdResult(flagOutput string, resp *pkg.CreateTempTokenResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Fatalf("create temp token output json to stdout failed: %s", err.Error())
		}
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"ID", "USERNAME", "TOKEN", "USER_TYPE", "CREATED_BY", "CREATED_AT", "DELETED_AT", "UPDATED_AT", "EXPIRES_AT",
		}
	}())
	tw.SetAutoMergeCells(true)
	data := resp.Data
	tw.Append(func() []string {
		return []string{
			strconv.Itoa(int(data.ID)),
			data.Username,
			data.Token,
			strconv.Itoa(int(data.UserType)),
			data.CreatedBy,
			data.CreatedAt.Format(timeFormatter),
			data.DeletedAt.Format(timeFormatter),
			data.UpdatedAt.Format(timeFormatter),
			data.ExpiresAt.Format(timeFormatter),
		}
	}())
	tw.Render()
}

// PrintCreateClientTokenCmdResult prints the response that create client token
func PrintCreateClientTokenCmdResult(flagOutput string, resp *pkg.CreateClientTokenResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Fatalf("create client token output json to stdout failed: %s", err.Error())
		}
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"TOKEN", "JWT", "STATUS", "EXPIRED_AT",
		}
	}())
	tw.SetAutoMergeCells(true)
	data := resp.Data
	tw.Append(func() []string {
		return []string{
			data.Token,
			data.JWT,
			data.Status.Val(),
			data.ExpiredAt.Format(timeFormatter),
		}
	}())
	tw.Render()
}

// PrintDeleteTokenCmdResult prints the response that delete token
func PrintDeleteTokenCmdResult(flagOutput string, resp *pkg.DeleteTokenResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Fatalf("delete token output json to stdout failed: %s", err.Error())
		}
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"TOKEN", "JWT", "STATUS", "EXPIRED_AT",
		}
	}())
	tw.Render()
}

// PrintGetTokenInfoCmdResult prints the response that get token
func PrintGetTokenInfoCmdResult(flagOutput string, resp *pkg.GetTokenResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Fatalf("get token output json to stdout failed: %s", err.Error())
		}
	}

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"TOKEN", "JWT", "STATUS", "EXPIRED_AT",
		}
	}())
	// 添加页脚
	tw.SetFooter([]string{"", "Total", strconv.Itoa(len(resp.Data))})
	// 合并相同值的列
	//tw.SetAutoMergeCells(true)
	for _, item := range resp.Data {
		tw.Append(func() []string {
			return []string{
				item.Token,
				item.JWT,
				item.Status.Val(),
				item.ExpiredAt.Format(timeFormatter),
			}
		}())
	}
	tw.Render()
}

// PrintGetTokenByUserAndClusterIDCmdResult prints the response that get token by user and clusterID
// nolint
func PrintGetTokenByUserAndClusterIDCmdResult(flagOutput string, resp *pkg.GetTokenByUserAndClusterIDResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Fatalf("get token by user and clusterID output json to stdout failed: %s", err.Error())
		}
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"USERNAME", "TOKEN", "STATUS", "EXPIRED_AT",
		}
	}())
	tw.SetAutoMergeCells(true)
	data := resp.Data
	tw.Append(func() []string {
		return []string{
			data.UserName,
			data.Token,
			data.Status.Val(),
			data.ExpiredAt.Format(timeFormatter),
		}
	}())
	tw.Render()
}

// PrintUpdateTokenCmdResult prints the response that update token
func PrintUpdateTokenCmdResult(flagOutput string, resp *pkg.UpdateTokenResponse) {
	if flagOutput == outputTypeJSON {
		if err := encodeJSON(resp); err != nil {
			klog.Fatalf("update token output json to stdout failed: %s", err.Error())
		}
	}
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(func() []string {
		return []string{
			"TOKEN", "JWT", "STATUS", "EXPIRED_AT",
		}
	}())
	tw.SetAutoMergeCells(true)
	data := resp.Data
	tw.Append(func() []string {
		return []string{
			data.Token,
			data.JWT,
			data.Status.Val(),
			data.ExpiredAt.Format(timeFormatter),
		}
	}())
	tw.Render()
}
