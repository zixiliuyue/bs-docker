

package main

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/cmd/webhook/options"
)

func main() {
	op := new(options.GitopsWebhookOptions)
	options.Parse(op)
	blog.InitLogs(op.LogConfig)
	defer blog.CloseLogs()

}
