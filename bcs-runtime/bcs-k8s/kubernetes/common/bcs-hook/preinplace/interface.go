

package preinplace

import (
	hookv1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
)

type PreInplaceHookObjectInterface interface {
	GetPreInplaceHook() *hookv1alpha1.HookStep
}

type PreInplaceHookStatusInterface interface {
	GetPreInplaceHookConditions() []hookv1alpha1.PreInplaceHookCondition
	SetPreInplaceHookConditions([]hookv1alpha1.PreInplaceHookCondition)
}
