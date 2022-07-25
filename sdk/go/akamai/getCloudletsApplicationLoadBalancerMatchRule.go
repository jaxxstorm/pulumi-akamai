// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
//
// Use the `getCloudletsApplicationLoadBalancerMatchRule` data source to build a match rule JSON object for the Application Load Balancer Cloudlet.
//
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `type` - The type of Cloudlet the rule is for.
// * `json` - A `matchRules` JSON structure generated from the API schema that defines the rules for this policy.
func GetCloudletsApplicationLoadBalancerMatchRule(ctx *pulumi.Context, args *GetCloudletsApplicationLoadBalancerMatchRuleArgs, opts ...pulumi.InvokeOption) (*GetCloudletsApplicationLoadBalancerMatchRuleResult, error) {
	var rv GetCloudletsApplicationLoadBalancerMatchRuleResult
	err := ctx.Invoke("akamai:index/getCloudletsApplicationLoadBalancerMatchRule:getCloudletsApplicationLoadBalancerMatchRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudletsApplicationLoadBalancerMatchRule.
type GetCloudletsApplicationLoadBalancerMatchRuleArgs struct {
	// - (Optional) A list of Cloudlet-specific match rules for a policy.
	MatchRules []GetCloudletsApplicationLoadBalancerMatchRuleMatchRule `pulumi:"matchRules"`
}

// A collection of values returned by getCloudletsApplicationLoadBalancerMatchRule.
type GetCloudletsApplicationLoadBalancerMatchRuleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                                                  `pulumi:"id"`
	Json       string                                                  `pulumi:"json"`
	MatchRules []GetCloudletsApplicationLoadBalancerMatchRuleMatchRule `pulumi:"matchRules"`
}

func GetCloudletsApplicationLoadBalancerMatchRuleOutput(ctx *pulumi.Context, args GetCloudletsApplicationLoadBalancerMatchRuleOutputArgs, opts ...pulumi.InvokeOption) GetCloudletsApplicationLoadBalancerMatchRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudletsApplicationLoadBalancerMatchRuleResult, error) {
			args := v.(GetCloudletsApplicationLoadBalancerMatchRuleArgs)
			r, err := GetCloudletsApplicationLoadBalancerMatchRule(ctx, &args, opts...)
			var s GetCloudletsApplicationLoadBalancerMatchRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudletsApplicationLoadBalancerMatchRuleResultOutput)
}

// A collection of arguments for invoking getCloudletsApplicationLoadBalancerMatchRule.
type GetCloudletsApplicationLoadBalancerMatchRuleOutputArgs struct {
	// - (Optional) A list of Cloudlet-specific match rules for a policy.
	MatchRules GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArrayInput `pulumi:"matchRules"`
}

func (GetCloudletsApplicationLoadBalancerMatchRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudletsApplicationLoadBalancerMatchRuleArgs)(nil)).Elem()
}

// A collection of values returned by getCloudletsApplicationLoadBalancerMatchRule.
type GetCloudletsApplicationLoadBalancerMatchRuleResultOutput struct{ *pulumi.OutputState }

func (GetCloudletsApplicationLoadBalancerMatchRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudletsApplicationLoadBalancerMatchRuleResult)(nil)).Elem()
}

func (o GetCloudletsApplicationLoadBalancerMatchRuleResultOutput) ToGetCloudletsApplicationLoadBalancerMatchRuleResultOutput() GetCloudletsApplicationLoadBalancerMatchRuleResultOutput {
	return o
}

func (o GetCloudletsApplicationLoadBalancerMatchRuleResultOutput) ToGetCloudletsApplicationLoadBalancerMatchRuleResultOutputWithContext(ctx context.Context) GetCloudletsApplicationLoadBalancerMatchRuleResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloudletsApplicationLoadBalancerMatchRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudletsApplicationLoadBalancerMatchRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudletsApplicationLoadBalancerMatchRuleResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudletsApplicationLoadBalancerMatchRuleResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetCloudletsApplicationLoadBalancerMatchRuleResultOutput) MatchRules() GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArrayOutput {
	return o.ApplyT(func(v GetCloudletsApplicationLoadBalancerMatchRuleResult) []GetCloudletsApplicationLoadBalancerMatchRuleMatchRule {
		return v.MatchRules
	}).(GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudletsApplicationLoadBalancerMatchRuleResultOutput{})
}
