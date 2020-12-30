// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use the `getAppSecConfigurationVersion` data source to retrieve information about the versions of a security configuration, or about a specific version.
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
// 		specificConfiguration, err := akamai.GetAppSecConfiguration(ctx, &akamai.GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		versions, err := akamai.GetAppSecConfigurationVersion(ctx, &akamai.GetAppSecConfigurationVersionArgs{
// 			ConfigId: specificConfiguration.ConfigId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("versionsOutputText", versions.OutputText)
// 		ctx.Export("versionsLatest", versions.LatestVersion)
// 		opt1 := 42
// 		specificVersion, err := akamai.GetAppSecConfigurationVersion(ctx, &akamai.GetAppSecConfigurationVersionArgs{
// 			ConfigId: specificConfiguration.ConfigId,
// 			Version:  &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("specificVersionVersion", specificVersion.Version)
// 		ctx.Export("specificVersionStaging", specificVersion.StagingStatus)
// 		ctx.Export("specificVersionProduction", specificVersion.ProductionStatus)
// 		return nil
// 	})
// }
// ```
func GetAppSecConfigurationVersion(ctx *pulumi.Context, args *GetAppSecConfigurationVersionArgs, opts ...pulumi.InvokeOption) (*GetAppSecConfigurationVersionResult, error) {
	var rv GetAppSecConfigurationVersionResult
	err := ctx.Invoke("akamai:index/getAppSecConfigurationVersion:getAppSecConfigurationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecConfigurationVersion.
type GetAppSecConfigurationVersionArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The version number of the security configuration to use. If not supplied, information about all versions of the specified security configuration is returned.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getAppSecConfigurationVersion.
type GetAppSecConfigurationVersionResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The last version of the security configuration created.
	LatestVersion int `pulumi:"latestVersion"`
	// A tabular display showing the following information about all versions of the security configuration: version number, staging status, and production status.
	OutputText string `pulumi:"outputText"`
	// The status of the specified version in production: "Active", "Inactive", or "Deactivated". Returned only if `version` was specified.
	ProductionStatus string `pulumi:"productionStatus"`
	// The status of the specified version in staging: "Active", "Inactive", or "Deactivated". Returned only if `version` was specified.
	StagingStatus string `pulumi:"stagingStatus"`
	Version       *int   `pulumi:"version"`
}