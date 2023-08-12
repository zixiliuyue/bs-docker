

package utils

import (
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/server/types"

	"github.com/emicklei/go-restful"
)

// WriteFuncFactory builds WriteXXX shortcut functions
func WriteFuncFactory(statusCode int) func(response *restful.Response, codeName, message string) {
	return func(response *restful.Response, codeName, message string) {
		response.WriteHeaderAndEntity(statusCode, types.ErrorResponse{
			CodeName: codeName,
			Message:  message,
		})
	}
}

// WriteClientError xxx
var WriteClientError = WriteFuncFactory(400)

// WriteUnauthorizedError xxx
var WriteUnauthorizedError = WriteFuncFactory(401)

// WriteForbiddenError xxx
var WriteForbiddenError = WriteFuncFactory(403)

// WriteNotFoundError xxx
var WriteNotFoundError = WriteFuncFactory(404)

// WriteServerError xxx
var WriteServerError = WriteFuncFactory(500)
