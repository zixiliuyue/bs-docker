

package v1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceMonitorSpec defines the desired state of ServiceMonitor
type ServiceMonitorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Endpoints []Endpoint    `json:"endpoints,omitempty"`
	Selector  LabelSelector `json:"selector,omitempty"`
}

type LabelSelector struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

type Endpoint struct {
	Port     string              `json:"port,omitempty"`
	Path     string              `json:"path,omitempty"`
	Interval string              `json:"interval,omitempty"`
	Params   map[string][]string `json:"params,omitempty"`
}

// ServiceMonitorStatus defines the observed state of ServiceMonitor
type ServiceMonitorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceMonitor is the Schema for the servicemonitors API
type ServiceMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceMonitorSpec   `json:"spec,omitempty"`
	Status ServiceMonitorStatus `json:"status,omitempty"`
}

func (s *ServiceMonitor) GetUuid() string {
	return fmt.Sprintf("%s.%s", s.Namespace, s.Name)
}

func (s *ServiceMonitor) GetSelector() labels.Requirements {
	rms := labels.Requirements{}
	for k, v := range s.Spec.Selector.MatchLabels {
		r, _ := labels.NewRequirement(k, selection.Equals, []string{v})
		rms = append(rms, *r)
	}
	return rms
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceMonitorList contains a list of ServiceMonitor
type ServiceMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceMonitor{}, &ServiceMonitorList{})
}
