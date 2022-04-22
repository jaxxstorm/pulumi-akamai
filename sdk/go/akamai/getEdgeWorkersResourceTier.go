// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getEdgeWorkersResourceTier` data source to list the available resource tiers for a specific contract ID. The resource tier defines the resource consumption [limits](https://techdocs.akamai.com/edgeworkers/docs/resource-tier-limitations) for an EdgeWorker ID.
//
// ## Example Usage
//
// This example returns the resource tier fields for an EdgeWorker ID:
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
// 		_, err := akamai.GetEdgeWorkersResourceTier(ctx, &GetEdgeWorkersResourceTierArgs{
// 			ContractId:       "1-ABC",
// 			ResourceTierName: "Basic Compute",
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
// This data source returns these attributes:
//
// * `resourceTierId` - Unique identifier of the resource tier.
func GetEdgeWorkersResourceTier(ctx *pulumi.Context, args *GetEdgeWorkersResourceTierArgs, opts ...pulumi.InvokeOption) (*GetEdgeWorkersResourceTierResult, error) {
	var rv GetEdgeWorkersResourceTierResult
	err := ctx.Invoke("akamai:index/getEdgeWorkersResourceTier:getEdgeWorkersResourceTier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEdgeWorkersResourceTier.
type GetEdgeWorkersResourceTierArgs struct {
	// Unique identifier of a contract.
	ContractId string `pulumi:"contractId"`
	// Unique name of the resource tier.
	ResourceTierName string `pulumi:"resourceTierName"`
}

// A collection of values returned by getEdgeWorkersResourceTier.
type GetEdgeWorkersResourceTierResult struct {
	ContractId string `pulumi:"contractId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	ResourceTierId   int    `pulumi:"resourceTierId"`
	ResourceTierName string `pulumi:"resourceTierName"`
}

func GetEdgeWorkersResourceTierOutput(ctx *pulumi.Context, args GetEdgeWorkersResourceTierOutputArgs, opts ...pulumi.InvokeOption) GetEdgeWorkersResourceTierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEdgeWorkersResourceTierResult, error) {
			args := v.(GetEdgeWorkersResourceTierArgs)
			r, err := GetEdgeWorkersResourceTier(ctx, &args, opts...)
			var s GetEdgeWorkersResourceTierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEdgeWorkersResourceTierResultOutput)
}

// A collection of arguments for invoking getEdgeWorkersResourceTier.
type GetEdgeWorkersResourceTierOutputArgs struct {
	// Unique identifier of a contract.
	ContractId pulumi.StringInput `pulumi:"contractId"`
	// Unique name of the resource tier.
	ResourceTierName pulumi.StringInput `pulumi:"resourceTierName"`
}

func (GetEdgeWorkersResourceTierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgeWorkersResourceTierArgs)(nil)).Elem()
}

// A collection of values returned by getEdgeWorkersResourceTier.
type GetEdgeWorkersResourceTierResultOutput struct{ *pulumi.OutputState }

func (GetEdgeWorkersResourceTierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgeWorkersResourceTierResult)(nil)).Elem()
}

func (o GetEdgeWorkersResourceTierResultOutput) ToGetEdgeWorkersResourceTierResultOutput() GetEdgeWorkersResourceTierResultOutput {
	return o
}

func (o GetEdgeWorkersResourceTierResultOutput) ToGetEdgeWorkersResourceTierResultOutputWithContext(ctx context.Context) GetEdgeWorkersResourceTierResultOutput {
	return o
}

func (o GetEdgeWorkersResourceTierResultOutput) ContractId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeWorkersResourceTierResult) string { return v.ContractId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEdgeWorkersResourceTierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeWorkersResourceTierResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEdgeWorkersResourceTierResultOutput) ResourceTierId() pulumi.IntOutput {
	return o.ApplyT(func(v GetEdgeWorkersResourceTierResult) int { return v.ResourceTierId }).(pulumi.IntOutput)
}

func (o GetEdgeWorkersResourceTierResultOutput) ResourceTierName() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeWorkersResourceTierResult) string { return v.ResourceTierName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEdgeWorkersResourceTierResultOutput{})
}
