// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `GtmAsmap` provides the resource for creating, configuring and importing a gtm AS Map to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existingDomainName`:`existingMapName`
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
// 		_, err := akamai.NewGtmAsmap(ctx, "demoAsmap", &akamai.GtmAsmapArgs{
// 			DefaultDatacenter: &akamai.GtmAsmapDefaultDatacenterArgs{
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
type GtmAsmap struct {
	pulumi.CustomResourceState

	// * `datacenterId`
	// * `nickname`
	Assignments       GtmAsmapAssignmentArrayOutput   `pulumi:"assignments"`
	DefaultDatacenter GtmAsmapDefaultDatacenterOutput `pulumi:"defaultDatacenter"`
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

// NewGtmAsmap registers a new resource with the given unique name, arguments, and options.
func NewGtmAsmap(ctx *pulumi.Context,
	name string, args *GtmAsmapArgs, opts ...pulumi.ResourceOption) (*GtmAsmap, error) {
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
			Type: pulumi.String("akamai:trafficmanagement/gtmASmap:GtmASmap"),
		},
	})
	opts = append(opts, aliases)
	var resource GtmAsmap
	err := ctx.RegisterResource("akamai:index/gtmAsmap:GtmAsmap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGtmAsmap gets an existing GtmAsmap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGtmAsmap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GtmAsmapState, opts ...pulumi.ResourceOption) (*GtmAsmap, error) {
	var resource GtmAsmap
	err := ctx.ReadResource("akamai:index/gtmAsmap:GtmAsmap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GtmAsmap resources.
type gtmAsmapState struct {
	// * `datacenterId`
	// * `nickname`
	Assignments       []GtmAsmapAssignment       `pulumi:"assignments"`
	DefaultDatacenter *GtmAsmapDefaultDatacenter `pulumi:"defaultDatacenter"`
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

type GtmAsmapState struct {
	// * `datacenterId`
	// * `nickname`
	Assignments       GtmAsmapAssignmentArrayInput
	DefaultDatacenter GtmAsmapDefaultDatacenterPtrInput
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

func (GtmAsmapState) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmAsmapState)(nil)).Elem()
}

type gtmAsmapArgs struct {
	// * `datacenterId`
	// * `nickname`
	Assignments       []GtmAsmapAssignment      `pulumi:"assignments"`
	DefaultDatacenter GtmAsmapDefaultDatacenter `pulumi:"defaultDatacenter"`
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

// The set of arguments for constructing a GtmAsmap resource.
type GtmAsmapArgs struct {
	// * `datacenterId`
	// * `nickname`
	Assignments       GtmAsmapAssignmentArrayInput
	DefaultDatacenter GtmAsmapDefaultDatacenterInput
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

func (GtmAsmapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmAsmapArgs)(nil)).Elem()
}

type GtmAsmapInput interface {
	pulumi.Input

	ToGtmAsmapOutput() GtmAsmapOutput
	ToGtmAsmapOutputWithContext(ctx context.Context) GtmAsmapOutput
}

func (*GtmAsmap) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmAsmap)(nil))
}

func (i *GtmAsmap) ToGtmAsmapOutput() GtmAsmapOutput {
	return i.ToGtmAsmapOutputWithContext(context.Background())
}

func (i *GtmAsmap) ToGtmAsmapOutputWithContext(ctx context.Context) GtmAsmapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmAsmapOutput)
}

func (i *GtmAsmap) ToGtmAsmapPtrOutput() GtmAsmapPtrOutput {
	return i.ToGtmAsmapPtrOutputWithContext(context.Background())
}

func (i *GtmAsmap) ToGtmAsmapPtrOutputWithContext(ctx context.Context) GtmAsmapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmAsmapPtrOutput)
}

type GtmAsmapPtrInput interface {
	pulumi.Input

	ToGtmAsmapPtrOutput() GtmAsmapPtrOutput
	ToGtmAsmapPtrOutputWithContext(ctx context.Context) GtmAsmapPtrOutput
}

type gtmAsmapPtrType GtmAsmapArgs

func (*gtmAsmapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmAsmap)(nil))
}

func (i *gtmAsmapPtrType) ToGtmAsmapPtrOutput() GtmAsmapPtrOutput {
	return i.ToGtmAsmapPtrOutputWithContext(context.Background())
}

func (i *gtmAsmapPtrType) ToGtmAsmapPtrOutputWithContext(ctx context.Context) GtmAsmapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmAsmapPtrOutput)
}

