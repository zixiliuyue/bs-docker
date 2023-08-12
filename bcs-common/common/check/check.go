

// Package check xxx
package check

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

const (
	prefixFlag  = "BCS_CHECK"
	successFlag = "BCS_MODULE_SUCCESS"
	failureFlag = "BCS_MODULE_FAILURE"
)

var (
	program = filepath.Base(os.Args[0])
)

type logFunc func(string, ...interface{})

func do(logger logFunc, flag, message string) {
	logger("%s | %s | %s", prefixFlag, flag, message)
}

// Succeed print success-message to log file
func Succeed() {
	do(blog.Infof, successFlag, fmt.Sprintf("%s is working.", program))
}

// Fail print failure-message to log file
func Fail(message string) {
	do(blog.Errorf, failureFlag, fmt.Sprintf("%s is not working. %s", program, message))
}
