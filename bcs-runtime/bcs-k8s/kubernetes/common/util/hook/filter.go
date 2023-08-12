

package hook

import (
	hookv1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
)

func FilterHookRuns(hrs []*hookv1alpha1.HookRun, cond func(hr *hookv1alpha1.HookRun) bool) ([]*hookv1alpha1.HookRun, []*hookv1alpha1.HookRun) {
	condTrue := []*hookv1alpha1.HookRun{}
	condFalse := []*hookv1alpha1.HookRun{}
	for _, hr := range hrs {
		if hr == nil {
			continue
		}
		if cond(hr) {
			condTrue = append(condTrue, hr)
		} else {
			condFalse = append(condFalse, hr)
		}
	}
	return condTrue, condFalse
}

func FilterHookRunsByName(hookRuns []*hookv1alpha1.HookRun, name string) *hookv1alpha1.HookRun {
	hookRunsByName, _ := FilterHookRuns(hookRuns, func(hr *hookv1alpha1.HookRun) bool {
		return hr.Name == name
	})
	if len(hookRunsByName) == 1 {
		return hookRunsByName[0]
	}
	return nil
}

func GetCurrentStepHookRun(currentHrs []*hookv1alpha1.HookRun) *hookv1alpha1.HookRun {
	for _, hr := range currentHrs {
		hookRunType, ok := hr.Labels[HookRunTypeLabel]
		if ok && hookRunType == HookRunTypeCanaryStepLabel {
			return hr
		}
	}
	return nil
}

func FilterHookRunsToDelete(hrs []*hookv1alpha1.HookRun, revision string) []*hookv1alpha1.HookRun {
	hrsToDelete := []*hookv1alpha1.HookRun{}
	for _, hr := range hrs {
		if hr.Labels[WorkloadRevisionUniqueLabel] != revision {
			hrsToDelete = append(hrsToDelete, hr)
		}
	}

	return hrsToDelete
}

func FilterPreDeleteHookRuns(hrs []*hookv1alpha1.HookRun) []*hookv1alpha1.HookRun {
	preDeleteHookRuns := []*hookv1alpha1.HookRun{}
	for _, hr := range hrs {
		hookRunType, ok := hr.Labels[HookRunTypeLabel]
		if ok && hookRunType == HookRunTypePreDeleteLabel {
			preDeleteHookRuns = append(preDeleteHookRuns, hr)
		}
	}

	return preDeleteHookRuns
}

func FilterPreInplaceHookRuns(hrs []*hookv1alpha1.HookRun) []*hookv1alpha1.HookRun {
	preInplaceHookRuns := []*hookv1alpha1.HookRun{}
	for _, hr := range hrs {
		hookRunType, ok := hr.Labels[HookRunTypeLabel]
		if ok && hookRunType == HookRunTypePreInplaceLabel {
			preInplaceHookRuns = append(preInplaceHookRuns, hr)
		}
	}

	return preInplaceHookRuns
}
