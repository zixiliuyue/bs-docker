

package utils

import (
	"encoding/json"
	"net/http"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// WriteKubeAPIError writes a standard error response
func WriteKubeAPIError(rw http.ResponseWriter, err *errors.StatusError) {
	payload, _ := json.Marshal(err.ErrStatus)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(int(err.ErrStatus.Code))
	rw.Write(payload)
}

// NewNotFound returns a new error which indicates that the resource of the kind and the name was not found.
func NewNotFound(qualifiedResource schema.GroupResource, name string, message string) *errors.StatusError {
	status := errors.NewNotFound(qualifiedResource, name)
	status.ErrStatus.Message = message
	return status
}

// NewForbidden xxx
func NewForbidden(qualifiedResource schema.GroupResource, name string, err error) *errors.StatusError {
	status := errors.NewForbidden(qualifiedResource, name, err)
	return status
}

// NewInternalError xxx
func NewInternalError(err error) *errors.StatusError {
	status := errors.NewInternalError(err)
	return status
}

// NewInvalid returns an error indicating the item is invalid and cannot be processed.
func NewInvalid(qualifiedKind schema.GroupKind, name string, fieldName string, err error) *errors.StatusError {
	fieldError := field.Error{
		Field:  fieldName,
		Detail: err.Error(),
	}
	errs := field.ErrorList{&fieldError}
	status := errors.NewInvalid(qualifiedKind, name, errs)
	return status
}

// NewUnauthorized returns an error indicating the client is not authorized to perform the requested action.
func NewUnauthorized(reason string) *errors.StatusError {
	return errors.NewUnauthorized(reason)
}
