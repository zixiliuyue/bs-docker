

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BcsDbPrivConfigSpec defines the desired state of BcsDbPrivConfig
type BcsDbPrivConfigSpec struct {
	PodSelector map[string]string `json:"podSelector"`
	AppName     string            `json:"appName"`
	TargetDb    string            `json:"targetDb"`
	DbType      string            `json:"dbType"`
	CallUser    string            `json:"callUser"`
	DbName      string            `json:"dbName"`
	Operator    string            `json:"operator"`
	UseCDP      bool              `json:"useCDP"`
}

// +kubebuilder:object:root=true
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BcsDbPrivConfig is the Schema for the bcsdbprivconfigs API
type BcsDbPrivConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BcsDbPrivConfigSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BcsDbPrivConfigList contains a list of BcsDbPrivConfig
type BcsDbPrivConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BcsDbPrivConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BcsDbPrivConfig{}, &BcsDbPrivConfigList{})
}
