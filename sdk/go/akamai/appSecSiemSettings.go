// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration
//
// Modifies SIEM (Security Information and Event Management) integration settings for a security configuration.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/siem](https://techdocs.akamai.com/application-security/reference/put-siem)
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
//			configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
//				Name: pulumi.StringRef("Documentation"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			siemDefinition, err := akamai.GetAppSecSiemDefinitions(ctx, &GetAppSecSiemDefinitionsArgs{
//				SiemDefinitionName: pulumi.StringRef("SIEM Version 01"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			securityPolicies, err := akamai.LookupAppSecSecurityPolicy(ctx, &GetAppSecSecurityPolicyArgs{
//				ConfigId: configuration.ConfigId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = akamai.NewAppSecSiemSettings(ctx, "siem", &akamai.AppSecSiemSettingsArgs{
//				ConfigId:             pulumi.Int(configuration.ConfigId),
//				EnableSiem:           pulumi.Bool(true),
//				EnableForAllPolicies: pulumi.Bool(false),
//				EnableBotmanSiem:     pulumi.Bool(true),
//				SiemId:               pulumi.String(siemDefinition.Id),
//				SecurityPolicyIds:    interface{}(securityPolicies.SecurityPolicyIdLists),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `outputText`. Tabular report showing the updated SIEM integration settings.
type AppSecSiemSettings struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the SIEM settings being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
	EnableBotmanSiem pulumi.BoolOutput `pulumi:"enableBotmanSiem"`
	// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `securityPolicyIds` argument.
	EnableForAllPolicies pulumi.BoolOutput `pulumi:"enableForAllPolicies"`
	// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
	EnableSiem pulumi.BoolOutput `pulumi:"enableSiem"`
	// JSON array of IDs for the security policies where SIEM integration is to be enabled.
	SecurityPolicyIds pulumi.StringArrayOutput `pulumi:"securityPolicyIds"`
	// . Unique identifier of the SIEM settings being modified.
	SiemId pulumi.IntOutput `pulumi:"siemId"`
}

// NewAppSecSiemSettings registers a new resource with the given unique name, arguments, and options.
func NewAppSecSiemSettings(ctx *pulumi.Context,
	name string, args *AppSecSiemSettingsArgs, opts ...pulumi.ResourceOption) (*AppSecSiemSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.EnableBotmanSiem == nil {
		return nil, errors.New("invalid value for required argument 'EnableBotmanSiem'")
	}
	if args.EnableForAllPolicies == nil {
		return nil, errors.New("invalid value for required argument 'EnableForAllPolicies'")
	}
	if args.EnableSiem == nil {
		return nil, errors.New("invalid value for required argument 'EnableSiem'")
	}
	if args.SiemId == nil {
		return nil, errors.New("invalid value for required argument 'SiemId'")
	}
	var resource AppSecSiemSettings
	err := ctx.RegisterResource("akamai:index/appSecSiemSettings:AppSecSiemSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecSiemSettings gets an existing AppSecSiemSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecSiemSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecSiemSettingsState, opts ...pulumi.ResourceOption) (*AppSecSiemSettings, error) {
	var resource AppSecSiemSettings
	err := ctx.ReadResource("akamai:index/appSecSiemSettings:AppSecSiemSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecSiemSettings resources.
type appSecSiemSettingsState struct {
	// . Unique identifier of the security configuration associated with the SIEM settings being modified.
	ConfigId *int `pulumi:"configId"`
	// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
	EnableBotmanSiem *bool `pulumi:"enableBotmanSiem"`
	// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `securityPolicyIds` argument.
	EnableForAllPolicies *bool `pulumi:"enableForAllPolicies"`
	// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
	EnableSiem *bool `pulumi:"enableSiem"`
	// JSON array of IDs for the security policies where SIEM integration is to be enabled.
	SecurityPolicyIds []string `pulumi:"securityPolicyIds"`
	// . Unique identifier of the SIEM settings being modified.
	SiemId *int `pulumi:"siemId"`
}

type AppSecSiemSettingsState struct {
	// . Unique identifier of the security configuration associated with the SIEM settings being modified.
	ConfigId pulumi.IntPtrInput
	// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
	EnableBotmanSiem pulumi.BoolPtrInput
	// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `securityPolicyIds` argument.
	EnableForAllPolicies pulumi.BoolPtrInput
	// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
	EnableSiem pulumi.BoolPtrInput
	// JSON array of IDs for the security policies where SIEM integration is to be enabled.
	SecurityPolicyIds pulumi.StringArrayInput
	// . Unique identifier of the SIEM settings being modified.
	SiemId pulumi.IntPtrInput
}

func (AppSecSiemSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSiemSettingsState)(nil)).Elem()
}

type appSecSiemSettingsArgs struct {
	// . Unique identifier of the security configuration associated with the SIEM settings being modified.
	ConfigId int `pulumi:"configId"`
	// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
	EnableBotmanSiem bool `pulumi:"enableBotmanSiem"`
	// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `securityPolicyIds` argument.
	EnableForAllPolicies bool `pulumi:"enableForAllPolicies"`
	// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
	EnableSiem bool `pulumi:"enableSiem"`
	// JSON array of IDs for the security policies where SIEM integration is to be enabled.
	SecurityPolicyIds []string `pulumi:"securityPolicyIds"`
	// . Unique identifier of the SIEM settings being modified.
	SiemId int `pulumi:"siemId"`
}

// The set of arguments for constructing a AppSecSiemSettings resource.
type AppSecSiemSettingsArgs struct {
	// . Unique identifier of the security configuration associated with the SIEM settings being modified.
	ConfigId pulumi.IntInput
	// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
	EnableBotmanSiem pulumi.BoolInput
	// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `securityPolicyIds` argument.
	EnableForAllPolicies pulumi.BoolInput
	// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
	EnableSiem pulumi.BoolInput
	// JSON array of IDs for the security policies where SIEM integration is to be enabled.
	SecurityPolicyIds pulumi.StringArrayInput
	// . Unique identifier of the SIEM settings being modified.
	SiemId pulumi.IntInput
}

func (AppSecSiemSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSiemSettingsArgs)(nil)).Elem()
}

