// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `GtmCidrmap` provides the resource for creating, configuring and importing a gtm Cidr Map to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existingDomainName`:`existingMapName`
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := akamai.NewGtmCidrmap(ctx, "demoCidrmap", &akamai.GtmCidrmapArgs{
// 			DefaultDatacenter: &akamai.GtmCidrmapDefaultDatacenterArgs{
// 				DatacenterId: pulumi.Int(5400),
// 				Nickname:     pulumi.String("All Other CIDR Blocks"),
// 			},
// 			Domain: pulumi.String("demo_domain.akadns.net"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GtmCidrmap struct {
	pulumi.CustomResourceState

	// * `datacenterId`
	// * `nickname`
	Assignments       GtmCidrmapAssignmentArrayOutput   `pulumi:"assignments"`
	DefaultDatacenter GtmCidrmapDefaultDatacenterOutput `pulumi:"defaultDatacenter"`
	// Domain name
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Resource name
	// * `defaultDatacenter`
	// * `datacenterId`
	// * `nickname`
	Name pulumi.StringOutput `pulumi:"name"`
	// Wait for transaction to complete
	WaitOnComplete pulumi.BoolPtrOutput `pulumi:"waitOnComplete"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("akamai:trafficmanagement/gtmCidrmap:GtmCidrmap"),
		},
	})
	opts = append(opts, aliases)
	var resource GtmCidrmap
	err := ctx.RegisterResource("akamai:index/gtmCidrmap:GtmCidrmap", name, args, &resource, opts...)
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
	err := ctx.ReadResource("akamai:index/gtmCidrmap:GtmCidrmap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GtmCidrmap resources.
type gtmCidrmapState struct {
	// * `datacenterId`
	// * `nickname`
	Assignments       []GtmCidrmapAssignment       `pulumi:"assignments"`
	DefaultDatacenter *GtmCidrmapDefaultDatacenter `pulumi:"defaultDatacenter"`
	// Domain name
	Domain *string `pulumi:"domain"`
	// Resource name
	// * `defaultDatacenter`
	// * `datacenterId`
	// * `nickname`
	Name *string `pulumi:"name"`
	// Wait for transaction to complete
	WaitOnComplete *bool `pulumi:"waitOnComplete"`
}

type GtmCidrmapState struct {
	// * `datacenterId`
	// * `nickname`
	Assignments       GtmCidrmapAssignmentArrayInput
	DefaultDatacenter GtmCidrmapDefaultDatacenterPtrInput
	// Domain name
	Domain pulumi.StringPtrInput
	// Resource name
	// * `defaultDatacenter`
	// * `datacenterId`
	// * `nickname`
	Name pulumi.StringPtrInput
	// Wait for transaction to complete
	WaitOnComplete pulumi.BoolPtrInput
}

func (GtmCidrmapState) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmCidrmapState)(nil)).Elem()
}

type gtmCidrmapArgs struct {
	// * `datacenterId`
	// * `nickname`
	Assignments       []GtmCidrmapAssignment      `pulumi:"assignments"`
	DefaultDatacenter GtmCidrmapDefaultDatacenter `pulumi:"defaultDatacenter"`
	// Domain name
	Domain string `pulumi:"domain"`
	// Resource name
	// * `defaultDatacenter`
	// * `datacenterId`
	// * `nickname`
	Name *string `pulumi:"name"`
	// Wait for transaction to complete
	WaitOnComplete *bool `pulumi:"waitOnComplete"`
}

// The set of arguments for constructing a GtmCidrmap resource.
type GtmCidrmapArgs struct {
	// * `datacenterId`
	// * `nickname`
	Assignments       GtmCidrmapAssignmentArrayInput
	DefaultDatacenter GtmCidrmapDefaultDatacenterInput
	// Domain name
	Domain pulumi.StringInput
	// Resource name
	// * `defaultDatacenter`
	// * `datacenterId`
	// * `nickname`
	Name pulumi.StringPtrInput
	// Wait for transaction to complete
	WaitOnComplete pulumi.BoolPtrInput
}

func (GtmCidrmapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmCidrmapArgs)(nil)).Elem()
}

type GtmCidrmapInput interface {
	pulumi.Input

	ToGtmCidrmapOutput() GtmCidrmapOutput
	ToGtmCidrmapOutputWithContext(ctx context.Context) GtmCidrmapOutput
}

func (GtmCidrmap) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmCidrmap)(nil)).Elem()
}

func (i GtmCidrmap) ToGtmCidrmapOutput() GtmCidrmapOutput {
	return i.ToGtmCidrmapOutputWithContext(context.Background())
}

func (i GtmCidrmap) ToGtmCidrmapOutputWithContext(ctx context.Context) GtmCidrmapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmCidrmapOutput)
}

type GtmCidrmapOutput struct {
	*pulumi.OutputState
}

func (GtmCidrmapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmCidrmapOutput)(nil)).Elem()
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