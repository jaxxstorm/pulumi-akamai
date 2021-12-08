// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Contract; group
//
// Returns information about the contracts and groups associated with your account. Among other things, this information is required to create a new security configuration and to return a list of the hostnames available for use in a security policy. The returned information for this data source is described in the [List contracts and groups](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getcontractsandgroupswithksdorwaf) of the Application Security API.
//
// **Related API Endpoint**: [/appsec/v1/contracts-groups](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getcontractsandgroupswithksdorwaf)
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
// 		opt0 := "5-2WA382"
// 		opt1 := 12198
// 		contractsGroups, err := akamai.GetAppSecContractsGroups(ctx, &GetAppSecContractsGroupsArgs{
// 			Contractid: &opt0,
// 			Groupid:    &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("contractsGroupsList", contractsGroups.OutputText)
// 		ctx.Export("contractsGroupsJson", contractsGroups.Json)
// 		ctx.Export("contractGroupsDefaultContractid", contractsGroups.DefaultContractid)
// 		ctx.Export("contractGroupsDefaultGroupid", contractsGroups.DefaultGroupid)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `json`. JSON-formatted list of contract and group information.
// - `outputText`. Tabular report of contract and group information.
// - `defaultContractid`. Default contract ID for the specified contract and group.
// - `defaultGroupid`. Default group ID for the specified contract and group.
func GetAppSecContractsGroups(ctx *pulumi.Context, args *GetAppSecContractsGroupsArgs, opts ...pulumi.InvokeOption) (*GetAppSecContractsGroupsResult, error) {
	var rv GetAppSecContractsGroupsResult
	err := ctx.Invoke("akamai:index/getAppSecContractsGroups:getAppSecContractsGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecContractsGroups.
type GetAppSecContractsGroupsArgs struct {
	// . Unique identifier of an Akamai contract. If not included, information is returned for all the Akamai contracts associated with your account.
	Contractid *string `pulumi:"contractid"`
	// . Unique identifier of a contract group. If not included, information is returned for all the groups associated with your account.
	Groupid *int `pulumi:"groupid"`
}

// A collection of values returned by getAppSecContractsGroups.
type GetAppSecContractsGroupsResult struct {
	Contractid        *string `pulumi:"contractid"`
	DefaultContractid string  `pulumi:"defaultContractid"`
	DefaultGroupid    int     `pulumi:"defaultGroupid"`
	Groupid           *int    `pulumi:"groupid"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Json       string `pulumi:"json"`
	OutputText string `pulumi:"outputText"`
}

func GetAppSecContractsGroupsOutput(ctx *pulumi.Context, args GetAppSecContractsGroupsOutputArgs, opts ...pulumi.InvokeOption) GetAppSecContractsGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecContractsGroupsResult, error) {
			args := v.(GetAppSecContractsGroupsArgs)
			r, err := GetAppSecContractsGroups(ctx, &args, opts...)
			return *r, err
		}).(GetAppSecContractsGroupsResultOutput)
}

// A collection of arguments for invoking getAppSecContractsGroups.
type GetAppSecContractsGroupsOutputArgs struct {
	// . Unique identifier of an Akamai contract. If not included, information is returned for all the Akamai contracts associated with your account.
	Contractid pulumi.StringPtrInput `pulumi:"contractid"`
	// . Unique identifier of a contract group. If not included, information is returned for all the groups associated with your account.
	Groupid pulumi.IntPtrInput `pulumi:"groupid"`
}

func (GetAppSecContractsGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecContractsGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecContractsGroups.
type GetAppSecContractsGroupsResultOutput struct{ *pulumi.OutputState }

func (GetAppSecContractsGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecContractsGroupsResult)(nil)).Elem()
}

func (o GetAppSecContractsGroupsResultOutput) ToGetAppSecContractsGroupsResultOutput() GetAppSecContractsGroupsResultOutput {
	return o
}

func (o GetAppSecContractsGroupsResultOutput) ToGetAppSecContractsGroupsResultOutputWithContext(ctx context.Context) GetAppSecContractsGroupsResultOutput {
	return o
}

func (o GetAppSecContractsGroupsResultOutput) Contractid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppSecContractsGroupsResult) *string { return v.Contractid }).(pulumi.StringPtrOutput)
}

func (o GetAppSecContractsGroupsResultOutput) DefaultContractid() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecContractsGroupsResult) string { return v.DefaultContractid }).(pulumi.StringOutput)
}

func (o GetAppSecContractsGroupsResultOutput) DefaultGroupid() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecContractsGroupsResult) int { return v.DefaultGroupid }).(pulumi.IntOutput)
}

func (o GetAppSecContractsGroupsResultOutput) Groupid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecContractsGroupsResult) *int { return v.Groupid }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecContractsGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecContractsGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecContractsGroupsResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecContractsGroupsResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetAppSecContractsGroupsResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecContractsGroupsResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecContractsGroupsResultOutput{})
}
