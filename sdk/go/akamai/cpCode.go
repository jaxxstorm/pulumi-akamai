// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `CpCode` resource lets you create or reuse content provider (CP) codes.  CP codes track web traffic handled by Akamai servers. Akamai gives you a CP code when you purchase a product. You need this code when you activate associated properties.
//
// You can create additional CP codes to support more detailed billing and reporting functions.
//
// By default, the Akamai Provider uses your existing CP code instead of creating a new one.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := akamai.NewCpCode(ctx, "cpCode", &akamai.CpCodeArgs{
//				ContractId: pulumi.Any(akamai_contract.Contract.Id),
//				GroupId:    pulumi.Any(akamai_group.Group.Id),
//				ProductId:  pulumi.String("prd_Object_Delivery"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Here's a real-life example that includes other data sources as dependencies:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			groupName := "example group name"
//			_ := "My CP Code"
//			exampleContract, err := akamai.GetContract(ctx, &GetContractArgs{
//				GroupName: pulumi.StringRef(groupName),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleGroup, err := akamai.GetGroup(ctx, &GetGroupArgs{
//				Name:       pulumi.StringRef(groupName),
//				ContractId: pulumi.StringRef(exampleContract.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = akamai.NewCpCode(ctx, "exampleCp", &akamai.CpCodeArgs{
//				GroupId:    pulumi.String(exampleGroup.Id),
//				ContractId: pulumi.String(exampleContract.Id),
//				ProductId:  pulumi.String("prd_Object_Delivery"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Attributes reference
//
// * `id` - The ID of the CP code.
//
// ## Import
//
// Basic Usagehcl resource "akamai_cp_code" "example" {
//
// # (resource arguments)
//
//	} You can import your Akamai CP codes using a comma-delimited string of the CP code, contract, and group IDs. You have to enter the IDs in this order`cpcode_id,contract_id,group_id` For example
//
// ```sh
//
//	$ pulumi import akamai:index/cpCode:CpCode example cpc_123,ctr_1-AB123,grp_123
//
// ```
type CpCode struct {
	pulumi.CustomResourceState

	// Replaced by `contractId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "contract" has been deprecated.
	Contract pulumi.StringOutput `pulumi:"contract"`
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId pulumi.StringOutput `pulumi:"contractId"`
	// Replaced by `groupId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "group" has been deprecated.
	Group pulumi.StringOutput `pulumi:"group"`
	// - (Required) A group's unique ID, including the `grp_` prefix.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
	Name pulumi.StringOutput `pulumi:"name"`
	// Replaced by `productId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "product" has been deprecated.
	Product   pulumi.StringOutput `pulumi:"product"`
	ProductId pulumi.StringOutput `pulumi:"productId"`
}

// NewCpCode registers a new resource with the given unique name, arguments, and options.
func NewCpCode(ctx *pulumi.Context,
	name string, args *CpCodeArgs, opts ...pulumi.ResourceOption) (*CpCode, error) {
	if args == nil {
		args = &CpCodeArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("akamai:properties/cpCode:CpCode"),
		},
	})
	opts = append(opts, aliases)
	var resource CpCode
	err := ctx.RegisterResource("akamai:index/cpCode:CpCode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCpCode gets an existing CpCode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCpCode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CpCodeState, opts ...pulumi.ResourceOption) (*CpCode, error) {
	var resource CpCode
	err := ctx.ReadResource("akamai:index/cpCode:CpCode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CpCode resources.
type cpCodeState struct {
	// Replaced by `contractId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "contract" has been deprecated.
	Contract *string `pulumi:"contract"`
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId *string `pulumi:"contractId"`
	// Replaced by `groupId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "group" has been deprecated.
	Group *string `pulumi:"group"`
	// - (Required) A group's unique ID, including the `grp_` prefix.
	GroupId *string `pulumi:"groupId"`
	// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
	Name *string `pulumi:"name"`
	// Replaced by `productId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "product" has been deprecated.
	Product   *string `pulumi:"product"`
	ProductId *string `pulumi:"productId"`
}

type CpCodeState struct {
	// Replaced by `contractId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "contract" has been deprecated.
	Contract pulumi.StringPtrInput
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId pulumi.StringPtrInput
	// Replaced by `groupId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "group" has been deprecated.
	Group pulumi.StringPtrInput
	// - (Required) A group's unique ID, including the `grp_` prefix.
	GroupId pulumi.StringPtrInput
	// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
	Name pulumi.StringPtrInput
	// Replaced by `productId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "product" has been deprecated.
	Product   pulumi.StringPtrInput
	ProductId pulumi.StringPtrInput
}

func (CpCodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*cpCodeState)(nil)).Elem()
}

