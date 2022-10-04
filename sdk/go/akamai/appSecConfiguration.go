// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Contract and group
//
// Creates a new WAP (Web Application Protector) or KSD (Kona Site Defender) security configuration. KSD security configurations start out empty (i.e., unconfigured), while WAP configurations are created using preset values. The contract referenced in the request body determines the type of configuration you can create.
//
// In addition to manually creating a new configuration, you can use the `createFromConfigId` argument to clone an existing configuration.
//
// **Related API Endpoint**: [/appsec/v1/configs](https://techdocs.akamai.com/application-security/reference/post-config)
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			selectableHostnames, err := akamai.GetAppSecSelectableHostnames(ctx, &GetAppSecSelectableHostnamesArgs{
//				ConfigId: pulumi.IntRef("Documentation"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			createConfig, err := akamai.NewAppSecConfiguration(ctx, "createConfig", &akamai.AppSecConfigurationArgs{
//				Description: pulumi.String("This configuration is used as a testing environment for the documentation team."),
//				ContractId:  pulumi.String("5-2WA382"),
//				GroupId:     pulumi.Int(12198),
//				HostNames: pulumi.StringArray{
//					pulumi.String("documentation.akamai.com"),
//					pulumi.String("training.akamai.com"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("createConfigId", createConfig.ConfigId)
//			cloneConfig, err := akamai.NewAppSecConfiguration(ctx, "cloneConfig", &akamai.AppSecConfigurationArgs{
//				Description:        pulumi.String("This configuration is used as a testing environment for the documentation team."),
//				CreateFromConfigId: pulumi.Any(data.Akamai_appsec_configuration.Configuration.Config_id),
//				CreateFromVersion:  pulumi.Any(data.Akamai_appsec_configuration.Configuration.Latest_version),
//				ContractId:         pulumi.String("5-2WA382"),
//				GroupId:            pulumi.Int(12198),
//				HostNames:          interface{}(selectableHostnames.Hostnames),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("cloneConfigId", cloneConfig.ConfigId)
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `configId`. ID of the new security configuration.
type AppSecConfiguration struct {
	pulumi.CustomResourceState

	// Unique identifier of the new security configuration
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Unique identifier of the Akamai contract associated with the new configuration.
	ContractId pulumi.StringOutput `pulumi:"contractId"`
	// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
	CreateFromConfigId pulumi.IntPtrOutput `pulumi:"createFromConfigId"`
	// . Version number of the security configuration being cloned.
	CreateFromVersion pulumi.IntPtrOutput `pulumi:"createFromVersion"`
	// . Brief description of the new configuration.
	Description pulumi.StringOutput `pulumi:"description"`
	// . Unique identifier of the contract group associated with the new configuration.
	GroupId pulumi.IntOutput `pulumi:"groupId"`
	// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
	HostNames pulumi.StringArrayOutput `pulumi:"hostNames"`
	// . Name of the new configuration.
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
	// Unique identifier of the new security configuration
	ConfigId *int `pulumi:"configId"`
	// . Unique identifier of the Akamai contract associated with the new configuration.
	ContractId *string `pulumi:"contractId"`
	// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
	CreateFromConfigId *int `pulumi:"createFromConfigId"`
	// . Version number of the security configuration being cloned.
	CreateFromVersion *int `pulumi:"createFromVersion"`
	// . Brief description of the new configuration.
	Description *string `pulumi:"description"`
	// . Unique identifier of the contract group associated with the new configuration.
	GroupId *int `pulumi:"groupId"`
	// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
	HostNames []string `pulumi:"hostNames"`
	// . Name of the new configuration.
	Name *string `pulumi:"name"`
}

type AppSecConfigurationState struct {
	// Unique identifier of the new security configuration
	ConfigId pulumi.IntPtrInput
	// . Unique identifier of the Akamai contract associated with the new configuration.
	ContractId pulumi.StringPtrInput
	// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
	CreateFromConfigId pulumi.IntPtrInput
	// . Version number of the security configuration being cloned.
	CreateFromVersion pulumi.IntPtrInput
	// . Brief description of the new configuration.
	Description pulumi.StringPtrInput
	// . Unique identifier of the contract group associated with the new configuration.
	GroupId pulumi.IntPtrInput
	// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
	HostNames pulumi.StringArrayInput
	// . Name of the new configuration.
	Name pulumi.StringPtrInput
}

func (AppSecConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecConfigurationState)(nil)).Elem()
}

