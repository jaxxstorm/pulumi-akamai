// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecPenaltyBox` data source to retrieve the penalty box settings for a specified security policy.
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
// 		penaltyBox, err := akamai.LookupAppSecPenaltyBox(ctx, &akamai.LookupAppSecPenaltyBoxArgs{
// 			ConfigId:         configuration.ConfigId,
// 			Version:          configuration.LatestVersion,
// 			SecurityPolicyId: _var.Security_policy_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("penaltyBoxAction", penaltyBox.Action)
// 		ctx.Export("penaltyBoxEnabled", penaltyBox.Enabled)
// 		ctx.Export("penaltyBoxText", penaltyBox.OutputText)
// 		return nil
// 	})
// }
// ```
func LookupAppSecPenaltyBox(ctx *pulumi.Context, args *LookupAppSecPenaltyBoxArgs, opts ...pulumi.InvokeOption) (*LookupAppSecPenaltyBoxResult, error) {
	var rv LookupAppSecPenaltyBoxResult
	err := ctx.Invoke("akamai:index/getAppSecPenaltyBox:getAppSecPenaltyBox", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecPenaltyBox.
type LookupAppSecPenaltyBoxArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	// The version number of the security configuration to use.
	Version int `pulumi:"version"`
}

// A collection of values returned by getAppSecPenaltyBox.
type LookupAppSecPenaltyBoxResult struct {
	// The action for the penalty box: `alert`, `deny`, or `none`.
	Action   string `pulumi:"action"`
	ConfigId int    `pulumi:"configId"`
	// Either `true` or `false`, indicating whether penalty box protection is enabled.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A tabular display of the `action` and `enabled` information.
	OutputText       string `pulumi:"outputText"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	Version          int    `pulumi:"version"`
}
