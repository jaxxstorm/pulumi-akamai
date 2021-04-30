// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: akamai.trafficmanagement.GtmProperty has been deprecated in favor of akamai.GtmProperty
type GtmProperty struct {
	pulumi.CustomResourceState

	BackupCname               pulumi.StringPtrOutput              `pulumi:"backupCname"`
	BackupIp                  pulumi.StringPtrOutput              `pulumi:"backupIp"`
	BalanceByDownloadScore    pulumi.BoolPtrOutput                `pulumi:"balanceByDownloadScore"`
	Cname                     pulumi.StringPtrOutput              `pulumi:"cname"`
	Comments                  pulumi.StringPtrOutput              `pulumi:"comments"`
	Domain                    pulumi.StringOutput                 `pulumi:"domain"`
	DynamicTtl                pulumi.IntPtrOutput                 `pulumi:"dynamicTtl"`
	FailbackDelay             pulumi.IntPtrOutput                 `pulumi:"failbackDelay"`
	FailoverDelay             pulumi.IntPtrOutput                 `pulumi:"failoverDelay"`
	GhostDemandReporting      pulumi.BoolPtrOutput                `pulumi:"ghostDemandReporting"`
	HandoutLimit              pulumi.IntOutput                    `pulumi:"handoutLimit"`
	HandoutMode               pulumi.StringOutput                 `pulumi:"handoutMode"`
	HealthMax                 pulumi.Float64PtrOutput             `pulumi:"healthMax"`
	HealthMultiplier          pulumi.Float64PtrOutput             `pulumi:"healthMultiplier"`
	HealthThreshold           pulumi.Float64PtrOutput             `pulumi:"healthThreshold"`
	Ipv6                      pulumi.BoolPtrOutput                `pulumi:"ipv6"`
	LivenessTests             GtmPropertyLivenessTestArrayOutput  `pulumi:"livenessTests"`
	LoadImbalancePercentage   pulumi.Float64PtrOutput             `pulumi:"loadImbalancePercentage"`
	MapName                   pulumi.StringPtrOutput              `pulumi:"mapName"`
	MaxUnreachablePenalty     pulumi.IntPtrOutput                 `pulumi:"maxUnreachablePenalty"`
	MinLiveFraction           pulumi.Float64PtrOutput             `pulumi:"minLiveFraction"`
	Name                      pulumi.StringOutput                 `pulumi:"name"`
	ScoreAggregationType      pulumi.StringOutput                 `pulumi:"scoreAggregationType"`
	StaticRrSets              GtmPropertyStaticRrSetArrayOutput   `pulumi:"staticRrSets"`
	StaticTtl                 pulumi.IntPtrOutput                 `pulumi:"staticTtl"`
	StickinessBonusConstant   pulumi.IntPtrOutput                 `pulumi:"stickinessBonusConstant"`
	StickinessBonusPercentage pulumi.IntPtrOutput                 `pulumi:"stickinessBonusPercentage"`
	TrafficTargets            GtmPropertyTrafficTargetArrayOutput `pulumi:"trafficTargets"`
	Type                      pulumi.StringOutput                 `pulumi:"type"`
	UnreachableThreshold      pulumi.Float64PtrOutput             `pulumi:"unreachableThreshold"`
	UseComputedTargets        pulumi.BoolPtrOutput                `pulumi:"useComputedTargets"`
	WaitOnComplete            pulumi.BoolPtrOutput                `pulumi:"waitOnComplete"`
	WeightedHashBitsForIpv4   pulumi.IntOutput                    `pulumi:"weightedHashBitsForIpv4"`
	WeightedHashBitsForIpv6   pulumi.IntOutput                    `pulumi:"weightedHashBitsForIpv6"`
}

