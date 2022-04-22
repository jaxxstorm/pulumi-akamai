// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package properties

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `Property` data source to query and list the property ID and rule tree based on the property name.
//
// ## Example Usage
//
// This example returns the property ID and rule tree based on the property name and optional version argument:
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
// 		example, err := akamai.LookupProperty(ctx, &GetPropertyArgs{
// 			Name:    "terraform-demo",
// 			Version: pulumi.IntRef(1),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("myPropertyID", example)
// 		return nil
// 	})
// }
// ```
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `property_ID` - A property's unique identifier, including the `prp_` prefix.
// * `rules` - A JSON-encoded rule tree for a given property.
//
// Deprecated: akamai.properties.getProperty has been deprecated in favor of akamai.getProperty
func LookupProperty(ctx *pulumi.Context, args *LookupPropertyArgs, opts ...pulumi.InvokeOption) (*LookupPropertyResult, error) {
	var rv LookupPropertyResult
	err := ctx.Invoke("akamai:properties/getProperty:getProperty", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProperty.
type LookupPropertyArgs struct {
	// - (Required) The property name.
	Name string `pulumi:"name"`
	// - (Optional) The version of the property whose ID you want to list.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getProperty.
type LookupPropertyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	Rules   string `pulumi:"rules"`
	Version *int   `pulumi:"version"`
}

func LookupPropertyOutput(ctx *pulumi.Context, args LookupPropertyOutputArgs, opts ...pulumi.InvokeOption) LookupPropertyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPropertyResult, error) {
			args := v.(LookupPropertyArgs)
			r, err := LookupProperty(ctx, &args, opts...)
			var s LookupPropertyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPropertyResultOutput)
}

// A collection of arguments for invoking getProperty.
type LookupPropertyOutputArgs struct {
	// - (Required) The property name.
	Name pulumi.StringInput `pulumi:"name"`
	// - (Optional) The version of the property whose ID you want to list.
	Version pulumi.IntPtrInput `pulumi:"version"`
}

func (LookupPropertyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyArgs)(nil)).Elem()
}

// A collection of values returned by getProperty.
type LookupPropertyResultOutput struct{ *pulumi.OutputState }

func (LookupPropertyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyResult)(nil)).Elem()
}

func (o LookupPropertyResultOutput) ToLookupPropertyResultOutput() LookupPropertyResultOutput {
	return o
}

func (o LookupPropertyResultOutput) ToLookupPropertyResultOutputWithContext(ctx context.Context) LookupPropertyResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPropertyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPropertyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPropertyResultOutput) Rules() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyResult) string { return v.Rules }).(pulumi.StringOutput)
}

func (o LookupPropertyResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPropertyResult) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPropertyResultOutput{})
}
