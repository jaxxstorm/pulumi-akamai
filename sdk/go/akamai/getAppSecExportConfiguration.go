// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use the `getAppSecExportConfiguration` data source to retrieve comprehensive details about a security configuration version, including rate and security policies, rules, hostnames, and other settings. You can retrieve the entire set of information in JSON format, or a subset of the information in tabular format.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		export, err := akamai.GetAppSecExportConfiguration(ctx, &akamai.GetAppSecExportConfigurationArgs{
// 			ConfigId: configuration.ConfigId,
// 			Version:  configuration.LatestVersion,
// 			Searches: []string{
// 				"securityPolicies",
// 				"selectedHosts",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("json", export.Json)
// 		ctx.Export("text", export.OutputText)
// 		return nil
// 	})
// }
// ```
func GetAppSecExportConfiguration(ctx *pulumi.Context, args *GetAppSecExportConfigurationArgs, opts ...pulumi.InvokeOption) (*GetAppSecExportConfigurationResult, error) {
	var rv GetAppSecExportConfigurationResult
	err := ctx.Invoke("akamai:index/getAppSecExportConfiguration:getAppSecExportConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecExportConfiguration.
type GetAppSecExportConfigurationArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// A bracket-delimited list of quoted strings specifying the types of information to be retrieved and made available for display in the `outputText` format. The following types are available:
	// * customRules
	// * matchTargets
	// * ratePolicies
	// * reputationProfiles
	// * rulesets
	// * securityPolicies
	// * selectableHosts
	// * selectedHosts
	Searches []string `pulumi:"searches"`
	// The version number of the security configuration to use.
	Version int `pulumi:"version"`
}

// A collection of values returned by getAppSecExportConfiguration.
type GetAppSecExportConfigurationResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The complete set of information about the specified security configuration version, in JSON format. This includes the types available for the `search` parameter, plus several additional fields such as createDate and createdBy.
	Json string `pulumi:"json"`
	// A tabular display showing the types of data specified in the `search` parameter. Included only if the `search` parameter specifies at least one type.
	OutputText string   `pulumi:"outputText"`
	Searches   []string `pulumi:"searches"`
	Version    int      `pulumi:"version"`
}
