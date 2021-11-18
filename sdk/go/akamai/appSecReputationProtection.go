// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecReputationProtection` resource to enable or disable reputation protection for a given configuration and security policy.
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
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecReputationProtection(ctx, "protection", &akamai.AppSecReputationProtectionArgs{
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
type AppSecReputationProtection struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// Whether to enable reputation controls: either `true` or `false`.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// A tabular display showing the current protection settings.
	OutputText pulumi.StringOutput `pulumi:"outputText"`
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecReputationProtection registers a new resource with the given unique name, arguments, and options.
func NewAppSecReputationProtection(ctx *pulumi.Context,
	name string, args *AppSecReputationProtectionArgs, opts ...pulumi.ResourceOption) (*AppSecReputationProtection, error) {
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
	var resource AppSecReputationProtection
	err := ctx.RegisterResource("akamai:index/appSecReputationProtection:AppSecReputationProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecReputationProtection gets an existing AppSecReputationProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecReputationProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecReputationProtectionState, opts ...pulumi.ResourceOption) (*AppSecReputationProtection, error) {
	var resource AppSecReputationProtection
	err := ctx.ReadResource("akamai:index/appSecReputationProtection:AppSecReputationProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecReputationProtection resources.
type appSecReputationProtectionState struct {
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// Whether to enable reputation controls: either `true` or `false`.
	Enabled *bool `pulumi:"enabled"`
	// A tabular display showing the current protection settings.
	OutputText *string `pulumi:"outputText"`
	// The ID of the security policy to use.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecReputationProtectionState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// Whether to enable reputation controls: either `true` or `false`.
	Enabled pulumi.BoolPtrInput
	// A tabular display showing the current protection settings.
	OutputText pulumi.StringPtrInput
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecReputationProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecReputationProtectionState)(nil)).Elem()
}

type appSecReputationProtectionArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// Whether to enable reputation controls: either `true` or `false`.
	Enabled bool `pulumi:"enabled"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecReputationProtection resource.
type AppSecReputationProtectionArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// Whether to enable reputation controls: either `true` or `false`.
	Enabled pulumi.BoolInput
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringInput
}

func (AppSecReputationProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecReputationProtectionArgs)(nil)).Elem()
}

type AppSecReputationProtectionInput interface {
	pulumi.Input

	ToAppSecReputationProtectionOutput() AppSecReputationProtectionOutput
	ToAppSecReputationProtectionOutputWithContext(ctx context.Context) AppSecReputationProtectionOutput
}

func (*AppSecReputationProtection) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecReputationProtection)(nil))
}

func (i *AppSecReputationProtection) ToAppSecReputationProtectionOutput() AppSecReputationProtectionOutput {
	return i.ToAppSecReputationProtectionOutputWithContext(context.Background())
}

func (i *AppSecReputationProtection) ToAppSecReputationProtectionOutputWithContext(ctx context.Context) AppSecReputationProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProtectionOutput)
}

func (i *AppSecReputationProtection) ToAppSecReputationProtectionPtrOutput() AppSecReputationProtectionPtrOutput {
	return i.ToAppSecReputationProtectionPtrOutputWithContext(context.Background())
}

func (i *AppSecReputationProtection) ToAppSecReputationProtectionPtrOutputWithContext(ctx context.Context) AppSecReputationProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProtectionPtrOutput)
}

type AppSecReputationProtectionPtrInput interface {
	pulumi.Input

	ToAppSecReputationProtectionPtrOutput() AppSecReputationProtectionPtrOutput
	ToAppSecReputationProtectionPtrOutputWithContext(ctx context.Context) AppSecReputationProtectionPtrOutput
}

type appSecReputationProtectionPtrType AppSecReputationProtectionArgs

func (*appSecReputationProtectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecReputationProtection)(nil))
}

func (i *appSecReputationProtectionPtrType) ToAppSecReputationProtectionPtrOutput() AppSecReputationProtectionPtrOutput {
	return i.ToAppSecReputationProtectionPtrOutputWithContext(context.Background())
}

func (i *appSecReputationProtectionPtrType) ToAppSecReputationProtectionPtrOutputWithContext(ctx context.Context) AppSecReputationProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProtectionPtrOutput)
}

// AppSecReputationProtectionArrayInput is an input type that accepts AppSecReputationProtectionArray and AppSecReputationProtectionArrayOutput values.
// You can construct a concrete instance of `AppSecReputationProtectionArrayInput` via:
//
//          AppSecReputationProtectionArray{ AppSecReputationProtectionArgs{...} }
type AppSecReputationProtectionArrayInput interface {
	pulumi.Input

	ToAppSecReputationProtectionArrayOutput() AppSecReputationProtectionArrayOutput
	ToAppSecReputationProtectionArrayOutputWithContext(context.Context) AppSecReputationProtectionArrayOutput
}

type AppSecReputationProtectionArray []AppSecReputationProtectionInput

func (AppSecReputationProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecReputationProtection)(nil)).Elem()
}

func (i AppSecReputationProtectionArray) ToAppSecReputationProtectionArrayOutput() AppSecReputationProtectionArrayOutput {
	return i.ToAppSecReputationProtectionArrayOutputWithContext(context.Background())
}

func (i AppSecReputationProtectionArray) ToAppSecReputationProtectionArrayOutputWithContext(ctx context.Context) AppSecReputationProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProtectionArrayOutput)
}

// AppSecReputationProtectionMapInput is an input type that accepts AppSecReputationProtectionMap and AppSecReputationProtectionMapOutput values.
// You can construct a concrete instance of `AppSecReputationProtectionMapInput` via:
//
//          AppSecReputationProtectionMap{ "key": AppSecReputationProtectionArgs{...} }
type AppSecReputationProtectionMapInput interface {
	pulumi.Input

	ToAppSecReputationProtectionMapOutput() AppSecReputationProtectionMapOutput
	ToAppSecReputationProtectionMapOutputWithContext(context.Context) AppSecReputationProtectionMapOutput
}

type AppSecReputationProtectionMap map[string]AppSecReputationProtectionInput

func (AppSecReputationProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecReputationProtection)(nil)).Elem()
}

func (i AppSecReputationProtectionMap) ToAppSecReputationProtectionMapOutput() AppSecReputationProtectionMapOutput {
	return i.ToAppSecReputationProtectionMapOutputWithContext(context.Background())
}

func (i AppSecReputationProtectionMap) ToAppSecReputationProtectionMapOutputWithContext(ctx context.Context) AppSecReputationProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProtectionMapOutput)
}

type AppSecReputationProtectionOutput struct{ *pulumi.OutputState }

func (AppSecReputationProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecReputationProtection)(nil))
}

func (o AppSecReputationProtectionOutput) ToAppSecReputationProtectionOutput() AppSecReputationProtectionOutput {
	return o
}

func (o AppSecReputationProtectionOutput) ToAppSecReputationProtectionOutputWithContext(ctx context.Context) AppSecReputationProtectionOutput {
	return o
}

func (o AppSecReputationProtectionOutput) ToAppSecReputationProtectionPtrOutput() AppSecReputationProtectionPtrOutput {
	return o.ToAppSecReputationProtectionPtrOutputWithContext(context.Background())
}

func (o AppSecReputationProtectionOutput) ToAppSecReputationProtectionPtrOutputWithContext(ctx context.Context) AppSecReputationProtectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSecReputationProtection) *AppSecReputationProtection {
		return &v
	}).(AppSecReputationProtectionPtrOutput)
}

type AppSecReputationProtectionPtrOutput struct{ *pulumi.OutputState }

func (AppSecReputationProtectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecReputationProtection)(nil))
}

func (o AppSecReputationProtectionPtrOutput) ToAppSecReputationProtectionPtrOutput() AppSecReputationProtectionPtrOutput {
	return o
}

func (o AppSecReputationProtectionPtrOutput) ToAppSecReputationProtectionPtrOutputWithContext(ctx context.Context) AppSecReputationProtectionPtrOutput {
	return o
}

func (o AppSecReputationProtectionPtrOutput) Elem() AppSecReputationProtectionOutput {
	return o.ApplyT(func(v *AppSecReputationProtection) AppSecReputationProtection {
		if v != nil {
			return *v
		}
		var ret AppSecReputationProtection
		return ret
	}).(AppSecReputationProtectionOutput)
}

type AppSecReputationProtectionArrayOutput struct{ *pulumi.OutputState }

func (AppSecReputationProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecReputationProtection)(nil))
}

func (o AppSecReputationProtectionArrayOutput) ToAppSecReputationProtectionArrayOutput() AppSecReputationProtectionArrayOutput {
	return o
}

func (o AppSecReputationProtectionArrayOutput) ToAppSecReputationProtectionArrayOutputWithContext(ctx context.Context) AppSecReputationProtectionArrayOutput {
	return o
}

func (o AppSecReputationProtectionArrayOutput) Index(i pulumi.IntInput) AppSecReputationProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecReputationProtection {
		return vs[0].([]AppSecReputationProtection)[vs[1].(int)]
	}).(AppSecReputationProtectionOutput)
}

type AppSecReputationProtectionMapOutput struct{ *pulumi.OutputState }

func (AppSecReputationProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecReputationProtection)(nil))
}

func (o AppSecReputationProtectionMapOutput) ToAppSecReputationProtectionMapOutput() AppSecReputationProtectionMapOutput {
	return o
}

func (o AppSecReputationProtectionMapOutput) ToAppSecReputationProtectionMapOutputWithContext(ctx context.Context) AppSecReputationProtectionMapOutput {
	return o
}

func (o AppSecReputationProtectionMapOutput) MapIndex(k pulumi.StringInput) AppSecReputationProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecReputationProtection {
		return vs[0].(map[string]AppSecReputationProtection)[vs[1].(string)]
	}).(AppSecReputationProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProtectionInput)(nil)).Elem(), &AppSecReputationProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProtectionPtrInput)(nil)).Elem(), &AppSecReputationProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProtectionArrayInput)(nil)).Elem(), AppSecReputationProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProtectionMapInput)(nil)).Elem(), AppSecReputationProtectionMap{})
	pulumi.RegisterOutputType(AppSecReputationProtectionOutput{})
	pulumi.RegisterOutputType(AppSecReputationProtectionPtrOutput{})
	pulumi.RegisterOutputType(AppSecReputationProtectionArrayOutput{})
	pulumi.RegisterOutputType(AppSecReputationProtectionMapOutput{})
}
