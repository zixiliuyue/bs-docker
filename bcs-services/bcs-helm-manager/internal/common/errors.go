/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"errors"
	"strconv"
)

// HelmManagerError describe the errors and its messages for helm manager
type HelmManagerError uint32

const (
	// ErrHelmManagerSuccess xxx
	ErrHelmManagerSuccess HelmManagerError = iota
	// ErrHelmManagerReqOrRespEmpty xxx
	ErrHelmManagerReqOrRespEmpty
	// ErrHelmManagerRequestParamInvalid xxx
	ErrHelmManagerRequestParamInvalid
	// ErrHelmManagerCreateActionFailed xxx
	ErrHelmManagerCreateActionFailed
	// ErrHelmManagerUpdateActionFailed xxx
	ErrHelmManagerUpdateActionFailed
	// ErrHelmManagerGetActionFailed xxx
	ErrHelmManagerGetActionFailed
	// ErrHelmManagerListActionFailed xxx
	ErrHelmManagerListActionFailed
	// ErrHelmManagerDeleteActionFailed xxx
	ErrHelmManagerDeleteActionFailed
	// ErrHelmManagerInstallActionFailed xxx
	ErrHelmManagerInstallActionFailed
	// ErrHelmManagerUninstallActionFailed xxx
	ErrHelmManagerUninstallActionFailed
	// ErrHelmManagerUpgradeActionFailed xxx
	ErrHelmManagerUpgradeActionFailed
	// ErrHelmManagerRollbackActionFailed xxx
	ErrHelmManagerRollbackActionFailed
	// ErrHelmManagerPreviewActionFailed xxx
	ErrHelmManagerPreviewActionFailed
	// ErrHelmManagerAuthFailed xxx
	ErrHelmManagerAuthFailed
	// ErrHelmManagerRequestComponentFailed xxx
	ErrHelmManagerRequestComponentFailed
	// ErrHelmManagerUploadChartFailed xxx
	ErrHelmManagerUploadChartFailed
	// ErrHelmManagerImportFailed xxx
	ErrHelmManagerImportFailed
	// NoPermissionErr auth faile
	NoPermissionErr = 40403
)

// Int32 return HelmManagerError's code value
func (hme HelmManagerError) Int32() uint32 {
	return uint32(hme)
}

// Error return HelmManagerError's builtin error message
func (hme HelmManagerError) Error() string {
	s, ok := errorCodeMapping[hme]
	if ok {
		return s
	}

	return "unknown error code " + strconv.Itoa(int(hme))
}

// ErrorMessage return HelmManagerError's builtin error message and add the custom message
func (hme HelmManagerError) ErrorMessage(message string) string {
	if hme == ErrHelmManagerSuccess {
		return hme.Error()
	}
	return message
}

// GenError return an error generated by ResourceManagerError
func (hme HelmManagerError) GenError() error {
	return errors.New(hme.Error())
}

// OK check if ResourceManagerError is success
func (hme HelmManagerError) OK() *bool {
	ok := hme == ErrHelmManagerSuccess
	return &ok
}

var errorCodeMapping = map[HelmManagerError]string{
	ErrHelmManagerSuccess:                "success",
	ErrHelmManagerReqOrRespEmpty:         "grpc req or resp is empty",
	ErrHelmManagerRequestParamInvalid:    "request param invalid",
	ErrHelmManagerCreateActionFailed:     "create action failed",
	ErrHelmManagerUpdateActionFailed:     "update action failed",
	ErrHelmManagerGetActionFailed:        "get action failed",
	ErrHelmManagerListActionFailed:       "list action failed",
	ErrHelmManagerDeleteActionFailed:     "delete action failed",
	ErrHelmManagerInstallActionFailed:    "install action failed",
	ErrHelmManagerUninstallActionFailed:  "uninstall action failed",
	ErrHelmManagerUpgradeActionFailed:    "upgrade action failed",
	ErrHelmManagerRollbackActionFailed:   "rollback action failed",
	ErrHelmManagerPreviewActionFailed:    "get preview action failed",
	ErrHelmManagerAuthFailed:             "user auth failed",
	ErrHelmManagerRequestComponentFailed: "request third party failed",
}