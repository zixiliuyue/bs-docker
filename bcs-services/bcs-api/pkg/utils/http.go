

package utils

import (
	"encoding/json"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/server/types"
	"net/http"
)

// WriteErrorResponse writes a standard error response
func WriteErrorResponse(rw http.ResponseWriter, statusCode int, err *types.ErrorResponse) {
	payload, _ := json.Marshal(err)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	rw.Write(payload)
}
