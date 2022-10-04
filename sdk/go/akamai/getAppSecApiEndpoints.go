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
// Returns information about the API endpoints associated with a security policy or configuration.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-endpoints](https://techdocs.akamai.com/application-security/reference/get-api-endpoints)
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
//			_, err := akamai.GetAppSecApiEndpoints(ctx, &GetAppSecApiEndpointsArgs{
//				ApiName:  pulumi.StringRef("Contracts"),
//				ConfigId: 58843,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `idList`. List of API endpoint IDs.
// - `json`. JSON-formatted list of information about the API endpoints.
// - `outputText`. Tabular report showing the ID and name of the API endpoints.
func GetAppSecApiEndpoints(ctx *pulumi.Context, args *GetAppSecApiEndpointsArgs, opts ...pulumi.InvokeOption) (*GetAppSecApiEndpointsResult, error) {
	var rv GetAppSecApiEndpointsResult
	err := ctx.Invoke("akamai:index/getAppSecApiEndpoints:getAppSecApiEndpoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecApiEndpoints.
type GetAppSecApiEndpointsArgs struct {
	// . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
	ApiName *string `pulumi:"apiName"`
	// . Unique identifier of the security configuration associated with the API endpoints.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecApiEndpoints.
type GetAppSecApiEndpointsResult struct {
	ApiName  *string `pulumi:"apiName"`
	ConfigId int     `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	IdLists          []int   `pulumi:"idLists"`
	Json             string  `pulumi:"json"`
	OutputText       string  `pulumi:"outputText"`
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

func GetAppSecApiEndpointsOutput(ctx *pulumi.Context, args GetAppSecApiEndpointsOutputArgs, opts ...pulumi.InvokeOption) GetAppSecApiEndpointsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecApiEndpointsResult, error) {
			args := v.(GetAppSecApiEndpointsArgs)
			r, err := GetAppSecApiEndpoints(ctx, &args, opts...)
			var s GetAppSecApiEndpointsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppSecApiEndpointsResultOutput)
}

// A collection of arguments for invoking getAppSecApiEndpoints.
type GetAppSecApiEndpointsOutputArgs struct {
	// . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
	ApiName pulumi.StringPtrInput `pulumi:"apiName"`
	// . Unique identifier of the security configuration associated with the API endpoints.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
	SecurityPolicyId pulumi.StringPtrInput `pulumi:"securityPolicyId"`
}

func (GetAppSecApiEndpointsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecApiEndpointsArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecApiEndpoints.
type GetAppSecApiEndpointsResultOutput struct{ *pulumi.OutputState }

func (GetAppSecApiEndpointsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecApiEndpointsResult)(nil)).Elem()
}

func (o GetAppSecApiEndpointsResultOutput) ToGetAppSecApiEndpointsResultOutput() GetAppSecApiEndpointsResultOutput {
	return o
}

func (o GetAppSecApiEndpointsResultOutput) ToGetAppSecApiEndpointsResultOutputWithContext(ctx context.Context) GetAppSecApiEndpointsResultOutput {
	return o
}

func (o GetAppSecApiEndpointsResultOutput) ApiName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppSecApiEndpointsResult) *string { return v.ApiName }).(pulumi.StringPtrOutput)
}

func (o GetAppSecApiEndpointsResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecApiEndpointsResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecApiEndpointsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecApiEndpointsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecApiEndpointsResultOutput) IdLists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetAppSecApiEndpointsResult) []int { return v.IdLists }).(pulumi.IntArrayOutput)
}

func (o GetAppSecApiEndpointsResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecApiEndpointsResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetAppSecApiEndpointsResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecApiEndpointsResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecApiEndpointsResultOutput) SecurityPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppSecApiEndpointsResult) *string { return v.SecurityPolicyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecApiEndpointsResultOutput{})
}
