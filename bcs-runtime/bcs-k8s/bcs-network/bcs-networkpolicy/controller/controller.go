

package controller

import (
	"sync"

	"k8s.io/client-go/tools/cache"
)

const (
	networkPolicyAnnotation = "net.beta.kubernetes.io/network-policy"
	//KubePodFirewallChainPrefix single pod forward chain
	KubePodFirewallChainPrefix = "KUBE-POD-FW-"
	//KubeNetworkPolicyChainPrefix network policy chain prefix
	KubeNetworkPolicyChainPrefix = "KUBE-NWPLCY-"
	//KubeSourceIPSetPrefix source ipset name prefix
	KubeSourceIPSetPrefix = "KUBE-SRC-"
	//KubeDestinationIPSetPrefix destination ipset name prefix
	KubeDestinationIPSetPrefix = "KUBE-DST-"

	// DockerTimeout used to create the docker context, unit second
	DockerTimeout = 5
	// PauseContainerCommand defines the command of pause container
	PauseContainerCommand = "/pause"
	// ContainerNamespaceLabel namespace label for container
	ContainerNamespaceLabel = "namespace"
)

// NetworkPolicyType defines the type of networkPolicy
type NetworkPolicyType string

const (
	// PolicyTypeIngress only have ingress rules
	PolicyTypeIngress NetworkPolicyType = "ingress"
	// PolicyTypeEgress onlye have egress rules
	PolicyTypeEgress NetworkPolicyType = "egress"
	// PolicyTypeBoth have ingress and egress rules
	PolicyTypeBoth NetworkPolicyType = "both"
)

// NetworkPolicyInfo internal structure to represent a network policy
type NetworkPolicyInfo struct {
	Name      string
	Namespace string
	Labels    map[string]string

	// set of pods matching network policy spec podselector label selector
	TargetPods map[string]PodInfo

	// whitelist ingress rules from the network policy spec
	IngressRules []IngressRule

	// whitelist egress rules from the network policy spec
	EgressRules []EgressRule

	// policy type "ingress" or "egress" or "both" as defined by PolicyType in the spec
	PolicyType NetworkPolicyType
}

// PodInfo internal structure to represent Pod
type PodInfo struct {
	IP        string
	Name      string
	Namespace string
	Labels    map[string]string
}

// IngressRule internal structure to represent NetworkPolicyIngressRule in the spec
type IngressRule struct {
	MatchAllPorts  bool
	Ports          []ProtocolAndPort
	NamedPorts     []EndPoints
	MatchAllSource bool
	SrcPods        []PodInfo
	SrcIPBlocks    [][]string
}

// EgressRule internal structure to represent NetworkPolicyEgressRule in the spec
type EgressRule struct {
	MatchAllPorts        bool
	Ports                []ProtocolAndPort
	NamedPorts           []EndPoints
	MatchAllDestinations bool
	DstPods              []PodInfo
	DstIPBlocks          [][]string
}

// ProtocolAndPort internal protocol and port
type ProtocolAndPort struct {
	Protocol string
	Port     string
}

//EndPoints endpoints
type EndPoints struct {
	IPs []string
	ProtocolAndPort
}

// NumericPort2eps internal definition
type NumericPort2eps map[string]*EndPoints

// Protocol2eps internal definition
type Protocol2eps map[string]NumericPort2eps

// NamedPort2eps internal definition
type NamedPort2eps map[string]Protocol2eps

// Controller policy controller interface definition
type Controller interface {
	// SetDataInformerSynced update the dataInformer sync status of networkPolicy controller
	SetDataInformerSynced()

	Run(stopCh <-chan struct{}, wg *sync.WaitGroup) error
	Sync() error
	Cleanup()
	// OnNamespaceUpdate Event callback
	OnNamespaceUpdate(obj interface{})
	OnNetworkPolicyUpdate(obj interface{})
	OnPodUpdate(obj interface{})
	// GetPodEventHandler Get Event Handler for injection
	GetPodEventHandler() cache.ResourceEventHandler
	GetNamespaceEventHandler() cache.ResourceEventHandler
	GetNetworkPolicyEventHandler() cache.ResourceEventHandler
}
