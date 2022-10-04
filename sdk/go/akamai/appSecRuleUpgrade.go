// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			ruleUpgrade, err := akamai.NewAppSecRuleUpgrade(ctx, "ruleUpgrade", &akamai.AppSecRuleUpgradeArgs{
//				ConfigId:         pulumi.Int(configuration.ConfigId),
//				SecurityPolicyId: pulumi.String("gms1_134637"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("ruleUpgradeCurrentRuleset", ruleUpgrade.CurrentRuleset)
//			ctx.Export("ruleUpgradeMode", ruleUpgrade.Mode)
//			ctx.Export("ruleUpgradeEvalStatus", ruleUpgrade.EvalStatus)
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned and how that returned information is formatted:
//
// - `currentRuleset`. Versioning information for your current KRS rule set.
// - `mode`. Specifies the current upgrade mode type. Valid values are:
//
//   - **KRS**. Rulesets must be manually upgraded.
//
//   - **AAG**. Rulesets are automatically upgraded by Akamai.
//
//   - **ASE_MANUAL**. Adaptive Security Engine rulesets must be manually upgraded.
//
//   - **ASE_AUTO**. Adaptive Security Engine rulesets are automatically updated by Akamai.
//
// - `evalStatus`. Returns **enabled** if an evaluation is currently in progress; otherwise returns **disabled**.
type AppSecRuleUpgrade struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the ruleset being upgraded.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// Versioning information for the current KRS rule set
	CurrentRuleset pulumi.StringOutput `pulumi:"currentRuleset"`
	// Whether an evaluation is currently in progress
	EvalStatus pulumi.StringOutput `pulumi:"evalStatus"`
	// Upgrade mode (KRS, AAG, ASE_MANUAL or ASE_AUTO)
	Mode pulumi.StringOutput `pulumi:"mode"`
	// . Unique identifier of the security policy associated with the ruleset being upgraded.
	// - `upgradeMode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
	// - **ASE_AUTO**. Akamai automatically updates your rulesets.
	// - **ASE_MANUAL**. Manually updates your rulesets.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
	// Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
	UpgradeMode pulumi.StringPtrOutput `pulumi:"upgradeMode"`
}

// NewAppSecRuleUpgrade registers a new resource with the given unique name, arguments, and options.
func NewAppSecRuleUpgrade(ctx *pulumi.Context,
	name string, args *AppSecRuleUpgradeArgs, opts ...pulumi.ResourceOption) (*AppSecRuleUpgrade, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	var resource AppSecRuleUpgrade
	err := ctx.RegisterResource("akamai:index/appSecRuleUpgrade:AppSecRuleUpgrade", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecRuleUpgrade gets an existing AppSecRuleUpgrade resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecRuleUpgrade(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecRuleUpgradeState, opts ...pulumi.ResourceOption) (*AppSecRuleUpgrade, error) {
	var resource AppSecRuleUpgrade
	err := ctx.ReadResource("akamai:index/appSecRuleUpgrade:AppSecRuleUpgrade", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecRuleUpgrade resources.
type appSecRuleUpgradeState struct {
	// . Unique identifier of the security configuration associated with the ruleset being upgraded.
	ConfigId *int `pulumi:"configId"`
	// Versioning information for the current KRS rule set
	CurrentRuleset *string `pulumi:"currentRuleset"`
	// Whether an evaluation is currently in progress
	EvalStatus *string `pulumi:"evalStatus"`
	// Upgrade mode (KRS, AAG, ASE_MANUAL or ASE_AUTO)
	Mode *string `pulumi:"mode"`
	// . Unique identifier of the security policy associated with the ruleset being upgraded.
	// - `upgradeMode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
	// - **ASE_AUTO**. Akamai automatically updates your rulesets.
	// - **ASE_MANUAL**. Manually updates your rulesets.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
	// Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
	UpgradeMode *string `pulumi:"upgradeMode"`
}

type AppSecRuleUpgradeState struct {
	// . Unique identifier of the security configuration associated with the ruleset being upgraded.
	ConfigId pulumi.IntPtrInput
	// Versioning information for the current KRS rule set
	CurrentRuleset pulumi.StringPtrInput
	// Whether an evaluation is currently in progress
	EvalStatus pulumi.StringPtrInput
	// Upgrade mode (KRS, AAG, ASE_MANUAL or ASE_AUTO)
	Mode pulumi.StringPtrInput
	// . Unique identifier of the security policy associated with the ruleset being upgraded.
	// - `upgradeMode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
	// - **ASE_AUTO**. Akamai automatically updates your rulesets.
	// - **ASE_MANUAL**. Manually updates your rulesets.
	SecurityPolicyId pulumi.StringPtrInput
	// Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
	UpgradeMode pulumi.StringPtrInput
}

func (AppSecRuleUpgradeState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRuleUpgradeState)(nil)).Elem()
}

type appSecRuleUpgradeArgs struct {
	// . Unique identifier of the security configuration associated with the ruleset being upgraded.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the ruleset being upgraded.
	// - `upgradeMode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
	// - **ASE_AUTO**. Akamai automatically updates your rulesets.
	// - **ASE_MANUAL**. Manually updates your rulesets.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	// Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
	UpgradeMode *string `pulumi:"upgradeMode"`
}

// The set of arguments for constructing a AppSecRuleUpgrade resource.
type AppSecRuleUpgradeArgs struct {
	// . Unique identifier of the security configuration associated with the ruleset being upgraded.
	ConfigId pulumi.IntInput
	// . Unique identifier of the security policy associated with the ruleset being upgraded.
	// - `upgradeMode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
	// - **ASE_AUTO**. Akamai automatically updates your rulesets.
	// - **ASE_MANUAL**. Manually updates your rulesets.
	SecurityPolicyId pulumi.StringInput
	// Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
	UpgradeMode pulumi.StringPtrInput
}

func (AppSecRuleUpgradeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRuleUpgradeArgs)(nil)).Elem()
}

