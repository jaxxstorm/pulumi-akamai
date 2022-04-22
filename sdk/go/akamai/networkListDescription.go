// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `NetworkListDescription` resource to update the name or description of an existing network list.
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
// 		_, err := akamai.NewNetworkListDescription(ctx, "networkListDescription", &akamai.NetworkListDescriptionArgs{
// 			NetworkListId: pulumi.Any(_var.Network_list_id),
// 			Description:   pulumi.String("Test network list updated description"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkListDescription struct {
	pulumi.CustomResourceState

	// The description to be assigned to the network list.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name to be assigned to the network list.
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique ID of the network list to use.
	NetworkListId pulumi.StringOutput `pulumi:"networkListId"`
}

// NewNetworkListDescription registers a new resource with the given unique name, arguments, and options.
func NewNetworkListDescription(ctx *pulumi.Context,
	name string, args *NetworkListDescriptionArgs, opts ...pulumi.ResourceOption) (*NetworkListDescription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.NetworkListId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkListId'")
	}
	var resource NetworkListDescription
	err := ctx.RegisterResource("akamai:index/networkListDescription:NetworkListDescription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkListDescription gets an existing NetworkListDescription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkListDescription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkListDescriptionState, opts ...pulumi.ResourceOption) (*NetworkListDescription, error) {
	var resource NetworkListDescription
	err := ctx.ReadResource("akamai:index/networkListDescription:NetworkListDescription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkListDescription resources.
type networkListDescriptionState struct {
	// The description to be assigned to the network list.
	Description *string `pulumi:"description"`
	// The name to be assigned to the network list.
	Name *string `pulumi:"name"`
	// The unique ID of the network list to use.
	NetworkListId *string `pulumi:"networkListId"`
}

type NetworkListDescriptionState struct {
	// The description to be assigned to the network list.
	Description pulumi.StringPtrInput
	// The name to be assigned to the network list.
	Name pulumi.StringPtrInput
	// The unique ID of the network list to use.
	NetworkListId pulumi.StringPtrInput
}

func (NetworkListDescriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkListDescriptionState)(nil)).Elem()
}

type networkListDescriptionArgs struct {
	// The description to be assigned to the network list.
	Description string `pulumi:"description"`
	// The name to be assigned to the network list.
	Name *string `pulumi:"name"`
	// The unique ID of the network list to use.
	NetworkListId string `pulumi:"networkListId"`
}

// The set of arguments for constructing a NetworkListDescription resource.
type NetworkListDescriptionArgs struct {
	// The description to be assigned to the network list.
	Description pulumi.StringInput
	// The name to be assigned to the network list.
	Name pulumi.StringPtrInput
	// The unique ID of the network list to use.
	NetworkListId pulumi.StringInput
}

func (NetworkListDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkListDescriptionArgs)(nil)).Elem()
}

type NetworkListDescriptionInput interface {
	pulumi.Input

	ToNetworkListDescriptionOutput() NetworkListDescriptionOutput
	ToNetworkListDescriptionOutputWithContext(ctx context.Context) NetworkListDescriptionOutput
}

func (*NetworkListDescription) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkListDescription)(nil)).Elem()
}

func (i *NetworkListDescription) ToNetworkListDescriptionOutput() NetworkListDescriptionOutput {
	return i.ToNetworkListDescriptionOutputWithContext(context.Background())
}

func (i *NetworkListDescription) ToNetworkListDescriptionOutputWithContext(ctx context.Context) NetworkListDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListDescriptionOutput)
}

// NetworkListDescriptionArrayInput is an input type that accepts NetworkListDescriptionArray and NetworkListDescriptionArrayOutput values.
// You can construct a concrete instance of `NetworkListDescriptionArrayInput` via:
//
//          NetworkListDescriptionArray{ NetworkListDescriptionArgs{...} }
type NetworkListDescriptionArrayInput interface {
	pulumi.Input

	ToNetworkListDescriptionArrayOutput() NetworkListDescriptionArrayOutput
	ToNetworkListDescriptionArrayOutputWithContext(context.Context) NetworkListDescriptionArrayOutput
}

