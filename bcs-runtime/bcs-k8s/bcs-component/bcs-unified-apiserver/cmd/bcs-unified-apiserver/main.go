

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-unified-apiserver/cmd/bcs-unified-apiserver/app"
)

func main() {
	ctx := context.Background()
	command := app.NewUnifiedAPIServer(ctx)

	if err := command.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
