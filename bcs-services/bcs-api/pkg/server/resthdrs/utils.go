

package resthdrs

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/server/resthdrs/utils"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/server/types"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/asaskevich/govalidator"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

// WriteClientError xxx
var WriteClientError = utils.WriteClientError

// WriteForbiddenError xxx
var WriteForbiddenError = utils.WriteForbiddenError

// WriteNotFoundError xxx
var WriteNotFoundError = utils.WriteNotFoundError

// WriteServerError xxx
var WriteServerError = utils.WriteServerError

func init() {
	// Use json tag name instead of the real struct field name
	// Source: https://github.com/go-playground/validator/issues/287
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}
		return name
	})

	validate.RegisterValidation("apiserver_addresses", ValidateAPIServerAddresses)
}

// ValidateAPIServerAddresses validates if given string is a valid apiserver addresses list.
// A valid addresses should be a list of URL concated with ';'
func ValidateAPIServerAddresses(fl validator.FieldLevel) bool {
	s := fl.Field().String()

	if s == "" {
		return false
	}
	for _, addr := range strings.Split(s, ";") {
		if !govalidator.IsURL(addr) {
			return false
		}
	}
	return true
}

// FormatValidationError turn the original validatetion errors into error response, it will only use the FIRST
// errorField to construct the error message.
func FormatValidationError(errList error) *types.ErrorResponse {
	var message string
	for _, err := range errList.(validator.ValidationErrors) {
		if err.Tag() == "required" {
			message = fmt.Sprintf("errcode: %d, ", common.BcsErrApiBadRequest) + fmt.Sprintf(`field '%s' is required`,
				err.Field())
			break
		}
		message = fmt.Sprintf("errcode: %d, ", common.BcsErrApiBadRequest) + fmt.Sprintf(`'%s' failed on the '%s' tag`,
			err.Field(), err.Tag())
	}
	return &types.ErrorResponse{
		CodeName: "VALIDATION_ERROR",
		Message:  message,
	}
}
