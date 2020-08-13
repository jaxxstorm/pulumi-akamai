// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package properties

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `properties.CpCode` resource allows you to create or re-use CP Codes.
//
// If the CP Code already exists it will be used instead of creating a new one.
//
// ## Example Usage
// ### Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/go/akamai/properties"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := properties.NewCpCode(ctx, "cpCode", &properties.CpCodeArgs{
// 			Contract: pulumi.Any(akamai_contract.Contract.Id),
// 			Group:    pulumi.Any(akamai_group.Group.Id),
// 			Product:  pulumi.String("prd_xxx"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CpCode struct {
	pulumi.CustomResourceState

	// — (Required) The Contract ID
	Contract pulumi.StringOutput `pulumi:"contract"`
	// — (Required) The Group ID
	Group pulumi.StringOutput `pulumi:"group"`
	// — (Required) The CP Code name
	Name pulumi.StringOutput `pulumi:"name"`
	// — (Required) The Product ID
	Product pulumi.StringOutput `pulumi:"product"`
}

// NewCpCode registers a new resource with the given unique name, arguments, and options.
func NewCpCode(ctx *pulumi.Context,
	name string, args *CpCodeArgs, opts ...pulumi.ResourceOption) (*CpCode, error) {
	if args == nil || args.Contract == nil {
		return nil, errors.New("missing required argument 'Contract'")
	}
	if args == nil || args.Group == nil {
		return nil, errors.New("missing required argument 'Group'")
	}
	if args == nil || args.Product == nil {
		return nil, errors.New("missing required argument 'Product'")
	}
	if args == nil {
		args = &CpCodeArgs{}
	}
	var resource CpCode
	err := ctx.RegisterResource("akamai:properties/cpCode:CpCode", name, args, &resource, opts...)
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
	err := ctx.ReadResource("akamai:properties/cpCode:CpCode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CpCode resources.
type cpCodeState struct {
	// — (Required) The Contract ID
	Contract *string `pulumi:"contract"`
	// — (Required) The Group ID
	Group *string `pulumi:"group"`
	// — (Required) The CP Code name
	Name *string `pulumi:"name"`
	// — (Required) The Product ID
	Product *string `pulumi:"product"`
}

type CpCodeState struct {
	// — (Required) The Contract ID
	Contract pulumi.StringPtrInput
	// — (Required) The Group ID
	Group pulumi.StringPtrInput
	// — (Required) The CP Code name
	Name pulumi.StringPtrInput
	// — (Required) The Product ID
	Product pulumi.StringPtrInput
}

func (CpCodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*cpCodeState)(nil)).Elem()
}

type cpCodeArgs struct {
	// — (Required) The Contract ID
	Contract string `pulumi:"contract"`
	// — (Required) The Group ID
	Group string `pulumi:"group"`
	// — (Required) The CP Code name
	Name *string `pulumi:"name"`
	// — (Required) The Product ID
	Product string `pulumi:"product"`
}

// The set of arguments for constructing a CpCode resource.
type CpCodeArgs struct {
	// — (Required) The Contract ID
	Contract pulumi.StringInput
	// — (Required) The Group ID
	Group pulumi.StringInput
	// — (Required) The CP Code name
	Name pulumi.StringPtrInput
	// — (Required) The Product ID
	Product pulumi.StringInput
}

func (CpCodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cpCodeArgs)(nil)).Elem()
}
