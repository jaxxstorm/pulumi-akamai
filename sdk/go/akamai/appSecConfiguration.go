// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `resourceAkamaiAppsecConfiguration` resource allows you to create a new WAP or KSD security configuration. KSD security configurations start out empty, and WAP configurations are created with preset values. The contract you pass in the request body determines which product you use. You can edit the default settings included in the WAP configuration, but you’ll need to run additional operations in this API to select specific protections for KSD. Your KSD configuration needs match targets and protection settings before it can be activated.
type AppSecConfiguration struct {
	pulumi.CustomResourceState

	// (Required) The ID of the security configuration.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// The contract ID of the configuration.
	ContractId pulumi.StringOutput `pulumi:"contractId"`
	// The config ID of the security configuration to clone from.
	CreateFromConfigId pulumi.IntPtrOutput `pulumi:"createFromConfigId"`
	// The version number of the security configuration to clone from.
	CreateFromVersion pulumi.IntPtrOutput `pulumi:"createFromVersion"`
	// A description of the configuration.
	Description pulumi.StringOutput `pulumi:"description"`
	// The group ID of the configuration.
	GroupId pulumi.IntOutput `pulumi:"groupId"`
	// The list of hostnames protected by this security configuration.
	HostNames pulumi.StringArrayOutput `pulumi:"hostNames"`
	// The name to be assigned to the configuration.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAppSecConfiguration registers a new resource with the given unique name, arguments, and options.
func NewAppSecConfiguration(ctx *pulumi.Context,
	name string, args *AppSecConfigurationArgs, opts ...pulumi.ResourceOption) (*AppSecConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContractId == nil {
		return nil, errors.New("invalid value for required argument 'ContractId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.HostNames == nil {
		return nil, errors.New("invalid value for required argument 'HostNames'")
	}
	var resource AppSecConfiguration
	err := ctx.RegisterResource("akamai:index/appSecConfiguration:AppSecConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecConfiguration gets an existing AppSecConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecConfigurationState, opts ...pulumi.ResourceOption) (*AppSecConfiguration, error) {
	var resource AppSecConfiguration
	err := ctx.ReadResource("akamai:index/appSecConfiguration:AppSecConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecConfiguration resources.
type appSecConfigurationState struct {
	// (Required) The ID of the security configuration.
	ConfigId *int `pulumi:"configId"`
	// The contract ID of the configuration.
	ContractId *string `pulumi:"contractId"`
	// The config ID of the security configuration to clone from.
	CreateFromConfigId *int `pulumi:"createFromConfigId"`
	// The version number of the security configuration to clone from.
	CreateFromVersion *int `pulumi:"createFromVersion"`
	// A description of the configuration.
	Description *string `pulumi:"description"`
	// The group ID of the configuration.
	GroupId *int `pulumi:"groupId"`
	// The list of hostnames protected by this security configuration.
	HostNames []string `pulumi:"hostNames"`
	// The name to be assigned to the configuration.
	Name *string `pulumi:"name"`
}

type AppSecConfigurationState struct {
	// (Required) The ID of the security configuration.
	ConfigId pulumi.IntPtrInput
	// The contract ID of the configuration.
	ContractId pulumi.StringPtrInput
	// The config ID of the security configuration to clone from.
	CreateFromConfigId pulumi.IntPtrInput
	// The version number of the security configuration to clone from.
	CreateFromVersion pulumi.IntPtrInput
	// A description of the configuration.
	Description pulumi.StringPtrInput
	// The group ID of the configuration.
	GroupId pulumi.IntPtrInput
	// The list of hostnames protected by this security configuration.
	HostNames pulumi.StringArrayInput
	// The name to be assigned to the configuration.
	Name pulumi.StringPtrInput
}

func (AppSecConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecConfigurationState)(nil)).Elem()
}

type appSecConfigurationArgs struct {
	// The contract ID of the configuration.
	ContractId string `pulumi:"contractId"`
	// The config ID of the security configuration to clone from.
	CreateFromConfigId *int `pulumi:"createFromConfigId"`
	// The version number of the security configuration to clone from.
	CreateFromVersion *int `pulumi:"createFromVersion"`
	// A description of the configuration.
	Description string `pulumi:"description"`
	// The group ID of the configuration.
	GroupId int `pulumi:"groupId"`
	// The list of hostnames protected by this security configuration.
	HostNames []string `pulumi:"hostNames"`
	// The name to be assigned to the configuration.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AppSecConfiguration resource.
type AppSecConfigurationArgs struct {
	// The contract ID of the configuration.
	ContractId pulumi.StringInput
	// The config ID of the security configuration to clone from.
	CreateFromConfigId pulumi.IntPtrInput
	// The version number of the security configuration to clone from.
	CreateFromVersion pulumi.IntPtrInput
	// A description of the configuration.
	Description pulumi.StringInput
	// The group ID of the configuration.
	GroupId pulumi.IntInput
	// The list of hostnames protected by this security configuration.
	HostNames pulumi.StringArrayInput
	// The name to be assigned to the configuration.
	Name pulumi.StringPtrInput
}

func (AppSecConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecConfigurationArgs)(nil)).Elem()
}

type AppSecConfigurationInput interface {
	pulumi.Input

	ToAppSecConfigurationOutput() AppSecConfigurationOutput
	ToAppSecConfigurationOutputWithContext(ctx context.Context) AppSecConfigurationOutput
}

func (*AppSecConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecConfiguration)(nil))
}

