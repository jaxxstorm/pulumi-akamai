// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GtmResource` lets you create, configure, and import a GTM resource. In GTM, a resource is anything you can measure whose scarcity affects load balancing. Examples of resources include bandwidth, CPU load average, database queries per second, or disk operations per second.
//
// > **Note** Import requires an ID with this format: `existingDomainName`:
// `existingResourceName`.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := akamai.NewGtmResource(ctx, "demoResource", &akamai.GtmResourceArgs{
// 			AggregationType: pulumi.String("latest"),
// 			Domain:          pulumi.String("demo_domain.akadns.net"),
// 			Type:            pulumi.String("XML load object via HTTP"),
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
// You can download the GTM Resource backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#resource) page.
//
// Deprecated: akamai.trafficmanagement.GtmResource has been deprecated in favor of akamai.GtmResource
type GtmResource struct {
	pulumi.CustomResourceState

	// Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
	AggregationType pulumi.StringOutput `pulumi:"aggregationType"`
	// Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
	ConstrainedProperty pulumi.StringPtrOutput `pulumi:"constrainedProperty"`
	// For Akamai internal use only. You can omit the value or set it to `null`.
	DecayRate pulumi.Float64PtrOutput `pulumi:"decayRate"`
	// A descriptive note to help you track what the resource constrains.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// DNS name for the GTM Domain set that includes this property.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Optionally specifies the host header used when fetching the load object.
	HostHeader pulumi.StringPtrOutput `pulumi:"hostHeader"`
	// Specifies the text that comes before the `loadObject`.
	LeaderString pulumi.StringPtrOutput `pulumi:"leaderString"`
	// For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
	LeastSquaresDecay       pulumi.Float64PtrOutput `pulumi:"leastSquaresDecay"`
	LoadImbalancePercentage pulumi.Float64PtrOutput `pulumi:"loadImbalancePercentage"`
	// For Akamai internal use only. You can omit the value or set it to `null`.
	MaxUMultiplicativeIncrement pulumi.Float64PtrOutput `pulumi:"maxUMultiplicativeIncrement"`
	// A descriptive label for the GTM resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
	ResourceInstances GtmResourceResourceInstanceArrayOutput `pulumi:"resourceInstances"`
	// Indicates the kind of `loadObject` format used to determine the load on the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// An optional sanity check that specifies the maximum allowed value for any component of the load object.
	UpperBound pulumi.IntPtrOutput `pulumi:"upperBound"`
	// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
	WaitOnComplete pulumi.BoolPtrOutput `pulumi:"waitOnComplete"`
}

