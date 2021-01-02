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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane/provider-aws/apis/ec2/v1alpha1"
)

// SecurityGroupName returns the spec.groupName of a SecurityGroup.
func SecurityGroupName() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		sg, ok := mg.(*SecurityGroup)
		if !ok {
			return ""
		}
		return sg.Spec.ForProvider.GroupName
	}
}

// ResolveReferences of this InternetGateway
func (mg *InternetGateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.vpcId
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: *mg.Spec.ForProvider.VPCID,
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To:           reference.To{Managed: &VPC{}, List: &VPCList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.vpcId")
	}
	mg.Spec.ForProvider.VPCID = aws.String(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroup
func (mg *SecurityGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.vpcId
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To:           reference.To{Managed: &VPC{}, List: &VPCList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.vpcId")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	for i, ing := range mg.Spec.ForProvider.Ingress {
		for j, pair := range ing.UserIDGroupPairs {
			// Resolve spec.forProvider.ingress[*].userIdGroupPairs[*]
			rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(pair.VPCID),
				Reference:    pair.VPCIDRef,
				Selector:     pair.VPCIDSelector,
				To:           reference.To{Managed: &VPC{}, List: &VPCList{}},
				Extract:      reference.ExternalName(),
			})
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("spec.forProvider.ingress[%d].userIdGroupPairs[%d]", i, j))
			}
			mg.Spec.ForProvider.Ingress[i].UserIDGroupPairs[j].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Ingress[i].UserIDGroupPairs[j].VPCIDRef = rsp.ResolvedReference
		}
	}

	return nil
}

// ResolveReferences of this Subnet
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.vpcId
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: *mg.Spec.ForProvider.VPCID,
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To:           reference.To{Managed: &VPC{}, List: &VPCList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.vpcId")
	}
	mg.Spec.ForProvider.VPCID = aws.String(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouteTable
func (mg *RouteTable) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.vpcId
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: aws.StringValue(mg.Spec.ForProvider.VPCID),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To:           reference.To{Managed: &VPC{}, List: &VPCList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.vpcId")
	}
	mg.Spec.ForProvider.VPCID = aws.String(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	// Resolve spec.forProvider.routes[].gatewayId
	for i := range mg.Spec.ForProvider.Routes {
		rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: aws.StringValue(mg.Spec.ForProvider.Routes[i].GatewayID),
			Reference:    mg.Spec.ForProvider.Routes[i].GatewayIDRef,
			Selector:     mg.Spec.ForProvider.Routes[i].GatewayIDSelector,
			To:           reference.To{Managed: &InternetGateway{}, List: &InternetGatewayList{}},
			Extract:      reference.ExternalName(),
		})
		if err != nil {
			return errors.Wrapf(err, "spec.forProvider.routes[%d].gatewayId", i)
		}
		mg.Spec.ForProvider.Routes[i].GatewayID = aws.String(rsp.ResolvedValue)
		mg.Spec.ForProvider.Routes[i].GatewayIDRef = rsp.ResolvedReference
	}

	// Resolve spec.forProvider.routes[].natGatewayId
	for i := range mg.Spec.ForProvider.Routes {
		rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: aws.StringValue(mg.Spec.ForProvider.Routes[i].NatGatewayID),
			Reference:    mg.Spec.ForProvider.Routes[i].NatGatewayIDRef,
			Selector:     mg.Spec.ForProvider.Routes[i].NatGatewayIDSelector,
			To:           reference.To{Managed: &NATGateway{}, List: &NATGatewayList{}},
			Extract:      reference.ExternalName(),
		})
		if err != nil {
			return errors.Wrapf(err, "spec.forProvider.routes[%d].natGatewayId", i)
		}
		mg.Spec.ForProvider.Routes[i].NatGatewayID = aws.String(rsp.ResolvedValue)
		mg.Spec.ForProvider.Routes[i].NatGatewayIDRef = rsp.ResolvedReference
	}

	// Resolve spec.associations[].subnetId
	for i := range mg.Spec.ForProvider.Associations {
		rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: aws.StringValue(mg.Spec.ForProvider.Associations[i].SubnetID),
			Reference:    mg.Spec.ForProvider.Associations[i].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Associations[i].SubnetIDSelector,
			To:           reference.To{Managed: &Subnet{}, List: &SubnetList{}},
			Extract:      reference.ExternalName(),
		})
		if err != nil {
			return errors.Wrapf(err, "spec.forProvider.associations[%d].subnetId", i)
		}
		mg.Spec.ForProvider.Associations[i].SubnetID = aws.String(rsp.ResolvedValue)
		mg.Spec.ForProvider.Associations[i].SubnetIDRef = rsp.ResolvedReference
	}

	return nil
}

// ResolveReferences of this NatGateway
func (mg *NATGateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// // Resolve spec.subnetId
	subnetIDResponse, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: aws.StringValue(mg.Spec.ForProvider.SubnetID),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To:           reference.To{Managed: &Subnet{}, List: &SubnetList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return err
	}
	mg.Spec.ForProvider.SubnetID = aws.String(subnetIDResponse.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = subnetIDResponse.ResolvedReference

	// // Resolve spec.elasticIp
	AllocationIDRespone, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: aws.StringValue(mg.Spec.ForProvider.AllocationID),
		Reference:    mg.Spec.ForProvider.AllocationIDRef,
		Selector:     mg.Spec.ForProvider.AllocationIDSelector,
		To:           reference.To{Managed: &v1alpha1.ElasticIP{}, List: &v1alpha1.ElasticIPList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return err
	}
	mg.Spec.ForProvider.AllocationID = aws.String(AllocationIDRespone.ResolvedValue)
	mg.Spec.ForProvider.AllocationIDRef = AllocationIDRespone.ResolvedReference

	return nil
}
