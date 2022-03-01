// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `NetworkListActivations` resource to activate a network list in either the STAGING or PRODUCTION
// environment.
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
// 		networkListsFilter, err := akamai.GetNetworkLists(ctx, &GetNetworkListsArgs{
// 			Name: pulumi.StringRef(_var.Network_list),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewNetworkListActivations(ctx, "activation", &akamai.NetworkListActivationsArgs{
// 			NetworkListId: pulumi.String(networkListsFilter.Lists[0]),
// 			Network:       pulumi.String("STAGING"),
// 			Notes:         pulumi.String("TEST Notes"),
// 			NotificationEmails: pulumi.StringArray{
// 				pulumi.String("user@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkListActivations struct {
	pulumi.CustomResourceState

	Activate pulumi.BoolPtrOutput `pulumi:"activate"`
	// The network to be used, either `STAGING` or `PRODUCTION`. If not supplied, defaults to
	// `STAGING`.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// The ID of the network list to be activated
	NetworkListId pulumi.StringOutput `pulumi:"networkListId"`
	// A comment describing the activation.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// A bracketed, comma-separated list of email addresses that will be notified when the
	// operation is complete.
	NotificationEmails pulumi.StringArrayOutput `pulumi:"notificationEmails"`
	// The string `ACTIVATED` if the activation was successful, or a string identifying the reason why the network
	// list was not activated.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewNetworkListActivations registers a new resource with the given unique name, arguments, and options.
func NewNetworkListActivations(ctx *pulumi.Context,
	name string, args *NetworkListActivationsArgs, opts ...pulumi.ResourceOption) (*NetworkListActivations, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkListId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkListId'")
	}
	if args.NotificationEmails == nil {
		return nil, errors.New("invalid value for required argument 'NotificationEmails'")
	}
	var resource NetworkListActivations
	err := ctx.RegisterResource("akamai:index/networkListActivations:NetworkListActivations", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkListActivations gets an existing NetworkListActivations resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkListActivations(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkListActivationsState, opts ...pulumi.ResourceOption) (*NetworkListActivations, error) {
	var resource NetworkListActivations
	err := ctx.ReadResource("akamai:index/networkListActivations:NetworkListActivations", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkListActivations resources.
type networkListActivationsState struct {
	Activate *bool `pulumi:"activate"`
	// The network to be used, either `STAGING` or `PRODUCTION`. If not supplied, defaults to
	// `STAGING`.
	Network *string `pulumi:"network"`
	// The ID of the network list to be activated
	NetworkListId *string `pulumi:"networkListId"`
	// A comment describing the activation.
	Notes *string `pulumi:"notes"`
	// A bracketed, comma-separated list of email addresses that will be notified when the
	// operation is complete.
	NotificationEmails []string `pulumi:"notificationEmails"`
	// The string `ACTIVATED` if the activation was successful, or a string identifying the reason why the network
	// list was not activated.
	Status *string `pulumi:"status"`
}

type NetworkListActivationsState struct {
	Activate pulumi.BoolPtrInput
	// The network to be used, either `STAGING` or `PRODUCTION`. If not supplied, defaults to
	// `STAGING`.
	Network pulumi.StringPtrInput
	// The ID of the network list to be activated
	NetworkListId pulumi.StringPtrInput
	// A comment describing the activation.
	Notes pulumi.StringPtrInput
	// A bracketed, comma-separated list of email addresses that will be notified when the
	// operation is complete.
	NotificationEmails pulumi.StringArrayInput
	// The string `ACTIVATED` if the activation was successful, or a string identifying the reason why the network
	// list was not activated.
	Status pulumi.StringPtrInput
}

func (NetworkListActivationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkListActivationsState)(nil)).Elem()
}

type networkListActivationsArgs struct {
	Activate *bool `pulumi:"activate"`
	// The network to be used, either `STAGING` or `PRODUCTION`. If not supplied, defaults to
	// `STAGING`.
	Network *string `pulumi:"network"`
	// The ID of the network list to be activated
	NetworkListId string `pulumi:"networkListId"`
	// A comment describing the activation.
	Notes *string `pulumi:"notes"`
	// A bracketed, comma-separated list of email addresses that will be notified when the
	// operation is complete.
	NotificationEmails []string `pulumi:"notificationEmails"`
}

// The set of arguments for constructing a NetworkListActivations resource.
type NetworkListActivationsArgs struct {
	Activate pulumi.BoolPtrInput
	// The network to be used, either `STAGING` or `PRODUCTION`. If not supplied, defaults to
	// `STAGING`.
	Network pulumi.StringPtrInput
	// The ID of the network list to be activated
	NetworkListId pulumi.StringInput
	// A comment describing the activation.
	Notes pulumi.StringPtrInput
	// A bracketed, comma-separated list of email addresses that will be notified when the
	// operation is complete.
	NotificationEmails pulumi.StringArrayInput
}

func (NetworkListActivationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkListActivationsArgs)(nil)).Elem()
}

