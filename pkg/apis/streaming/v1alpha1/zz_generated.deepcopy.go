// +build !ignore_autogenerated

/*
Copyright 2019 the original author or authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	duckv1 "github.com/projectriff/system/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingReference) DeepCopyInto(out *BindingReference) {
	*out = *in
	out.MetadataRef = in.MetadataRef
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingReference.
func (in *BindingReference) DeepCopy() *BindingReference {
	if in == nil {
		return nil
	}
	out := new(BindingReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Build) DeepCopyInto(out *Build) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Build.
func (in *Build) DeepCopy() *Build {
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayList) DeepCopyInto(out *GatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayList.
func (in *GatewayList) DeepCopy() *GatewayList {
	if in == nil {
		return nil
	}
	out := new(GatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpec) DeepCopyInto(out *GatewaySpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpec.
func (in *GatewaySpec) DeepCopy() *GatewaySpec {
	if in == nil {
		return nil
	}
	out := new(GatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayStatus) DeepCopyInto(out *GatewayStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1.Addressable)
		**out = **in
	}
	if in.DeploymentRef != nil {
		in, out := &in.DeploymentRef, &out.DeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.ServiceRef != nil {
		in, out := &in.ServiceRef, &out.ServiceRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayStatus.
func (in *GatewayStatus) DeepCopy() *GatewayStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryGateway) DeepCopyInto(out *InMemoryGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryGateway.
func (in *InMemoryGateway) DeepCopy() *InMemoryGateway {
	if in == nil {
		return nil
	}
	out := new(InMemoryGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryGatewayList) DeepCopyInto(out *InMemoryGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InMemoryGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryGatewayList.
func (in *InMemoryGatewayList) DeepCopy() *InMemoryGatewayList {
	if in == nil {
		return nil
	}
	out := new(InMemoryGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryGatewaySpec) DeepCopyInto(out *InMemoryGatewaySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryGatewaySpec.
func (in *InMemoryGatewaySpec) DeepCopy() *InMemoryGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(InMemoryGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryGatewayStatus) DeepCopyInto(out *InMemoryGatewayStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1.Addressable)
		**out = **in
	}
	if in.GatewayRef != nil {
		in, out := &in.GatewayRef, &out.GatewayRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryGatewayStatus.
func (in *InMemoryGatewayStatus) DeepCopy() *InMemoryGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(InMemoryGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputStreamBinding) DeepCopyInto(out *InputStreamBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputStreamBinding.
func (in *InputStreamBinding) DeepCopy() *InputStreamBinding {
	if in == nil {
		return nil
	}
	out := new(InputStreamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaGateway) DeepCopyInto(out *KafkaGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaGateway.
func (in *KafkaGateway) DeepCopy() *KafkaGateway {
	if in == nil {
		return nil
	}
	out := new(KafkaGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaGatewayList) DeepCopyInto(out *KafkaGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaGatewayList.
func (in *KafkaGatewayList) DeepCopy() *KafkaGatewayList {
	if in == nil {
		return nil
	}
	out := new(KafkaGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaGatewaySpec) DeepCopyInto(out *KafkaGatewaySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaGatewaySpec.
func (in *KafkaGatewaySpec) DeepCopy() *KafkaGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(KafkaGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaGatewayStatus) DeepCopyInto(out *KafkaGatewayStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1.Addressable)
		**out = **in
	}
	if in.GatewayRef != nil {
		in, out := &in.GatewayRef, &out.GatewayRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaGatewayStatus.
func (in *KafkaGatewayStatus) DeepCopy() *KafkaGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputStreamBinding) DeepCopyInto(out *OutputStreamBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputStreamBinding.
func (in *OutputStreamBinding) DeepCopy() *OutputStreamBinding {
	if in == nil {
		return nil
	}
	out := new(OutputStreamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Processor) DeepCopyInto(out *Processor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Processor.
func (in *Processor) DeepCopy() *Processor {
	if in == nil {
		return nil
	}
	out := new(Processor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Processor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessorList) DeepCopyInto(out *ProcessorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Processor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessorList.
func (in *ProcessorList) DeepCopy() *ProcessorList {
	if in == nil {
		return nil
	}
	out := new(ProcessorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProcessorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessorSpec) DeepCopyInto(out *ProcessorSpec) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Build)
		**out = **in
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]InputStreamBinding, len(*in))
		copy(*out, *in)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]OutputStreamBinding, len(*in))
		copy(*out, *in)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessorSpec.
func (in *ProcessorSpec) DeepCopy() *ProcessorSpec {
	if in == nil {
		return nil
	}
	out := new(ProcessorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessorStatus) DeepCopyInto(out *ProcessorStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.DeploymentRef != nil {
		in, out := &in.DeploymentRef, &out.DeploymentRef
		*out = (*in).DeepCopy()
	}
	if in.ScaledObjectRef != nil {
		in, out := &in.ScaledObjectRef, &out.ScaledObjectRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessorStatus.
func (in *ProcessorStatus) DeepCopy() *ProcessorStatus {
	if in == nil {
		return nil
	}
	out := new(ProcessorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarGateway) DeepCopyInto(out *PulsarGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarGateway.
func (in *PulsarGateway) DeepCopy() *PulsarGateway {
	if in == nil {
		return nil
	}
	out := new(PulsarGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarGatewayList) DeepCopyInto(out *PulsarGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulsarGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarGatewayList.
func (in *PulsarGatewayList) DeepCopy() *PulsarGatewayList {
	if in == nil {
		return nil
	}
	out := new(PulsarGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarGatewaySpec) DeepCopyInto(out *PulsarGatewaySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarGatewaySpec.
func (in *PulsarGatewaySpec) DeepCopy() *PulsarGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(PulsarGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarGatewayStatus) DeepCopyInto(out *PulsarGatewayStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(duckv1.Addressable)
		**out = **in
	}
	if in.GatewayRef != nil {
		in, out := &in.GatewayRef, &out.GatewayRef
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarGatewayStatus.
func (in *PulsarGatewayStatus) DeepCopy() *PulsarGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(PulsarGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stream) DeepCopyInto(out *Stream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stream.
func (in *Stream) DeepCopy() *Stream {
	if in == nil {
		return nil
	}
	out := new(Stream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamList) DeepCopyInto(out *StreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamList.
func (in *StreamList) DeepCopy() *StreamList {
	if in == nil {
		return nil
	}
	out := new(StreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSpec) DeepCopyInto(out *StreamSpec) {
	*out = *in
	out.Gateway = in.Gateway
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSpec.
func (in *StreamSpec) DeepCopy() *StreamSpec {
	if in == nil {
		return nil
	}
	out := new(StreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamStatus) DeepCopyInto(out *StreamStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	out.Binding = in.Binding
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamStatus.
func (in *StreamStatus) DeepCopy() *StreamStatus {
	if in == nil {
		return nil
	}
	out := new(StreamStatus)
	in.DeepCopyInto(out)
	return out
}
