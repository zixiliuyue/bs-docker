

package main

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/cmd/app"
	"os"
)

func main() {

	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