// GtmAsmapArrayInput is an input type that accepts GtmAsmapArray and GtmAsmapArrayOutput values.
// You can construct a concrete instance of `GtmAsmapArrayInput` via:
//
//          GtmAsmapArray{ GtmAsmapArgs{...} }
type GtmAsmapArrayInput interface {
	pulumi.Input

	ToGtmAsmapArrayOutput() GtmAsmapArrayOutput
	ToGtmAsmapArrayOutputWithContext(context.Context) GtmAsmapArrayOutput
}

type GtmAsmapArray []GtmAsmapInput

func (GtmAsmapArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GtmAsmap)(nil))
}

func (i GtmAsmapArray) ToGtmAsmapArrayOutput() GtmAsmapArrayOutput {
	return i.ToGtmAsmapArrayOutputWithContext(context.Background())
}

func (i GtmAsmapArray) ToGtmAsmapArrayOutputWithContext(ctx context.Context) GtmAsmapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmAsmapArrayOutput)
}

// GtmAsmapMapInput is an input type that accepts GtmAsmapMap and GtmAsmapMapOutput values.
// You can construct a concrete instance of `GtmAsmapMapInput` via:
//
//          GtmAsmapMap{ "key": GtmAsmapArgs{...} }
type GtmAsmapMapInput interface {
	pulumi.Input

	ToGtmAsmapMapOutput() GtmAsmapMapOutput
	ToGtmAsmapMapOutputWithContext(context.Context) GtmAsmapMapOutput
}

type GtmAsmapMap map[string]GtmAsmapInput

func (GtmAsmapMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GtmAsmap)(nil))
}

func (i GtmAsmapMap) ToGtmAsmapMapOutput() GtmAsmapMapOutput {
	return i.ToGtmAsmapMapOutputWithContext(context.Background())
}

func (i GtmAsmapMap) ToGtmAsmapMapOutputWithContext(ctx context.Context) GtmAsmapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmAsmapMapOutput)
}

type GtmAsmapOutput struct {
	*pulumi.OutputState
}

func (GtmAsmapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmAsmap)(nil))
}

func (o GtmAsmapOutput) ToGtmAsmapOutput() GtmAsmapOutput {
	return o
}

func (o GtmAsmapOutput) ToGtmAsmapOutputWithContext(ctx context.Context) GtmAsmapOutput {
	return o
}

func (o GtmAsmapOutput) ToGtmAsmapPtrOutput() GtmAsmapPtrOutput {
	return o.ToGtmAsmapPtrOutputWithContext(context.Background())
}

func (o GtmAsmapOutput) ToGtmAsmapPtrOutputWithContext(ctx context.Context) GtmAsmapPtrOutput {
	return o.ApplyT(func(v GtmAsmap) *GtmAsmap {
		return &v
	}).(GtmAsmapPtrOutput)
}

type GtmAsmapPtrOutput struct {
	*pulumi.OutputState
}

func (GtmAsmapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmAsmap)(nil))
}

func (o GtmAsmapPtrOutput) ToGtmAsmapPtrOutput() GtmAsmapPtrOutput {
	return o
}

func (o GtmAsmapPtrOutput) ToGtmAsmapPtrOutputWithContext(ctx context.Context) GtmAsmapPtrOutput {
	return o
}

type GtmAsmapArrayOutput struct{ *pulumi.OutputState }

func (GtmAsmapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GtmAsmap)(nil))
}

func (o GtmAsmapArrayOutput) ToGtmAsmapArrayOutput() GtmAsmapArrayOutput {
	return o
}

func (o GtmAsmapArrayOutput) ToGtmAsmapArrayOutputWithContext(ctx context.Context) GtmAsmapArrayOutput {
	return o
}

func (o GtmAsmapArrayOutput) Index(i pulumi.IntInput) GtmAsmapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GtmAsmap {
		return vs[0].([]GtmAsmap)[vs[1].(int)]
	}).(GtmAsmapOutput)
}

type GtmAsmapMapOutput struct{ *pulumi.OutputState }

func (GtmAsmapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GtmAsmap)(nil))
}

func (o GtmAsmapMapOutput) ToGtmAsmapMapOutput() GtmAsmapMapOutput {
	return o
}

func (o GtmAsmapMapOutput) ToGtmAsmapMapOutputWithContext(ctx context.Context) GtmAsmapMapOutput {
	return o
}

func (o GtmAsmapMapOutput) MapIndex(k pulumi.StringInput) GtmAsmapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GtmAsmap {
		return vs[0].(map[string]GtmAsmap)[vs[1].(string)]
	}).(GtmAsmapOutput)
}

func init() {
	pulumi.RegisterOutputType(GtmAsmapOutput{})
	pulumi.RegisterOutputType(GtmAsmapPtrOutput{})
	pulumi.RegisterOutputType(GtmAsmapArrayOutput{})
	pulumi.RegisterOutputType(GtmAsmapMapOutput{})
}