type appSecConfigurationArgs struct {
	// . Unique identifier of the Akamai contract associated with the new configuration.
	ContractId string `pulumi:"contractId"`
	// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
	CreateFromConfigId *int `pulumi:"createFromConfigId"`
	// . Version number of the security configuration being cloned.
	CreateFromVersion *int `pulumi:"createFromVersion"`
	// . Brief description of the new configuration.
	Description string `pulumi:"description"`
	// . Unique identifier of the contract group associated with the new configuration.
	GroupId int `pulumi:"groupId"`
	// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
	HostNames []string `pulumi:"hostNames"`
	// . Name of the new configuration.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AppSecConfiguration resource.
type AppSecConfigurationArgs struct {
	// . Unique identifier of the Akamai contract associated with the new configuration.
	ContractId pulumi.StringInput
	// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
	CreateFromConfigId pulumi.IntPtrInput
	// . Version number of the security configuration being cloned.
	CreateFromVersion pulumi.IntPtrInput
	// . Brief description of the new configuration.
	Description pulumi.StringInput
	// . Unique identifier of the contract group associated with the new configuration.
	GroupId pulumi.IntInput
	// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
	HostNames pulumi.StringArrayInput
	// . Name of the new configuration.
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
	return reflect.TypeOf((**AppSecConfiguration)(nil)).Elem()
}

func (i *AppSecConfiguration) ToAppSecConfigurationOutput() AppSecConfigurationOutput {
	return i.ToAppSecConfigurationOutputWithContext(context.Background())
}

func (i *AppSecConfiguration) ToAppSecConfigurationOutputWithContext(ctx context.Context) AppSecConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationOutput)
}

// AppSecConfigurationArrayInput is an input type that accepts AppSecConfigurationArray and AppSecConfigurationArrayOutput values.
// You can construct a concrete instance of `AppSecConfigurationArrayInput` via:
//
//	AppSecConfigurationArray{ AppSecConfigurationArgs{...} }
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
//	AppSecConfigurationMap{ "key": AppSecConfigurationArgs{...} }
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
	return reflect.TypeOf((**AppSecConfiguration)(nil)).Elem()
}

func (o AppSecConfigurationOutput) ToAppSecConfigurationOutput() AppSecConfigurationOutput {
	return o
}

func (o AppSecConfigurationOutput) ToAppSecConfigurationOutputWithContext(ctx context.Context) AppSecConfigurationOutput {
	return o
}

// Unique identifier of the new security configuration
func (o AppSecConfigurationOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecConfiguration) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// . Unique identifier of the Akamai contract associated with the new configuration.
func (o AppSecConfigurationOutput) ContractId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecConfiguration) pulumi.StringOutput { return v.ContractId }).(pulumi.StringOutput)
}

// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
func (o AppSecConfigurationOutput) CreateFromConfigId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppSecConfiguration) pulumi.IntPtrOutput { return v.CreateFromConfigId }).(pulumi.IntPtrOutput)
}

// . Version number of the security configuration being cloned.
func (o AppSecConfigurationOutput) CreateFromVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppSecConfiguration) pulumi.IntPtrOutput { return v.CreateFromVersion }).(pulumi.IntPtrOutput)
}

// . Brief description of the new configuration.
func (o AppSecConfigurationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecConfiguration) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// . Unique identifier of the contract group associated with the new configuration.
func (o AppSecConfigurationOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecConfiguration) pulumi.IntOutput { return v.GroupId }).(pulumi.IntOutput)
}

// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
func (o AppSecConfigurationOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppSecConfiguration) pulumi.StringArrayOutput { return v.HostNames }).(pulumi.StringArrayOutput)
}

// . Name of the new configuration.
func (o AppSecConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type AppSecConfigurationArrayOutput struct{ *pulumi.OutputState }

func (AppSecConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecConfiguration)(nil)).Elem()
}

func (o AppSecConfigurationArrayOutput) ToAppSecConfigurationArrayOutput() AppSecConfigurationArrayOutput {
	return o
}

func (o AppSecConfigurationArrayOutput) ToAppSecConfigurationArrayOutputWithContext(ctx context.Context) AppSecConfigurationArrayOutput {
	return o
}

func (o AppSecConfigurationArrayOutput) Index(i pulumi.IntInput) AppSecConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecConfiguration {
		return vs[0].([]*AppSecConfiguration)[vs[1].(int)]
	}).(AppSecConfigurationOutput)
}

type AppSecConfigurationMapOutput struct{ *pulumi.OutputState }

func (AppSecConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecConfiguration)(nil)).Elem()
}

func (o AppSecConfigurationMapOutput) ToAppSecConfigurationMapOutput() AppSecConfigurationMapOutput {
	return o
}

func (o AppSecConfigurationMapOutput) ToAppSecConfigurationMapOutputWithContext(ctx context.Context) AppSecConfigurationMapOutput {
	return o
}

func (o AppSecConfigurationMapOutput) MapIndex(k pulumi.StringInput) AppSecConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecConfiguration {
		return vs[0].(map[string]*AppSecConfiguration)[vs[1].(string)]
	}).(AppSecConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecConfigurationInput)(nil)).Elem(), &AppSecConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecConfigurationArrayInput)(nil)).Elem(), AppSecConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecConfigurationMapInput)(nil)).Elem(), AppSecConfigurationMap{})
	pulumi.RegisterOutputType(AppSecConfigurationOutput{})
	pulumi.RegisterOutputType(AppSecConfigurationArrayOutput{})
	pulumi.RegisterOutputType(AppSecConfigurationMapOutput{})
}
