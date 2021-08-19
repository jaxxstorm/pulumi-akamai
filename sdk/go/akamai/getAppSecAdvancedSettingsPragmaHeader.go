// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecAdvancedSettingsPragmaHeader` data source to retrieve pragma header settings for a configuration or a security policy. Additional information is available [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getpragmaheaderconfiguration).
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
// 		pragmaHeader, err := akamai.LookupAppSecAdvancedSettingsPragmaHeader(ctx, &akamai.LookupAppSecAdvancedSettingsPragmaHeaderArgs{
// 			ConfigId: configuration.ConfigId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("advancedSettingsPragmaHeaderOutput", pragmaHeader.OutputText)
// 		ctx.Export("advancedSettingsPragmaHeaderJson", pragmaHeader.Json)
// 		opt1 := _var.Security_policy_id
// 		policyPragmaHeader, err := akamai.LookupAppSecAdvancedSettingsPragmaHeader(ctx, &akamai.LookupAppSecAdvancedSettingsPragmaHeaderArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("advancedSettingsPolicyPragmaHeaderOutput", policyPragmaHeader.OutputText)
// 		ctx.Export("advancedSettingsPolicyPragmaHeaderJson", policyPragmaHeader.Json)
// 		return nil
// 	})
// }
// ```
func LookupAppSecAdvancedSettingsPragmaHeader(ctx *pulumi.Context, args *LookupAppSecAdvancedSettingsPragmaHeaderArgs, opts ...pulumi.InvokeOption) (*LookupAppSecAdvancedSettingsPragmaHeaderResult, error) {
	var rv LookupAppSecAdvancedSettingsPragmaHeaderResult
	err := ctx.Invoke("akamai:index/getAppSecAdvancedSettingsPragmaHeader:getAppSecAdvancedSettingsPragmaHeader", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecAdvancedSettingsPragmaHeader.
type LookupAppSecAdvancedSettingsPragmaHeaderArgs struct {
	// The configuration ID.
	ConfigId int `pulumi:"configId"`
	// The ID of the security policy to use.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecAdvancedSettingsPragmaHeader.
type LookupAppSecAdvancedSettingsPragmaHeaderResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putpragmaheaderpolicy)) list of information about the pragma header settings.
	Json string `pulumi:"json"`
	// A tabular display showing the pragma header settings.
	OutputText       string  `pulumi:"outputText"`
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}