type AppSecSiemSettingsInput interface {
	pulumi.Input

	ToAppSecSiemSettingsOutput() AppSecSiemSettingsOutput
	ToAppSecSiemSettingsOutputWithContext(ctx context.Context) AppSecSiemSettingsOutput
}

func (*AppSecSiemSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSiemSettings)(nil)).Elem()
}

func (i *AppSecSiemSettings) ToAppSecSiemSettingsOutput() AppSecSiemSettingsOutput {
	return i.ToAppSecSiemSettingsOutputWithContext(context.Background())
}

func (i *AppSecSiemSettings) ToAppSecSiemSettingsOutputWithContext(ctx context.Context) AppSecSiemSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSiemSettingsOutput)
}

// AppSecSiemSettingsArrayInput is an input type that accepts AppSecSiemSettingsArray and AppSecSiemSettingsArrayOutput values.
// You can construct a concrete instance of `AppSecSiemSettingsArrayInput` via:
//
//	AppSecSiemSettingsArray{ AppSecSiemSettingsArgs{...} }
type AppSecSiemSettingsArrayInput interface {
	pulumi.Input

	ToAppSecSiemSettingsArrayOutput() AppSecSiemSettingsArrayOutput
	ToAppSecSiemSettingsArrayOutputWithContext(context.Context) AppSecSiemSettingsArrayOutput
}

type AppSecSiemSettingsArray []AppSecSiemSettingsInput

func (AppSecSiemSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecSiemSettings)(nil)).Elem()
}

func (i AppSecSiemSettingsArray) ToAppSecSiemSettingsArrayOutput() AppSecSiemSettingsArrayOutput {
	return i.ToAppSecSiemSettingsArrayOutputWithContext(context.Background())
}

func (i AppSecSiemSettingsArray) ToAppSecSiemSettingsArrayOutputWithContext(ctx context.Context) AppSecSiemSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSiemSettingsArrayOutput)
}

// AppSecSiemSettingsMapInput is an input type that accepts AppSecSiemSettingsMap and AppSecSiemSettingsMapOutput values.
// You can construct a concrete instance of `AppSecSiemSettingsMapInput` via:
//
//	AppSecSiemSettingsMap{ "key": AppSecSiemSettingsArgs{...} }
type AppSecSiemSettingsMapInput interface {
	pulumi.Input

	ToAppSecSiemSettingsMapOutput() AppSecSiemSettingsMapOutput
	ToAppSecSiemSettingsMapOutputWithContext(context.Context) AppSecSiemSettingsMapOutput
}

type AppSecSiemSettingsMap map[string]AppSecSiemSettingsInput

func (AppSecSiemSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecSiemSettings)(nil)).Elem()
}

