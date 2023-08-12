

package context

import (
	hookv1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
)

type CanaryContext interface {
	CurrentHookRuns() []*hookv1alpha1.HookRun
	OtherHookRuns() []*hookv1alpha1.HookRun
	SetCurrentHookRuns(hrs []*hookv1alpha1.HookRun)
	AddPauseCondition(reason hookv1alpha1.PauseReason)
	HasAddPause() bool
}
