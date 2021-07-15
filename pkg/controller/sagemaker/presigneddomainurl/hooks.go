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

package presigneddomainurl

import (
	"context"

	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/sagemaker/v1alpha1"
)

// SetupPresignedDomainURL adds a controller that reconciles PresignedDomainURL.
func SetupPresignedDomainURL(mgr ctrl.Manager, l logging.Logger) error {
	name := managed.ControllerName(svcapitypes.PresignedDomainURLGroupKind)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&svcapitypes.PresignedDomainURL{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.PresignedDomainURLGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	// TODO: implement me!
	return managed.ExternalObservation{}, nil
}

func (*external) preCreate(context.Context, *svcapitypes.PresignedDomainURL) error {
	return nil
}

func (*external) postCreate(_ context.Context, _ *svcapitypes.PresignedDomainURL, _ *svcsdk.CreatePresignedDomainUrlOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}

func (*external) preUpdate(context.Context, *svcapitypes.PresignedDomainURL) error {
	return nil
}

func (*external) postUpdate(_ context.Context, _ *svcapitypes.PresignedDomainURL, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}

func preGenerateCreatePresignedDomainUrlInput(_ *svcapitypes.PresignedDomainURL, obj *svcsdk.CreatePresignedDomainUrlInput) *svcsdk.CreatePresignedDomainUrlInput {
	return obj
}

func postGenerateCreatePresignedDomainUrlInput(_ *svcapitypes.PresignedDomainURL, obj *svcsdk.CreatePresignedDomainUrlInput) *svcsdk.CreatePresignedDomainUrlInput {
	return obj
}
func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	// TODO: implement me!
	return nil
}