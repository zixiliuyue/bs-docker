// +build !ignore_autogenerated



// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsDbPrivConfig) DeepCopyInto(out *BcsDbPrivConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsDbPrivConfig.
func (in *BcsDbPrivConfig) DeepCopy() *BcsDbPrivConfig {
	if in == nil {
		return nil
	}
	out := new(BcsDbPrivConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BcsDbPrivConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsDbPrivConfigList) DeepCopyInto(out *BcsDbPrivConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BcsDbPrivConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsDbPrivConfigList.
func (in *BcsDbPrivConfigList) DeepCopy() *BcsDbPrivConfigList {
	if in == nil {
		return nil
	}
	out := new(BcsDbPrivConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BcsDbPrivConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsDbPrivConfigSpec) DeepCopyInto(out *BcsDbPrivConfigSpec) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsDbPrivConfigSpec.
func (in *BcsDbPrivConfigSpec) DeepCopy() *BcsDbPrivConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BcsDbPrivConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsLogConfig) DeepCopyInto(out *BcsLogConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsLogConfig.
func (in *BcsLogConfig) DeepCopy() *BcsLogConfig {
	if in == nil {
		return nil
	}
	out := new(BcsLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BcsLogConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsLogConfigList) DeepCopyInto(out *BcsLogConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BcsLogConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsLogConfigList.
func (in *BcsLogConfigList) DeepCopy() *BcsLogConfigList {
	if in == nil {
		return nil
	}
	out := new(BcsLogConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BcsLogConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsLogConfigSpec) DeepCopyInto(out *BcsLogConfigSpec) {
	*out = *in
	if in.LogPaths != nil {
		in, out := &in.LogPaths, &out.LogPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostPaths != nil {
		in, out := &in.HostPaths, &out.HostPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogTags != nil {
		in, out := &in.LogTags, &out.LogTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Multiline != nil {
		in, out := &in.Multiline, &out.Multiline
		*out = new(MultilineConf)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerConfs != nil {
		in, out := &in.ContainerConfs, &out.ContainerConfs
		*out = make([]ContainerConf, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsLogConfigSpec.
func (in *BcsLogConfigSpec) DeepCopy() *BcsLogConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BcsLogConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerConf) DeepCopyInto(out *ContainerConf) {
	*out = *in
	if in.HostPaths != nil {
		in, out := &in.HostPaths, &out.HostPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogPaths != nil {
		in, out := &in.LogPaths, &out.LogPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogTags != nil {
		in, out := &in.LogTags, &out.LogTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Multiline != nil {
		in, out := &in.Multiline, &out.Multiline
		*out = new(MultilineConf)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerConf.
func (in *ContainerConf) DeepCopy() *ContainerConf {
	if in == nil {
		return nil
	}
	out := new(ContainerConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultilineConf) DeepCopyInto(out *MultilineConf) {
	*out = *in
	if in.MaxLines != nil {
		in, out := &in.MaxLines, &out.MaxLines
		*out = new(int)
		**out = **in
	}
	if in.SkipNewline != nil {
		in, out := &in.SkipNewline, &out.SkipNewline
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultilineConf.
func (in *MultilineConf) DeepCopy() *MultilineConf {
	if in == nil {
		return nil
	}
	out := new(MultilineConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSelector) DeepCopyInto(out *PodSelector) {
	*out = *in
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]SelectorExpression, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSelector.
func (in *PodSelector) DeepCopy() *PodSelector {
	if in == nil {
		return nil
	}
	out := new(PodSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorExpression) DeepCopyInto(out *SelectorExpression) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorExpression.
func (in *SelectorExpression) DeepCopy() *SelectorExpression {
	if in == nil {
		return nil
	}
	out := new(SelectorExpression)
	in.DeepCopyInto(out)
	return out
}