type NetworkListDescriptionArray []NetworkListDescriptionInput

func (NetworkListDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkListDescription)(nil)).Elem()
}

func (i NetworkListDescriptionArray) ToNetworkListDescriptionArrayOutput() NetworkListDescriptionArrayOutput {
	return i.ToNetworkListDescriptionArrayOutputWithContext(context.Background())
}

func (i NetworkListDescriptionArray) ToNetworkListDescriptionArrayOutputWithContext(ctx context.Context) NetworkListDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListDescriptionArrayOutput)
}

// NetworkListDescriptionMapInput is an input type that accepts NetworkListDescriptionMap and NetworkListDescriptionMapOutput values.
// You can construct a concrete instance of `NetworkListDescriptionMapInput` via:
//
//          NetworkListDescriptionMap{ "key": NetworkListDescriptionArgs{...} }
type NetworkListDescriptionMapInput interface {
	pulumi.Input

	ToNetworkListDescriptionMapOutput() NetworkListDescriptionMapOutput
	ToNetworkListDescriptionMapOutputWithContext(context.Context) NetworkListDescriptionMapOutput
}

type NetworkListDescriptionMap map[string]NetworkListDescriptionInput

func (NetworkListDescriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkListDescription)(nil)).Elem()
}

func (i NetworkListDescriptionMap) ToNetworkListDescriptionMapOutput() NetworkListDescriptionMapOutput {
	return i.ToNetworkListDescriptionMapOutputWithContext(context.Background())
}

func (i NetworkListDescriptionMap) ToNetworkListDescriptionMapOutputWithContext(ctx context.Context) NetworkListDescriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListDescriptionMapOutput)
}

type NetworkListDescriptionOutput struct{ *pulumi.OutputState }

func (NetworkListDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkListDescription)(nil)).Elem()
}

func (o NetworkListDescriptionOutput) ToNetworkListDescriptionOutput() NetworkListDescriptionOutput {
	return o
}

func (o NetworkListDescriptionOutput) ToNetworkListDescriptionOutputWithContext(ctx context.Context) NetworkListDescriptionOutput {
	return o
}

type NetworkListDescriptionArrayOutput struct{ *pulumi.OutputState }

func (NetworkListDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkListDescription)(nil)).Elem()
}

func (o NetworkListDescriptionArrayOutput) ToNetworkListDescriptionArrayOutput() NetworkListDescriptionArrayOutput {
	return o
}

func (o NetworkListDescriptionArrayOutput) ToNetworkListDescriptionArrayOutputWithContext(ctx context.Context) NetworkListDescriptionArrayOutput {
	return o
}

func (o NetworkListDescriptionArrayOutput) Index(i pulumi.IntInput) NetworkListDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkListDescription {
		return vs[0].([]*NetworkListDescription)[vs[1].(int)]
	}).(NetworkListDescriptionOutput)
}

type NetworkListDescriptionMapOutput struct{ *pulumi.OutputState }

func (NetworkListDescriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkListDescription)(nil)).Elem()
}

func (o NetworkListDescriptionMapOutput) ToNetworkListDescriptionMapOutput() NetworkListDescriptionMapOutput {
	return o
}

func (o NetworkListDescriptionMapOutput) ToNetworkListDescriptionMapOutputWithContext(ctx context.Context) NetworkListDescriptionMapOutput {
	return o
}

func (o NetworkListDescriptionMapOutput) MapIndex(k pulumi.StringInput) NetworkListDescriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkListDescription {
		return vs[0].(map[string]*NetworkListDescription)[vs[1].(string)]
	}).(NetworkListDescriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListDescriptionInput)(nil)).Elem(), &NetworkListDescription{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListDescriptionArrayInput)(nil)).Elem(), NetworkListDescriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListDescriptionMapInput)(nil)).Elem(), NetworkListDescriptionMap{})
	pulumi.RegisterOutputType(NetworkListDescriptionOutput{})
	pulumi.RegisterOutputType(NetworkListDescriptionArrayOutput{})
	pulumi.RegisterOutputType(NetworkListDescriptionMapOutput{})
}
