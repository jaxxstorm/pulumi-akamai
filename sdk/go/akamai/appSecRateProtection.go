// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy
//
// Enables or disables rate protection for a security policy.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/protections](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putprotections)
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
// 		opt0 := "Documentation"
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecRateProtection(ctx, "protection", &akamai.AppSecRateProtectionArgs{
// 			ConfigId:         pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId: pulumi.String("gms1_134637"),
// 			Enabled:          pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `outputText`. Tabular report showing the current protection settings.
type AppSecRateProtection struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the rate protection settings being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Set to **true** to enable rate protection; set to **false** to disable rate protection.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Text Export representation
	OutputText pulumi.StringOutput `pulumi:"outputText"`
	// . Unique identifier of the security policy associated with the rate protection settings being modified.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecRateProtection registers a new resource with the given unique name, arguments, and options.
func NewAppSecRateProtection(ctx *pulumi.Context,
	name string, args *AppSecRateProtectionArgs, opts ...pulumi.ResourceOption) (*AppSecRateProtection, error) {
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
	var resource AppSecRateProtection
	err := ctx.RegisterResource("akamai:index/appSecRateProtection:AppSecRateProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecRateProtection gets an existing AppSecRateProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecRateProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecRateProtectionState, opts ...pulumi.ResourceOption) (*AppSecRateProtection, error) {
	var resource AppSecRateProtection
	err := ctx.ReadResource("akamai:index/appSecRateProtection:AppSecRateProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecRateProtection resources.
type appSecRateProtectionState struct {
	// . Unique identifier of the security configuration associated with the rate protection settings being modified.
	ConfigId *int `pulumi:"configId"`
	// . Set to **true** to enable rate protection; set to **false** to disable rate protection.
	Enabled *bool `pulumi:"enabled"`
	// Text Export representation
	OutputText *string `pulumi:"outputText"`
	// . Unique identifier of the security policy associated with the rate protection settings being modified.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecRateProtectionState struct {
	// . Unique identifier of the security configuration associated with the rate protection settings being modified.
	ConfigId pulumi.IntPtrInput
	// . Set to **true** to enable rate protection; set to **false** to disable rate protection.
	Enabled pulumi.BoolPtrInput
	// Text Export representation
	OutputText pulumi.StringPtrInput
	// . Unique identifier of the security policy associated with the rate protection settings being modified.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecRateProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRateProtectionState)(nil)).Elem()
}

type appSecRateProtectionArgs struct {
	// . Unique identifier of the security configuration associated with the rate protection settings being modified.
	ConfigId int `pulumi:"configId"`
	// . Set to **true** to enable rate protection; set to **false** to disable rate protection.
	Enabled bool `pulumi:"enabled"`
	// . Unique identifier of the security policy associated with the rate protection settings being modified.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecRateProtection resource.
type AppSecRateProtectionArgs struct {
	// . Unique identifier of the security configuration associated with the rate protection settings being modified.
	ConfigId pulumi.IntInput
	// . Set to **true** to enable rate protection; set to **false** to disable rate protection.
	Enabled pulumi.BoolInput
	// . Unique identifier of the security policy associated with the rate protection settings being modified.
	SecurityPolicyId pulumi.StringInput
}

func (AppSecRateProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRateProtectionArgs)(nil)).Elem()
}

type AppSecRateProtectionInput interface {
	pulumi.Input

	ToAppSecRateProtectionOutput() AppSecRateProtectionOutput
	ToAppSecRateProtectionOutputWithContext(ctx context.Context) AppSecRateProtectionOutput
}

func (*AppSecRateProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRateProtection)(nil)).Elem()
}

func (i *AppSecRateProtection) ToAppSecRateProtectionOutput() AppSecRateProtectionOutput {
	return i.ToAppSecRateProtectionOutputWithContext(context.Background())
}

func (i *AppSecRateProtection) ToAppSecRateProtectionOutputWithContext(ctx context.Context) AppSecRateProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRateProtectionOutput)
}

// AppSecRateProtectionArrayInput is an input type that accepts AppSecRateProtectionArray and AppSecRateProtectionArrayOutput values.
// You can construct a concrete instance of `AppSecRateProtectionArrayInput` via:
//
//          AppSecRateProtectionArray{ AppSecRateProtectionArgs{...} }
type AppSecRateProtectionArrayInput interface {
	pulumi.Input

	ToAppSecRateProtectionArrayOutput() AppSecRateProtectionArrayOutput
	ToAppSecRateProtectionArrayOutputWithContext(context.Context) AppSecRateProtectionArrayOutput
}

