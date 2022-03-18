// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `NetworkList` resource to create a network list, or to modify an existing list.
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
// 		_, err := akamai.NewNetworkList(ctx, "networkList", &akamai.NetworkListArgs{
// 			Type:        pulumi.String("IP"),
// 			Description: pulumi.String("network list description"),
// 			Lists:       pulumi.Any(_var.List),
// 			Mode:        pulumi.String("APPEND"),
// 			ContractId:  pulumi.String("ABC-123"),
// 			GroupId:     pulumi.Int(12345),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkList struct {
	pulumi.CustomResourceState

	// The contract ID of the network list. If supplied, groupId must also be supplied. The
	// contractId value of an existing network list may not be modified.
	ContractId pulumi.StringPtrOutput `pulumi:"contractId"`
	// The description to be assigned to the network list.
	Description pulumi.StringOutput `pulumi:"description"`
	// The group ID of the network list. If supplied, contractId must also be supplied. The
	// groupId value of an existing network list may not be modified.
	GroupId pulumi.IntPtrOutput `pulumi:"groupId"`
	// : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
	// removed from an existing list.
	Lists pulumi.StringArrayOutput `pulumi:"lists"`
	// A string specifying the interpretation of the `list` parameter. Must be one of the following:
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The name to be assigned to the network list.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the network list.
	NetworkListId pulumi.StringOutput `pulumi:"networkListId"`
	// An integer that identifies the current version of the network list; this value is incremented each time
	// the list is modified.
	SyncPoint pulumi.IntOutput `pulumi:"syncPoint"`
	// The type of the network list; must be either "IP" or "GEO".
	Type pulumi.StringOutput `pulumi:"type"`
	// unique ID
	Uniqueid pulumi.StringOutput `pulumi:"uniqueid"`
}

// NewNetworkList registers a new resource with the given unique name, arguments, and options.
func NewNetworkList(ctx *pulumi.Context,
	name string, args *NetworkListArgs, opts ...pulumi.ResourceOption) (*NetworkList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource NetworkList
	err := ctx.RegisterResource("akamai:index/networkList:NetworkList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkList gets an existing NetworkList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkListState, opts ...pulumi.ResourceOption) (*NetworkList, error) {
	var resource NetworkList
	err := ctx.ReadResource("akamai:index/networkList:NetworkList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkList resources.
type networkListState struct {
	// The contract ID of the network list. If supplied, groupId must also be supplied. The
	// contractId value of an existing network list may not be modified.
	ContractId *string `pulumi:"contractId"`
	// The description to be assigned to the network list.
	Description *string `pulumi:"description"`
	// The group ID of the network list. If supplied, contractId must also be supplied. The
	// groupId value of an existing network list may not be modified.
	GroupId *int `pulumi:"groupId"`
	// : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
	// removed from an existing list.
	Lists []string `pulumi:"lists"`
	// A string specifying the interpretation of the `list` parameter. Must be one of the following:
	Mode *string `pulumi:"mode"`
	// The name to be assigned to the network list.
	Name *string `pulumi:"name"`
	// The ID of the network list.
	NetworkListId *string `pulumi:"networkListId"`
	// An integer that identifies the current version of the network list; this value is incremented each time
	// the list is modified.
	SyncPoint *int `pulumi:"syncPoint"`
	// The type of the network list; must be either "IP" or "GEO".
	Type *string `pulumi:"type"`
	// unique ID
	Uniqueid *string `pulumi:"uniqueid"`
}

type NetworkListState struct {
	// The contract ID of the network list. If supplied, groupId must also be supplied. The
	// contractId value of an existing network list may not be modified.
	ContractId pulumi.StringPtrInput
	// The description to be assigned to the network list.
	Description pulumi.StringPtrInput
	// The group ID of the network list. If supplied, contractId must also be supplied. The
	// groupId value of an existing network list may not be modified.
	GroupId pulumi.IntPtrInput
	// : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
	// removed from an existing list.
	Lists pulumi.StringArrayInput
	// A string specifying the interpretation of the `list` parameter. Must be one of the following:
	Mode pulumi.StringPtrInput
	// The name to be assigned to the network list.
	Name pulumi.StringPtrInput
	// The ID of the network list.
	NetworkListId pulumi.StringPtrInput
	// An integer that identifies the current version of the network list; this value is incremented each time
	// the list is modified.
	SyncPoint pulumi.IntPtrInput
	// The type of the network list; must be either "IP" or "GEO".
	Type pulumi.StringPtrInput
	// unique ID
	Uniqueid pulumi.StringPtrInput
}

func (NetworkListState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkListState)(nil)).Elem()
}

type networkListArgs struct {
	// The contract ID of the network list. If supplied, groupId must also be supplied. The
	// contractId value of an existing network list may not be modified.
	ContractId *string `pulumi:"contractId"`
	// The description to be assigned to the network list.
	Description string `pulumi:"description"`
	// The group ID of the network list. If supplied, contractId must also be supplied. The
	// groupId value of an existing network list may not be modified.
	GroupId *int `pulumi:"groupId"`
	// : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
	// removed from an existing list.
	Lists []string `pulumi:"lists"`
	// A string specifying the interpretation of the `list` parameter. Must be one of the following:
	Mode string `pulumi:"mode"`
	// The name to be assigned to the network list.
	Name *string `pulumi:"name"`
	// The type of the network list; must be either "IP" or "GEO".
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a NetworkList resource.
type NetworkListArgs struct {
	// The contract ID of the network list. If supplied, groupId must also be supplied. The
	// contractId value of an existing network list may not be modified.
	ContractId pulumi.StringPtrInput
	// The description to be assigned to the network list.
	Description pulumi.StringInput
	// The group ID of the network list. If supplied, contractId must also be supplied. The
	// groupId value of an existing network list may not be modified.
	GroupId pulumi.IntPtrInput
	// : (Optional) A list of IP addresses or locations to be included in the list, added to an existing list, or
	// removed from an existing list.
	Lists pulumi.StringArrayInput
	// A string specifying the interpretation of the `list` parameter. Must be one of the following:
	Mode pulumi.StringInput
	// The name to be assigned to the network list.
	Name pulumi.StringPtrInput
	// The type of the network list; must be either "IP" or "GEO".
	Type pulumi.StringInput
}

func (NetworkListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkListArgs)(nil)).Elem()
}

