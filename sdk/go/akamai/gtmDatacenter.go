// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `GtmDatacenter` resource to create, configure, and import a GTM data center. A GTM data center represents a customer data center and is also known as a traffic target, a location containing many servers GTM can direct traffic to.
//
// GTM uses data centers to scale load balancing. For example, you might have data centers in both New York and Amsterdam and want to balance load between them. You can configure GTM to send US users to the New York data center and European users to the data center in Amsterdam.
//
// > **Note** Import requires an ID with this format: `existingDomainName`:`existingDatacenterId`.
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
// 		_, err := akamai.NewGtmDatacenter(ctx, "demoDatacenter", &akamai.GtmDatacenterArgs{
// 			Domain:   pulumi.String("demo_domain.akadns.net"),
// 			Nickname: pulumi.String("demo_datacenter"),
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
// * `domain` - (Required) The GTM domain name for the data center.
// * `waitOnComplete` - (Optional) A boolean, that if set to `true`, waits for transaction to complete.
// * `nickname` - (Optional) A descriptive label for the data center.
// * `defaultLoadObject` - (Optional) Specifies the load reporting interface between you and the GTM system. If used, requires these additional arguments:
//   * `loadObject` - A load object is a file that provides real-time information about the current load, maximum allowable load, and target load on each resource.
//   * `loadObjectPort` - Specifies the TCP port to connect to when requesting the load object.
//   * `loadServers` - Specifies a list of servers to request the load object from.
// * `city` - (Optional) The name of the city where the data center is located.
// * `cloneOf` - (Optional) Identifies the data center’s `datacenterId` of which this data center is a clone.
// * `cloudServerTargeting` - (Optional) A boolean indicating whether to balance load between two or more servers in a cloud environment.
// * `cloudServerHostHeaderOverride` - (Optional) A boolean that, if set to `true`, Akamai's liveness test agents use the Host header configured in the liveness test.
// * `continent` - (Optional) A two-letter code that specifies the continent where the data center maps to.
// * `country` - (Optional) A two-letter ISO 3166 country code that specifies the country where the data center maps to.
// * `latitude` - (Optional) Specifies the geographical latitude of the data center’s position. See also longitude within this object.
// * `longitude` - (Optional) Specifies the geographic longitude of the data center’s position. See also latitude within this object.
// * `stateOrProvince` - (Optional) Specifies a two-letter ISO 3166 country code for the state or province where the data center is located.
//
// ## Attribute reference
//
// This resource returns these computed attributes in the state file:
//
// * `datacenterId` - A unique identifier for an existing data center in the domain.
// * `pingInterval`
// * `pingPacketSize`
// * `scorePenalty`
// * `servermonitorLivenessCount`
// * `servermonitorLoadCount`
// * `servermonitorPool`
// * `virtual` - A boolean indicating whether the data center is virtual or physical, the latter meaning the data center has an Akamai Network Agent installed, and its physical location (`latitude`, `longitude`) is fixed. Either `true` if virtual or `false` if physical.
//
// ## Schema reference
//
// You can download the GTM Data Center backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#datacenter) page.
type GtmDatacenter struct {
	pulumi.CustomResourceState

	City                          pulumi.StringPtrOutput                  `pulumi:"city"`
	CloneOf                       pulumi.IntPtrOutput                     `pulumi:"cloneOf"`
	CloudServerHostHeaderOverride pulumi.BoolPtrOutput                    `pulumi:"cloudServerHostHeaderOverride"`
	CloudServerTargeting          pulumi.BoolPtrOutput                    `pulumi:"cloudServerTargeting"`
	Continent                     pulumi.StringPtrOutput                  `pulumi:"continent"`
	Country                       pulumi.StringPtrOutput                  `pulumi:"country"`
	DatacenterId                  pulumi.IntOutput                        `pulumi:"datacenterId"`
	DefaultLoadObject             GtmDatacenterDefaultLoadObjectPtrOutput `pulumi:"defaultLoadObject"`
	Domain                        pulumi.StringOutput                     `pulumi:"domain"`
	Latitude                      pulumi.Float64PtrOutput                 `pulumi:"latitude"`
	Longitude                     pulumi.Float64PtrOutput                 `pulumi:"longitude"`
	Nickname                      pulumi.StringPtrOutput                  `pulumi:"nickname"`
	PingInterval                  pulumi.IntOutput                        `pulumi:"pingInterval"`
	PingPacketSize                pulumi.IntOutput                        `pulumi:"pingPacketSize"`
	ScorePenalty                  pulumi.IntOutput                        `pulumi:"scorePenalty"`
	ServermonitorLivenessCount    pulumi.IntOutput                        `pulumi:"servermonitorLivenessCount"`
	ServermonitorLoadCount        pulumi.IntOutput                        `pulumi:"servermonitorLoadCount"`
	ServermonitorPool             pulumi.StringOutput                     `pulumi:"servermonitorPool"`
	StateOrProvince               pulumi.StringPtrOutput                  `pulumi:"stateOrProvince"`
	Virtual                       pulumi.BoolOutput                       `pulumi:"virtual"`
	WaitOnComplete                pulumi.BoolPtrOutput                    `pulumi:"waitOnComplete"`
}

