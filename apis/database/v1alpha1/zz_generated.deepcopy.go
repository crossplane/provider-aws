// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributeDefinition) DeepCopyInto(out *AttributeDefinition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributeDefinition.
func (in *AttributeDefinition) DeepCopy() *AttributeDefinition {
	if in == nil {
		return nil
	}
	out := new(AttributeDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoTable) DeepCopyInto(out *DynamoTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoTable.
func (in *DynamoTable) DeepCopy() *DynamoTable {
	if in == nil {
		return nil
	}
	out := new(DynamoTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoTableList) DeepCopyInto(out *DynamoTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamoTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoTableList.
func (in *DynamoTableList) DeepCopy() *DynamoTableList {
	if in == nil {
		return nil
	}
	out := new(DynamoTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoTableObservation) DeepCopyInto(out *DynamoTableObservation) {
	*out = *in
	if in.AttributeDefinitions != nil {
		in, out := &in.AttributeDefinitions, &out.AttributeDefinitions
		*out = make([]AttributeDefinition, len(*in))
		copy(*out, *in)
	}
	if in.GlobalSecondaryIndexes != nil {
		in, out := &in.GlobalSecondaryIndexes, &out.GlobalSecondaryIndexes
		*out = make([]GlobalSecondaryIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KeySchema != nil {
		in, out := &in.KeySchema, &out.KeySchema
		*out = make([]KeySchemaElement, len(*in))
		copy(*out, *in)
	}
	if in.LocalSecondaryIndexes != nil {
		in, out := &in.LocalSecondaryIndexes, &out.LocalSecondaryIndexes
		*out = make([]LocalSecondaryIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ProvisionedThroughput.DeepCopyInto(&out.ProvisionedThroughput)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoTableObservation.
func (in *DynamoTableObservation) DeepCopy() *DynamoTableObservation {
	if in == nil {
		return nil
	}
	out := new(DynamoTableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoTableParameters) DeepCopyInto(out *DynamoTableParameters) {
	*out = *in
	if in.AttributeDefinitions != nil {
		in, out := &in.AttributeDefinitions, &out.AttributeDefinitions
		*out = make([]AttributeDefinition, len(*in))
		copy(*out, *in)
	}
	if in.GlobalSecondaryIndexes != nil {
		in, out := &in.GlobalSecondaryIndexes, &out.GlobalSecondaryIndexes
		*out = make([]GlobalSecondaryIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KeySchema != nil {
		in, out := &in.KeySchema, &out.KeySchema
		*out = make([]KeySchemaElement, len(*in))
		copy(*out, *in)
	}
	if in.LocalSecondaryIndexes != nil {
		in, out := &in.LocalSecondaryIndexes, &out.LocalSecondaryIndexes
		*out = make([]LocalSecondaryIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisionedThroughput != nil {
		in, out := &in.ProvisionedThroughput, &out.ProvisionedThroughput
		*out = new(ProvisionedThroughput)
		(*in).DeepCopyInto(*out)
	}
	if in.SSESpecification != nil {
		in, out := &in.SSESpecification, &out.SSESpecification
		*out = new(SSESpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.StreamSpecification != nil {
		in, out := &in.StreamSpecification, &out.StreamSpecification
		*out = new(StreamSpecification)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoTableParameters.
func (in *DynamoTableParameters) DeepCopy() *DynamoTableParameters {
	if in == nil {
		return nil
	}
	out := new(DynamoTableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoTableSpec) DeepCopyInto(out *DynamoTableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoTableSpec.
func (in *DynamoTableSpec) DeepCopy() *DynamoTableSpec {
	if in == nil {
		return nil
	}
	out := new(DynamoTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoTableStatus) DeepCopyInto(out *DynamoTableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoTableStatus.
func (in *DynamoTableStatus) DeepCopy() *DynamoTableStatus {
	if in == nil {
		return nil
	}
	out := new(DynamoTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalSecondaryIndex) DeepCopyInto(out *GlobalSecondaryIndex) {
	*out = *in
	if in.IndexName != nil {
		in, out := &in.IndexName, &out.IndexName
		*out = new(string)
		**out = **in
	}
	if in.KeySchema != nil {
		in, out := &in.KeySchema, &out.KeySchema
		*out = make([]KeySchemaElement, len(*in))
		copy(*out, *in)
	}
	if in.Projection != nil {
		in, out := &in.Projection, &out.Projection
		*out = new(Projection)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisionedThroughput != nil {
		in, out := &in.ProvisionedThroughput, &out.ProvisionedThroughput
		*out = new(ProvisionedThroughput)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalSecondaryIndex.
func (in *GlobalSecondaryIndex) DeepCopy() *GlobalSecondaryIndex {
	if in == nil {
		return nil
	}
	out := new(GlobalSecondaryIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeySchemaElement) DeepCopyInto(out *KeySchemaElement) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeySchemaElement.
func (in *KeySchemaElement) DeepCopy() *KeySchemaElement {
	if in == nil {
		return nil
	}
	out := new(KeySchemaElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSecondaryIndex) DeepCopyInto(out *LocalSecondaryIndex) {
	*out = *in
	if in.IndexName != nil {
		in, out := &in.IndexName, &out.IndexName
		*out = new(string)
		**out = **in
	}
	if in.KeySchema != nil {
		in, out := &in.KeySchema, &out.KeySchema
		*out = make([]KeySchemaElement, len(*in))
		copy(*out, *in)
	}
	if in.Projection != nil {
		in, out := &in.Projection, &out.Projection
		*out = new(Projection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSecondaryIndex.
func (in *LocalSecondaryIndex) DeepCopy() *LocalSecondaryIndex {
	if in == nil {
		return nil
	}
	out := new(LocalSecondaryIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Projection) DeepCopyInto(out *Projection) {
	*out = *in
	if in.NonKeyAttributes != nil {
		in, out := &in.NonKeyAttributes, &out.NonKeyAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Projection.
func (in *Projection) DeepCopy() *Projection {
	if in == nil {
		return nil
	}
	out := new(Projection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedThroughput) DeepCopyInto(out *ProvisionedThroughput) {
	*out = *in
	if in.ReadCapacityUnits != nil {
		in, out := &in.ReadCapacityUnits, &out.ReadCapacityUnits
		*out = new(int64)
		**out = **in
	}
	if in.WriteCapacityUnits != nil {
		in, out := &in.WriteCapacityUnits, &out.WriteCapacityUnits
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedThroughput.
func (in *ProvisionedThroughput) DeepCopy() *ProvisionedThroughput {
	if in == nil {
		return nil
	}
	out := new(ProvisionedThroughput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSESpecification) DeepCopyInto(out *SSESpecification) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KMSMasterKeyID != nil {
		in, out := &in.KMSMasterKeyID, &out.KMSMasterKeyID
		*out = new(string)
		**out = **in
	}
	if in.SSEType != nil {
		in, out := &in.SSEType, &out.SSEType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSESpecification.
func (in *SSESpecification) DeepCopy() *SSESpecification {
	if in == nil {
		return nil
	}
	out := new(SSESpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSpecification) DeepCopyInto(out *StreamSpecification) {
	*out = *in
	if in.StreamEnabled != nil {
		in, out := &in.StreamEnabled, &out.StreamEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSpecification.
func (in *StreamSpecification) DeepCopy() *StreamSpecification {
	if in == nil {
		return nil
	}
	out := new(StreamSpecification)
	in.DeepCopyInto(out)
	return out
}