// NewGtmResource registers a new resource with the given unique name, arguments, and options.
func NewGtmResource(ctx *pulumi.Context,
	name string, args *GtmResourceArgs, opts ...pulumi.ResourceOption) (*GtmResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AggregationType == nil {
		return nil, errors.New("invalid value for required argument 'AggregationType'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource GtmResource
	err := ctx.RegisterResource("akamai:trafficmanagement/gtmResource:GtmResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGtmResource gets an existing GtmResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGtmResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GtmResourceState, opts ...pulumi.ResourceOption) (*GtmResource, error) {
	var resource GtmResource
	err := ctx.ReadResource("akamai:trafficmanagement/gtmResource:GtmResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GtmResource resources.
type gtmResourceState struct {
	// Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
	AggregationType *string `pulumi:"aggregationType"`
	// Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
	ConstrainedProperty *string `pulumi:"constrainedProperty"`
	// For Akamai internal use only. You can omit the value or set it to `null`.
	DecayRate *float64 `pulumi:"decayRate"`
	// A descriptive note to help you track what the resource constrains.
	Description *string `pulumi:"description"`
	// DNS name for the GTM Domain set that includes this property.
	Domain *string `pulumi:"domain"`
	// Optionally specifies the host header used when fetching the load object.
	HostHeader *string `pulumi:"hostHeader"`
	// Specifies the text that comes before the `loadObject`.
	LeaderString *string `pulumi:"leaderString"`
	// For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
	LeastSquaresDecay       *float64 `pulumi:"leastSquaresDecay"`
	LoadImbalancePercentage *float64 `pulumi:"loadImbalancePercentage"`
	// For Akamai internal use only. You can omit the value or set it to `null`.
	MaxUMultiplicativeIncrement *float64 `pulumi:"maxUMultiplicativeIncrement"`
	// A descriptive label for the GTM resource.
	Name *string `pulumi:"name"`
	// (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
	ResourceInstances []GtmResourceResourceInstance `pulumi:"resourceInstances"`
	// Indicates the kind of `loadObject` format used to determine the load on the resource.
	Type *string `pulumi:"type"`
	// An optional sanity check that specifies the maximum allowed value for any component of the load object.
	UpperBound *int `pulumi:"upperBound"`
	// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
	WaitOnComplete *bool `pulumi:"waitOnComplete"`
}

type GtmResourceState struct {
	// Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
	AggregationType pulumi.StringPtrInput
	// Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
	ConstrainedProperty pulumi.StringPtrInput
	// For Akamai internal use only. You can omit the value or set it to `null`.
	DecayRate pulumi.Float64PtrInput
	// A descriptive note to help you track what the resource constrains.
	Description pulumi.StringPtrInput
	// DNS name for the GTM Domain set that includes this property.
	Domain pulumi.StringPtrInput
	// Optionally specifies the host header used when fetching the load object.
	HostHeader pulumi.StringPtrInput
	// Specifies the text that comes before the `loadObject`.
	LeaderString pulumi.StringPtrInput
	// For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
	LeastSquaresDecay       pulumi.Float64PtrInput
	LoadImbalancePercentage pulumi.Float64PtrInput
	// For Akamai internal use only. You can omit the value or set it to `null`.
	MaxUMultiplicativeIncrement pulumi.Float64PtrInput
	// A descriptive label for the GTM resource.
	Name pulumi.StringPtrInput
	// (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
	ResourceInstances GtmResourceResourceInstanceArrayInput
	// Indicates the kind of `loadObject` format used to determine the load on the resource.
	Type pulumi.StringPtrInput
	// An optional sanity check that specifies the maximum allowed value for any component of the load object.
	UpperBound pulumi.IntPtrInput
	// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
	WaitOnComplete pulumi.BoolPtrInput
}

func (GtmResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmResourceState)(nil)).Elem()
}

type gtmResourceArgs struct {
	// Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
	AggregationType string `pulumi:"aggregationType"`
	// Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
	ConstrainedProperty *string `pulumi:"constrainedProperty"`
	// For Akamai internal use only. You can omit the value or set it to `null`.
	DecayRate *float64 `pulumi:"decayRate"`
	// A descriptive note to help you track what the resource constrains.
	Description *string `pulumi:"description"`
	// DNS name for the GTM Domain set that includes this property.
	Domain string `pulumi:"domain"`
	// Optionally specifies the host header used when fetching the load object.
	HostHeader *string `pulumi:"hostHeader"`
	// Specifies the text that comes before the `loadObject`.
	LeaderString *string `pulumi:"leaderString"`
	// For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
	LeastSquaresDecay       *float64 `pulumi:"leastSquaresDecay"`
	LoadImbalancePercentage *float64 `pulumi:"loadImbalancePercentage"`
	// For Akamai internal use only. You can omit the value or set it to `null`.
	MaxUMultiplicativeIncrement *float64 `pulumi:"maxUMultiplicativeIncrement"`
	// A descriptive label for the GTM resource.
	Name *string `pulumi:"name"`
	// (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
	ResourceInstances []GtmResourceResourceInstance `pulumi:"resourceInstances"`
	// Indicates the kind of `loadObject` format used to determine the load on the resource.
	Type string `pulumi:"type"`
	// An optional sanity check that specifies the maximum allowed value for any component of the load object.
	UpperBound *int `pulumi:"upperBound"`
	// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
	WaitOnComplete *bool `pulumi:"waitOnComplete"`
}

// The set of arguments for constructing a GtmResource resource.
type GtmResourceArgs struct {
	// Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
	AggregationType pulumi.StringInput
	// Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
	ConstrainedProperty pulumi.StringPtrInput
	// For Akamai internal use only. You can omit the value or set it to `null`.
	DecayRate pulumi.Float64PtrInput
	// A descriptive note to help you track what the resource constrains.
	Description pulumi.StringPtrInput
	// DNS name for the GTM Domain set that includes this property.
	Domain pulumi.StringInput
	// Optionally specifies the host header used when fetching the load object.
	HostHeader pulumi.StringPtrInput
	// Specifies the text that comes before the `loadObject`.
	LeaderString pulumi.StringPtrInput
	// For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
	LeastSquaresDecay       pulumi.Float64PtrInput
	LoadImbalancePercentage pulumi.Float64PtrInput
	// For Akamai internal use only. You can omit the value or set it to `null`.
	MaxUMultiplicativeIncrement pulumi.Float64PtrInput
	// A descriptive label for the GTM resource.
	Name pulumi.StringPtrInput
	// (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
	ResourceInstances GtmResourceResourceInstanceArrayInput
	// Indicates the kind of `loadObject` format used to determine the load on the resource.
	Type pulumi.StringInput
	// An optional sanity check that specifies the maximum allowed value for any component of the load object.
	UpperBound pulumi.IntPtrInput
	// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
	WaitOnComplete pulumi.BoolPtrInput
}

func (GtmResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmResourceArgs)(nil)).Elem()
}

type GtmResourceInput interface {
	pulumi.Input

	ToGtmResourceOutput() GtmResourceOutput
	ToGtmResourceOutputWithContext(ctx context.Context) GtmResourceOutput
}

func (*GtmResource) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmResource)(nil)).Elem()
}

