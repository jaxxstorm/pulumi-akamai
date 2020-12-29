// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type EdgeHostName struct {
	pulumi.CustomResourceState

	Certificate pulumi.IntPtrOutput `pulumi:"certificate"`
	// Deprecated: use "contract_id" attribute instead
	Contract     pulumi.StringOutput `pulumi:"contract"`
	ContractId   pulumi.StringOutput `pulumi:"contractId"`
	EdgeHostname pulumi.StringOutput `pulumi:"edgeHostname"`
	// Deprecated: use "group_id" attribute instead
	Group      pulumi.StringOutput `pulumi:"group"`
	GroupId    pulumi.StringOutput `pulumi:"groupId"`
	IpBehavior pulumi.StringOutput `pulumi:"ipBehavior"`
	// Deprecated: use "product_id" attribute instead
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
	// Deprecated: use "contract_id" attribute instead
	Contract     *string `pulumi:"contract"`
	ContractId   *string `pulumi:"contractId"`
	EdgeHostname *string `pulumi:"edgeHostname"`
	// Deprecated: use "group_id" attribute instead
	Group      *string `pulumi:"group"`
	GroupId    *string `pulumi:"groupId"`
	IpBehavior *string `pulumi:"ipBehavior"`
	// Deprecated: use "product_id" attribute instead
	Product   *string `pulumi:"product"`
	ProductId *string `pulumi:"productId"`
}

type EdgeHostNameState struct {
	Certificate pulumi.IntPtrInput
	// Deprecated: use "contract_id" attribute instead
	Contract     pulumi.StringPtrInput
	ContractId   pulumi.StringPtrInput
	EdgeHostname pulumi.StringPtrInput
	// Deprecated: use "group_id" attribute instead
	Group      pulumi.StringPtrInput
	GroupId    pulumi.StringPtrInput
	IpBehavior pulumi.StringPtrInput
	// Deprecated: use "product_id" attribute instead
	Product   pulumi.StringPtrInput
	ProductId pulumi.StringPtrInput
}

func (EdgeHostNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeHostNameState)(nil)).Elem()
}

type edgeHostNameArgs struct {
	Certificate *int `pulumi:"certificate"`
	// Deprecated: use "contract_id" attribute instead
	Contract     *string `pulumi:"contract"`
	ContractId   *string `pulumi:"contractId"`
	EdgeHostname string  `pulumi:"edgeHostname"`
	// Deprecated: use "group_id" attribute instead
	Group      *string `pulumi:"group"`
	GroupId    *string `pulumi:"groupId"`
	IpBehavior string  `pulumi:"ipBehavior"`
	// Deprecated: use "product_id" attribute instead
	Product   *string `pulumi:"product"`
	ProductId *string `pulumi:"productId"`
}

// The set of arguments for constructing a EdgeHostName resource.
type EdgeHostNameArgs struct {
	Certificate pulumi.IntPtrInput
	// Deprecated: use "contract_id" attribute instead
	Contract     pulumi.StringPtrInput
	ContractId   pulumi.StringPtrInput
	EdgeHostname pulumi.StringInput
	// Deprecated: use "group_id" attribute instead
	Group      pulumi.StringPtrInput
	GroupId    pulumi.StringPtrInput
	IpBehavior pulumi.StringInput
	// Deprecated: use "product_id" attribute instead
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

func (EdgeHostName) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeHostName)(nil)).Elem()
}

func (i EdgeHostName) ToEdgeHostNameOutput() EdgeHostNameOutput {
	return i.ToEdgeHostNameOutputWithContext(context.Background())
}

func (i EdgeHostName) ToEdgeHostNameOutputWithContext(ctx context.Context) EdgeHostNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeHostNameOutput)
}

type EdgeHostNameOutput struct {
	*pulumi.OutputState
}

func (EdgeHostNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeHostNameOutput)(nil)).Elem()
}

func (o EdgeHostNameOutput) ToEdgeHostNameOutput() EdgeHostNameOutput {
	return o
}

func (o EdgeHostNameOutput) ToEdgeHostNameOutputWithContext(ctx context.Context) EdgeHostNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EdgeHostNameOutput{})
}
