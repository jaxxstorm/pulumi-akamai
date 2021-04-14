// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAppSecCustomRuleActions` data source to retrieve information about the actions defined for the custom rules associated with a specific security configuration, version and security policy.
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
// 		opt0 := "Akamai Tools"
// 		configuration, err := akamai.GetAppSecConfiguration(ctx, &akamai.GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		customRuleActionsAppSecCustomRuleActions, err := akamai.GetAppSecCustomRuleActions(ctx, &akamai.GetAppSecCustomRuleActionsArgs{
// 			ConfigId: configuration.ConfigId,
// 			Version:  configuration.LatestVersion,
// 			PolicyId: "crAP_75829",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("customRuleActions", customRuleActionsAppSecCustomRuleActions.OutputText)
// 		return nil
// 	})
// }
// ```
func GetAppSecCustomRuleActions(ctx *pulumi.Context, args *GetAppSecCustomRuleActionsArgs, opts ...pulumi.InvokeOption) (*GetAppSecCustomRuleActionsResult, error) {
	var rv GetAppSecCustomRuleActionsResult
	err := ctx.Invoke("akamai:index/getAppSecCustomRuleActions:getAppSecCustomRuleActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecCustomRuleActions.
type GetAppSecCustomRuleActionsArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the security policy to use
	PolicyId string `pulumi:"policyId"`
	// The version number of the security configuration to use.
	Version int `pulumi:"version"`
}

// A collection of values returned by getAppSecCustomRuleActions.
type GetAppSecCustomRuleActionsResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A tabular display showing the ID, name, and action of all custom rules associated with the specified security configuration, version and security policy.
	OutputText string `pulumi:"outputText"`
	PolicyId   string `pulumi:"policyId"`
	Version    int    `pulumi:"version"`
}
