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
// Returns the slow POST protection settings for the specified security configuration and policy. Slow POST protections help defend a site against attacks that try to tie up the site by using extremely slow requests and responses: the idea is to keep the site occupied waiting for these requests and responses to finish instead of being occupied with new (and legitimate) transactions.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/slow-post](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getslowpostprotectionsettings)
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
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: pulumi.StringRef("Documentation"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		slowPost, err := akamai.LookupAppSecSlowPost(ctx, &GetAppSecSlowPostArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: "gms1_134637",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("slowPostOutputText", slowPost.OutputText)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `outputText`. Tabular report including the following:
//   - **ACTION**. Action taken any time slow POST protection is triggered. Valid values are:
//     - **alert**. Record the event.
//     - **abort**. Block the request.
//   - **SLOW_RATE_THRESHOLD RATE**. Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
//   - **SLOW_RATE_THRESHOLD PERIOD**. Amount of time (in seconds) that the server should allow a request before marking the request as being too slow
//   - **DURATION_THRESHOLD TIMEOUT**. Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in order to avoid triggering the specified action.
func LookupAppSecSlowPost(ctx *pulumi.Context, args *LookupAppSecSlowPostArgs, opts ...pulumi.InvokeOption) (*LookupAppSecSlowPostResult, error) {
	var rv LookupAppSecSlowPostResult
	err := ctx.Invoke("akamai:index/getAppSecSlowPost:getAppSecSlowPost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecSlowPost.
type LookupAppSecSlowPostArgs struct {
	// . Unique identifier of the security configuration associated with the slow POST settings.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the slow POST settings.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecSlowPost.
type LookupAppSecSlowPostResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	Json             string `pulumi:"json"`
	OutputText       string `pulumi:"outputText"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

func LookupAppSecSlowPostOutput(ctx *pulumi.Context, args LookupAppSecSlowPostOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecSlowPostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecSlowPostResult, error) {
			args := v.(LookupAppSecSlowPostArgs)
			r, err := LookupAppSecSlowPost(ctx, &args, opts...)
			return *r, err
		}).(LookupAppSecSlowPostResultOutput)
}

// A collection of arguments for invoking getAppSecSlowPost.
type LookupAppSecSlowPostOutputArgs struct {
	// . Unique identifier of the security configuration associated with the slow POST settings.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the slow POST settings.
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (LookupAppSecSlowPostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecSlowPostArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecSlowPost.
type LookupAppSecSlowPostResultOutput struct{ *pulumi.OutputState }

func (LookupAppSecSlowPostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecSlowPostResult)(nil)).Elem()
}

func (o LookupAppSecSlowPostResultOutput) ToLookupAppSecSlowPostResultOutput() LookupAppSecSlowPostResultOutput {
	return o
}

func (o LookupAppSecSlowPostResultOutput) ToLookupAppSecSlowPostResultOutputWithContext(ctx context.Context) LookupAppSecSlowPostResultOutput {
	return o
}

func (o LookupAppSecSlowPostResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecSlowPostResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAppSecSlowPostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecSlowPostResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppSecSlowPostResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecSlowPostResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o LookupAppSecSlowPostResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecSlowPostResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o LookupAppSecSlowPostResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecSlowPostResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecSlowPostResultOutput{})
}
