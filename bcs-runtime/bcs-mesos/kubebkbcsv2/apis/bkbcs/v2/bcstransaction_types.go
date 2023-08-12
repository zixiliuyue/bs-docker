

package v2

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BcsTransactionSpec defines the desired state of BcsTransaction
type BcsTransactionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	types.Transaction
}

// BcsTransactionStatus defines the observed state of BcsTransaction
type BcsTransactionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BcsTransaction is the Schema for the bcstransactions API
type BcsTransaction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BcsTransactionSpec   `json:"spec,omitempty"`
	Status BcsTransactionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BcsTransactionList contains a list of BcsTransaction
type BcsTransactionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BcsTransaction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BcsTransaction{}, &BcsTransactionList{})
}
