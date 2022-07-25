// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Rule
//
// Modifies a Kona Rule Set rule's action, conditions, and exceptions.
//
// **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/{ruleId}](https://techdocs.akamai.com/application-security/reference/put-rule) *and* [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/{ruleId}/condition-exception](https://techdocs.akamai.com/application-security/reference/put-rule-condition-exception)
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: pulumi.StringRef("Documentation"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecRule(ctx, "rule", &akamai.AppSecRuleArgs{
// 			ConfigId:           pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId:   pulumi.String("gms1_134637"),
// 			RuleId:             pulumi.Int(60029316),
// 			RuleAction:         pulumi.String("deny"),
// 			ConditionException: readFileOrPanic(fmt.Sprintf("%v/condition_exception.json", path.Module)),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecRule struct {
	pulumi.CustomResourceState

	// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
	ConditionException pulumi.StringPtrOutput `pulumi:"conditionException"`
	// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// Allowed values are:
	// - **alert**. Record the event.
	// - **deny**. Block the request.
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action. or `none` to take no action.
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// . Unique identifier of the rule being modified.
	RuleId pulumi.IntOutput `pulumi:"ruleId"`
	// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecRule registers a new resource with the given unique name, arguments, and options.
func NewAppSecRule(ctx *pulumi.Context,
	name string, args *AppSecRuleArgs, opts ...pulumi.ResourceOption) (*AppSecRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	var resource AppSecRule
	err := ctx.RegisterResource("akamai:index/appSecRule:AppSecRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecRule gets an existing AppSecRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecRuleState, opts ...pulumi.ResourceOption) (*AppSecRule, error) {
	var resource AppSecRule
	err := ctx.ReadResource("akamai:index/appSecRule:AppSecRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecRule resources.
type appSecRuleState struct {
	// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
	ConditionException *string `pulumi:"conditionException"`
	// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
	ConfigId *int `pulumi:"configId"`
	// Allowed values are:
	// - **alert**. Record the event.
	// - **deny**. Block the request.
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action. or `none` to take no action.
	RuleAction *string `pulumi:"ruleAction"`
	// . Unique identifier of the rule being modified.
	RuleId *int `pulumi:"ruleId"`
	// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecRuleState struct {
	// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
	ConditionException pulumi.StringPtrInput
	// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
	ConfigId pulumi.IntPtrInput
	// Allowed values are:
	// - **alert**. Record the event.
	// - **deny**. Block the request.
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action. or `none` to take no action.
	RuleAction pulumi.StringPtrInput
	// . Unique identifier of the rule being modified.
	RuleId pulumi.IntPtrInput
	// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRuleState)(nil)).Elem()
}

type appSecRuleArgs struct {
	// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
	ConditionException *string `pulumi:"conditionException"`
	// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
	ConfigId int `pulumi:"configId"`
	// Allowed values are:
	// - **alert**. Record the event.
	// - **deny**. Block the request.
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action. or `none` to take no action.
	RuleAction *string `pulumi:"ruleAction"`
	// . Unique identifier of the rule being modified.
	RuleId int `pulumi:"ruleId"`
	// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecRule resource.
type AppSecRuleArgs struct {
	// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
	ConditionException pulumi.StringPtrInput
	// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
	ConfigId pulumi.IntInput
	// Allowed values are:
	// - **alert**. Record the event.
	// - **deny**. Block the request.
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action. or `none` to take no action.
	RuleAction pulumi.StringPtrInput
	// . Unique identifier of the rule being modified.
	RuleId pulumi.IntInput
	// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
	SecurityPolicyId pulumi.StringInput
}

func (AppSecRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRuleArgs)(nil)).Elem()
}

type AppSecRuleInput interface {
	pulumi.Input

	ToAppSecRuleOutput() AppSecRuleOutput
	ToAppSecRuleOutputWithContext(ctx context.Context) AppSecRuleOutput
}

func (*AppSecRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRule)(nil)).Elem()
}

func (i *AppSecRule) ToAppSecRuleOutput() AppSecRuleOutput {
	return i.ToAppSecRuleOutputWithContext(context.Background())
}

func (i *AppSecRule) ToAppSecRuleOutputWithContext(ctx context.Context) AppSecRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRuleOutput)
}

// AppSecRuleArrayInput is an input type that accepts AppSecRuleArray and AppSecRuleArrayOutput values.
// You can construct a concrete instance of `AppSecRuleArrayInput` via:
//
//          AppSecRuleArray{ AppSecRuleArgs{...} }
type AppSecRuleArrayInput interface {
	pulumi.Input

	ToAppSecRuleArrayOutput() AppSecRuleArrayOutput
	ToAppSecRuleArrayOutputWithContext(context.Context) AppSecRuleArrayOutput
}

type AppSecRuleArray []AppSecRuleInput

func (AppSecRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecRule)(nil)).Elem()
}

