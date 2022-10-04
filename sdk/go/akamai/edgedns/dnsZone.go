// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edgedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: akamai.edgedns.DnsZone has been deprecated in favor of akamai.DnsZone
type DnsZone struct {
	pulumi.CustomResourceState

	ActivationState pulumi.StringOutput `pulumi:"activationState"`
	AliasCount      pulumi.IntOutput    `pulumi:"aliasCount"`
	// A descriptive comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The contract ID.
	Contract pulumi.StringOutput `pulumi:"contract"`
	// A free form identifier for the zone.
	EndCustomerId pulumi.StringPtrOutput `pulumi:"endCustomerId"`
	// The currently selected group ID.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// The names or IP addresses of the nameservers that the zone data should be retrieved from.
	Masters pulumi.StringArrayOutput `pulumi:"masters"`
	// Whether DNSSEC Sign and Serve is enabled.
	SignAndServe pulumi.BoolPtrOutput `pulumi:"signAndServe"`
	// The algorithm used by Sign and Serve.
	SignAndServeAlgorithm pulumi.StringPtrOutput `pulumi:"signAndServeAlgorithm"`
	// The name of the zone whose configuration this zone will copy.
	Target pulumi.StringPtrOutput `pulumi:"target"`
	// The TSIG Key used in secure zone transfers. If used, requires these arguments:
	TsigKey DnsZoneTsigKeyPtrOutput `pulumi:"tsigKey"`
	// Whether the zone is `primary`, `secondary`, or `alias`.
	Type      pulumi.StringOutput `pulumi:"type"`
	VersionId pulumi.StringOutput `pulumi:"versionId"`
	// The domain zone, encapsulating any nested subdomains.
	Zone pulumi.StringOutput `pulumi:"zone"`
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
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource DnsZone
	err := ctx.RegisterResource("akamai:edgedns/dnsZone:DnsZone", name, args, &resource, opts...)
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
	err := ctx.ReadResource("akamai:edgedns/dnsZone:DnsZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsZone resources.
type dnsZoneState struct {
	ActivationState *string `pulumi:"activationState"`
	AliasCount      *int    `pulumi:"aliasCount"`
	// A descriptive comment.
	Comment *string `pulumi:"comment"`
	// The contract ID.
	Contract *string `pulumi:"contract"`
	// A free form identifier for the zone.
	EndCustomerId *string `pulumi:"endCustomerId"`
	// The currently selected group ID.
	Group *string `pulumi:"group"`
	// The names or IP addresses of the nameservers that the zone data should be retrieved from.
	Masters []string `pulumi:"masters"`
	// Whether DNSSEC Sign and Serve is enabled.
	SignAndServe *bool `pulumi:"signAndServe"`
	// The algorithm used by Sign and Serve.
	SignAndServeAlgorithm *string `pulumi:"signAndServeAlgorithm"`
	// The name of the zone whose configuration this zone will copy.
	Target *string `pulumi:"target"`
	// The TSIG Key used in secure zone transfers. If used, requires these arguments:
	TsigKey *DnsZoneTsigKey `pulumi:"tsigKey"`
	// Whether the zone is `primary`, `secondary`, or `alias`.
	Type      *string `pulumi:"type"`
	VersionId *string `pulumi:"versionId"`
	// The domain zone, encapsulating any nested subdomains.
	Zone *string `pulumi:"zone"`
}

type DnsZoneState struct {
	ActivationState pulumi.StringPtrInput
	AliasCount      pulumi.IntPtrInput
	// A descriptive comment.
	Comment pulumi.StringPtrInput
	// The contract ID.
	Contract pulumi.StringPtrInput
	// A free form identifier for the zone.
	EndCustomerId pulumi.StringPtrInput
	// The currently selected group ID.
	Group pulumi.StringPtrInput
	// The names or IP addresses of the nameservers that the zone data should be retrieved from.
	Masters pulumi.StringArrayInput
	// Whether DNSSEC Sign and Serve is enabled.
	SignAndServe pulumi.BoolPtrInput
	// The algorithm used by Sign and Serve.
	SignAndServeAlgorithm pulumi.StringPtrInput
	// The name of the zone whose configuration this zone will copy.
	Target pulumi.StringPtrInput
	// The TSIG Key used in secure zone transfers. If used, requires these arguments:
	TsigKey DnsZoneTsigKeyPtrInput
	// Whether the zone is `primary`, `secondary`, or `alias`.
	Type      pulumi.StringPtrInput
	VersionId pulumi.StringPtrInput
	// The domain zone, encapsulating any nested subdomains.
	Zone pulumi.StringPtrInput
}

func (DnsZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsZoneState)(nil)).Elem()
}

