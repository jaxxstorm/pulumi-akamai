// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use `getIamSupportedLangs` to list all the possible languages Akamai supports. Use the values from this API to set the preferred language for a user. Users should see Control Center in the language you set for them. The default language is English.
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
//			supportedLangs, err := akamai.GetIamSupportedLangs(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("supportedSupportedLangs", supportedLangs)
//			return nil
//		})
//	}
//
// ```
// ## Attributes reference
//
// These attributes are returned:
//
// * `languages` — Languages supported by Akamai
//
// [API Reference](https://techdocs.akamai.com/iam-api/reference/get-user-languages)
func GetIamSupportedLangs(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIamSupportedLangsResult, error) {
	var rv GetIamSupportedLangsResult
	err := ctx.Invoke("akamai:index/getIamSupportedLangs:getIamSupportedLangs", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIamSupportedLangs.
type GetIamSupportedLangsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Languages []string `pulumi:"languages"`
}
