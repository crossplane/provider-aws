/*
Copyright 2021 The Crossplane Authors.

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

package deployment

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetDeploymentInput returns input for read
// operation.
func GenerateGetDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.GetDeploymentInput {
	res := &svcsdk.GetDeploymentInput{}

	if cr.Status.AtProvider.DeploymentID != nil {
		res.SetDeploymentId(*cr.Status.AtProvider.DeploymentID)
	}

	return res
}

// GenerateDeployment returns the current state in the form of *svcapitypes.Deployment.
func GenerateDeployment(resp *svcsdk.GetDeploymentOutput) *svcapitypes.Deployment {
	cr := &svcapitypes.Deployment{}

	if resp.AutoDeployed != nil {
		cr.Status.AtProvider.AutoDeployed = resp.AutoDeployed
	}
	if resp.CreatedDate != nil {
		cr.Status.AtProvider.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.DeploymentId != nil {
		cr.Status.AtProvider.DeploymentID = resp.DeploymentId
	}
	if resp.DeploymentStatus != nil {
		cr.Status.AtProvider.DeploymentStatus = resp.DeploymentStatus
	}
	if resp.DeploymentStatusMessage != nil {
		cr.Status.AtProvider.DeploymentStatusMessage = resp.DeploymentStatusMessage
	}

	return cr
}

// GenerateCreateDeploymentInput returns a create input.
func GenerateCreateDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.CreateDeploymentInput {
	res := &svcsdk.CreateDeploymentInput{}

	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.StageName != nil {
		res.SetStageName(*cr.Spec.ForProvider.StageName)
	}

	return res
}

// GenerateUpdateDeploymentInput returns an update input.
func GenerateUpdateDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.UpdateDeploymentInput {
	res := &svcsdk.UpdateDeploymentInput{}

	if cr.Status.AtProvider.DeploymentID != nil {
		res.SetDeploymentId(*cr.Status.AtProvider.DeploymentID)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}

	return res
}

// GenerateDeleteDeploymentInput returns a deletion input.
func GenerateDeleteDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.DeleteDeploymentInput {
	res := &svcsdk.DeleteDeploymentInput{}

	if cr.Status.AtProvider.DeploymentID != nil {
		res.SetDeploymentId(*cr.Status.AtProvider.DeploymentID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
