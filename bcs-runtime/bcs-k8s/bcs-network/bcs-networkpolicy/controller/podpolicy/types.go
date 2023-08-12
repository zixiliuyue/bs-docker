

package podpolicy

// EventType used to define the event type of informer's event
type EventType string

const (
	// PodUpdate event change of pod
	PodUpdate EventType = "PodUpdate"
	// NamespaceUpdate event change of namespace
	NamespaceUpdate EventType = "NamespaceUpdate"
	// NetworkPolicyUpdate event change of networkPolicy
	NetworkPolicyUpdate EventType = "NetworkPolicyUpdate"
)

// resourceEvent defines the change of informer received
type resourceEvent struct {
	Type      EventType
	Namespace string
	Name      string
}
