

package predelete

import (
	hookv1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
)

type PreDeleteHookObjectInterface interface {
	GetPreDeleteHook() *hookv1alpha1.HookStep
}

type PreDeleteHookStatusInterface interface {
	GetPreDeleteHookConditions() []hookv1alpha1.PreDeleteHookCondition
	SetPreDeleteHookConditions([]hookv1alpha1.PreDeleteHookCondition)
}
