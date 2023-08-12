

package types

// BcsEndpoint endpoint definition
type BcsEndpoint struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`
	Endpoints  []Endpoint `json:"eps"`
}

// Endpoint endpoint info for ip address
type Endpoint struct {
	NetworkMode string          `json:"networkMode"`
	NodeIP      string          `json:"nodeIP"`
	ContainerIP string          `json:"containerIP"`
	Target      TargetRef       `json:"targetRef,omitempty"`
	Ports       []ContainerPort `json:"ports,omitempty"`
}

// TargetRef reference for endpoint
type TargetRef struct {
	Kind      string `json:"kind"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	// container images
	Image string `json:"image"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsEndpoint) DeepCopyInto(out *BcsEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsEndpointSpec.
func (in *BcsEndpoint) DeepCopy() *BcsEndpoint {
	if in == nil {
		return nil
	}
	out := new(BcsEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	out.Target = in.Target
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]ContainerPort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}
