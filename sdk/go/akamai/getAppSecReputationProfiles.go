// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration; reputation profile
//
// Returns information about your reputation profiles. Reputation profiles grade the security risk of an IP address based on previous activities associated with that address. Depending on the reputation score, and depending on how your configuration has been set up, requests from a specific IP address can trigger an alert, or even be blocked.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/reputation-profiles](https://techdocs.akamai.com/application-security/reference/get-reputation-profiles)
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
//			configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
//				Name: pulumi.StringRef("Documentation"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			reputationProfiles, err := akamai.GetAppSecReputationProfiles(ctx, &GetAppSecReputationProfilesArgs{
//				ConfigId: configuration.ConfigId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("reputationProfilesOutput", reputationProfiles.OutputText)
//			ctx.Export("reputationProfilesJson", reputationProfiles.Json)
//			reputationProfile, err := akamai.GetAppSecReputationProfiles(ctx, &GetAppSecReputationProfilesArgs{
//				ConfigId:            configuration.ConfigId,
//				ReputationProfileId: pulumi.IntRef(12345),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("reputationProfileJson", reputationProfile.Json)
//			ctx.Export("reputationProfileOutput", reputationProfile.OutputText)
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `outputText`. Tabular report of the details about the specified reputation profile or profiles.
// - `json`. JSON-formatted report of the details about the specified reputation profile or profiles.
func GetAppSecReputationProfiles(ctx *pulumi.Context, args *GetAppSecReputationProfilesArgs, opts ...pulumi.InvokeOption) (*GetAppSecReputationProfilesResult, error) {
	var rv GetAppSecReputationProfilesResult
	err := ctx.Invoke("akamai:index/getAppSecReputationProfiles:getAppSecReputationProfiles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecReputationProfiles.
type GetAppSecReputationProfilesArgs struct {
	// . Unique identifier of the security configuration associated with the reputation profiles.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the reputation profile you want to return information for. If not included, information is returned for all your reputation profiles.
	ReputationProfileId *int `pulumi:"reputationProfileId"`
}

// A collection of values returned by getAppSecReputationProfiles.
type GetAppSecReputationProfilesResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string `pulumi:"id"`
	Json                string `pulumi:"json"`
	OutputText          string `pulumi:"outputText"`
	ReputationProfileId *int   `pulumi:"reputationProfileId"`
}

func GetAppSecReputationProfilesOutput(ctx *pulumi.Context, args GetAppSecReputationProfilesOutputArgs, opts ...pulumi.InvokeOption) GetAppSecReputationProfilesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecReputationProfilesResult, error) {
			args := v.(GetAppSecReputationProfilesArgs)
			r, err := GetAppSecReputationProfiles(ctx, &args, opts...)
			var s GetAppSecReputationProfilesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppSecReputationProfilesResultOutput)
}

// A collection of arguments for invoking getAppSecReputationProfiles.
type GetAppSecReputationProfilesOutputArgs struct {
	// . Unique identifier of the security configuration associated with the reputation profiles.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . Unique identifier of the reputation profile you want to return information for. If not included, information is returned for all your reputation profiles.
	ReputationProfileId pulumi.IntPtrInput `pulumi:"reputationProfileId"`
}

func (GetAppSecReputationProfilesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecReputationProfilesArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecReputationProfiles.
type GetAppSecReputationProfilesResultOutput struct{ *pulumi.OutputState }

func (GetAppSecReputationProfilesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecReputationProfilesResult)(nil)).Elem()
}

func (o GetAppSecReputationProfilesResultOutput) ToGetAppSecReputationProfilesResultOutput() GetAppSecReputationProfilesResultOutput {
	return o
}

func (o GetAppSecReputationProfilesResultOutput) ToGetAppSecReputationProfilesResultOutputWithContext(ctx context.Context) GetAppSecReputationProfilesResultOutput {
	return o
}

func (o GetAppSecReputationProfilesResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecReputationProfilesResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecReputationProfilesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecReputationProfilesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecReputationProfilesResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecReputationProfilesResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetAppSecReputationProfilesResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecReputationProfilesResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecReputationProfilesResultOutput) ReputationProfileId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAppSecReputationProfilesResult) *int { return v.ReputationProfileId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecReputationProfilesResultOutput{})
}
