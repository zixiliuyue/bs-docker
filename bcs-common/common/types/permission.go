

package types

// BcsPermission xxx
type BcsPermission struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata"`

	Spec BcsPermissionSpec `json:"spec"`
}

// BcsPermissionSpec xxx
type BcsPermissionSpec struct {
	Permissions []Permission `json:"permissions"`
}

// Permission xxx
type Permission struct {
	UserName     string `json:"user_name"`
	ResourceType string `json:"resource_type"`
	Resource     string `json:"resource"`
	Role         string `json:"role"`
}
