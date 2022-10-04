// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppSecActivations struct {
	pulumi.CustomResourceState

	// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting activate has been deprecated; "terraform apply" will always perform activation. (Use "terraform destroy" for deactivation.)
	Activate pulumi.BoolPtrOutput `pulumi:"activate"`
	// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
	// * **PRODUCTION**
	// * **STAGING**
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting notes has been deprecated. Use "note" instead.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
	NotificationEmails pulumi.StringArrayOutput `pulumi:"notificationEmails"`
	// The results of the activation
	Status pulumi.StringOutput `pulumi:"status"`
	// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
	Version pulumi.IntOutput `pulumi:"version"`
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
	if args.NotificationEmails == nil {
		return nil, errors.New("invalid value for required argument 'NotificationEmails'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
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
	// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting activate has been deprecated; "terraform apply" will always perform activation. (Use "terraform destroy" for deactivation.)
	Activate *bool `pulumi:"activate"`
	// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
	ConfigId *int `pulumi:"configId"`
	// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
	// * **PRODUCTION**
	// * **STAGING**
	Network *string `pulumi:"network"`
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
	Note *string `pulumi:"note"`
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting notes has been deprecated. Use "note" instead.
	Notes *string `pulumi:"notes"`
	// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
	NotificationEmails []string `pulumi:"notificationEmails"`
	// The results of the activation
	Status *string `pulumi:"status"`
	// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
	Version *int `pulumi:"version"`
}

type AppSecActivationsState struct {
	// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting activate has been deprecated; "terraform apply" will always perform activation. (Use "terraform destroy" for deactivation.)
	Activate pulumi.BoolPtrInput
	// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
	ConfigId pulumi.IntPtrInput
	// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
	// * **PRODUCTION**
	// * **STAGING**
	Network pulumi.StringPtrInput
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
	Note pulumi.StringPtrInput
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting notes has been deprecated. Use "note" instead.
	Notes pulumi.StringPtrInput
	// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
	NotificationEmails pulumi.StringArrayInput
	// The results of the activation
	Status pulumi.StringPtrInput
	// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
	Version pulumi.IntPtrInput
}

func (AppSecActivationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecActivationsState)(nil)).Elem()
}

type appSecActivationsArgs struct {
	// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting activate has been deprecated; "terraform apply" will always perform activation. (Use "terraform destroy" for deactivation.)
	Activate *bool `pulumi:"activate"`
	// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
	ConfigId int `pulumi:"configId"`
	// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
	// * **PRODUCTION**
	// * **STAGING**
	Network *string `pulumi:"network"`
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
	Note *string `pulumi:"note"`
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting notes has been deprecated. Use "note" instead.
	Notes *string `pulumi:"notes"`
	// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
	NotificationEmails []string `pulumi:"notificationEmails"`
	// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
	Version int `pulumi:"version"`
}

// The set of arguments for constructing a AppSecActivations resource.
type AppSecActivationsArgs struct {
	// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting activate has been deprecated; "terraform apply" will always perform activation. (Use "terraform destroy" for deactivation.)
	Activate pulumi.BoolPtrInput
	// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
	ConfigId pulumi.IntInput
	// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
	// * **PRODUCTION**
	// * **STAGING**
	Network pulumi.StringPtrInput
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
	Note pulumi.StringPtrInput
	// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
	//
	// Deprecated: The setting notes has been deprecated. Use "note" instead.
	Notes pulumi.StringPtrInput
	// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
	NotificationEmails pulumi.StringArrayInput
	// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
	Version pulumi.IntInput
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
	return reflect.TypeOf((**AppSecActivations)(nil)).Elem()
}

func (i *AppSecActivations) ToAppSecActivationsOutput() AppSecActivationsOutput {
	return i.ToAppSecActivationsOutputWithContext(context.Background())
}

func (i *AppSecActivations) ToAppSecActivationsOutputWithContext(ctx context.Context) AppSecActivationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecActivationsOutput)
}

// AppSecActivationsArrayInput is an input type that accepts AppSecActivationsArray and AppSecActivationsArrayOutput values.
// You can construct a concrete instance of `AppSecActivationsArrayInput` via:
//
//	AppSecActivationsArray{ AppSecActivationsArgs{...} }
type AppSecActivationsArrayInput interface {
	pulumi.Input

	ToAppSecActivationsArrayOutput() AppSecActivationsArrayOutput
	ToAppSecActivationsArrayOutputWithContext(context.Context) AppSecActivationsArrayOutput
}

