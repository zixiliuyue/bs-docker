

// Package plugin xxx
package plugin

// InitPluginParameter xxx
type InitPluginParameter struct {
	ConfPath string `json:"confPath"`
}

// HostPluginParameter xxx
type HostPluginParameter struct {
	Ips       []string `json:"ips"`
	ClusterId string   `json:"clusterId"`
}

// HostAttributes xxx
type HostAttributes struct {
	Ip         string       `json:"ip"`
	Attributes []*Attribute `json:"attributes"`
}

// Attribute xxx
type Attribute struct {
	Name   string         `json:"name,omitempty"`
	Type   Value_Type     `json:"type,omitempty"`
	Scalar Value_Scalar   `json:"scalar,omitempty"`
	Ranges []Value_Ranges `json:"ranges,omitempty"`
	Set    Value_Set      `json:"set,omitempty"`
	Text   Value_Text     `json:"text,omitempty"`
}

// Value_Type xxx
type Value_Type uint8

const (
	// ValueScalar xxx
	ValueScalar Value_Type = 0
	// ValueRanges xxx
	ValueRanges Value_Type = 1
	// ValueSet xxx
	ValueSet Value_Type = 2
	// ValueText xxx
	ValueText Value_Type = 3
)

// Value_Scalar xxx
type Value_Scalar struct {
	Value float64 `json:"value,omitempty"`
}

// Value_Ranges xxx
type Value_Ranges struct {
	Begin int `json:"begin,omitempty"`
	End   int `json:"end,omitempty"`
}

// Value_Set xxx
type Value_Set struct {
	Item []string `json:"item,omitempty"`
}

// Value_Text xxx
type Value_Text struct {
	Text string `json:"text,omitempty"`
}

const (
	// SlaveAttributeIpResources xxx
	SlaveAttributeIpResources = "ip-resources"
)
