// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy
//
// Enables or disables threat intelligence for a security policy. This resource is only available to organizations running the Adaptive Security Engine (ASE) beta Please contact your Akamai representative for more information.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/threat-intel](https://techdocs.akamai.com/application-security/reference/put-rules-threat-intel)
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: pulumi.StringRef("Documentation"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecThreatIntel(ctx, "threatIntel", &akamai.AppSecThreatIntelArgs{
// 			ConfigId:         pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId: pulumi.String("gms1_134637"),
// 			ThreatIntel:      pulumi.String("on"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecThreatIntel struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
	// . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
	ThreatIntel pulumi.StringOutput `pulumi:"threatIntel"`
}

// NewAppSecThreatIntel registers a new resource with the given unique name, arguments, and options.
func NewAppSecThreatIntel(ctx *pulumi.Context,
	name string, args *AppSecThreatIntelArgs, opts ...pulumi.ResourceOption) (*AppSecThreatIntel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	if args.ThreatIntel == nil {
		return nil, errors.New("invalid value for required argument 'ThreatIntel'")
	}
	var resource AppSecThreatIntel
	err := ctx.RegisterResource("akamai:index/appSecThreatIntel:AppSecThreatIntel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecThreatIntel gets an existing AppSecThreatIntel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecThreatIntel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecThreatIntelState, opts ...pulumi.ResourceOption) (*AppSecThreatIntel, error) {
	var resource AppSecThreatIntel
	err := ctx.ReadResource("akamai:index/appSecThreatIntel:AppSecThreatIntel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecThreatIntel resources.
type appSecThreatIntelState struct {
	// . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
	ConfigId *int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
	// . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
	ThreatIntel *string `pulumi:"threatIntel"`
}

type AppSecThreatIntelState struct {
	// . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
	ConfigId pulumi.IntPtrInput
	// . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
	SecurityPolicyId pulumi.StringPtrInput
	// . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
	ThreatIntel pulumi.StringPtrInput
}

func (AppSecThreatIntelState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecThreatIntelState)(nil)).Elem()
}

type appSecThreatIntelArgs struct {
	// . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	// . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
	ThreatIntel string `pulumi:"threatIntel"`
}

// The set of arguments for constructing a AppSecThreatIntel resource.
type AppSecThreatIntelArgs struct {
	// . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
	ConfigId pulumi.IntInput
	// . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
	SecurityPolicyId pulumi.StringInput
	// . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
	ThreatIntel pulumi.StringInput
}

func (AppSecThreatIntelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecThreatIntelArgs)(nil)).Elem()
}

type AppSecThreatIntelInput interface {
	pulumi.Input

	ToAppSecThreatIntelOutput() AppSecThreatIntelOutput
	ToAppSecThreatIntelOutputWithContext(ctx context.Context) AppSecThreatIntelOutput
}

func (*AppSecThreatIntel) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecThreatIntel)(nil)).Elem()
}

func (i *AppSecThreatIntel) ToAppSecThreatIntelOutput() AppSecThreatIntelOutput {
	return i.ToAppSecThreatIntelOutputWithContext(context.Background())
}

func (i *AppSecThreatIntel) ToAppSecThreatIntelOutputWithContext(ctx context.Context) AppSecThreatIntelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecThreatIntelOutput)
}

// AppSecThreatIntelArrayInput is an input type that accepts AppSecThreatIntelArray and AppSecThreatIntelArrayOutput values.
// You can construct a concrete instance of `AppSecThreatIntelArrayInput` via:
//
//          AppSecThreatIntelArray{ AppSecThreatIntelArgs{...} }
type AppSecThreatIntelArrayInput interface {
	pulumi.Input

	ToAppSecThreatIntelArrayOutput() AppSecThreatIntelArrayOutput
	ToAppSecThreatIntelArrayOutputWithContext(context.Context) AppSecThreatIntelArrayOutput
}

type AppSecThreatIntelArray []AppSecThreatIntelInput

func (AppSecThreatIntelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecThreatIntel)(nil)).Elem()
}

func (i AppSecThreatIntelArray) ToAppSecThreatIntelArrayOutput() AppSecThreatIntelArrayOutput {
	return i.ToAppSecThreatIntelArrayOutputWithContext(context.Background())
}

