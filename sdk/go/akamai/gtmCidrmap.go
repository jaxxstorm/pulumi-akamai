// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `GtmCidrmap` resource to create, configure, and import a GTM Classless Inter-Domain Routing (CIDR) map. CIDR mapping uses the IP addresses of the requesting name server to provide IP-specific CNAME entries. CNAMEs let you direct internal users to a specific environment or direct them to the origin. This lets you provide different responses to an internal corporate DNS infrastructure, such as internal test environments and another answer for all other name servers (`defaultDatacenter`).
//
//  CIDR maps split the Internet into multiple CIDR block zones. Properties that use a map can specify a handout CNAME for each zone on the property’s editing page. To configure a property for CIDR mapping, your domain needs at least one CIDR map defined.
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
// 		_, err := akamai.NewGtmCidrmap(ctx, "demoCidrmap", &akamai.GtmCidrmapArgs{
// 			DefaultDatacenter: &GtmCidrmapDefaultDatacenterArgs{
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
// ## Argument reference
//
// This resource supports these arguments:
//
// * `domain` - (Required) GTM Domain name for the AS Map.
// * `name` - (Required) A descriptive label for the CIDR map, up to 255 characters.
// * `defaultDatacenter` - (Required) A placeholder for all other CIDR zones not found in these CIDR zones. Requires these additional arguments:
//   * `datacenterId` - (Required) For each property, an identifier for all other CIDR zones.
//   * `nickname` - (Required) A descriptive label for the all other CIDR blocks.
// * `waitOnComplete` - (Optional) A boolean that, if set to `true`, waits for transaction to complete.
// * `assignment` - (Optional) Contains information about the CIDR zone groupings of CIDR blocks. You can have multiple entries with this argument. If used, requires these additional arguments:
//   * `datacenterId` - (Optional) A unique identifier for an existing data center in the domain.
//   * `nickname` - (Optional) A descriptive label for the CIDR zone group, up to 256 characters.
//   * `blocks` - (Optional, list) Specifies an array of CIDR blocks.
//
// ## Schema reference
//
// You can download the GTM CIDR Map backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#cidrmap) page.
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

func (i *GtmCidrmap) ToGtmCidrmapPtrOutput() GtmCidrmapPtrOutput {
	return i.ToGtmCidrmapPtrOutputWithContext(context.Background())
}

func (i *GtmCidrmap) ToGtmCidrmapPtrOutputWithContext(ctx context.Context) GtmCidrmapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmCidrmapPtrOutput)
}

type GtmCidrmapPtrInput interface {
	pulumi.Input

	ToGtmCidrmapPtrOutput() GtmCidrmapPtrOutput
	ToGtmCidrmapPtrOutputWithContext(ctx context.Context) GtmCidrmapPtrOutput
}

type gtmCidrmapPtrType GtmCidrmapArgs

func (*gtmCidrmapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmCidrmap)(nil))
}

func (i *gtmCidrmapPtrType) ToGtmCidrmapPtrOutput() GtmCidrmapPtrOutput {
	return i.ToGtmCidrmapPtrOutputWithContext(context.Background())
}

func (i *gtmCidrmapPtrType) ToGtmCidrmapPtrOutputWithContext(ctx context.Context) GtmCidrmapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmCidrmapPtrOutput)
}

// GtmCidrmapArrayInput is an input type that accepts GtmCidrmapArray and GtmCidrmapArrayOutput values.
// You can construct a concrete instance of `GtmCidrmapArrayInput` via:
//
//          GtmCidrmapArray{ GtmCidrmapArgs{...} }
type GtmCidrmapArrayInput interface {
	pulumi.Input

	ToGtmCidrmapArrayOutput() GtmCidrmapArrayOutput
	ToGtmCidrmapArrayOutputWithContext(context.Context) GtmCidrmapArrayOutput
}

type GtmCidrmapArray []GtmCidrmapInput

func (GtmCidrmapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GtmCidrmap)(nil)).Elem()
}

func (i GtmCidrmapArray) ToGtmCidrmapArrayOutput() GtmCidrmapArrayOutput {
	return i.ToGtmCidrmapArrayOutputWithContext(context.Background())
}

func (i GtmCidrmapArray) ToGtmCidrmapArrayOutputWithContext(ctx context.Context) GtmCidrmapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmCidrmapArrayOutput)
}

