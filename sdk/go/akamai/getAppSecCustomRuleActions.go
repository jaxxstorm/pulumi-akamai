// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAppSecCustomRuleActions` data source to retrieve information about the actions defined for the custom rules, or a specific custom rule, associated with a specific security configuration and security policy.
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
// 		opt0 := "Akamai Tools"
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		customRuleActionsAppSecCustomRuleActions, err := akamai.GetAppSecCustomRuleActions(ctx, &GetAppSecCustomRuleActionsArgs{
// 			ConfigId:         configuration.ConfigId,
// 			SecurityPolicyId: "crAP_75829",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("customRuleActions", customRuleActionsAppSecCustomRuleActions.OutputText)
// 		return nil
// 	})
// }
// ```
func GetAppSecCustomRuleActions(ctx *pulumi.Context, args *GetAppSecCustomRuleActionsArgs, opts ...pulumi.InvokeOption) (*GetAppSecCustomRuleActionsResult, error) {
	var rv GetAppSecCustomRuleActionsResult
	err := ctx.Invoke("akamai:index/getAppSecCustomRuleActions:getAppSecCustomRuleActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecCustomRuleActions.
type GetAppSecCustomRuleActionsArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// A specific custom rule for which to retrieve information. If not supplied, information about all custom rules will be returned.
	CustomRuleId *int `pulumi:"customRuleId"`
	// The ID of the security policy to use
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecCustomRuleActions.
type GetAppSecCustomRuleActionsResult struct {
	ConfigId     int  `pulumi:"configId"`
	CustomRuleId *int `pulumi:"customRuleId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A tabular display showing the ID, name, and action of all custom rules, or of the specific custom rule, associated with the specified security configuration, version and security policy.
	OutputText       string `pulumi:"outputText"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

func GetAppSecCustomRuleActionsOutput(ctx *pulumi.Context, args GetAppSecCustomRuleActionsOutputArgs, opts ...pulumi.InvokeOption) GetAppSecCustomRuleActionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecCustomRuleActionsResult, error) {
			args := v.(GetAppSecCustomRuleActionsArgs)
			r, err := GetAppSecCustomRuleActions(ctx, &args, opts...)
			return *r, err
		}).(GetAppSecCustomRuleActionsResultOutput)
}

// A collection of arguments for invoking getAppSecCustomRuleActions.
type GetAppSecCustomRuleActionsOutputArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// A specific custom rule for which to retrieve information. If not supplied, information about all custom rules will be returned.
	CustomRuleId pulumi.IntPtrInput `pulumi:"customRuleId"`
	// The ID of the security policy to use
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (GetAppSecCustomRuleActionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecCustomRuleActionsArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecCustomRuleActions.
type GetAppSecCustomRuleActionsResultOutput struct{ *pulumi.OutputState }

func (GetAppSecCustomRuleActionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecCustomRuleActionsResult)(nil)).Elem()
}

func (o GetAppSecCustomRuleActionsResultOutput) ToGetAppSecCustomRuleActionsResultOutput() GetAppSecCustomRuleActionsResultOutput {
	return o
}

func (o GetAppSecCustomRuleActionsResultOutput) ToGetAppSecCustomRuleActionsResultOutputWithContext(ctx context.Context) GetAppSecCustomRuleActionsResultOutput {
	return o
}

func (o GetAppSecCustomRuleActionsResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecCustomRuleActionsResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

func (o GetAppSecCustomRuleActionsResultOutput) CustomRuleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecCustomRuleActionsResult) *int { return v.CustomRuleId }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecCustomRuleActionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecCustomRuleActionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A tabular display showing the ID, name, and action of all custom rules, or of the specific custom rule, associated with the specified security configuration, version and security policy.
func (o GetAppSecCustomRuleActionsResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecCustomRuleActionsResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecCustomRuleActionsResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecCustomRuleActionsResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecCustomRuleActionsResultOutput{})
}
