// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `AppSecActivations` resource allows you to activate or deactivate a given security configuration version.
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
// 		_, err = akamai.NewAppSecActivations(ctx, "activation", &akamai.AppSecActivationsArgs{
// 			ConfigId: pulumi.Int(configuration.ConfigId),
// 			Network:  pulumi.String("STAGING"),
// 			Notes:    pulumi.String("TEST Notes"),
// 			NotificationEmails: pulumi.StringArray{
// 				pulumi.String("user@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecActivations struct {
	pulumi.CustomResourceState

	// A boolean indicating whether to activate the specified configuration version. If not supplied, True is assumed.
	Activate pulumi.BoolPtrOutput `pulumi:"activate"`
	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// The network in which the security configuration should be activated. If supplied, must be either STAGING or PRODUCTION. If not supplied, STAGING will be assumed.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// A text note describing this operation. If no attributes were changed since the last time a security
	// configuration was updated using the AppSecActivations resource, an activation will not occur. To ensure an activation
	// is called, please update one of the attributes, e.g. the notes attribute.
	Notes pulumi.StringOutput `pulumi:"notes"`
	// A bracketed, comma-separated list of email addresses that will be notified when the operation is complete.
	NotificationEmails pulumi.StringArrayOutput `pulumi:"notificationEmails"`
	// The status of the operation. The following values are may be returned:
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAppSecActivations registers a new resource with the given unique name, arguments, and options.
func NewAppSecActivations(ctx *pulumi.Context,
	name string, args *AppSecActivationsArgs, opts ...pulumi.ResourceOption) (*AppSecActivations, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.Notes == nil {
		return nil, errors.New("invalid value for required argument 'Notes'")
	}
	if args.NotificationEmails == nil {
		return nil, errors.New("invalid value for required argument 'NotificationEmails'")
	}
	var resource AppSecActivations
	err := ctx.RegisterResource("akamai:index/appSecActivations:AppSecActivations", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecActivations gets an existing AppSecActivations resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecActivations(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecActivationsState, opts ...pulumi.ResourceOption) (*AppSecActivations, error) {
	var resource AppSecActivations
	err := ctx.ReadResource("akamai:index/appSecActivations:AppSecActivations", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecActivations resources.
type appSecActivationsState struct {
	// A boolean indicating whether to activate the specified configuration version. If not supplied, True is assumed.
	Activate *bool `pulumi:"activate"`
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// The network in which the security configuration should be activated. If supplied, must be either STAGING or PRODUCTION. If not supplied, STAGING will be assumed.
	Network *string `pulumi:"network"`
	// A text note describing this operation. If no attributes were changed since the last time a security
	// configuration was updated using the AppSecActivations resource, an activation will not occur. To ensure an activation
	// is called, please update one of the attributes, e.g. the notes attribute.
	Notes *string `pulumi:"notes"`
	// A bracketed, comma-separated list of email addresses that will be notified when the operation is complete.
	NotificationEmails []string `pulumi:"notificationEmails"`
	// The status of the operation. The following values are may be returned:
	Status *string `pulumi:"status"`
}

type AppSecActivationsState struct {
	// A boolean indicating whether to activate the specified configuration version. If not supplied, True is assumed.
	Activate pulumi.BoolPtrInput
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// The network in which the security configuration should be activated. If supplied, must be either STAGING or PRODUCTION. If not supplied, STAGING will be assumed.
	Network pulumi.StringPtrInput
	// A text note describing this operation. If no attributes were changed since the last time a security
	// configuration was updated using the AppSecActivations resource, an activation will not occur. To ensure an activation
	// is called, please update one of the attributes, e.g. the notes attribute.
	Notes pulumi.StringPtrInput
	// A bracketed, comma-separated list of email addresses that will be notified when the operation is complete.
	NotificationEmails pulumi.StringArrayInput
	// The status of the operation. The following values are may be returned:
	Status pulumi.StringPtrInput
}

func (AppSecActivationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecActivationsState)(nil)).Elem()
}

type appSecActivationsArgs struct {
	// A boolean indicating whether to activate the specified configuration version. If not supplied, True is assumed.
	Activate *bool `pulumi:"activate"`
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The network in which the security configuration should be activated. If supplied, must be either STAGING or PRODUCTION. If not supplied, STAGING will be assumed.
	Network *string `pulumi:"network"`
	// A text note describing this operation. If no attributes were changed since the last time a security
	// configuration was updated using the AppSecActivations resource, an activation will not occur. To ensure an activation
	// is called, please update one of the attributes, e.g. the notes attribute.
	Notes string `pulumi:"notes"`
	// A bracketed, comma-separated list of email addresses that will be notified when the operation is complete.
	NotificationEmails []string `pulumi:"notificationEmails"`
}

// The set of arguments for constructing a AppSecActivations resource.
type AppSecActivationsArgs struct {
	// A boolean indicating whether to activate the specified configuration version. If not supplied, True is assumed.
	Activate pulumi.BoolPtrInput
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// The network in which the security configuration should be activated. If supplied, must be either STAGING or PRODUCTION. If not supplied, STAGING will be assumed.
	Network pulumi.StringPtrInput
	// A text note describing this operation. If no attributes were changed since the last time a security
	// configuration was updated using the AppSecActivations resource, an activation will not occur. To ensure an activation
	// is called, please update one of the attributes, e.g. the notes attribute.
	Notes pulumi.StringInput
	// A bracketed, comma-separated list of email addresses that will be notified when the operation is complete.
	NotificationEmails pulumi.StringArrayInput
}

func (AppSecActivationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecActivationsArgs)(nil)).Elem()
}

