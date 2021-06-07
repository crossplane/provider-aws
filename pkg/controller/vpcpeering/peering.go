package vpcpeering

import (
	"context"
	"github.com/crossplane/provider-aws/pkg/clients/peering"

	"github.com/aws/aws-sdk-go/aws"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/vpcpeering/v1alpha1"
	awsclients "github.com/crossplane/provider-aws/pkg/clients"

	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	svcapi "github.com/aws/aws-sdk-go/service/ec2"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	awsclient "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

type VPCPeeringConnectionStateReasonCode string

const (
	VPCPeeringConnectionStateReasonCode_initiating_request VPCPeeringConnectionStateReasonCode = "initiating-request"
	VPCPeeringConnectionStateReasonCode_pending_acceptance VPCPeeringConnectionStateReasonCode = "pending-acceptance"
	VPCPeeringConnectionStateReasonCode_active             VPCPeeringConnectionStateReasonCode = "active"
	VPCPeeringConnectionStateReasonCode_deleted            VPCPeeringConnectionStateReasonCode = "deleted"
	VPCPeeringConnectionStateReasonCode_rejected           VPCPeeringConnectionStateReasonCode = "rejected"
	VPCPeeringConnectionStateReasonCode_failed             VPCPeeringConnectionStateReasonCode = "failed"
	VPCPeeringConnectionStateReasonCode_expired            VPCPeeringConnectionStateReasonCode = "expired"
	VPCPeeringConnectionStateReasonCode_provisioning       VPCPeeringConnectionStateReasonCode = "provisioning"
	VPCPeeringConnectionStateReasonCode_deleting           VPCPeeringConnectionStateReasonCode = "deleting"
)

const (
	errUnexpectedObject = "managed resource is not an VPCPeeringConnection resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create VPCPeeringConnection in AWS"
	errUpdate        = "cannot update VPCPeeringConnection in AWS"
	errDescribe      = "failed to describe VPCPeeringConnection"
	errDelete        = "failed to delete VPCPeeringConnection"
)

// SetupVPCPeeringConnection adds a controller that reconciles VPCPeeringConnection.
func SetupVPCPeeringConnection(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(svcapitypes.VPCPeeringConnectionGroupKind)
	opts := []option{
		func(e *external) {
			c := &custom{client: e.client, kube: e.kube}
			e.postObserve = c.postObserve
			e.postCreate = c.postCreate
			e.preCreate = preCreate
			e.isUpToDate = c.isUpToDate
			e.filterList = filterList
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&svcapitypes.VPCPeeringConnection{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.VPCPeeringConnectionGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func filterList(cr *svcapitypes.VPCPeeringConnection, obj *svcsdk.DescribeVpcPeeringConnectionsOutput) *svcsdk.DescribeVpcPeeringConnectionsOutput {
	connectionIdentifier := aws.String(meta.GetExternalName(cr))
	resp := &svcsdk.DescribeVpcPeeringConnectionsOutput{}
	for _, vpcPeeringConnection := range obj.VpcPeeringConnections {
		if aws.StringValue(vpcPeeringConnection.VpcPeeringConnectionId) == aws.StringValue(connectionIdentifier) {
			resp.VpcPeeringConnections = append(resp.VpcPeeringConnections, vpcPeeringConnection)
			break
		}
	}
	return resp
}

type custom struct {
	kube   client.Client
	client svcsdkapi.EC2API
}

func (e *custom) postObserve(_ context.Context, cr *svcapitypes.VPCPeeringConnection, obj *svcsdk.DescribeVpcPeeringConnectionsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}

	if awsclients.StringValue(obj.VpcPeeringConnections[0].Status.Code) == "pending-acceptance" && cr.Spec.ForProvider.AcceptRequest {
		req := svcsdk.AcceptVpcPeeringConnectionInput{
			VpcPeeringConnectionId: awsclients.String(*obj.VpcPeeringConnections[0].VpcPeeringConnectionId),
		}
		request, _ := e.client.AcceptVpcPeeringConnectionRequest(&req)
		err := request.Send()
		if err != nil {
			return obs, err
		}
	}

	available := setCondition(obj.VpcPeeringConnections[0].Status, cr)
	if !available {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	return obs, nil
}

func setCondition(code *svcsdk.VpcPeeringConnectionStateReason, cr *svcapitypes.VPCPeeringConnection) bool {
	switch aws.StringValue(code.Code) {
	case string(VPCPeeringConnectionStateReasonCode_pending_acceptance):
		cr.SetConditions(xpv1.Creating())
		return true
	case string(VPCPeeringConnectionStateReasonCode_deleted):
		cr.SetConditions(xpv1.Unavailable())
		return false
	case string(VPCPeeringConnectionStateReasonCode_active):
		cr.SetConditions(xpv1.Available())
		return true
	}
	return false
}

func (e *custom) isUpToDate(cr *svcapitypes.VPCPeeringConnection, obj *svcsdk.DescribeVpcPeeringConnectionsOutput) (bool, error) {
	return true, nil
}

func preCreate(ctx context.Context, cr *svcapitypes.VPCPeeringConnection, obj *svcsdk.CreateVpcPeeringConnectionInput) error {
	// set external name as tag on the vpc peering connection
	resType := "vpc-peering-connection"
	key := "crossplane-claim-name"
	value := cr.ObjectMeta.Name

	spec := svcsdk.TagSpecification{
		ResourceType: &resType,
		Tags: []*svcsdk.Tag{
			{
				Key:   &key,
				Value: &value,
			},
		},
	}
	obj.TagSpecifications = append(obj.TagSpecifications, &spec)

	return nil
}

func (e *custom) postCreate(ctx context.Context, cr *svcapitypes.VPCPeeringConnection, obj *svcsdk.CreateVpcPeeringConnectionOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	if err != nil {
		return managed.ExternalCreation{}, err
	}
	// set peering connection id as external name annotation on k8s object after creation

	meta.SetExternalName(cr, aws.StringValue(obj.VpcPeeringConnection.VpcPeeringConnectionId))
	cre.ExternalNameAssigned = true
	return cre, nil
}

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.VPCPeeringConnection)
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
	cr, ok := mg.(*svcapitypes.VPCPeeringConnection)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := peering.GenerateDescribeVpcPeeringConnectionsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeVpcPeeringConnectionsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.VpcPeeringConnections) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	peering.GenerateVPCPeeringConnection(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.VPCPeeringConnection)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := peering.GenerateCreateVpcPeeringConnectionInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateVpcPeeringConnectionWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.VpcPeeringConnection.AccepterVpcInfo != nil {
		f0 := &svcapitypes.VPCPeeringConnectionVPCInfo{}
		if resp.VpcPeeringConnection.AccepterVpcInfo.CidrBlock != nil {
			f0.CIDRBlock = resp.VpcPeeringConnection.AccepterVpcInfo.CidrBlock
		}
		if resp.VpcPeeringConnection.AccepterVpcInfo.CidrBlockSet != nil {
			f0f1 := []*svcapitypes.CIDRBlock{}
			for _, f0f1iter := range resp.VpcPeeringConnection.AccepterVpcInfo.CidrBlockSet {
				f0f1elem := &svcapitypes.CIDRBlock{}
				if f0f1iter.CidrBlock != nil {
					f0f1elem.CIDRBlock = f0f1iter.CidrBlock
				}
				f0f1 = append(f0f1, f0f1elem)
			}
			f0.CIDRBlockSet = f0f1
		}
		if resp.VpcPeeringConnection.AccepterVpcInfo.Ipv6CidrBlockSet != nil {
			f0f2 := []*svcapitypes.IPv6CIDRBlock{}
			for _, f0f2iter := range resp.VpcPeeringConnection.AccepterVpcInfo.Ipv6CidrBlockSet {
				f0f2elem := &svcapitypes.IPv6CIDRBlock{}
				if f0f2iter.Ipv6CidrBlock != nil {
					f0f2elem.IPv6CIDRBlock = f0f2iter.Ipv6CidrBlock
				}
				f0f2 = append(f0f2, f0f2elem)
			}
			f0.IPv6CIDRBlockSet = f0f2
		}
		if resp.VpcPeeringConnection.AccepterVpcInfo.OwnerId != nil {
			f0.OwnerID = resp.VpcPeeringConnection.AccepterVpcInfo.OwnerId
		}
		if resp.VpcPeeringConnection.AccepterVpcInfo.PeeringOptions != nil {
			f0f4 := &svcapitypes.VPCPeeringConnectionOptionsDescription{}
			if resp.VpcPeeringConnection.AccepterVpcInfo.PeeringOptions.AllowDnsResolutionFromRemoteVpc != nil {
				f0f4.AllowDNSResolutionFromRemoteVPC = resp.VpcPeeringConnection.AccepterVpcInfo.PeeringOptions.AllowDnsResolutionFromRemoteVpc
			}
			if resp.VpcPeeringConnection.AccepterVpcInfo.PeeringOptions.AllowEgressFromLocalClassicLinkToRemoteVpc != nil {
				f0f4.AllowEgressFromLocalClassicLinkToRemoteVPC = resp.VpcPeeringConnection.AccepterVpcInfo.PeeringOptions.AllowEgressFromLocalClassicLinkToRemoteVpc
			}
			if resp.VpcPeeringConnection.AccepterVpcInfo.PeeringOptions.AllowEgressFromLocalVpcToRemoteClassicLink != nil {
				f0f4.AllowEgressFromLocalVPCToRemoteClassicLink = resp.VpcPeeringConnection.AccepterVpcInfo.PeeringOptions.AllowEgressFromLocalVpcToRemoteClassicLink
			}
			f0.PeeringOptions = f0f4
		}
		if resp.VpcPeeringConnection.AccepterVpcInfo.Region != nil {
			f0.Region = resp.VpcPeeringConnection.AccepterVpcInfo.Region
		}
		if resp.VpcPeeringConnection.AccepterVpcInfo.VpcId != nil {
			f0.VPCID = resp.VpcPeeringConnection.AccepterVpcInfo.VpcId
		}
		cr.Status.AtProvider.AccepterVPCInfo = f0
	} else {
		cr.Status.AtProvider.AccepterVPCInfo = nil
	}
	if resp.VpcPeeringConnection.ExpirationTime != nil {
		cr.Status.AtProvider.ExpirationTime = &metav1.Time{*resp.VpcPeeringConnection.ExpirationTime}
	} else {
		cr.Status.AtProvider.ExpirationTime = nil
	}
	if resp.VpcPeeringConnection.RequesterVpcInfo != nil {
		f2 := &svcapitypes.VPCPeeringConnectionVPCInfo{}
		if resp.VpcPeeringConnection.RequesterVpcInfo.CidrBlock != nil {
			f2.CIDRBlock = resp.VpcPeeringConnection.RequesterVpcInfo.CidrBlock
		}
		if resp.VpcPeeringConnection.RequesterVpcInfo.CidrBlockSet != nil {
			f2f1 := []*svcapitypes.CIDRBlock{}
			for _, f2f1iter := range resp.VpcPeeringConnection.RequesterVpcInfo.CidrBlockSet {
				f2f1elem := &svcapitypes.CIDRBlock{}
				if f2f1iter.CidrBlock != nil {
					f2f1elem.CIDRBlock = f2f1iter.CidrBlock
				}
				f2f1 = append(f2f1, f2f1elem)
			}
			f2.CIDRBlockSet = f2f1
		}
		if resp.VpcPeeringConnection.RequesterVpcInfo.Ipv6CidrBlockSet != nil {
			f2f2 := []*svcapitypes.IPv6CIDRBlock{}
			for _, f2f2iter := range resp.VpcPeeringConnection.RequesterVpcInfo.Ipv6CidrBlockSet {
				f2f2elem := &svcapitypes.IPv6CIDRBlock{}
				if f2f2iter.Ipv6CidrBlock != nil {
					f2f2elem.IPv6CIDRBlock = f2f2iter.Ipv6CidrBlock
				}
				f2f2 = append(f2f2, f2f2elem)
			}
			f2.IPv6CIDRBlockSet = f2f2
		}
		if resp.VpcPeeringConnection.RequesterVpcInfo.OwnerId != nil {
			f2.OwnerID = resp.VpcPeeringConnection.RequesterVpcInfo.OwnerId
		}
		if resp.VpcPeeringConnection.RequesterVpcInfo.PeeringOptions != nil {
			f2f4 := &svcapitypes.VPCPeeringConnectionOptionsDescription{}
			if resp.VpcPeeringConnection.RequesterVpcInfo.PeeringOptions.AllowDnsResolutionFromRemoteVpc != nil {
				f2f4.AllowDNSResolutionFromRemoteVPC = resp.VpcPeeringConnection.RequesterVpcInfo.PeeringOptions.AllowDnsResolutionFromRemoteVpc
			}
			if resp.VpcPeeringConnection.RequesterVpcInfo.PeeringOptions.AllowEgressFromLocalClassicLinkToRemoteVpc != nil {
				f2f4.AllowEgressFromLocalClassicLinkToRemoteVPC = resp.VpcPeeringConnection.RequesterVpcInfo.PeeringOptions.AllowEgressFromLocalClassicLinkToRemoteVpc
			}
			if resp.VpcPeeringConnection.RequesterVpcInfo.PeeringOptions.AllowEgressFromLocalVpcToRemoteClassicLink != nil {
				f2f4.AllowEgressFromLocalVPCToRemoteClassicLink = resp.VpcPeeringConnection.RequesterVpcInfo.PeeringOptions.AllowEgressFromLocalVpcToRemoteClassicLink
			}
			f2.PeeringOptions = f2f4
		}
		if resp.VpcPeeringConnection.RequesterVpcInfo.Region != nil {
			f2.Region = resp.VpcPeeringConnection.RequesterVpcInfo.Region
		}
		if resp.VpcPeeringConnection.RequesterVpcInfo.VpcId != nil {
			f2.VPCID = resp.VpcPeeringConnection.RequesterVpcInfo.VpcId
		}
		cr.Status.AtProvider.RequesterVPCInfo = f2
	} else {
		cr.Status.AtProvider.RequesterVPCInfo = nil
	}
	if resp.VpcPeeringConnection.Status != nil {
		f3 := &svcapitypes.VPCPeeringConnectionStateReason{}
		if resp.VpcPeeringConnection.Status.Code != nil {
			f3.Code = resp.VpcPeeringConnection.Status.Code
		}
		if resp.VpcPeeringConnection.Status.Message != nil {
			f3.Message = resp.VpcPeeringConnection.Status.Message
		}
		cr.Status.AtProvider.Status = f3
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.VpcPeeringConnection.Tags != nil {
		f4 := []*svcapitypes.Tag{}
		for _, f4iter := range resp.VpcPeeringConnection.Tags {
			f4elem := &svcapitypes.Tag{}
			if f4iter.Key != nil {
				f4elem.Key = f4iter.Key
			}
			if f4iter.Value != nil {
				f4elem.Value = f4iter.Value
			}
			f4 = append(f4, f4elem)
		}
		cr.Status.AtProvider.Tags = f4
	} else {
		cr.Status.AtProvider.Tags = nil
	}
	if resp.VpcPeeringConnection.VpcPeeringConnectionId != nil {
		cr.Status.AtProvider.VPCPeeringConnectionID = resp.VpcPeeringConnection.VpcPeeringConnectionId
	} else {
		cr.Status.AtProvider.VPCPeeringConnectionID = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	return e.update(ctx, mg)

}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.VPCPeeringConnection)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := peering.GenerateDeleteVpcPeeringConnectionInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteVpcPeeringConnectionWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EC2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		update:         nopUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.EC2API
	preObserve     func(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.DescribeVpcPeeringConnectionsInput) error
	postObserve    func(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.DescribeVpcPeeringConnectionsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.VPCPeeringConnection, *svcsdk.DescribeVpcPeeringConnectionsOutput) *svcsdk.DescribeVpcPeeringConnectionsOutput
	lateInitialize func(*svcapitypes.VPCPeeringConnectionParameters, *svcsdk.DescribeVpcPeeringConnectionsOutput) error
	isUpToDate     func(*svcapitypes.VPCPeeringConnection, *svcsdk.DescribeVpcPeeringConnectionsOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.CreateVpcPeeringConnectionInput) error
	postCreate     func(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.CreateVpcPeeringConnectionOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.DeleteVpcPeeringConnectionInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.DeleteVpcPeeringConnectionOutput, error) error
	update         func(context.Context, cpresource.Managed) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.DescribeVpcPeeringConnectionsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.VPCPeeringConnection, _ *svcsdk.DescribeVpcPeeringConnectionsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.VPCPeeringConnection, list *svcsdk.DescribeVpcPeeringConnectionsOutput) *svcsdk.DescribeVpcPeeringConnectionsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.VPCPeeringConnectionParameters, *svcsdk.DescribeVpcPeeringConnectionsOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.VPCPeeringConnection, *svcsdk.DescribeVpcPeeringConnectionsOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.CreateVpcPeeringConnectionInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.VPCPeeringConnection, _ *svcsdk.CreateVpcPeeringConnectionOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.VPCPeeringConnection, *svcsdk.DeleteVpcPeeringConnectionInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.VPCPeeringConnection, _ *svcsdk.DeleteVpcPeeringConnectionOutput, err error) error {
	return err
}
func nopUpdate(context.Context, cpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
