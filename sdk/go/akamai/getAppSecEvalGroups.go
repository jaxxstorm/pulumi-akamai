// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppSecEvalGroups(ctx *pulumi.Context, args *GetAppSecEvalGroupsArgs, opts ...pulumi.InvokeOption) (*GetAppSecEvalGroupsResult, error) {
	var rv GetAppSecEvalGroupsResult
	err := ctx.Invoke("akamai:index/getAppSecEvalGroups:getAppSecEvalGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecEvalGroups.
type GetAppSecEvalGroupsArgs struct {
	// . Unique identifier of the evaluation attack group you want to return information for. If not included, information is returned for all your evaluation attack groups.
	AttackGroup *string `pulumi:"attackGroup"`
	// . Unique identifier of the security configuration associated with the evaluation attack group.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the evaluation attack group.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecEvalGroups.
type GetAppSecEvalGroupsResult struct {
	AttackGroup        *string `pulumi:"attackGroup"`
	AttackGroupAction  string  `pulumi:"attackGroupAction"`
	ConditionException string  `pulumi:"conditionException"`
	ConfigId           int     `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	Json             string `pulumi:"json"`
	OutputText       string `pulumi:"outputText"`
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

func GetAppSecEvalGroupsOutput(ctx *pulumi.Context, args GetAppSecEvalGroupsOutputArgs, opts ...pulumi.InvokeOption) GetAppSecEvalGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecEvalGroupsResult, error) {
			args := v.(GetAppSecEvalGroupsArgs)
			r, err := GetAppSecEvalGroups(ctx, &args, opts...)
			return *r, err
		}).(GetAppSecEvalGroupsResultOutput)
}

// A collection of arguments for invoking getAppSecEvalGroups.
type GetAppSecEvalGroupsOutputArgs struct {
	// . Unique identifier of the evaluation attack group you want to return information for. If not included, information is returned for all your evaluation attack groups.
	AttackGroup pulumi.StringPtrInput `pulumi:"attackGroup"`
	// . Unique identifier of the security configuration associated with the evaluation attack group.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the evaluation attack group.
	SecurityPolicyId pulumi.StringInput `pulumi:"securityPolicyId"`
}

func (GetAppSecEvalGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecEvalGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecEvalGroups.
type GetAppSecEvalGroupsResultOutput struct{ *pulumi.OutputState }

func (GetAppSecEvalGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecEvalGroupsResult)(nil)).Elem()
}

func (o GetAppSecEvalGroupsResultOutput) ToGetAppSecEvalGroupsResultOutput() GetAppSecEvalGroupsResultOutput {
	return o
}

func (o GetAppSecEvalGroupsResultOutput) ToGetAppSecEvalGroupsResultOutputWithContext(ctx context.Context) GetAppSecEvalGroupsResultOutput {
	return o
}

func (o GetAppSecEvalGroupsResultOutput) AttackGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppSecEvalGroupsResult) *string { return v.AttackGroup }).(pulumi.StringPtrOutput)
}

func (o GetAppSecEvalGroupsResultOutput) AttackGroupAction() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalGroupsResult) string { return v.AttackGroupAction }).(pulumi.StringOutput)
}

func (o GetAppSecEvalGroupsResultOutput) ConditionException() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalGroupsResult) string { return v.ConditionException }).(pulumi.StringOutput)
}

func (o GetAppSecEvalGroupsResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecEvalGroupsResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecEvalGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecEvalGroupsResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalGroupsResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetAppSecEvalGroupsResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalGroupsResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecEvalGroupsResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecEvalGroupsResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecEvalGroupsResultOutput{})
}