type AppSecActivationsInput interface {
	pulumi.Input

	ToAppSecActivationsOutput() AppSecActivationsOutput
	ToAppSecActivationsOutputWithContext(ctx context.Context) AppSecActivationsOutput
}

func (*AppSecActivations) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecActivations)(nil))
}

func (i *AppSecActivations) ToAppSecActivationsOutput() AppSecActivationsOutput {
	return i.ToAppSecActivationsOutputWithContext(context.Background())
}

func (i *AppSecActivations) ToAppSecActivationsOutputWithContext(ctx context.Context) AppSecActivationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecActivationsOutput)
}

func (i *AppSecActivations) ToAppSecActivationsPtrOutput() AppSecActivationsPtrOutput {
	return i.ToAppSecActivationsPtrOutputWithContext(context.Background())
}

func (i *AppSecActivations) ToAppSecActivationsPtrOutputWithContext(ctx context.Context) AppSecActivationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecActivationsPtrOutput)
}

type AppSecActivationsPtrInput interface {
	pulumi.Input

	ToAppSecActivationsPtrOutput() AppSecActivationsPtrOutput
	ToAppSecActivationsPtrOutputWithContext(ctx context.Context) AppSecActivationsPtrOutput
}

type appSecActivationsPtrType AppSecActivationsArgs

func (*appSecActivationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecActivations)(nil))
}

func (i *appSecActivationsPtrType) ToAppSecActivationsPtrOutput() AppSecActivationsPtrOutput {
	return i.ToAppSecActivationsPtrOutputWithContext(context.Background())
}

func (i *appSecActivationsPtrType) ToAppSecActivationsPtrOutputWithContext(ctx context.Context) AppSecActivationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecActivationsPtrOutput)
}

// AppSecActivationsArrayInput is an input type that accepts AppSecActivationsArray and AppSecActivationsArrayOutput values.
// You can construct a concrete instance of `AppSecActivationsArrayInput` via:
//
//          AppSecActivationsArray{ AppSecActivationsArgs{...} }
type AppSecActivationsArrayInput interface {
	pulumi.Input

	ToAppSecActivationsArrayOutput() AppSecActivationsArrayOutput
	ToAppSecActivationsArrayOutputWithContext(context.Context) AppSecActivationsArrayOutput
}

type AppSecActivationsArray []AppSecActivationsInput

func (AppSecActivationsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecActivations)(nil))
}

func (i AppSecActivationsArray) ToAppSecActivationsArrayOutput() AppSecActivationsArrayOutput {
	return i.ToAppSecActivationsArrayOutputWithContext(context.Background())
}

