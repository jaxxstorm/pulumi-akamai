// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy; rate policy
//
// Returns information about your rate policy actions. Actions specify what happens any time a rate policy is triggered: the issue could be ignored, the request could be denied, or an alert could be generated.
//
// **Related API Endpoint:** [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rate-policies](https://techdocs.akamai.com/application-security/reference/get-rate-policies-actions)
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
//			ratePolicyActionsAppSecRatePolicyActions, err := akamai.GetAppSecRatePolicyActions(ctx, &GetAppSecRatePolicyActionsArgs{
//				ConfigId:         configuration.ConfigId,
//				SecurityPolicyId: "gms1_134637",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ratePolicyActions", ratePolicyActionsAppSecRatePolicyActions.OutputText)
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `outputText`. Tabular report showing the ID, IPv4 action, and IPv6 action of the rate policies.
func GetAppSecRatePolicyActions(ctx *pulumi.Context, args *GetAppSecRatePolicyActionsArgs, opts ...pulumi.InvokeOption) (*GetAppSecRatePolicyActionsResult, error) {
	var rv GetAppSecRatePolicyActionsResult
	err := ctx.Invoke("akamai:index/getAppSecRatePolicyActions:getAppSecRatePolicyActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecRatePolicyActions.
type GetAppSecRatePolicyActionsArgs struct {
	// . Unique identifier of the security configuration associated with the rate policies and rate policy actions.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the rate policy you want to return action information for. If not included, action information is returned for all your rate policies.
	RatePolicyId *int `pulumi:"ratePolicyId"`
	// . Unique identifier of the security policy associated with the rate policies and rate policy actions.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecRatePolicyActions.
type GetAppSecRatePolicyActionsResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	OutputText       string `pulumi:"outputText"`
	RatePolicyId     *int   `pulumi:"ratePolicyId"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

func GetAppSecRatePolicyActionsOutput(ctx *pulumi.Context, args GetAppSecRatePolicyActionsOutputArgs, opts ...pulumi.InvokeOption) GetAppSecRatePolicyActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecRatePolicyActionsResult, error) {
			args := v.(GetAppSecRatePolicyActionsArgs)
			r, err := GetAppSecRatePolicyActions(ctx, &args, opts...)
			var s GetAppSecRatePolicyActionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppSecRatePolicyActionsResultOutput)
}

// A collection of arguments for invoking getAppSecRatePolicyActions.
type GetAppSecRatePolicyActionsOutputArgs struct {
	// . Unique identifier of the security configuration associated with the rate policies and rate policy actions.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the rate policy you want to return action information for. If not included, action information is returned for all your rate policies.
	RatePolicyId pulumi.IntPtrInput `pulumi:"ratePolicyId"`
	// . Unique identifier of the security policy associated with the rate policies and rate policy actions.
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (GetAppSecRatePolicyActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecRatePolicyActionsArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecRatePolicyActions.
type GetAppSecRatePolicyActionsResultOutput struct{ *pulumi.OutputState }

func (GetAppSecRatePolicyActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecRatePolicyActionsResult)(nil)).Elem()
}

func (o GetAppSecRatePolicyActionsResultOutput) ToGetAppSecRatePolicyActionsResultOutput() GetAppSecRatePolicyActionsResultOutput {
	return o
}

func (o GetAppSecRatePolicyActionsResultOutput) ToGetAppSecRatePolicyActionsResultOutputWithContext(ctx context.Context) GetAppSecRatePolicyActionsResultOutput {
	return o
}

func (o GetAppSecRatePolicyActionsResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecRatePolicyActionsResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecRatePolicyActionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRatePolicyActionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecRatePolicyActionsResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRatePolicyActionsResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecRatePolicyActionsResultOutput) RatePolicyId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecRatePolicyActionsResult) *int { return v.RatePolicyId }).(pulumi.IntPtrOutput)
}

func (o GetAppSecRatePolicyActionsResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecRatePolicyActionsResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecRatePolicyActionsResultOutput{})
}
