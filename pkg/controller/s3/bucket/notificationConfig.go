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

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	awss3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/google/go-cmp/cmp"

	"github.com/crossplane/provider-aws/apis/s3/v1beta1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/crossplane/provider-aws/pkg/clients/s3"
)

const (
	notificationGetFailed = "cannot get Bucket notification"
	notificationPutFailed = "cannot put Bucket notification"
)

// NotificationConfigurationClient is the client for API methods and reconciling the LifecycleConfiguration
type NotificationConfigurationClient struct {
	client s3.BucketClient
}

// LateInitialize is responsible for initializing the resource based on the external value
func (in *NotificationConfigurationClient) LateInitialize(ctx context.Context, bucket *v1beta1.Bucket) error {
	external, err := in.client.GetBucketNotificationConfiguration(ctx, &awss3.GetBucketNotificationConfigurationInput{Bucket: awsclient.String(meta.GetExternalName(bucket))})
	if err != nil {
		return awsclient.Wrap(err, notificationGetFailed)
	}
	if emptyConfiguration(external) {
		// There is nothing to initialize from AWS
		return nil
	}
	config := bucket.Spec.ForProvider.NotificationConfiguration
	if config == nil {
		// We need the configuration to exist so we can initialize
		bucket.Spec.ForProvider.NotificationConfiguration = &v1beta1.NotificationConfiguration{}
		config = bucket.Spec.ForProvider.NotificationConfiguration
	}

	// A list is provided by AWS
	if external.LambdaFunctionConfigurations != nil {
		if config.LambdaFunctionConfigurations == nil {
			config.LambdaFunctionConfigurations = make([]v1beta1.LambdaFunctionConfiguration, len(external.LambdaFunctionConfigurations))
		}
		LateInitializeLambda(external.LambdaFunctionConfigurations, config.LambdaFunctionConfigurations)
	}

	// A list is provided by AWS
	if external.QueueConfigurations != nil {
		if config.QueueConfigurations == nil {
			config.QueueConfigurations = make([]v1beta1.QueueConfiguration, len(external.QueueConfigurations))
		}
		LateInitializeQueue(external.QueueConfigurations, config.QueueConfigurations)
	}

	// A list is provided by AWS
	if external.TopicConfigurations != nil {
		if config.TopicConfigurations == nil {
			config.TopicConfigurations = make([]v1beta1.TopicConfiguration, len(external.TopicConfigurations))
		}
		LateInitializeTopic(external.TopicConfigurations, config.TopicConfigurations)
	}
	return nil
}

// LateInitializeFilter initializes the external awss3types.NotificationConfigurationFilter to a local v1beta.NotificationConfigurationFilter
func LateInitializeFilter(local *v1beta1.NotificationConfigurationFilter, external *awss3types.NotificationConfigurationFilter) *v1beta1.NotificationConfigurationFilter {
	if local != nil {
		return local
	}
	if external == nil {
		return nil
	}
	local = &v1beta1.NotificationConfigurationFilter{}
	if external.Key == nil {
		return local
	}
	local.Key = &v1beta1.S3KeyFilter{}
	if external.Key.FilterRules != nil {
		local.Key.FilterRules = make([]v1beta1.FilterRule, len(external.Key.FilterRules))
		for i, v := range external.Key.FilterRules {
			local.Key.FilterRules[i] = v1beta1.FilterRule{
				Name:  string(v.Name),
				Value: v.Value,
			}
		}
	}
	return local
}

// LateInitializeEvents initializes the external []awss3types.Event to a local []string
func LateInitializeEvents(local []string, external []awss3types.Event) []string {
	if local != nil {
		return local
	}
	newLocal := make([]string, len(external))
	for i, v := range external {
		newLocal[i] = string(v)
	}
	return newLocal
}

// LateInitializeLambda initializes the external awss3types.LambdaFunctionConfiguration to a local v1beta.LambdaFunctionConfiguration
func LateInitializeLambda(external []awss3types.LambdaFunctionConfiguration, local []v1beta1.LambdaFunctionConfiguration) {
	for i, v := range external {
		if i >= len(local) {
			break
		}
		local[i] = v1beta1.LambdaFunctionConfiguration{
			Events:            LateInitializeEvents(local[i].Events, v.Events),
			Filter:            LateInitializeFilter(local[i].Filter, v.Filter),
			ID:                awsclient.LateInitializeStringPtr(local[i].ID, v.Id),
			LambdaFunctionArn: awsclient.LateInitializeString(local[i].LambdaFunctionArn, v.LambdaFunctionArn),
		}
	}
}

