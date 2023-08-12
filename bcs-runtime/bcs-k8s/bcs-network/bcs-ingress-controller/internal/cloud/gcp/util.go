

package gcp

import (
	"net/http"

	"google.golang.org/api/googleapi"
)

// isNotFound returns true if the error is resource not found
func isNotFound(err error) bool {
	if err == nil {
		return false
	}
	ae, ok := err.(*googleapi.Error)
	return ok && ae.Code == http.StatusNotFound
}

func isError(err error) bool {
	if err != nil && !isNotFound(err) {
		return true
	}
	return false
}
