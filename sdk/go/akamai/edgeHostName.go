// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `EdgeHostName` resource lets you configure a secure edge hostname. Your
// edge hostname determines how requests for your site, app, or content are mapped to
// Akamai edge servers.
//
// An edge hostname is the CNAME target you use when directing your end user traffic to
// Akamai. Each hostname assigned to a property has a corresponding edge hostname.
//
// Akamai supports three types of edge hostnames, depending on the level of security
// you need for your traffic: Standard TLS, Enhanced TLS, and Shared Certificate. When
// entering the `edgeHostname` attribute, you need to include a specific domain suffix
// for your edge hostname type:
//
// | Edge hostname type | Domain suffix |
// |------|-------|
// | Enhanced TLS | edgekey.net |
// | Standard TLS | edgesuite.net |
// | Shared Cert | akamaized.net |
//
// For example, if you use Standard TLS and have `www.example.com` as a hostname, your edge hostname would be `www.example.com.edgesuite.net`. If you wanted to use Enhanced TLS with the same hostname, your edge hostname would be `www.example.com.edgekey.net`. See the [Property Manager API (PAPI)](https://developer.akamai.com/api/core_features/property_manager/v1.html#createedgehostnames) for more information.
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
// 		_, err := akamai.NewEdgeHostName(ctx, "provider_demo", &akamai.EdgeHostNameArgs{
// 			ContractId:   pulumi.String("ctr_1-AB123"),
// 			EdgeHostname: pulumi.String("www.example.org.edgesuite.net"),
// 			GroupId:      pulumi.String("grp_123"),
// 			ProductId:    pulumi.String("prd_Object_Delivery"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Argument reference
//
// This resource supports these arguments:
//
// * `name` - (Required) The name of the edge hostname.
// * `contractId` - (Required) A contract's unique ID, including the `ctr_` prefix.
// * `groupId` - (Required) A group's unique ID, including the `grp_` prefix.
// * `productId` - (Required) A product's unique ID, including the `prd_` prefix.
// * `edgeHostname` - (Required) One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
// * `certificate` - (Optional) Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
// * `ipBehavior` - (Required) Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6. The default value is `IPV4`.
// * `contract` - (Deprecated) Replaced by `contractId`. Maintained for legacy purposes.
// * `group` - (Deprecated) Replaced by `groupId`. Maintained for legacy purposes.
// * `product` - (Deprecated) Replaced by `productId`. Maintained for legacy purposes.
//
// ## Attributes reference
//
// This resource returns this attribute:
//
// * `ipBehavior` - Returns the IP protocol the hostname will use, either `IPV4` for version 4, IPV6_PERFORMANCE`for version 6, or`IPV6_COMPLIANCE` for both.
//
// ## Import
//
// Basic Usagehcl resource "akamai_edge_hostname" "example" {
//
// # (resource arguments)
//
//  } You can import Akamai edge hostnames using a comma-delimited string of edge
//
// hostname, contract ID, and group ID. You have to enter the values in this order:
//
//  `edge_hostname, contract_id, group_id`
//
// For example
//
// ```sh
//  $ pulumi import akamai:index/edgeHostName:EdgeHostName example ehn_123,ctr_1-AB123,grp_123
// ```
type EdgeHostName struct {
	pulumi.CustomResourceState

	Certificate pulumi.IntPtrOutput `pulumi:"certificate"`
	// Deprecated: The setting "contract" has been deprecated.
	Contract     pulumi.StringOutput `pulumi:"contract"`
	ContractId   pulumi.StringOutput `pulumi:"contractId"`
	EdgeHostname pulumi.StringOutput `pulumi:"edgeHostname"`
	// Deprecated: The setting "group" has been deprecated.
	Group      pulumi.StringOutput `pulumi:"group"`
	GroupId    pulumi.StringOutput `pulumi:"groupId"`
	IpBehavior pulumi.StringOutput `pulumi:"ipBehavior"`
	// Deprecated: The setting "product" has been deprecated.
	Product   pulumi.StringOutput `pulumi:"product"`
	ProductId pulumi.StringOutput `pulumi:"productId"`
}

