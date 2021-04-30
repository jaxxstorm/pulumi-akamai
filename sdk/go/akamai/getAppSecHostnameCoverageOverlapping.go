// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAppSecHostnameCoverageOverlapping` data source to retrieve information about the configuration versions that contain a hostname also included in the current configuration version. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoverageoverlapping).
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
// 		_, err := akamai.GetAppSecHostnameCoverageOverlapping(ctx, &akamai.GetAppSecHostnameCoverageOverlappingArgs{
// 			ConfigId: 43253,
// 			Hostname: "example.com",
// 			Version:  7,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAppSecHostnameCoverageOverlapping(ctx *pulumi.Context, args *GetAppSecHostnameCoverageOverlappingArgs, opts ...pulumi.InvokeOption) (*GetAppSecHostnameCoverageOverlappingResult, error) {
	var rv GetAppSecHostnameCoverageOverlappingResult
	err := ctx.Invoke("akamai:index/getAppSecHostnameCoverageOverlapping:getAppSecHostnameCoverageOverlapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecHostnameCoverageOverlapping.
type GetAppSecHostnameCoverageOverlappingArgs struct {
	// The configuration ID.
	ConfigId int `pulumi:"configId"`
	// The hostname for which to retrieve information.
	Hostname string `pulumi:"hostname"`
	// The version number of the configuration.
	Version int `pulumi:"version"`
}

// A collection of values returned by getAppSecHostnameCoverageOverlapping.
type GetAppSecHostnameCoverageOverlappingResult struct {
	ConfigId int    `pulumi:"configId"`
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of the overlap information.
	Json string `pulumi:"json"`
	// A tabular display of the overlap information.
	OutputText string `pulumi:"outputText"`
	Version    int    `pulumi:"version"`
}
