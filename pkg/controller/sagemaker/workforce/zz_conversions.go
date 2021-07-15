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

package workforce

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"

	svcapitypes "github.com/crossplane/provider-aws/apis/sagemaker/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeWorkforceInput returns input for read
// operation.
func GenerateDescribeWorkforceInput(cr *svcapitypes.Workforce) *svcsdk.DescribeWorkforceInput {
	res := preGenerateDescribeWorkforceInput(cr, &svcsdk.DescribeWorkforceInput{})

	if cr.Spec.ForProvider.WorkforceName != nil {
		res.SetWorkforceName(*cr.Spec.ForProvider.WorkforceName)
	}

	return postGenerateDescribeWorkforceInput(cr, res)
}

// GenerateWorkforce returns the current state in the form of *svcapitypes.Workforce.
func GenerateWorkforce(resp *svcsdk.DescribeWorkforceOutput) *svcapitypes.Workforce {
	cr := &svcapitypes.Workforce{}

	if resp.Workforce.WorkforceArn != nil {
		cr.Status.AtProvider.WorkforceARN = resp.Workforce.WorkforceArn
	}

	return cr
}

// GenerateCreateWorkforceInput returns a create input.
func GenerateCreateWorkforceInput(cr *svcapitypes.Workforce) *svcsdk.CreateWorkforceInput {
	res := preGenerateCreateWorkforceInput(cr, &svcsdk.CreateWorkforceInput{})

	if cr.Spec.ForProvider.CognitoConfig != nil {
		f0 := &svcsdk.CognitoConfig{}
		if cr.Spec.ForProvider.CognitoConfig.ClientID != nil {
			f0.SetClientId(*cr.Spec.ForProvider.CognitoConfig.ClientID)
		}
		if cr.Spec.ForProvider.CognitoConfig.UserPool != nil {
			f0.SetUserPool(*cr.Spec.ForProvider.CognitoConfig.UserPool)
		}
		res.SetCognitoConfig(f0)
	}
	if cr.Spec.ForProvider.OidcConfig != nil {
		f1 := &svcsdk.OidcConfig{}
		if cr.Spec.ForProvider.OidcConfig.AuthorizationEndpoint != nil {
			f1.SetAuthorizationEndpoint(*cr.Spec.ForProvider.OidcConfig.AuthorizationEndpoint)
		}
		if cr.Spec.ForProvider.OidcConfig.ClientID != nil {
			f1.SetClientId(*cr.Spec.ForProvider.OidcConfig.ClientID)
		}
		if cr.Spec.ForProvider.OidcConfig.ClientSecret != nil {
			f1.SetClientSecret(*cr.Spec.ForProvider.OidcConfig.ClientSecret)
		}
		if cr.Spec.ForProvider.OidcConfig.Issuer != nil {
			f1.SetIssuer(*cr.Spec.ForProvider.OidcConfig.Issuer)
		}
		if cr.Spec.ForProvider.OidcConfig.JwksURI != nil {
			f1.SetJwksUri(*cr.Spec.ForProvider.OidcConfig.JwksURI)
		}
		if cr.Spec.ForProvider.OidcConfig.LogoutEndpoint != nil {
			f1.SetLogoutEndpoint(*cr.Spec.ForProvider.OidcConfig.LogoutEndpoint)
		}
		if cr.Spec.ForProvider.OidcConfig.TokenEndpoint != nil {
			f1.SetTokenEndpoint(*cr.Spec.ForProvider.OidcConfig.TokenEndpoint)
		}
		if cr.Spec.ForProvider.OidcConfig.UserInfoEndpoint != nil {
			f1.SetUserInfoEndpoint(*cr.Spec.ForProvider.OidcConfig.UserInfoEndpoint)
		}
		res.SetOidcConfig(f1)
	}
	if cr.Spec.ForProvider.SourceIPConfig != nil {
		f2 := &svcsdk.SourceIpConfig{}
		if cr.Spec.ForProvider.SourceIPConfig.CIDRs != nil {
			f2f0 := []*string{}
			for _, f2f0iter := range cr.Spec.ForProvider.SourceIPConfig.CIDRs {
				var f2f0elem string
				f2f0elem = *f2f0iter
				f2f0 = append(f2f0, &f2f0elem)
			}
			f2.SetCidrs(f2f0)
		}
		res.SetSourceIpConfig(f2)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f3 := []*svcsdk.Tag{}
		for _, f3iter := range cr.Spec.ForProvider.Tags {
			f3elem := &svcsdk.Tag{}
			if f3iter.Key != nil {
				f3elem.SetKey(*f3iter.Key)
			}
			if f3iter.Value != nil {
				f3elem.SetValue(*f3iter.Value)
			}
			f3 = append(f3, f3elem)
		}
		res.SetTags(f3)
	}
	if cr.Spec.ForProvider.WorkforceName != nil {
		res.SetWorkforceName(*cr.Spec.ForProvider.WorkforceName)
	}

	return postGenerateCreateWorkforceInput(cr, res)
}

// GenerateDeleteWorkforceInput returns a deletion input.
func GenerateDeleteWorkforceInput(cr *svcapitypes.Workforce) *svcsdk.DeleteWorkforceInput {
	res := preGenerateDeleteWorkforceInput(cr, &svcsdk.DeleteWorkforceInput{})

	if cr.Spec.ForProvider.WorkforceName != nil {
		res.SetWorkforceName(*cr.Spec.ForProvider.WorkforceName)
	}

	return postGenerateDeleteWorkforceInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}