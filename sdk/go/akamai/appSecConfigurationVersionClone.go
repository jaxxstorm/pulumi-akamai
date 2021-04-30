// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `AppSecConfigurationVersionClone` resource allows you to create a new version of a security configuration by cloning an existing version.
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
// 		opt0 := "Akamai Tools"
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &akamai.LookupAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		clone, err := akamai.NewAppSecConfigurationVersionClone(ctx, "clone", &akamai.AppSecConfigurationVersionCloneArgs{
// 			ConfigId:          pulumi.Int(configuration.ConfigId),
// 			CreateFromVersion: pulumi.Int(configuration.LatestVersion),
// 			RuleUpdate:        pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloneVersion", clone.Version)
// 		return nil
// 	})
// }
// ```
type AppSecConfigurationVersionClone struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// The version number of the security configuration to clone.
	CreateFromVersion pulumi.IntOutput `pulumi:"createFromVersion"`
	// A boolean indicating whether to update the rules of the new version. If not supplied, False is assumed.
	RuleUpdate pulumi.BoolPtrOutput `pulumi:"ruleUpdate"`
	// The number of the cloned version.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewAppSecConfigurationVersionClone registers a new resource with the given unique name, arguments, and options.
func NewAppSecConfigurationVersionClone(ctx *pulumi.Context,
	name string, args *AppSecConfigurationVersionCloneArgs, opts ...pulumi.ResourceOption) (*AppSecConfigurationVersionClone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.CreateFromVersion == nil {
		return nil, errors.New("invalid value for required argument 'CreateFromVersion'")
	}
	var resource AppSecConfigurationVersionClone
	err := ctx.RegisterResource("akamai:index/appSecConfigurationVersionClone:AppSecConfigurationVersionClone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecConfigurationVersionClone gets an existing AppSecConfigurationVersionClone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecConfigurationVersionClone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecConfigurationVersionCloneState, opts ...pulumi.ResourceOption) (*AppSecConfigurationVersionClone, error) {
	var resource AppSecConfigurationVersionClone
	err := ctx.ReadResource("akamai:index/appSecConfigurationVersionClone:AppSecConfigurationVersionClone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecConfigurationVersionClone resources.
type appSecConfigurationVersionCloneState struct {
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// The version number of the security configuration to clone.
	CreateFromVersion *int `pulumi:"createFromVersion"`
	// A boolean indicating whether to update the rules of the new version. If not supplied, False is assumed.
	RuleUpdate *bool `pulumi:"ruleUpdate"`
	// The number of the cloned version.
	Version *int `pulumi:"version"`
}

type AppSecConfigurationVersionCloneState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// The version number of the security configuration to clone.
	CreateFromVersion pulumi.IntPtrInput
	// A boolean indicating whether to update the rules of the new version. If not supplied, False is assumed.
	RuleUpdate pulumi.BoolPtrInput
	// The number of the cloned version.
	Version pulumi.IntPtrInput
}

func (AppSecConfigurationVersionCloneState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecConfigurationVersionCloneState)(nil)).Elem()
}

type appSecConfigurationVersionCloneArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The version number of the security configuration to clone.
	CreateFromVersion int `pulumi:"createFromVersion"`
	// A boolean indicating whether to update the rules of the new version. If not supplied, False is assumed.
	RuleUpdate *bool `pulumi:"ruleUpdate"`
}

// The set of arguments for constructing a AppSecConfigurationVersionClone resource.
type AppSecConfigurationVersionCloneArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// The version number of the security configuration to clone.
	CreateFromVersion pulumi.IntInput
	// A boolean indicating whether to update the rules of the new version. If not supplied, False is assumed.
	RuleUpdate pulumi.BoolPtrInput
}

func (AppSecConfigurationVersionCloneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecConfigurationVersionCloneArgs)(nil)).Elem()
}

type AppSecConfigurationVersionCloneInput interface {
	pulumi.Input

	ToAppSecConfigurationVersionCloneOutput() AppSecConfigurationVersionCloneOutput
	ToAppSecConfigurationVersionCloneOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneOutput
}

func (*AppSecConfigurationVersionClone) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecConfigurationVersionClone)(nil))
}

