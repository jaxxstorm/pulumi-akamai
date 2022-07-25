// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy
//
// Returns hostnames currently protected or being evaluated by a configuration and security policy.
// This resource is available only to organizations running Web Application Protector (WAP).
// Note that the WAP selected hostnames feature is currently in beta.
// Please contact your Akamai representative for more information.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
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
// 		wapSelectedHostnames, err := akamai.LookupAppSecWapSelectedHostnames(ctx, &GetAppSecWapSelectedHostnamesArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: "gms1_134637",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("protectedHostnames", wapSelectedHostnames.ProtectedHosts)
// 		ctx.Export("evaluatedHostnames", wapSelectedHostnames.EvaluatedHosts)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned and how that returned information is formatted:
//
// - `protectedHostnames`. List of hostnames currently protected under the security configuration and security policy.
// - `evaluatedHostnames`. List of hostnames currently being evaluated under the security configuration and security policy.
// - `hostnamesJson`. JSON-formatted report of the protected and evaluated hostnames.
// - `outputText`. Tabular reports of the protected and evaluated hostnames.
func LookupAppSecWapSelectedHostnames(ctx *pulumi.Context, args *LookupAppSecWapSelectedHostnamesArgs, opts ...pulumi.InvokeOption) (*LookupAppSecWapSelectedHostnamesResult, error) {
	var rv LookupAppSecWapSelectedHostnamesResult
	err := ctx.Invoke("akamai:index/getAppSecWapSelectedHostnames:getAppSecWapSelectedHostnames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecWapSelectedHostnames.
type LookupAppSecWapSelectedHostnamesArgs struct {
	// . Unique identifier of the security configuration associated with the hostnames.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the hostnames.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecWapSelectedHostnames.
type LookupAppSecWapSelectedHostnamesResult struct {
	ConfigId       int      `pulumi:"configId"`
	EvaluatedHosts []string `pulumi:"evaluatedHosts"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	Json             string   `pulumi:"json"`
	MatchTargets     string   `pulumi:"matchTargets"`
	OutputText       string   `pulumi:"outputText"`
	ProtectedHosts   []string `pulumi:"protectedHosts"`
	SecurityPolicyId string   `pulumi:"securityPolicyId"`
	SelectedHosts    []string `pulumi:"selectedHosts"`
}

func LookupAppSecWapSelectedHostnamesOutput(ctx *pulumi.Context, args LookupAppSecWapSelectedHostnamesOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecWapSelectedHostnamesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecWapSelectedHostnamesResult, error) {
			args := v.(LookupAppSecWapSelectedHostnamesArgs)
			r, err := LookupAppSecWapSelectedHostnames(ctx, &args, opts...)
			var s LookupAppSecWapSelectedHostnamesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppSecWapSelectedHostnamesResultOutput)
}

// A collection of arguments for invoking getAppSecWapSelectedHostnames.
type LookupAppSecWapSelectedHostnamesOutputArgs struct {
	// . Unique identifier of the security configuration associated with the hostnames.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the hostnames.
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (LookupAppSecWapSelectedHostnamesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecWapSelectedHostnamesArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecWapSelectedHostnames.
type LookupAppSecWapSelectedHostnamesResultOutput struct{ *pulumi.OutputState }

func (LookupAppSecWapSelectedHostnamesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecWapSelectedHostnamesResult)(nil)).Elem()
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) ToLookupAppSecWapSelectedHostnamesResultOutput() LookupAppSecWapSelectedHostnamesResultOutput {
	return o
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) ToLookupAppSecWapSelectedHostnamesResultOutputWithContext(ctx context.Context) LookupAppSecWapSelectedHostnamesResultOutput {
	return o
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) EvaluatedHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) []string { return v.EvaluatedHosts }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAppSecWapSelectedHostnamesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) MatchTargets() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) string { return v.MatchTargets }).(pulumi.StringOutput)
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) ProtectedHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) []string { return v.ProtectedHosts }).(pulumi.StringArrayOutput)
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func (o LookupAppSecWapSelectedHostnamesResultOutput) SelectedHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppSecWapSelectedHostnamesResult) []string { return v.SelectedHosts }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecWapSelectedHostnamesResultOutput{})
}