// NewGtmDatacenter registers a new resource with the given unique name, arguments, and options.
func NewGtmDatacenter(ctx *pulumi.Context,
	name string, args *GtmDatacenterArgs, opts ...pulumi.ResourceOption) (*GtmDatacenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("akamai:trafficmanagement/gtmDatacenter:GtmDatacenter"),
		},
	})
	opts = append(opts, aliases)
	var resource GtmDatacenter
	err := ctx.RegisterResource("akamai:index/gtmDatacenter:GtmDatacenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGtmDatacenter gets an existing GtmDatacenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGtmDatacenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GtmDatacenterState, opts ...pulumi.ResourceOption) (*GtmDatacenter, error) {
	var resource GtmDatacenter
	err := ctx.ReadResource("akamai:index/gtmDatacenter:GtmDatacenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GtmDatacenter resources.
type gtmDatacenterState struct {
	City                          *string                         `pulumi:"city"`
	CloneOf                       *int                            `pulumi:"cloneOf"`
	CloudServerHostHeaderOverride *bool                           `pulumi:"cloudServerHostHeaderOverride"`
	CloudServerTargeting          *bool                           `pulumi:"cloudServerTargeting"`
	Continent                     *string                         `pulumi:"continent"`
	Country                       *string                         `pulumi:"country"`
	DatacenterId                  *int                            `pulumi:"datacenterId"`
	DefaultLoadObject             *GtmDatacenterDefaultLoadObject `pulumi:"defaultLoadObject"`
	Domain                        *string                         `pulumi:"domain"`
	Latitude                      *float64                        `pulumi:"latitude"`
	Longitude                     *float64                        `pulumi:"longitude"`
	Nickname                      *string                         `pulumi:"nickname"`
	PingInterval                  *int                            `pulumi:"pingInterval"`
	PingPacketSize                *int                            `pulumi:"pingPacketSize"`
	ScorePenalty                  *int                            `pulumi:"scorePenalty"`
	ServermonitorLivenessCount    *int                            `pulumi:"servermonitorLivenessCount"`
	ServermonitorLoadCount        *int                            `pulumi:"servermonitorLoadCount"`
	ServermonitorPool             *string                         `pulumi:"servermonitorPool"`
	StateOrProvince               *string                         `pulumi:"stateOrProvince"`
	Virtual                       *bool                           `pulumi:"virtual"`
	WaitOnComplete                *bool                           `pulumi:"waitOnComplete"`
}

type GtmDatacenterState struct {
	City                          pulumi.StringPtrInput
	CloneOf                       pulumi.IntPtrInput
	CloudServerHostHeaderOverride pulumi.BoolPtrInput
	CloudServerTargeting          pulumi.BoolPtrInput
	Continent                     pulumi.StringPtrInput
	Country                       pulumi.StringPtrInput
	DatacenterId                  pulumi.IntPtrInput
	DefaultLoadObject             GtmDatacenterDefaultLoadObjectPtrInput
	Domain                        pulumi.StringPtrInput
	Latitude                      pulumi.Float64PtrInput
	Longitude                     pulumi.Float64PtrInput
	Nickname                      pulumi.StringPtrInput
	PingInterval                  pulumi.IntPtrInput
	PingPacketSize                pulumi.IntPtrInput
	ScorePenalty                  pulumi.IntPtrInput
	ServermonitorLivenessCount    pulumi.IntPtrInput
	ServermonitorLoadCount        pulumi.IntPtrInput
	ServermonitorPool             pulumi.StringPtrInput
	StateOrProvince               pulumi.StringPtrInput
	Virtual                       pulumi.BoolPtrInput
	WaitOnComplete                pulumi.BoolPtrInput
}

func (GtmDatacenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmDatacenterState)(nil)).Elem()
}

