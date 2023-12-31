

package v2

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AdmissionWebhookConfigurationSpec defines the desired state of AdmissionWebhookConfiguration
type AdmissionWebhookConfigurationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	types.AdmissionWebhookConfiguration
}

// AdmissionWebhookConfigurationStatus defines the observed state of AdmissionWebhookConfiguration
type AdmissionWebhookConfigurationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AdmissionWebhookConfiguration is the Schema for the admissionwebhookconfigurations API
type AdmissionWebhookConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AdmissionWebhookConfigurationSpec   `json:"spec,omitempty"`
	Status AdmissionWebhookConfigurationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AdmissionWebhookConfigurationList contains a list of AdmissionWebhookConfiguration
type AdmissionWebhookConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmissionWebhookConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmissionWebhookConfiguration{}, &AdmissionWebhookConfigurationList{})
}
