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

package api

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an API resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create API in AWS"
	errDescribe      = "failed to describe API"
	errDelete        = "failed to delete API"
)

type connector struct {
	kube client.Client
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, err
	}
	return &external{client: svcapi.New(sess), kube: c.kube}, errors.Wrap(err, errCreateSession)
}

type external struct {
	kube   client.Client
	client svcsdkapi.ApiGatewayV2API
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if err := e.preObserve(ctx, cr); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetApiInput(cr)
	resp, err := e.client.GetApiWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errors.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	lateInitialize(&cr.Spec.ForProvider, resp)
	GenerateAPI(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        true,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(runtimev1alpha1.Creating())
	if err := e.preCreate(ctx, cr); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	input := GenerateCreateApiInput(cr)
	resp, err := e.client.CreateApiWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreate)
	}

	if resp.ApiEndpoint != nil {
		cr.Status.AtProvider.APIEndpoint = resp.ApiEndpoint
	}
	if resp.ApiGatewayManaged != nil {
		cr.Status.AtProvider.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.ApiId != nil {
		cr.Status.AtProvider.APIID = resp.ApiId
	}
	if resp.CreatedDate != nil {
		cr.Status.AtProvider.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.ImportInfo != nil {
		f9 := []*string{}
		for _, f9iter := range resp.ImportInfo {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		cr.Status.AtProvider.ImportInfo = f9
	}
	if resp.Warnings != nil {
		f15 := []*string{}
		for _, f15iter := range resp.Warnings {
			var f15elem string
			f15elem = *f15iter
			f15 = append(f15, &f15elem)
		}
		cr.Status.AtProvider.Warnings = f15
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	if err := e.preUpdate(ctx, cr); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	return e.postUpdate(ctx, cr, managed.ExternalUpdate{}, nil)
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(runtimev1alpha1.Deleting())
	input := GenerateDeleteApiInput(cr)
	_, err := e.client.DeleteApiWithContext(ctx, input)
	return errors.Wrap(cpresource.Ignore(IsNotFound, err), errDelete)
}
