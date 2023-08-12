// NOCC:tosa/license(设计如此)
//go:build !linux
// +build !linux



package ipvs

import (
	utilsexec "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-apiserver-proxy/pkg/ipvs/exec"
)

// RequiredIPVSKernelModulesAvailableCheck tests IPVS required kernel modules.
type RequiredIPVSKernelModulesAvailableCheck struct {
	Executor utilsexec.Interface
}

// Name returns label for RequiredIPVSKernelModulesAvailableCheck
func (r RequiredIPVSKernelModulesAvailableCheck) Name() string {
	return "RequiredIPVSKernelModulesAvailable"
}

// Check try to validates IPVS required kernel modules exists or not.
func (r RequiredIPVSKernelModulesAvailableCheck) Check() (warnings, errors []error) {

	return nil, nil
}
