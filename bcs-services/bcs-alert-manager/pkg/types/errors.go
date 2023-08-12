

package types

import "github.com/Tencent/bk-bcs/bcs-common/common"

const (
	// BcsErrAlertManagerSuccess success code
	BcsErrAlertManagerSuccess = 0
	// BcsErrAlertManagerSuccessStr success string
	BcsErrAlertManagerSuccessStr = "success"

	// BcsErrAlertManagerInvalidParameter invalid request parameter
	BcsErrAlertManagerInvalidParameter = common.AdditionErrorCode + 500

	// BcsErrAlertManagerAlertClientOperationFailed invalid request parameter
	BcsErrAlertManagerAlertClientOperationFailed = common.AdditionErrorCode + 501

	// BcsErrAlertManagerUnknown unknown error
	BcsErrAlertManagerUnknown = common.AdditionErrorCode + 502
	// BcsErrAlertManagerUnknownStr unknown error msg
	BcsErrAlertManagerUnknownStr = "unknown error"
)
