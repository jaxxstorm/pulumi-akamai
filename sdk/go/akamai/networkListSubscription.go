// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `NetworkListSubscription` resource to specify a set of email addresses to be notified of changes to any
// of a set of network lists.
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
// 		opt0 := _var.Network_list
// 		networkListsFilter, err := akamai.GetNetworkLists(ctx, &akamai.GetNetworkListsArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewNetworkListSubscription(ctx, "subscribe", &akamai.NetworkListSubscriptionArgs{
// 			NetworkLists: toPulumiStringArray(networkListsFilter.Lists),
// 			Recipients: pulumi.StringArray{
// 				pulumi.String("user@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// func toPulumiStringArray(arr []string) pulumi.StringArray {
// 	var pulumiArr pulumi.StringArray
// 	for _, v := range arr {
// 		pulumiArr = append(pulumiArr, pulumi.String(v))
// 	}
// 	return pulumiArr
// }
// ```
type NetworkListSubscription struct {
	pulumi.CustomResourceState

	// A list containing one or more IDs of the network lists to which the indicated email
	// addresses should be subscribed.
	NetworkLists pulumi.StringArrayOutput `pulumi:"networkLists"`
	// A bracketed, comma-separated list of email addresses that will be notified of changes to any
	// of the specified network lists.
	Recipients pulumi.StringArrayOutput `pulumi:"recipients"`
}

// NewNetworkListSubscription registers a new resource with the given unique name, arguments, and options.
func NewNetworkListSubscription(ctx *pulumi.Context,
	name string, args *NetworkListSubscriptionArgs, opts ...pulumi.ResourceOption) (*NetworkListSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkLists == nil {
		return nil, errors.New("invalid value for required argument 'NetworkLists'")
	}
	if args.Recipients == nil {
		return nil, errors.New("invalid value for required argument 'Recipients'")
	}
	var resource NetworkListSubscription
	err := ctx.RegisterResource("akamai:index/networkListSubscription:NetworkListSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkListSubscription gets an existing NetworkListSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkListSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkListSubscriptionState, opts ...pulumi.ResourceOption) (*NetworkListSubscription, error) {
	var resource NetworkListSubscription
	err := ctx.ReadResource("akamai:index/networkListSubscription:NetworkListSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkListSubscription resources.
type networkListSubscriptionState struct {
	// A list containing one or more IDs of the network lists to which the indicated email
	// addresses should be subscribed.
	NetworkLists []string `pulumi:"networkLists"`
	// A bracketed, comma-separated list of email addresses that will be notified of changes to any
	// of the specified network lists.
	Recipients []string `pulumi:"recipients"`
}

type NetworkListSubscriptionState struct {
	// A list containing one or more IDs of the network lists to which the indicated email
	// addresses should be subscribed.
	NetworkLists pulumi.StringArrayInput
	// A bracketed, comma-separated list of email addresses that will be notified of changes to any
	// of the specified network lists.
	Recipients pulumi.StringArrayInput
}

func (NetworkListSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkListSubscriptionState)(nil)).Elem()
}

type networkListSubscriptionArgs struct {
	// A list containing one or more IDs of the network lists to which the indicated email
	// addresses should be subscribed.
	NetworkLists []string `pulumi:"networkLists"`
	// A bracketed, comma-separated list of email addresses that will be notified of changes to any
	// of the specified network lists.
	Recipients []string `pulumi:"recipients"`
}

// The set of arguments for constructing a NetworkListSubscription resource.
type NetworkListSubscriptionArgs struct {
	// A list containing one or more IDs of the network lists to which the indicated email
	// addresses should be subscribed.
	NetworkLists pulumi.StringArrayInput
	// A bracketed, comma-separated list of email addresses that will be notified of changes to any
	// of the specified network lists.
	Recipients pulumi.StringArrayInput
}

func (NetworkListSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkListSubscriptionArgs)(nil)).Elem()
}

type NetworkListSubscriptionInput interface {
	pulumi.Input

	ToNetworkListSubscriptionOutput() NetworkListSubscriptionOutput
	ToNetworkListSubscriptionOutputWithContext(ctx context.Context) NetworkListSubscriptionOutput
}

func (*NetworkListSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkListSubscription)(nil))
}

func (i *NetworkListSubscription) ToNetworkListSubscriptionOutput() NetworkListSubscriptionOutput {
	return i.ToNetworkListSubscriptionOutputWithContext(context.Background())
}

func (i *NetworkListSubscription) ToNetworkListSubscriptionOutputWithContext(ctx context.Context) NetworkListSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListSubscriptionOutput)
}

func (i *NetworkListSubscription) ToNetworkListSubscriptionPtrOutput() NetworkListSubscriptionPtrOutput {
	return i.ToNetworkListSubscriptionPtrOutputWithContext(context.Background())
}

func (i *NetworkListSubscription) ToNetworkListSubscriptionPtrOutputWithContext(ctx context.Context) NetworkListSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListSubscriptionPtrOutput)
}

type NetworkListSubscriptionPtrInput interface {
	pulumi.Input

	ToNetworkListSubscriptionPtrOutput() NetworkListSubscriptionPtrOutput
	ToNetworkListSubscriptionPtrOutputWithContext(ctx context.Context) NetworkListSubscriptionPtrOutput
}

type networkListSubscriptionPtrType NetworkListSubscriptionArgs

func (*networkListSubscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkListSubscription)(nil))
}

func (i *networkListSubscriptionPtrType) ToNetworkListSubscriptionPtrOutput() NetworkListSubscriptionPtrOutput {
	return i.ToNetworkListSubscriptionPtrOutputWithContext(context.Background())
}

