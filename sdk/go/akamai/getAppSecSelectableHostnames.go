// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration; contract; group
//
// Returns the list of hostnames that can be (but aren't yet) protected by a security configuration. You can specify the set of hostnames to be retrieved either by supplying the name of a security configuration or by supplying an Akamai group ID and contract ID.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selectable-hostnames](https://techdocs.akamai.com/application-security/reference/get-selectable-hostnames)
//
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `hostnames`. List of selectable hostnames.
// - `hostnamesJson`. JSON-formatted list of selectable hostnames.
// - `outputText`. Tabular report of the selectable hostnames showing the name and configId of the security configuration under which the host is protected in production.
func GetAppSecSelectableHostnames(ctx *pulumi.Context, args *GetAppSecSelectableHostnamesArgs, opts ...pulumi.InvokeOption) (*GetAppSecSelectableHostnamesResult, error) {
	var rv GetAppSecSelectableHostnamesResult
	err := ctx.Invoke("akamai:index/getAppSecSelectableHostnames:getAppSecSelectableHostnames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecSelectableHostnames.
type GetAppSecSelectableHostnamesArgs struct {
	ActiveInProduction *bool `pulumi:"activeInProduction"`
	ActiveInStaging    *bool `pulumi:"activeInStaging"`
	// . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can't be used with either the `contractid` or the `groupid` arguments.
	ConfigId *int `pulumi:"configId"`
	// . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can't be used with the `configId` argument.
	Contractid *string `pulumi:"contractid"`
	// . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can't be used with the `configId` argument.
	Groupid *int `pulumi:"groupid"`
}

// A collection of values returned by getAppSecSelectableHostnames.
type GetAppSecSelectableHostnamesResult struct {
	ActiveInProduction *bool    `pulumi:"activeInProduction"`
	ActiveInStaging    *bool    `pulumi:"activeInStaging"`
	ConfigId           *int     `pulumi:"configId"`
	Contractid         *string  `pulumi:"contractid"`
	Groupid            *int     `pulumi:"groupid"`
	Hostnames          []string `pulumi:"hostnames"`
	HostnamesJson      string   `pulumi:"hostnamesJson"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	OutputText string `pulumi:"outputText"`
}

func GetAppSecSelectableHostnamesOutput(ctx *pulumi.Context, args GetAppSecSelectableHostnamesOutputArgs, opts ...pulumi.InvokeOption) GetAppSecSelectableHostnamesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecSelectableHostnamesResult, error) {
			args := v.(GetAppSecSelectableHostnamesArgs)
			r, err := GetAppSecSelectableHostnames(ctx, &args, opts...)
			var s GetAppSecSelectableHostnamesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppSecSelectableHostnamesResultOutput)
}

// A collection of arguments for invoking getAppSecSelectableHostnames.
type GetAppSecSelectableHostnamesOutputArgs struct {
	ActiveInProduction pulumi.BoolPtrInput `pulumi:"activeInProduction"`
	ActiveInStaging    pulumi.BoolPtrInput `pulumi:"activeInStaging"`
	// . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can't be used with either the `contractid` or the `groupid` arguments.
	ConfigId pulumi.IntPtrInput `pulumi:"configId"`
	// . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can't be used with the `configId` argument.
	Contractid pulumi.StringPtrInput `pulumi:"contractid"`
	// . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can't be used with the `configId` argument.
	Groupid pulumi.IntPtrInput `pulumi:"groupid"`
}

func (GetAppSecSelectableHostnamesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecSelectableHostnamesArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecSelectableHostnames.
type GetAppSecSelectableHostnamesResultOutput struct{ *pulumi.OutputState }

func (GetAppSecSelectableHostnamesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecSelectableHostnamesResult)(nil)).Elem()
}

func (o GetAppSecSelectableHostnamesResultOutput) ToGetAppSecSelectableHostnamesResultOutput() GetAppSecSelectableHostnamesResultOutput {
	return o
}

func (o GetAppSecSelectableHostnamesResultOutput) ToGetAppSecSelectableHostnamesResultOutputWithContext(ctx context.Context) GetAppSecSelectableHostnamesResultOutput {
	return o
}

func (o GetAppSecSelectableHostnamesResultOutput) ActiveInProduction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) *bool { return v.ActiveInProduction }).(pulumi.BoolPtrOutput)
}

func (o GetAppSecSelectableHostnamesResultOutput) ActiveInStaging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) *bool { return v.ActiveInStaging }).(pulumi.BoolPtrOutput)
}

func (o GetAppSecSelectableHostnamesResultOutput) ConfigId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) *int { return v.ConfigId }).(pulumi.IntPtrOutput)
}

func (o GetAppSecSelectableHostnamesResultOutput) Contractid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) *string { return v.Contractid }).(pulumi.StringPtrOutput)
}

func (o GetAppSecSelectableHostnamesResultOutput) Groupid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) *int { return v.Groupid }).(pulumi.IntPtrOutput)
}

func (o GetAppSecSelectableHostnamesResultOutput) Hostnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) []string { return v.Hostnames }).(pulumi.StringArrayOutput)
}

func (o GetAppSecSelectableHostnamesResultOutput) HostnamesJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) string { return v.HostnamesJson }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecSelectableHostnamesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecSelectableHostnamesResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecSelectableHostnamesResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecSelectableHostnamesResultOutput{})
}
