// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edgedns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAuthoritiesSet` data source to retrieve a contract's authorities set. You use the authorities set when creating new zones.
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
// 		_, err := akamai.GetAuthoritiesSet(ctx, &GetAuthoritiesSetArgs{
// 			Contract: "ctr_1-AB123",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Attributes reference
//
// This data source supports this attribute:
//
// * `authorities` - A list of authorities.
//
// Deprecated: akamai.edgedns.getAuthoritiesSet has been deprecated in favor of akamai.getAuthoritiesSet
func GetAuthoritiesSet(ctx *pulumi.Context, args *GetAuthoritiesSetArgs, opts ...pulumi.InvokeOption) (*GetAuthoritiesSetResult, error) {
	var rv GetAuthoritiesSetResult
	err := ctx.Invoke("akamai:edgedns/getAuthoritiesSet:getAuthoritiesSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthoritiesSet.
type GetAuthoritiesSetArgs struct {
	// The contract ID.
	Contract string `pulumi:"contract"`
}

// A collection of values returned by getAuthoritiesSet.
type GetAuthoritiesSetResult struct {
	Authorities []string `pulumi:"authorities"`
	Contract    string   `pulumi:"contract"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetAuthoritiesSetOutput(ctx *pulumi.Context, args GetAuthoritiesSetOutputArgs, opts ...pulumi.InvokeOption) GetAuthoritiesSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAuthoritiesSetResult, error) {
			args := v.(GetAuthoritiesSetArgs)
			r, err := GetAuthoritiesSet(ctx, &args, opts...)
			var s GetAuthoritiesSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAuthoritiesSetResultOutput)
}

// A collection of arguments for invoking getAuthoritiesSet.
type GetAuthoritiesSetOutputArgs struct {
	// The contract ID.
	Contract pulumi.StringInput `pulumi:"contract"`
}

func (GetAuthoritiesSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthoritiesSetArgs)(nil)).Elem()
}

// A collection of values returned by getAuthoritiesSet.
type GetAuthoritiesSetResultOutput struct{ *pulumi.OutputState }

func (GetAuthoritiesSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthoritiesSetResult)(nil)).Elem()
}

func (o GetAuthoritiesSetResultOutput) ToGetAuthoritiesSetResultOutput() GetAuthoritiesSetResultOutput {
	return o
}

func (o GetAuthoritiesSetResultOutput) ToGetAuthoritiesSetResultOutputWithContext(ctx context.Context) GetAuthoritiesSetResultOutput {
	return o
}

func (o GetAuthoritiesSetResultOutput) Authorities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAuthoritiesSetResult) []string { return v.Authorities }).(pulumi.StringArrayOutput)
}

func (o GetAuthoritiesSetResultOutput) Contract() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthoritiesSetResult) string { return v.Contract }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAuthoritiesSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthoritiesSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAuthoritiesSetResultOutput{})
}