type AppSecActivationsArray []AppSecActivationsInput

func (AppSecActivationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecActivations)(nil)).Elem()
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
//	AppSecActivationsMap{ "key": AppSecActivationsArgs{...} }
type AppSecActivationsMapInput interface {
	pulumi.Input

	ToAppSecActivationsMapOutput() AppSecActivationsMapOutput
	ToAppSecActivationsMapOutputWithContext(context.Context) AppSecActivationsMapOutput
}

type AppSecActivationsMap map[string]AppSecActivationsInput

func (AppSecActivationsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecActivations)(nil)).Elem()
}

func (i AppSecActivationsMap) ToAppSecActivationsMapOutput() AppSecActivationsMapOutput {
	return i.ToAppSecActivationsMapOutputWithContext(context.Background())
}

func (i AppSecActivationsMap) ToAppSecActivationsMapOutputWithContext(ctx context.Context) AppSecActivationsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecActivationsMapOutput)
}

type AppSecActivationsOutput struct{ *pulumi.OutputState }

func (AppSecActivationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecActivations)(nil)).Elem()
}

func (o AppSecActivationsOutput) ToAppSecActivationsOutput() AppSecActivationsOutput {
	return o
}

func (o AppSecActivationsOutput) ToAppSecActivationsOutputWithContext(ctx context.Context) AppSecActivationsOutput {
	return o
}

// . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
//
// Deprecated: The setting activate has been deprecated; "terraform apply" will always perform activation. (Use "terraform destroy" for deactivation.)
func (o AppSecActivationsOutput) Activate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppSecActivations) pulumi.BoolPtrOutput { return v.Activate }).(pulumi.BoolPtrOutput)
}

// . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
func (o AppSecActivationsOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecActivations) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
// * **PRODUCTION**
// * **STAGING**
func (o AppSecActivationsOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSecActivations) pulumi.StringPtrOutput { return v.Network }).(pulumi.StringPtrOutput)
}

// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
func (o AppSecActivationsOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSecActivations) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

// . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
//
// Deprecated: The setting notes has been deprecated. Use "note" instead.
func (o AppSecActivationsOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSecActivations) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
func (o AppSecActivationsOutput) NotificationEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppSecActivations) pulumi.StringArrayOutput { return v.NotificationEmails }).(pulumi.StringArrayOutput)
}

// The results of the activation
func (o AppSecActivationsOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecActivations) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
func (o AppSecActivationsOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecActivations) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type AppSecActivationsArrayOutput struct{ *pulumi.OutputState }

func (AppSecActivationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecActivations)(nil)).Elem()
}

func (o AppSecActivationsArrayOutput) ToAppSecActivationsArrayOutput() AppSecActivationsArrayOutput {
	return o
}

func (o AppSecActivationsArrayOutput) ToAppSecActivationsArrayOutputWithContext(ctx context.Context) AppSecActivationsArrayOutput {
	return o
}

func (o AppSecActivationsArrayOutput) Index(i pulumi.IntInput) AppSecActivationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecActivations {
		return vs[0].([]*AppSecActivations)[vs[1].(int)]
	}).(AppSecActivationsOutput)
}

type AppSecActivationsMapOutput struct{ *pulumi.OutputState }

func (AppSecActivationsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecActivations)(nil)).Elem()
}

func (o AppSecActivationsMapOutput) ToAppSecActivationsMapOutput() AppSecActivationsMapOutput {
	return o
}

func (o AppSecActivationsMapOutput) ToAppSecActivationsMapOutputWithContext(ctx context.Context) AppSecActivationsMapOutput {
	return o
}

func (o AppSecActivationsMapOutput) MapIndex(k pulumi.StringInput) AppSecActivationsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecActivations {
		return vs[0].(map[string]*AppSecActivations)[vs[1].(string)]
	}).(AppSecActivationsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecActivationsInput)(nil)).Elem(), &AppSecActivations{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecActivationsArrayInput)(nil)).Elem(), AppSecActivationsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecActivationsMapInput)(nil)).Elem(), AppSecActivationsMap{})
	pulumi.RegisterOutputType(AppSecActivationsOutput{})
	pulumi.RegisterOutputType(AppSecActivationsArrayOutput{})
	pulumi.RegisterOutputType(AppSecActivationsMapOutput{})
}
