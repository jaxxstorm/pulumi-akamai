// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAppSecEvalRules` data source to list the action and condition-exception information
// for a rule or rules you want to evaluate.
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
// 		opt0 := _var.Security_configuration
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := _var.Rule_id
// 		_, err = akamai.GetAppSecEvalRules(ctx, &GetAppSecEvalRulesArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: _var.Security_policy_id,
// 			RuleId:           &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("evalRuleAction", akamai_appsec_eval_rules.Eval_rule.Eval_rule_action)
// 		ctx.Export("conditionException", akamai_appsec_eval_rules.Eval_rule.Condition_exception)
// 		ctx.Export("json", akamai_appsec_eval_rules.Eval_rule.Json)
// 		ctx.Export("outputText", akamai_appsec_eval_rules.Eval_rule.Output_text)
// 		return nil
// 	})
// }
// ```
func GetAppSecEvalRules(ctx *pulumi.Context, args *GetAppSecEvalRulesArgs, opts ...pulumi.InvokeOption) (*GetAppSecEvalRulesResult, error) {
	var rv GetAppSecEvalRulesResult
	err := ctx.Invoke("akamai:index/getAppSecEvalRules:getAppSecEvalRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecEvalRules.
type GetAppSecEvalRulesArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the rule to use. If not specified, information about all rules will be returned.
	RuleId *int `pulumi:"ruleId"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecEvalRules.
type GetAppSecEvalRulesResult struct {
	// The eval rule's conditions and exceptions.
	ConditionException string `pulumi:"conditionException"`
	ConfigId           int    `pulumi:"configId"`
	// The eval rule's action, either `alert`, `deny`, or `none`.
	EvalRuleAction string `pulumi:"evalRuleAction"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of the action and condition-exception information for the specified eval rule.
	// This output is only generated if an eval rule is specified.
	Json string `pulumi:"json"`
	// A tabular display showing, for the specified eval rule or rules, the rule action and boolean values
	// indicating whether conditions and exceptions are present.
	OutputText       string `pulumi:"outputText"`
	RuleId           *int   `pulumi:"ruleId"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

func GetAppSecEvalRulesOutput(ctx *pulumi.Context, args GetAppSecEvalRulesOutputArgs, opts ...pulumi.InvokeOption) GetAppSecEvalRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecEvalRulesResult, error) {
			args := v.(GetAppSecEvalRulesArgs)
			r, err := GetAppSecEvalRules(ctx, &args, opts...)
			return *r, err
		}).(GetAppSecEvalRulesResultOutput)
}

// A collection of arguments for invoking getAppSecEvalRules.
type GetAppSecEvalRulesOutputArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// The ID of the rule to use. If not specified, information about all rules will be returned.
	RuleId pulumi.IntPtrInput `pulumi:"ruleId"`
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (GetAppSecEvalRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecEvalRulesArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecEvalRules.
type GetAppSecEvalRulesResultOutput struct{ *pulumi.OutputState }

func (GetAppSecEvalRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecEvalRulesResult)(nil)).Elem()
}

func (o GetAppSecEvalRulesResultOutput) ToGetAppSecEvalRulesResultOutput() GetAppSecEvalRulesResultOutput {
	return o
}

func (o GetAppSecEvalRulesResultOutput) ToGetAppSecEvalRulesResultOutputWithContext(ctx context.Context) GetAppSecEvalRulesResultOutput {
	return o
}

// The eval rule's conditions and exceptions.
func (o GetAppSecEvalRulesResultOutput) ConditionException() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalRulesResult) string { return v.ConditionException }).(pulumi.StringOutput)
}

func (o GetAppSecEvalRulesResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecEvalRulesResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The eval rule's action, either `alert`, `deny`, or `none`.
func (o GetAppSecEvalRulesResultOutput) EvalRuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalRulesResult) string { return v.EvalRuleAction }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecEvalRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A JSON-formatted list of the action and condition-exception information for the specified eval rule.
// This output is only generated if an eval rule is specified.
func (o GetAppSecEvalRulesResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalRulesResult) string { return v.Json }).(pulumi.StringOutput)
}

// A tabular display showing, for the specified eval rule or rules, the rule action and boolean values
// indicating whether conditions and exceptions are present.
func (o GetAppSecEvalRulesResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalRulesResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecEvalRulesResultOutput) RuleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecEvalRulesResult) *int { return v.RuleId }).(pulumi.IntPtrOutput)
}

func (o GetAppSecEvalRulesResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalRulesResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecEvalRulesResultOutput{})
}