type AppSecRuleUpgradeInput interface {
	pulumi.Input

	ToAppSecRuleUpgradeOutput() AppSecRuleUpgradeOutput
	ToAppSecRuleUpgradeOutputWithContext(ctx context.Context) AppSecRuleUpgradeOutput
}

func (*AppSecRuleUpgrade) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRuleUpgrade)(nil)).Elem()
}

func (i *AppSecRuleUpgrade) ToAppSecRuleUpgradeOutput() AppSecRuleUpgradeOutput {
	return i.ToAppSecRuleUpgradeOutputWithContext(context.Background())
}

func (i *AppSecRuleUpgrade) ToAppSecRuleUpgradeOutputWithContext(ctx context.Context) AppSecRuleUpgradeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRuleUpgradeOutput)
}

// AppSecRuleUpgradeArrayInput is an input type that accepts AppSecRuleUpgradeArray and AppSecRuleUpgradeArrayOutput values.
// You can construct a concrete instance of `AppSecRuleUpgradeArrayInput` via:
//
//	AppSecRuleUpgradeArray{ AppSecRuleUpgradeArgs{...} }
type AppSecRuleUpgradeArrayInput interface {
	pulumi.Input

	ToAppSecRuleUpgradeArrayOutput() AppSecRuleUpgradeArrayOutput
	ToAppSecRuleUpgradeArrayOutputWithContext(context.Context) AppSecRuleUpgradeArrayOutput
}

type AppSecRuleUpgradeArray []AppSecRuleUpgradeInput

func (AppSecRuleUpgradeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecRuleUpgrade)(nil)).Elem()
}

func (i AppSecRuleUpgradeArray) ToAppSecRuleUpgradeArrayOutput() AppSecRuleUpgradeArrayOutput {
	return i.ToAppSecRuleUpgradeArrayOutputWithContext(context.Background())
}

func (i AppSecRuleUpgradeArray) ToAppSecRuleUpgradeArrayOutputWithContext(ctx context.Context) AppSecRuleUpgradeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRuleUpgradeArrayOutput)
}

// AppSecRuleUpgradeMapInput is an input type that accepts AppSecRuleUpgradeMap and AppSecRuleUpgradeMapOutput values.
// You can construct a concrete instance of `AppSecRuleUpgradeMapInput` via:
//
//	AppSecRuleUpgradeMap{ "key": AppSecRuleUpgradeArgs{...} }
type AppSecRuleUpgradeMapInput interface {
	pulumi.Input

	ToAppSecRuleUpgradeMapOutput() AppSecRuleUpgradeMapOutput
	ToAppSecRuleUpgradeMapOutputWithContext(context.Context) AppSecRuleUpgradeMapOutput
}

type AppSecRuleUpgradeMap map[string]AppSecRuleUpgradeInput

func (AppSecRuleUpgradeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecRuleUpgrade)(nil)).Elem()
}

