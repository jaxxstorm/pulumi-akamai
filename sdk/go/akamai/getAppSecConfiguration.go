// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppSecConfiguration(ctx *pulumi.Context, args *LookupAppSecConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAppSecConfigurationResult, error) {
	var rv LookupAppSecConfigurationResult
	err := ctx.Invoke("akamai:index/getAppSecConfiguration:getAppSecConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecConfiguration.
type LookupAppSecConfigurationArgs struct {
	// . Name of the security configuration you want to return information for. If not included, information is returned for all your security configurations.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAppSecConfiguration.
type LookupAppSecConfigurationResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id                string  `pulumi:"id"`
	LatestVersion     int     `pulumi:"latestVersion"`
	Name              *string `pulumi:"name"`
	OutputText        string  `pulumi:"outputText"`
	ProductionVersion int     `pulumi:"productionVersion"`
	StagingVersion    int     `pulumi:"stagingVersion"`
}

func LookupAppSecConfigurationOutput(ctx *pulumi.Context, args LookupAppSecConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupAppSecConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppSecConfigurationResult, error) {
			args := v.(LookupAppSecConfigurationArgs)
			r, err := LookupAppSecConfiguration(ctx, &args, opts...)
			return *r, err
		}).(LookupAppSecConfigurationResultOutput)
}

// A collection of arguments for invoking getAppSecConfiguration.
type LookupAppSecConfigurationOutputArgs struct {
	// . Name of the security configuration you want to return information for. If not included, information is returned for all your security configurations.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAppSecConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecConfigurationArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecConfiguration.
type LookupAppSecConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupAppSecConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppSecConfigurationResult)(nil)).Elem()
}

func (o LookupAppSecConfigurationResultOutput) ToLookupAppSecConfigurationResultOutput() LookupAppSecConfigurationResultOutput {
	return o
}

func (o LookupAppSecConfigurationResultOutput) ToLookupAppSecConfigurationResultOutputWithContext(ctx context.Context) LookupAppSecConfigurationResultOutput {
	return o
}

func (o LookupAppSecConfigurationResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecConfigurationResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAppSecConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppSecConfigurationResultOutput) LatestVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecConfigurationResult) int { return v.LatestVersion }).(pulumi.IntOutput)
}

func (o LookupAppSecConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppSecConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAppSecConfigurationResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppSecConfigurationResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func (o LookupAppSecConfigurationResultOutput) ProductionVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecConfigurationResult) int { return v.ProductionVersion }).(pulumi.IntOutput)
}

func (o LookupAppSecConfigurationResultOutput) StagingVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppSecConfigurationResult) int { return v.StagingVersion }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppSecConfigurationResultOutput{})
}
