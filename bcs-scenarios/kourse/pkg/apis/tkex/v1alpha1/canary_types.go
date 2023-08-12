

// Package v1alpha1 is the v1alpha1 version of the API.
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// CanaryStrategy defines the strategy of canary update
type CanaryStrategy struct {
	// +kubebuilder:validation:Required
	Steps []CanaryStep `json:"steps,omitempty"`
}

// CanaryStep defines the steps of canary update
type CanaryStep struct {
	Partition *intstr.IntOrString `json:"partition,omitempty"`
	Pause     *CanaryPause        `json:"pause,omitempty"`
	Hook      *HookStep           `json:"hook,omitempty"`
}

// CanaryPause defines the pause time of canary update
type CanaryPause struct {
	// Duration the amount of time to wait before moving to the next step.
	// +optional
	Duration *int32 `json:"duration,omitempty"`
}

// CanaryStatus defines the status of canary update
type CanaryStatus struct {
	Revision           string       `json:"revision,omitempty"`
	PauseStartTime     *metav1.Time `json:"pauseStartTime,omitempty"`
	CurrentStepHookRun string       `json:"currentStepHookRun,omitempty"`
}
