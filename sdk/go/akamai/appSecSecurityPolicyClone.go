// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `AppSecSecurityPolicyClone` resource allows you to create a new security policy by cloning an existing policy.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Akamai Tools"
// 		configuration, err := akamai.GetAppSecConfiguration(ctx, &akamai.GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		securityPolicyCloneAppSecSecurityPolicyClone, err := akamai.NewAppSecSecurityPolicyClone(ctx, "securityPolicyCloneAppSecSecurityPolicyClone", &akamai.AppSecSecurityPolicyCloneArgs{
// 			ConfigId:                 pulumi.Int(configuration.ConfigId),
// 			Version:                  pulumi.Int(configuration.LatestVersion),
// 			CreateFromSecurityPolicy: pulumi.String("crAP_75829"),
// 			PolicyName:               pulumi.String("Test Policy 1"),
// 			PolicyPrefix:             pulumi.String("TP1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("securityPolicyClone", securityPolicyCloneAppSecSecurityPolicyClone.PolicyId)
// 		return nil
// 	})
// }
// ```
type AppSecSecurityPolicyClone struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// The ID of the security policy to clone.
	CreateFromSecurityPolicy pulumi.StringOutput `pulumi:"createFromSecurityPolicy"`
	// The ID of the new security policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
	PolicyName pulumi.StringPtrOutput `pulumi:"policyName"`
	// The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
	PolicyPrefix pulumi.StringPtrOutput `pulumi:"policyPrefix"`
	Version      pulumi.IntOutput       `pulumi:"version"`
}

// NewAppSecSecurityPolicyClone registers a new resource with the given unique name, arguments, and options.
func NewAppSecSecurityPolicyClone(ctx *pulumi.Context,
	name string, args *AppSecSecurityPolicyCloneArgs, opts ...pulumi.ResourceOption) (*AppSecSecurityPolicyClone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.CreateFromSecurityPolicy == nil {
		return nil, errors.New("invalid value for required argument 'CreateFromSecurityPolicy'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource AppSecSecurityPolicyClone
	err := ctx.RegisterResource("akamai:index/appSecSecurityPolicyClone:AppSecSecurityPolicyClone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecSecurityPolicyClone gets an existing AppSecSecurityPolicyClone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecSecurityPolicyClone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecSecurityPolicyCloneState, opts ...pulumi.ResourceOption) (*AppSecSecurityPolicyClone, error) {
	var resource AppSecSecurityPolicyClone
	err := ctx.ReadResource("akamai:index/appSecSecurityPolicyClone:AppSecSecurityPolicyClone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecSecurityPolicyClone resources.
type appSecSecurityPolicyCloneState struct {
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// The ID of the security policy to clone.
	CreateFromSecurityPolicy *string `pulumi:"createFromSecurityPolicy"`
	// The ID of the new security policy.
	PolicyId *string `pulumi:"policyId"`
	// The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
	PolicyName *string `pulumi:"policyName"`
	// The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
	PolicyPrefix *string `pulumi:"policyPrefix"`
	Version      *int    `pulumi:"version"`
}

type AppSecSecurityPolicyCloneState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// The ID of the security policy to clone.
	CreateFromSecurityPolicy pulumi.StringPtrInput
	// The ID of the new security policy.
	PolicyId pulumi.StringPtrInput
	// The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
	PolicyName pulumi.StringPtrInput
	// The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
	PolicyPrefix pulumi.StringPtrInput
	Version      pulumi.IntPtrInput
}

func (AppSecSecurityPolicyCloneState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSecurityPolicyCloneState)(nil)).Elem()
}

type appSecSecurityPolicyCloneArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the security policy to clone.
	CreateFromSecurityPolicy string `pulumi:"createFromSecurityPolicy"`
	// The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
	PolicyName *string `pulumi:"policyName"`
	// The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
	PolicyPrefix *string `pulumi:"policyPrefix"`
	Version      int     `pulumi:"version"`
}

// The set of arguments for constructing a AppSecSecurityPolicyClone resource.
type AppSecSecurityPolicyCloneArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// The ID of the security policy to clone.
	CreateFromSecurityPolicy pulumi.StringInput
	// The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
	PolicyName pulumi.StringPtrInput
	// The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
	PolicyPrefix pulumi.StringPtrInput
	Version      pulumi.IntInput
}

func (AppSecSecurityPolicyCloneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSecurityPolicyCloneArgs)(nil)).Elem()
}

