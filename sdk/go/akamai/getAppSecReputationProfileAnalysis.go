// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecReputationProfileAnalysis` data source to retrieve information about the current reputation analysis settings. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getreputationanalysis).
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
// 		reputationAnalysis, err := akamai.LookupAppSecReputationProfileAnalysis(ctx, &akamai.LookupAppSecReputationProfileAnalysisArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: _var.Security_policy_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("reputationAnalysisText", reputationAnalysis.OutputText)
// 		ctx.Export("reputationAnalysisJson", reputationAnalysis.Json)
// 		return nil
// 	})
// }
// ```
func LookupAppSecReputationProfileAnalysis(ctx *pulumi.Context, args *LookupAppSecReputationProfileAnalysisArgs, opts ...pulumi.InvokeOption) (*LookupAppSecReputationProfileAnalysisResult, error) {
	var rv LookupAppSecReputationProfileAnalysisResult
	err := ctx.Invoke("akamai:index/getAppSecReputationProfileAnalysis:getAppSecReputationProfileAnalysis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecReputationProfileAnalysis.
type LookupAppSecReputationProfileAnalysisArgs struct {
	// The configuration ID to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecReputationProfileAnalysis.
type LookupAppSecReputationProfileAnalysisResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of the reputation analysis settings.
	Json string `pulumi:"json"`
	// A tabular display showing the reputation analysis settings.
	OutputText       string `pulumi:"outputText"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}
