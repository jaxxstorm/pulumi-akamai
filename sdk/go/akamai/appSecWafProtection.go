// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecWafProtection` resource to enable or disable WAF protection for a given configuration and security policy.
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
// 		opt0 := _var.Security_configuration
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &akamai.LookupAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecWafProtection(ctx, "protection", &akamai.AppSecWafProtectionArgs{
// 			ConfigId:         pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId: pulumi.Any(_var.Security_policy_id),
// 			Enabled:          pulumi.Any(_var.Enabled),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecWafProtection struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// Whether to enable WAF controls: either `true` or `false`.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// A tabular display showing the current protection settings.
	OutputText pulumi.StringOutput `pulumi:"outputText"`
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecWafProtection registers a new resource with the given unique name, arguments, and options.
func NewAppSecWafProtection(ctx *pulumi.Context,
	name string, args *AppSecWafProtectionArgs, opts ...pulumi.ResourceOption) (*AppSecWafProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	var resource AppSecWafProtection
	err := ctx.RegisterResource("akamai:index/appSecWafProtection:AppSecWafProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecWafProtection gets an existing AppSecWafProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecWafProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecWafProtectionState, opts ...pulumi.ResourceOption) (*AppSecWafProtection, error) {
	var resource AppSecWafProtection
	err := ctx.ReadResource("akamai:index/appSecWafProtection:AppSecWafProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecWafProtection resources.
type appSecWafProtectionState struct {
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// Whether to enable WAF controls: either `true` or `false`.
	Enabled *bool `pulumi:"enabled"`
	// A tabular display showing the current protection settings.
	OutputText *string `pulumi:"outputText"`
	// The ID of the security policy to use.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecWafProtectionState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// Whether to enable WAF controls: either `true` or `false`.
	Enabled pulumi.BoolPtrInput
	// A tabular display showing the current protection settings.
	OutputText pulumi.StringPtrInput
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecWafProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecWafProtectionState)(nil)).Elem()
}

type appSecWafProtectionArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// Whether to enable WAF controls: either `true` or `false`.
	Enabled bool `pulumi:"enabled"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecWafProtection resource.
type AppSecWafProtectionArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// Whether to enable WAF controls: either `true` or `false`.
	Enabled pulumi.BoolInput
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringInput
}

func (AppSecWafProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecWafProtectionArgs)(nil)).Elem()
}

type AppSecWafProtectionInput interface {
	pulumi.Input

	ToAppSecWafProtectionOutput() AppSecWafProtectionOutput
	ToAppSecWafProtectionOutputWithContext(ctx context.Context) AppSecWafProtectionOutput
}

func (*AppSecWafProtection) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecWafProtection)(nil))
}

func (i *AppSecWafProtection) ToAppSecWafProtectionOutput() AppSecWafProtectionOutput {
	return i.ToAppSecWafProtectionOutputWithContext(context.Background())
}

func (i *AppSecWafProtection) ToAppSecWafProtectionOutputWithContext(ctx context.Context) AppSecWafProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecWafProtectionOutput)
}

func (i *AppSecWafProtection) ToAppSecWafProtectionPtrOutput() AppSecWafProtectionPtrOutput {
	return i.ToAppSecWafProtectionPtrOutputWithContext(context.Background())
}

func (i *AppSecWafProtection) ToAppSecWafProtectionPtrOutputWithContext(ctx context.Context) AppSecWafProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecWafProtectionPtrOutput)
}

type AppSecWafProtectionPtrInput interface {
	pulumi.Input

	ToAppSecWafProtectionPtrOutput() AppSecWafProtectionPtrOutput
	ToAppSecWafProtectionPtrOutputWithContext(ctx context.Context) AppSecWafProtectionPtrOutput
}

type appSecWafProtectionPtrType AppSecWafProtectionArgs

func (*appSecWafProtectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecWafProtection)(nil))
}

func (i *appSecWafProtectionPtrType) ToAppSecWafProtectionPtrOutput() AppSecWafProtectionPtrOutput {
	return i.ToAppSecWafProtectionPtrOutputWithContext(context.Background())
}

func (i *appSecWafProtectionPtrType) ToAppSecWafProtectionPtrOutputWithContext(ctx context.Context) AppSecWafProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecWafProtectionPtrOutput)
}

// AppSecWafProtectionArrayInput is an input type that accepts AppSecWafProtectionArray and AppSecWafProtectionArrayOutput values.
// You can construct a concrete instance of `AppSecWafProtectionArrayInput` via:
//
//          AppSecWafProtectionArray{ AppSecWafProtectionArgs{...} }
type AppSecWafProtectionArrayInput interface {
	pulumi.Input

	ToAppSecWafProtectionArrayOutput() AppSecWafProtectionArrayOutput
	ToAppSecWafProtectionArrayOutputWithContext(context.Context) AppSecWafProtectionArrayOutput
}

type AppSecWafProtectionArray []AppSecWafProtectionInput

func (AppSecWafProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecWafProtection)(nil))
}

func (i AppSecWafProtectionArray) ToAppSecWafProtectionArrayOutput() AppSecWafProtectionArrayOutput {
	return i.ToAppSecWafProtectionArrayOutputWithContext(context.Background())
}

func (i AppSecWafProtectionArray) ToAppSecWafProtectionArrayOutputWithContext(ctx context.Context) AppSecWafProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecWafProtectionArrayOutput)
}

// AppSecWafProtectionMapInput is an input type that accepts AppSecWafProtectionMap and AppSecWafProtectionMapOutput values.
// You can construct a concrete instance of `AppSecWafProtectionMapInput` via:
//
//          AppSecWafProtectionMap{ "key": AppSecWafProtectionArgs{...} }
type AppSecWafProtectionMapInput interface {
	pulumi.Input

	ToAppSecWafProtectionMapOutput() AppSecWafProtectionMapOutput
	ToAppSecWafProtectionMapOutputWithContext(context.Context) AppSecWafProtectionMapOutput
}

type AppSecWafProtectionMap map[string]AppSecWafProtectionInput

func (AppSecWafProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecWafProtection)(nil))
}

func (i AppSecWafProtectionMap) ToAppSecWafProtectionMapOutput() AppSecWafProtectionMapOutput {
	return i.ToAppSecWafProtectionMapOutputWithContext(context.Background())
}

func (i AppSecWafProtectionMap) ToAppSecWafProtectionMapOutputWithContext(ctx context.Context) AppSecWafProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecWafProtectionMapOutput)
}

type AppSecWafProtectionOutput struct {
	*pulumi.OutputState
}

func (AppSecWafProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecWafProtection)(nil))
}

func (o AppSecWafProtectionOutput) ToAppSecWafProtectionOutput() AppSecWafProtectionOutput {
	return o
}

func (o AppSecWafProtectionOutput) ToAppSecWafProtectionOutputWithContext(ctx context.Context) AppSecWafProtectionOutput {
	return o
}

func (o AppSecWafProtectionOutput) ToAppSecWafProtectionPtrOutput() AppSecWafProtectionPtrOutput {
	return o.ToAppSecWafProtectionPtrOutputWithContext(context.Background())
}

func (o AppSecWafProtectionOutput) ToAppSecWafProtectionPtrOutputWithContext(ctx context.Context) AppSecWafProtectionPtrOutput {
	return o.ApplyT(func(v AppSecWafProtection) *AppSecWafProtection {
		return &v
	}).(AppSecWafProtectionPtrOutput)
}

type AppSecWafProtectionPtrOutput struct {
	*pulumi.OutputState
}

func (AppSecWafProtectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecWafProtection)(nil))
}

func (o AppSecWafProtectionPtrOutput) ToAppSecWafProtectionPtrOutput() AppSecWafProtectionPtrOutput {
	return o
}

func (o AppSecWafProtectionPtrOutput) ToAppSecWafProtectionPtrOutputWithContext(ctx context.Context) AppSecWafProtectionPtrOutput {
	return o
}

type AppSecWafProtectionArrayOutput struct{ *pulumi.OutputState }

func (AppSecWafProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecWafProtection)(nil))
}

func (o AppSecWafProtectionArrayOutput) ToAppSecWafProtectionArrayOutput() AppSecWafProtectionArrayOutput {
	return o
}

func (o AppSecWafProtectionArrayOutput) ToAppSecWafProtectionArrayOutputWithContext(ctx context.Context) AppSecWafProtectionArrayOutput {
	return o
}

func (o AppSecWafProtectionArrayOutput) Index(i pulumi.IntInput) AppSecWafProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecWafProtection {
		return vs[0].([]AppSecWafProtection)[vs[1].(int)]
	}).(AppSecWafProtectionOutput)
}

type AppSecWafProtectionMapOutput struct{ *pulumi.OutputState }

func (AppSecWafProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecWafProtection)(nil))
}

func (o AppSecWafProtectionMapOutput) ToAppSecWafProtectionMapOutput() AppSecWafProtectionMapOutput {
	return o
}

func (o AppSecWafProtectionMapOutput) ToAppSecWafProtectionMapOutputWithContext(ctx context.Context) AppSecWafProtectionMapOutput {
	return o
}

func (o AppSecWafProtectionMapOutput) MapIndex(k pulumi.StringInput) AppSecWafProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecWafProtection {
		return vs[0].(map[string]AppSecWafProtection)[vs[1].(string)]
	}).(AppSecWafProtectionOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecWafProtectionOutput{})
	pulumi.RegisterOutputType(AppSecWafProtectionPtrOutput{})
	pulumi.RegisterOutputType(AppSecWafProtectionArrayOutput{})
	pulumi.RegisterOutputType(AppSecWafProtectionMapOutput{})
}
