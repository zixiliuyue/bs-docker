

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CloudIPQuotaSpec defines the desired state of CloudIPQuota
type CloudIPQuotaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Cluster clusterid of quota
	Cluster string `json:"cluster"`
	// Limit maximum available quantity
	Limit int64 `json:"limit"`
}

// CloudIPQuotaStatus defines the observed state of CloudIPQuota
type CloudIPQuotaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudIPQuota is the Schema for the cloudipquota API
type CloudIPQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIPQuotaSpec   `json:"spec,omitempty"`
	Status CloudIPQuotaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudIPQuotaList contains a list of CloudIPQuota
type CloudIPQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIPQuota `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudIPQuota{}, &CloudIPQuotaList{})
}
