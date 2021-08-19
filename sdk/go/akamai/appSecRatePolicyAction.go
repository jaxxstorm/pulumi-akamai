// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `resourceAkamaiAppsecRatePolicyAction` resource allows you to create, modify or delete the actions in a rate policy.
type AppSecRatePolicyAction struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// The ipv4 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv4Action pulumi.StringOutput `pulumi:"ipv4Action"`
	// The ipv6 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv6Action pulumi.StringOutput `pulumi:"ipv6Action"`
	// The ID of the rate policy to use.
	RatePolicyId     pulumi.IntOutput    `pulumi:"ratePolicyId"`
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecRatePolicyAction registers a new resource with the given unique name, arguments, and options.
func NewAppSecRatePolicyAction(ctx *pulumi.Context,
	name string, args *AppSecRatePolicyActionArgs, opts ...pulumi.ResourceOption) (*AppSecRatePolicyAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.Ipv4Action == nil {
		return nil, errors.New("invalid value for required argument 'Ipv4Action'")
	}
	if args.Ipv6Action == nil {
		return nil, errors.New("invalid value for required argument 'Ipv6Action'")
	}
	if args.RatePolicyId == nil {
		return nil, errors.New("invalid value for required argument 'RatePolicyId'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	var resource AppSecRatePolicyAction
	err := ctx.RegisterResource("akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecRatePolicyAction gets an existing AppSecRatePolicyAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecRatePolicyAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecRatePolicyActionState, opts ...pulumi.ResourceOption) (*AppSecRatePolicyAction, error) {
	var resource AppSecRatePolicyAction
	err := ctx.ReadResource("akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecRatePolicyAction resources.
type appSecRatePolicyActionState struct {
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// The ipv4 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv4Action *string `pulumi:"ipv4Action"`
	// The ipv6 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv6Action *string `pulumi:"ipv6Action"`
	// The ID of the rate policy to use.
	RatePolicyId     *int    `pulumi:"ratePolicyId"`
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecRatePolicyActionState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// The ipv4 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv4Action pulumi.StringPtrInput
	// The ipv6 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv6Action pulumi.StringPtrInput
	// The ID of the rate policy to use.
	RatePolicyId     pulumi.IntPtrInput
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecRatePolicyActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRatePolicyActionState)(nil)).Elem()
}

type appSecRatePolicyActionArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ipv4 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv4Action string `pulumi:"ipv4Action"`
	// The ipv6 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv6Action string `pulumi:"ipv6Action"`
	// The ID of the rate policy to use.
	RatePolicyId     int    `pulumi:"ratePolicyId"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecRatePolicyAction resource.
type AppSecRatePolicyActionArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// The ipv4 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv4Action pulumi.StringInput
	// The ipv6 action to assign to this rate policy, either `alert`, `deny`, `deny_custom_{custom_deny_id}`, or `none`. If the action is none, the rate policy is inactive in the policy.
	Ipv6Action pulumi.StringInput
	// The ID of the rate policy to use.
	RatePolicyId     pulumi.IntInput
	SecurityPolicyId pulumi.StringInput
}

func (AppSecRatePolicyActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRatePolicyActionArgs)(nil)).Elem()
}

type AppSecRatePolicyActionInput interface {
	pulumi.Input

	ToAppSecRatePolicyActionOutput() AppSecRatePolicyActionOutput
	ToAppSecRatePolicyActionOutputWithContext(ctx context.Context) AppSecRatePolicyActionOutput
}

func (*AppSecRatePolicyAction) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecRatePolicyAction)(nil))
}

func (i *AppSecRatePolicyAction) ToAppSecRatePolicyActionOutput() AppSecRatePolicyActionOutput {
	return i.ToAppSecRatePolicyActionOutputWithContext(context.Background())
}

func (i *AppSecRatePolicyAction) ToAppSecRatePolicyActionOutputWithContext(ctx context.Context) AppSecRatePolicyActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyActionOutput)
}

func (i *AppSecRatePolicyAction) ToAppSecRatePolicyActionPtrOutput() AppSecRatePolicyActionPtrOutput {
	return i.ToAppSecRatePolicyActionPtrOutputWithContext(context.Background())
}

func (i *AppSecRatePolicyAction) ToAppSecRatePolicyActionPtrOutputWithContext(ctx context.Context) AppSecRatePolicyActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyActionPtrOutput)
}

type AppSecRatePolicyActionPtrInput interface {
	pulumi.Input

	ToAppSecRatePolicyActionPtrOutput() AppSecRatePolicyActionPtrOutput
	ToAppSecRatePolicyActionPtrOutputWithContext(ctx context.Context) AppSecRatePolicyActionPtrOutput
}

type appSecRatePolicyActionPtrType AppSecRatePolicyActionArgs

func (*appSecRatePolicyActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRatePolicyAction)(nil))
}

func (i *appSecRatePolicyActionPtrType) ToAppSecRatePolicyActionPtrOutput() AppSecRatePolicyActionPtrOutput {
	return i.ToAppSecRatePolicyActionPtrOutputWithContext(context.Background())
}

func (i *appSecRatePolicyActionPtrType) ToAppSecRatePolicyActionPtrOutputWithContext(ctx context.Context) AppSecRatePolicyActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyActionPtrOutput)
}

// AppSecRatePolicyActionArrayInput is an input type that accepts AppSecRatePolicyActionArray and AppSecRatePolicyActionArrayOutput values.
// You can construct a concrete instance of `AppSecRatePolicyActionArrayInput` via:
//
//          AppSecRatePolicyActionArray{ AppSecRatePolicyActionArgs{...} }
type AppSecRatePolicyActionArrayInput interface {
	pulumi.Input

	ToAppSecRatePolicyActionArrayOutput() AppSecRatePolicyActionArrayOutput
	ToAppSecRatePolicyActionArrayOutputWithContext(context.Context) AppSecRatePolicyActionArrayOutput
}

type AppSecRatePolicyActionArray []AppSecRatePolicyActionInput

func (AppSecRatePolicyActionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecRatePolicyAction)(nil))
}

func (i AppSecRatePolicyActionArray) ToAppSecRatePolicyActionArrayOutput() AppSecRatePolicyActionArrayOutput {
	return i.ToAppSecRatePolicyActionArrayOutputWithContext(context.Background())
}

func (i AppSecRatePolicyActionArray) ToAppSecRatePolicyActionArrayOutputWithContext(ctx context.Context) AppSecRatePolicyActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyActionArrayOutput)
}

// AppSecRatePolicyActionMapInput is an input type that accepts AppSecRatePolicyActionMap and AppSecRatePolicyActionMapOutput values.
// You can construct a concrete instance of `AppSecRatePolicyActionMapInput` via:
//
//          AppSecRatePolicyActionMap{ "key": AppSecRatePolicyActionArgs{...} }
type AppSecRatePolicyActionMapInput interface {
	pulumi.Input

	ToAppSecRatePolicyActionMapOutput() AppSecRatePolicyActionMapOutput
	ToAppSecRatePolicyActionMapOutputWithContext(context.Context) AppSecRatePolicyActionMapOutput
}

type AppSecRatePolicyActionMap map[string]AppSecRatePolicyActionInput

func (AppSecRatePolicyActionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecRatePolicyAction)(nil))
}

func (i AppSecRatePolicyActionMap) ToAppSecRatePolicyActionMapOutput() AppSecRatePolicyActionMapOutput {
	return i.ToAppSecRatePolicyActionMapOutputWithContext(context.Background())
}

func (i AppSecRatePolicyActionMap) ToAppSecRatePolicyActionMapOutputWithContext(ctx context.Context) AppSecRatePolicyActionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyActionMapOutput)
}

type AppSecRatePolicyActionOutput struct {
	*pulumi.OutputState
}

func (AppSecRatePolicyActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecRatePolicyAction)(nil))
}

func (o AppSecRatePolicyActionOutput) ToAppSecRatePolicyActionOutput() AppSecRatePolicyActionOutput {
	return o
}

func (o AppSecRatePolicyActionOutput) ToAppSecRatePolicyActionOutputWithContext(ctx context.Context) AppSecRatePolicyActionOutput {
	return o
}

func (o AppSecRatePolicyActionOutput) ToAppSecRatePolicyActionPtrOutput() AppSecRatePolicyActionPtrOutput {
	return o.ToAppSecRatePolicyActionPtrOutputWithContext(context.Background())
}

func (o AppSecRatePolicyActionOutput) ToAppSecRatePolicyActionPtrOutputWithContext(ctx context.Context) AppSecRatePolicyActionPtrOutput {
	return o.ApplyT(func(v AppSecRatePolicyAction) *AppSecRatePolicyAction {
		return &v
	}).(AppSecRatePolicyActionPtrOutput)
}

type AppSecRatePolicyActionPtrOutput struct {
	*pulumi.OutputState
}

func (AppSecRatePolicyActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRatePolicyAction)(nil))
}

func (o AppSecRatePolicyActionPtrOutput) ToAppSecRatePolicyActionPtrOutput() AppSecRatePolicyActionPtrOutput {
	return o
}

func (o AppSecRatePolicyActionPtrOutput) ToAppSecRatePolicyActionPtrOutputWithContext(ctx context.Context) AppSecRatePolicyActionPtrOutput {
	return o
}

type AppSecRatePolicyActionArrayOutput struct{ *pulumi.OutputState }

func (AppSecRatePolicyActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecRatePolicyAction)(nil))
}

func (o AppSecRatePolicyActionArrayOutput) ToAppSecRatePolicyActionArrayOutput() AppSecRatePolicyActionArrayOutput {
	return o
}

func (o AppSecRatePolicyActionArrayOutput) ToAppSecRatePolicyActionArrayOutputWithContext(ctx context.Context) AppSecRatePolicyActionArrayOutput {
	return o
}

func (o AppSecRatePolicyActionArrayOutput) Index(i pulumi.IntInput) AppSecRatePolicyActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecRatePolicyAction {
		return vs[0].([]AppSecRatePolicyAction)[vs[1].(int)]
	}).(AppSecRatePolicyActionOutput)
}

type AppSecRatePolicyActionMapOutput struct{ *pulumi.OutputState }

func (AppSecRatePolicyActionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecRatePolicyAction)(nil))
}

func (o AppSecRatePolicyActionMapOutput) ToAppSecRatePolicyActionMapOutput() AppSecRatePolicyActionMapOutput {
	return o
}

func (o AppSecRatePolicyActionMapOutput) ToAppSecRatePolicyActionMapOutputWithContext(ctx context.Context) AppSecRatePolicyActionMapOutput {
	return o
}

func (o AppSecRatePolicyActionMapOutput) MapIndex(k pulumi.StringInput) AppSecRatePolicyActionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecRatePolicyAction {
		return vs[0].(map[string]AppSecRatePolicyAction)[vs[1].(string)]
	}).(AppSecRatePolicyActionOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecRatePolicyActionOutput{})
	pulumi.RegisterOutputType(AppSecRatePolicyActionPtrOutput{})
	pulumi.RegisterOutputType(AppSecRatePolicyActionArrayOutput{})
	pulumi.RegisterOutputType(AppSecRatePolicyActionMapOutput{})
}