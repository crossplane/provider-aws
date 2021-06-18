// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHTTPNamespaceParameters) DeepCopyInto(out *CustomHTTPNamespaceParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHTTPNamespaceParameters.
func (in *CustomHTTPNamespaceParameters) DeepCopy() *CustomHTTPNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(CustomHTTPNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPrivateDNSNamespaceParameters) DeepCopyInto(out *CustomPrivateDNSNamespaceParameters) {
	*out = *in
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(string)
		**out = **in
	}
	if in.VPCRef != nil {
		in, out := &in.VPCRef, &out.VPCRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.VPCSelector != nil {
		in, out := &in.VPCSelector, &out.VPCSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPrivateDNSNamespaceParameters.
func (in *CustomPrivateDNSNamespaceParameters) DeepCopy() *CustomPrivateDNSNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(CustomPrivateDNSNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPublicDNSNamespaceParameters) DeepCopyInto(out *CustomPublicDNSNamespaceParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPublicDNSNamespaceParameters.
func (in *CustomPublicDNSNamespaceParameters) DeepCopy() *CustomPublicDNSNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(CustomPublicDNSNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomServiceParameters) DeepCopyInto(out *CustomServiceParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomServiceParameters.
func (in *CustomServiceParameters) DeepCopy() *CustomServiceParameters {
	if in == nil {
		return nil
	}
	out := new(CustomServiceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPInstanceSummary) DeepCopyInto(out *HTTPInstanceSummary) {
	*out = *in
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPInstanceSummary.
func (in *HTTPInstanceSummary) DeepCopy() *HTTPInstanceSummary {
	if in == nil {
		return nil
	}
	out := new(HTTPInstanceSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPNamespace) DeepCopyInto(out *HTTPNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPNamespace.
func (in *HTTPNamespace) DeepCopy() *HTTPNamespace {
	if in == nil {
		return nil
	}
	out := new(HTTPNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPNamespaceList) DeepCopyInto(out *HTTPNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPNamespaceList.
func (in *HTTPNamespaceList) DeepCopy() *HTTPNamespaceList {
	if in == nil {
		return nil
	}
	out := new(HTTPNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPNamespaceObservation) DeepCopyInto(out *HTTPNamespaceObservation) {
	*out = *in
	if in.OperationID != nil {
		in, out := &in.OperationID, &out.OperationID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPNamespaceObservation.
func (in *HTTPNamespaceObservation) DeepCopy() *HTTPNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(HTTPNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPNamespaceParameters) DeepCopyInto(out *HTTPNamespaceParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.CustomHTTPNamespaceParameters = in.CustomHTTPNamespaceParameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPNamespaceParameters.
func (in *HTTPNamespaceParameters) DeepCopy() *HTTPNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPNamespaceSpec) DeepCopyInto(out *HTTPNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPNamespaceSpec.
func (in *HTTPNamespaceSpec) DeepCopy() *HTTPNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPNamespaceStatus) DeepCopyInto(out *HTTPNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPNamespaceStatus.
func (in *HTTPNamespaceStatus) DeepCopy() *HTTPNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(HTTPNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProperties) DeepCopyInto(out *HTTPProperties) {
	*out = *in
	if in.HTTPName != nil {
		in, out := &in.HTTPName, &out.HTTPName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProperties.
func (in *HTTPProperties) DeepCopy() *HTTPProperties {
	if in == nil {
		return nil
	}
	out := new(HTTPProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Namespace) DeepCopyInto(out *Namespace) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namespace.
func (in *Namespace) DeepCopy() *Namespace {
	if in == nil {
		return nil
	}
	out := new(Namespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSummary) DeepCopyInto(out *NamespaceSummary) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSummary.
func (in *NamespaceSummary) DeepCopy() *NamespaceSummary {
	if in == nil {
		return nil
	}
	out := new(NamespaceSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operation) DeepCopyInto(out *Operation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationSummary) DeepCopyInto(out *OperationSummary) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationSummary.
func (in *OperationSummary) DeepCopy() *OperationSummary {
	if in == nil {
		return nil
	}
	out := new(OperationSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDNSNamespace) DeepCopyInto(out *PrivateDNSNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDNSNamespace.
func (in *PrivateDNSNamespace) DeepCopy() *PrivateDNSNamespace {
	if in == nil {
		return nil
	}
	out := new(PrivateDNSNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateDNSNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDNSNamespaceList) DeepCopyInto(out *PrivateDNSNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrivateDNSNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDNSNamespaceList.
func (in *PrivateDNSNamespaceList) DeepCopy() *PrivateDNSNamespaceList {
	if in == nil {
		return nil
	}
	out := new(PrivateDNSNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateDNSNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDNSNamespaceObservation) DeepCopyInto(out *PrivateDNSNamespaceObservation) {
	*out = *in
	if in.OperationID != nil {
		in, out := &in.OperationID, &out.OperationID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDNSNamespaceObservation.
func (in *PrivateDNSNamespaceObservation) DeepCopy() *PrivateDNSNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(PrivateDNSNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDNSNamespaceParameters) DeepCopyInto(out *PrivateDNSNamespaceParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.CustomPrivateDNSNamespaceParameters.DeepCopyInto(&out.CustomPrivateDNSNamespaceParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDNSNamespaceParameters.
func (in *PrivateDNSNamespaceParameters) DeepCopy() *PrivateDNSNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(PrivateDNSNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDNSNamespaceSpec) DeepCopyInto(out *PrivateDNSNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDNSNamespaceSpec.
func (in *PrivateDNSNamespaceSpec) DeepCopy() *PrivateDNSNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(PrivateDNSNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDNSNamespaceStatus) DeepCopyInto(out *PrivateDNSNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDNSNamespaceStatus.
func (in *PrivateDNSNamespaceStatus) DeepCopy() *PrivateDNSNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(PrivateDNSNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDNSNamespace) DeepCopyInto(out *PublicDNSNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDNSNamespace.
func (in *PublicDNSNamespace) DeepCopy() *PublicDNSNamespace {
	if in == nil {
		return nil
	}
	out := new(PublicDNSNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicDNSNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDNSNamespaceList) DeepCopyInto(out *PublicDNSNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublicDNSNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDNSNamespaceList.
func (in *PublicDNSNamespaceList) DeepCopy() *PublicDNSNamespaceList {
	if in == nil {
		return nil
	}
	out := new(PublicDNSNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicDNSNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDNSNamespaceObservation) DeepCopyInto(out *PublicDNSNamespaceObservation) {
	*out = *in
	if in.OperationID != nil {
		in, out := &in.OperationID, &out.OperationID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDNSNamespaceObservation.
func (in *PublicDNSNamespaceObservation) DeepCopy() *PublicDNSNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(PublicDNSNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDNSNamespaceParameters) DeepCopyInto(out *PublicDNSNamespaceParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.CustomPublicDNSNamespaceParameters = in.CustomPublicDNSNamespaceParameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDNSNamespaceParameters.
func (in *PublicDNSNamespaceParameters) DeepCopy() *PublicDNSNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(PublicDNSNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDNSNamespaceSpec) DeepCopyInto(out *PublicDNSNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDNSNamespaceSpec.
func (in *PublicDNSNamespaceSpec) DeepCopy() *PublicDNSNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(PublicDNSNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDNSNamespaceStatus) DeepCopyInto(out *PublicDNSNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDNSNamespaceStatus.
func (in *PublicDNSNamespaceStatus) DeepCopy() *PublicDNSNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(PublicDNSNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceChange) DeepCopyInto(out *ServiceChange) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceChange.
func (in *ServiceChange) DeepCopy() *ServiceChange {
	if in == nil {
		return nil
	}
	out := new(ServiceChange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSummary) DeepCopyInto(out *ServiceSummary) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSummary.
func (in *ServiceSummary) DeepCopy() *ServiceSummary {
	if in == nil {
		return nil
	}
	out := new(ServiceSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
