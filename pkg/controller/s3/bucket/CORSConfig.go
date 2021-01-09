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

package bucket

import (
	"context"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	awss3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/google/go-cmp/cmp"

	"github.com/crossplane/provider-aws/apis/s3/v1beta1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/crossplane/provider-aws/pkg/clients/s3"
)

const (
	corsGetFailed    = "cannot get Bucket CORS configuration"
	corsPutFailed    = "cannot put Bucket cors"
	corsDeleteFailed = "cannot delete Bucket CORS configuration"
)

// CORSConfigurationClient is the client for API methods and reconciling the CORSConfiguration
type CORSConfigurationClient struct {
	client s3.BucketClient
}

// LateInitialize does nothing because CORSConfiguration might have been deleted
// by the user.
func (*CORSConfigurationClient) LateInitialize(_ context.Context, _ *v1beta1.Bucket) error {
	return nil
}

// NewCORSConfigurationClient creates the client for CORS Configuration
func NewCORSConfigurationClient(client s3.BucketClient) *CORSConfigurationClient {
	return &CORSConfigurationClient{client: client}
}

// CompareCORS compares the external and internal representations for the list of CORSRules
func CompareCORS(local []v1beta1.CORSRule, external []awss3types.CORSRule) ResourceStatus { // nolint:gocyclo
	switch {
	case len(local) == 0 && len(external) != 0:
		return NeedsDeletion
	case len(local) == 0 && len(external) == 0:
		return Updated
	case len(local) != len(external):
		return NeedsUpdate
	}

	for i := range local {
		outputRule := external[i]
		if !(cmp.Equal(local[i].AllowedHeaders, outputRule.AllowedHeaders) &&
			cmp.Equal(local[i].AllowedMethods, outputRule.AllowedMethods) &&
			cmp.Equal(local[i].AllowedOrigins, outputRule.AllowedOrigins) &&
			cmp.Equal(local[i].ExposeHeaders, outputRule.ExposeHeaders) &&
			cmp.Equal(aws.ToInt32(local[i].MaxAgeSeconds), outputRule.MaxAgeSeconds)) {
			return NeedsUpdate
		}
	}

	return Updated
}

// Observe checks if the resource exists and if it matches the local configuration
func (in *CORSConfigurationClient) Observe(ctx context.Context, bucket *v1beta1.Bucket) (ResourceStatus, error) {
	result, err := in.client.GetBucketCors(ctx, &awss3.GetBucketCorsInput{Bucket: awsclient.String(meta.GetExternalName(bucket))})
	if resource.Ignore(s3.CORSConfigurationNotFound, err) != nil {
		return NeedsUpdate, awsclient.Wrap(err, corsGetFailed)
	}
	var local []v1beta1.CORSRule
	if bucket.Spec.ForProvider.CORSConfiguration != nil {
		local = bucket.Spec.ForProvider.CORSConfiguration.CORSRules
	}
	var external []awss3types.CORSRule
	if result != nil {
		external = result.CORSRules
	}
	return CompareCORS(local, external), nil
}

// GeneratePutBucketCorsInput creates the input for the PutBucketCors request for the S3 Client
func GeneratePutBucketCorsInput(name string, config *v1beta1.CORSConfiguration) *awss3.PutBucketCorsInput {
	bci := &awss3.PutBucketCorsInput{
		Bucket:            awsclient.String(name),
		CORSConfiguration: &awss3types.CORSConfiguration{CORSRules: make([]awss3types.CORSRule, 0)},
	}
	for _, cors := range config.CORSRules {
		bci.CORSConfiguration.CORSRules = append(bci.CORSConfiguration.CORSRules, awss3types.CORSRule{
			AllowedHeaders: cors.AllowedHeaders,
			AllowedMethods: cors.AllowedMethods,
			AllowedOrigins: cors.AllowedOrigins,
			ExposeHeaders:  cors.ExposeHeaders,
			MaxAgeSeconds:  aws.ToInt32(cors.MaxAgeSeconds),
		})
	}
	return bci
}

// CreateOrUpdate sends a request to have resource created on AWS
func (in *CORSConfigurationClient) CreateOrUpdate(ctx context.Context, bucket *v1beta1.Bucket) error {
	if bucket.Spec.ForProvider.CORSConfiguration == nil {
		return nil
	}
	input := GeneratePutBucketCorsInput(meta.GetExternalName(bucket), bucket.Spec.ForProvider.CORSConfiguration)
	_, err := in.client.PutBucketCors(ctx, input)
	return awsclient.Wrap(err, corsPutFailed)
}

// Delete creates the request to delete the resource on AWS or set it to the default value.
func (in *CORSConfigurationClient) Delete(ctx context.Context, bucket *v1beta1.Bucket) error {
	_, err := in.client.DeleteBucketCors(ctx,
		&awss3.DeleteBucketCorsInput{
			Bucket: awsclient.String(meta.GetExternalName(bucket)),
		},
	)
	return awsclient.Wrap(err, corsDeleteFailed)
}
