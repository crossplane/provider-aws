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

// Code generated by ack-generate. DO NOT EDIT.

package integrationResponse

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.
// TODO(muvaf): We can generate one-time boilerplate for these hooks but currently
// ACK doesn't support not generating if file exists.
// GenerateGetIntegrationResponsesInput returns input for read
// operation.
func GenerateGetIntegrationResponsesInput(cr *svcapitypes.IntegrationResponse) *svcsdk.GetIntegrationResponsesInput {
	res := preGenerateGetIntegrationResponsesInput(cr, &svcsdk.GetIntegrationResponsesInput{})

	if cr.Spec.ForProvider.APIID != nil {
		res.SetApiId(*cr.Spec.ForProvider.APIID)
	}
	if cr.Spec.ForProvider.IntegrationID != nil {
		res.SetIntegrationId(*cr.Spec.ForProvider.IntegrationID)
	}

	return postGenerateGetIntegrationResponsesInput(cr, res)
}

// GenerateIntegrationResponse returns the current state in the form of *svcapitypes.IntegrationResponse.
func GenerateIntegrationResponse(resp *svcsdk.GetIntegrationResponsesOutput) *svcapitypes.IntegrationResponse {
	cr := &svcapitypes.IntegrationResponse{}

	found := false
	for _, elem := range resp.Items {
		if elem.ContentHandlingStrategy != nil {
			cr.Spec.ForProvider.ContentHandlingStrategy = elem.ContentHandlingStrategy
		}
		if elem.IntegrationResponseId != nil {
			cr.Status.AtProvider.IntegrationResponseID = elem.IntegrationResponseId
		}
		if elem.IntegrationResponseKey != nil {
			cr.Spec.ForProvider.IntegrationResponseKey = elem.IntegrationResponseKey
		}
		if elem.ResponseParameters != nil {
			f3 := map[string]*string{}
			for f3key, f3valiter := range elem.ResponseParameters {
				var f3val string
				f3val = *f3valiter
				f3[f3key] = &f3val
			}
			cr.Spec.ForProvider.ResponseParameters = f3
		}
		if elem.ResponseTemplates != nil {
			f4 := map[string]*string{}
			for f4key, f4valiter := range elem.ResponseTemplates {
				var f4val string
				f4val = *f4valiter
				f4[f4key] = &f4val
			}
			cr.Spec.ForProvider.ResponseTemplates = f4
		}
		if elem.TemplateSelectionExpression != nil {
			cr.Spec.ForProvider.TemplateSelectionExpression = elem.TemplateSelectionExpression
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateIntegrationResponseInput returns a create input.
func GenerateCreateIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.CreateIntegrationResponseInput {
	res := preGenerateCreateIntegrationResponseInput(cr, &svcsdk.CreateIntegrationResponseInput{})

	if cr.Spec.ForProvider.APIID != nil {
		res.SetApiId(*cr.Spec.ForProvider.APIID)
	}
	if cr.Spec.ForProvider.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*cr.Spec.ForProvider.ContentHandlingStrategy)
	}
	if cr.Spec.ForProvider.IntegrationID != nil {
		res.SetIntegrationId(*cr.Spec.ForProvider.IntegrationID)
	}
	if cr.Spec.ForProvider.IntegrationResponseKey != nil {
		res.SetIntegrationResponseKey(*cr.Spec.ForProvider.IntegrationResponseKey)
	}
	if cr.Spec.ForProvider.ResponseParameters != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range cr.Spec.ForProvider.ResponseParameters {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		res.SetResponseParameters(f4)
	}
	if cr.Spec.ForProvider.ResponseTemplates != nil {
		f5 := map[string]*string{}
		for f5key, f5valiter := range cr.Spec.ForProvider.ResponseTemplates {
			var f5val string
			f5val = *f5valiter
			f5[f5key] = &f5val
		}
		res.SetResponseTemplates(f5)
	}
	if cr.Spec.ForProvider.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*cr.Spec.ForProvider.TemplateSelectionExpression)
	}

	return postGenerateCreateIntegrationResponseInput(cr, res)
}

// GenerateDeleteIntegrationResponseInput returns a deletion input.
func GenerateDeleteIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.DeleteIntegrationResponseInput {
	res := preGenerateDeleteIntegrationResponseInput(cr, &svcsdk.DeleteIntegrationResponseInput{})

	if cr.Spec.ForProvider.APIID != nil {
		res.SetApiId(*cr.Spec.ForProvider.APIID)
	}
	if cr.Spec.ForProvider.IntegrationID != nil {
		res.SetIntegrationId(*cr.Spec.ForProvider.IntegrationID)
	}
	if cr.Status.AtProvider.IntegrationResponseID != nil {
		res.SetIntegrationResponseId(*cr.Status.AtProvider.IntegrationResponseID)
	}

	return postGenerateDeleteIntegrationResponseInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
