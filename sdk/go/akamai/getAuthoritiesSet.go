// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetAuthoritiesSet(ctx *pulumi.Context, args *GetAuthoritiesSetArgs, opts ...pulumi.InvokeOption) (*GetAuthoritiesSetResult, error) {
	var rv GetAuthoritiesSetResult
	err := ctx.Invoke("akamai:index/getAuthoritiesSet:getAuthoritiesSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthoritiesSet.
type GetAuthoritiesSetArgs struct {
	Contract string `pulumi:"contract"`
}

// A collection of values returned by getAuthoritiesSet.
type GetAuthoritiesSetResult struct {
	Authorities []string `pulumi:"authorities"`
	Contract    string   `pulumi:"contract"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