func (i *AppSecConfigurationVersionClone) ToAppSecConfigurationVersionCloneOutput() AppSecConfigurationVersionCloneOutput {
	return i.ToAppSecConfigurationVersionCloneOutputWithContext(context.Background())
}

func (i *AppSecConfigurationVersionClone) ToAppSecConfigurationVersionCloneOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationVersionCloneOutput)
}

func (i *AppSecConfigurationVersionClone) ToAppSecConfigurationVersionClonePtrOutput() AppSecConfigurationVersionClonePtrOutput {
	return i.ToAppSecConfigurationVersionClonePtrOutputWithContext(context.Background())
}

func (i *AppSecConfigurationVersionClone) ToAppSecConfigurationVersionClonePtrOutputWithContext(ctx context.Context) AppSecConfigurationVersionClonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationVersionClonePtrOutput)
}

type AppSecConfigurationVersionClonePtrInput interface {
	pulumi.Input

	ToAppSecConfigurationVersionClonePtrOutput() AppSecConfigurationVersionClonePtrOutput
	ToAppSecConfigurationVersionClonePtrOutputWithContext(ctx context.Context) AppSecConfigurationVersionClonePtrOutput
}

type appSecConfigurationVersionClonePtrType AppSecConfigurationVersionCloneArgs

func (*appSecConfigurationVersionClonePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecConfigurationVersionClone)(nil))
}

func (i *appSecConfigurationVersionClonePtrType) ToAppSecConfigurationVersionClonePtrOutput() AppSecConfigurationVersionClonePtrOutput {
	return i.ToAppSecConfigurationVersionClonePtrOutputWithContext(context.Background())
}

func (i *appSecConfigurationVersionClonePtrType) ToAppSecConfigurationVersionClonePtrOutputWithContext(ctx context.Context) AppSecConfigurationVersionClonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationVersionClonePtrOutput)
}

// AppSecConfigurationVersionCloneArrayInput is an input type that accepts AppSecConfigurationVersionCloneArray and AppSecConfigurationVersionCloneArrayOutput values.
// You can construct a concrete instance of `AppSecConfigurationVersionCloneArrayInput` via:
//
//          AppSecConfigurationVersionCloneArray{ AppSecConfigurationVersionCloneArgs{...} }
type AppSecConfigurationVersionCloneArrayInput interface {
	pulumi.Input

	ToAppSecConfigurationVersionCloneArrayOutput() AppSecConfigurationVersionCloneArrayOutput
	ToAppSecConfigurationVersionCloneArrayOutputWithContext(context.Context) AppSecConfigurationVersionCloneArrayOutput
}

type AppSecConfigurationVersionCloneArray []AppSecConfigurationVersionCloneInput

func (AppSecConfigurationVersionCloneArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecConfigurationVersionClone)(nil))
}

func (i AppSecConfigurationVersionCloneArray) ToAppSecConfigurationVersionCloneArrayOutput() AppSecConfigurationVersionCloneArrayOutput {
	return i.ToAppSecConfigurationVersionCloneArrayOutputWithContext(context.Background())
}

func (i AppSecConfigurationVersionCloneArray) ToAppSecConfigurationVersionCloneArrayOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationVersionCloneArrayOutput)
}

// AppSecConfigurationVersionCloneMapInput is an input type that accepts AppSecConfigurationVersionCloneMap and AppSecConfigurationVersionCloneMapOutput values.
// You can construct a concrete instance of `AppSecConfigurationVersionCloneMapInput` via:
//
//          AppSecConfigurationVersionCloneMap{ "key": AppSecConfigurationVersionCloneArgs{...} }
type AppSecConfigurationVersionCloneMapInput interface {
	pulumi.Input

	ToAppSecConfigurationVersionCloneMapOutput() AppSecConfigurationVersionCloneMapOutput
	ToAppSecConfigurationVersionCloneMapOutputWithContext(context.Context) AppSecConfigurationVersionCloneMapOutput
}

type AppSecConfigurationVersionCloneMap map[string]AppSecConfigurationVersionCloneInput

func (AppSecConfigurationVersionCloneMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecConfigurationVersionClone)(nil))
}

