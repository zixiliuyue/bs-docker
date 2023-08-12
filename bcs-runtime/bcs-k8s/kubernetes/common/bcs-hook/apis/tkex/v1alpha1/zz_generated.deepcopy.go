// +build !ignore_autogenerated



// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Argument) DeepCopyInto(out *Argument) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Argument.
func (in *Argument) DeepCopy() *Argument {
	if in == nil {
		return nil
	}
	out := new(Argument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookRun) DeepCopyInto(out *HookRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookRun.
func (in *HookRun) DeepCopy() *HookRun {
	if in == nil {
		return nil
	}
	out := new(HookRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HookRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookRunArgument) DeepCopyInto(out *HookRunArgument) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookRunArgument.
func (in *HookRunArgument) DeepCopy() *HookRunArgument {
	if in == nil {
		return nil
	}
	out := new(HookRunArgument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookRunList) DeepCopyInto(out *HookRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HookRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookRunList.
func (in *HookRunList) DeepCopy() *HookRunList {
	if in == nil {
		return nil
	}
	out := new(HookRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HookRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookRunSpec) DeepCopyInto(out *HookRunSpec) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]Metric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]Argument, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookRunSpec.
func (in *HookRunSpec) DeepCopy() *HookRunSpec {
	if in == nil {
		return nil
	}
	out := new(HookRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookRunStatus) DeepCopyInto(out *HookRunStatus) {
	*out = *in
	if in.MetricResults != nil {
		in, out := &in.MetricResults, &out.MetricResults
		*out = make([]MetricResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookRunStatus.
func (in *HookRunStatus) DeepCopy() *HookRunStatus {
	if in == nil {
		return nil
	}
	out := new(HookRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookStep) DeepCopyInto(out *HookStep) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]HookRunArgument, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookStep.
func (in *HookStep) DeepCopy() *HookStep {
	if in == nil {
		return nil
	}
	out := new(HookStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookTemplate) DeepCopyInto(out *HookTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookTemplate.
func (in *HookTemplate) DeepCopy() *HookTemplate {
	if in == nil {
		return nil
	}
	out := new(HookTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HookTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookTemplateList) DeepCopyInto(out *HookTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HookTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookTemplateList.
func (in *HookTemplateList) DeepCopy() *HookTemplateList {
	if in == nil {
		return nil
	}
	out := new(HookTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HookTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookTemplateSpec) DeepCopyInto(out *HookTemplateSpec) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]Metric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]Argument, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookTemplateSpec.
func (in *HookTemplateSpec) DeepCopy() *HookTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HookTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Measurement) DeepCopyInto(out *Measurement) {
	*out = *in
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	if in.FinishedAt != nil {
		in, out := &in.FinishedAt, &out.FinishedAt
		*out = (*in).DeepCopy()
	}
	if in.ResumeAt != nil {
		in, out := &in.ResumeAt, &out.ResumeAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Measurement.
func (in *Measurement) DeepCopy() *Measurement {
	if in == nil {
		return nil
	}
	out := new(Measurement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metric) DeepCopyInto(out *Metric) {
	*out = *in
	if in.ConsecutiveErrorLimit != nil {
		in, out := &in.ConsecutiveErrorLimit, &out.ConsecutiveErrorLimit
		*out = new(int32)
		**out = **in
	}
	if in.ConsecutiveSuccessfulLimit != nil {
		in, out := &in.ConsecutiveSuccessfulLimit, &out.ConsecutiveSuccessfulLimit
		*out = new(int32)
		**out = **in
	}
	in.Provider.DeepCopyInto(&out.Provider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metric.
func (in *Metric) DeepCopy() *Metric {
	if in == nil {
		return nil
	}
	out := new(Metric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricProvider) DeepCopyInto(out *MetricProvider) {
	*out = *in
	if in.Web != nil {
		in, out := &in.Web, &out.Web
		*out = new(WebMetric)
		(*in).DeepCopyInto(*out)
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusMetric)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricProvider.
func (in *MetricProvider) DeepCopy() *MetricProvider {
	if in == nil {
		return nil
	}
	out := new(MetricProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricResult) DeepCopyInto(out *MetricResult) {
	*out = *in
	if in.Measurements != nil {
		in, out := &in.Measurements, &out.Measurements
		*out = make([]Measurement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricResult.
func (in *MetricResult) DeepCopy() *MetricResult {
	if in == nil {
		return nil
	}
	out := new(MetricResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PauseCondition) DeepCopyInto(out *PauseCondition) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PauseCondition.
func (in *PauseCondition) DeepCopy() *PauseCondition {
	if in == nil {
		return nil
	}
	out := new(PauseCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreDeleteHookCondition) DeepCopyInto(out *PreDeleteHookCondition) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreDeleteHookCondition.
func (in *PreDeleteHookCondition) DeepCopy() *PreDeleteHookCondition {
	if in == nil {
		return nil
	}
	out := new(PreDeleteHookCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreInplaceHookCondition) DeepCopyInto(out *PreInplaceHookCondition) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreInplaceHookCondition.
func (in *PreInplaceHookCondition) DeepCopy() *PreInplaceHookCondition {
	if in == nil {
		return nil
	}
	out := new(PreInplaceHookCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMetric) DeepCopyInto(out *PrometheusMetric) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMetric.
func (in *PrometheusMetric) DeepCopy() *PrometheusMetric {
	if in == nil {
		return nil
	}
	out := new(PrometheusMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebMetric) DeepCopyInto(out *WebMetric) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]WebMetricHeader, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebMetric.
func (in *WebMetric) DeepCopy() *WebMetric {
	if in == nil {
		return nil
	}
	out := new(WebMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebMetricHeader) DeepCopyInto(out *WebMetricHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebMetricHeader.
func (in *WebMetricHeader) DeepCopy() *WebMetricHeader {
	if in == nil {
		return nil
	}
	out := new(WebMetricHeader)
	in.DeepCopyInto(out)
	return out
}