type NetworkListActivationsInput interface {
	pulumi.Input

	ToNetworkListActivationsOutput() NetworkListActivationsOutput
	ToNetworkListActivationsOutputWithContext(ctx context.Context) NetworkListActivationsOutput
}

func (*NetworkListActivations) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkListActivations)(nil)).Elem()
}

func (i *NetworkListActivations) ToNetworkListActivationsOutput() NetworkListActivationsOutput {
	return i.ToNetworkListActivationsOutputWithContext(context.Background())
}

func (i *NetworkListActivations) ToNetworkListActivationsOutputWithContext(ctx context.Context) NetworkListActivationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListActivationsOutput)
}

// NetworkListActivationsArrayInput is an input type that accepts NetworkListActivationsArray and NetworkListActivationsArrayOutput values.
// You can construct a concrete instance of `NetworkListActivationsArrayInput` via:
//
//          NetworkListActivationsArray{ NetworkListActivationsArgs{...} }
type NetworkListActivationsArrayInput interface {
	pulumi.Input

	ToNetworkListActivationsArrayOutput() NetworkListActivationsArrayOutput
	ToNetworkListActivationsArrayOutputWithContext(context.Context) NetworkListActivationsArrayOutput
}

type NetworkListActivationsArray []NetworkListActivationsInput

func (NetworkListActivationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkListActivations)(nil)).Elem()
}

func (i NetworkListActivationsArray) ToNetworkListActivationsArrayOutput() NetworkListActivationsArrayOutput {
	return i.ToNetworkListActivationsArrayOutputWithContext(context.Background())
}

func (i NetworkListActivationsArray) ToNetworkListActivationsArrayOutputWithContext(ctx context.Context) NetworkListActivationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListActivationsArrayOutput)
}

// NetworkListActivationsMapInput is an input type that accepts NetworkListActivationsMap and NetworkListActivationsMapOutput values.
// You can construct a concrete instance of `NetworkListActivationsMapInput` via:
//
//          NetworkListActivationsMap{ "key": NetworkListActivationsArgs{...} }
type NetworkListActivationsMapInput interface {
	pulumi.Input

	ToNetworkListActivationsMapOutput() NetworkListActivationsMapOutput
	ToNetworkListActivationsMapOutputWithContext(context.Context) NetworkListActivationsMapOutput
}

type NetworkListActivationsMap map[string]NetworkListActivationsInput

func (NetworkListActivationsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkListActivations)(nil)).Elem()
}

func (i NetworkListActivationsMap) ToNetworkListActivationsMapOutput() NetworkListActivationsMapOutput {
	return i.ToNetworkListActivationsMapOutputWithContext(context.Background())
}

func (i NetworkListActivationsMap) ToNetworkListActivationsMapOutputWithContext(ctx context.Context) NetworkListActivationsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkListActivationsMapOutput)
}

type NetworkListActivationsOutput struct{ *pulumi.OutputState }

func (NetworkListActivationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkListActivations)(nil)).Elem()
}

func (o NetworkListActivationsOutput) ToNetworkListActivationsOutput() NetworkListActivationsOutput {
	return o
}

func (o NetworkListActivationsOutput) ToNetworkListActivationsOutputWithContext(ctx context.Context) NetworkListActivationsOutput {
	return o
}

type NetworkListActivationsArrayOutput struct{ *pulumi.OutputState }

func (NetworkListActivationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkListActivations)(nil)).Elem()
}

func (o NetworkListActivationsArrayOutput) ToNetworkListActivationsArrayOutput() NetworkListActivationsArrayOutput {
	return o
}

func (o NetworkListActivationsArrayOutput) ToNetworkListActivationsArrayOutputWithContext(ctx context.Context) NetworkListActivationsArrayOutput {
	return o
}

func (o NetworkListActivationsArrayOutput) Index(i pulumi.IntInput) NetworkListActivationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkListActivations {
		return vs[0].([]*NetworkListActivations)[vs[1].(int)]
	}).(NetworkListActivationsOutput)
}

type NetworkListActivationsMapOutput struct{ *pulumi.OutputState }

func (NetworkListActivationsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkListActivations)(nil)).Elem()
}

func (o NetworkListActivationsMapOutput) ToNetworkListActivationsMapOutput() NetworkListActivationsMapOutput {
	return o
}

func (o NetworkListActivationsMapOutput) ToNetworkListActivationsMapOutputWithContext(ctx context.Context) NetworkListActivationsMapOutput {
	return o
}

func (o NetworkListActivationsMapOutput) MapIndex(k pulumi.StringInput) NetworkListActivationsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkListActivations {
		return vs[0].(map[string]*NetworkListActivations)[vs[1].(string)]
	}).(NetworkListActivationsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListActivationsInput)(nil)).Elem(), &NetworkListActivations{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListActivationsArrayInput)(nil)).Elem(), NetworkListActivationsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkListActivationsMapInput)(nil)).Elem(), NetworkListActivationsMap{})
	pulumi.RegisterOutputType(NetworkListActivationsOutput{})
	pulumi.RegisterOutputType(NetworkListActivationsArrayOutput{})
	pulumi.RegisterOutputType(NetworkListActivationsMapOutput{})
}