func (i AppSecSiemSettingsMap) ToAppSecSiemSettingsMapOutput() AppSecSiemSettingsMapOutput {
	return i.ToAppSecSiemSettingsMapOutputWithContext(context.Background())
}

func (i AppSecSiemSettingsMap) ToAppSecSiemSettingsMapOutputWithContext(ctx context.Context) AppSecSiemSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSiemSettingsMapOutput)
}

type AppSecSiemSettingsOutput struct{ *pulumi.OutputState }

func (AppSecSiemSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSiemSettings)(nil)).Elem()
}

func (o AppSecSiemSettingsOutput) ToAppSecSiemSettingsOutput() AppSecSiemSettingsOutput {
	return o
}

func (o AppSecSiemSettingsOutput) ToAppSecSiemSettingsOutputWithContext(ctx context.Context) AppSecSiemSettingsOutput {
	return o
}

// . Unique identifier of the security configuration associated with the SIEM settings being modified.
func (o AppSecSiemSettingsOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecSiemSettings) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// . Set to **true** to include Bot Manager events in your SIEM events; set to **false** to exclude Bot Manager events from your SIEM events.
func (o AppSecSiemSettingsOutput) EnableBotmanSiem() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppSecSiemSettings) pulumi.BoolOutput { return v.EnableBotmanSiem }).(pulumi.BoolOutput)
}

// . Set to **true** to enable SIEM on all security policies in the security configuration; set to **false** to only enable SIEM on the security policies specified by the `securityPolicyIds` argument.
func (o AppSecSiemSettingsOutput) EnableForAllPolicies() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppSecSiemSettings) pulumi.BoolOutput { return v.EnableForAllPolicies }).(pulumi.BoolOutput)
}

// . Set to **true** to enable SIEM; set to **false** to disable SIEM.
func (o AppSecSiemSettingsOutput) EnableSiem() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppSecSiemSettings) pulumi.BoolOutput { return v.EnableSiem }).(pulumi.BoolOutput)
}

// JSON array of IDs for the security policies where SIEM integration is to be enabled.
func (o AppSecSiemSettingsOutput) SecurityPolicyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppSecSiemSettings) pulumi.StringArrayOutput { return v.SecurityPolicyIds }).(pulumi.StringArrayOutput)
}

// . Unique identifier of the SIEM settings being modified.
func (o AppSecSiemSettingsOutput) SiemId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecSiemSettings) pulumi.IntOutput { return v.SiemId }).(pulumi.IntOutput)
}

type AppSecSiemSettingsArrayOutput struct{ *pulumi.OutputState }

func (AppSecSiemSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecSiemSettings)(nil)).Elem()
}

func (o AppSecSiemSettingsArrayOutput) ToAppSecSiemSettingsArrayOutput() AppSecSiemSettingsArrayOutput {
	return o
}

func (o AppSecSiemSettingsArrayOutput) ToAppSecSiemSettingsArrayOutputWithContext(ctx context.Context) AppSecSiemSettingsArrayOutput {
	return o
}

func (o AppSecSiemSettingsArrayOutput) Index(i pulumi.IntInput) AppSecSiemSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecSiemSettings {
		return vs[0].([]*AppSecSiemSettings)[vs[1].(int)]
	}).(AppSecSiemSettingsOutput)
}

type AppSecSiemSettingsMapOutput struct{ *pulumi.OutputState }

func (AppSecSiemSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecSiemSettings)(nil)).Elem()
}

func (o AppSecSiemSettingsMapOutput) ToAppSecSiemSettingsMapOutput() AppSecSiemSettingsMapOutput {
	return o
}

func (o AppSecSiemSettingsMapOutput) ToAppSecSiemSettingsMapOutputWithContext(ctx context.Context) AppSecSiemSettingsMapOutput {
	return o
}

func (o AppSecSiemSettingsMapOutput) MapIndex(k pulumi.StringInput) AppSecSiemSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecSiemSettings {
		return vs[0].(map[string]*AppSecSiemSettings)[vs[1].(string)]
	}).(AppSecSiemSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSiemSettingsInput)(nil)).Elem(), &AppSecSiemSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSiemSettingsArrayInput)(nil)).Elem(), AppSecSiemSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSiemSettingsMapInput)(nil)).Elem(), AppSecSiemSettingsMap{})
	pulumi.RegisterOutputType(AppSecSiemSettingsOutput{})
	pulumi.RegisterOutputType(AppSecSiemSettingsArrayOutput{})
	pulumi.RegisterOutputType(AppSecSiemSettingsMapOutput{})
}