// NewGtmProperty registers a new resource with the given unique name, arguments, and options.
func NewGtmProperty(ctx *pulumi.Context,
	name string, args *GtmPropertyArgs, opts ...pulumi.ResourceOption) (*GtmProperty, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.HandoutLimit == nil {
		return nil, errors.New("invalid value for required argument 'HandoutLimit'")
	}
	if args.HandoutMode == nil {
		return nil, errors.New("invalid value for required argument 'HandoutMode'")
	}
	if args.ScoreAggregationType == nil {
		return nil, errors.New("invalid value for required argument 'ScoreAggregationType'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource GtmProperty
	err := ctx.RegisterResource("akamai:trafficmanagement/gtmProperty:GtmProperty", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGtmProperty gets an existing GtmProperty resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGtmProperty(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GtmPropertyState, opts ...pulumi.ResourceOption) (*GtmProperty, error) {
	var resource GtmProperty
	err := ctx.ReadResource("akamai:trafficmanagement/gtmProperty:GtmProperty", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GtmProperty resources.
type gtmPropertyState struct {
	BackupCname               *string                    `pulumi:"backupCname"`
	BackupIp                  *string                    `pulumi:"backupIp"`
	BalanceByDownloadScore    *bool                      `pulumi:"balanceByDownloadScore"`
	Cname                     *string                    `pulumi:"cname"`
	Comments                  *string                    `pulumi:"comments"`
	Domain                    *string                    `pulumi:"domain"`
	DynamicTtl                *int                       `pulumi:"dynamicTtl"`
	FailbackDelay             *int                       `pulumi:"failbackDelay"`
	FailoverDelay             *int                       `pulumi:"failoverDelay"`
	GhostDemandReporting      *bool                      `pulumi:"ghostDemandReporting"`
	HandoutLimit              *int                       `pulumi:"handoutLimit"`
	HandoutMode               *string                    `pulumi:"handoutMode"`
	HealthMax                 *float64                   `pulumi:"healthMax"`
	HealthMultiplier          *float64                   `pulumi:"healthMultiplier"`
	HealthThreshold           *float64                   `pulumi:"healthThreshold"`
	Ipv6                      *bool                      `pulumi:"ipv6"`
	LivenessTests             []GtmPropertyLivenessTest  `pulumi:"livenessTests"`
	LoadImbalancePercentage   *float64                   `pulumi:"loadImbalancePercentage"`
	MapName                   *string                    `pulumi:"mapName"`
	MaxUnreachablePenalty     *int                       `pulumi:"maxUnreachablePenalty"`
	MinLiveFraction           *float64                   `pulumi:"minLiveFraction"`
	Name                      *string                    `pulumi:"name"`
	ScoreAggregationType      *string                    `pulumi:"scoreAggregationType"`
	StaticRrSets              []GtmPropertyStaticRrSet   `pulumi:"staticRrSets"`
	StaticTtl                 *int                       `pulumi:"staticTtl"`
	StickinessBonusConstant   *int                       `pulumi:"stickinessBonusConstant"`
	StickinessBonusPercentage *int                       `pulumi:"stickinessBonusPercentage"`
	TrafficTargets            []GtmPropertyTrafficTarget `pulumi:"trafficTargets"`
	Type                      *string                    `pulumi:"type"`
	UnreachableThreshold      *float64                   `pulumi:"unreachableThreshold"`
	UseComputedTargets        *bool                      `pulumi:"useComputedTargets"`
	WaitOnComplete            *bool                      `pulumi:"waitOnComplete"`
	WeightedHashBitsForIpv4   *int                       `pulumi:"weightedHashBitsForIpv4"`
	WeightedHashBitsForIpv6   *int                       `pulumi:"weightedHashBitsForIpv6"`
}

type GtmPropertyState struct {
	BackupCname               pulumi.StringPtrInput
	BackupIp                  pulumi.StringPtrInput
	BalanceByDownloadScore    pulumi.BoolPtrInput
	Cname                     pulumi.StringPtrInput
	Comments                  pulumi.StringPtrInput
	Domain                    pulumi.StringPtrInput
	DynamicTtl                pulumi.IntPtrInput
	FailbackDelay             pulumi.IntPtrInput
	FailoverDelay             pulumi.IntPtrInput
	GhostDemandReporting      pulumi.BoolPtrInput
	HandoutLimit              pulumi.IntPtrInput
	HandoutMode               pulumi.StringPtrInput
	HealthMax                 pulumi.Float64PtrInput
	HealthMultiplier          pulumi.Float64PtrInput
	HealthThreshold           pulumi.Float64PtrInput
	Ipv6                      pulumi.BoolPtrInput
	LivenessTests             GtmPropertyLivenessTestArrayInput
	LoadImbalancePercentage   pulumi.Float64PtrInput
	MapName                   pulumi.StringPtrInput
	MaxUnreachablePenalty     pulumi.IntPtrInput
	MinLiveFraction           pulumi.Float64PtrInput
	Name                      pulumi.StringPtrInput
	ScoreAggregationType      pulumi.StringPtrInput
	StaticRrSets              GtmPropertyStaticRrSetArrayInput
	StaticTtl                 pulumi.IntPtrInput
	StickinessBonusConstant   pulumi.IntPtrInput
	StickinessBonusPercentage pulumi.IntPtrInput
	TrafficTargets            GtmPropertyTrafficTargetArrayInput
	Type                      pulumi.StringPtrInput
	UnreachableThreshold      pulumi.Float64PtrInput
	UseComputedTargets        pulumi.BoolPtrInput
	WaitOnComplete            pulumi.BoolPtrInput
	WeightedHashBitsForIpv4   pulumi.IntPtrInput
	WeightedHashBitsForIpv6   pulumi.IntPtrInput
}

func (GtmPropertyState) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmPropertyState)(nil)).Elem()
}

type gtmPropertyArgs struct {
	BackupCname               *string                    `pulumi:"backupCname"`
	BackupIp                  *string                    `pulumi:"backupIp"`
	BalanceByDownloadScore    *bool                      `pulumi:"balanceByDownloadScore"`
	Cname                     *string                    `pulumi:"cname"`
	Comments                  *string                    `pulumi:"comments"`
	Domain                    string                     `pulumi:"domain"`
	DynamicTtl                *int                       `pulumi:"dynamicTtl"`
	FailbackDelay             *int                       `pulumi:"failbackDelay"`
	FailoverDelay             *int                       `pulumi:"failoverDelay"`
	GhostDemandReporting      *bool                      `pulumi:"ghostDemandReporting"`
	HandoutLimit              int                        `pulumi:"handoutLimit"`
	HandoutMode               string                     `pulumi:"handoutMode"`
	HealthMax                 *float64                   `pulumi:"healthMax"`
	HealthMultiplier          *float64                   `pulumi:"healthMultiplier"`
	HealthThreshold           *float64                   `pulumi:"healthThreshold"`
	Ipv6                      *bool                      `pulumi:"ipv6"`
	LivenessTests             []GtmPropertyLivenessTest  `pulumi:"livenessTests"`
	LoadImbalancePercentage   *float64                   `pulumi:"loadImbalancePercentage"`
	MapName                   *string                    `pulumi:"mapName"`
	MaxUnreachablePenalty     *int                       `pulumi:"maxUnreachablePenalty"`
	MinLiveFraction           *float64                   `pulumi:"minLiveFraction"`
	Name                      *string                    `pulumi:"name"`
	ScoreAggregationType      string                     `pulumi:"scoreAggregationType"`
	StaticRrSets              []GtmPropertyStaticRrSet   `pulumi:"staticRrSets"`
	StaticTtl                 *int                       `pulumi:"staticTtl"`
	StickinessBonusConstant   *int                       `pulumi:"stickinessBonusConstant"`
	StickinessBonusPercentage *int                       `pulumi:"stickinessBonusPercentage"`
	TrafficTargets            []GtmPropertyTrafficTarget `pulumi:"trafficTargets"`
	Type                      string                     `pulumi:"type"`
	UnreachableThreshold      *float64                   `pulumi:"unreachableThreshold"`
	UseComputedTargets        *bool                      `pulumi:"useComputedTargets"`
	WaitOnComplete            *bool                      `pulumi:"waitOnComplete"`
}

// The set of arguments for constructing a GtmProperty resource.
type GtmPropertyArgs struct {
	BackupCname               pulumi.StringPtrInput
	BackupIp                  pulumi.StringPtrInput
	BalanceByDownloadScore    pulumi.BoolPtrInput
	Cname                     pulumi.StringPtrInput
	Comments                  pulumi.StringPtrInput
	Domain                    pulumi.StringInput
	DynamicTtl                pulumi.IntPtrInput
	FailbackDelay             pulumi.IntPtrInput
	FailoverDelay             pulumi.IntPtrInput
	GhostDemandReporting      pulumi.BoolPtrInput
	HandoutLimit              pulumi.IntInput
	HandoutMode               pulumi.StringInput
	HealthMax                 pulumi.Float64PtrInput
	HealthMultiplier          pulumi.Float64PtrInput
	HealthThreshold           pulumi.Float64PtrInput
	Ipv6                      pulumi.BoolPtrInput
	LivenessTests             GtmPropertyLivenessTestArrayInput
	LoadImbalancePercentage   pulumi.Float64PtrInput
	MapName                   pulumi.StringPtrInput
	MaxUnreachablePenalty     pulumi.IntPtrInput
	MinLiveFraction           pulumi.Float64PtrInput
	Name                      pulumi.StringPtrInput
	ScoreAggregationType      pulumi.StringInput
	StaticRrSets              GtmPropertyStaticRrSetArrayInput
	StaticTtl                 pulumi.IntPtrInput
	StickinessBonusConstant   pulumi.IntPtrInput
	StickinessBonusPercentage pulumi.IntPtrInput
	TrafficTargets            GtmPropertyTrafficTargetArrayInput
	Type                      pulumi.StringInput
	UnreachableThreshold      pulumi.Float64PtrInput
	UseComputedTargets        pulumi.BoolPtrInput
	WaitOnComplete            pulumi.BoolPtrInput
}

func (GtmPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gtmPropertyArgs)(nil)).Elem()
}

type GtmPropertyInput interface {
	pulumi.Input

	ToGtmPropertyOutput() GtmPropertyOutput
	ToGtmPropertyOutputWithContext(ctx context.Context) GtmPropertyOutput
}

func (*GtmProperty) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmProperty)(nil))
}

