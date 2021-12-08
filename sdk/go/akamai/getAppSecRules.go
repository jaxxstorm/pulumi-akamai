// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppSecRules(ctx *pulumi.Context, args *GetAppSecRulesArgs, opts ...pulumi.InvokeOption) (*GetAppSecRulesResult, error) {
	var rv GetAppSecRulesResult
	err := ctx.Invoke("akamai:index/getAppSecRules:getAppSecRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecRules.
type GetAppSecRulesArgs struct {
	// . Unique identifier of the security configuration associated with the rules.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the Kona Rule Set rule you want to return information for. If not included, information is returned for all your KRS rules.
	RuleId *int `pulumi:"ruleId"`
	// . Unique identifier of the security policy associated with the rules.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecRules.
type GetAppSecRulesResult struct {
	ConditionException string `pulumi:"conditionException"`
	ConfigId           int    `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	Json             string `pulumi:"json"`
	OutputText       string `pulumi:"outputText"`
	RuleAction       string `pulumi:"ruleAction"`
	RuleId           *int   `pulumi:"ruleId"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

func GetAppSecRulesOutput(ctx *pulumi.Context, args GetAppSecRulesOutputArgs, opts ...pulumi.InvokeOption) GetAppSecRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecRulesResult, error) {
			args := v.(GetAppSecRulesArgs)
			r, err := GetAppSecRules(ctx, &args, opts...)
			return *r, err
		}).(GetAppSecRulesResultOutput)
}

// A collection of arguments for invoking getAppSecRules.
type GetAppSecRulesOutputArgs struct {
	// . Unique identifier of the security configuration associated with the rules.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the Kona Rule Set rule you want to return information for. If not included, information is returned for all your KRS rules.
	RuleId pulumi.IntPtrInput `pulumi:"ruleId"`
	// . Unique identifier of the security policy associated with the rules.
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (GetAppSecRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecRulesArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecRules.
type GetAppSecRulesResultOutput struct{ *pulumi.OutputState }

func (GetAppSecRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecRulesResult)(nil)).Elem()
}

func (o GetAppSecRulesResultOutput) ToGetAppSecRulesResultOutput() GetAppSecRulesResultOutput {
	return o
}

func (o GetAppSecRulesResultOutput) ToGetAppSecRulesResultOutputWithContext(ctx context.Context) GetAppSecRulesResultOutput {
	return o
}

func (o GetAppSecRulesResultOutput) ConditionException() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRulesResult) string { return v.ConditionException }).(pulumi.StringOutput)
}

func (o GetAppSecRulesResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecRulesResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecRulesResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRulesResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetAppSecRulesResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRulesResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecRulesResultOutput) RuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRulesResult) string { return v.RuleAction }).(pulumi.StringOutput)
}

func (o GetAppSecRulesResultOutput) RuleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecRulesResult) *int { return v.RuleId }).(pulumi.IntPtrOutput)
}

func (o GetAppSecRulesResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRulesResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecRulesResultOutput{})
}
