// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `DnsZone` resource to configure a DNS zone that integrates with your existing DNS infrastructure.
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
// 		_, err := akamai.NewDnsZone(ctx, "demozone", &akamai.DnsZoneArgs{
// 			Comment:  pulumi.String("some comment"),
// 			Contract: pulumi.String("ctr_1-AB123"),
// 			Group:    pulumi.String("100"),
// 			Masters: pulumi.StringArray{
// 				pulumi.String("1.2.3.4"),
// 				pulumi.String("1.2.3.5"),
// 			},
// 			SignAndServe: pulumi.Bool(false),
// 			Type:         pulumi.String("secondary"),
// 			Zone:         pulumi.String("example.com"),
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
// * `comment` - (Required) A descriptive comment.
// * `contract` - (Required) The contract ID.
// * `group` - (Required) The currently selected group ID.
// * `zone` - (Required) The domain zone, encapsulating any nested subdomains.
// * `type` - (Required) Whether the zone is `primary`, `secondary`, or `alias`.
// * `masters` - (Required for `secondary` zones) The names or IP addresses of the nameservers that the zone data should be retrieved from.
// * `target` - (Required for `alias` zones) The name of the zone whose configuration this zone will copy.
// * `signAndServe` - (Optional) Whether DNSSEC Sign and Serve is enabled.
// * `signAndServeAlgorithm` - (Optional) The algorithm used by Sign and Serve.
// * `tsigKey` - (Optional) The TSIG Key used in secure zone transfers. If used, requires these arguments:
//     * `name` - The key name.
//     * `algorithm` - The hashing algorithm.
//     * `secret` - String known between transfer endpoints.
// * `endCustomerId` - (Optional) A free form identifier for the zone.
type DnsZone struct {
	pulumi.CustomResourceState

	ActivationState       pulumi.StringOutput      `pulumi:"activationState"`
	AliasCount            pulumi.IntOutput         `pulumi:"aliasCount"`
	Comment               pulumi.StringPtrOutput   `pulumi:"comment"`
	Contract              pulumi.StringOutput      `pulumi:"contract"`
	EndCustomerId         pulumi.StringPtrOutput   `pulumi:"endCustomerId"`
	Group                 pulumi.StringOutput      `pulumi:"group"`
	Masters               pulumi.StringArrayOutput `pulumi:"masters"`
	SignAndServe          pulumi.BoolPtrOutput     `pulumi:"signAndServe"`
	SignAndServeAlgorithm pulumi.StringPtrOutput   `pulumi:"signAndServeAlgorithm"`
	Target                pulumi.StringPtrOutput   `pulumi:"target"`
	TsigKey               DnsZoneTsigKeyPtrOutput  `pulumi:"tsigKey"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
	VersionId             pulumi.StringOutput      `pulumi:"versionId"`
	Zone                  pulumi.StringOutput      `pulumi:"zone"`
}

// NewDnsZone registers a new resource with the given unique name, arguments, and options.
func NewDnsZone(ctx *pulumi.Context,
	name string, args *DnsZoneArgs, opts ...pulumi.ResourceOption) (*DnsZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Contract == nil {
		return nil, errors.New("invalid value for required argument 'Contract'")
	}
	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("akamai:edgedns/dnsZone:DnsZone"),
		},
	})
	opts = append(opts, aliases)
	var resource DnsZone
	err := ctx.RegisterResource("akamai:index/dnsZone:DnsZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsZone gets an existing DnsZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsZoneState, opts ...pulumi.ResourceOption) (*DnsZone, error) {
	var resource DnsZone
	err := ctx.ReadResource("akamai:index/dnsZone:DnsZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsZone resources.
type dnsZoneState struct {
	ActivationState       *string         `pulumi:"activationState"`
	AliasCount            *int            `pulumi:"aliasCount"`
	Comment               *string         `pulumi:"comment"`
	Contract              *string         `pulumi:"contract"`
	EndCustomerId         *string         `pulumi:"endCustomerId"`
	Group                 *string         `pulumi:"group"`
	Masters               []string        `pulumi:"masters"`
	SignAndServe          *bool           `pulumi:"signAndServe"`
	SignAndServeAlgorithm *string         `pulumi:"signAndServeAlgorithm"`
	Target                *string         `pulumi:"target"`
	TsigKey               *DnsZoneTsigKey `pulumi:"tsigKey"`
	Type                  *string         `pulumi:"type"`
	VersionId             *string         `pulumi:"versionId"`
	Zone                  *string         `pulumi:"zone"`
}

type DnsZoneState struct {
	ActivationState       pulumi.StringPtrInput
	AliasCount            pulumi.IntPtrInput
	Comment               pulumi.StringPtrInput
	Contract              pulumi.StringPtrInput
	EndCustomerId         pulumi.StringPtrInput
	Group                 pulumi.StringPtrInput
	Masters               pulumi.StringArrayInput
	SignAndServe          pulumi.BoolPtrInput
	SignAndServeAlgorithm pulumi.StringPtrInput
	Target                pulumi.StringPtrInput
	TsigKey               DnsZoneTsigKeyPtrInput
	Type                  pulumi.StringPtrInput
	VersionId             pulumi.StringPtrInput
	Zone                  pulumi.StringPtrInput
}

func (DnsZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsZoneState)(nil)).Elem()
}

type dnsZoneArgs struct {
	Comment               *string         `pulumi:"comment"`
	Contract              string          `pulumi:"contract"`
	EndCustomerId         *string         `pulumi:"endCustomerId"`
	Group                 string          `pulumi:"group"`
	Masters               []string        `pulumi:"masters"`
	SignAndServe          *bool           `pulumi:"signAndServe"`
	SignAndServeAlgorithm *string         `pulumi:"signAndServeAlgorithm"`
	Target                *string         `pulumi:"target"`
	TsigKey               *DnsZoneTsigKey `pulumi:"tsigKey"`
	Type                  string          `pulumi:"type"`
	Zone                  string          `pulumi:"zone"`
}

// The set of arguments for constructing a DnsZone resource.
type DnsZoneArgs struct {
	Comment               pulumi.StringPtrInput
	Contract              pulumi.StringInput
	EndCustomerId         pulumi.StringPtrInput
	Group                 pulumi.StringInput
	Masters               pulumi.StringArrayInput
	SignAndServe          pulumi.BoolPtrInput
	SignAndServeAlgorithm pulumi.StringPtrInput
	Target                pulumi.StringPtrInput
	TsigKey               DnsZoneTsigKeyPtrInput
	Type                  pulumi.StringInput
	Zone                  pulumi.StringInput
}

func (DnsZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsZoneArgs)(nil)).Elem()
}

type DnsZoneInput interface {
	pulumi.Input

	ToDnsZoneOutput() DnsZoneOutput
	ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput
}

func (*DnsZone) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsZone)(nil))
}

func (i *DnsZone) ToDnsZoneOutput() DnsZoneOutput {
	return i.ToDnsZoneOutputWithContext(context.Background())
}

func (i *DnsZone) ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneOutput)
}

func (i *DnsZone) ToDnsZonePtrOutput() DnsZonePtrOutput {
	return i.ToDnsZonePtrOutputWithContext(context.Background())
}

func (i *DnsZone) ToDnsZonePtrOutputWithContext(ctx context.Context) DnsZonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZonePtrOutput)
}

type DnsZonePtrInput interface {
	pulumi.Input

	ToDnsZonePtrOutput() DnsZonePtrOutput
	ToDnsZonePtrOutputWithContext(ctx context.Context) DnsZonePtrOutput
}

type dnsZonePtrType DnsZoneArgs

func (*dnsZonePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsZone)(nil))
}

func (i *dnsZonePtrType) ToDnsZonePtrOutput() DnsZonePtrOutput {
	return i.ToDnsZonePtrOutputWithContext(context.Background())
}

func (i *dnsZonePtrType) ToDnsZonePtrOutputWithContext(ctx context.Context) DnsZonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZonePtrOutput)
}

// DnsZoneArrayInput is an input type that accepts DnsZoneArray and DnsZoneArrayOutput values.
// You can construct a concrete instance of `DnsZoneArrayInput` via:
//
//          DnsZoneArray{ DnsZoneArgs{...} }
type DnsZoneArrayInput interface {
	pulumi.Input

	ToDnsZoneArrayOutput() DnsZoneArrayOutput
	ToDnsZoneArrayOutputWithContext(context.Context) DnsZoneArrayOutput
}

type DnsZoneArray []DnsZoneInput

func (DnsZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DnsZone)(nil))
}

func (i DnsZoneArray) ToDnsZoneArrayOutput() DnsZoneArrayOutput {
	return i.ToDnsZoneArrayOutputWithContext(context.Background())
}

func (i DnsZoneArray) ToDnsZoneArrayOutputWithContext(ctx context.Context) DnsZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneArrayOutput)
}

// DnsZoneMapInput is an input type that accepts DnsZoneMap and DnsZoneMapOutput values.
// You can construct a concrete instance of `DnsZoneMapInput` via:
//
//          DnsZoneMap{ "key": DnsZoneArgs{...} }
type DnsZoneMapInput interface {
	pulumi.Input

	ToDnsZoneMapOutput() DnsZoneMapOutput
	ToDnsZoneMapOutputWithContext(context.Context) DnsZoneMapOutput
}

type DnsZoneMap map[string]DnsZoneInput

func (DnsZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DnsZone)(nil))
}

func (i DnsZoneMap) ToDnsZoneMapOutput() DnsZoneMapOutput {
	return i.ToDnsZoneMapOutputWithContext(context.Background())
}

func (i DnsZoneMap) ToDnsZoneMapOutputWithContext(ctx context.Context) DnsZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneMapOutput)
}

type DnsZoneOutput struct {
	*pulumi.OutputState
}

func (DnsZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsZone)(nil))
}

func (o DnsZoneOutput) ToDnsZoneOutput() DnsZoneOutput {
	return o
}

func (o DnsZoneOutput) ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput {
	return o
}

func (o DnsZoneOutput) ToDnsZonePtrOutput() DnsZonePtrOutput {
	return o.ToDnsZonePtrOutputWithContext(context.Background())
}

func (o DnsZoneOutput) ToDnsZonePtrOutputWithContext(ctx context.Context) DnsZonePtrOutput {
	return o.ApplyT(func(v DnsZone) *DnsZone {
		return &v
	}).(DnsZonePtrOutput)
}

type DnsZonePtrOutput struct {
	*pulumi.OutputState
}

func (DnsZonePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsZone)(nil))
}

func (o DnsZonePtrOutput) ToDnsZonePtrOutput() DnsZonePtrOutput {
	return o
}

func (o DnsZonePtrOutput) ToDnsZonePtrOutputWithContext(ctx context.Context) DnsZonePtrOutput {
	return o
}

type DnsZoneArrayOutput struct{ *pulumi.OutputState }

func (DnsZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DnsZone)(nil))
}

func (o DnsZoneArrayOutput) ToDnsZoneArrayOutput() DnsZoneArrayOutput {
	return o
}

func (o DnsZoneArrayOutput) ToDnsZoneArrayOutputWithContext(ctx context.Context) DnsZoneArrayOutput {
	return o
}

func (o DnsZoneArrayOutput) Index(i pulumi.IntInput) DnsZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DnsZone {
		return vs[0].([]DnsZone)[vs[1].(int)]
	}).(DnsZoneOutput)
}

type DnsZoneMapOutput struct{ *pulumi.OutputState }

func (DnsZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DnsZone)(nil))
}

func (o DnsZoneMapOutput) ToDnsZoneMapOutput() DnsZoneMapOutput {
	return o
}

func (o DnsZoneMapOutput) ToDnsZoneMapOutputWithContext(ctx context.Context) DnsZoneMapOutput {
	return o
}

func (o DnsZoneMapOutput) MapIndex(k pulumi.StringInput) DnsZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DnsZone {
		return vs[0].(map[string]DnsZone)[vs[1].(string)]
	}).(DnsZoneOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsZoneOutput{})
	pulumi.RegisterOutputType(DnsZonePtrOutput{})
	pulumi.RegisterOutputType(DnsZoneArrayOutput{})
	pulumi.RegisterOutputType(DnsZoneMapOutput{})
}
