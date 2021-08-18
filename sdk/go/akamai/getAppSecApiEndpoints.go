// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAppSecApiEndpoints` data source to retrieve information about the API Endpoints associated with a security policy or configuration. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapiendpoints).
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
// 		opt0 := "TestEndpoint"
// 		_, err := akamai.GetAppSecApiEndpoints(ctx, &akamai.GetAppSecApiEndpointsArgs{
// 			ApiName:  &opt0,
// 			ConfigId: 43253,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAppSecApiEndpoints(ctx *pulumi.Context, args *GetAppSecApiEndpointsArgs, opts ...pulumi.InvokeOption) (*GetAppSecApiEndpointsResult, error) {
	var rv GetAppSecApiEndpointsResult
	err := ctx.Invoke("akamai:index/getAppSecApiEndpoints:getAppSecApiEndpoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecApiEndpoints.
type GetAppSecApiEndpointsArgs struct {
	// The name of a specific endpoint.
	ApiName *string `pulumi:"apiName"`
	// The configuration ID.
	ConfigId int `pulumi:"configId"`
	// The ID of the security policy to use.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecApiEndpoints.
type GetAppSecApiEndpointsResult struct {
	ApiName  *string `pulumi:"apiName"`
	ConfigId int     `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of IDs of the API endpoints.
	IdLists []int `pulumi:"idLists"`
	// A JSON-formatted list of information about the API endpoints.
	Json string `pulumi:"json"`
	// A tabular display showing the ID and name of the API endpoints.
	OutputText       string  `pulumi:"outputText"`
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}
