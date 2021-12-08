// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy
//
// Returns threat intelligence settings for a security policy Note that this data source is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information on ASE, please contact your Akamai representative.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/threat-intel](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getthreatintelligence)l
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
// 		opt0 := "Documentation"
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		threatIntelAppSecThreatIntel, err := akamai.LookupAppSecThreatIntel(ctx, &GetAppSecThreatIntelArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: "gms1_134637",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("threatIntel", threatIntelAppSecThreatIntel.ThreatIntel)
// 		ctx.Export("json", threatIntelAppSecThreatIntel.Json)
// 		ctx.Export("outputText", threatIntelAppSecThreatIntel.OutputText)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `threatIntel`. Reports the threat Intelligence setting, either **on** or **off**.
// - `json`. JSON-formatted threat intelligence report
// - `outputText`. Tabular report of the threat intelligence information.
func LookupAppSecThreatIntel(ctx *pulumi.Context, args *LookupAppSecThreatIntelArgs, opts ...pulumi.InvokeOption) (*LookupAppSecThreatIntelResult, error) {
	var rv LookupAppSecThreatIntelResult
	err := ctx.Invoke("akamai:index/getAppSecThreatIntel:getAppSecThreatIntel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecThreatIntel.
type LookupAppSecThreatIntelArgs struct {
	// . Unique identifier of the security configuration associated with the threat intelligence settings.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the threat intelligence settings.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecThreatIntel.
type LookupAppSecThreatIntelResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	Json             string `pulumi:"json"`
	OutputText       string `pulumi:"outputText"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	ThreatIntel      string `pulumi:"threatIntel"`
}

func LookupAppSecThreatIntelOutput(ctx *pulumi.Context, args LookupAppSecThreatIntelOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecThreatIntelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecThreatIntelResult, error) {
			args := v.(LookupAppSecThreatIntelArgs)
			r, err := LookupAppSecThreatIntel(ctx, &args, opts...)
			return *r, err
		}).(LookupAppSecThreatIntelResultOutput)
}

// A collection of arguments for invoking getAppSecThreatIntel.
type LookupAppSecThreatIntelOutputArgs struct {
	// . Unique identifier of the security configuration associated with the threat intelligence settings.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the threat intelligence settings.
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (LookupAppSecThreatIntelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecThreatIntelArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecThreatIntel.
type LookupAppSecThreatIntelResultOutput struct{ *pulumi.OutputState }

func (LookupAppSecThreatIntelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecThreatIntelResult)(nil)).Elem()
}

func (o LookupAppSecThreatIntelResultOutput) ToLookupAppSecThreatIntelResultOutput() LookupAppSecThreatIntelResultOutput {
	return o
}

func (o LookupAppSecThreatIntelResultOutput) ToLookupAppSecThreatIntelResultOutputWithContext(ctx context.Context) LookupAppSecThreatIntelResultOutput {
	return o
}

func (o LookupAppSecThreatIntelResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecThreatIntelResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAppSecThreatIntelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecThreatIntelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppSecThreatIntelResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecThreatIntelResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o LookupAppSecThreatIntelResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecThreatIntelResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o LookupAppSecThreatIntelResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecThreatIntelResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func (o LookupAppSecThreatIntelResultOutput) ThreatIntel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecThreatIntelResult) string { return v.ThreatIntel }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecThreatIntelResultOutput{})
}
