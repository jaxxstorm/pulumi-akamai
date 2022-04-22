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
// Use the `getCloudletsForwardRewriteMatchRule` data source to build a match rule JSON object for the Forward Rewrite Cloudlet.
//
// ## Basic usage
//
// This example returns the JSON-encoded rules for the Forward Rewrite Cloudlet:
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
// 		_, err := akamai.GetCloudletsForwardRewriteMatchRule(ctx, &GetCloudletsForwardRewriteMatchRuleArgs{
// 			MatchRules: []GetCloudletsForwardRewriteMatchRuleMatchRule{
// 				GetCloudletsForwardRewriteMatchRuleMatchRule{
// 					ForwardSettings: GetCloudletsForwardRewriteMatchRuleMatchRuleForwardSettings{
// 						OriginId:               pulumi.StringRef("1234"),
// 						PathAndQs:              pulumi.StringRef("/path"),
// 						UseIncomingQueryString: pulumi.BoolRef(true),
// 					},
// 					Matches: []GetCloudletsForwardRewriteMatchRuleMatchRuleMatch{
// 						GetCloudletsForwardRewriteMatchRuleMatchRuleMatch{
// 							CaseSensitive: pulumi.BoolRef(false),
// 							CheckIps:      pulumi.StringRef("CONNECTING_IP XFF_HEADERS"),
// 							MatchOperator: pulumi.StringRef("equals"),
// 							MatchType:     pulumi.StringRef("clientip"),
// 							MatchValue:    pulumi.StringRef("192.0.2.0"),
// 							Negate:        pulumi.BoolRef(false),
// 						},
// 					},
// 					Name: pulumi.StringRef("rule"),
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `type` - The type of Cloudlet the rule is for.
// * `json` - A `matchRules` JSON structure generated from the API schema that defines the rules for this policy.
func GetCloudletsForwardRewriteMatchRule(ctx *pulumi.Context, args *GetCloudletsForwardRewriteMatchRuleArgs, opts ...pulumi.InvokeOption) (*GetCloudletsForwardRewriteMatchRuleResult, error) {
	var rv GetCloudletsForwardRewriteMatchRuleResult
	err := ctx.Invoke("akamai:index/getCloudletsForwardRewriteMatchRule:getCloudletsForwardRewriteMatchRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudletsForwardRewriteMatchRule.
type GetCloudletsForwardRewriteMatchRuleArgs struct {
	// - (Optional) A list of Cloudlet-specific match rules for a policy.
	MatchRules []GetCloudletsForwardRewriteMatchRuleMatchRule `pulumi:"matchRules"`
}

// A collection of values returned by getCloudletsForwardRewriteMatchRule.
type GetCloudletsForwardRewriteMatchRuleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                                         `pulumi:"id"`
	Json       string                                         `pulumi:"json"`
	MatchRules []GetCloudletsForwardRewriteMatchRuleMatchRule `pulumi:"matchRules"`
}

func GetCloudletsForwardRewriteMatchRuleOutput(ctx *pulumi.Context, args GetCloudletsForwardRewriteMatchRuleOutputArgs, opts ...pulumi.InvokeOption) GetCloudletsForwardRewriteMatchRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudletsForwardRewriteMatchRuleResult, error) {
			args := v.(GetCloudletsForwardRewriteMatchRuleArgs)
			r, err := GetCloudletsForwardRewriteMatchRule(ctx, &args, opts...)
			var s GetCloudletsForwardRewriteMatchRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudletsForwardRewriteMatchRuleResultOutput)
}

// A collection of arguments for invoking getCloudletsForwardRewriteMatchRule.
type GetCloudletsForwardRewriteMatchRuleOutputArgs struct {
	// - (Optional) A list of Cloudlet-specific match rules for a policy.
	MatchRules GetCloudletsForwardRewriteMatchRuleMatchRuleArrayInput `pulumi:"matchRules"`
}

func (GetCloudletsForwardRewriteMatchRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudletsForwardRewriteMatchRuleArgs)(nil)).Elem()
}

// A collection of values returned by getCloudletsForwardRewriteMatchRule.
type GetCloudletsForwardRewriteMatchRuleResultOutput struct{ *pulumi.OutputState }

func (GetCloudletsForwardRewriteMatchRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudletsForwardRewriteMatchRuleResult)(nil)).Elem()
}

func (o GetCloudletsForwardRewriteMatchRuleResultOutput) ToGetCloudletsForwardRewriteMatchRuleResultOutput() GetCloudletsForwardRewriteMatchRuleResultOutput {
	return o
}

func (o GetCloudletsForwardRewriteMatchRuleResultOutput) ToGetCloudletsForwardRewriteMatchRuleResultOutputWithContext(ctx context.Context) GetCloudletsForwardRewriteMatchRuleResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCloudletsForwardRewriteMatchRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudletsForwardRewriteMatchRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudletsForwardRewriteMatchRuleResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudletsForwardRewriteMatchRuleResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetCloudletsForwardRewriteMatchRuleResultOutput) MatchRules() GetCloudletsForwardRewriteMatchRuleMatchRuleArrayOutput {
	return o.ApplyT(func(v GetCloudletsForwardRewriteMatchRuleResult) []GetCloudletsForwardRewriteMatchRuleMatchRule {
		return v.MatchRules
	}).(GetCloudletsForwardRewriteMatchRuleMatchRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudletsForwardRewriteMatchRuleResultOutput{})
}