type cpCodeArgs struct {
	// Replaced by `contractId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "contract" has been deprecated.
	Contract *string `pulumi:"contract"`
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId *string `pulumi:"contractId"`
	// Replaced by `groupId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "group" has been deprecated.
	Group *string `pulumi:"group"`
	// - (Required) A group's unique ID, including the `grp_` prefix.
	GroupId *string `pulumi:"groupId"`
	// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
	Name *string `pulumi:"name"`
	// Replaced by `productId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "product" has been deprecated.
	Product   *string `pulumi:"product"`
	ProductId *string `pulumi:"productId"`
}

// The set of arguments for constructing a CpCode resource.
type CpCodeArgs struct {
	// Replaced by `contractId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "contract" has been deprecated.
	Contract pulumi.StringPtrInput
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId pulumi.StringPtrInput
	// Replaced by `groupId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "group" has been deprecated.
	Group pulumi.StringPtrInput
	// - (Required) A group's unique ID, including the `grp_` prefix.
	GroupId pulumi.StringPtrInput
	// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
	Name pulumi.StringPtrInput
	// Replaced by `productId`. Maintained for legacy purposes.
	//
	// Deprecated: The setting "product" has been deprecated.
	Product   pulumi.StringPtrInput
	ProductId pulumi.StringPtrInput
}

func (CpCodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cpCodeArgs)(nil)).Elem()
}

type CpCodeInput interface {
	pulumi.Input

	ToCpCodeOutput() CpCodeOutput
	ToCpCodeOutputWithContext(ctx context.Context) CpCodeOutput
}

func (*CpCode) ElementType() reflect.Type {
	return reflect.TypeOf((**CpCode)(nil)).Elem()
}

func (i *CpCode) ToCpCodeOutput() CpCodeOutput {
	return i.ToCpCodeOutputWithContext(context.Background())
}

func (i *CpCode) ToCpCodeOutputWithContext(ctx context.Context) CpCodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpCodeOutput)
}

// CpCodeArrayInput is an input type that accepts CpCodeArray and CpCodeArrayOutput values.
// You can construct a concrete instance of `CpCodeArrayInput` via:
//
//	CpCodeArray{ CpCodeArgs{...} }
type CpCodeArrayInput interface {
	pulumi.Input

	ToCpCodeArrayOutput() CpCodeArrayOutput
	ToCpCodeArrayOutputWithContext(context.Context) CpCodeArrayOutput
}

type CpCodeArray []CpCodeInput

func (CpCodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CpCode)(nil)).Elem()
}

func (i CpCodeArray) ToCpCodeArrayOutput() CpCodeArrayOutput {
	return i.ToCpCodeArrayOutputWithContext(context.Background())
}

func (i CpCodeArray) ToCpCodeArrayOutputWithContext(ctx context.Context) CpCodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpCodeArrayOutput)
}