func (i AppSecConfigurationVersionCloneMap) ToAppSecConfigurationVersionCloneMapOutput() AppSecConfigurationVersionCloneMapOutput {
	return i.ToAppSecConfigurationVersionCloneMapOutputWithContext(context.Background())
}

func (i AppSecConfigurationVersionCloneMap) ToAppSecConfigurationVersionCloneMapOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationVersionCloneMapOutput)
}

type AppSecConfigurationVersionCloneOutput struct {
	*pulumi.OutputState
}

func (AppSecConfigurationVersionCloneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecConfigurationVersionClone)(nil))
}

func (o AppSecConfigurationVersionCloneOutput) ToAppSecConfigurationVersionCloneOutput() AppSecConfigurationVersionCloneOutput {
	return o
}

func (o AppSecConfigurationVersionCloneOutput) ToAppSecConfigurationVersionCloneOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneOutput {
	return o
}

func (o AppSecConfigurationVersionCloneOutput) ToAppSecConfigurationVersionClonePtrOutput() AppSecConfigurationVersionClonePtrOutput {
	return o.ToAppSecConfigurationVersionClonePtrOutputWithContext(context.Background())
}

func (o AppSecConfigurationVersionCloneOutput) ToAppSecConfigurationVersionClonePtrOutputWithContext(ctx context.Context) AppSecConfigurationVersionClonePtrOutput {
	return o.ApplyT(func(v AppSecConfigurationVersionClone) *AppSecConfigurationVersionClone {
		return &v
	}).(AppSecConfigurationVersionClonePtrOutput)
}

type AppSecConfigurationVersionClonePtrOutput struct {
	*pulumi.OutputState
}

func (AppSecConfigurationVersionClonePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecConfigurationVersionClone)(nil))
}

func (o AppSecConfigurationVersionClonePtrOutput) ToAppSecConfigurationVersionClonePtrOutput() AppSecConfigurationVersionClonePtrOutput {
	return o
}

func (o AppSecConfigurationVersionClonePtrOutput) ToAppSecConfigurationVersionClonePtrOutputWithContext(ctx context.Context) AppSecConfigurationVersionClonePtrOutput {
	return o
}

type AppSecConfigurationVersionCloneArrayOutput struct{ *pulumi.OutputState }

func (AppSecConfigurationVersionCloneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecConfigurationVersionClone)(nil))
}

func (o AppSecConfigurationVersionCloneArrayOutput) ToAppSecConfigurationVersionCloneArrayOutput() AppSecConfigurationVersionCloneArrayOutput {
	return o
}

func (o AppSecConfigurationVersionCloneArrayOutput) ToAppSecConfigurationVersionCloneArrayOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneArrayOutput {
	return o
}

func (o AppSecConfigurationVersionCloneArrayOutput) Index(i pulumi.IntInput) AppSecConfigurationVersionCloneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecConfigurationVersionClone {
		return vs[0].([]AppSecConfigurationVersionClone)[vs[1].(int)]
	}).(AppSecConfigurationVersionCloneOutput)
}

type AppSecConfigurationVersionCloneMapOutput struct{ *pulumi.OutputState }

func (AppSecConfigurationVersionCloneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecConfigurationVersionClone)(nil))
}

func (o AppSecConfigurationVersionCloneMapOutput) ToAppSecConfigurationVersionCloneMapOutput() AppSecConfigurationVersionCloneMapOutput {
	return o
}

func (o AppSecConfigurationVersionCloneMapOutput) ToAppSecConfigurationVersionCloneMapOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneMapOutput {
	return o
}

func (o AppSecConfigurationVersionCloneMapOutput) MapIndex(k pulumi.StringInput) AppSecConfigurationVersionCloneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecConfigurationVersionClone {
		return vs[0].(map[string]AppSecConfigurationVersionClone)[vs[1].(string)]
	}).(AppSecConfigurationVersionCloneOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecConfigurationVersionCloneOutput{})
	pulumi.RegisterOutputType(AppSecConfigurationVersionClonePtrOutput{})
	pulumi.RegisterOutputType(AppSecConfigurationVersionCloneArrayOutput{})
	pulumi.RegisterOutputType(AppSecConfigurationVersionCloneMapOutput{})
}
