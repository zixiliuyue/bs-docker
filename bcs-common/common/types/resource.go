

package types

// BcsClusterResource xxx
type BcsClusterResource struct {
	DiskTotal float64               `json:"disktotal"`
	MemTotal  float64               `json:"memtotal"`
	CpuTotal  float64               `json:"cputotal"`
	DiskUsed  float64               `json:"diskused"`
	MemUsed   float64               `json:"memused"`
	CpuUsed   float64               `json:"cpuused"`
	Agents    []BcsClusterAgentInfo `json:"agents"`
}

// BcsClusterAgentInfo xxx
type BcsClusterAgentInfo struct {
	HostName  string  `json:"hostname"`
	IP        string  `json:"ip"`
	DiskTotal float64 `json:"disktotal"`
	MemTotal  float64 `json:"memtotal"`
	CpuTotal  float64 `json:"cputotal"`
	DiskUsed  float64 `json:"diskused"`
	MemUsed   float64 `json:"memused"`
	CpuUsed   float64 `json:"cpuused"`

	Disabled       bool                 `json:"disabled"`
	HostAttributes []*BcsAgentAttribute `json:"host_attributes"`
	Attributes     []*BcsAgentAttribute `json:"attributes"`

	RegisteredTime   int64 `json:"registered_time"`
	ReRegisteredTime int64 `json:"reregistered_time"`
}

// MesosValue_Scalar xxx
type MesosValue_Scalar struct {
	Value float64 `json:"value"`
}

// MesosValue_Ranges xxx
type MesosValue_Ranges struct {
	Begin uint64 `json:"begin"`
	End   uint64 `json:"end"`
}

// MesosValue_Text xxx
type MesosValue_Text struct {
	Value string `json:"value"`
}

// MesosValue_Set xxx
type MesosValue_Set struct {
	Item []string `json:"item"`
}

// MesosValue_Type xxx
type MesosValue_Type int32

const (
	// MesosValueType_UNKNOW xxx
	MesosValueType_UNKNOW MesosValue_Type = 0
	// MesosValueType_Scalar xxx
	MesosValueType_Scalar MesosValue_Type = 101
	// MesosValueType_Ranges xxx
	MesosValueType_Ranges MesosValue_Type = 102
	// MesosValueType_Text xxx
	MesosValueType_Text MesosValue_Type = 103
	// MesosValueType_Set xxx
	MesosValueType_Set MesosValue_Type = 104
)

// BcsClusterAgentSetting xxx
type BcsClusterAgentSetting struct {
	// agent ip
	InnerIP string `json:"innerIP"`
	// whether disable scheduler container
	Disabled bool `json:"disabled"`
	// agent settings
	AttrStrings map[string]MesosValue_Text `json:"strings"`
	// agent settings
	AttrScalars map[string]MesosValue_Scalar `json:"scalars"`
	// taint agent
	NoSchedule map[string]string `json:"noSchedule"`
	// Pods index
	Pods []string `json:"pods"`
	// External Resources, key=ExtendedResource.Name
	ExtendedResources map[string]*ExtendedResource
	// Populated by the system.
	// Read-only.
	// Value must be treated as opaque by clients and .
	ResourceVersion string `json:"-"`
}

// ExtendedResource xxx
type ExtendedResource struct {
	// InnerIP, agent ip
	InnerIP string
	// external resource name, example: bkbcs/cpuset
	Name string
	// Value, container need extended resource value, allocated value
	Value float64
	// Capacity, extended resource total value
	Capacity float64
	// device plugin socket address, example: /data/bcs/cpuset.socket
	Socket string
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsClusterAgentSetting) DeepCopyInto(out *BcsClusterAgentSetting) {
	*out = *in
	if in.AttrStrings != nil {
		in, out := &in.AttrStrings, &out.AttrStrings
		*out = make(map[string]MesosValue_Text, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AttrScalars != nil {
		in, out := &in.AttrScalars, &out.AttrScalars
		*out = make(map[string]MesosValue_Scalar, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NoSchedule != nil {
		in, out := &in.NoSchedule, &out.NoSchedule
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}

	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSchedInfoSpec.
func (in *BcsClusterAgentSetting) DeepCopy() *BcsClusterAgentSetting {
	if in == nil {
		return nil
	}
	out := new(BcsClusterAgentSetting)
	in.DeepCopyInto(out)
	return out
}

// BcsClusterAgentSettingUpdate xxx
type BcsClusterAgentSettingUpdate struct {
	IPs         []string           `json:"ips"`
	SettingName string             `json:"name"`
	ValueType   MesosValue_Type    `json:"valuetype"`
	ValueText   *MesosValue_Text   `json:"text"`
	ValueScalar *MesosValue_Scalar `json:"scalar"`
}

// BcsAgentAttribute xxx
type BcsAgentAttribute struct {
	Name   string               `json:"name,omitempty"`
	Type   MesosValue_Type      `json:"type,omitempty"`
	Scalar *MesosValue_Scalar   `json:"scalar,omitempty"`
	Ranges []*MesosValue_Ranges `json:"ranges,omitempty"`
	Set    *MesosValue_Set      `json:"set,omitempty"`
	Text   *MesosValue_Text     `json:"text,omitempty"`
}
