

package types

import (
	"github.com/emicklei/go-restful"
)

// EmptyResponse xxx
type EmptyResponse struct{}

// ErrorResponse xxx
type ErrorResponse struct {
	CodeName string `json:"code_name"`
	Message  string `json:"message"`
}

// ErrorRespWithStatus is useful for common helper functions
type ErrorRespWithStatus struct {
	Resp       *ErrorResponse
	StatusCode int
}

// WriteToResp xxx
func (resp *ErrorRespWithStatus) WriteToResp(response *restful.Response) {
	response.WriteHeaderAndEntity(resp.StatusCode, *resp.Resp)
}