func (i AppSecRuleUpgradeMap) ToAppSecRuleUpgradeMapOutput() AppSecRuleUpgradeMapOutput {
	return i.ToAppSecRuleUpgradeMapOutputWithContext(context.Background())
}

func (i AppSecRuleUpgradeMap) ToAppSecRuleUpgradeMapOutputWithContext(ctx context.Context) AppSecRuleUpgradeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRuleUpgradeMapOutput)
}

type AppSecRuleUpgradeOutput struct{ *pulumi.OutputState }

func (AppSecRuleUpgradeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRuleUpgrade)(nil)).Elem()
}

func (o AppSecRuleUpgradeOutput) ToAppSecRuleUpgradeOutput() AppSecRuleUpgradeOutput {
	return o
}

func (o AppSecRuleUpgradeOutput) ToAppSecRuleUpgradeOutputWithContext(ctx context.Context) AppSecRuleUpgradeOutput {
	return o
}

// . Unique identifier of the security configuration associated with the ruleset being upgraded.
func (o AppSecRuleUpgradeOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecRuleUpgrade) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// Versioning information for the current KRS rule set
func (o AppSecRuleUpgradeOutput) CurrentRuleset() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecRuleUpgrade) pulumi.StringOutput { return v.CurrentRuleset }).(pulumi.StringOutput)
}

// Whether an evaluation is currently in progress
func (o AppSecRuleUpgradeOutput) EvalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecRuleUpgrade) pulumi.StringOutput { return v.EvalStatus }).(pulumi.StringOutput)
}

// Upgrade mode (KRS, AAG, ASE_MANUAL or ASE_AUTO)
func (o AppSecRuleUpgradeOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecRuleUpgrade) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// . Unique identifier of the security policy associated with the ruleset being upgraded.
// - `upgradeMode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
// - **ASE_AUTO**. Akamai automatically updates your rulesets.
// - **ASE_MANUAL**. Manually updates your rulesets.
func (o AppSecRuleUpgradeOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecRuleUpgrade) pulumi.StringOutput { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

// Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
func (o AppSecRuleUpgradeOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSecRuleUpgrade) pulumi.StringPtrOutput { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

type AppSecRuleUpgradeArrayOutput struct{ *pulumi.OutputState }

func (AppSecRuleUpgradeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecRuleUpgrade)(nil)).Elem()
}

func (o AppSecRuleUpgradeArrayOutput) ToAppSecRuleUpgradeArrayOutput() AppSecRuleUpgradeArrayOutput {
	return o
}

func (o AppSecRuleUpgradeArrayOutput) ToAppSecRuleUpgradeArrayOutputWithContext(ctx context.Context) AppSecRuleUpgradeArrayOutput {
	return o
}

func (o AppSecRuleUpgradeArrayOutput) Index(i pulumi.IntInput) AppSecRuleUpgradeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecRuleUpgrade {
		return vs[0].([]*AppSecRuleUpgrade)[vs[1].(int)]
	}).(AppSecRuleUpgradeOutput)
}

type AppSecRuleUpgradeMapOutput struct{ *pulumi.OutputState }

func (AppSecRuleUpgradeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecRuleUpgrade)(nil)).Elem()
}

func (o AppSecRuleUpgradeMapOutput) ToAppSecRuleUpgradeMapOutput() AppSecRuleUpgradeMapOutput {
	return o
}

func (o AppSecRuleUpgradeMapOutput) ToAppSecRuleUpgradeMapOutputWithContext(ctx context.Context) AppSecRuleUpgradeMapOutput {
	return o
}

func (o AppSecRuleUpgradeMapOutput) MapIndex(k pulumi.StringInput) AppSecRuleUpgradeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecRuleUpgrade {
		return vs[0].(map[string]*AppSecRuleUpgrade)[vs[1].(string)]
	}).(AppSecRuleUpgradeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRuleUpgradeInput)(nil)).Elem(), &AppSecRuleUpgrade{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRuleUpgradeArrayInput)(nil)).Elem(), AppSecRuleUpgradeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRuleUpgradeMapInput)(nil)).Elem(), AppSecRuleUpgradeMap{})
	pulumi.RegisterOutputType(AppSecRuleUpgradeOutput{})
	pulumi.RegisterOutputType(AppSecRuleUpgradeArrayOutput{})
	pulumi.RegisterOutputType(AppSecRuleUpgradeMapOutput{})
}
