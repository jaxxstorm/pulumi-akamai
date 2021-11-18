// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecVersionNodes` data source to retrieve the most recent version notes for a configuration.
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
// 		opt0 := _var.Security_configuration
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		versionNotes, err := akamai.GetAppSecVersionNotes(ctx, &GetAppSecVersionNotesArgs{
// 			ConfigId: configuration.ConfigId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("versionNotesText", versionNotes.OutputText)
// 		ctx.Export("versionNotesJson", versionNotes.Json)
// 		return nil
// 	})
// }
// ```
func GetAppSecVersionNotes(ctx *pulumi.Context, args *GetAppSecVersionNotesArgs, opts ...pulumi.InvokeOption) (*GetAppSecVersionNotesResult, error) {
	var rv GetAppSecVersionNotesResult
	err := ctx.Invoke("akamai:index/getAppSecVersionNotes:getAppSecVersionNotes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecVersionNotes.
type GetAppSecVersionNotesArgs struct {
	// The configuration ID to use.
	ConfigId int `pulumi:"configId"`
}

// A collection of values returned by getAppSecVersionNotes.
type GetAppSecVersionNotesResult struct {
	ConfigId int `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted list showing the version notes.
	Json string `pulumi:"json"`
	// A tabular display showing the version notes.
	OutputText string `pulumi:"outputText"`
}

func GetAppSecVersionNotesOutput(ctx *pulumi.Context, args GetAppSecVersionNotesOutputArgs, opts ...pulumi.InvokeOption) GetAppSecVersionNotesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppSecVersionNotesResult, error) {
			args := v.(GetAppSecVersionNotesArgs)
			r, err := GetAppSecVersionNotes(ctx, &args, opts...)
			return *r, err
		}).(GetAppSecVersionNotesResultOutput)
}

// A collection of arguments for invoking getAppSecVersionNotes.
type GetAppSecVersionNotesOutputArgs struct {
	// The configuration ID to use.
	ConfigId pulumi.IntInput `pulumi:"configId"`
}

func (GetAppSecVersionNotesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecVersionNotesArgs)(nil)).Elem()
}

// A collection of values returned by getAppSecVersionNotes.
type GetAppSecVersionNotesResultOutput struct{ *pulumi.OutputState }

func (GetAppSecVersionNotesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppSecVersionNotesResult)(nil)).Elem()
}

func (o GetAppSecVersionNotesResultOutput) ToGetAppSecVersionNotesResultOutput() GetAppSecVersionNotesResultOutput {
	return o
}

func (o GetAppSecVersionNotesResultOutput) ToGetAppSecVersionNotesResultOutputWithContext(ctx context.Context) GetAppSecVersionNotesResultOutput {
	return o
}

func (o GetAppSecVersionNotesResultOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v GetAppSecVersionNotesResult) int { return v.ConfigId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppSecVersionNotesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecVersionNotesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A JSON-formatted list showing the version notes.
func (o GetAppSecVersionNotesResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecVersionNotesResult) string { return v.Json }).(pulumi.StringOutput)
}

// A tabular display showing the version notes.
func (o GetAppSecVersionNotesResultOutput) OutputText() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppSecVersionNotesResult) string { return v.OutputText }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppSecVersionNotesResultOutput{})
}
