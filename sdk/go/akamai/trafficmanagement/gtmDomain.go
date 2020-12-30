// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain
type GtmDomain struct {
	pulumi.CustomResourceState

	CnameCoalescingEnabled       pulumi.BoolPtrOutput     `pulumi:"cnameCoalescingEnabled"`
	Comment                      pulumi.StringPtrOutput   `pulumi:"comment"`
	Contract                     pulumi.StringPtrOutput   `pulumi:"contract"`
	DefaultErrorPenalty          pulumi.IntPtrOutput      `pulumi:"defaultErrorPenalty"`
	DefaultHealthMax             pulumi.Float64Output     `pulumi:"defaultHealthMax"`
	DefaultHealthMultiplier      pulumi.Float64Output     `pulumi:"defaultHealthMultiplier"`
	DefaultHealthThreshold       pulumi.Float64Output     `pulumi:"defaultHealthThreshold"`
	DefaultMaxUnreachablePenalty pulumi.IntOutput         `pulumi:"defaultMaxUnreachablePenalty"`
	DefaultSslClientCertificate  pulumi.StringPtrOutput   `pulumi:"defaultSslClientCertificate"`
	DefaultSslClientPrivateKey   pulumi.StringPtrOutput   `pulumi:"defaultSslClientPrivateKey"`
	DefaultTimeoutPenalty        pulumi.IntPtrOutput      `pulumi:"defaultTimeoutPenalty"`
	DefaultUnreachableThreshold  pulumi.Float64Output     `pulumi:"defaultUnreachableThreshold"`
	EmailNotificationLists       pulumi.StringArrayOutput `pulumi:"emailNotificationLists"`
	EndUserMappingEnabled        pulumi.BoolPtrOutput     `pulumi:"endUserMappingEnabled"`
	Group                        pulumi.StringPtrOutput   `pulumi:"group"`
	LoadFeedback                 pulumi.BoolPtrOutput     `pulumi:"loadFeedback"`
	LoadImbalancePercentage      pulumi.Float64PtrOutput  `pulumi:"loadImbalancePercentage"`
	MapUpdateInterval            pulumi.IntOutput         `pulumi:"mapUpdateInterval"`
	MaxProperties                pulumi.IntOutput         `pulumi:"maxProperties"`
	MaxResources                 pulumi.IntOutput         `pulumi:"maxResources"`
	MaxTestTimeout               pulumi.Float64Output     `pulumi:"maxTestTimeout"`
	MaxTtl                       pulumi.IntOutput         `pulumi:"maxTtl"`
	MinPingableRegionFraction    pulumi.Float64Output     `pulumi:"minPingableRegionFraction"`
	MinTestInterval              pulumi.IntOutput         `pulumi:"minTestInterval"`
	MinTtl                       pulumi.IntOutput         `pulumi:"minTtl"`
	Name                         pulumi.StringOutput      `pulumi:"name"`
	PingInterval                 pulumi.IntOutput         `pulumi:"pingInterval"`
	PingPacketSize               pulumi.IntOutput         `pulumi:"pingPacketSize"`
	RoundRobinPrefix             pulumi.StringOutput      `pulumi:"roundRobinPrefix"`
	ServermonitorLivenessCount   pulumi.IntOutput         `pulumi:"servermonitorLivenessCount"`
	ServermonitorLoadCount       pulumi.IntOutput         `pulumi:"servermonitorLoadCount"`
	ServermonitorPool            pulumi.StringOutput      `pulumi:"servermonitorPool"`
	Type                         pulumi.StringOutput      `pulumi:"type"`
	WaitOnComplete               pulumi.BoolPtrOutput     `pulumi:"waitOnComplete"`
}

