// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAppSecFailoverHostnames` data source to retrieve a list of the failover hostnames in a configuration. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getfailoverhostnames).
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
// 		failoverHostnamesAppSecFailoverHostnames, err := akamai.GetAppSecFailoverHostnames(ctx, &akamai.GetAppSecFailoverHostnamesArgs{
// 			ConfigId: configuration.ConfigId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("failoverHostnames", failoverHostnamesAppSecFailoverHostnames.Hostnames)
// 		ctx.Export("failoverHostnamesOutput", failoverHostnamesAppSecFailoverHostnames.OutputText)
// 		ctx.Export("failoverHostnamesJson", failoverHostnamesAppSecFailoverHostnames.Json)
// 		return nil
// 	})
// }
// ```
func GetAppSecFailoverHostnames(ctx *pulumi.Context, args *GetAppSecFailoverHostnamesArgs, opts ...pulumi.InvokeOption) (*GetAppSecFailoverHostnamesResult, error) {
	var rv GetAppSecFailoverHostnamesResult
	err := ctx.Invoke("akamai:index/getAppSecFailoverHostnames:getAppSecFailoverHostnames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecFailoverHostnames.
type GetAppSecFailoverHostnamesArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
}

// A collection of values returned by getAppSecFailoverHostnames.
type GetAppSecFailoverHostnamesResult struct {
	ConfigId int `pulumi:"configId"`
	// A list of the failover hostnames.
	Hostnames []string `pulumi:"hostnames"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of the failover hostnames.
	Json string `pulumi:"json"`
	// A tabular display showing the failover hostnames.
	OutputText string `pulumi:"outputText"`
}
