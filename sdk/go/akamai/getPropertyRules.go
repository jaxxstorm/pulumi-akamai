// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getPropertyRules` data source to query and retrieve the rule tree of
// an existing property version. This data source lets you search across the contracts
// and groups you have access to.
//
// ## Basic usage
//
// This example returns the rule tree for version 3 of a property based on the selected contract and group:
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
// 		opt0 := "grp_12345"
// 		opt1 := "ctr_1-AB123"
// 		opt2 := 3
// 		my_example, err := akamai.GetPropertyRules(ctx, &GetPropertyRulesArgs{
// 			PropertyId: "prp_123",
// 			GroupId:    &opt0,
// 			ContractId: &opt1,
// 			Version:    &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("propertyMatch", my_example)
// 		return nil
// 	})
// }
// ```
//
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `ruleFormat` - The rule tree version used. Property rule objects are versioned infrequently, and are known as rule formats. See [About rule formats](https://developer.akamai.com/api/core_features/property_manager/vlatest.html#rf) to learn more.
// * `rules` - A JSON-encoded rule tree for the property.
// * `errors` - A list of validation errors for the rule tree object returned. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the Property Manager API documentation.
func GetPropertyRules(ctx *pulumi.Context, args *GetPropertyRulesArgs, opts ...pulumi.InvokeOption) (*GetPropertyRulesResult, error) {
	var rv GetPropertyRulesResult
	err := ctx.Invoke("akamai:index/getPropertyRules:getPropertyRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPropertyRules.
type GetPropertyRulesArgs struct {
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId *string `pulumi:"contractId"`
	// - (Required) A group's unique ID, including the `grp_` prefix.
	GroupId *string `pulumi:"groupId"`
	// - (Required) A property's unique ID, including the `prp_` prefix.
	PropertyId string  `pulumi:"propertyId"`
	RuleFormat *string `pulumi:"ruleFormat"`
	// - (Optional) The version to return. Returns the latest version by default.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getPropertyRules.
type GetPropertyRulesResult struct {
	ContractId string `pulumi:"contractId"`
	Errors     string `pulumi:"errors"`
	GroupId    string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	PropertyId string  `pulumi:"propertyId"`
	RuleFormat *string `pulumi:"ruleFormat"`
	Rules      string  `pulumi:"rules"`
	Version    int     `pulumi:"version"`
}

func GetPropertyRulesOutput(ctx *pulumi.Context, args GetPropertyRulesOutputArgs, opts ...pulumi.InvokeOption) GetPropertyRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPropertyRulesResult, error) {
			args := v.(GetPropertyRulesArgs)
			r, err := GetPropertyRules(ctx, &args, opts...)
			return *r, err
		}).(GetPropertyRulesResultOutput)
}

// A collection of arguments for invoking getPropertyRules.
type GetPropertyRulesOutputArgs struct {
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId pulumi.StringPtrInput `pulumi:"contractId"`
	// - (Required) A group's unique ID, including the `grp_` prefix.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// - (Required) A property's unique ID, including the `prp_` prefix.
	PropertyId pulumi.StringInput    `pulumi:"propertyId"`
	RuleFormat pulumi.StringPtrInput `pulumi:"ruleFormat"`
	// - (Optional) The version to return. Returns the latest version by default.
	Version pulumi.IntPtrInput `pulumi:"version"`
}

func (GetPropertyRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPropertyRulesArgs)(nil)).Elem()
}

// A collection of values returned by getPropertyRules.
type GetPropertyRulesResultOutput struct{ *pulumi.OutputState }

func (GetPropertyRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPropertyRulesResult)(nil)).Elem()
}

func (o GetPropertyRulesResultOutput) ToGetPropertyRulesResultOutput() GetPropertyRulesResultOutput {
	return o
}

func (o GetPropertyRulesResultOutput) ToGetPropertyRulesResultOutputWithContext(ctx context.Context) GetPropertyRulesResultOutput {
	return o
}

func (o GetPropertyRulesResultOutput) ContractId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPropertyRulesResult) string { return v.ContractId }).(pulumi.StringOutput)
}

func (o GetPropertyRulesResultOutput) Errors() pulumi.StringOutput {
	return o.ApplyT(func(v GetPropertyRulesResult) string { return v.Errors }).(pulumi.StringOutput)
}

func (o GetPropertyRulesResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPropertyRulesResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPropertyRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPropertyRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPropertyRulesResultOutput) PropertyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPropertyRulesResult) string { return v.PropertyId }).(pulumi.StringOutput)
}

func (o GetPropertyRulesResultOutput) RuleFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPropertyRulesResult) *string { return v.RuleFormat }).(pulumi.StringPtrOutput)
}

func (o GetPropertyRulesResultOutput) Rules() pulumi.StringOutput {
	return o.ApplyT(func(v GetPropertyRulesResult) string { return v.Rules }).(pulumi.StringOutput)
}

func (o GetPropertyRulesResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v GetPropertyRulesResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPropertyRulesResultOutput{})
}
