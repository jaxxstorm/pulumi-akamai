// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration; security policy
//
// Returns pragma header settings information. This HTTP header provides information about such things as: the edge routers used in a transaction; the Akamai IP addresses involved; information about whether a request was cached or not; and so on. By default, pragma headers are removed from all responses.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/pragma-header](https://techdocs.akamai.com/application-security/reference/get-advanced-settings-pragma-header)
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
//				Name: pulumi.StringRef("Documentation"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			pragmaHeader, err := akamai.LookupAppSecAdvancedSettingsPragmaHeader(ctx, &GetAppSecAdvancedSettingsPragmaHeaderArgs{
//				ConfigId: configuration.ConfigId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("advancedSettingsPragmaHeaderOutput", pragmaHeader.OutputText)
//			ctx.Export("advancedSettingsPragmaHeaderJson", pragmaHeader.Json)
//			policyPragmaHeader, err := akamai.LookupAppSecAdvancedSettingsPragmaHeader(ctx, &GetAppSecAdvancedSettingsPragmaHeaderArgs{
//				ConfigId:         configuration.ConfigId,
//				SecurityPolicyId: pulumi.StringRef("gms1_134637"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("advancedSettingsPolicyPragmaHeaderOutput", policyPragmaHeader.OutputText)
//			ctx.Export("advancedSettingsPolicyPragmaHeaderJson", policyPragmaHeader.Json)
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `json`. JSON-formatted list of information about the pragma header settings.
// - `outputText`. Tabular report showing the pragma header settings.
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
	// . Unique identifier of the security configuration associated with the pragma header settings.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the pragma header settings. If not included, information is returned for all your security policies.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecAdvancedSettingsPragmaHeader.
type LookupAppSecAdvancedSettingsPragmaHeaderResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Json             string  `pulumi:"json"`
	OutputText       string  `pulumi:"outputText"`
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

func LookupAppSecAdvancedSettingsPragmaHeaderOutput(ctx *pulumi.Context, args LookupAppSecAdvancedSettingsPragmaHeaderOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecAdvancedSettingsPragmaHeaderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecAdvancedSettingsPragmaHeaderResult, error) {
			args := v.(LookupAppSecAdvancedSettingsPragmaHeaderArgs)
			r, err := LookupAppSecAdvancedSettingsPragmaHeader(ctx, &args, opts...)
			var s LookupAppSecAdvancedSettingsPragmaHeaderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppSecAdvancedSettingsPragmaHeaderResultOutput)
}

// A collection of arguments for invoking getAppSecAdvancedSettingsPragmaHeader.
type LookupAppSecAdvancedSettingsPragmaHeaderOutputArgs struct {
	// . Unique identifier of the security configuration associated with the pragma header settings.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the pragma header settings. If not included, information is returned for all your security policies.
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

func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPragmaHeaderResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPragmaHeaderResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o LookupAppSecAdvancedSettingsPragmaHeaderResultOutput) SecurityPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPragmaHeaderResult) *string { return v.SecurityPolicyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecAdvancedSettingsPragmaHeaderResultOutput{})
}
