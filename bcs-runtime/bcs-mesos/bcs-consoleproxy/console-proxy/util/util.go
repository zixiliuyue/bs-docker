

// Package util xxx
package util

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
)

// CreateResponeData create response data
func CreateResponeData(err error, code int, msg string, data interface{}) []byte {
	var rpyErr error
	if err != nil {
		rpyErr = bhttp.InternalError(code, msg)
	} else {
		rpyErr = fmt.Errorf(bhttp.GetRespone(0, "successful ", data))
	}

	blog.V(3).Infof("createRespone: %s", rpyErr.Error())

	return []byte(rpyErr.Error())
}
