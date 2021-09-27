// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecWafMode` data source to retrieve the mode that indicates how the WAF rules of the given security configuration and security policy will be updated.
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
// 		wafMode, err := akamai.LookupAppSecWafMode(ctx, &akamai.LookupAppSecWafModeArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: _var.Policy_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("wafModeMode", wafMode.Mode)
// 		ctx.Export("wafModeCurrentRuleset", wafMode.CurrentRuleset)
// 		ctx.Export("wafModeEvalStatus", wafMode.EvalStatus)
// 		ctx.Export("wafModeEvalRuleset", wafMode.EvalRuleset)
// 		ctx.Export("wafModeEvalExpirationDate", wafMode.EvalExpirationDate)
// 		ctx.Export("wafModeText", wafMode.OutputText)
// 		ctx.Export("wafModeJson", wafMode.Json)
// 		return nil
// 	})
// }
// ```
func LookupAppSecWafMode(ctx *pulumi.Context, args *LookupAppSecWafModeArgs, opts ...pulumi.InvokeOption) (*LookupAppSecWafModeResult, error) {
	var rv LookupAppSecWafModeResult
	err := ctx.Invoke("akamai:index/getAppSecWafMode:getAppSecWafMode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecWafMode.
type LookupAppSecWafModeArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecWafMode.
type LookupAppSecWafModeResult struct {
	ConfigId int `pulumi:"configId"`
	// The current rule set version and the ISO 8601 date the rule set version was introduced; this date acts like a version number.
	CurrentRuleset string `pulumi:"currentRuleset"`
	// The ISO 8601 time stamp when the evaluation is expiring. This value only appears when `eval` is set to "enabled".
	EvalExpirationDate string `pulumi:"evalExpirationDate"`
	// The evaluation rule set version and the ISO 8601 date the evaluation starts.
	EvalRuleset string `pulumi:"evalRuleset"`
	// Whether the evaluation mode is enabled or disabled."
	EvalStatus string `pulumi:"evalStatus"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of the mode information.
	Json string `pulumi:"json"`
	// The security policy mode, either `KRS` (update manually) or `AAG` (update automatically), For Adaptive Security Engine (ASE) __BETA__, use `ASE_AUTO` for automatic updates or `ASE_MANUAL` to manually get current rules. Please contact your Akamai representative to learn more about ASE.
	Mode string `pulumi:"mode"`
	// A tabular display of the mode information.
	OutputText       string `pulumi:"outputText"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}
