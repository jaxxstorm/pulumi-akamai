// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getGroup` data source to get a group by name.
//
// Each account features a hierarchy of groups, which control access to your
// Akamai configurations and help consolidate reporting functions, typically
// mapping to an organizational hierarchy. Using either Control Center or the
// [Identity Management: User Administration API](https://developer.akamai.com/en-us/api/core_features/identity_management_user_admin/v2.html),
// account administrators can assign properties to specific groups, each with
// its own set of users and accompanying roles.
//
// ## Attributes reference
//
// This data source returns this attribute:
//
// * `id` - The group's unique ID, including the `grp_` prefix.
func GetGroup(ctx *pulumi.Context, args *GetGroupArgs, opts ...pulumi.InvokeOption) (*GetGroupResult, error) {
	var rv GetGroupResult
	err := ctx.Invoke("akamai:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type GetGroupArgs struct {
	// Replaced by `contractId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "contract" has been deprecated.
	Contract *string `pulumi:"contract"`
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId *string `pulumi:"contractId"`
	// The group name.
	GroupName *string `pulumi:"groupName"`
	// Replaced by `groupName`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "name" has been deprecated.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getGroup.
type GetGroupResult struct {
	// Deprecated: The setting "contract" has been deprecated.
	Contract   string `pulumi:"contract"`
	ContractId string `pulumi:"contractId"`
	GroupName  string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Deprecated: The setting "name" has been deprecated.
	Name string `pulumi:"name"`
}

func GetGroupOutput(ctx *pulumi.Context, args GetGroupOutputArgs, opts ...pulumi.InvokeOption) GetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupResult, error) {
			args := v.(GetGroupArgs)
			r, err := GetGroup(ctx, &args, opts...)
			var s GetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type GetGroupOutputArgs struct {
	// Replaced by `contractId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "contract" has been deprecated.
	Contract pulumi.StringPtrInput `pulumi:"contract"`
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId pulumi.StringPtrInput `pulumi:"contractId"`
	// The group name.
	GroupName pulumi.StringPtrInput `pulumi:"groupName"`
	// Replaced by `groupName`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "name" has been deprecated.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type GetGroupResultOutput struct{ *pulumi.OutputState }

func (GetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupResult)(nil)).Elem()
}

func (o GetGroupResultOutput) ToGetGroupResultOutput() GetGroupResultOutput {
	return o
}

func (o GetGroupResultOutput) ToGetGroupResultOutputWithContext(ctx context.Context) GetGroupResultOutput {
	return o
}

// Deprecated: The setting "contract" has been deprecated.
func (o GetGroupResultOutput) Contract() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupResult) string { return v.Contract }).(pulumi.StringOutput)
}

func (o GetGroupResultOutput) ContractId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupResult) string { return v.ContractId }).(pulumi.StringOutput)
}

func (o GetGroupResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupResult) string { return v.GroupName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Deprecated: The setting "name" has been deprecated.
func (o GetGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupResultOutput{})
}
