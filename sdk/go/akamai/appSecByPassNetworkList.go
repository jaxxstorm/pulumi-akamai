// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecByPassNetworkList` resource to update which network lists to use in the
// bypass network lists settings. Note: this resource is only applicable to WAP (Web Application
// Protector) configurations.
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
// 		opt0 := _var.Security_configuration
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &akamai.LookupAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecByPassNetworkList(ctx, "bypassNetworkLists", &akamai.AppSecByPassNetworkListArgs{
// 			ConfigId: pulumi.Int(configuration.ConfigId),
// 			BypassNetworkLists: pulumi.StringArray{
// 				pulumi.String("id1"),
// 				pulumi.String("id2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecByPassNetworkList struct {
	pulumi.CustomResourceState

	// A list containing the IDs of the network lists to use.
	BypassNetworkLists pulumi.StringArrayOutput `pulumi:"bypassNetworkLists"`
	// The configuration ID to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
}

// NewAppSecByPassNetworkList registers a new resource with the given unique name, arguments, and options.
func NewAppSecByPassNetworkList(ctx *pulumi.Context,
	name string, args *AppSecByPassNetworkListArgs, opts ...pulumi.ResourceOption) (*AppSecByPassNetworkList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BypassNetworkLists == nil {
		return nil, errors.New("invalid value for required argument 'BypassNetworkLists'")
	}
	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	var resource AppSecByPassNetworkList
	err := ctx.RegisterResource("akamai:index/appSecByPassNetworkList:AppSecByPassNetworkList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecByPassNetworkList gets an existing AppSecByPassNetworkList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecByPassNetworkList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecByPassNetworkListState, opts ...pulumi.ResourceOption) (*AppSecByPassNetworkList, error) {
	var resource AppSecByPassNetworkList
	err := ctx.ReadResource("akamai:index/appSecByPassNetworkList:AppSecByPassNetworkList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecByPassNetworkList resources.
type appSecByPassNetworkListState struct {
	// A list containing the IDs of the network lists to use.
	BypassNetworkLists []string `pulumi:"bypassNetworkLists"`
	// The configuration ID to use.
	ConfigId *int `pulumi:"configId"`
}

type AppSecByPassNetworkListState struct {
	// A list containing the IDs of the network lists to use.
	BypassNetworkLists pulumi.StringArrayInput
	// The configuration ID to use.
	ConfigId pulumi.IntPtrInput
}

func (AppSecByPassNetworkListState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecByPassNetworkListState)(nil)).Elem()
}

type appSecByPassNetworkListArgs struct {
	// A list containing the IDs of the network lists to use.
	BypassNetworkLists []string `pulumi:"bypassNetworkLists"`
	// The configuration ID to use.
	ConfigId int `pulumi:"configId"`
}

// The set of arguments for constructing a AppSecByPassNetworkList resource.
type AppSecByPassNetworkListArgs struct {
	// A list containing the IDs of the network lists to use.
	BypassNetworkLists pulumi.StringArrayInput
	// The configuration ID to use.
	ConfigId pulumi.IntInput
}

func (AppSecByPassNetworkListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecByPassNetworkListArgs)(nil)).Elem()
}

type AppSecByPassNetworkListInput interface {
	pulumi.Input

	ToAppSecByPassNetworkListOutput() AppSecByPassNetworkListOutput
	ToAppSecByPassNetworkListOutputWithContext(ctx context.Context) AppSecByPassNetworkListOutput
}

func (*AppSecByPassNetworkList) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecByPassNetworkList)(nil))
}

func (i *AppSecByPassNetworkList) ToAppSecByPassNetworkListOutput() AppSecByPassNetworkListOutput {
	return i.ToAppSecByPassNetworkListOutputWithContext(context.Background())
}

func (i *AppSecByPassNetworkList) ToAppSecByPassNetworkListOutputWithContext(ctx context.Context) AppSecByPassNetworkListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecByPassNetworkListOutput)
}

func (i *AppSecByPassNetworkList) ToAppSecByPassNetworkListPtrOutput() AppSecByPassNetworkListPtrOutput {
	return i.ToAppSecByPassNetworkListPtrOutputWithContext(context.Background())
}

func (i *AppSecByPassNetworkList) ToAppSecByPassNetworkListPtrOutputWithContext(ctx context.Context) AppSecByPassNetworkListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecByPassNetworkListPtrOutput)
}

type AppSecByPassNetworkListPtrInput interface {
	pulumi.Input

	ToAppSecByPassNetworkListPtrOutput() AppSecByPassNetworkListPtrOutput
	ToAppSecByPassNetworkListPtrOutputWithContext(ctx context.Context) AppSecByPassNetworkListPtrOutput
}

type appSecByPassNetworkListPtrType AppSecByPassNetworkListArgs

func (*appSecByPassNetworkListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecByPassNetworkList)(nil))
}

func (i *appSecByPassNetworkListPtrType) ToAppSecByPassNetworkListPtrOutput() AppSecByPassNetworkListPtrOutput {
	return i.ToAppSecByPassNetworkListPtrOutputWithContext(context.Background())
}

func (i *appSecByPassNetworkListPtrType) ToAppSecByPassNetworkListPtrOutputWithContext(ctx context.Context) AppSecByPassNetworkListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecByPassNetworkListPtrOutput)
}

// AppSecByPassNetworkListArrayInput is an input type that accepts AppSecByPassNetworkListArray and AppSecByPassNetworkListArrayOutput values.
// You can construct a concrete instance of `AppSecByPassNetworkListArrayInput` via:
//
//          AppSecByPassNetworkListArray{ AppSecByPassNetworkListArgs{...} }
type AppSecByPassNetworkListArrayInput interface {
	pulumi.Input

	ToAppSecByPassNetworkListArrayOutput() AppSecByPassNetworkListArrayOutput
	ToAppSecByPassNetworkListArrayOutputWithContext(context.Context) AppSecByPassNetworkListArrayOutput
}