func (i *GtmProperty) ToGtmPropertyOutput() GtmPropertyOutput {
	return i.ToGtmPropertyOutputWithContext(context.Background())
}

func (i *GtmProperty) ToGtmPropertyOutputWithContext(ctx context.Context) GtmPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmPropertyOutput)
}

func (i *GtmProperty) ToGtmPropertyPtrOutput() GtmPropertyPtrOutput {
	return i.ToGtmPropertyPtrOutputWithContext(context.Background())
}

func (i *GtmProperty) ToGtmPropertyPtrOutputWithContext(ctx context.Context) GtmPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmPropertyPtrOutput)
}

type GtmPropertyPtrInput interface {
	pulumi.Input

	ToGtmPropertyPtrOutput() GtmPropertyPtrOutput
	ToGtmPropertyPtrOutputWithContext(ctx context.Context) GtmPropertyPtrOutput
}

type gtmPropertyPtrType GtmPropertyArgs

func (*gtmPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmProperty)(nil))
}

func (i *gtmPropertyPtrType) ToGtmPropertyPtrOutput() GtmPropertyPtrOutput {
	return i.ToGtmPropertyPtrOutputWithContext(context.Background())
}

func (i *gtmPropertyPtrType) ToGtmPropertyPtrOutputWithContext(ctx context.Context) GtmPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmPropertyPtrOutput)
}

