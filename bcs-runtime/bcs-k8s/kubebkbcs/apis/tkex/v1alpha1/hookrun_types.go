

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HookRun is the Schema for the hookruns API
type HookRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HookRunSpec   `json:"spec"`
	Status HookRunStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HookRunList contains a list of HookRun
type HookRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HookRun `json:"items"`
}

type HookRunSpec struct {
	Metrics   []Metric   `json:"metrics"`
	Args      []Argument `json:"args,omitempty"`
	Terminate bool       `json:"terminate,omitempty"`
}

type HookRunStatus struct {
	Phase         HookPhase      `json:"phase"`
	Message       string         `json:"message,omitempty"`
	MetricResults []MetricResult `json:"metricResults,omitempty"`
	StartedAt     *metav1.Time   `json:"startedAt,omitempty"`
}

type MetricResult struct {
	Name                  string        `json:"name"`
	Phase                 HookPhase     `json:"phase"`
	Measurements          []Measurement `json:"measurements,omitempty"`
	Message               string        `json:"message,omitempty"`
	Count                 int32         `json:"count,omitempty"`
	Successful            int32         `json:"successful,omitempty"`
	Failed                int32         `json:"failed,omitempty"`
	Inconclusive          int32         `json:"inconclusive,omitempty"`
	Error                 int32         `json:"error,omitempty"`
	ConsecutiveError      int32         `json:"consecutiveError,omitempty"`
	ConsecutiveSuccessful int32         `json:"consecutiveSuccessful,omitempty"`
}

type Measurement struct {
	Phase      HookPhase    `json:"phase"`
	Message    string       `json:"message,omitempty"`
	StartedAt  *metav1.Time `json:"startedAt,omitempty"`
	FinishedAt *metav1.Time `json:"finishedAt,omitempty"`
	ResumeAt   *metav1.Time `json:"resumeAt,omitempty"`
	Value      string       `json:"value,omitempty"`
}

type PreDeleteHookCondition struct {
	PodName   string      `json:"podName"`
	StartTime metav1.Time `json:"startTime"`
	HookPhase HookPhase   `json:"phase"`
}

type PreInplaceHookCondition struct {
	PodName   string      `json:"podName"`
	StartTime metav1.Time `json:"startTime"`
	HookPhase HookPhase   `json:"phase"`
}

type HookStep struct {
	// +kubebuilder:validation:Required
	TemplateName string            `json:"templateName"`
	Args         []HookRunArgument `json:"args,omitempty"`
}

type HookRunArgument struct {
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// +kubebuilder:validation:Required
	Value string `json:"value,omitempty"`
}

type PauseReason string

const (
	PauseReasonCanaryPauseStep PauseReason = "PausedByCanaryPauseStep"
	PauseReasonStepBasedHook   PauseReason = "PausedByStepBasedHook"
)

type PauseCondition struct {
	Reason    PauseReason `json:"reason"`
	StartTime metav1.Time `json:"startTime"`
}

func init() {
	SchemeBuilder.Register(&HookRun{}, &HookRunList{})
}
