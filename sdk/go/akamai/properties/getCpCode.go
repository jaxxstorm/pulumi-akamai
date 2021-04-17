// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package properties

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: akamai.properties.getCpCode has been deprecated in favor of akamai.getCpCode
func LookupCpCode(ctx *pulumi.Context, args *LookupCpCodeArgs, opts ...pulumi.InvokeOption) (*LookupCpCodeResult, error) {
	var rv LookupCpCodeResult
	err := ctx.Invoke("akamai:properties/getCpCode:getCpCode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCpCode.
type LookupCpCodeArgs struct {
	// Deprecated: The setting "contract" has been deprecated.
	Contract   *string `pulumi:"contract"`
	ContractId *string `pulumi:"contractId"`
	// Deprecated: The setting "group" has been deprecated.
	Group   *string `pulumi:"group"`
	GroupId *string `pulumi:"groupId"`
	Name    string  `pulumi:"name"`
}

// A collection of values returned by getCpCode.
type LookupCpCodeResult struct {
	// Deprecated: The setting "contract" has been deprecated.
	Contract   string `pulumi:"contract"`
	ContractId string `pulumi:"contractId"`
	// Deprecated: The setting "group" has been deprecated.
	Group   string `pulumi:"group"`
	GroupId string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Name       string   `pulumi:"name"`
	ProductIds []string `pulumi:"productIds"`
}
