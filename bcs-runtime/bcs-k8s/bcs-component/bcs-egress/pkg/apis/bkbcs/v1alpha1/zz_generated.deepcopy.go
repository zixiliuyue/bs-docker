// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BCSEgress) DeepCopyInto(out *BCSEgress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BCSEgress.
func (in *BCSEgress) DeepCopy() *BCSEgress {
	if in == nil {
		return nil
	}
	out := new(BCSEgress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BCSEgress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BCSEgressList) DeepCopyInto(out *BCSEgressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BCSEgress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BCSEgressList.
func (in *BCSEgressList) DeepCopy() *BCSEgressList {
	if in == nil {
		return nil
	}
	out := new(BCSEgressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BCSEgressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BCSEgressSpec) DeepCopyInto(out *BCSEgressSpec) {
	*out = *in
	out.Controller = in.Controller
	if in.HTTPS != nil {
		in, out := &in.HTTPS, &out.HTTPS
		*out = make([]HTTP, len(*in))
		copy(*out, *in)
	}
	if in.TCPS != nil {
		in, out := &in.TCPS, &out.TCPS
		*out = make([]TCP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BCSEgressSpec.
func (in *BCSEgressSpec) DeepCopy() *BCSEgressSpec {
	if in == nil {
		return nil
	}
	out := new(BCSEgressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BCSEgressStatus) DeepCopyInto(out *BCSEgressStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BCSEgressStatus.
func (in *BCSEgressStatus) DeepCopy() *BCSEgressStatus {
	if in == nil {
		return nil
	}
	out := new(BCSEgressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRef) DeepCopyInto(out *ControllerRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRef.
func (in *ControllerRef) DeepCopy() *ControllerRef {
	if in == nil {
		return nil
	}
	out := new(ControllerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTP) DeepCopyInto(out *HTTP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTP.
func (in *HTTP) DeepCopy() *HTTP {
	if in == nil {
		return nil
	}
	out := new(HTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCP) DeepCopyInto(out *TCP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCP.
func (in *TCP) DeepCopy() *TCP {
	if in == nil {
		return nil
	}
	out := new(TCP)
	in.DeepCopyInto(out)
	return out
}
