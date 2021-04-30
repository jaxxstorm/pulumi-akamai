// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecEvalRuleConditionException` data source to list the conditions and exceptions to a rule you want to evaluate.
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
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &akamai.LookupAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		conditionException, err := akamai.LookupAppSecEvalRuleConditionException(ctx, &akamai.LookupAppSecEvalRuleConditionExceptionArgs{
// 			ConfigId:         configuration.ConfigId,
// 			Version:          configuration.LatestVersion,
// 			SecurityPolicyId: _var.Security_policy_id,
// 			RuleId:           _var.Rule_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("conditionExceptionText", conditionException.OutputText)
// 		ctx.Export("conditionExceptionJson", conditionException.Json)
// 		return nil
// 	})
// }
// ```
func LookupAppSecEvalRuleConditionException(ctx *pulumi.Context, args *LookupAppSecEvalRuleConditionExceptionArgs, opts ...pulumi.InvokeOption) (*LookupAppSecEvalRuleConditionExceptionResult, error) {
	var rv LookupAppSecEvalRuleConditionExceptionResult
	err := ctx.Invoke("akamai:index/getAppSecEvalRuleConditionException:getAppSecEvalRuleConditionException", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecEvalRuleConditionException.
type LookupAppSecEvalRuleConditionExceptionArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the rule to use.
	RuleId int `pulumi:"ruleId"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	// The version number of the security configuration to use.
	Version int `pulumi:"version"`
}

// A collection of values returned by getAppSecEvalRuleConditionException.
type LookupAppSecEvalRuleConditionExceptionResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of the condition and exception information for the specified rule.
	Json string `pulumi:"json"`
	// A tabular display showing boolean values indicating whether conditions and exceptions are present.
	OutputText       string `pulumi:"outputText"`
	RuleId           int    `pulumi:"ruleId"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	Version          int    `pulumi:"version"`
}
