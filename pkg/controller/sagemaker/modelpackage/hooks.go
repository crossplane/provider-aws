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

package modelpackage

import (
	"context"

	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	ctrl "sigs.k8s.io/controller-runtime"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/sagemaker/v1alpha1"
)

// SetupModelPackage adds a controller that reconciles ModelPackage.
func SetupModelPackage(mgr ctrl.Manager, l logging.Logger) error {
	name := managed.ControllerName(svcapitypes.ModelPackageGroupKind)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&svcapitypes.ModelPackage{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.ModelPackageGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func (*external) preObserve(context.Context, *svcapitypes.ModelPackage) error {
	return nil
}
func (*external) postObserve(_ context.Context, cr *svcapitypes.ModelPackage, _ *svcsdk.DescribeModelPackageOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	cr.SetConditions(xpv1.Available())
	return obs, err
}

func (*external) preCreate(context.Context, *svcapitypes.ModelPackage) error {
	return nil
}

func (*external) postCreate(_ context.Context, _ *svcapitypes.ModelPackage, _ *svcsdk.CreateModelPackageOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}

func (*external) preUpdate(context.Context, *svcapitypes.ModelPackage) error {
	return nil
}

func (*external) postUpdate(_ context.Context, _ *svcapitypes.ModelPackage, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
func lateInitialize(*svcapitypes.ModelPackageParameters, *svcsdk.DescribeModelPackageOutput) error {
	return nil
}

func isUpToDate(*svcapitypes.ModelPackage, *svcsdk.DescribeModelPackageOutput) bool {
	return true
}

func preGenerateDescribeModelPackageInput(_ *svcapitypes.ModelPackage, obj *svcsdk.DescribeModelPackageInput) *svcsdk.DescribeModelPackageInput {
	return obj
}

func postGenerateDescribeModelPackageInput(_ *svcapitypes.ModelPackage, obj *svcsdk.DescribeModelPackageInput) *svcsdk.DescribeModelPackageInput {
	return obj
}

func preGenerateCreateModelPackageInput(_ *svcapitypes.ModelPackage, obj *svcsdk.CreateModelPackageInput) *svcsdk.CreateModelPackageInput {
	return obj
}

func postGenerateCreateModelPackageInput(_ *svcapitypes.ModelPackage, obj *svcsdk.CreateModelPackageInput) *svcsdk.CreateModelPackageInput {
	return obj
}
func preGenerateDeleteModelPackageInput(_ *svcapitypes.ModelPackage, obj *svcsdk.DeleteModelPackageInput) *svcsdk.DeleteModelPackageInput {
	return obj
}

func postGenerateDeleteModelPackageInput(_ *svcapitypes.ModelPackage, obj *svcsdk.DeleteModelPackageInput) *svcsdk.DeleteModelPackageInput {
	return obj
}