// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy; API endpoint
//
// Returns information about API endpoint constraints and actions. The returned information is described in the [API Constraints members](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapirequestconstraints) section of the Application Security API.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-request-constraints](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapirequestconstraints)
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
// 		apisRequestConstraints, err := akamai.LookupAppSecApiRequestConstraints(ctx, &GetAppSecApiRequestConstraintsArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: "gms1_134637",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("apisConstraintsText", apisRequestConstraints.OutputText)
// 		ctx.Export("apisConstraintsJson", apisRequestConstraints.Json)
// 		apiRequestConstraints, err := akamai.LookupAppSecApiRequestConstraints(ctx, &GetAppSecApiRequestConstraintsArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: "gms1_134637",
// 			ApiId:            pulumi.IntRef(624913),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("apiConstraintsText", apiRequestConstraints.OutputText)
// 		ctx.Export("apiConstraintsJson", apiRequestConstraints.Json)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `json`. JSON-formatted list of information about the APIs, their constraints, and their actions.
// - `outputText`. Tabular report of the APIs, their constraints, and their actions.
func LookupAppSecApiRequestConstraints(ctx *pulumi.Context, args *LookupAppSecApiRequestConstraintsArgs, opts ...pulumi.InvokeOption) (*LookupAppSecApiRequestConstraintsResult, error) {
	var rv LookupAppSecApiRequestConstraintsResult
	err := ctx.Invoke("akamai:index/getAppSecApiRequestConstraints:getAppSecApiRequestConstraints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecApiRequestConstraints.
type LookupAppSecApiRequestConstraintsArgs struct {
	// . Unique identifier of the API endpoint you want to return constraint information for.
	ApiId *int `pulumi:"apiId"`
	// . Unique identifier of the security configuration associated with the API constraints.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the API constraints.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecApiRequestConstraints.
type LookupAppSecApiRequestConstraintsResult struct {
	ApiId    *int `pulumi:"apiId"`
	ConfigId int  `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	Json             string `pulumi:"json"`
	OutputText       string `pulumi:"outputText"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

func LookupAppSecApiRequestConstraintsOutput(ctx *pulumi.Context, args LookupAppSecApiRequestConstraintsOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecApiRequestConstraintsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecApiRequestConstraintsResult, error) {
			args := v.(LookupAppSecApiRequestConstraintsArgs)
			r, err := LookupAppSecApiRequestConstraints(ctx, &args, opts...)
			return *r, err
		}).(LookupAppSecApiRequestConstraintsResultOutput)
}

// A collection of arguments for invoking getAppSecApiRequestConstraints.
type LookupAppSecApiRequestConstraintsOutputArgs struct {
	// . Unique identifier of the API endpoint you want to return constraint information for.
	ApiId pulumi.IntPtrInput `pulumi:"apiId"`
	// . Unique identifier of the security configuration associated with the API constraints.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the API constraints.
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (LookupAppSecApiRequestConstraintsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecApiRequestConstraintsArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecApiRequestConstraints.
type LookupAppSecApiRequestConstraintsResultOutput struct{ *pulumi.OutputState }

func (LookupAppSecApiRequestConstraintsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecApiRequestConstraintsResult)(nil)).Elem()
}

func (o LookupAppSecApiRequestConstraintsResultOutput) ToLookupAppSecApiRequestConstraintsResultOutput() LookupAppSecApiRequestConstraintsResultOutput {
	return o
}

func (o LookupAppSecApiRequestConstraintsResultOutput) ToLookupAppSecApiRequestConstraintsResultOutputWithContext(ctx context.Context) LookupAppSecApiRequestConstraintsResultOutput {
	return o
}

func (o LookupAppSecApiRequestConstraintsResultOutput) ApiId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppSecApiRequestConstraintsResult) *int { return v.ApiId }).(pulumi.IntPtrOutput)
}

func (o LookupAppSecApiRequestConstraintsResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecApiRequestConstraintsResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAppSecApiRequestConstraintsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecApiRequestConstraintsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppSecApiRequestConstraintsResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecApiRequestConstraintsResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o LookupAppSecApiRequestConstraintsResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecApiRequestConstraintsResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o LookupAppSecApiRequestConstraintsResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecApiRequestConstraintsResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecApiRequestConstraintsResultOutput{})
}
