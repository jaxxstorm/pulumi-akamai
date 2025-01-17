// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `EdgeKv` resource lets you control EdgeKV database functions outside EdgeWorkers JavaScript code. Refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/welcome-to-edgekv) for more information.
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
//			_, err := akamai.NewEdgeKv(ctx, "testStaging", &akamai.EdgeKvArgs{
//				GeoLocation: pulumi.String("US"),
//				GroupId:     pulumi.Int(4284),
//				InitialDatas: EdgeKvInitialDataArray{
//					&EdgeKvInitialDataArgs{
//						Group: pulumi.String("translations"),
//						Key:   pulumi.String("lang"),
//						Value: pulumi.String("English"),
//					},
//				},
//				NamespaceName:      pulumi.String("Marketing"),
//				Network:            pulumi.String("staging"),
//				RetentionInSeconds: pulumi.Int(15724800),
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
// There are no supported arguments for this resource.
type EdgeKv struct {
	pulumi.CustomResourceState

	// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
	GeoLocation pulumi.StringPtrOutput `pulumi:"geoLocation"`
	// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
	GroupId pulumi.IntOutput `pulumi:"groupId"`
	// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
	InitialDatas EdgeKvInitialDataArrayOutput `pulumi:"initialDatas"`
	// - (Required) The name of the namespace.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network pulumi.StringOutput `pulumi:"network"`
	// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
	RetentionInSeconds pulumi.IntOutput `pulumi:"retentionInSeconds"`
}

// NewEdgeKv registers a new resource with the given unique name, arguments, and options.
func NewEdgeKv(ctx *pulumi.Context,
	name string, args *EdgeKvArgs, opts ...pulumi.ResourceOption) (*EdgeKv, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.RetentionInSeconds == nil {
		return nil, errors.New("invalid value for required argument 'RetentionInSeconds'")
	}
	var resource EdgeKv
	err := ctx.RegisterResource("akamai:index/edgeKv:EdgeKv", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeKv gets an existing EdgeKv resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeKv(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeKvState, opts ...pulumi.ResourceOption) (*EdgeKv, error) {
	var resource EdgeKv
	err := ctx.ReadResource("akamai:index/edgeKv:EdgeKv", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeKv resources.
type edgeKvState struct {
	// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
	GeoLocation *string `pulumi:"geoLocation"`
	// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
	GroupId *int `pulumi:"groupId"`
	// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
	InitialDatas []EdgeKvInitialData `pulumi:"initialDatas"`
	// - (Required) The name of the namespace.
	NamespaceName *string `pulumi:"namespaceName"`
	// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network *string `pulumi:"network"`
	// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
	RetentionInSeconds *int `pulumi:"retentionInSeconds"`
}

type EdgeKvState struct {
	// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
	GeoLocation pulumi.StringPtrInput
	// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
	GroupId pulumi.IntPtrInput
	// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
	InitialDatas EdgeKvInitialDataArrayInput
	// - (Required) The name of the namespace.
	NamespaceName pulumi.StringPtrInput
	// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network pulumi.StringPtrInput
	// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
	RetentionInSeconds pulumi.IntPtrInput
}

func (EdgeKvState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeKvState)(nil)).Elem()
}

type edgeKvArgs struct {
	// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
	GeoLocation *string `pulumi:"geoLocation"`
	// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
	GroupId int `pulumi:"groupId"`
	// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
	InitialDatas []EdgeKvInitialData `pulumi:"initialDatas"`
	// - (Required) The name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network string `pulumi:"network"`
	// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
	RetentionInSeconds int `pulumi:"retentionInSeconds"`
}

// The set of arguments for constructing a EdgeKv resource.
type EdgeKvArgs struct {
	// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
	GeoLocation pulumi.StringPtrInput
	// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
	GroupId pulumi.IntInput
	// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
	InitialDatas EdgeKvInitialDataArrayInput
	// - (Required) The name of the namespace.
	NamespaceName pulumi.StringInput
	// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network pulumi.StringInput
	// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
	RetentionInSeconds pulumi.IntInput
}

func (EdgeKvArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeKvArgs)(nil)).Elem()
}

type EdgeKvInput interface {
	pulumi.Input

	ToEdgeKvOutput() EdgeKvOutput
	ToEdgeKvOutputWithContext(ctx context.Context) EdgeKvOutput
}

func (*EdgeKv) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeKv)(nil)).Elem()
}

func (i *EdgeKv) ToEdgeKvOutput() EdgeKvOutput {
	return i.ToEdgeKvOutputWithContext(context.Background())
}

func (i *EdgeKv) ToEdgeKvOutputWithContext(ctx context.Context) EdgeKvOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeKvOutput)
}

// EdgeKvArrayInput is an input type that accepts EdgeKvArray and EdgeKvArrayOutput values.
// You can construct a concrete instance of `EdgeKvArrayInput` via:
//
//	EdgeKvArray{ EdgeKvArgs{...} }
type EdgeKvArrayInput interface {
	pulumi.Input

	ToEdgeKvArrayOutput() EdgeKvArrayOutput
	ToEdgeKvArrayOutputWithContext(context.Context) EdgeKvArrayOutput
}