type dnsZoneArgs struct {
	// A descriptive comment.
	Comment *string `pulumi:"comment"`
	// The contract ID.
	Contract string `pulumi:"contract"`
	// A free form identifier for the zone.
	EndCustomerId *string `pulumi:"endCustomerId"`
	// The currently selected group ID.
	Group *string `pulumi:"group"`
	// The names or IP addresses of the nameservers that the zone data should be retrieved from.
	Masters []string `pulumi:"masters"`
	// Whether DNSSEC Sign and Serve is enabled.
	SignAndServe *bool `pulumi:"signAndServe"`
	// The algorithm used by Sign and Serve.
	SignAndServeAlgorithm *string `pulumi:"signAndServeAlgorithm"`
	// The name of the zone whose configuration this zone will copy.
	Target *string `pulumi:"target"`
	// The TSIG Key used in secure zone transfers. If used, requires these arguments:
	TsigKey *DnsZoneTsigKey `pulumi:"tsigKey"`
	// Whether the zone is `primary`, `secondary`, or `alias`.
	Type string `pulumi:"type"`
	// The domain zone, encapsulating any nested subdomains.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a DnsZone resource.
type DnsZoneArgs struct {
	// A descriptive comment.
	Comment pulumi.StringPtrInput
	// The contract ID.
	Contract pulumi.StringInput
	// A free form identifier for the zone.
	EndCustomerId pulumi.StringPtrInput
	// The currently selected group ID.
	Group pulumi.StringPtrInput
	// The names or IP addresses of the nameservers that the zone data should be retrieved from.
	Masters pulumi.StringArrayInput
	// Whether DNSSEC Sign and Serve is enabled.
	SignAndServe pulumi.BoolPtrInput
	// The algorithm used by Sign and Serve.
	SignAndServeAlgorithm pulumi.StringPtrInput
	// The name of the zone whose configuration this zone will copy.
	Target pulumi.StringPtrInput
	// The TSIG Key used in secure zone transfers. If used, requires these arguments:
	TsigKey DnsZoneTsigKeyPtrInput
	// Whether the zone is `primary`, `secondary`, or `alias`.
	Type pulumi.StringInput
	// The domain zone, encapsulating any nested subdomains.
	Zone pulumi.StringInput
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
	return reflect.TypeOf((**DnsZone)(nil)).Elem()
}

func (i *DnsZone) ToDnsZoneOutput() DnsZoneOutput {
	return i.ToDnsZoneOutputWithContext(context.Background())
}

func (i *DnsZone) ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneOutput)
}

// DnsZoneArrayInput is an input type that accepts DnsZoneArray and DnsZoneArrayOutput values.
// You can construct a concrete instance of `DnsZoneArrayInput` via:
//
//	DnsZoneArray{ DnsZoneArgs{...} }
type DnsZoneArrayInput interface {
	pulumi.Input

	ToDnsZoneArrayOutput() DnsZoneArrayOutput
	ToDnsZoneArrayOutputWithContext(context.Context) DnsZoneArrayOutput
}

type DnsZoneArray []DnsZoneInput

func (DnsZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsZone)(nil)).Elem()
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
//	DnsZoneMap{ "key": DnsZoneArgs{...} }
type DnsZoneMapInput interface {
	pulumi.Input

	ToDnsZoneMapOutput() DnsZoneMapOutput
	ToDnsZoneMapOutputWithContext(context.Context) DnsZoneMapOutput
}

type DnsZoneMap map[string]DnsZoneInput

func (DnsZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsZone)(nil)).Elem()
}