func (i AppSecThreatIntelArray) ToAppSecThreatIntelArrayOutputWithContext(ctx context.Context) AppSecThreatIntelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecThreatIntelArrayOutput)
}

// AppSecThreatIntelMapInput is an input type that accepts AppSecThreatIntelMap and AppSecThreatIntelMapOutput values.
// You can construct a concrete instance of `AppSecThreatIntelMapInput` via:
//
//          AppSecThreatIntelMap{ "key": AppSecThreatIntelArgs{...} }
type AppSecThreatIntelMapInput interface {
	pulumi.Input

	ToAppSecThreatIntelMapOutput() AppSecThreatIntelMapOutput
	ToAppSecThreatIntelMapOutputWithContext(context.Context) AppSecThreatIntelMapOutput
}

type AppSecThreatIntelMap map[string]AppSecThreatIntelInput

func (AppSecThreatIntelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecThreatIntel)(nil)).Elem()
}

func (i AppSecThreatIntelMap) ToAppSecThreatIntelMapOutput() AppSecThreatIntelMapOutput {
	return i.ToAppSecThreatIntelMapOutputWithContext(context.Background())
}

func (i AppSecThreatIntelMap) ToAppSecThreatIntelMapOutputWithContext(ctx context.Context) AppSecThreatIntelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecThreatIntelMapOutput)
}

type AppSecThreatIntelOutput struct{ *pulumi.OutputState }

func (AppSecThreatIntelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecThreatIntel)(nil)).Elem()
}

func (o AppSecThreatIntelOutput) ToAppSecThreatIntelOutput() AppSecThreatIntelOutput {
	return o
}

func (o AppSecThreatIntelOutput) ToAppSecThreatIntelOutputWithContext(ctx context.Context) AppSecThreatIntelOutput {
	return o
}

// . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
func (o AppSecThreatIntelOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecThreatIntel) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
func (o AppSecThreatIntelOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecThreatIntel) pulumi.StringOutput { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

// . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
func (o AppSecThreatIntelOutput) ThreatIntel() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecThreatIntel) pulumi.StringOutput { return v.ThreatIntel }).(pulumi.StringOutput)
}

type AppSecThreatIntelArrayOutput struct{ *pulumi.OutputState }

func (AppSecThreatIntelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecThreatIntel)(nil)).Elem()
}

func (o AppSecThreatIntelArrayOutput) ToAppSecThreatIntelArrayOutput() AppSecThreatIntelArrayOutput {
	return o
}

func (o AppSecThreatIntelArrayOutput) ToAppSecThreatIntelArrayOutputWithContext(ctx context.Context) AppSecThreatIntelArrayOutput {
	return o
}

func (o AppSecThreatIntelArrayOutput) Index(i pulumi.IntInput) AppSecThreatIntelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecThreatIntel {
		return vs[0].([]*AppSecThreatIntel)[vs[1].(int)]
	}).(AppSecThreatIntelOutput)
}

type AppSecThreatIntelMapOutput struct{ *pulumi.OutputState }

func (AppSecThreatIntelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecThreatIntel)(nil)).Elem()
}

func (o AppSecThreatIntelMapOutput) ToAppSecThreatIntelMapOutput() AppSecThreatIntelMapOutput {
	return o
}

func (o AppSecThreatIntelMapOutput) ToAppSecThreatIntelMapOutputWithContext(ctx context.Context) AppSecThreatIntelMapOutput {
	return o
}

func (o AppSecThreatIntelMapOutput) MapIndex(k pulumi.StringInput) AppSecThreatIntelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecThreatIntel {
		return vs[0].(map[string]*AppSecThreatIntel)[vs[1].(string)]
	}).(AppSecThreatIntelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecThreatIntelInput)(nil)).Elem(), &AppSecThreatIntel{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecThreatIntelArrayInput)(nil)).Elem(), AppSecThreatIntelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecThreatIntelMapInput)(nil)).Elem(), AppSecThreatIntelMap{})
	pulumi.RegisterOutputType(AppSecThreatIntelOutput{})
	pulumi.RegisterOutputType(AppSecThreatIntelArrayOutput{})
	pulumi.RegisterOutputType(AppSecThreatIntelMapOutput{})
}