func (i *networkListSubscriptionPtrType) ToNetworkListSubscriptionPtrOutputWithContext(ctx context.Context) NetworkListSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListSubscriptionPtrOutput)
}

// NetworkListSubscriptionArrayInput is an input type that accepts NetworkListSubscriptionArray and NetworkListSubscriptionArrayOutput values.
// You can construct a concrete instance of `NetworkListSubscriptionArrayInput` via:
//
//          NetworkListSubscriptionArray{ NetworkListSubscriptionArgs{...} }
type NetworkListSubscriptionArrayInput interface {
	pulumi.Input

	ToNetworkListSubscriptionArrayOutput() NetworkListSubscriptionArrayOutput
	ToNetworkListSubscriptionArrayOutputWithContext(context.Context) NetworkListSubscriptionArrayOutput
}

type NetworkListSubscriptionArray []NetworkListSubscriptionInput

func (NetworkListSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NetworkListSubscription)(nil))
}

func (i NetworkListSubscriptionArray) ToNetworkListSubscriptionArrayOutput() NetworkListSubscriptionArrayOutput {
	return i.ToNetworkListSubscriptionArrayOutputWithContext(context.Background())
}

func (i NetworkListSubscriptionArray) ToNetworkListSubscriptionArrayOutputWithContext(ctx context.Context) NetworkListSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListSubscriptionArrayOutput)
}

// NetworkListSubscriptionMapInput is an input type that accepts NetworkListSubscriptionMap and NetworkListSubscriptionMapOutput values.
// You can construct a concrete instance of `NetworkListSubscriptionMapInput` via:
//
//          NetworkListSubscriptionMap{ "key": NetworkListSubscriptionArgs{...} }
type NetworkListSubscriptionMapInput interface {
	pulumi.Input

	ToNetworkListSubscriptionMapOutput() NetworkListSubscriptionMapOutput
	ToNetworkListSubscriptionMapOutputWithContext(context.Context) NetworkListSubscriptionMapOutput
}

type NetworkListSubscriptionMap map[string]NetworkListSubscriptionInput

func (NetworkListSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NetworkListSubscription)(nil))
}

func (i NetworkListSubscriptionMap) ToNetworkListSubscriptionMapOutput() NetworkListSubscriptionMapOutput {
	return i.ToNetworkListSubscriptionMapOutputWithContext(context.Background())
}

func (i NetworkListSubscriptionMap) ToNetworkListSubscriptionMapOutputWithContext(ctx context.Context) NetworkListSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListSubscriptionMapOutput)
}

type NetworkListSubscriptionOutput struct {
	*pulumi.OutputState
}

func (NetworkListSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkListSubscription)(nil))
}

func (o NetworkListSubscriptionOutput) ToNetworkListSubscriptionOutput() NetworkListSubscriptionOutput {
	return o
}

func (o NetworkListSubscriptionOutput) ToNetworkListSubscriptionOutputWithContext(ctx context.Context) NetworkListSubscriptionOutput {
	return o
}

func (o NetworkListSubscriptionOutput) ToNetworkListSubscriptionPtrOutput() NetworkListSubscriptionPtrOutput {
	return o.ToNetworkListSubscriptionPtrOutputWithContext(context.Background())
}

func (o NetworkListSubscriptionOutput) ToNetworkListSubscriptionPtrOutputWithContext(ctx context.Context) NetworkListSubscriptionPtrOutput {
	return o.ApplyT(func(v NetworkListSubscription) *NetworkListSubscription {
		return &v
	}).(NetworkListSubscriptionPtrOutput)
}

type NetworkListSubscriptionPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkListSubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkListSubscription)(nil))
}

func (o NetworkListSubscriptionPtrOutput) ToNetworkListSubscriptionPtrOutput() NetworkListSubscriptionPtrOutput {
	return o
}

func (o NetworkListSubscriptionPtrOutput) ToNetworkListSubscriptionPtrOutputWithContext(ctx context.Context) NetworkListSubscriptionPtrOutput {
	return o
}

type NetworkListSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (NetworkListSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkListSubscription)(nil))
}

func (o NetworkListSubscriptionArrayOutput) ToNetworkListSubscriptionArrayOutput() NetworkListSubscriptionArrayOutput {
	return o
}

func (o NetworkListSubscriptionArrayOutput) ToNetworkListSubscriptionArrayOutputWithContext(ctx context.Context) NetworkListSubscriptionArrayOutput {
	return o
}

func (o NetworkListSubscriptionArrayOutput) Index(i pulumi.IntInput) NetworkListSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkListSubscription {
		return vs[0].([]NetworkListSubscription)[vs[1].(int)]
	}).(NetworkListSubscriptionOutput)
}

type NetworkListSubscriptionMapOutput struct{ *pulumi.OutputState }

func (NetworkListSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkListSubscription)(nil))
}

func (o NetworkListSubscriptionMapOutput) ToNetworkListSubscriptionMapOutput() NetworkListSubscriptionMapOutput {
	return o
}

func (o NetworkListSubscriptionMapOutput) ToNetworkListSubscriptionMapOutputWithContext(ctx context.Context) NetworkListSubscriptionMapOutput {
	return o
}

func (o NetworkListSubscriptionMapOutput) MapIndex(k pulumi.StringInput) NetworkListSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkListSubscription {
		return vs[0].(map[string]NetworkListSubscription)[vs[1].(string)]
	}).(NetworkListSubscriptionOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkListSubscriptionOutput{})
	pulumi.RegisterOutputType(NetworkListSubscriptionPtrOutput{})
	pulumi.RegisterOutputType(NetworkListSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(NetworkListSubscriptionMapOutput{})
}