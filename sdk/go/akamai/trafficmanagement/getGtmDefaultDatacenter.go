// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: akamai.trafficmanagement.getGtmDefaultDatacenter has been deprecated in favor of akamai.getGtmDefaultDatacenter
func GetGtmDefaultDatacenter(ctx *pulumi.Context, args *GetGtmDefaultDatacenterArgs, opts ...pulumi.InvokeOption) (*GetGtmDefaultDatacenterResult, error) {
	var rv GetGtmDefaultDatacenterResult
	err := ctx.Invoke("akamai:trafficmanagement/getGtmDefaultDatacenter:getGtmDefaultDatacenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGtmDefaultDatacenter.
type GetGtmDefaultDatacenterArgs struct {
	Datacenter *int   `pulumi:"datacenter"`
	Domain     string `pulumi:"domain"`
}

// A collection of values returned by getGtmDefaultDatacenter.
type GetGtmDefaultDatacenterResult struct {
	Datacenter   *int   `pulumi:"datacenter"`
	DatacenterId int    `pulumi:"datacenterId"`
	Domain       string `pulumi:"domain"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Nickname string `pulumi:"nickname"`
}