// NewEdgeHostName registers a new resource with the given unique name, arguments, and options.
func NewEdgeHostName(ctx *pulumi.Context,
	name string, args *EdgeHostNameArgs, opts ...pulumi.ResourceOption) (*EdgeHostName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeHostname == nil {
		return nil, errors.New("invalid value for required argument 'EdgeHostname'")
	}
	if args.IpBehavior == nil {
		return nil, errors.New("invalid value for required argument 'IpBehavior'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("akamai:properties/edgeHostName:EdgeHostName"),
		},
	})
	opts = append(opts, aliases)
	var resource EdgeHostName
	err := ctx.RegisterResource("akamai:index/edgeHostName:EdgeHostName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeHostName gets an existing EdgeHostName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeHostName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeHostNameState, opts ...pulumi.ResourceOption) (*EdgeHostName, error) {
	var resource EdgeHostName
	err := ctx.ReadResource("akamai:index/edgeHostName:EdgeHostName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeHostName resources.
type edgeHostNameState struct {
	Certificate *int `pulumi:"certificate"`
	// Deprecated: The setting "contract" has been deprecated.
	Contract     *string `pulumi:"contract"`
	ContractId   *string `pulumi:"contractId"`
	EdgeHostname *string `pulumi:"edgeHostname"`
	// Deprecated: The setting "group" has been deprecated.
	Group      *string `pulumi:"group"`
	GroupId    *string `pulumi:"groupId"`
	IpBehavior *string `pulumi:"ipBehavior"`
	// Deprecated: The setting "product" has been deprecated.
	Product   *string `pulumi:"product"`
	ProductId *string `pulumi:"productId"`
}

type EdgeHostNameState struct {
	Certificate pulumi.IntPtrInput
	// Deprecated: The setting "contract" has been deprecated.
	Contract     pulumi.StringPtrInput
	ContractId   pulumi.StringPtrInput
	EdgeHostname pulumi.StringPtrInput
	// Deprecated: The setting "group" has been deprecated.
	Group      pulumi.StringPtrInput
	GroupId    pulumi.StringPtrInput
	IpBehavior pulumi.StringPtrInput
	// Deprecated: The setting "product" has been deprecated.
	Product   pulumi.StringPtrInput
	ProductId pulumi.StringPtrInput
}

func (EdgeHostNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeHostNameState)(nil)).Elem()
}

type edgeHostNameArgs struct {
	Certificate *int `pulumi:"certificate"`
	// Deprecated: The setting "contract" has been deprecated.
	Contract     *string `pulumi:"contract"`
	ContractId   *string `pulumi:"contractId"`
	EdgeHostname string  `pulumi:"edgeHostname"`
	// Deprecated: The setting "group" has been deprecated.
	Group      *string `pulumi:"group"`
	GroupId    *string `pulumi:"groupId"`
	IpBehavior string  `pulumi:"ipBehavior"`
	// Deprecated: The setting "product" has been deprecated.
	Product   *string `pulumi:"product"`
	ProductId *string `pulumi:"productId"`
}

// The set of arguments for constructing a EdgeHostName resource.
type EdgeHostNameArgs struct {
	Certificate pulumi.IntPtrInput
	// Deprecated: The setting "contract" has been deprecated.
	Contract     pulumi.StringPtrInput
	ContractId   pulumi.StringPtrInput
	EdgeHostname pulumi.StringInput
	// Deprecated: The setting "group" has been deprecated.
	Group      pulumi.StringPtrInput
	GroupId    pulumi.StringPtrInput
	IpBehavior pulumi.StringInput
	// Deprecated: The setting "product" has been deprecated.
	Product   pulumi.StringPtrInput
	ProductId pulumi.StringPtrInput
}

func (EdgeHostNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeHostNameArgs)(nil)).Elem()
}

type EdgeHostNameInput interface {
	pulumi.Input

	ToEdgeHostNameOutput() EdgeHostNameOutput
	ToEdgeHostNameOutputWithContext(ctx context.Context) EdgeHostNameOutput
}

func (*EdgeHostName) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeHostName)(nil))
}

func (i *EdgeHostName) ToEdgeHostNameOutput() EdgeHostNameOutput {
	return i.ToEdgeHostNameOutputWithContext(context.Background())
}

func (i *EdgeHostName) ToEdgeHostNameOutputWithContext(ctx context.Context) EdgeHostNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeHostNameOutput)
}

func (i *EdgeHostName) ToEdgeHostNamePtrOutput() EdgeHostNamePtrOutput {
	return i.ToEdgeHostNamePtrOutputWithContext(context.Background())
}

func (i *EdgeHostName) ToEdgeHostNamePtrOutputWithContext(ctx context.Context) EdgeHostNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeHostNamePtrOutput)
}

type EdgeHostNamePtrInput interface {
	pulumi.Input

	ToEdgeHostNamePtrOutput() EdgeHostNamePtrOutput
	ToEdgeHostNamePtrOutputWithContext(ctx context.Context) EdgeHostNamePtrOutput
}

type edgeHostNamePtrType EdgeHostNameArgs

func (*edgeHostNamePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeHostName)(nil))
}

func (i *edgeHostNamePtrType) ToEdgeHostNamePtrOutput() EdgeHostNamePtrOutput {
	return i.ToEdgeHostNamePtrOutputWithContext(context.Background())
}

func (i *edgeHostNamePtrType) ToEdgeHostNamePtrOutputWithContext(ctx context.Context) EdgeHostNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeHostNamePtrOutput)
}

