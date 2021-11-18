// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `CpCode` data source to retrieve the ID for a content provider (CP) code.
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
// 		opt0 := "ctr_1-AB123"
// 		opt1 := "grp_123"
// 		_, err := akamai.LookupCpCode(ctx, &GetCpCodeArgs{
// 			ContractId: &opt0,
// 			GroupId:    &opt1,
// 			Name:       "my cpcode name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Here's a real-world example that includes other data sources as dependencies:
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
// 		groupName := "example group name"
// 		cpcodeName := "My CP code Name"
// 		opt0 := groupName
// 		exampleContract, err := akamai.GetContract(ctx, &GetContractArgs{
// 			GroupName: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := groupName
// 		opt2 := exampleContract.Id
// 		exampleGroup, err := akamai.GetGroup(ctx, &GetGroupArgs{
// 			GroupName:  &opt1,
// 			ContractId: &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt3 := exampleGroup.Id
// 		opt4 := exampleContract.Id
// 		_, err = akamai.LookupCpCode(ctx, &GetCpCodeArgs{
// 			Name:       cpcodeName,
// 			GroupId:    &opt3,
// 			ContractId: &opt4,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Argument reference
//
// This data source supports these arguments:
//
// * `name` - (Required) The name of the CP code.
// * `groupId` - (Required) The group's unique ID, including the `grp_` prefix.
// * `contractId` - (Required) A contract's unique ID, including the `ctr_` prefix.
//
// ### Deprecated arguments
// * `contract` - (Deprecated) Replaced by `contractId`. Maintained for legacy purposes.
// * `group` - (Deprecated) Replaced by `groupId`. Maintained for legacy purposes.
//
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `id` - The ID of the CP code, including the `cpc_` prefix.
// * `productIds` - An array of product IDs associated with this CP code. Each ID returned includes the `prd_` prefix.
func LookupCpCode(ctx *pulumi.Context, args *LookupCpCodeArgs, opts ...pulumi.InvokeOption) (*LookupCpCodeResult, error) {
	var rv LookupCpCodeResult
	err := ctx.Invoke("akamai:index/getCpCode:getCpCode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCpCode.
type LookupCpCodeArgs struct {
	// Deprecated: The setting "contract" has been deprecated.
	Contract   *string `pulumi:"contract"`
	ContractId *string `pulumi:"contractId"`
	// Deprecated: The setting "group" has been deprecated.
	Group   *string `pulumi:"group"`
	GroupId *string `pulumi:"groupId"`
	Name    string  `pulumi:"name"`
}

// A collection of values returned by getCpCode.
type LookupCpCodeResult struct {
	// Deprecated: The setting "contract" has been deprecated.
	Contract   string `pulumi:"contract"`
	ContractId string `pulumi:"contractId"`
	// Deprecated: The setting "group" has been deprecated.
	Group   string `pulumi:"group"`
	GroupId string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Name       string   `pulumi:"name"`
	ProductIds []string `pulumi:"productIds"`
}

func LookupCpCodeOutput(ctx *pulumi.Context, args LookupCpCodeOutputArgs, opts ...pulumi.InvokeOption) LookupCpCodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCpCodeResult, error) {
			args := v.(LookupCpCodeArgs)
			r, err := LookupCpCode(ctx, &args, opts...)
			return *r, err
		}).(LookupCpCodeResultOutput)
}

// A collection of arguments for invoking getCpCode.
type LookupCpCodeOutputArgs struct {
	// Deprecated: The setting "contract" has been deprecated.
	Contract   pulumi.StringPtrInput `pulumi:"contract"`
	ContractId pulumi.StringPtrInput `pulumi:"contractId"`
	// Deprecated: The setting "group" has been deprecated.
	Group   pulumi.StringPtrInput `pulumi:"group"`
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	Name    pulumi.StringInput    `pulumi:"name"`
}

func (LookupCpCodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCpCodeArgs)(nil)).Elem()
}

// A collection of values returned by getCpCode.
type LookupCpCodeResultOutput struct{ *pulumi.OutputState }

func (LookupCpCodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCpCodeResult)(nil)).Elem()
}

func (o LookupCpCodeResultOutput) ToLookupCpCodeResultOutput() LookupCpCodeResultOutput {
	return o
}

func (o LookupCpCodeResultOutput) ToLookupCpCodeResultOutputWithContext(ctx context.Context) LookupCpCodeResultOutput {
	return o
}

// Deprecated: The setting "contract" has been deprecated.
func (o LookupCpCodeResultOutput) Contract() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCpCodeResult) string { return v.Contract }).(pulumi.StringOutput)
}

func (o LookupCpCodeResultOutput) ContractId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCpCodeResult) string { return v.ContractId }).(pulumi.StringOutput)
}

// Deprecated: The setting "group" has been deprecated.
func (o LookupCpCodeResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCpCodeResult) string { return v.Group }).(pulumi.StringOutput)
}

func (o LookupCpCodeResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCpCodeResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCpCodeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCpCodeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCpCodeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCpCodeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCpCodeResultOutput) ProductIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCpCodeResult) []string { return v.ProductIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCpCodeResultOutput{})
}
