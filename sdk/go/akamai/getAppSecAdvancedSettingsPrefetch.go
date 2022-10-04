// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration
//
// Returns information about your prefetch request settings. By default, Web Application Firewall inspects only external requests — requests originating outside of your firewall or Akamai's edge servers. When prefetch is enabled, requests between your origin servers and Akamai's edge servers can also be inspected by the firewall.
//
// **Related** **API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/prefetch](https://techdocs.akamai.com/application-security/reference/get-advanced-settings-prefetch)
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
//			prefetch, err := akamai.LookupAppSecAdvancedSettingsPrefetch(ctx, &GetAppSecAdvancedSettingsPrefetchArgs{
//				ConfigId: configuration.ConfigId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("advancedSettingsPrefetchOutput", prefetch.OutputText)
//			ctx.Export("advancedSettingsPrefetchJson", prefetch.Json)
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `json`. JSON-formatted list of information about the prefetch request settings.
// - `outputText`. Tabular report showing the prefetch request settings.
func LookupAppSecAdvancedSettingsPrefetch(ctx *pulumi.Context, args *LookupAppSecAdvancedSettingsPrefetchArgs, opts ...pulumi.InvokeOption) (*LookupAppSecAdvancedSettingsPrefetchResult, error) {
	var rv LookupAppSecAdvancedSettingsPrefetchResult
	err := ctx.Invoke("akamai:index/getAppSecAdvancedSettingsPrefetch:getAppSecAdvancedSettingsPrefetch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecAdvancedSettingsPrefetch.
type LookupAppSecAdvancedSettingsPrefetchArgs struct {
	// . Unique identifier of the security configuration associated with the prefetch settings.
	ConfigId int `pulumi:"configId"`
}

// A collection of values returned by getAppSecAdvancedSettingsPrefetch.
type LookupAppSecAdvancedSettingsPrefetchResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Json       string `pulumi:"json"`
	OutputText string `pulumi:"outputText"`
}

func LookupAppSecAdvancedSettingsPrefetchOutput(ctx *pulumi.Context, args LookupAppSecAdvancedSettingsPrefetchOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecAdvancedSettingsPrefetchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecAdvancedSettingsPrefetchResult, error) {
			args := v.(LookupAppSecAdvancedSettingsPrefetchArgs)
			r, err := LookupAppSecAdvancedSettingsPrefetch(ctx, &args, opts...)
			var s LookupAppSecAdvancedSettingsPrefetchResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppSecAdvancedSettingsPrefetchResultOutput)
}

// A collection of arguments for invoking getAppSecAdvancedSettingsPrefetch.
type LookupAppSecAdvancedSettingsPrefetchOutputArgs struct {
	// . Unique identifier of the security configuration associated with the prefetch settings.
	ConfigId pulumi.IntInput `pulumi:"configId"`
}

func (LookupAppSecAdvancedSettingsPrefetchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecAdvancedSettingsPrefetchArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecAdvancedSettingsPrefetch.
type LookupAppSecAdvancedSettingsPrefetchResultOutput struct{ *pulumi.OutputState }

func (LookupAppSecAdvancedSettingsPrefetchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecAdvancedSettingsPrefetchResult)(nil)).Elem()
}

func (o LookupAppSecAdvancedSettingsPrefetchResultOutput) ToLookupAppSecAdvancedSettingsPrefetchResultOutput() LookupAppSecAdvancedSettingsPrefetchResultOutput {
	return o
}

func (o LookupAppSecAdvancedSettingsPrefetchResultOutput) ToLookupAppSecAdvancedSettingsPrefetchResultOutputWithContext(ctx context.Context) LookupAppSecAdvancedSettingsPrefetchResultOutput {
	return o
}

func (o LookupAppSecAdvancedSettingsPrefetchResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPrefetchResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAppSecAdvancedSettingsPrefetchResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPrefetchResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppSecAdvancedSettingsPrefetchResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPrefetchResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o LookupAppSecAdvancedSettingsPrefetchResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecAdvancedSettingsPrefetchResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecAdvancedSettingsPrefetchResultOutput{})
}
