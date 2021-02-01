// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap
type GtmCidrmap struct {
	pulumi.CustomResourceState

	Assignments       GtmCidrmapAssignmentArrayOutput   `pulumi:"assignments"`
	DefaultDatacenter GtmCidrmapDefaultDatacenterOutput `pulumi:"defaultDatacenter"`
	Domain            pulumi.StringOutput               `pulumi:"domain"`
	Name              pulumi.StringOutput               `pulumi:"name"`
	WaitOnComplete    pulumi.BoolPtrOutput              `pulumi:"waitOnComplete"`
}

// NewGtmCidrmap registers a new resource with the given unique name, arguments, and options.
func NewGtmCidrmap(ctx *pulumi.Context,
	name string, args *GtmCidrmapArgs, opts ...pulumi.ResourceOption) (*GtmCidrmap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultDatacenter == nil {
		return nil, errors.New("invalid value for required argument 'DefaultDatacenter'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	var resource GtmCidrmap
	err := ctx.RegisterResource("akamai:trafficmanagement/gtmCidrmap:GtmCidrmap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGtmCidrmap gets an existing GtmCidrmap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGtmCidrmap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GtmCidrmapState, opts ...pulumi.ResourceOption) (*GtmCidrmap, error) {
	var resource GtmCidrmap
	err := ctx.ReadResource("akamai:trafficmanagement/gtmCidrmap:GtmCidrmap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GtmCidrmap resources.
type gtmCidrmapState struct {
	Assignments       []GtmCidrmapAssignment       `pulumi:"assignments"`
	DefaultDatacenter *GtmCidrmapDefaultDatacenter `pulumi:"defaultDatacenter"`
	Domain            *string                      `pulumi:"domain"`
	Name              *string                      `pulumi:"name"`
	WaitOnComplete    *bool                        `pulumi:"waitOnComplete"`
}

type GtmCidrmapState struct {
	Assignments       GtmCidrmapAssignmentArrayInput
	DefaultDatacenter GtmCidrmapDefaultDatacenterPtrInput
	Domain            pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	WaitOnComplete    pulumi.BoolPtrInput
}

func (GtmCidrmapState) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmCidrmapState)(nil)).Elem()
}

type gtmCidrmapArgs struct {
	Assignments       []GtmCidrmapAssignment      `pulumi:"assignments"`
	DefaultDatacenter GtmCidrmapDefaultDatacenter `pulumi:"defaultDatacenter"`
	Domain            string                      `pulumi:"domain"`
	Name              *string                     `pulumi:"name"`
	WaitOnComplete    *bool                       `pulumi:"waitOnComplete"`
}

// The set of arguments for constructing a GtmCidrmap resource.
type GtmCidrmapArgs struct {
	Assignments       GtmCidrmapAssignmentArrayInput
	DefaultDatacenter GtmCidrmapDefaultDatacenterInput
	Domain            pulumi.StringInput
	Name              pulumi.StringPtrInput
	WaitOnComplete    pulumi.BoolPtrInput
}

func (GtmCidrmapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmCidrmapArgs)(nil)).Elem()
}

type GtmCidrmapInput interface {
	pulumi.Input

	ToGtmCidrmapOutput() GtmCidrmapOutput
	ToGtmCidrmapOutputWithContext(ctx context.Context) GtmCidrmapOutput
}

func (*GtmCidrmap) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmCidrmap)(nil))
}

func (i *GtmCidrmap) ToGtmCidrmapOutput() GtmCidrmapOutput {
	return i.ToGtmCidrmapOutputWithContext(context.Background())
}

func (i *GtmCidrmap) ToGtmCidrmapOutputWithContext(ctx context.Context) GtmCidrmapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmCidrmapOutput)
}

type GtmCidrmapOutput struct {
	*pulumi.OutputState
}

func (GtmCidrmapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmCidrmap)(nil))
}

func (o GtmCidrmapOutput) ToGtmCidrmapOutput() GtmCidrmapOutput {
	return o
}

func (o GtmCidrmapOutput) ToGtmCidrmapOutputWithContext(ctx context.Context) GtmCidrmapOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GtmCidrmapOutput{})
}