type NetworkListInput interface {
	pulumi.Input

	ToNetworkListOutput() NetworkListOutput
	ToNetworkListOutputWithContext(ctx context.Context) NetworkListOutput
}

func (*NetworkList) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkList)(nil)).Elem()
}

func (i *NetworkList) ToNetworkListOutput() NetworkListOutput {
	return i.ToNetworkListOutputWithContext(context.Background())
}

func (i *NetworkList) ToNetworkListOutputWithContext(ctx context.Context) NetworkListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListOutput)
}

// NetworkListArrayInput is an input type that accepts NetworkListArray and NetworkListArrayOutput values.
// You can construct a concrete instance of `NetworkListArrayInput` via:
//
//          NetworkListArray{ NetworkListArgs{...} }
type NetworkListArrayInput interface {
	pulumi.Input

	ToNetworkListArrayOutput() NetworkListArrayOutput
	ToNetworkListArrayOutputWithContext(context.Context) NetworkListArrayOutput
}

type NetworkListArray []NetworkListInput

func (NetworkListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkList)(nil)).Elem()
}

func (i NetworkListArray) ToNetworkListArrayOutput() NetworkListArrayOutput {
	return i.ToNetworkListArrayOutputWithContext(context.Background())
}

func (i NetworkListArray) ToNetworkListArrayOutputWithContext(ctx context.Context) NetworkListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListArrayOutput)
}

// NetworkListMapInput is an input type that accepts NetworkListMap and NetworkListMapOutput values.
// You can construct a concrete instance of `NetworkListMapInput` via:
//
//          NetworkListMap{ "key": NetworkListArgs{...} }
type NetworkListMapInput interface {
	pulumi.Input

	ToNetworkListMapOutput() NetworkListMapOutput
	ToNetworkListMapOutputWithContext(context.Context) NetworkListMapOutput
}

type NetworkListMap map[string]NetworkListInput

func (NetworkListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkList)(nil)).Elem()
}

func (i NetworkListMap) ToNetworkListMapOutput() NetworkListMapOutput {
	return i.ToNetworkListMapOutputWithContext(context.Background())
}

func (i NetworkListMap) ToNetworkListMapOutputWithContext(ctx context.Context) NetworkListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListMapOutput)
}

type NetworkListOutput struct{ *pulumi.OutputState }

func (NetworkListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkList)(nil)).Elem()
}

func (o NetworkListOutput) ToNetworkListOutput() NetworkListOutput {
	return o
}

func (o NetworkListOutput) ToNetworkListOutputWithContext(ctx context.Context) NetworkListOutput {
	return o
}

type NetworkListArrayOutput struct{ *pulumi.OutputState }

func (NetworkListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkList)(nil)).Elem()
}

func (o NetworkListArrayOutput) ToNetworkListArrayOutput() NetworkListArrayOutput {
	return o
}

func (o NetworkListArrayOutput) ToNetworkListArrayOutputWithContext(ctx context.Context) NetworkListArrayOutput {
	return o
}

func (o NetworkListArrayOutput) Index(i pulumi.IntInput) NetworkListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkList {
		return vs[0].([]*NetworkList)[vs[1].(int)]
	}).(NetworkListOutput)
}

type NetworkListMapOutput struct{ *pulumi.OutputState }

func (NetworkListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkList)(nil)).Elem()
}

func (o NetworkListMapOutput) ToNetworkListMapOutput() NetworkListMapOutput {
	return o
}

func (o NetworkListMapOutput) ToNetworkListMapOutputWithContext(ctx context.Context) NetworkListMapOutput {
	return o
}

func (o NetworkListMapOutput) MapIndex(k pulumi.StringInput) NetworkListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkList {
		return vs[0].(map[string]*NetworkList)[vs[1].(string)]
	}).(NetworkListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListInput)(nil)).Elem(), &NetworkList{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListArrayInput)(nil)).Elem(), NetworkListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListMapInput)(nil)).Elem(), NetworkListMap{})
	pulumi.RegisterOutputType(NetworkListOutput{})
	pulumi.RegisterOutputType(NetworkListArrayOutput{})
	pulumi.RegisterOutputType(NetworkListMapOutput{})
}
