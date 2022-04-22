// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration; rate policy
//
// Returns information about your rate policies. Rate polices help you monitor and moderate the number and rate of all the requests you receive; in turn, this helps you prevent your website from being overwhelmed by a dramatic, and unexpected, surge in traffic.
//
// **Related API Endpoint:** [/appsec/v1/configs/{configId}/versions/{versionNumber}/rate-policies](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getratepolicies)
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
// 		ratePolicies, err := akamai.GetAppSecRatePolicies(ctx, &GetAppSecRatePoliciesArgs{
// 			ConfigId: configuration.ConfigId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ratePoliciesOutput", ratePolicies.OutputText)
// 		ctx.Export("ratePoliciesJson", ratePolicies.Json)
// 		ratePolicy, err := akamai.GetAppSecRatePolicies(ctx, &GetAppSecRatePoliciesArgs{
// 			ConfigId:     configuration.ConfigId,
// 			RatePolicyId: pulumi.IntRef(122149),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ratePolicyJson", ratePolicy.Json)
// 		ctx.Export("ratePolicyOutput", ratePolicy.OutputText)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `outputText`. Tabular report showing the ID and name of the rate policies.
// - `json`. JSON-formatted list of the rate policy information.
func GetAppSecRatePolicies(ctx *pulumi.Context, args *GetAppSecRatePoliciesArgs, opts ...pulumi.InvokeOption) (*GetAppSecRatePoliciesResult, error) {
	var rv GetAppSecRatePoliciesResult
	err := ctx.Invoke("akamai:index/getAppSecRatePolicies:getAppSecRatePolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecRatePolicies.
type GetAppSecRatePoliciesArgs struct {
	// . Unique identifier of the security configuration associated with the rate policies.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the rate policy you want to return information for. If not included, information is returned for all your rate policies.
	RatePolicyId *int `pulumi:"ratePolicyId"`
}

// A collection of values returned by getAppSecRatePolicies.
type GetAppSecRatePoliciesResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	Json         string `pulumi:"json"`
	OutputText   string `pulumi:"outputText"`
	RatePolicyId *int   `pulumi:"ratePolicyId"`
}

func GetAppSecRatePoliciesOutput(ctx *pulumi.Context, args GetAppSecRatePoliciesOutputArgs, opts ...pulumi.InvokeOption) GetAppSecRatePoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecRatePoliciesResult, error) {
			args := v.(GetAppSecRatePoliciesArgs)
			r, err := GetAppSecRatePolicies(ctx, &args, opts...)
			var s GetAppSecRatePoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppSecRatePoliciesResultOutput)
}

// A collection of arguments for invoking getAppSecRatePolicies.
type GetAppSecRatePoliciesOutputArgs struct {
	// . Unique identifier of the security configuration associated with the rate policies.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the rate policy you want to return information for. If not included, information is returned for all your rate policies.
	RatePolicyId pulumi.IntPtrInput `pulumi:"ratePolicyId"`
}

func (GetAppSecRatePoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecRatePoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecRatePolicies.
type GetAppSecRatePoliciesResultOutput struct{ *pulumi.OutputState }

func (GetAppSecRatePoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecRatePoliciesResult)(nil)).Elem()
}

func (o GetAppSecRatePoliciesResultOutput) ToGetAppSecRatePoliciesResultOutput() GetAppSecRatePoliciesResultOutput {
	return o
}

func (o GetAppSecRatePoliciesResultOutput) ToGetAppSecRatePoliciesResultOutputWithContext(ctx context.Context) GetAppSecRatePoliciesResultOutput {
	return o
}

func (o GetAppSecRatePoliciesResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecRatePoliciesResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecRatePoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRatePoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecRatePoliciesResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRatePoliciesResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetAppSecRatePoliciesResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRatePoliciesResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecRatePoliciesResultOutput) RatePolicyId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecRatePoliciesResult) *int { return v.RatePolicyId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecRatePoliciesResultOutput{})
}