type gtmDatacenterArgs struct {
	City                          *string                         `pulumi:"city"`
	CloneOf                       *int                            `pulumi:"cloneOf"`
	CloudServerHostHeaderOverride *bool                           `pulumi:"cloudServerHostHeaderOverride"`
	CloudServerTargeting          *bool                           `pulumi:"cloudServerTargeting"`
	Continent                     *string                         `pulumi:"continent"`
	Country                       *string                         `pulumi:"country"`
	DefaultLoadObject             *GtmDatacenterDefaultLoadObject `pulumi:"defaultLoadObject"`
	Domain                        string                          `pulumi:"domain"`
	Latitude                      *float64                        `pulumi:"latitude"`
	Longitude                     *float64                        `pulumi:"longitude"`
	Nickname                      *string                         `pulumi:"nickname"`
	StateOrProvince               *string                         `pulumi:"stateOrProvince"`
	WaitOnComplete                *bool                           `pulumi:"waitOnComplete"`
}

// The set of arguments for constructing a GtmDatacenter resource.
type GtmDatacenterArgs struct {
	City                          pulumi.StringPtrInput
	CloneOf                       pulumi.IntPtrInput
	CloudServerHostHeaderOverride pulumi.BoolPtrInput
	CloudServerTargeting          pulumi.BoolPtrInput
	Continent                     pulumi.StringPtrInput
	Country                       pulumi.StringPtrInput
	DefaultLoadObject             GtmDatacenterDefaultLoadObjectPtrInput
	Domain                        pulumi.StringInput
	Latitude                      pulumi.Float64PtrInput
	Longitude                     pulumi.Float64PtrInput
	Nickname                      pulumi.StringPtrInput
	StateOrProvince               pulumi.StringPtrInput
	WaitOnComplete                pulumi.BoolPtrInput
}

func (GtmDatacenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmDatacenterArgs)(nil)).Elem()
}

type GtmDatacenterInput interface {
	pulumi.Input

	ToGtmDatacenterOutput() GtmDatacenterOutput
	ToGtmDatacenterOutputWithContext(ctx context.Context) GtmDatacenterOutput
}

func (*GtmDatacenter) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmDatacenter)(nil))
}

func (i *GtmDatacenter) ToGtmDatacenterOutput() GtmDatacenterOutput {
	return i.ToGtmDatacenterOutputWithContext(context.Background())
}

func (i *GtmDatacenter) ToGtmDatacenterOutputWithContext(ctx context.Context) GtmDatacenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmDatacenterOutput)
}

func (i *GtmDatacenter) ToGtmDatacenterPtrOutput() GtmDatacenterPtrOutput {
	return i.ToGtmDatacenterPtrOutputWithContext(context.Background())
}

func (i *GtmDatacenter) ToGtmDatacenterPtrOutputWithContext(ctx context.Context) GtmDatacenterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmDatacenterPtrOutput)
}

type GtmDatacenterPtrInput interface {
	pulumi.Input

	ToGtmDatacenterPtrOutput() GtmDatacenterPtrOutput
	ToGtmDatacenterPtrOutputWithContext(ctx context.Context) GtmDatacenterPtrOutput
}

type gtmDatacenterPtrType GtmDatacenterArgs

func (*gtmDatacenterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmDatacenter)(nil))
}

func (i *gtmDatacenterPtrType) ToGtmDatacenterPtrOutput() GtmDatacenterPtrOutput {
	return i.ToGtmDatacenterPtrOutputWithContext(context.Background())
}

func (i *gtmDatacenterPtrType) ToGtmDatacenterPtrOutputWithContext(ctx context.Context) GtmDatacenterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmDatacenterPtrOutput)
}

// GtmDatacenterArrayInput is an input type that accepts GtmDatacenterArray and GtmDatacenterArrayOutput values.
// You can construct a concrete instance of `GtmDatacenterArrayInput` via:
//
//          GtmDatacenterArray{ GtmDatacenterArgs{...} }
type GtmDatacenterArrayInput interface {
	pulumi.Input

	ToGtmDatacenterArrayOutput() GtmDatacenterArrayOutput
	ToGtmDatacenterArrayOutputWithContext(context.Context) GtmDatacenterArrayOutput
}

type GtmDatacenterArray []GtmDatacenterInput

func (GtmDatacenterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GtmDatacenter)(nil))
}