func (i DnsZoneMap) ToDnsZoneMapOutput() DnsZoneMapOutput {
	return i.ToDnsZoneMapOutputWithContext(context.Background())
}

func (i DnsZoneMap) ToDnsZoneMapOutputWithContext(ctx context.Context) DnsZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneMapOutput)
}

type DnsZoneOutput struct{ *pulumi.OutputState }

func (DnsZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsZone)(nil)).Elem()
}

func (o DnsZoneOutput) ToDnsZoneOutput() DnsZoneOutput {
	return o
}

func (o DnsZoneOutput) ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput {
	return o
}

func (o DnsZoneOutput) ActivationState() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.ActivationState }).(pulumi.StringOutput)
}

func (o DnsZoneOutput) AliasCount() pulumi.IntOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.IntOutput { return v.AliasCount }).(pulumi.IntOutput)
}

// A descriptive comment.
func (o DnsZoneOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// The contract ID.
func (o DnsZoneOutput) Contract() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.Contract }).(pulumi.StringOutput)
}

// A free form identifier for the zone.
func (o DnsZoneOutput) EndCustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.EndCustomerId }).(pulumi.StringPtrOutput)
}

// The currently selected group ID.
func (o DnsZoneOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

// The names or IP addresses of the nameservers that the zone data should be retrieved from.
func (o DnsZoneOutput) Masters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringArrayOutput { return v.Masters }).(pulumi.StringArrayOutput)
}

// Whether DNSSEC Sign and Serve is enabled.
func (o DnsZoneOutput) SignAndServe() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.BoolPtrOutput { return v.SignAndServe }).(pulumi.BoolPtrOutput)
}

// The algorithm used by Sign and Serve.
func (o DnsZoneOutput) SignAndServeAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.SignAndServeAlgorithm }).(pulumi.StringPtrOutput)
}

// The name of the zone whose configuration this zone will copy.
func (o DnsZoneOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Target }).(pulumi.StringPtrOutput)
}

// The TSIG Key used in secure zone transfers. If used, requires these arguments:
func (o DnsZoneOutput) TsigKey() DnsZoneTsigKeyPtrOutput {
	return o.ApplyT(func(v *DnsZone) DnsZoneTsigKeyPtrOutput { return v.TsigKey }).(DnsZoneTsigKeyPtrOutput)
}

// Whether the zone is `primary`, `secondary`, or `alias`.
func (o DnsZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DnsZoneOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.VersionId }).(pulumi.StringOutput)
}

// The domain zone, encapsulating any nested subdomains.
func (o DnsZoneOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type DnsZoneArrayOutput struct{ *pulumi.OutputState }

func (DnsZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsZone)(nil)).Elem()
}

func (o DnsZoneArrayOutput) ToDnsZoneArrayOutput() DnsZoneArrayOutput {
	return o
}

func (o DnsZoneArrayOutput) ToDnsZoneArrayOutputWithContext(ctx context.Context) DnsZoneArrayOutput {
	return o
}

func (o DnsZoneArrayOutput) Index(i pulumi.IntInput) DnsZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsZone {
		return vs[0].([]*DnsZone)[vs[1].(int)]
	}).(DnsZoneOutput)
}

type DnsZoneMapOutput struct{ *pulumi.OutputState }

func (DnsZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsZone)(nil)).Elem()
}

func (o DnsZoneMapOutput) ToDnsZoneMapOutput() DnsZoneMapOutput {
	return o
}

func (o DnsZoneMapOutput) ToDnsZoneMapOutputWithContext(ctx context.Context) DnsZoneMapOutput {
	return o
}

func (o DnsZoneMapOutput) MapIndex(k pulumi.StringInput) DnsZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsZone {
		return vs[0].(map[string]*DnsZone)[vs[1].(string)]
	}).(DnsZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneInput)(nil)).Elem(), &DnsZone{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneArrayInput)(nil)).Elem(), DnsZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneMapInput)(nil)).Elem(), DnsZoneMap{})
	pulumi.RegisterOutputType(DnsZoneOutput{})
	pulumi.RegisterOutputType(DnsZoneArrayOutput{})
	pulumi.RegisterOutputType(DnsZoneMapOutput{})
}