func (i *AppSecConfiguration) ToAppSecConfigurationOutput() AppSecConfigurationOutput {
	return i.ToAppSecConfigurationOutputWithContext(context.Background())
}

func (i *AppSecConfiguration) ToAppSecConfigurationOutputWithContext(ctx context.Context) AppSecConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationOutput)
}

func (i *AppSecConfiguration) ToAppSecConfigurationPtrOutput() AppSecConfigurationPtrOutput {
	return i.ToAppSecConfigurationPtrOutputWithContext(context.Background())
}

func (i *AppSecConfiguration) ToAppSecConfigurationPtrOutputWithContext(ctx context.Context) AppSecConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationPtrOutput)
}

type AppSecConfigurationPtrInput interface {
	pulumi.Input

	ToAppSecConfigurationPtrOutput() AppSecConfigurationPtrOutput
	ToAppSecConfigurationPtrOutputWithContext(ctx context.Context) AppSecConfigurationPtrOutput
}

type appSecConfigurationPtrType AppSecConfigurationArgs

func (*appSecConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecConfiguration)(nil))
}

func (i *appSecConfigurationPtrType) ToAppSecConfigurationPtrOutput() AppSecConfigurationPtrOutput {
	return i.ToAppSecConfigurationPtrOutputWithContext(context.Background())
}

func (i *appSecConfigurationPtrType) ToAppSecConfigurationPtrOutputWithContext(ctx context.Context) AppSecConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationPtrOutput)
}

// AppSecConfigurationArrayInput is an input type that accepts AppSecConfigurationArray and AppSecConfigurationArrayOutput values.
// You can construct a concrete instance of `AppSecConfigurationArrayInput` via:
//
//          AppSecConfigurationArray{ AppSecConfigurationArgs{...} }
type AppSecConfigurationArrayInput interface {
	pulumi.Input

	ToAppSecConfigurationArrayOutput() AppSecConfigurationArrayOutput
	ToAppSecConfigurationArrayOutputWithContext(context.Context) AppSecConfigurationArrayOutput
}

type AppSecConfigurationArray []AppSecConfigurationInput

func (AppSecConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecConfiguration)(nil)).Elem()
}

func (i AppSecConfigurationArray) ToAppSecConfigurationArrayOutput() AppSecConfigurationArrayOutput {
	return i.ToAppSecConfigurationArrayOutputWithContext(context.Background())
}

func (i AppSecConfigurationArray) ToAppSecConfigurationArrayOutputWithContext(ctx context.Context) AppSecConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationArrayOutput)
}

// AppSecConfigurationMapInput is an input type that accepts AppSecConfigurationMap and AppSecConfigurationMapOutput values.
// You can construct a concrete instance of `AppSecConfigurationMapInput` via:
//
//          AppSecConfigurationMap{ "key": AppSecConfigurationArgs{...} }
type AppSecConfigurationMapInput interface {
	pulumi.Input

	ToAppSecConfigurationMapOutput() AppSecConfigurationMapOutput
	ToAppSecConfigurationMapOutputWithContext(context.Context) AppSecConfigurationMapOutput
}

