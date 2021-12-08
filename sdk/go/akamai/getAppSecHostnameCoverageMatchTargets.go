// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Hostname
//
// Returns information about the API and website match targets used to protect a hostname. The returned information is described in the [Get the hostname coverage match targets](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getfailoverhostnames) section of the Application Security API.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/hostname-coverage/match-targets?hostname={host}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoveragematchtargets)
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
// 		opt0 := "Documentation"
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.GetAppSecHostnameCoverageMatchTargets(ctx, &GetAppSecHostnameCoverageMatchTargetsArgs{
// 			ConfigId: configuration.ConfigId,
// 			Hostname: "documentation.akamai.com",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `json`. JSON-formatted list of the coverage information.
// - `outputText`. Tabular report of the coverage information.
func GetAppSecHostnameCoverageMatchTargets(ctx *pulumi.Context, args *GetAppSecHostnameCoverageMatchTargetsArgs, opts ...pulumi.InvokeOption) (*GetAppSecHostnameCoverageMatchTargetsResult, error) {
	var rv GetAppSecHostnameCoverageMatchTargetsResult
	err := ctx.Invoke("akamai:index/getAppSecHostnameCoverageMatchTargets:getAppSecHostnameCoverageMatchTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecHostnameCoverageMatchTargets.
type GetAppSecHostnameCoverageMatchTargetsArgs struct {
	ConfigId int `pulumi:"configId"`
	// . Name of the host you want to return information for. You can only return information for a single host and hostname at a time.
	Hostname string `pulumi:"hostname"`
}

// A collection of values returned by getAppSecHostnameCoverageMatchTargets.
type GetAppSecHostnameCoverageMatchTargetsResult struct {
	ConfigId int    `pulumi:"configId"`
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Json       string `pulumi:"json"`
	OutputText string `pulumi:"outputText"`
}

func GetAppSecHostnameCoverageMatchTargetsOutput(ctx *pulumi.Context, args GetAppSecHostnameCoverageMatchTargetsOutputArgs, opts ...pulumi.InvokeOption) GetAppSecHostnameCoverageMatchTargetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecHostnameCoverageMatchTargetsResult, error) {
			args := v.(GetAppSecHostnameCoverageMatchTargetsArgs)
			r, err := GetAppSecHostnameCoverageMatchTargets(ctx, &args, opts...)
			return *r, err
		}).(GetAppSecHostnameCoverageMatchTargetsResultOutput)
}

// A collection of arguments for invoking getAppSecHostnameCoverageMatchTargets.
type GetAppSecHostnameCoverageMatchTargetsOutputArgs struct {
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Name of the host you want to return information for. You can only return information for a single host and hostname at a time.
	Hostname pulumi.StringInput `pulumi:"hostname"`
}

func (GetAppSecHostnameCoverageMatchTargetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecHostnameCoverageMatchTargetsArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecHostnameCoverageMatchTargets.
type GetAppSecHostnameCoverageMatchTargetsResultOutput struct{ *pulumi.OutputState }

func (GetAppSecHostnameCoverageMatchTargetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecHostnameCoverageMatchTargetsResult)(nil)).Elem()
}

func (o GetAppSecHostnameCoverageMatchTargetsResultOutput) ToGetAppSecHostnameCoverageMatchTargetsResultOutput() GetAppSecHostnameCoverageMatchTargetsResultOutput {
	return o
}

func (o GetAppSecHostnameCoverageMatchTargetsResultOutput) ToGetAppSecHostnameCoverageMatchTargetsResultOutputWithContext(ctx context.Context) GetAppSecHostnameCoverageMatchTargetsResultOutput {
	return o
}

func (o GetAppSecHostnameCoverageMatchTargetsResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecHostnameCoverageMatchTargetsResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

func (o GetAppSecHostnameCoverageMatchTargetsResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecHostnameCoverageMatchTargetsResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecHostnameCoverageMatchTargetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecHostnameCoverageMatchTargetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecHostnameCoverageMatchTargetsResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecHostnameCoverageMatchTargetsResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetAppSecHostnameCoverageMatchTargetsResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecHostnameCoverageMatchTargetsResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecHostnameCoverageMatchTargetsResultOutput{})
}
