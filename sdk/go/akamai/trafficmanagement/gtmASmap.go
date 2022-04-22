// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `GtmAsmap` resource to create, configure, and import a GTM Autonomous System (AS) map. AS mapping lets you configure a GTM property that returns a CNAME based on the AS number associated with the requester's IP address.
//
// You can reuse maps for multiple properties or create new ones. AS maps split the Internet into multiple AS block zones. Properties that use AS maps can specify handout integers for each zone. AS mapping lets you configure a property that directs users to a specific environment or to the origin.
//
// > **Note** Import requires an ID with this format: `existingDomainName`:`existingMapName`.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/v2/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := akamai.NewGtmAsmap(ctx, "demoAsmap", &akamai.GtmAsmapArgs{
// 			DefaultDatacenter: &GtmAsmapDefaultDatacenterArgs{
// 				DatacenterId: pulumi.Int(5400),
// 				Nickname:     pulumi.String("All Other AS numbers"),
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
// ## Schema reference
//
// You can download the GTM AS Map backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#asmap) page.
//
// Deprecated: akamai.trafficmanagement.GtmASmap has been deprecated in favor of akamai.GtmAsmap
type GtmASmap struct {
	pulumi.CustomResourceState

	// Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
	Assignments GtmASmapAssignmentArrayOutput `pulumi:"assignments"`
	// A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
	DefaultDatacenter GtmASmapDefaultDatacenterOutput `pulumi:"defaultDatacenter"`
	// The GTM Domain name for the AS map.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
	Name pulumi.StringOutput `pulumi:"name"`
	// A boolean that, if `true`, waits for transaction to complete.
	WaitOnComplete pulumi.BoolPtrOutput `pulumi:"waitOnComplete"`
}

// NewGtmASmap registers a new resource with the given unique name, arguments, and options.
func NewGtmASmap(ctx *pulumi.Context,
	name string, args *GtmASmapArgs, opts ...pulumi.ResourceOption) (*GtmASmap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultDatacenter == nil {
		return nil, errors.New("invalid value for required argument 'DefaultDatacenter'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	var resource GtmASmap
	err := ctx.RegisterResource("akamai:trafficmanagement/gtmASmap:GtmASmap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGtmASmap gets an existing GtmASmap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGtmASmap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GtmASmapState, opts ...pulumi.ResourceOption) (*GtmASmap, error) {
	var resource GtmASmap
	err := ctx.ReadResource("akamai:trafficmanagement/gtmASmap:GtmASmap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GtmASmap resources.
type gtmASmapState struct {
	// Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
	Assignments []GtmASmapAssignment `pulumi:"assignments"`
	// A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
	DefaultDatacenter *GtmASmapDefaultDatacenter `pulumi:"defaultDatacenter"`
	// The GTM Domain name for the AS map.
	Domain *string `pulumi:"domain"`
	// A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
	Name *string `pulumi:"name"`
	// A boolean that, if `true`, waits for transaction to complete.
	WaitOnComplete *bool `pulumi:"waitOnComplete"`
}

type GtmASmapState struct {
	// Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
	Assignments GtmASmapAssignmentArrayInput
	// A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
	DefaultDatacenter GtmASmapDefaultDatacenterPtrInput
	// The GTM Domain name for the AS map.
	Domain pulumi.StringPtrInput
	// A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
	Name pulumi.StringPtrInput
	// A boolean that, if `true`, waits for transaction to complete.
	WaitOnComplete pulumi.BoolPtrInput
}

func (GtmASmapState) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmASmapState)(nil)).Elem()
}

type gtmASmapArgs struct {
	// Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
	Assignments []GtmASmapAssignment `pulumi:"assignments"`
	// A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
	DefaultDatacenter GtmASmapDefaultDatacenter `pulumi:"defaultDatacenter"`
	// The GTM Domain name for the AS map.
	Domain string `pulumi:"domain"`
	// A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
	Name *string `pulumi:"name"`
	// A boolean that, if `true`, waits for transaction to complete.
	WaitOnComplete *bool `pulumi:"waitOnComplete"`
}

// The set of arguments for constructing a GtmASmap resource.
type GtmASmapArgs struct {
	// Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
	Assignments GtmASmapAssignmentArrayInput
	// A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
	DefaultDatacenter GtmASmapDefaultDatacenterInput
	// The GTM Domain name for the AS map.
	Domain pulumi.StringInput
	// A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
	Name pulumi.StringPtrInput
	// A boolean that, if `true`, waits for transaction to complete.
	WaitOnComplete pulumi.BoolPtrInput
}

func (GtmASmapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmASmapArgs)(nil)).Elem()
}