// GtmPropertyArrayInput is an input type that accepts GtmPropertyArray and GtmPropertyArrayOutput values.
// You can construct a concrete instance of `GtmPropertyArrayInput` via:
//
//          GtmPropertyArray{ GtmPropertyArgs{...} }
type GtmPropertyArrayInput interface {
	pulumi.Input

	ToGtmPropertyArrayOutput() GtmPropertyArrayOutput
	ToGtmPropertyArrayOutputWithContext(context.Context) GtmPropertyArrayOutput
}

type GtmPropertyArray []GtmPropertyInput

func (GtmPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GtmProperty)(nil))
}

func (i GtmPropertyArray) ToGtmPropertyArrayOutput() GtmPropertyArrayOutput {
	return i.ToGtmPropertyArrayOutputWithContext(context.Background())
}

func (i GtmPropertyArray) ToGtmPropertyArrayOutputWithContext(ctx context.Context) GtmPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmPropertyArrayOutput)
}

// GtmPropertyMapInput is an input type that accepts GtmPropertyMap and GtmPropertyMapOutput values.
// You can construct a concrete instance of `GtmPropertyMapInput` via:
//
//          GtmPropertyMap{ "key": GtmPropertyArgs{...} }
type GtmPropertyMapInput interface {
	pulumi.Input

	ToGtmPropertyMapOutput() GtmPropertyMapOutput
	ToGtmPropertyMapOutputWithContext(context.Context) GtmPropertyMapOutput
}