func (i GtmDatacenterArray) ToGtmDatacenterArrayOutput() GtmDatacenterArrayOutput {
	return i.ToGtmDatacenterArrayOutputWithContext(context.Background())
}

func (i GtmDatacenterArray) ToGtmDatacenterArrayOutputWithContext(ctx context.Context) GtmDatacenterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmDatacenterArrayOutput)
}

// GtmDatacenterMapInput is an input type that accepts GtmDatacenterMap and GtmDatacenterMapOutput values.
// You can construct a concrete instance of `GtmDatacenterMapInput` via:
//
//          GtmDatacenterMap{ "key": GtmDatacenterArgs{...} }
type GtmDatacenterMapInput interface {
	pulumi.Input

	ToGtmDatacenterMapOutput() GtmDatacenterMapOutput
	ToGtmDatacenterMapOutputWithContext(context.Context) GtmDatacenterMapOutput
}

type GtmDatacenterMap map[string]GtmDatacenterInput

func (GtmDatacenterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GtmDatacenter)(nil))
}

func (i GtmDatacenterMap) ToGtmDatacenterMapOutput() GtmDatacenterMapOutput {
	return i.ToGtmDatacenterMapOutputWithContext(context.Background())
}

func (i GtmDatacenterMap) ToGtmDatacenterMapOutputWithContext(ctx context.Context) GtmDatacenterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmDatacenterMapOutput)
}

type GtmDatacenterOutput struct {
	*pulumi.OutputState
}

func (GtmDatacenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmDatacenter)(nil))
}

func (o GtmDatacenterOutput) ToGtmDatacenterOutput() GtmDatacenterOutput {
	return o
}

func (o GtmDatacenterOutput) ToGtmDatacenterOutputWithContext(ctx context.Context) GtmDatacenterOutput {
	return o
}

func (o GtmDatacenterOutput) ToGtmDatacenterPtrOutput() GtmDatacenterPtrOutput {
	return o.ToGtmDatacenterPtrOutputWithContext(context.Background())
}

func (o GtmDatacenterOutput) ToGtmDatacenterPtrOutputWithContext(ctx context.Context) GtmDatacenterPtrOutput {
	return o.ApplyT(func(v GtmDatacenter) *GtmDatacenter {
		return &v
	}).(GtmDatacenterPtrOutput)
}

type GtmDatacenterPtrOutput struct {
	*pulumi.OutputState
}

func (GtmDatacenterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmDatacenter)(nil))
}

func (o GtmDatacenterPtrOutput) ToGtmDatacenterPtrOutput() GtmDatacenterPtrOutput {
	return o
}

func (o GtmDatacenterPtrOutput) ToGtmDatacenterPtrOutputWithContext(ctx context.Context) GtmDatacenterPtrOutput {
	return o
}

type GtmDatacenterArrayOutput struct{ *pulumi.OutputState }

func (GtmDatacenterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GtmDatacenter)(nil))
}

func (o GtmDatacenterArrayOutput) ToGtmDatacenterArrayOutput() GtmDatacenterArrayOutput {
	return o
}

func (o GtmDatacenterArrayOutput) ToGtmDatacenterArrayOutputWithContext(ctx context.Context) GtmDatacenterArrayOutput {
	return o
}

func (o GtmDatacenterArrayOutput) Index(i pulumi.IntInput) GtmDatacenterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GtmDatacenter {
		return vs[0].([]GtmDatacenter)[vs[1].(int)]
	}).(GtmDatacenterOutput)
}

type GtmDatacenterMapOutput struct{ *pulumi.OutputState }

func (GtmDatacenterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GtmDatacenter)(nil))
}

func (o GtmDatacenterMapOutput) ToGtmDatacenterMapOutput() GtmDatacenterMapOutput {
	return o
}

func (o GtmDatacenterMapOutput) ToGtmDatacenterMapOutputWithContext(ctx context.Context) GtmDatacenterMapOutput {
	return o
}

func (o GtmDatacenterMapOutput) MapIndex(k pulumi.StringInput) GtmDatacenterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GtmDatacenter {
		return vs[0].(map[string]GtmDatacenter)[vs[1].(string)]
	}).(GtmDatacenterOutput)
}

func init() {
	pulumi.RegisterOutputType(GtmDatacenterOutput{})
	pulumi.RegisterOutputType(GtmDatacenterPtrOutput{})
	pulumi.RegisterOutputType(GtmDatacenterArrayOutput{})
	pulumi.RegisterOutputType(GtmDatacenterMapOutput{})
}