// CpCodeMapInput is an input type that accepts CpCodeMap and CpCodeMapOutput values.
// You can construct a concrete instance of `CpCodeMapInput` via:
//
//	CpCodeMap{ "key": CpCodeArgs{...} }
type CpCodeMapInput interface {
	pulumi.Input

	ToCpCodeMapOutput() CpCodeMapOutput
	ToCpCodeMapOutputWithContext(context.Context) CpCodeMapOutput
}

type CpCodeMap map[string]CpCodeInput

func (CpCodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CpCode)(nil)).Elem()
}

func (i CpCodeMap) ToCpCodeMapOutput() CpCodeMapOutput {
	return i.ToCpCodeMapOutputWithContext(context.Background())
}

func (i CpCodeMap) ToCpCodeMapOutputWithContext(ctx context.Context) CpCodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpCodeMapOutput)
}

type CpCodeOutput struct{ *pulumi.OutputState }

func (CpCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CpCode)(nil)).Elem()
}

func (o CpCodeOutput) ToCpCodeOutput() CpCodeOutput {
	return o
}

func (o CpCodeOutput) ToCpCodeOutputWithContext(ctx context.Context) CpCodeOutput {
	return o
}

// Replaced by `contractId`. Maintained for legacy purposes.
//
// Deprecated: The setting "contract" has been deprecated.
func (o CpCodeOutput) Contract() pulumi.StringOutput {
	return o.ApplyT(func(v *CpCode) pulumi.StringOutput { return v.Contract }).(pulumi.StringOutput)
}

// - (Required) A contract's unique ID, including the `ctr_` prefix.
func (o CpCodeOutput) ContractId() pulumi.StringOutput {
	return o.ApplyT(func(v *CpCode) pulumi.StringOutput { return v.ContractId }).(pulumi.StringOutput)
}

// Replaced by `groupId`. Maintained for legacy purposes.
//
// Deprecated: The setting "group" has been deprecated.
func (o CpCodeOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *CpCode) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// - (Required) A group's unique ID, including the `grp_` prefix.
func (o CpCodeOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *CpCode) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
func (o CpCodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CpCode) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Replaced by `productId`. Maintained for legacy purposes.
//
// Deprecated: The setting "product" has been deprecated.
func (o CpCodeOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v *CpCode) pulumi.StringOutput { return v.Product }).(pulumi.StringOutput)
}

func (o CpCodeOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *CpCode) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

type CpCodeArrayOutput struct{ *pulumi.OutputState }

func (CpCodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CpCode)(nil)).Elem()
}

func (o CpCodeArrayOutput) ToCpCodeArrayOutput() CpCodeArrayOutput {
	return o
}

func (o CpCodeArrayOutput) ToCpCodeArrayOutputWithContext(ctx context.Context) CpCodeArrayOutput {
	return o
}

func (o CpCodeArrayOutput) Index(i pulumi.IntInput) CpCodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CpCode {
		return vs[0].([]*CpCode)[vs[1].(int)]
	}).(CpCodeOutput)
}

type CpCodeMapOutput struct{ *pulumi.OutputState }

func (CpCodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CpCode)(nil)).Elem()
}

func (o CpCodeMapOutput) ToCpCodeMapOutput() CpCodeMapOutput {
	return o
}

func (o CpCodeMapOutput) ToCpCodeMapOutputWithContext(ctx context.Context) CpCodeMapOutput {
	return o
}

func (o CpCodeMapOutput) MapIndex(k pulumi.StringInput) CpCodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CpCode {
		return vs[0].(map[string]*CpCode)[vs[1].(string)]
	}).(CpCodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CpCodeInput)(nil)).Elem(), &CpCode{})
	pulumi.RegisterInputType(reflect.TypeOf((*CpCodeArrayInput)(nil)).Elem(), CpCodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CpCodeMapInput)(nil)).Elem(), CpCodeMap{})
	pulumi.RegisterOutputType(CpCodeOutput{})
	pulumi.RegisterOutputType(CpCodeArrayOutput{})
	pulumi.RegisterOutputType(CpCodeMapOutput{})
}
