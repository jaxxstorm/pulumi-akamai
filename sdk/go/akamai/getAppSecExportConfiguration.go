// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration and version
//
// Returns comprehensive details about a security configuration, including rate policies, security policies, rules, hostnames, and match targets.
//
// **Related API Endpoint**: [/appsec/v1/export/configs/{configId}/versions/{versionNumber}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getconfigurationversionexport)
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
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: pulumi.StringRef("Documentation"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		export, err := akamai.GetAppSecExportConfiguration(ctx, &GetAppSecExportConfigurationArgs{
// 			ConfigId: configuration.ConfigId,
// 			Version:  configuration.LatestVersion,
// 			Searches: []string{
// 				"securityPolicies",
// 				"selectedHosts",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("json", export.Json)
// 		ctx.Export("text", export.OutputText)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `json`. Complete set of information about the specified security configuration version in JSON format. Includes the types available for the `search` parameter as well as additional fields such as `createDate` and `createdBy`.
// - `outputText`. Tabular report showing the types of data specified in the `search` parameter. Valid only if the `search` parameter references at least one type.
func GetAppSecExportConfiguration(ctx *pulumi.Context, args *GetAppSecExportConfigurationArgs, opts ...pulumi.InvokeOption) (*GetAppSecExportConfigurationResult, error) {
	var rv GetAppSecExportConfigurationResult
	err := ctx.Invoke("akamai:index/getAppSecExportConfiguration:getAppSecExportConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecExportConfiguration.
type GetAppSecExportConfigurationArgs struct {
	// . Unique identifier of the security configuration you want to return information for.
	ConfigId int `pulumi:"configId"`
	// . JSON array of strings specifying the types of information to be retrieved. Allowed values include:
	// > - **AdvancedSettingsLogging**
	// > - **AdvancedSettingsPrefetch**
	// > - **ApiRequestConstraints**
	// > - **AttackGroup**
	// > - **AttackGroupConditionException**
	// > - **Eval**
	// > - **EvalRuleConditionException**
	// > - **CustomDeny**
	// > - **CustomRule**
	// > - **CustomRuleAction**
	// > - **IPGeoFirewall**
	// > - **MatchTarget**
	// > - **PenaltyBox**
	// > - **RatePolicy**
	// > - **RatePolicyAction**
	// > - **ReputationProfile**
	// > - **ReputationProfileAction**
	// > - **Rule**
	// > - **RuleConditionException**
	// > - **SecurityPolicy**
	// > - **SiemSettings**
	// > - **SlowPost**
	Searches []string `pulumi:"searches"`
	// . Version number of the security configuration.
	Version int `pulumi:"version"`
}

// A collection of values returned by getAppSecExportConfiguration.
type GetAppSecExportConfigurationResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Json       string   `pulumi:"json"`
	OutputText string   `pulumi:"outputText"`
	Searches   []string `pulumi:"searches"`
	Version    int      `pulumi:"version"`
}

func GetAppSecExportConfigurationOutput(ctx *pulumi.Context, args GetAppSecExportConfigurationOutputArgs, opts ...pulumi.InvokeOption) GetAppSecExportConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecExportConfigurationResult, error) {
			args := v.(GetAppSecExportConfigurationArgs)
			r, err := GetAppSecExportConfiguration(ctx, &args, opts...)
			var s GetAppSecExportConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppSecExportConfigurationResultOutput)
}

// A collection of arguments for invoking getAppSecExportConfiguration.
type GetAppSecExportConfigurationOutputArgs struct {
	// . Unique identifier of the security configuration you want to return information for.
	ConfigId pulumi.IntInput `pulumi:"configId"`
	// . JSON array of strings specifying the types of information to be retrieved. Allowed values include:
	// > - **AdvancedSettingsLogging**
	// > - **AdvancedSettingsPrefetch**
	// > - **ApiRequestConstraints**
	// > - **AttackGroup**
	// > - **AttackGroupConditionException**
	// > - **Eval**
	// > - **EvalRuleConditionException**
	// > - **CustomDeny**
	// > - **CustomRule**
	// > - **CustomRuleAction**
	// > - **IPGeoFirewall**
	// > - **MatchTarget**
	// > - **PenaltyBox**
	// > - **RatePolicy**
	// > - **RatePolicyAction**
	// > - **ReputationProfile**
	// > - **ReputationProfileAction**
	// > - **Rule**
	// > - **RuleConditionException**
	// > - **SecurityPolicy**
	// > - **SiemSettings**
	// > - **SlowPost**
	Searches pulumi.StringArrayInput `pulumi:"searches"`
	// . Version number of the security configuration.
	Version pulumi.IntInput `pulumi:"version"`
}

func (GetAppSecExportConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecExportConfigurationArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecExportConfiguration.
type GetAppSecExportConfigurationResultOutput struct{ *pulumi.OutputState }

func (GetAppSecExportConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecExportConfigurationResult)(nil)).Elem()
}

func (o GetAppSecExportConfigurationResultOutput) ToGetAppSecExportConfigurationResultOutput() GetAppSecExportConfigurationResultOutput {
	return o
}

func (o GetAppSecExportConfigurationResultOutput) ToGetAppSecExportConfigurationResultOutputWithContext(ctx context.Context) GetAppSecExportConfigurationResultOutput {
	return o
}

func (o GetAppSecExportConfigurationResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecExportConfigurationResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecExportConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecExportConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppSecExportConfigurationResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecExportConfigurationResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetAppSecExportConfigurationResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecExportConfigurationResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o GetAppSecExportConfigurationResultOutput) Searches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppSecExportConfigurationResult) []string { return v.Searches }).(pulumi.StringArrayOutput)
}

func (o GetAppSecExportConfigurationResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecExportConfigurationResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecExportConfigurationResultOutput{})
}
