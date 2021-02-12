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

package routeresponse

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an RouteResponse resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create RouteResponse in AWS"
	errUpdate        = "cannot update RouteResponse in AWS"
	errDescribe      = "failed to describe RouteResponse"
	errDelete        = "failed to delete RouteResponse"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.RouteResponse)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.RouteResponse)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetRouteResponseInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetRouteResponseWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateRouteResponse(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.RouteResponse)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateRouteResponseInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateRouteResponseWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.RouteResponseId != nil {
		cr.Status.AtProvider.RouteResponseID = resp.RouteResponseId
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.RouteResponse)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateRouteResponseInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateRouteResponseWithContext(ctx, input)
	if err != nil {
		return managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate)
	}
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, err)
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.RouteResponse)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteRouteResponseInput(cr)
	if err := e.preDelete(ctx, cr, input); err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	_, err := e.client.DeleteRouteResponseWithContext(ctx, input)
	return awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete)
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ApiGatewayV2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.ApiGatewayV2API
	preObserve     func(context.Context, *svcapitypes.RouteResponse, *svcsdk.GetRouteResponseInput) error
	postObserve    func(context.Context, *svcapitypes.RouteResponse, *svcsdk.GetRouteResponseOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.RouteResponseParameters, *svcsdk.GetRouteResponseOutput) error
	isUpToDate     func(*svcapitypes.RouteResponse, *svcsdk.GetRouteResponseOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.RouteResponse, *svcsdk.CreateRouteResponseInput) error
	postCreate     func(context.Context, *svcapitypes.RouteResponse, *svcsdk.CreateRouteResponseOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.RouteResponse, *svcsdk.DeleteRouteResponseInput) error
	preUpdate      func(context.Context, *svcapitypes.RouteResponse, *svcsdk.UpdateRouteResponseInput) error
	postUpdate     func(context.Context, *svcapitypes.RouteResponse, *svcsdk.UpdateRouteResponseOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.RouteResponse, *svcsdk.GetRouteResponseInput) error {
	return nil
}
func nopPostObserve(context.Context, *svcapitypes.RouteResponse, *svcsdk.GetRouteResponseOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error) {
	return managed.ExternalObservation{}, nil
}
func nopLateInitialize(*svcapitypes.RouteResponseParameters, *svcsdk.GetRouteResponseOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.RouteResponse, *svcsdk.GetRouteResponseOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.RouteResponse, *svcsdk.CreateRouteResponseInput) error {
	return nil
}
func nopPostCreate(context.Context, *svcapitypes.RouteResponse, *svcsdk.CreateRouteResponseOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error) {
	return managed.ExternalCreation{}, nil
}
func nopPreDelete(context.Context, *svcapitypes.RouteResponse, *svcsdk.DeleteRouteResponseInput) error {
	return nil
}
func nopPreUpdate(context.Context, *svcapitypes.RouteResponse, *svcsdk.UpdateRouteResponseInput) error {
	return nil
}
func nopPostUpdate(context.Context, *svcapitypes.RouteResponse, *svcsdk.UpdateRouteResponseOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
