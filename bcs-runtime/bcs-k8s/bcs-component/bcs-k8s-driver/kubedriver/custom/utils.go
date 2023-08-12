

package custom

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
)

const (
	// RequestFailedCode starting code for error response
	RequestFailedCode = 4001
	// RequestSuccessCode code for success
	RequestSuccessCode = 0
)

// CustomHTTPResponse : custom http response
func CustomHTTPResponse(response *restful.Response, httpCode int, result bool, msg string, code int, data interface{}) {
	response.WriteHeaderAndEntity(httpCode, APIResponse{
		Result:  result,
		Message: msg,
		Code:    code,
		Data:    data,
	})
}

// CustomServerErrorResponse : custom http response for inner server error
func CustomServerErrorResponse(response *restful.Response, msg string) {
	CustomHTTPResponse(response, http.StatusInternalServerError, false, msg, RequestFailedCode, nil)
}

// CustomSuccessResponse : custom http response for success
func CustomSuccessResponse(response *restful.Response, msg string, data interface{}) {
	CustomHTTPResponse(response, http.StatusOK, true, msg, RequestSuccessCode, data)
}

// CustomSimpleHTTPResponse : simple http response
func CustomSimpleHTTPResponse(response *restful.Response, httpCode int, resp APIResponse) {
	response.WriteHeaderAndEntity(httpCode, resp)
}
