

package types

// ServiceLoadBalance loadbalance between multiple service with weight
type ServiceLoadBalance struct {
	ServiceName string `json:"name"`
	Protocol    string `json:"protocol"`
	Domain      string `json:"domain,omitempty"`
	ExportPort  int    `json:"exportPort,omitempty"`
	Weight      uint   `json:"weight"`
}

// BcsLoadBalance loadbalance for bcs-api
type BcsLoadBalance struct {
	TypeMeta `json:",inline"`
	// AppMeta     `json:",inline"`
	ObjectMeta  `json:"metadata"`
	Protocol    string               `json:"protocol"`
	Port        int                  `json:"port"`
	DomainName  string               `json:"domainName,omitempty"`
	ClusterIP   []string             `json:"clusterIP,omitempty"`
	LoadBalance []ServiceLoadBalance `json:"loadBalance"`
}