type AppSecRateProtectionArray []AppSecRateProtectionInput

func (AppSecRateProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecRateProtection)(nil)).Elem()
}

func (i AppSecRateProtectionArray) ToAppSecRateProtectionArrayOutput() AppSecRateProtectionArrayOutput {
	return i.ToAppSecRateProtectionArrayOutputWithContext(context.Background())
}

func (i AppSecRateProtectionArray) ToAppSecRateProtectionArrayOutputWithContext(ctx context.Context) AppSecRateProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRateProtectionArrayOutput)
}

// AppSecRateProtectionMapInput is an input type that accepts AppSecRateProtectionMap and AppSecRateProtectionMapOutput values.
// You can construct a concrete instance of `AppSecRateProtectionMapInput` via:
//
//          AppSecRateProtectionMap{ "key": AppSecRateProtectionArgs{...} }
type AppSecRateProtectionMapInput interface {
	pulumi.Input

	ToAppSecRateProtectionMapOutput() AppSecRateProtectionMapOutput
	ToAppSecRateProtectionMapOutputWithContext(context.Context) AppSecRateProtectionMapOutput
}

type AppSecRateProtectionMap map[string]AppSecRateProtectionInput

func (AppSecRateProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecRateProtection)(nil)).Elem()
}

func (i AppSecRateProtectionMap) ToAppSecRateProtectionMapOutput() AppSecRateProtectionMapOutput {
	return i.ToAppSecRateProtectionMapOutputWithContext(context.Background())
}

func (i AppSecRateProtectionMap) ToAppSecRateProtectionMapOutputWithContext(ctx context.Context) AppSecRateProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRateProtectionMapOutput)
}

type AppSecRateProtectionOutput struct{ *pulumi.OutputState }

func (AppSecRateProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRateProtection)(nil)).Elem()
}

func (o AppSecRateProtectionOutput) ToAppSecRateProtectionOutput() AppSecRateProtectionOutput {
	return o
}

func (o AppSecRateProtectionOutput) ToAppSecRateProtectionOutputWithContext(ctx context.Context) AppSecRateProtectionOutput {
	return o
}

type AppSecRateProtectionArrayOutput struct{ *pulumi.OutputState }

func (AppSecRateProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecRateProtection)(nil)).Elem()
}

func (o AppSecRateProtectionArrayOutput) ToAppSecRateProtectionArrayOutput() AppSecRateProtectionArrayOutput {
	return o
}

func (o AppSecRateProtectionArrayOutput) ToAppSecRateProtectionArrayOutputWithContext(ctx context.Context) AppSecRateProtectionArrayOutput {
	return o
}

func (o AppSecRateProtectionArrayOutput) Index(i pulumi.IntInput) AppSecRateProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecRateProtection {
		return vs[0].([]*AppSecRateProtection)[vs[1].(int)]
	}).(AppSecRateProtectionOutput)
}

type AppSecRateProtectionMapOutput struct{ *pulumi.OutputState }

func (AppSecRateProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecRateProtection)(nil)).Elem()
}

func (o AppSecRateProtectionMapOutput) ToAppSecRateProtectionMapOutput() AppSecRateProtectionMapOutput {
	return o
}

func (o AppSecRateProtectionMapOutput) ToAppSecRateProtectionMapOutputWithContext(ctx context.Context) AppSecRateProtectionMapOutput {
	return o
}

func (o AppSecRateProtectionMapOutput) MapIndex(k pulumi.StringInput) AppSecRateProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecRateProtection {
		return vs[0].(map[string]*AppSecRateProtection)[vs[1].(string)]
	}).(AppSecRateProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRateProtectionInput)(nil)).Elem(), &AppSecRateProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRateProtectionArrayInput)(nil)).Elem(), AppSecRateProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRateProtectionMapInput)(nil)).Elem(), AppSecRateProtectionMap{})
	pulumi.RegisterOutputType(AppSecRateProtectionOutput{})
	pulumi.RegisterOutputType(AppSecRateProtectionArrayOutput{})
	pulumi.RegisterOutputType(AppSecRateProtectionMapOutput{})
}