func (i AppSecRuleArray) ToAppSecRuleArrayOutput() AppSecRuleArrayOutput {
	return i.ToAppSecRuleArrayOutputWithContext(context.Background())
}

func (i AppSecRuleArray) ToAppSecRuleArrayOutputWithContext(ctx context.Context) AppSecRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRuleArrayOutput)
}

// AppSecRuleMapInput is an input type that accepts AppSecRuleMap and AppSecRuleMapOutput values.
// You can construct a concrete instance of `AppSecRuleMapInput` via:
//
//          AppSecRuleMap{ "key": AppSecRuleArgs{...} }
type AppSecRuleMapInput interface {
	pulumi.Input

	ToAppSecRuleMapOutput() AppSecRuleMapOutput
	ToAppSecRuleMapOutputWithContext(context.Context) AppSecRuleMapOutput
}

type AppSecRuleMap map[string]AppSecRuleInput

func (AppSecRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecRule)(nil)).Elem()
}

func (i AppSecRuleMap) ToAppSecRuleMapOutput() AppSecRuleMapOutput {
	return i.ToAppSecRuleMapOutputWithContext(context.Background())
}

func (i AppSecRuleMap) ToAppSecRuleMapOutputWithContext(ctx context.Context) AppSecRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRuleMapOutput)
}

type AppSecRuleOutput struct{ *pulumi.OutputState }

func (AppSecRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRule)(nil)).Elem()
}

func (o AppSecRuleOutput) ToAppSecRuleOutput() AppSecRuleOutput {
	return o
}

func (o AppSecRuleOutput) ToAppSecRuleOutputWithContext(ctx context.Context) AppSecRuleOutput {
	return o
}

// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
func (o AppSecRuleOutput) ConditionException() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSecRule) pulumi.StringPtrOutput { return v.ConditionException }).(pulumi.StringPtrOutput)
}

// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
func (o AppSecRuleOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecRule) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// Allowed values are:
// - **alert**. Record the event.
// - **deny**. Block the request.
// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
// - **none**. Take no action. or `none` to take no action.
func (o AppSecRuleOutput) RuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecRule) pulumi.StringOutput { return v.RuleAction }).(pulumi.StringOutput)
}

// . Unique identifier of the rule being modified.
func (o AppSecRuleOutput) RuleId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecRule) pulumi.IntOutput { return v.RuleId }).(pulumi.IntOutput)
}

// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
func (o AppSecRuleOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecRule) pulumi.StringOutput { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

type AppSecRuleArrayOutput struct{ *pulumi.OutputState }

func (AppSecRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecRule)(nil)).Elem()
}

func (o AppSecRuleArrayOutput) ToAppSecRuleArrayOutput() AppSecRuleArrayOutput {
	return o
}

func (o AppSecRuleArrayOutput) ToAppSecRuleArrayOutputWithContext(ctx context.Context) AppSecRuleArrayOutput {
	return o
}

func (o AppSecRuleArrayOutput) Index(i pulumi.IntInput) AppSecRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecRule {
		return vs[0].([]*AppSecRule)[vs[1].(int)]
	}).(AppSecRuleOutput)
}

type AppSecRuleMapOutput struct{ *pulumi.OutputState }

func (AppSecRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecRule)(nil)).Elem()
}

func (o AppSecRuleMapOutput) ToAppSecRuleMapOutput() AppSecRuleMapOutput {
	return o
}

func (o AppSecRuleMapOutput) ToAppSecRuleMapOutputWithContext(ctx context.Context) AppSecRuleMapOutput {
	return o
}

func (o AppSecRuleMapOutput) MapIndex(k pulumi.StringInput) AppSecRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecRule {
		return vs[0].(map[string]*AppSecRule)[vs[1].(string)]
	}).(AppSecRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRuleInput)(nil)).Elem(), &AppSecRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRuleArrayInput)(nil)).Elem(), AppSecRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRuleMapInput)(nil)).Elem(), AppSecRuleMap{})
	pulumi.RegisterOutputType(AppSecRuleOutput{})
	pulumi.RegisterOutputType(AppSecRuleArrayOutput{})
	pulumi.RegisterOutputType(AppSecRuleMapOutput{})
}
