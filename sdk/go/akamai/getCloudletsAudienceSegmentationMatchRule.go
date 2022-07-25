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
// Use the `getCloudletsAudienceSegmentationMatchRule` data source to build a match rule JSON object for the Audience Segmentation Cloudlet.
//
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `type` - The type of Cloudlet the rule is for.
// * `json` - A `matchRules` JSON structure generated from the API schema that defines the rules for this policy.
func GetCloudletsAudienceSegmentationMatchRule(ctx *pulumi.Context, args *GetCloudletsAudienceSegmentationMatchRuleArgs, opts ...pulumi.InvokeOption) (*GetCloudletsAudienceSegmentationMatchRuleResult, error) {
	var rv GetCloudletsAudienceSegmentationMatchRuleResult
	err := ctx.Invoke("akamai:index/getCloudletsAudienceSegmentationMatchRule:getCloudletsAudienceSegmentationMatchRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudletsAudienceSegmentationMatchRule.
type GetCloudletsAudienceSegmentationMatchRuleArgs struct {
	// - (Optional) A list of Cloudlet-specific match rules for a policy.
	MatchRules []GetCloudletsAudienceSegmentationMatchRuleMatchRule `pulumi:"matchRules"`
}

// A collection of values returned by getCloudletsAudienceSegmentationMatchRule.
type GetCloudletsAudienceSegmentationMatchRuleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                                               `pulumi:"id"`
	Json       string                                               `pulumi:"json"`
	MatchRules []GetCloudletsAudienceSegmentationMatchRuleMatchRule `pulumi:"matchRules"`
}

func GetCloudletsAudienceSegmentationMatchRuleOutput(ctx *pulumi.Context, args GetCloudletsAudienceSegmentationMatchRuleOutputArgs, opts ...pulumi.InvokeOption) GetCloudletsAudienceSegmentationMatchRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudletsAudienceSegmentationMatchRuleResult, error) {
			args := v.(GetCloudletsAudienceSegmentationMatchRuleArgs)
			r, err := GetCloudletsAudienceSegmentationMatchRule(ctx, &args, opts...)
			var s GetCloudletsAudienceSegmentationMatchRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudletsAudienceSegmentationMatchRuleResultOutput)
}

// A collection of arguments for invoking getCloudletsAudienceSegmentationMatchRule.
type GetCloudletsAudienceSegmentationMatchRuleOutputArgs struct {
	// - (Optional) A list of Cloudlet-specific match rules for a policy.
	MatchRules GetCloudletsAudienceSegmentationMatchRuleMatchRuleArrayInput `pulumi:"matchRules"`
}

func (GetCloudletsAudienceSegmentationMatchRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudletsAudienceSegmentationMatchRuleArgs)(nil)).Elem()
}

// A collection of values returned by getCloudletsAudienceSegmentationMatchRule.
type GetCloudletsAudienceSegmentationMatchRuleResultOutput struct{ *pulumi.OutputState }

func (GetCloudletsAudienceSegmentationMatchRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudletsAudienceSegmentationMatchRuleResult)(nil)).Elem()
}

func (o GetCloudletsAudienceSegmentationMatchRuleResultOutput) ToGetCloudletsAudienceSegmentationMatchRuleResultOutput() GetCloudletsAudienceSegmentationMatchRuleResultOutput {
	return o
}

func (o GetCloudletsAudienceSegmentationMatchRuleResultOutput) ToGetCloudletsAudienceSegmentationMatchRuleResultOutputWithContext(ctx context.Context) GetCloudletsAudienceSegmentationMatchRuleResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloudletsAudienceSegmentationMatchRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudletsAudienceSegmentationMatchRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudletsAudienceSegmentationMatchRuleResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudletsAudienceSegmentationMatchRuleResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetCloudletsAudienceSegmentationMatchRuleResultOutput) MatchRules() GetCloudletsAudienceSegmentationMatchRuleMatchRuleArrayOutput {
	return o.ApplyT(func(v GetCloudletsAudienceSegmentationMatchRuleResult) []GetCloudletsAudienceSegmentationMatchRuleMatchRule {
		return v.MatchRules
	}).(GetCloudletsAudienceSegmentationMatchRuleMatchRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudletsAudienceSegmentationMatchRuleResultOutput{})
}
