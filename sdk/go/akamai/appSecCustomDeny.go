// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `resourceAkamaiAppsecCustomDeny` resource allows you to create a new custom deny action for a specific configuration.
type AppSecCustomDeny struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
	CustomDeny pulumi.StringOutput `pulumi:"customDeny"`
	// custom_deny_id
	CustomDenyId pulumi.StringOutput `pulumi:"customDenyId"`
}

// NewAppSecCustomDeny registers a new resource with the given unique name, arguments, and options.
func NewAppSecCustomDeny(ctx *pulumi.Context,
	name string, args *AppSecCustomDenyArgs, opts ...pulumi.ResourceOption) (*AppSecCustomDeny, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.CustomDeny == nil {
		return nil, errors.New("invalid value for required argument 'CustomDeny'")
	}
	var resource AppSecCustomDeny
	err := ctx.RegisterResource("akamai:index/appSecCustomDeny:AppSecCustomDeny", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecCustomDeny gets an existing AppSecCustomDeny resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecCustomDeny(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecCustomDenyState, opts ...pulumi.ResourceOption) (*AppSecCustomDeny, error) {
	var resource AppSecCustomDeny
	err := ctx.ReadResource("akamai:index/appSecCustomDeny:AppSecCustomDeny", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecCustomDeny resources.
type appSecCustomDenyState struct {
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
	CustomDeny *string `pulumi:"customDeny"`
	// custom_deny_id
	CustomDenyId *string `pulumi:"customDenyId"`
}

type AppSecCustomDenyState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
	CustomDeny pulumi.StringPtrInput
	// custom_deny_id
	CustomDenyId pulumi.StringPtrInput
}

func (AppSecCustomDenyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecCustomDenyState)(nil)).Elem()
}

type appSecCustomDenyArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
	CustomDeny string `pulumi:"customDeny"`
}

// The set of arguments for constructing a AppSecCustomDeny resource.
type AppSecCustomDenyArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
	CustomDeny pulumi.StringInput
}

func (AppSecCustomDenyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecCustomDenyArgs)(nil)).Elem()
}

type AppSecCustomDenyInput interface {
	pulumi.Input

	ToAppSecCustomDenyOutput() AppSecCustomDenyOutput
	ToAppSecCustomDenyOutputWithContext(ctx context.Context) AppSecCustomDenyOutput
}

func (*AppSecCustomDeny) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecCustomDeny)(nil))
}

func (i *AppSecCustomDeny) ToAppSecCustomDenyOutput() AppSecCustomDenyOutput {
	return i.ToAppSecCustomDenyOutputWithContext(context.Background())
}

func (i *AppSecCustomDeny) ToAppSecCustomDenyOutputWithContext(ctx context.Context) AppSecCustomDenyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecCustomDenyOutput)
}

func (i *AppSecCustomDeny) ToAppSecCustomDenyPtrOutput() AppSecCustomDenyPtrOutput {
	return i.ToAppSecCustomDenyPtrOutputWithContext(context.Background())
}

func (i *AppSecCustomDeny) ToAppSecCustomDenyPtrOutputWithContext(ctx context.Context) AppSecCustomDenyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecCustomDenyPtrOutput)
}

type AppSecCustomDenyPtrInput interface {
	pulumi.Input

	ToAppSecCustomDenyPtrOutput() AppSecCustomDenyPtrOutput
	ToAppSecCustomDenyPtrOutputWithContext(ctx context.Context) AppSecCustomDenyPtrOutput
}

type appSecCustomDenyPtrType AppSecCustomDenyArgs

func (*appSecCustomDenyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecCustomDeny)(nil))
}

func (i *appSecCustomDenyPtrType) ToAppSecCustomDenyPtrOutput() AppSecCustomDenyPtrOutput {
	return i.ToAppSecCustomDenyPtrOutputWithContext(context.Background())
}

func (i *appSecCustomDenyPtrType) ToAppSecCustomDenyPtrOutputWithContext(ctx context.Context) AppSecCustomDenyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecCustomDenyPtrOutput)
}

// AppSecCustomDenyArrayInput is an input type that accepts AppSecCustomDenyArray and AppSecCustomDenyArrayOutput values.
// You can construct a concrete instance of `AppSecCustomDenyArrayInput` via:
//
//          AppSecCustomDenyArray{ AppSecCustomDenyArgs{...} }
type AppSecCustomDenyArrayInput interface {
	pulumi.Input

	ToAppSecCustomDenyArrayOutput() AppSecCustomDenyArrayOutput
	ToAppSecCustomDenyArrayOutputWithContext(context.Context) AppSecCustomDenyArrayOutput
}

type AppSecCustomDenyArray []AppSecCustomDenyInput

func (AppSecCustomDenyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecCustomDeny)(nil))
}