// LateInitializeQueue initializes the external awss3types.QueueConfiguration to a local v1beta.QueueConfiguration
func LateInitializeQueue(external []awss3types.QueueConfiguration, local []v1beta1.QueueConfiguration) {
	for i, v := range external {
		if i >= len(local) {
			break
		}
		local[i] = v1beta1.QueueConfiguration{
			Events:   LateInitializeEvents(local[i].Events, v.Events),
			Filter:   LateInitializeFilter(local[i].Filter, v.Filter),
			ID:       awsclient.LateInitializeStringPtr(local[i].ID, v.Id),
			QueueArn: awsclient.LateInitializeString(local[i].QueueArn, v.QueueArn),
		}
	}
}

// LateInitializeTopic initializes the external awss3types.TopicConfiguration to a local v1beta.TopicConfiguration
func LateInitializeTopic(external []awss3types.TopicConfiguration, local []v1beta1.TopicConfiguration) {
	for i, v := range external {
		if i >= len(local) {
			break
		}
		local[i] = v1beta1.TopicConfiguration{
			Events:   LateInitializeEvents(local[i].Events, v.Events),
			Filter:   LateInitializeFilter(local[i].Filter, v.Filter),
			ID:       awsclient.LateInitializeStringPtr(local[i].ID, v.Id),
			TopicArn: awsclient.LateInitializeStringPtr(local[i].TopicArn, v.TopicArn),
		}
	}
}

// NewNotificationConfigurationClient creates the client for Accelerate Configuration
func NewNotificationConfigurationClient(client s3.BucketClient) *NotificationConfigurationClient {
	return &NotificationConfigurationClient{client: client}
}

func emptyConfiguration(external *awss3.GetBucketNotificationConfigurationOutput) bool {
	return external == nil || len(external.TopicConfigurations) == 0 || len(external.QueueConfigurations) == 0 || len(external.LambdaFunctionConfigurations) == 0
}

func bucketStatus(config *v1beta1.NotificationConfiguration, external *awss3.GetBucketNotificationConfigurationOutput) ResourceStatus { // nolint:gocyclo
	if config == nil && len(external.QueueConfigurations) == 0 && len(external.LambdaFunctionConfigurations) == 0 && len(external.TopicConfigurations) == 0 {
		return Updated
	} else if config == nil && (len(external.QueueConfigurations) != 0 || len(external.LambdaFunctionConfigurations) != 0 || len(external.TopicConfigurations) != 0) {
		return NeedsDeletion
	}
	return NeedsUpdate
}

// Observe checks if the resource exists and if it matches the local configuration
func (in *NotificationConfigurationClient) Observe(ctx context.Context, bucket *v1beta1.Bucket) (ResourceStatus, error) {
	external, err := in.client.GetBucketNotificationConfiguration(ctx, &awss3.GetBucketNotificationConfigurationInput{Bucket: awsclient.String(meta.GetExternalName(bucket))})
	if err != nil {
		return NeedsUpdate, awsclient.Wrap(err, notificationGetFailed)
	}

	config := bucket.Spec.ForProvider.NotificationConfiguration
	status := bucketStatus(config, external)
	switch status { // nolint:exhaustive
	case Updated, NeedsDeletion:
		return status, nil
	}

	generated := GenerateConfiguration(config)

	if cmp.Equal(external.LambdaFunctionConfigurations, generated.LambdaFunctionConfigurations) &&
		cmp.Equal(external.QueueConfigurations, generated.QueueConfigurations) &&
		cmp.Equal(external.TopicConfigurations, generated.TopicConfigurations) {
		return Updated, nil
	}

	return NeedsUpdate, nil
}

func copyEvents(src []string) []awss3types.Event {
	if len(src) == 0 {
		return nil
	}
	out := make([]awss3types.Event, len(src))
	for i, v := range src {
		cast := awss3types.Event(v)
		out[i] = cast
	}
	return out
}

func generateFilter(src *v1beta1.NotificationConfigurationFilter) *awss3types.NotificationConfigurationFilter {
	if src == nil || src.Key == nil {
		return nil
	}
	out := &awss3types.NotificationConfigurationFilter{Key: &awss3types.S3KeyFilter{}}
	if src.Key.FilterRules == nil {
		return out
	}
	out.Key.FilterRules = make([]awss3types.FilterRule, len(src.Key.FilterRules))
	for i, v := range src.Key.FilterRules {
		out.Key.FilterRules[i] = awss3types.FilterRule{
			Name:  awss3types.FilterRuleName(v.Name),
			Value: v.Value,
		}
	}
	return out
}