func (i *GtmResource) ToGtmResourceOutput() GtmResourceOutput {
	return i.ToGtmResourceOutputWithContext(context.Background())
}

func (i *GtmResource) ToGtmResourceOutputWithContext(ctx context.Context) GtmResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmResourceOutput)
}

// GtmResourceArrayInput is an input type that accepts GtmResourceArray and GtmResourceArrayOutput values.
// You can construct a concrete instance of `GtmResourceArrayInput` via:
//
//          GtmResourceArray{ GtmResourceArgs{...} }
type GtmResourceArrayInput interface {
	pulumi.Input

	ToGtmResourceArrayOutput() GtmResourceArrayOutput
	ToGtmResourceArrayOutputWithContext(context.Context) GtmResourceArrayOutput
}

type GtmResourceArray []GtmResourceInput

func (GtmResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GtmResource)(nil)).Elem()
}

func (i GtmResourceArray) ToGtmResourceArrayOutput() GtmResourceArrayOutput {
	return i.ToGtmResourceArrayOutputWithContext(context.Background())
}

func (i GtmResourceArray) ToGtmResourceArrayOutputWithContext(ctx context.Context) GtmResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmResourceArrayOutput)
}

// GtmResourceMapInput is an input type that accepts GtmResourceMap and GtmResourceMapOutput values.
// You can construct a concrete instance of `GtmResourceMapInput` via:
//
//          GtmResourceMap{ "key": GtmResourceArgs{...} }
type GtmResourceMapInput interface {
	pulumi.Input

	ToGtmResourceMapOutput() GtmResourceMapOutput
	ToGtmResourceMapOutputWithContext(context.Context) GtmResourceMapOutput
}

type GtmResourceMap map[string]GtmResourceInput

func (GtmResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GtmResource)(nil)).Elem()
}

func (i GtmResourceMap) ToGtmResourceMapOutput() GtmResourceMapOutput {
	return i.ToGtmResourceMapOutputWithContext(context.Background())
}

func (i GtmResourceMap) ToGtmResourceMapOutputWithContext(ctx context.Context) GtmResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmResourceMapOutput)
}

type GtmResourceOutput struct{ *pulumi.OutputState }

func (GtmResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmResource)(nil)).Elem()
}

func (o GtmResourceOutput) ToGtmResourceOutput() GtmResourceOutput {
	return o
}

func (o GtmResourceOutput) ToGtmResourceOutputWithContext(ctx context.Context) GtmResourceOutput {
	return o
}

// Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
func (o GtmResourceOutput) AggregationType() pulumi.StringOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.StringOutput { return v.AggregationType }).(pulumi.StringOutput)
}

// Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
func (o GtmResourceOutput) ConstrainedProperty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.StringPtrOutput { return v.ConstrainedProperty }).(pulumi.StringPtrOutput)
}

// For Akamai internal use only. You can omit the value or set it to `null`.
func (o GtmResourceOutput) DecayRate() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.Float64PtrOutput { return v.DecayRate }).(pulumi.Float64PtrOutput)
}

// A descriptive note to help you track what the resource constrains.
func (o GtmResourceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// DNS name for the GTM Domain set that includes this property.
func (o GtmResourceOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Optionally specifies the host header used when fetching the load object.
func (o GtmResourceOutput) HostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.StringPtrOutput { return v.HostHeader }).(pulumi.StringPtrOutput)
}

// Specifies the text that comes before the `loadObject`.
func (o GtmResourceOutput) LeaderString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.StringPtrOutput { return v.LeaderString }).(pulumi.StringPtrOutput)
}

// For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
func (o GtmResourceOutput) LeastSquaresDecay() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.Float64PtrOutput { return v.LeastSquaresDecay }).(pulumi.Float64PtrOutput)
}

func (o GtmResourceOutput) LoadImbalancePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.Float64PtrOutput { return v.LoadImbalancePercentage }).(pulumi.Float64PtrOutput)
}

// For Akamai internal use only. You can omit the value or set it to `null`.
func (o GtmResourceOutput) MaxUMultiplicativeIncrement() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.Float64PtrOutput { return v.MaxUMultiplicativeIncrement }).(pulumi.Float64PtrOutput)
}

// A descriptive label for the GTM resource.
func (o GtmResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resourceInstance` entries. Requires these arguments:
func (o GtmResourceOutput) ResourceInstances() GtmResourceResourceInstanceArrayOutput {
	return o.ApplyT(func(v *GtmResource) GtmResourceResourceInstanceArrayOutput { return v.ResourceInstances }).(GtmResourceResourceInstanceArrayOutput)
}

// Indicates the kind of `loadObject` format used to determine the load on the resource.
func (o GtmResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// An optional sanity check that specifies the maximum allowed value for any component of the load object.
func (o GtmResourceOutput) UpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.IntPtrOutput { return v.UpperBound }).(pulumi.IntPtrOutput)
}

// A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
func (o GtmResourceOutput) WaitOnComplete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GtmResource) pulumi.BoolPtrOutput { return v.WaitOnComplete }).(pulumi.BoolPtrOutput)
}

type GtmResourceArrayOutput struct{ *pulumi.OutputState }

func (GtmResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GtmResource)(nil)).Elem()
}

func (o GtmResourceArrayOutput) ToGtmResourceArrayOutput() GtmResourceArrayOutput {
	return o
}

func (o GtmResourceArrayOutput) ToGtmResourceArrayOutputWithContext(ctx context.Context) GtmResourceArrayOutput {
	return o
}

func (o GtmResourceArrayOutput) Index(i pulumi.IntInput) GtmResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GtmResource {
		return vs[0].([]*GtmResource)[vs[1].(int)]
	}).(GtmResourceOutput)
}

type GtmResourceMapOutput struct{ *pulumi.OutputState }

func (GtmResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GtmResource)(nil)).Elem()
}

func (o GtmResourceMapOutput) ToGtmResourceMapOutput() GtmResourceMapOutput {
	return o
}

func (o GtmResourceMapOutput) ToGtmResourceMapOutputWithContext(ctx context.Context) GtmResourceMapOutput {
	return o
}

func (o GtmResourceMapOutput) MapIndex(k pulumi.StringInput) GtmResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GtmResource {
		return vs[0].(map[string]*GtmResource)[vs[1].(string)]
	}).(GtmResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GtmResourceInput)(nil)).Elem(), &GtmResource{})
	pulumi.RegisterInputType(reflect.TypeOf((*GtmResourceArrayInput)(nil)).Elem(), GtmResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GtmResourceMapInput)(nil)).Elem(), GtmResourceMap{})
	pulumi.RegisterOutputType(GtmResourceOutput{})
	pulumi.RegisterOutputType(GtmResourceArrayOutput{})
	pulumi.RegisterOutputType(GtmResourceMapOutput{})
}
