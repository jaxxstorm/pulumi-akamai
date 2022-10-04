// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `IamGroup` resource to list details about groups. Groups are organizational containers for the objects you use.  Groups can contain other groups, primary objects like properties, and secondary objects like edge hostnames or content provider (CP) codes.
//
// ## Basic usage
//
// This example returns the policy details based on the policy ID and optionally, a version:
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
//			_, err := akamai.NewIamGroup(ctx, "example", &akamai.IamGroupArgs{
//				ParentGroupId: pulumi.Int(12345),
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
// ## Attributes reference
//
// This resource returns this attribute:
//
// * `subGroups` - Sub-groups that are related to this group. Each identifier must be an integer.
type IamGroup struct {
	pulumi.CustomResourceState

	// Human readable name for a group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A unique identifier for the parent group. Each identifier must be an integer.
	ParentGroupId pulumi.IntOutput `pulumi:"parentGroupId"`
	// Subgroups IDs
	SubGroups pulumi.IntArrayOutput `pulumi:"subGroups"`
}

// NewIamGroup registers a new resource with the given unique name, arguments, and options.
func NewIamGroup(ctx *pulumi.Context,
	name string, args *IamGroupArgs, opts ...pulumi.ResourceOption) (*IamGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ParentGroupId'")
	}
	var resource IamGroup
	err := ctx.RegisterResource("akamai:index/iamGroup:IamGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamGroup gets an existing IamGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamGroupState, opts ...pulumi.ResourceOption) (*IamGroup, error) {
	var resource IamGroup
	err := ctx.ReadResource("akamai:index/iamGroup:IamGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamGroup resources.
type iamGroupState struct {
	// Human readable name for a group.
	Name *string `pulumi:"name"`
	// A unique identifier for the parent group. Each identifier must be an integer.
	ParentGroupId *int `pulumi:"parentGroupId"`
	// Subgroups IDs
	SubGroups []int `pulumi:"subGroups"`
}

type IamGroupState struct {
	// Human readable name for a group.
	Name pulumi.StringPtrInput
	// A unique identifier for the parent group. Each identifier must be an integer.
	ParentGroupId pulumi.IntPtrInput
	// Subgroups IDs
	SubGroups pulumi.IntArrayInput
}

func (IamGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupState)(nil)).Elem()
}

type iamGroupArgs struct {
	// Human readable name for a group.
	Name *string `pulumi:"name"`
	// A unique identifier for the parent group. Each identifier must be an integer.
	ParentGroupId int `pulumi:"parentGroupId"`
}

// The set of arguments for constructing a IamGroup resource.
type IamGroupArgs struct {
	// Human readable name for a group.
	Name pulumi.StringPtrInput
	// A unique identifier for the parent group. Each identifier must be an integer.
	ParentGroupId pulumi.IntInput
}

func (IamGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupArgs)(nil)).Elem()
}

type IamGroupInput interface {
	pulumi.Input

	ToIamGroupOutput() IamGroupOutput
	ToIamGroupOutputWithContext(ctx context.Context) IamGroupOutput
}

func (*IamGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroup)(nil)).Elem()
}

func (i *IamGroup) ToIamGroupOutput() IamGroupOutput {
	return i.ToIamGroupOutputWithContext(context.Background())
}

func (i *IamGroup) ToIamGroupOutputWithContext(ctx context.Context) IamGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupOutput)
}

// IamGroupArrayInput is an input type that accepts IamGroupArray and IamGroupArrayOutput values.
// You can construct a concrete instance of `IamGroupArrayInput` via:
//
//	IamGroupArray{ IamGroupArgs{...} }
type IamGroupArrayInput interface {
	pulumi.Input

	ToIamGroupArrayOutput() IamGroupArrayOutput
	ToIamGroupArrayOutputWithContext(context.Context) IamGroupArrayOutput
}

type IamGroupArray []IamGroupInput

func (IamGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroup)(nil)).Elem()
}

func (i IamGroupArray) ToIamGroupArrayOutput() IamGroupArrayOutput {
	return i.ToIamGroupArrayOutputWithContext(context.Background())
}

func (i IamGroupArray) ToIamGroupArrayOutputWithContext(ctx context.Context) IamGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupArrayOutput)
}

// IamGroupMapInput is an input type that accepts IamGroupMap and IamGroupMapOutput values.
// You can construct a concrete instance of `IamGroupMapInput` via:
//
//	IamGroupMap{ "key": IamGroupArgs{...} }
type IamGroupMapInput interface {
	pulumi.Input

	ToIamGroupMapOutput() IamGroupMapOutput
	ToIamGroupMapOutputWithContext(context.Context) IamGroupMapOutput
}

type IamGroupMap map[string]IamGroupInput

func (IamGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroup)(nil)).Elem()
}

func (i IamGroupMap) ToIamGroupMapOutput() IamGroupMapOutput {
	return i.ToIamGroupMapOutputWithContext(context.Background())
}

func (i IamGroupMap) ToIamGroupMapOutputWithContext(ctx context.Context) IamGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupMapOutput)
}

type IamGroupOutput struct{ *pulumi.OutputState }

func (IamGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroup)(nil)).Elem()
}

func (o IamGroupOutput) ToIamGroupOutput() IamGroupOutput {
	return o
}

func (o IamGroupOutput) ToIamGroupOutputWithContext(ctx context.Context) IamGroupOutput {
	return o
}

// Human readable name for a group.
func (o IamGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A unique identifier for the parent group. Each identifier must be an integer.
func (o IamGroupOutput) ParentGroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.IntOutput { return v.ParentGroupId }).(pulumi.IntOutput)
}

// Subgroups IDs
func (o IamGroupOutput) SubGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.IntArrayOutput { return v.SubGroups }).(pulumi.IntArrayOutput)
}

type IamGroupArrayOutput struct{ *pulumi.OutputState }

func (IamGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroup)(nil)).Elem()
}

func (o IamGroupArrayOutput) ToIamGroupArrayOutput() IamGroupArrayOutput {
	return o
}

func (o IamGroupArrayOutput) ToIamGroupArrayOutputWithContext(ctx context.Context) IamGroupArrayOutput {
	return o
}

func (o IamGroupArrayOutput) Index(i pulumi.IntInput) IamGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamGroup {
		return vs[0].([]*IamGroup)[vs[1].(int)]
	}).(IamGroupOutput)
}

type IamGroupMapOutput struct{ *pulumi.OutputState }

func (IamGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroup)(nil)).Elem()
}

func (o IamGroupMapOutput) ToIamGroupMapOutput() IamGroupMapOutput {
	return o
}

func (o IamGroupMapOutput) ToIamGroupMapOutputWithContext(ctx context.Context) IamGroupMapOutput {
	return o
}

func (o IamGroupMapOutput) MapIndex(k pulumi.StringInput) IamGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamGroup {
		return vs[0].(map[string]*IamGroup)[vs[1].(string)]
	}).(IamGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupInput)(nil)).Elem(), &IamGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupArrayInput)(nil)).Elem(), IamGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupMapInput)(nil)).Elem(), IamGroupMap{})
	pulumi.RegisterOutputType(IamGroupOutput{})
	pulumi.RegisterOutputType(IamGroupArrayOutput{})
	pulumi.RegisterOutputType(IamGroupMapOutput{})
}
