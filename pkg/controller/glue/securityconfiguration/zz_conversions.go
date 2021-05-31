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

package securityconfiguration

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/glue"

	svcapitypes "github.com/crossplane/provider-aws/apis/glue/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetSecurityConfigurationInput returns input for read
// operation.
func GenerateGetSecurityConfigurationInput(cr *svcapitypes.SecurityConfiguration) *svcsdk.GetSecurityConfigurationInput {
	res := &svcsdk.GetSecurityConfigurationInput{}

	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}

	return res
}

// GenerateSecurityConfiguration returns the current state in the form of *svcapitypes.SecurityConfiguration.
func GenerateSecurityConfiguration(resp *svcsdk.GetSecurityConfigurationOutput) *svcapitypes.SecurityConfiguration {
	cr := &svcapitypes.SecurityConfiguration{}

	return cr
}

// GenerateCreateSecurityConfigurationInput returns a create input.
func GenerateCreateSecurityConfigurationInput(cr *svcapitypes.SecurityConfiguration) *svcsdk.CreateSecurityConfigurationInput {
	res := &svcsdk.CreateSecurityConfigurationInput{}

	if cr.Spec.ForProvider.EncryptionConfiguration != nil {
		f0 := &svcsdk.EncryptionConfiguration{}
		if cr.Spec.ForProvider.EncryptionConfiguration.CloudWatchEncryption != nil {
			f0f0 := &svcsdk.CloudWatchEncryption{}
			if cr.Spec.ForProvider.EncryptionConfiguration.CloudWatchEncryption.CloudWatchEncryptionMode != nil {
				f0f0.SetCloudWatchEncryptionMode(*cr.Spec.ForProvider.EncryptionConfiguration.CloudWatchEncryption.CloudWatchEncryptionMode)
			}
			if cr.Spec.ForProvider.EncryptionConfiguration.CloudWatchEncryption.KMSKeyARN != nil {
				f0f0.SetKmsKeyArn(*cr.Spec.ForProvider.EncryptionConfiguration.CloudWatchEncryption.KMSKeyARN)
			}
			f0.SetCloudWatchEncryption(f0f0)
		}
		if cr.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption != nil {
			f0f1 := &svcsdk.JobBookmarksEncryption{}
			if cr.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.JobBookmarksEncryptionMode != nil {
				f0f1.SetJobBookmarksEncryptionMode(*cr.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.JobBookmarksEncryptionMode)
			}
			if cr.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyARN != nil {
				f0f1.SetKmsKeyArn(*cr.Spec.ForProvider.EncryptionConfiguration.JobBookmarksEncryption.KMSKeyARN)
			}
			f0.SetJobBookmarksEncryption(f0f1)
		}
		if cr.Spec.ForProvider.EncryptionConfiguration.S3Encryption != nil {
			f0f2 := []*svcsdk.S3Encryption{}
			for _, f0f2iter := range cr.Spec.ForProvider.EncryptionConfiguration.S3Encryption {
				f0f2elem := &svcsdk.S3Encryption{}
				if f0f2iter.KMSKeyARN != nil {
					f0f2elem.SetKmsKeyArn(*f0f2iter.KMSKeyARN)
				}
				if f0f2iter.S3EncryptionMode != nil {
					f0f2elem.SetS3EncryptionMode(*f0f2iter.S3EncryptionMode)
				}
				f0f2 = append(f0f2, f0f2elem)
			}
			f0.SetS3Encryption(f0f2)
		}
		res.SetEncryptionConfiguration(f0)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}

	return res
}

// GenerateDeleteSecurityConfigurationInput returns a deletion input.
func GenerateDeleteSecurityConfigurationInput(cr *svcapitypes.SecurityConfiguration) *svcsdk.DeleteSecurityConfigurationInput {
	res := &svcsdk.DeleteSecurityConfigurationInput{}

	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "EntityNotFoundException"
}
