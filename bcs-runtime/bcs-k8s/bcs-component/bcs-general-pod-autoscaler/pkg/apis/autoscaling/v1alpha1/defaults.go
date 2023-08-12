

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_GeneralPodAutoscaler sets default values
func SetDefaults_GeneralPodAutoscaler(obj *GeneralPodAutoscaler) {
	if obj.Spec.Behavior == nil {
		obj.Spec.Behavior = new(GeneralPodAutoscalerBehavior)
	}

	// set defaults for scale down behavior
	if obj.Spec.Behavior.ScaleDown == nil {
		obj.Spec.Behavior.ScaleDown = new(GPAScalingRules)
	}
	if obj.Spec.Behavior.ScaleDown.StabilizationWindowSeconds == nil {
		obj.Spec.Behavior.ScaleDown.StabilizationWindowSeconds = new(int32)
		*obj.Spec.Behavior.ScaleDown.StabilizationWindowSeconds = 300
	}
	if obj.Spec.Behavior.ScaleDown.SelectPolicy == nil {
		obj.Spec.Behavior.ScaleDown.SelectPolicy = new(ScalingPolicySelect)
		*obj.Spec.Behavior.ScaleDown.SelectPolicy = MaxPolicySelect

	}
	if len(obj.Spec.Behavior.ScaleDown.Policies) == 0 {
		obj.Spec.Behavior.ScaleDown.Policies = append(obj.Spec.Behavior.ScaleDown.Policies,
			GPAScalingPolicy{
				Type:          PercentScalingPolicy,
				Value:         100,
				PeriodSeconds: 15,
			})
	}

	// set defaults for scale up behavior
	if obj.Spec.Behavior.ScaleUp == nil {
		obj.Spec.Behavior.ScaleUp = new(GPAScalingRules)
	}
	if obj.Spec.Behavior.ScaleUp.StabilizationWindowSeconds == nil {
		obj.Spec.Behavior.ScaleUp.StabilizationWindowSeconds = new(int32)
		*obj.Spec.Behavior.ScaleUp.StabilizationWindowSeconds = 0
	}
	if obj.Spec.Behavior.ScaleUp.SelectPolicy == nil {
		obj.Spec.Behavior.ScaleUp.SelectPolicy = new(ScalingPolicySelect)
		*obj.Spec.Behavior.ScaleUp.SelectPolicy = MaxPolicySelect

	}
	if len(obj.Spec.Behavior.ScaleUp.Policies) == 0 {
		obj.Spec.Behavior.ScaleUp.Policies = append(obj.Spec.Behavior.ScaleUp.Policies,
			GPAScalingPolicy{
				Type:          PercentScalingPolicy,
				Value:         100,
				PeriodSeconds: 15,
			})
		obj.Spec.Behavior.ScaleUp.Policies = append(obj.Spec.Behavior.ScaleUp.Policies,
			GPAScalingPolicy{
				Type:          PodsScalingPolicy,
				Value:         4,
				PeriodSeconds: 15,
			})
	}
}
