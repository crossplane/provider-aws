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

package v1alpha2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSPrincipal) DeepCopyInto(out *AWSPrincipal) {
	*out = *in
	if in.IAMUserARN != nil {
		in, out := &in.IAMUserARN, &out.IAMUserARN
		*out = new(string)
		**out = **in
	}
	if in.IAMUserARNRef != nil {
		in, out := &in.IAMUserARNRef, &out.IAMUserARNRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.IAMUserARNSelector != nil {
		in, out := &in.IAMUserARNSelector, &out.IAMUserARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AWSAccountID != nil {
		in, out := &in.AWSAccountID, &out.AWSAccountID
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleARN != nil {
		in, out := &in.IAMRoleARN, &out.IAMRoleARN
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleARNRef != nil {
		in, out := &in.IAMRoleARNRef, &out.IAMRoleARNRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.IAMRoleARNSelector != nil {
		in, out := &in.IAMRoleARNSelector, &out.IAMRoleARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSPrincipal.
func (in *AWSPrincipal) DeepCopy() *AWSPrincipal {
	if in == nil {
		return nil
	}
	out := new(AWSPrincipal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicy) DeepCopyInto(out *BucketPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicy.
func (in *BucketPolicy) DeepCopy() *BucketPolicy {
	if in == nil {
		return nil
	}
	out := new(BucketPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyList) DeepCopyInto(out *BucketPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyList.
func (in *BucketPolicyList) DeepCopy() *BucketPolicyList {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyParameters) DeepCopyInto(out *BucketPolicyParameters) {
	*out = *in
	if in.Statements != nil {
		in, out := &in.Statements, &out.Statements
		*out = make([]BucketPolicyStatement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.BucketNameRef != nil {
		in, out := &in.BucketNameRef, &out.BucketNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.BucketNameSelector != nil {
		in, out := &in.BucketNameSelector, &out.BucketNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyParameters.
func (in *BucketPolicyParameters) DeepCopy() *BucketPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicySpec) DeepCopyInto(out *BucketPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.PolicyBody.DeepCopyInto(&out.PolicyBody)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicySpec.
func (in *BucketPolicySpec) DeepCopy() *BucketPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BucketPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyStatement) DeepCopyInto(out *BucketPolicyStatement) {
	*out = *in
	if in.SID != nil {
		in, out := &in.SID, &out.SID
		*out = new(string)
		**out = **in
	}
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(BucketPrincipal)
		(*in).DeepCopyInto(*out)
	}
	if in.NotPrincipal != nil {
		in, out := &in.NotPrincipal, &out.NotPrincipal
		*out = new(BucketPrincipal)
		(*in).DeepCopyInto(*out)
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotAction != nil {
		in, out := &in.NotAction, &out.NotAction
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotResource != nil {
		in, out := &in.NotResource, &out.NotResource
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyStatement.
func (in *BucketPolicyStatement) DeepCopy() *BucketPolicyStatement {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyStatement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyStatus) DeepCopyInto(out *BucketPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyStatus.
func (in *BucketPolicyStatus) DeepCopy() *BucketPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPrincipal) DeepCopyInto(out *BucketPrincipal) {
	*out = *in
	if in.AWSPrincipals != nil {
		in, out := &in.AWSPrincipals, &out.AWSPrincipals
		*out = make([]AWSPrincipal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Federated != nil {
		in, out := &in.Federated, &out.Federated
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPrincipal.
func (in *BucketPrincipal) DeepCopy() *BucketPrincipal {
	if in == nil {
		return nil
	}
	out := new(BucketPrincipal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionPair, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionPair) DeepCopyInto(out *ConditionPair) {
	*out = *in
	if in.ConditionStringValue != nil {
		in, out := &in.ConditionStringValue, &out.ConditionStringValue
		*out = new(string)
		**out = **in
	}
	if in.ConditionDateValue != nil {
		in, out := &in.ConditionDateValue, &out.ConditionDateValue
		*out = (*in).DeepCopy()
	}
	if in.ConditionNumericValue != nil {
		in, out := &in.ConditionNumericValue, &out.ConditionNumericValue
		*out = new(int64)
		**out = **in
	}
	if in.ConditionBooleanValue != nil {
		in, out := &in.ConditionBooleanValue, &out.ConditionBooleanValue
		*out = new(bool)
		**out = **in
	}
	if in.ConditionListValue != nil {
		in, out := &in.ConditionListValue, &out.ConditionListValue
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionPair.
func (in *ConditionPair) DeepCopy() *ConditionPair {
	if in == nil {
		return nil
	}
	out := new(ConditionPair)
	in.DeepCopyInto(out)
	return out
}
