// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use `getIamTimezones` to list all time zones Akamai supports. Time zones are in ISO 8601 format. Use the values from this data source to set the time zone for a user. Administrators use this data source to set a user's time zone. The default time zone is GMT.
//
// ## Attributes reference
//
// These attributes are returned:
//
// * `timezones` — Supported timezones.
//   - `timezone` - The time zone ID.
//   - `description` - The description of a time zone, including the GMT +/-.
//   - `offset` - The time zone offset from GMT.
//   - `posix` - The time zone posix.
//
// [API Reference](https://techdocs.akamai.com/iam-api/reference/get-common-timezones)
func GetIamTimezones(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIamTimezonesResult, error) {
	var rv GetIamTimezonesResult
	err := ctx.Invoke("akamai:index/getIamTimezones:getIamTimezones", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIamTimezones.
type GetIamTimezonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string                    `pulumi:"id"`
	Timezones []GetIamTimezonesTimezone `pulumi:"timezones"`
}
