// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List which grantable roles you can include in a new custom role or add to an existing custom role.
//
// ## Attributes reference
//
// This resource returns this attribute:
//
// * `grantableRoles` - Lists which grantable roles you can include in a new custom role or add to an existing custom role.
//   - `grantedRoleId` - Granted role ID.
//   - `name` - Granted role name.
//   - `description` - Granted role description.
func GetIamGrantableRoles(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIamGrantableRolesResult, error) {
	var rv GetIamGrantableRolesResult
	err := ctx.Invoke("akamai:index/getIamGrantableRoles:getIamGrantableRoles", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIamGrantableRoles.
type GetIamGrantableRolesResult struct {
	GrantableRoles []GetIamGrantableRolesGrantableRole `pulumi:"grantableRoles"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
