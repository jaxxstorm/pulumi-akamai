// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppSecSecurityPolicy(ctx *pulumi.Context, args *LookupAppSecSecurityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAppSecSecurityPolicyResult, error) {
	var rv LookupAppSecSecurityPolicyResult
	err := ctx.Invoke("akamai:index/getAppSecSecurityPolicy:getAppSecSecurityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecSecurityPolicy.
type LookupAppSecSecurityPolicyArgs struct {
	// . Unique identifier of the security configuration associated with the security policies.
	// - `securityPolicyName`. (Optional). Name of the security policy you want to return information for (be sure to reference the policy name and not the policy ID). If not included, information is returned for all your security policies.
	ConfigId           int     `pulumi:"configId"`
	SecurityPolicyName *string `pulumi:"securityPolicyName"`
}

// A collection of values returned by getAppSecSecurityPolicy.
type LookupAppSecSecurityPolicyResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string   `pulumi:"id"`
	OutputText            string   `pulumi:"outputText"`
	SecurityPolicyId      string   `pulumi:"securityPolicyId"`
	SecurityPolicyIdLists []string `pulumi:"securityPolicyIdLists"`
	SecurityPolicyName    *string  `pulumi:"securityPolicyName"`
}

func LookupAppSecSecurityPolicyOutput(ctx *pulumi.Context, args LookupAppSecSecurityPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecSecurityPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecSecurityPolicyResult, error) {
			args := v.(LookupAppSecSecurityPolicyArgs)
			r, err := LookupAppSecSecurityPolicy(ctx, &args, opts...)
			var s LookupAppSecSecurityPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppSecSecurityPolicyResultOutput)
}

// A collection of arguments for invoking getAppSecSecurityPolicy.
type LookupAppSecSecurityPolicyOutputArgs struct {
	// . Unique identifier of the security configuration associated with the security policies.
	// - `securityPolicyName`. (Optional). Name of the security policy you want to return information for (be sure to reference the policy name and not the policy ID). If not included, information is returned for all your security policies.
	ConfigId           pulumi.IntInput       `pulumi:"configId"`
	SecurityPolicyName pulumi.StringPtrInput `pulumi:"securityPolicyName"`
}

func (LookupAppSecSecurityPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecSecurityPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecSecurityPolicy.
type LookupAppSecSecurityPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAppSecSecurityPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecSecurityPolicyResult)(nil)).Elem()
}

func (o LookupAppSecSecurityPolicyResultOutput) ToLookupAppSecSecurityPolicyResultOutput() LookupAppSecSecurityPolicyResultOutput {
	return o
}

func (o LookupAppSecSecurityPolicyResultOutput) ToLookupAppSecSecurityPolicyResultOutputWithContext(ctx context.Context) LookupAppSecSecurityPolicyResultOutput {
	return o
}

func (o LookupAppSecSecurityPolicyResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecSecurityPolicyResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAppSecSecurityPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecSecurityPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppSecSecurityPolicyResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecSecurityPolicyResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o LookupAppSecSecurityPolicyResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecSecurityPolicyResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

func (o LookupAppSecSecurityPolicyResultOutput) SecurityPolicyIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppSecSecurityPolicyResult) []string { return v.SecurityPolicyIdLists }).(pulumi.StringArrayOutput)
}

func (o LookupAppSecSecurityPolicyResultOutput) SecurityPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppSecSecurityPolicyResult) *string { return v.SecurityPolicyName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecSecurityPolicyResultOutput{})
}