type EdgeKvArray []EdgeKvInput

func (EdgeKvArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeKv)(nil)).Elem()
}

func (i EdgeKvArray) ToEdgeKvArrayOutput() EdgeKvArrayOutput {
	return i.ToEdgeKvArrayOutputWithContext(context.Background())
}

func (i EdgeKvArray) ToEdgeKvArrayOutputWithContext(ctx context.Context) EdgeKvArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeKvArrayOutput)
}

// EdgeKvMapInput is an input type that accepts EdgeKvMap and EdgeKvMapOutput values.
// You can construct a concrete instance of `EdgeKvMapInput` via:
//
//	EdgeKvMap{ "key": EdgeKvArgs{...} }
type EdgeKvMapInput interface {
	pulumi.Input

	ToEdgeKvMapOutput() EdgeKvMapOutput
	ToEdgeKvMapOutputWithContext(context.Context) EdgeKvMapOutput
}

type EdgeKvMap map[string]EdgeKvInput

func (EdgeKvMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeKv)(nil)).Elem()
}

func (i EdgeKvMap) ToEdgeKvMapOutput() EdgeKvMapOutput {
	return i.ToEdgeKvMapOutputWithContext(context.Background())
}

func (i EdgeKvMap) ToEdgeKvMapOutputWithContext(ctx context.Context) EdgeKvMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeKvMapOutput)
}

type EdgeKvOutput struct{ *pulumi.OutputState }

func (EdgeKvOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeKv)(nil)).Elem()
}

func (o EdgeKvOutput) ToEdgeKvOutput() EdgeKvOutput {
	return o
}

func (o EdgeKvOutput) ToEdgeKvOutputWithContext(ctx context.Context) EdgeKvOutput {
	return o
}

// Storage location for data when creating a namespace on the production network. This can help optimize performance by storing data where most or all of your users are located. The value defaults to `US` on the `STAGING` and `PRODUCTION` networks. For a list of supported geoLocations on the `PRODUCTION` network refer to the [EdgeKV documentation](https://techdocs.akamai.com/edgekv/docs/edgekv-data-model#namespace).
func (o EdgeKvOutput) GeoLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeKv) pulumi.StringPtrOutput { return v.GeoLocation }).(pulumi.StringPtrOutput)
}

// - (Required) The `group ID` for the EdgeKV namespace. This numeric value will be required in the next EdgeKV API version.
func (o EdgeKvOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *EdgeKv) pulumi.IntOutput { return v.GroupId }).(pulumi.IntOutput)
}

// List of key-value pairs called items to initialize the namespace. These items are valid only for database creation, updates are ignored.
func (o EdgeKvOutput) InitialDatas() EdgeKvInitialDataArrayOutput {
	return o.ApplyT(func(v *EdgeKv) EdgeKvInitialDataArrayOutput { return v.InitialDatas }).(EdgeKvInitialDataArrayOutput)
}

// - (Required) The name of the namespace.
func (o EdgeKvOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeKv) pulumi.StringOutput { return v.NamespaceName }).(pulumi.StringOutput)
}

// The network you want to activate the EdgeKV database on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
func (o EdgeKvOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeKv) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// - (Required) Retention period for data in this namespace, or 0 for indefinite. An update of this value will just affect new EdgeKV items.
func (o EdgeKvOutput) RetentionInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *EdgeKv) pulumi.IntOutput { return v.RetentionInSeconds }).(pulumi.IntOutput)
}

type EdgeKvArrayOutput struct{ *pulumi.OutputState }

func (EdgeKvArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeKv)(nil)).Elem()
}

func (o EdgeKvArrayOutput) ToEdgeKvArrayOutput() EdgeKvArrayOutput {
	return o
}

func (o EdgeKvArrayOutput) ToEdgeKvArrayOutputWithContext(ctx context.Context) EdgeKvArrayOutput {
	return o
}

func (o EdgeKvArrayOutput) Index(i pulumi.IntInput) EdgeKvOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EdgeKv {
		return vs[0].([]*EdgeKv)[vs[1].(int)]
	}).(EdgeKvOutput)
}

type EdgeKvMapOutput struct{ *pulumi.OutputState }

func (EdgeKvMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeKv)(nil)).Elem()
}

func (o EdgeKvMapOutput) ToEdgeKvMapOutput() EdgeKvMapOutput {
	return o
}

func (o EdgeKvMapOutput) ToEdgeKvMapOutputWithContext(ctx context.Context) EdgeKvMapOutput {
	return o
}

func (o EdgeKvMapOutput) MapIndex(k pulumi.StringInput) EdgeKvOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EdgeKv {
		return vs[0].(map[string]*EdgeKv)[vs[1].(string)]
	}).(EdgeKvOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeKvInput)(nil)).Elem(), &EdgeKv{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeKvArrayInput)(nil)).Elem(), EdgeKvArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeKvMapInput)(nil)).Elem(), EdgeKvMap{})
	pulumi.RegisterOutputType(EdgeKvOutput{})
	pulumi.RegisterOutputType(EdgeKvArrayOutput{})
	pulumi.RegisterOutputType(EdgeKvMapOutput{})
}