// EdgeHostNameArrayInput is an input type that accepts EdgeHostNameArray and EdgeHostNameArrayOutput values.
// You can construct a concrete instance of `EdgeHostNameArrayInput` via:
//
//          EdgeHostNameArray{ EdgeHostNameArgs{...} }
type EdgeHostNameArrayInput interface {
	pulumi.Input

	ToEdgeHostNameArrayOutput() EdgeHostNameArrayOutput
	ToEdgeHostNameArrayOutputWithContext(context.Context) EdgeHostNameArrayOutput
}

type EdgeHostNameArray []EdgeHostNameInput

func (EdgeHostNameArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EdgeHostName)(nil))
}

func (i EdgeHostNameArray) ToEdgeHostNameArrayOutput() EdgeHostNameArrayOutput {
	return i.ToEdgeHostNameArrayOutputWithContext(context.Background())
}

func (i EdgeHostNameArray) ToEdgeHostNameArrayOutputWithContext(ctx context.Context) EdgeHostNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeHostNameArrayOutput)
}

// EdgeHostNameMapInput is an input type that accepts EdgeHostNameMap and EdgeHostNameMapOutput values.
// You can construct a concrete instance of `EdgeHostNameMapInput` via:
//
//          EdgeHostNameMap{ "key": EdgeHostNameArgs{...} }
type EdgeHostNameMapInput interface {
	pulumi.Input

	ToEdgeHostNameMapOutput() EdgeHostNameMapOutput
	ToEdgeHostNameMapOutputWithContext(context.Context) EdgeHostNameMapOutput
}

type EdgeHostNameMap map[string]EdgeHostNameInput

func (EdgeHostNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EdgeHostName)(nil))
}

func (i EdgeHostNameMap) ToEdgeHostNameMapOutput() EdgeHostNameMapOutput {
	return i.ToEdgeHostNameMapOutputWithContext(context.Background())
}

func (i EdgeHostNameMap) ToEdgeHostNameMapOutputWithContext(ctx context.Context) EdgeHostNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeHostNameMapOutput)
}

type EdgeHostNameOutput struct {
	*pulumi.OutputState
}

func (EdgeHostNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeHostName)(nil))
}

func (o EdgeHostNameOutput) ToEdgeHostNameOutput() EdgeHostNameOutput {
	return o
}

func (o EdgeHostNameOutput) ToEdgeHostNameOutputWithContext(ctx context.Context) EdgeHostNameOutput {
	return o
}

func (o EdgeHostNameOutput) ToEdgeHostNamePtrOutput() EdgeHostNamePtrOutput {
	return o.ToEdgeHostNamePtrOutputWithContext(context.Background())
}

func (o EdgeHostNameOutput) ToEdgeHostNamePtrOutputWithContext(ctx context.Context) EdgeHostNamePtrOutput {
	return o.ApplyT(func(v EdgeHostName) *EdgeHostName {
		return &v
	}).(EdgeHostNamePtrOutput)
}

type EdgeHostNamePtrOutput struct {
	*pulumi.OutputState
}

func (EdgeHostNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeHostName)(nil))
}

func (o EdgeHostNamePtrOutput) ToEdgeHostNamePtrOutput() EdgeHostNamePtrOutput {
	return o
}

func (o EdgeHostNamePtrOutput) ToEdgeHostNamePtrOutputWithContext(ctx context.Context) EdgeHostNamePtrOutput {
	return o
}

type EdgeHostNameArrayOutput struct{ *pulumi.OutputState }

func (EdgeHostNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdgeHostName)(nil))
}

func (o EdgeHostNameArrayOutput) ToEdgeHostNameArrayOutput() EdgeHostNameArrayOutput {
	return o
}

func (o EdgeHostNameArrayOutput) ToEdgeHostNameArrayOutputWithContext(ctx context.Context) EdgeHostNameArrayOutput {
	return o
}

func (o EdgeHostNameArrayOutput) Index(i pulumi.IntInput) EdgeHostNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdgeHostName {
		return vs[0].([]EdgeHostName)[vs[1].(int)]
	}).(EdgeHostNameOutput)
}

type EdgeHostNameMapOutput struct{ *pulumi.OutputState }

func (EdgeHostNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EdgeHostName)(nil))
}

func (o EdgeHostNameMapOutput) ToEdgeHostNameMapOutput() EdgeHostNameMapOutput {
	return o
}

func (o EdgeHostNameMapOutput) ToEdgeHostNameMapOutputWithContext(ctx context.Context) EdgeHostNameMapOutput {
	return o
}

func (o EdgeHostNameMapOutput) MapIndex(k pulumi.StringInput) EdgeHostNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EdgeHostName {
		return vs[0].(map[string]EdgeHostName)[vs[1].(string)]
	}).(EdgeHostNameOutput)
}

func init() {
	pulumi.RegisterOutputType(EdgeHostNameOutput{})
	pulumi.RegisterOutputType(EdgeHostNamePtrOutput{})
	pulumi.RegisterOutputType(EdgeHostNameArrayOutput{})
	pulumi.RegisterOutputType(EdgeHostNameMapOutput{})
}