type AppSecConfigurationMap map[string]AppSecConfigurationInput

func (AppSecConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecConfiguration)(nil)).Elem()
}

func (i AppSecConfigurationMap) ToAppSecConfigurationMapOutput() AppSecConfigurationMapOutput {
	return i.ToAppSecConfigurationMapOutputWithContext(context.Background())
}

func (i AppSecConfigurationMap) ToAppSecConfigurationMapOutputWithContext(ctx context.Context) AppSecConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationMapOutput)
}

type AppSecConfigurationOutput struct{ *pulumi.OutputState }

func (AppSecConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecConfiguration)(nil))
}

func (o AppSecConfigurationOutput) ToAppSecConfigurationOutput() AppSecConfigurationOutput {
	return o
}

func (o AppSecConfigurationOutput) ToAppSecConfigurationOutputWithContext(ctx context.Context) AppSecConfigurationOutput {
	return o
}

func (o AppSecConfigurationOutput) ToAppSecConfigurationPtrOutput() AppSecConfigurationPtrOutput {
	return o.ToAppSecConfigurationPtrOutputWithContext(context.Background())
}

func (o AppSecConfigurationOutput) ToAppSecConfigurationPtrOutputWithContext(ctx context.Context) AppSecConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSecConfiguration) *AppSecConfiguration {
		return &v
	}).(AppSecConfigurationPtrOutput)
}

type AppSecConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AppSecConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecConfiguration)(nil))
}

func (o AppSecConfigurationPtrOutput) ToAppSecConfigurationPtrOutput() AppSecConfigurationPtrOutput {
	return o
}

func (o AppSecConfigurationPtrOutput) ToAppSecConfigurationPtrOutputWithContext(ctx context.Context) AppSecConfigurationPtrOutput {
	return o
}

func (o AppSecConfigurationPtrOutput) Elem() AppSecConfigurationOutput {
	return o.ApplyT(func(v *AppSecConfiguration) AppSecConfiguration {
		if v != nil {
			return *v
		}
		var ret AppSecConfiguration
		return ret
	}).(AppSecConfigurationOutput)
}

type AppSecConfigurationArrayOutput struct{ *pulumi.OutputState }

func (AppSecConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecConfiguration)(nil))
}

func (o AppSecConfigurationArrayOutput) ToAppSecConfigurationArrayOutput() AppSecConfigurationArrayOutput {
	return o
}

func (o AppSecConfigurationArrayOutput) ToAppSecConfigurationArrayOutputWithContext(ctx context.Context) AppSecConfigurationArrayOutput {
	return o
}

func (o AppSecConfigurationArrayOutput) Index(i pulumi.IntInput) AppSecConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecConfiguration {
		return vs[0].([]AppSecConfiguration)[vs[1].(int)]
	}).(AppSecConfigurationOutput)
}

type AppSecConfigurationMapOutput struct{ *pulumi.OutputState }

func (AppSecConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecConfiguration)(nil))
}

func (o AppSecConfigurationMapOutput) ToAppSecConfigurationMapOutput() AppSecConfigurationMapOutput {
	return o
}

func (o AppSecConfigurationMapOutput) ToAppSecConfigurationMapOutputWithContext(ctx context.Context) AppSecConfigurationMapOutput {
	return o
}

func (o AppSecConfigurationMapOutput) MapIndex(k pulumi.StringInput) AppSecConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecConfiguration {
		return vs[0].(map[string]AppSecConfiguration)[vs[1].(string)]
	}).(AppSecConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecConfigurationInput)(nil)).Elem(), &AppSecConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecConfigurationPtrInput)(nil)).Elem(), &AppSecConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecConfigurationArrayInput)(nil)).Elem(), AppSecConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecConfigurationMapInput)(nil)).Elem(), AppSecConfigurationMap{})
	pulumi.RegisterOutputType(AppSecConfigurationOutput{})
	pulumi.RegisterOutputType(AppSecConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AppSecConfigurationArrayOutput{})
	pulumi.RegisterOutputType(AppSecConfigurationMapOutput{})
}