type GtmASmapInput interface {
	pulumi.Input

	ToGtmASmapOutput() GtmASmapOutput
	ToGtmASmapOutputWithContext(ctx context.Context) GtmASmapOutput
}

func (*GtmASmap) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmASmap)(nil)).Elem()
}

func (i *GtmASmap) ToGtmASmapOutput() GtmASmapOutput {
	return i.ToGtmASmapOutputWithContext(context.Background())
}

func (i *GtmASmap) ToGtmASmapOutputWithContext(ctx context.Context) GtmASmapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmASmapOutput)
}

// GtmASmapArrayInput is an input type that accepts GtmASmapArray and GtmASmapArrayOutput values.
// You can construct a concrete instance of `GtmASmapArrayInput` via:
//
//          GtmASmapArray{ GtmASmapArgs{...} }
type GtmASmapArrayInput interface {
	pulumi.Input

	ToGtmASmapArrayOutput() GtmASmapArrayOutput
	ToGtmASmapArrayOutputWithContext(context.Context) GtmASmapArrayOutput
}

type GtmASmapArray []GtmASmapInput

func (GtmASmapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GtmASmap)(nil)).Elem()
}

func (i GtmASmapArray) ToGtmASmapArrayOutput() GtmASmapArrayOutput {
	return i.ToGtmASmapArrayOutputWithContext(context.Background())
}

func (i GtmASmapArray) ToGtmASmapArrayOutputWithContext(ctx context.Context) GtmASmapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmASmapArrayOutput)
}

// GtmASmapMapInput is an input type that accepts GtmASmapMap and GtmASmapMapOutput values.
// You can construct a concrete instance of `GtmASmapMapInput` via:
//
//          GtmASmapMap{ "key": GtmASmapArgs{...} }
type GtmASmapMapInput interface {
	pulumi.Input

	ToGtmASmapMapOutput() GtmASmapMapOutput
	ToGtmASmapMapOutputWithContext(context.Context) GtmASmapMapOutput
}

type GtmASmapMap map[string]GtmASmapInput

func (GtmASmapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GtmASmap)(nil)).Elem()
}

func (i GtmASmapMap) ToGtmASmapMapOutput() GtmASmapMapOutput {
	return i.ToGtmASmapMapOutputWithContext(context.Background())
}

func (i GtmASmapMap) ToGtmASmapMapOutputWithContext(ctx context.Context) GtmASmapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmASmapMapOutput)
}

type GtmASmapOutput struct{ *pulumi.OutputState }

func (GtmASmapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmASmap)(nil)).Elem()
}

func (o GtmASmapOutput) ToGtmASmapOutput() GtmASmapOutput {
	return o
}

func (o GtmASmapOutput) ToGtmASmapOutputWithContext(ctx context.Context) GtmASmapOutput {
	return o
}

type GtmASmapArrayOutput struct{ *pulumi.OutputState }

func (GtmASmapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GtmASmap)(nil)).Elem()
}

func (o GtmASmapArrayOutput) ToGtmASmapArrayOutput() GtmASmapArrayOutput {
	return o
}

func (o GtmASmapArrayOutput) ToGtmASmapArrayOutputWithContext(ctx context.Context) GtmASmapArrayOutput {
	return o
}

func (o GtmASmapArrayOutput) Index(i pulumi.IntInput) GtmASmapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GtmASmap {
		return vs[0].([]*GtmASmap)[vs[1].(int)]
	}).(GtmASmapOutput)
}

type GtmASmapMapOutput struct{ *pulumi.OutputState }

func (GtmASmapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GtmASmap)(nil)).Elem()
}

func (o GtmASmapMapOutput) ToGtmASmapMapOutput() GtmASmapMapOutput {
	return o
}

func (o GtmASmapMapOutput) ToGtmASmapMapOutputWithContext(ctx context.Context) GtmASmapMapOutput {
	return o
}

func (o GtmASmapMapOutput) MapIndex(k pulumi.StringInput) GtmASmapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GtmASmap {
		return vs[0].(map[string]*GtmASmap)[vs[1].(string)]
	}).(GtmASmapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GtmASmapInput)(nil)).Elem(), &GtmASmap{})
	pulumi.RegisterInputType(reflect.TypeOf((*GtmASmapArrayInput)(nil)).Elem(), GtmASmapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GtmASmapMapInput)(nil)).Elem(), GtmASmapMap{})
	pulumi.RegisterOutputType(GtmASmapOutput{})
	pulumi.RegisterOutputType(GtmASmapArrayOutput{})
	pulumi.RegisterOutputType(GtmASmapMapOutput{})
}
