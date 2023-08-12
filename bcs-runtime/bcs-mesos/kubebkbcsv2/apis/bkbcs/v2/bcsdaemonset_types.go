

package v2

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BcsDaemonsetSpec defines the desired state of BcsDaemonset
type BcsDaemonsetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	types.BcsDaemonset
}

// BcsDaemonsetStatus defines the observed state of BcsDaemonset
type BcsDaemonsetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BcsDaemonset is the Schema for the bcsdaemonsets API
type BcsDaemonset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BcsDaemonsetSpec   `json:"spec,omitempty"`
	Status BcsDaemonsetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BcsDaemonsetList contains a list of BcsDaemonset
type BcsDaemonsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BcsDaemonset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BcsDaemonset{}, &BcsDaemonsetList{})
}