type AppSecSecurityPolicyCloneInput interface {
	pulumi.Input

	ToAppSecSecurityPolicyCloneOutput() AppSecSecurityPolicyCloneOutput
	ToAppSecSecurityPolicyCloneOutputWithContext(ctx context.Context) AppSecSecurityPolicyCloneOutput
}

func (*AppSecSecurityPolicyClone) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecSecurityPolicyClone)(nil))
}

func (i *AppSecSecurityPolicyClone) ToAppSecSecurityPolicyCloneOutput() AppSecSecurityPolicyCloneOutput {
	return i.ToAppSecSecurityPolicyCloneOutputWithContext(context.Background())
}

func (i *AppSecSecurityPolicyClone) ToAppSecSecurityPolicyCloneOutputWithContext(ctx context.Context) AppSecSecurityPolicyCloneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyCloneOutput)
}

func (i *AppSecSecurityPolicyClone) ToAppSecSecurityPolicyClonePtrOutput() AppSecSecurityPolicyClonePtrOutput {
	return i.ToAppSecSecurityPolicyClonePtrOutputWithContext(context.Background())
}

func (i *AppSecSecurityPolicyClone) ToAppSecSecurityPolicyClonePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyClonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyClonePtrOutput)
}

type AppSecSecurityPolicyClonePtrInput interface {
	pulumi.Input

	ToAppSecSecurityPolicyClonePtrOutput() AppSecSecurityPolicyClonePtrOutput
	ToAppSecSecurityPolicyClonePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyClonePtrOutput
}

type appSecSecurityPolicyClonePtrType AppSecSecurityPolicyCloneArgs

func (*appSecSecurityPolicyClonePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSecurityPolicyClone)(nil))
}

func (i *appSecSecurityPolicyClonePtrType) ToAppSecSecurityPolicyClonePtrOutput() AppSecSecurityPolicyClonePtrOutput {
	return i.ToAppSecSecurityPolicyClonePtrOutputWithContext(context.Background())
}

func (i *appSecSecurityPolicyClonePtrType) ToAppSecSecurityPolicyClonePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyClonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyClonePtrOutput)
}

// AppSecSecurityPolicyCloneArrayInput is an input type that accepts AppSecSecurityPolicyCloneArray and AppSecSecurityPolicyCloneArrayOutput values.
// You can construct a concrete instance of `AppSecSecurityPolicyCloneArrayInput` via:
//
//          AppSecSecurityPolicyCloneArray{ AppSecSecurityPolicyCloneArgs{...} }
type AppSecSecurityPolicyCloneArrayInput interface {
	pulumi.Input

	ToAppSecSecurityPolicyCloneArrayOutput() AppSecSecurityPolicyCloneArrayOutput
	ToAppSecSecurityPolicyCloneArrayOutputWithContext(context.Context) AppSecSecurityPolicyCloneArrayOutput
}

type AppSecSecurityPolicyCloneArray []AppSecSecurityPolicyCloneInput

func (AppSecSecurityPolicyCloneArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecSecurityPolicyClone)(nil))
}

func (i AppSecSecurityPolicyCloneArray) ToAppSecSecurityPolicyCloneArrayOutput() AppSecSecurityPolicyCloneArrayOutput {
	return i.ToAppSecSecurityPolicyCloneArrayOutputWithContext(context.Background())
}

func (i AppSecSecurityPolicyCloneArray) ToAppSecSecurityPolicyCloneArrayOutputWithContext(ctx context.Context) AppSecSecurityPolicyCloneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyCloneArrayOutput)
}

// AppSecSecurityPolicyCloneMapInput is an input type that accepts AppSecSecurityPolicyCloneMap and AppSecSecurityPolicyCloneMapOutput values.
// You can construct a concrete instance of `AppSecSecurityPolicyCloneMapInput` via:
//
//          AppSecSecurityPolicyCloneMap{ "key": AppSecSecurityPolicyCloneArgs{...} }
type AppSecSecurityPolicyCloneMapInput interface {
	pulumi.Input

	ToAppSecSecurityPolicyCloneMapOutput() AppSecSecurityPolicyCloneMapOutput
	ToAppSecSecurityPolicyCloneMapOutputWithContext(context.Context) AppSecSecurityPolicyCloneMapOutput
}

type AppSecSecurityPolicyCloneMap map[string]AppSecSecurityPolicyCloneInput

func (AppSecSecurityPolicyCloneMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecSecurityPolicyClone)(nil))
}

func (i AppSecSecurityPolicyCloneMap) ToAppSecSecurityPolicyCloneMapOutput() AppSecSecurityPolicyCloneMapOutput {
	return i.ToAppSecSecurityPolicyCloneMapOutputWithContext(context.Background())
}

func (i AppSecSecurityPolicyCloneMap) ToAppSecSecurityPolicyCloneMapOutputWithContext(ctx context.Context) AppSecSecurityPolicyCloneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyCloneMapOutput)
}

type AppSecSecurityPolicyCloneOutput struct {
	*pulumi.OutputState
}

func (AppSecSecurityPolicyCloneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecSecurityPolicyClone)(nil))
}

func (o AppSecSecurityPolicyCloneOutput) ToAppSecSecurityPolicyCloneOutput() AppSecSecurityPolicyCloneOutput {
	return o
}

func (o AppSecSecurityPolicyCloneOutput) ToAppSecSecurityPolicyCloneOutputWithContext(ctx context.Context) AppSecSecurityPolicyCloneOutput {
	return o
}

func (o AppSecSecurityPolicyCloneOutput) ToAppSecSecurityPolicyClonePtrOutput() AppSecSecurityPolicyClonePtrOutput {
	return o.ToAppSecSecurityPolicyClonePtrOutputWithContext(context.Background())
}

func (o AppSecSecurityPolicyCloneOutput) ToAppSecSecurityPolicyClonePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyClonePtrOutput {
	return o.ApplyT(func(v AppSecSecurityPolicyClone) *AppSecSecurityPolicyClone {
		return &v
	}).(AppSecSecurityPolicyClonePtrOutput)
}

type AppSecSecurityPolicyClonePtrOutput struct {
	*pulumi.OutputState
}

func (AppSecSecurityPolicyClonePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSecurityPolicyClone)(nil))
}

func (o AppSecSecurityPolicyClonePtrOutput) ToAppSecSecurityPolicyClonePtrOutput() AppSecSecurityPolicyClonePtrOutput {
	return o
}

func (o AppSecSecurityPolicyClonePtrOutput) ToAppSecSecurityPolicyClonePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyClonePtrOutput {
	return o
}

type AppSecSecurityPolicyCloneArrayOutput struct{ *pulumi.OutputState }

func (AppSecSecurityPolicyCloneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecSecurityPolicyClone)(nil))
}

func (o AppSecSecurityPolicyCloneArrayOutput) ToAppSecSecurityPolicyCloneArrayOutput() AppSecSecurityPolicyCloneArrayOutput {
	return o
}

func (o AppSecSecurityPolicyCloneArrayOutput) ToAppSecSecurityPolicyCloneArrayOutputWithContext(ctx context.Context) AppSecSecurityPolicyCloneArrayOutput {
	return o
}

func (o AppSecSecurityPolicyCloneArrayOutput) Index(i pulumi.IntInput) AppSecSecurityPolicyCloneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecSecurityPolicyClone {
		return vs[0].([]AppSecSecurityPolicyClone)[vs[1].(int)]
	}).(AppSecSecurityPolicyCloneOutput)
}

type AppSecSecurityPolicyCloneMapOutput struct{ *pulumi.OutputState }

func (AppSecSecurityPolicyCloneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecSecurityPolicyClone)(nil))
}

func (o AppSecSecurityPolicyCloneMapOutput) ToAppSecSecurityPolicyCloneMapOutput() AppSecSecurityPolicyCloneMapOutput {
	return o
}

func (o AppSecSecurityPolicyCloneMapOutput) ToAppSecSecurityPolicyCloneMapOutputWithContext(ctx context.Context) AppSecSecurityPolicyCloneMapOutput {
	return o
}

func (o AppSecSecurityPolicyCloneMapOutput) MapIndex(k pulumi.StringInput) AppSecSecurityPolicyCloneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecSecurityPolicyClone {
		return vs[0].(map[string]AppSecSecurityPolicyClone)[vs[1].(string)]
	}).(AppSecSecurityPolicyCloneOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecSecurityPolicyCloneOutput{})
	pulumi.RegisterOutputType(AppSecSecurityPolicyClonePtrOutput{})
	pulumi.RegisterOutputType(AppSecSecurityPolicyCloneArrayOutput{})
	pulumi.RegisterOutputType(AppSecSecurityPolicyCloneMapOutput{})
}