// GenerateLambdaConfiguration creates []awss3types.LambdaFunctionConfiguration from the local NotificationConfiguration
func GenerateLambdaConfiguration(config *v1beta1.NotificationConfiguration) []awss3types.LambdaFunctionConfiguration {
	// NOTE(muvaf): We skip prealloc because the behavior of AWS SDK differs when
	// the array is 0 element vs nil.
	var configurations []awss3types.LambdaFunctionConfiguration // nolint:prealloc
	for _, v := range config.LambdaFunctionConfigurations {
		conf := awss3types.LambdaFunctionConfiguration{
			Filter:            nil,
			Id:                v.ID,
			LambdaFunctionArn: awsclient.String(v.LambdaFunctionArn),
		}
		if v.Events != nil {
			conf.Events = copyEvents(v.Events)
		}
		if v.Filter != nil {
			conf.Filter = generateFilter(v.Filter)
		}
		configurations = append(configurations, conf)
	}
	return configurations
}

// GenerateTopicConfigurations creates []awss3types.TopicConfiguration from the local NotificationConfiguration
func GenerateTopicConfigurations(config *v1beta1.NotificationConfiguration) []awss3types.TopicConfiguration {
	// NOTE(muvaf): We skip prealloc because the behavior of AWS SDK differs when
	// the array is 0 element vs nil.
	var configurations []awss3types.TopicConfiguration // nolint:prealloc
	for _, v := range config.TopicConfigurations {
		conf := awss3types.TopicConfiguration{
			Id:       v.ID,
			TopicArn: v.TopicArn,
		}
		if v.Events != nil {
			conf.Events = copyEvents(v.Events)
		}
		if v.Filter != nil {
			conf.Filter = generateFilter(v.Filter)
		}
		configurations = append(configurations, conf)
	}
	return configurations
}

// GenerateQueueConfigurations creates []awss3types.QueueConfiguration from the local NotificationConfiguration
func GenerateQueueConfigurations(config *v1beta1.NotificationConfiguration) []awss3types.QueueConfiguration {
	// NOTE(muvaf): We skip prealloc because the behavior of AWS SDK differs when
	// the array is 0 element vs nil.
	var configurations []awss3types.QueueConfiguration // nolint:prealloc
	for _, v := range config.QueueConfigurations {
		conf := awss3types.QueueConfiguration{
			Id:       v.ID,
			QueueArn: awsclient.String(v.QueueArn),
		}
		if v.Events != nil {
			conf.Events = copyEvents(v.Events)
		}
		if v.Filter != nil {
			conf.Filter = generateFilter(v.Filter)
		}
		configurations = append(configurations, conf)
	}
	return configurations
}

// GenerateConfiguration creates the external aws NotificationConfiguration from the local representation
func GenerateConfiguration(config *v1beta1.NotificationConfiguration) *awss3types.NotificationConfiguration {
	return &awss3types.NotificationConfiguration{
		LambdaFunctionConfigurations: GenerateLambdaConfiguration(config),
		QueueConfigurations:          GenerateQueueConfigurations(config),
		TopicConfigurations:          GenerateTopicConfigurations(config),
	}
}

// GenerateNotificationConfigurationInput creates the input for the LifecycleConfiguration request for the S3 Client
func GenerateNotificationConfigurationInput(name string, config *v1beta1.NotificationConfiguration) *awss3.PutBucketNotificationConfigurationInput {
	return &awss3.PutBucketNotificationConfigurationInput{
		Bucket:                    awsclient.String(name),
		NotificationConfiguration: GenerateConfiguration(config),
	}
}

// CreateOrUpdate sends a request to have resource created on AWS
func (in *NotificationConfigurationClient) CreateOrUpdate(ctx context.Context, bucket *v1beta1.Bucket) error {
	if bucket.Spec.ForProvider.NotificationConfiguration == nil {
		return nil
	}
	input := GenerateNotificationConfigurationInput(meta.GetExternalName(bucket), bucket.Spec.ForProvider.NotificationConfiguration)
	_, err := in.client.PutBucketNotificationConfiguration(ctx, input)
	return awsclient.Wrap(err, notificationPutFailed)
}

// Delete does nothing because there is no corresponding deletion call in awsclient.
func (*NotificationConfigurationClient) Delete(_ context.Context, _ *v1beta1.Bucket) error {
	return nil
}