func (i AppSecCustomDenyArray) ToAppSecCustomDenyArrayOutput() AppSecCustomDenyArrayOutput {
	return i.ToAppSecCustomDenyArrayOutputWithContext(context.Background())
}

func (i AppSecCustomDenyArray) ToAppSecCustomDenyArrayOutputWithContext(ctx context.Context) AppSecCustomDenyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecCustomDenyArrayOutput)
}

// AppSecCustomDenyMapInput is an input type that accepts AppSecCustomDenyMap and AppSecCustomDenyMapOutput values.
// You can construct a concrete instance of `AppSecCustomDenyMapInput` via:
//
//          AppSecCustomDenyMap{ "key": AppSecCustomDenyArgs{...} }
type AppSecCustomDenyMapInput interface {
	pulumi.Input

	ToAppSecCustomDenyMapOutput() AppSecCustomDenyMapOutput
	ToAppSecCustomDenyMapOutputWithContext(context.Context) AppSecCustomDenyMapOutput
}

type AppSecCustomDenyMap map[string]AppSecCustomDenyInput

func (AppSecCustomDenyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecCustomDeny)(nil))
}

func (i AppSecCustomDenyMap) ToAppSecCustomDenyMapOutput() AppSecCustomDenyMapOutput {
	return i.ToAppSecCustomDenyMapOutputWithContext(context.Background())
}

func (i AppSecCustomDenyMap) ToAppSecCustomDenyMapOutputWithContext(ctx context.Context) AppSecCustomDenyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecCustomDenyMapOutput)
}

type AppSecCustomDenyOutput struct {
	*pulumi.OutputState
}

func (AppSecCustomDenyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecCustomDeny)(nil))
}

func (o AppSecCustomDenyOutput) ToAppSecCustomDenyOutput() AppSecCustomDenyOutput {
	return o
}

func (o AppSecCustomDenyOutput) ToAppSecCustomDenyOutputWithContext(ctx context.Context) AppSecCustomDenyOutput {
	return o
}

func (o AppSecCustomDenyOutput) ToAppSecCustomDenyPtrOutput() AppSecCustomDenyPtrOutput {
	return o.ToAppSecCustomDenyPtrOutputWithContext(context.Background())
}

func (o AppSecCustomDenyOutput) ToAppSecCustomDenyPtrOutputWithContext(ctx context.Context) AppSecCustomDenyPtrOutput {
	return o.ApplyT(func(v AppSecCustomDeny) *AppSecCustomDeny {
		return &v
	}).(AppSecCustomDenyPtrOutput)
}

type AppSecCustomDenyPtrOutput struct {
	*pulumi.OutputState
}

func (AppSecCustomDenyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecCustomDeny)(nil))
}

func (o AppSecCustomDenyPtrOutput) ToAppSecCustomDenyPtrOutput() AppSecCustomDenyPtrOutput {
	return o
}

func (o AppSecCustomDenyPtrOutput) ToAppSecCustomDenyPtrOutputWithContext(ctx context.Context) AppSecCustomDenyPtrOutput {
	return o
}

type AppSecCustomDenyArrayOutput struct{ *pulumi.OutputState }

func (AppSecCustomDenyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecCustomDeny)(nil))
}

func (o AppSecCustomDenyArrayOutput) ToAppSecCustomDenyArrayOutput() AppSecCustomDenyArrayOutput {
	return o
}

func (o AppSecCustomDenyArrayOutput) ToAppSecCustomDenyArrayOutputWithContext(ctx context.Context) AppSecCustomDenyArrayOutput {
	return o
}

func (o AppSecCustomDenyArrayOutput) Index(i pulumi.IntInput) AppSecCustomDenyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecCustomDeny {
		return vs[0].([]AppSecCustomDeny)[vs[1].(int)]
	}).(AppSecCustomDenyOutput)
}

type AppSecCustomDenyMapOutput struct{ *pulumi.OutputState }

func (AppSecCustomDenyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecCustomDeny)(nil))
}

func (o AppSecCustomDenyMapOutput) ToAppSecCustomDenyMapOutput() AppSecCustomDenyMapOutput {
	return o
}

func (o AppSecCustomDenyMapOutput) ToAppSecCustomDenyMapOutputWithContext(ctx context.Context) AppSecCustomDenyMapOutput {
	return o
}

func (o AppSecCustomDenyMapOutput) MapIndex(k pulumi.StringInput) AppSecCustomDenyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecCustomDeny {
		return vs[0].(map[string]AppSecCustomDeny)[vs[1].(string)]
	}).(AppSecCustomDenyOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecCustomDenyOutput{})
	pulumi.RegisterOutputType(AppSecCustomDenyPtrOutput{})
	pulumi.RegisterOutputType(AppSecCustomDenyArrayOutput{})
	pulumi.RegisterOutputType(AppSecCustomDenyMapOutput{})
}
