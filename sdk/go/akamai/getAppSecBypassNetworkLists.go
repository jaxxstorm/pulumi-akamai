// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecByPassNetworkList` data source to retrieve information about which network
// lists are used in the bypass network lists settings.  The information available is described
// [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getbypassnetworklistsforawapconfigversion).
// Note: this data source is only applicable to WAP (Web Application Protector) configurations.
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
// 		bypassNetworkLists, err := akamai.GetAppSecBypassNetworkLists(ctx, &akamai.GetAppSecBypassNetworkListsArgs{
// 			ConfigId: configuration.ConfigId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("bypassNetworkListsOutput", bypassNetworkLists.OutputText)
// 		ctx.Export("bypassNetworkListsJson", bypassNetworkLists.Json)
// 		ctx.Export("bypassNetworkListsIdList", bypassNetworkLists.BypassNetworkLists)
// 		return nil
// 	})
// }
// ```
func GetAppSecBypassNetworkLists(ctx *pulumi.Context, args *GetAppSecBypassNetworkListsArgs, opts ...pulumi.InvokeOption) (*GetAppSecBypassNetworkListsResult, error) {
	var rv GetAppSecBypassNetworkListsResult
	err := ctx.Invoke("akamai:index/getAppSecBypassNetworkLists:getAppSecBypassNetworkLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecBypassNetworkLists.
type GetAppSecBypassNetworkListsArgs struct {
	// The configuration ID to use.
	ConfigId int `pulumi:"configId"`
}

// A collection of values returned by getAppSecBypassNetworkLists.
type GetAppSecBypassNetworkListsResult struct {
	// A list of strings containing the network list IDs.
	BypassNetworkLists []string `pulumi:"bypassNetworkLists"`
	ConfigId           int      `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of information about the bypass network lists.
	Json string `pulumi:"json"`
	// A tabular display showing the bypass network list information.
	OutputText string `pulumi:"outputText"`
}