type GtmPropertyMap map[string]GtmPropertyInput

func (GtmPropertyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GtmProperty)(nil))
}

func (i GtmPropertyMap) ToGtmPropertyMapOutput() GtmPropertyMapOutput {
	return i.ToGtmPropertyMapOutputWithContext(context.Background())
}

func (i GtmPropertyMap) ToGtmPropertyMapOutputWithContext(ctx context.Context) GtmPropertyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GtmPropertyMapOutput)
}

type GtmPropertyOutput struct {
	*pulumi.OutputState
}

func (GtmPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GtmProperty)(nil))
}

func (o GtmPropertyOutput) ToGtmPropertyOutput() GtmPropertyOutput {
	return o
}

func (o GtmPropertyOutput) ToGtmPropertyOutputWithContext(ctx context.Context) GtmPropertyOutput {
	return o
}

func (o GtmPropertyOutput) ToGtmPropertyPtrOutput() GtmPropertyPtrOutput {
	return o.ToGtmPropertyPtrOutputWithContext(context.Background())
}

func (o GtmPropertyOutput) ToGtmPropertyPtrOutputWithContext(ctx context.Context) GtmPropertyPtrOutput {
	return o.ApplyT(func(v GtmProperty) *GtmProperty {
		return &v
	}).(GtmPropertyPtrOutput)
}

type GtmPropertyPtrOutput struct {
	*pulumi.OutputState
}

func (GtmPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GtmProperty)(nil))
}

func (o GtmPropertyPtrOutput) ToGtmPropertyPtrOutput() GtmPropertyPtrOutput {
	return o
}

func (o GtmPropertyPtrOutput) ToGtmPropertyPtrOutputWithContext(ctx context.Context) GtmPropertyPtrOutput {
	return o
}

type GtmPropertyArrayOutput struct{ *pulumi.OutputState }

func (GtmPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GtmProperty)(nil))
}

func (o GtmPropertyArrayOutput) ToGtmPropertyArrayOutput() GtmPropertyArrayOutput {
	return o
}

func (o GtmPropertyArrayOutput) ToGtmPropertyArrayOutputWithContext(ctx context.Context) GtmPropertyArrayOutput {
	return o
}

func (o GtmPropertyArrayOutput) Index(i pulumi.IntInput) GtmPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GtmProperty {
		return vs[0].([]GtmProperty)[vs[1].(int)]
	}).(GtmPropertyOutput)
}

type GtmPropertyMapOutput struct{ *pulumi.OutputState }

func (GtmPropertyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GtmProperty)(nil))
}

func (o GtmPropertyMapOutput) ToGtmPropertyMapOutput() GtmPropertyMapOutput {
	return o
}

func (o GtmPropertyMapOutput) ToGtmPropertyMapOutputWithContext(ctx context.Context) GtmPropertyMapOutput {
	return o
}

func (o GtmPropertyMapOutput) MapIndex(k pulumi.StringInput) GtmPropertyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GtmProperty {
		return vs[0].(map[string]GtmProperty)[vs[1].(string)]
	}).(GtmPropertyOutput)
}

func init() {
	pulumi.RegisterOutputType(GtmPropertyOutput{})
	pulumi.RegisterOutputType(GtmPropertyPtrOutput{})
	pulumi.RegisterOutputType(GtmPropertyArrayOutput{})
	pulumi.RegisterOutputType(GtmPropertyMapOutput{})
}