// GtmCidrmapMapInput is an input type that accepts GtmCidrmapMap and GtmCidrmapMapOutput values.
// You can construct a concrete instance of `GtmCidrmapMapInput` via:
//
//          GtmCidrmapMap{ "key": GtmCidrmapArgs{...} }
type GtmCidrmapMapInput interface {
	pulumi.Input

	ToGtmCidrmapMapOutput() GtmCidrmapMapOutput
	ToGtmCidrmapMapOutputWithContext(context.Context) GtmCidrmapMapOutput
}

type GtmCidrmapMap map[string]GtmCidrmapInput

func (GtmCidrmapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GtmCidrmap)(nil)).Elem()
}

func (i GtmCidrmapMap) ToGtmCidrmapMapOutput() GtmCidrmapMapOutput {
	return i.ToGtmCidrmapMapOutputWithContext(context.Background())
}

func (i GtmCidrmapMap) ToGtmCidrmapMapOutputWithContext(ctx context.Context) GtmCidrmapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmCidrmapMapOutput)
}

type GtmCidrmapOutput struct{ *pulumi.OutputState }

func (GtmCidrmapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmCidrmap)(nil))
}

func (o GtmCidrmapOutput) ToGtmCidrmapOutput() GtmCidrmapOutput {
	return o
}

func (o GtmCidrmapOutput) ToGtmCidrmapOutputWithContext(ctx context.Context) GtmCidrmapOutput {
	return o
}

func (o GtmCidrmapOutput) ToGtmCidrmapPtrOutput() GtmCidrmapPtrOutput {
	return o.ToGtmCidrmapPtrOutputWithContext(context.Background())
}

func (o GtmCidrmapOutput) ToGtmCidrmapPtrOutputWithContext(ctx context.Context) GtmCidrmapPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GtmCidrmap) *GtmCidrmap {
		return &v
	}).(GtmCidrmapPtrOutput)
}

type GtmCidrmapPtrOutput struct{ *pulumi.OutputState }

func (GtmCidrmapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmCidrmap)(nil))
}

func (o GtmCidrmapPtrOutput) ToGtmCidrmapPtrOutput() GtmCidrmapPtrOutput {
	return o
}

func (o GtmCidrmapPtrOutput) ToGtmCidrmapPtrOutputWithContext(ctx context.Context) GtmCidrmapPtrOutput {
	return o
}

func (o GtmCidrmapPtrOutput) Elem() GtmCidrmapOutput {
	return o.ApplyT(func(v *GtmCidrmap) GtmCidrmap {
		if v != nil {
			return *v
		}
		var ret GtmCidrmap
		return ret
	}).(GtmCidrmapOutput)
}

type GtmCidrmapArrayOutput struct{ *pulumi.OutputState }

func (GtmCidrmapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GtmCidrmap)(nil))
}

func (o GtmCidrmapArrayOutput) ToGtmCidrmapArrayOutput() GtmCidrmapArrayOutput {
	return o
}

func (o GtmCidrmapArrayOutput) ToGtmCidrmapArrayOutputWithContext(ctx context.Context) GtmCidrmapArrayOutput {
	return o
}

func (o GtmCidrmapArrayOutput) Index(i pulumi.IntInput) GtmCidrmapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GtmCidrmap {
		return vs[0].([]GtmCidrmap)[vs[1].(int)]
	}).(GtmCidrmapOutput)
}

type GtmCidrmapMapOutput struct{ *pulumi.OutputState }

func (GtmCidrmapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GtmCidrmap)(nil))
}

func (o GtmCidrmapMapOutput) ToGtmCidrmapMapOutput() GtmCidrmapMapOutput {
	return o
}

func (o GtmCidrmapMapOutput) ToGtmCidrmapMapOutputWithContext(ctx context.Context) GtmCidrmapMapOutput {
	return o
}

func (o GtmCidrmapMapOutput) MapIndex(k pulumi.StringInput) GtmCidrmapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GtmCidrmap {
		return vs[0].(map[string]GtmCidrmap)[vs[1].(string)]
	}).(GtmCidrmapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GtmCidrmapInput)(nil)).Elem(), &GtmCidrmap{})
	pulumi.RegisterInputType(reflect.TypeOf((*GtmCidrmapPtrInput)(nil)).Elem(), &GtmCidrmap{})
	pulumi.RegisterInputType(reflect.TypeOf((*GtmCidrmapArrayInput)(nil)).Elem(), GtmCidrmapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GtmCidrmapMapInput)(nil)).Elem(), GtmCidrmapMap{})
	pulumi.RegisterOutputType(GtmCidrmapOutput{})
	pulumi.RegisterOutputType(GtmCidrmapPtrOutput{})
	pulumi.RegisterOutputType(GtmCidrmapArrayOutput{})
	pulumi.RegisterOutputType(GtmCidrmapMapOutput{})
}