// NewGtmDomain registers a new resource with the given unique name, arguments, and options.
func NewGtmDomain(ctx *pulumi.Context,
	name string, args *GtmDomainArgs, opts ...pulumi.ResourceOption) (*GtmDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource GtmDomain
	err := ctx.RegisterResource("akamai:trafficmanagement/gtmDomain:GtmDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGtmDomain gets an existing GtmDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGtmDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GtmDomainState, opts ...pulumi.ResourceOption) (*GtmDomain, error) {
	var resource GtmDomain
	err := ctx.ReadResource("akamai:trafficmanagement/gtmDomain:GtmDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GtmDomain resources.
type gtmDomainState struct {
	CnameCoalescingEnabled       *bool    `pulumi:"cnameCoalescingEnabled"`
	Comment                      *string  `pulumi:"comment"`
	Contract                     *string  `pulumi:"contract"`
	DefaultErrorPenalty          *int     `pulumi:"defaultErrorPenalty"`
	DefaultHealthMax             *float64 `pulumi:"defaultHealthMax"`
	DefaultHealthMultiplier      *float64 `pulumi:"defaultHealthMultiplier"`
	DefaultHealthThreshold       *float64 `pulumi:"defaultHealthThreshold"`
	DefaultMaxUnreachablePenalty *int     `pulumi:"defaultMaxUnreachablePenalty"`
	DefaultSslClientCertificate  *string  `pulumi:"defaultSslClientCertificate"`
	DefaultSslClientPrivateKey   *string  `pulumi:"defaultSslClientPrivateKey"`
	DefaultTimeoutPenalty        *int     `pulumi:"defaultTimeoutPenalty"`
	DefaultUnreachableThreshold  *float64 `pulumi:"defaultUnreachableThreshold"`
	EmailNotificationLists       []string `pulumi:"emailNotificationLists"`
	EndUserMappingEnabled        *bool    `pulumi:"endUserMappingEnabled"`
	Group                        *string  `pulumi:"group"`
	LoadFeedback                 *bool    `pulumi:"loadFeedback"`
	LoadImbalancePercentage      *float64 `pulumi:"loadImbalancePercentage"`
	MapUpdateInterval            *int     `pulumi:"mapUpdateInterval"`
	MaxProperties                *int     `pulumi:"maxProperties"`
	MaxResources                 *int     `pulumi:"maxResources"`
	MaxTestTimeout               *float64 `pulumi:"maxTestTimeout"`
	MaxTtl                       *int     `pulumi:"maxTtl"`
	MinPingableRegionFraction    *float64 `pulumi:"minPingableRegionFraction"`
	MinTestInterval              *int     `pulumi:"minTestInterval"`
	MinTtl                       *int     `pulumi:"minTtl"`
	Name                         *string  `pulumi:"name"`
	PingInterval                 *int     `pulumi:"pingInterval"`
	PingPacketSize               *int     `pulumi:"pingPacketSize"`
	RoundRobinPrefix             *string  `pulumi:"roundRobinPrefix"`
	ServermonitorLivenessCount   *int     `pulumi:"servermonitorLivenessCount"`
	ServermonitorLoadCount       *int     `pulumi:"servermonitorLoadCount"`
	ServermonitorPool            *string  `pulumi:"servermonitorPool"`
	Type                         *string  `pulumi:"type"`
	WaitOnComplete               *bool    `pulumi:"waitOnComplete"`
}

type GtmDomainState struct {
	CnameCoalescingEnabled       pulumi.BoolPtrInput
	Comment                      pulumi.StringPtrInput
	Contract                     pulumi.StringPtrInput
	DefaultErrorPenalty          pulumi.IntPtrInput
	DefaultHealthMax             pulumi.Float64PtrInput
	DefaultHealthMultiplier      pulumi.Float64PtrInput
	DefaultHealthThreshold       pulumi.Float64PtrInput
	DefaultMaxUnreachablePenalty pulumi.IntPtrInput
	DefaultSslClientCertificate  pulumi.StringPtrInput
	DefaultSslClientPrivateKey   pulumi.StringPtrInput
	DefaultTimeoutPenalty        pulumi.IntPtrInput
	DefaultUnreachableThreshold  pulumi.Float64PtrInput
	EmailNotificationLists       pulumi.StringArrayInput
	EndUserMappingEnabled        pulumi.BoolPtrInput
	Group                        pulumi.StringPtrInput
	LoadFeedback                 pulumi.BoolPtrInput
	LoadImbalancePercentage      pulumi.Float64PtrInput
	MapUpdateInterval            pulumi.IntPtrInput
	MaxProperties                pulumi.IntPtrInput
	MaxResources                 pulumi.IntPtrInput
	MaxTestTimeout               pulumi.Float64PtrInput
	MaxTtl                       pulumi.IntPtrInput
	MinPingableRegionFraction    pulumi.Float64PtrInput
	MinTestInterval              pulumi.IntPtrInput
	MinTtl                       pulumi.IntPtrInput
	Name                         pulumi.StringPtrInput
	PingInterval                 pulumi.IntPtrInput
	PingPacketSize               pulumi.IntPtrInput
	RoundRobinPrefix             pulumi.StringPtrInput
	ServermonitorLivenessCount   pulumi.IntPtrInput
	ServermonitorLoadCount       pulumi.IntPtrInput
	ServermonitorPool            pulumi.StringPtrInput
	Type                         pulumi.StringPtrInput
	WaitOnComplete               pulumi.BoolPtrInput
}

func (GtmDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmDomainState)(nil)).Elem()
}

type gtmDomainArgs struct {
	CnameCoalescingEnabled      *bool    `pulumi:"cnameCoalescingEnabled"`
	Comment                     *string  `pulumi:"comment"`
	Contract                    *string  `pulumi:"contract"`
	DefaultErrorPenalty         *int     `pulumi:"defaultErrorPenalty"`
	DefaultSslClientCertificate *string  `pulumi:"defaultSslClientCertificate"`
	DefaultSslClientPrivateKey  *string  `pulumi:"defaultSslClientPrivateKey"`
	DefaultTimeoutPenalty       *int     `pulumi:"defaultTimeoutPenalty"`
	EmailNotificationLists      []string `pulumi:"emailNotificationLists"`
	EndUserMappingEnabled       *bool    `pulumi:"endUserMappingEnabled"`
	Group                       *string  `pulumi:"group"`
	LoadFeedback                *bool    `pulumi:"loadFeedback"`
	LoadImbalancePercentage     *float64 `pulumi:"loadImbalancePercentage"`
	Name                        *string  `pulumi:"name"`
	Type                        string   `pulumi:"type"`
	WaitOnComplete              *bool    `pulumi:"waitOnComplete"`
}

// The set of arguments for constructing a GtmDomain resource.
type GtmDomainArgs struct {
	CnameCoalescingEnabled      pulumi.BoolPtrInput
	Comment                     pulumi.StringPtrInput
	Contract                    pulumi.StringPtrInput
	DefaultErrorPenalty         pulumi.IntPtrInput
	DefaultSslClientCertificate pulumi.StringPtrInput
	DefaultSslClientPrivateKey  pulumi.StringPtrInput
	DefaultTimeoutPenalty       pulumi.IntPtrInput
	EmailNotificationLists      pulumi.StringArrayInput
	EndUserMappingEnabled       pulumi.BoolPtrInput
	Group                       pulumi.StringPtrInput
	LoadFeedback                pulumi.BoolPtrInput
	LoadImbalancePercentage     pulumi.Float64PtrInput
	Name                        pulumi.StringPtrInput
	Type                        pulumi.StringInput
	WaitOnComplete              pulumi.BoolPtrInput
}

func (GtmDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmDomainArgs)(nil)).Elem()
}

type GtmDomainInput interface {
	pulumi.Input

	ToGtmDomainOutput() GtmDomainOutput
	ToGtmDomainOutputWithContext(ctx context.Context) GtmDomainOutput
}

func (GtmDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmDomain)(nil)).Elem()
}

func (i GtmDomain) ToGtmDomainOutput() GtmDomainOutput {
	return i.ToGtmDomainOutputWithContext(context.Background())
}

func (i GtmDomain) ToGtmDomainOutputWithContext(ctx context.Context) GtmDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmDomainOutput)
}

type GtmDomainOutput struct {
	*pulumi.OutputState
}

func (GtmDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmDomainOutput)(nil)).Elem()
}

func (o GtmDomainOutput) ToGtmDomainOutput() GtmDomainOutput {
	return o
}

func (o GtmDomainOutput) ToGtmDomainOutputWithContext(ctx context.Context) GtmDomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GtmDomainOutput{})
}