type AppSecByPassNetworkListArray []AppSecByPassNetworkListInput

func (AppSecByPassNetworkListArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecByPassNetworkList)(nil))
}

func (i AppSecByPassNetworkListArray) ToAppSecByPassNetworkListArrayOutput() AppSecByPassNetworkListArrayOutput {
	return i.ToAppSecByPassNetworkListArrayOutputWithContext(context.Background())
}

func (i AppSecByPassNetworkListArray) ToAppSecByPassNetworkListArrayOutputWithContext(ctx context.Context) AppSecByPassNetworkListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecByPassNetworkListArrayOutput)
}

// AppSecByPassNetworkListMapInput is an input type that accepts AppSecByPassNetworkListMap and AppSecByPassNetworkListMapOutput values.
// You can construct a concrete instance of `AppSecByPassNetworkListMapInput` via:
//
//          AppSecByPassNetworkListMap{ "key": AppSecByPassNetworkListArgs{...} }
type AppSecByPassNetworkListMapInput interface {
	pulumi.Input

	ToAppSecByPassNetworkListMapOutput() AppSecByPassNetworkListMapOutput
	ToAppSecByPassNetworkListMapOutputWithContext(context.Context) AppSecByPassNetworkListMapOutput
}

type AppSecByPassNetworkListMap map[string]AppSecByPassNetworkListInput

func (AppSecByPassNetworkListMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecByPassNetworkList)(nil))
}

func (i AppSecByPassNetworkListMap) ToAppSecByPassNetworkListMapOutput() AppSecByPassNetworkListMapOutput {
	return i.ToAppSecByPassNetworkListMapOutputWithContext(context.Background())
}

func (i AppSecByPassNetworkListMap) ToAppSecByPassNetworkListMapOutputWithContext(ctx context.Context) AppSecByPassNetworkListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecByPassNetworkListMapOutput)
}

type AppSecByPassNetworkListOutput struct {
	*pulumi.OutputState
}

func (AppSecByPassNetworkListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecByPassNetworkList)(nil))
}

func (o AppSecByPassNetworkListOutput) ToAppSecByPassNetworkListOutput() AppSecByPassNetworkListOutput {
	return o
}

func (o AppSecByPassNetworkListOutput) ToAppSecByPassNetworkListOutputWithContext(ctx context.Context) AppSecByPassNetworkListOutput {
	return o
}

func (o AppSecByPassNetworkListOutput) ToAppSecByPassNetworkListPtrOutput() AppSecByPassNetworkListPtrOutput {
	return o.ToAppSecByPassNetworkListPtrOutputWithContext(context.Background())
}

func (o AppSecByPassNetworkListOutput) ToAppSecByPassNetworkListPtrOutputWithContext(ctx context.Context) AppSecByPassNetworkListPtrOutput {
	return o.ApplyT(func(v AppSecByPassNetworkList) *AppSecByPassNetworkList {
		return &v
	}).(AppSecByPassNetworkListPtrOutput)
}

type AppSecByPassNetworkListPtrOutput struct {
	*pulumi.OutputState
}

func (AppSecByPassNetworkListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecByPassNetworkList)(nil))
}

func (o AppSecByPassNetworkListPtrOutput) ToAppSecByPassNetworkListPtrOutput() AppSecByPassNetworkListPtrOutput {
	return o
}

func (o AppSecByPassNetworkListPtrOutput) ToAppSecByPassNetworkListPtrOutputWithContext(ctx context.Context) AppSecByPassNetworkListPtrOutput {
	return o
}

type AppSecByPassNetworkListArrayOutput struct{ *pulumi.OutputState }

func (AppSecByPassNetworkListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecByPassNetworkList)(nil))
}

func (o AppSecByPassNetworkListArrayOutput) ToAppSecByPassNetworkListArrayOutput() AppSecByPassNetworkListArrayOutput {
	return o
}

func (o AppSecByPassNetworkListArrayOutput) ToAppSecByPassNetworkListArrayOutputWithContext(ctx context.Context) AppSecByPassNetworkListArrayOutput {
	return o
}

func (o AppSecByPassNetworkListArrayOutput) Index(i pulumi.IntInput) AppSecByPassNetworkListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecByPassNetworkList {
		return vs[0].([]AppSecByPassNetworkList)[vs[1].(int)]
	}).(AppSecByPassNetworkListOutput)
}

type AppSecByPassNetworkListMapOutput struct{ *pulumi.OutputState }

func (AppSecByPassNetworkListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecByPassNetworkList)(nil))
}

func (o AppSecByPassNetworkListMapOutput) ToAppSecByPassNetworkListMapOutput() AppSecByPassNetworkListMapOutput {
	return o
}

func (o AppSecByPassNetworkListMapOutput) ToAppSecByPassNetworkListMapOutputWithContext(ctx context.Context) AppSecByPassNetworkListMapOutput {
	return o
}

func (o AppSecByPassNetworkListMapOutput) MapIndex(k pulumi.StringInput) AppSecByPassNetworkListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecByPassNetworkList {
		return vs[0].(map[string]AppSecByPassNetworkList)[vs[1].(string)]
	}).(AppSecByPassNetworkListOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecByPassNetworkListOutput{})
	pulumi.RegisterOutputType(AppSecByPassNetworkListPtrOutput{})
	pulumi.RegisterOutputType(AppSecByPassNetworkListArrayOutput{})
	pulumi.RegisterOutputType(AppSecByPassNetworkListMapOutput{})
}
