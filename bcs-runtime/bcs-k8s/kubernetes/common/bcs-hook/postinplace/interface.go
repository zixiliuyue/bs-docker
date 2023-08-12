

package postinplace

import (
	hookv1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
)

type PostInplaceHookObjectInterface interface {
	GetPostInplaceHook() *hookv1alpha1.HookStep
}

type PostInplaceHookStatusInterface interface {
	GetPostInplaceHookConditions() []hookv1alpha1.PostInplaceHookCondition
	SetPostInplaceHookConditions([]hookv1alpha1.PostInplaceHookCondition)
}
