

package main

import (
	"fmt"
	"os"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-container-executor/app"
)

func main() {
	cmd := app.ParseCmdFlags()
	// rnning with
	if err := app.Run(cmd); err != nil {
		fmt.Fprintf(os.Stderr, "bcs-container-executor running failed: %s\n", err.Error())
		os.Exit(1)
	}
}
