// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAppSecMatchTargets` data source to retrieve information about the match targets associated with a given configuration version, or about a specific match target.
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
// 		matchTargetsAppSecMatchTargets, err := akamai.GetAppSecMatchTargets(ctx, &akamai.GetAppSecMatchTargetsArgs{
// 			ConfigId: configuration.ConfigId,
// 			Version:  configuration.LatestVersion,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("matchTargets", matchTargetsAppSecMatchTargets.OutputText)
// 		opt1 := _var.Match_target_id
// 		matchTarget, err := akamai.GetAppSecMatchTargets(ctx, &akamai.GetAppSecMatchTargetsArgs{
// 			ConfigId:      configuration.ConfigId,
// 			Version:       configuration.LatestVersion,
// 			MatchTargetId: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("matchTargetOutput", matchTarget.OutputText)
// 		return nil
// 	})
// }
// ```
func GetAppSecMatchTargets(ctx *pulumi.Context, args *GetAppSecMatchTargetsArgs, opts ...pulumi.InvokeOption) (*GetAppSecMatchTargetsResult, error) {
	var rv GetAppSecMatchTargetsResult
	err := ctx.Invoke("akamai:index/getAppSecMatchTargets:getAppSecMatchTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecMatchTargets.
type GetAppSecMatchTargetsArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the match target to use. If not supplied, information about all match targets is returned.
	MatchTargetId *int `pulumi:"matchTargetId"`
	// The version number of the security configuration to use.
	Version int `pulumi:"version"`
}

// A collection of values returned by getAppSecMatchTargets.
type GetAppSecMatchTargetsResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of the match target information.
	Json          string `pulumi:"json"`
	MatchTargetId *int   `pulumi:"matchTargetId"`
	// A tabular display showing the ID and Policy ID of all match targets associated with the specified security configuration and version, or of the specific match target if `matchTargetId` was supplied.
	OutputText string `pulumi:"outputText"`
	Version    int    `pulumi:"version"`
}
