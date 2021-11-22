// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

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
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		pragmaHeader, err := akamai.LookupAppSecAdvancedSettingsPragmaHeader(ctx, &GetAppSecAdvancedSettingsPragmaHeaderArgs{
// 			ConfigId: configuration.ConfigId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("advancedSettingsPragmaHeaderOutput", pragmaHeader.OutputText)
// 		ctx.Export("advancedSettingsPragmaHeaderJson", pragmaHeader.Json)
// 		opt1 := _var.Security_policy_id
// 		policyPragmaHeader, err := akamai.LookupAppSecAdvancedSettingsPragmaHeader(ctx, &GetAppSecAdvancedSettingsPragmaHeaderArgs{
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

func LookupAppSecAdvancedSettingsPragmaHeaderOutput(ctx *pulumi.Context, args LookupAppSecAdvancedSettingsPragmaHeaderOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecAdvancedSettingsPragmaHeaderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecAdvancedSettingsPragmaHeaderResult, error) {
			args := v.(LookupAppSecAdvancedSettingsPragmaHeaderArgs)
			r, err := LookupAppSecAdvancedSettingsPragmaHeader(ctx, &args, opts...)
			return *r, err
		}).(LookupAppSecAdvancedSettingsPragmaHeaderResultOutput)
}

// A collection of arguments for invoking getAppSecAdvancedSettingsPragmaHeader.
type LookupAppSecAdvancedSettingsPragmaHeaderOutputArgs struct {
	// The configuration ID.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringPtrInput `pulumi:"securityPolicyId"`
}

func (LookupAppSecAdvancedSettingsPragmaHeaderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecAdvancedSettingsPragmaHeaderArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecAdvancedSettingsPragmaHeader.
type LookupAppSecAdvancedSettingsPragmaHeaderResultOutput struct{ *pulumi.OutputState }

func (LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecAdvancedSettingsPragmaHeaderResult)(nil)).Elem()
}

func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) ToLookupAppSecAdvancedSettingsPragmaHeaderResultOutput() LookupAppSecAdvancedSettingsPragmaHeaderResultOutput {
	return o
}

func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) ToLookupAppSecAdvancedSettingsPragmaHeaderResultOutputWithContext(ctx context.Context) LookupAppSecAdvancedSettingsPragmaHeaderResultOutput {
	return o
}

func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPragmaHeaderResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPragmaHeaderResult) string { return v.Id }).(pulumi.StringOutput)
}

// A JSON-formatted ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putpragmaheaderpolicy)) list of information about the pragma header settings.
func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPragmaHeaderResult) string { return v.Json }).(pulumi.StringOutput)
}

// A tabular display showing the pragma header settings.
func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPragmaHeaderResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) SecurityPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPragmaHeaderResult) *string { return v.SecurityPolicyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecAdvancedSettingsPragmaHeaderResultOutput{})
}
