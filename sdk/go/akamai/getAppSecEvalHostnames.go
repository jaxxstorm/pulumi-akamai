// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecEvalHostnames` data source to retrieve the evaluation hostnames for a configuration version. Evaluation mode for hostnames is only available for Web Application Protector. Run hostnames in evaluation mode to see how your configuration settings protect traffic for that hostname before adding a hostname directly to a live configuration. An evaluation period lasts four weeks unless you stop the evaluation. Once you begin, the hostnames you evaluate start responding to traffic as if they are your current hostnames. However, instead of taking an action the evaluation hostnames log which action they would have taken if they were your actively-protected hostnames and not a test.
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
// 		evalHostnamesAppSecEvalHostnames, err := akamai.LookupAppSecEvalHostnames(ctx, &akamai.LookupAppSecEvalHostnamesArgs{
// 			ConfigId: configuration.ConfigId,
// 			Version:  configuration.LatestVersion,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("evalHostnames", evalHostnamesAppSecEvalHostnames.Hostnames)
// 		ctx.Export("evalHostnamesOutput", evalHostnamesAppSecEvalHostnames.OutputText)
// 		ctx.Export("evalHostnamesJson", evalHostnamesAppSecEvalHostnames.Json)
// 		return nil
// 	})
// }
// ```
func LookupAppSecEvalHostnames(ctx *pulumi.Context, args *LookupAppSecEvalHostnamesArgs, opts ...pulumi.InvokeOption) (*LookupAppSecEvalHostnamesResult, error) {
	var rv LookupAppSecEvalHostnamesResult
	err := ctx.Invoke("akamai:index/getAppSecEvalHostnames:getAppSecEvalHostnames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecEvalHostnames.
type LookupAppSecEvalHostnamesArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The version number of the security configuration to use.
	Version int `pulumi:"version"`
}

// A collection of values returned by getAppSecEvalHostnames.
type LookupAppSecEvalHostnamesResult struct {
	ConfigId int `pulumi:"configId"`
	// A list of the evaluation hostnames.
	Hostnames []string `pulumi:"hostnames"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list of the evaluation hostnames.
	Json string `pulumi:"json"`
	// A tabular display showing the evaluation hostnames.
	OutputText string `pulumi:"outputText"`
	Version    int    `pulumi:"version"`
}
