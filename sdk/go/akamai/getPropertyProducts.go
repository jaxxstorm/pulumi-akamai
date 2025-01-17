// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getPropertyProducts` data source to list the products included on a contract.
//
// ## Example Usage
//
// This example returns products associated with the [EdgeGrid client token](https://techdocs.akamai.com/developer/docs/authenticate-with-edgegrid) for a given contract:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ctx.Export("propertyMatch", data.Akamai_property_products.My-example)
//			return nil
//		})
//	}
//
// ```
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `products` - A list of supported products for the contract, including:
//   - `productId` - The product's unique ID, including the `prd_` prefix.
//   - `productName` - A string containing the product name.
func GetPropertyProducts(ctx *pulumi.Context, args *GetPropertyProductsArgs, opts ...pulumi.InvokeOption) (*GetPropertyProductsResult, error) {
	var rv GetPropertyProductsResult
	err := ctx.Invoke("akamai:index/getPropertyProducts:getPropertyProducts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPropertyProducts.
type GetPropertyProductsArgs struct {
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId string `pulumi:"contractId"`
}

// A collection of values returned by getPropertyProducts.
type GetPropertyProductsResult struct {
	ContractId string `pulumi:"contractId"`
	// The provider-assigned unique ID for this managed resource.
	Id       string                       `pulumi:"id"`
	Products []GetPropertyProductsProduct `pulumi:"products"`
}

func GetPropertyProductsOutput(ctx *pulumi.Context, args GetPropertyProductsOutputArgs, opts ...pulumi.InvokeOption) GetPropertyProductsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPropertyProductsResult, error) {
			args := v.(GetPropertyProductsArgs)
			r, err := GetPropertyProducts(ctx, &args, opts...)
			var s GetPropertyProductsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPropertyProductsResultOutput)
}

// A collection of arguments for invoking getPropertyProducts.
type GetPropertyProductsOutputArgs struct {
	// - (Required) A contract's unique ID, including the `ctr_` prefix.
	ContractId pulumi.StringInput `pulumi:"contractId"`
}

func (GetPropertyProductsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPropertyProductsArgs)(nil)).Elem()
}

// A collection of values returned by getPropertyProducts.
type GetPropertyProductsResultOutput struct{ *pulumi.OutputState }

func (GetPropertyProductsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPropertyProductsResult)(nil)).Elem()
}

func (o GetPropertyProductsResultOutput) ToGetPropertyProductsResultOutput() GetPropertyProductsResultOutput {
	return o
}

func (o GetPropertyProductsResultOutput) ToGetPropertyProductsResultOutputWithContext(ctx context.Context) GetPropertyProductsResultOutput {
	return o
}

func (o GetPropertyProductsResultOutput) ContractId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPropertyProductsResult) string { return v.ContractId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPropertyProductsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPropertyProductsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPropertyProductsResultOutput) Products() GetPropertyProductsProductArrayOutput {
	return o.ApplyT(func(v GetPropertyProductsResult) []GetPropertyProductsProduct { return v.Products }).(GetPropertyProductsProductArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPropertyProductsResultOutput{})
}
