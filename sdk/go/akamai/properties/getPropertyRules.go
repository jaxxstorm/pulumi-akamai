// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package properties

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `properties.PropertyRules` data source allows you to configure a nested block of property rules, criteria, and behaviors. A property’s main functionality is encapsulated in its set of rules and rules are composed of the matches and the behavior that applies under those matches.
//
// ## Example Usage
// ### Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/go/akamai/properties"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		examplePropertyRules, err := properties.LookupPropertyRules(ctx, &properties.LookupPropertyRulesArgs{
// 			Rules: []properties.GetPropertyRulesRule{
// 				properties.GetPropertyRulesRule{
// 					Behaviors: []properties.GetPropertyRulesRuleBehavior{
// 						properties.GetPropertyRulesRuleBehavior{
// 							Name: "downstreamCache",
// 							Option: []map[string]interface{}{
// 								map[string]interface{}{
// 									"key":   "behavior",
// 									"value": "TUNNEL_ORIGIN",
// 								},
// 							},
// 						},
// 					},
// 					Rules: []properties.GetPropertyRulesRuleRule{
// 						properties.GetPropertyRulesRuleRule{
// 							Name: "Performance",
// 							Rule: []map[string]interface{}{
// 								map[string]interface{}{
// 									"behavior": []map[string]interface{}{
// 										map[string]interface{}{
// 											"name": "adaptiveImageCompression",
// 											"option": []map[string]interface{}{
// 												map[string]interface{}{
// 													"key":   "tier1MobileCompressionMethod",
// 													"value": "COMPRESS",
// 												},
// 												map[string]interface{}{
// 													"key":   "tier1MobileCompressionValue",
// 													"value": "80",
// 												},
// 												map[string]interface{}{
// 													"key":   "tier2MobileCompressionMethod",
// 													"value": "COMPRESS",
// 												},
// 											},
// 										},
// 									},
// 									"name": "JPEG Images",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = properties.NewProperty(ctx, "exampleProperty", &properties.PropertyArgs{
// 			Rules: pulumi.String(examplePropertyRules.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPropertyRules(ctx *pulumi.Context, args *LookupPropertyRulesArgs, opts ...pulumi.InvokeOption) (*LookupPropertyRulesResult, error) {
	var rv LookupPropertyRulesResult
	err := ctx.Invoke("akamai:properties/getPropertyRules:getPropertyRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPropertyRules.
type LookupPropertyRulesArgs struct {
	Rules     []GetPropertyRulesRule `pulumi:"rules"`
	Variables *string                `pulumi:"variables"`
}

// A collection of values returned by getPropertyRules.
type LookupPropertyRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string                 `pulumi:"id"`
	Json      string                 `pulumi:"json"`
	Rules     []GetPropertyRulesRule `pulumi:"rules"`
	Variables *string                `pulumi:"variables"`
}
