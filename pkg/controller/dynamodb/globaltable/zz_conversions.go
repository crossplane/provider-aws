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

// Code generated by ack-generate. DO NOT EDIT.

package globaltable

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/dynamodb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/dynamodb/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.
// TODO(muvaf): We can generate one-time boilerplate for these hooks but currently
// ACK doesn't support not generating if file exists.

// GenerateDescribeGlobalTableInput returns input for read
// operation.
func GenerateDescribeGlobalTableInput(cr *svcapitypes.GlobalTable) *svcsdk.DescribeGlobalTableInput {
	res := preGenerateDescribeGlobalTableInput(cr, &svcsdk.DescribeGlobalTableInput{})

	if cr.Status.AtProvider.GlobalTableName != nil {
		res.SetGlobalTableName(*cr.Status.AtProvider.GlobalTableName)
	}

	return postGenerateDescribeGlobalTableInput(cr, res)
}

// GenerateGlobalTable returns the current state in the form of *svcapitypes.GlobalTable.
func GenerateGlobalTable(resp *svcsdk.DescribeGlobalTableOutput) *svcapitypes.GlobalTable {
	cr := &svcapitypes.GlobalTable{}

	if resp.GlobalTableDescription.CreationDateTime != nil {
		cr.Status.AtProvider.CreationDateTime = &metav1.Time{*resp.GlobalTableDescription.CreationDateTime}
	}
	if resp.GlobalTableDescription.GlobalTableArn != nil {
		cr.Status.AtProvider.GlobalTableARN = resp.GlobalTableDescription.GlobalTableArn
	}
	if resp.GlobalTableDescription.GlobalTableName != nil {
		cr.Status.AtProvider.GlobalTableName = resp.GlobalTableDescription.GlobalTableName
	}
	if resp.GlobalTableDescription.GlobalTableStatus != nil {
		cr.Status.AtProvider.GlobalTableStatus = resp.GlobalTableDescription.GlobalTableStatus
	}

	return cr
}

// GenerateCreateGlobalTableInput returns a create input.
func GenerateCreateGlobalTableInput(cr *svcapitypes.GlobalTable) *svcsdk.CreateGlobalTableInput {
	res := preGenerateCreateGlobalTableInput(cr, &svcsdk.CreateGlobalTableInput{})

	if cr.Spec.ForProvider.ReplicationGroup != nil {
		f0 := []*svcsdk.Replica{}
		for _, f0iter := range cr.Spec.ForProvider.ReplicationGroup {
			f0elem := &svcsdk.Replica{}
			if f0iter.RegionName != nil {
				f0elem.SetRegionName(*f0iter.RegionName)
			}
			f0 = append(f0, f0elem)
		}
		res.SetReplicationGroup(f0)
	}

	return postGenerateCreateGlobalTableInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