func (i AppSecActivationsArray) ToAppSecActivationsArrayOutputWithContext(ctx context.Context) AppSecActivationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecActivationsArrayOutput)
}

// AppSecActivationsMapInput is an input type that accepts AppSecActivationsMap and AppSecActivationsMapOutput values.
// You can construct a concrete instance of `AppSecActivationsMapInput` via:
//
//          AppSecActivationsMap{ "key": AppSecActivationsArgs{...} }
type AppSecActivationsMapInput interface {
	pulumi.Input

	ToAppSecActivationsMapOutput() AppSecActivationsMapOutput
	ToAppSecActivationsMapOutputWithContext(context.Context) AppSecActivationsMapOutput
}

type AppSecActivationsMap map[string]AppSecActivationsInput

func (AppSecActivationsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecActivations)(nil))
}

func (i AppSecActivationsMap) ToAppSecActivationsMapOutput() AppSecActivationsMapOutput {
	return i.ToAppSecActivationsMapOutputWithContext(context.Background())
}

func (i AppSecActivationsMap) ToAppSecActivationsMapOutputWithContext(ctx context.Context) AppSecActivationsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecActivationsMapOutput)
}

type AppSecActivationsOutput struct {
	*pulumi.OutputState
}

func (AppSecActivationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecActivations)(nil))
}

func (o AppSecActivationsOutput) ToAppSecActivationsOutput() AppSecActivationsOutput {
	return o
}

func (o AppSecActivationsOutput) ToAppSecActivationsOutputWithContext(ctx context.Context) AppSecActivationsOutput {
	return o
}

func (o AppSecActivationsOutput) ToAppSecActivationsPtrOutput() AppSecActivationsPtrOutput {
	return o.ToAppSecActivationsPtrOutputWithContext(context.Background())
}

func (o AppSecActivationsOutput) ToAppSecActivationsPtrOutputWithContext(ctx context.Context) AppSecActivationsPtrOutput {
	return o.ApplyT(func(v AppSecActivations) *AppSecActivations {
		return &v
	}).(AppSecActivationsPtrOutput)
}

type AppSecActivationsPtrOutput struct {
	*pulumi.OutputState
}

func (AppSecActivationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecActivations)(nil))
}

func (o AppSecActivationsPtrOutput) ToAppSecActivationsPtrOutput() AppSecActivationsPtrOutput {
	return o
}

func (o AppSecActivationsPtrOutput) ToAppSecActivationsPtrOutputWithContext(ctx context.Context) AppSecActivationsPtrOutput {
	return o
}

type AppSecActivationsArrayOutput struct{ *pulumi.OutputState }

func (AppSecActivationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecActivations)(nil))
}

func (o AppSecActivationsArrayOutput) ToAppSecActivationsArrayOutput() AppSecActivationsArrayOutput {
	return o
}

func (o AppSecActivationsArrayOutput) ToAppSecActivationsArrayOutputWithContext(ctx context.Context) AppSecActivationsArrayOutput {
	return o
}

func (o AppSecActivationsArrayOutput) Index(i pulumi.IntInput) AppSecActivationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecActivations {
		return vs[0].([]AppSecActivations)[vs[1].(int)]
	}).(AppSecActivationsOutput)
}

type AppSecActivationsMapOutput struct{ *pulumi.OutputState }

func (AppSecActivationsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecActivations)(nil))
}

func (o AppSecActivationsMapOutput) ToAppSecActivationsMapOutput() AppSecActivationsMapOutput {
	return o
}

func (o AppSecActivationsMapOutput) ToAppSecActivationsMapOutputWithContext(ctx context.Context) AppSecActivationsMapOutput {
	return o
}

func (o AppSecActivationsMapOutput) MapIndex(k pulumi.StringInput) AppSecActivationsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecActivations {
		return vs[0].(map[string]AppSecActivations)[vs[1].(string)]
	}).(AppSecActivationsOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecActivationsOutput{})
	pulumi.RegisterOutputType(AppSecActivationsPtrOutput{})
	pulumi.RegisterOutputType(AppSecActivationsArrayOutput{})
	pulumi.RegisterOutputType(AppSecActivationsMapOutput{})
}
