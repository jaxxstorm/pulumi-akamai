// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `resourceAkamaiAppsecReputationProfileAnalysis` resource allows you to toggle the reputation analysis settings for a given security policy. The `forwardToHttpHeader` parameter indicates whether to add client reputation details to requests forwarded to origin in an HTTP header. The `forwardSharedIpToHttpHeaderSiem` parameter indicates whether to add value indicating that shared IPs are included in HTTP header and SIEM integration.
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
// 		_, err = akamai.NewAppSecReputationProfileAnalysis(ctx, "reputationAnalysis", &akamai.AppSecReputationProfileAnalysisArgs{
// 			ConfigId:                        pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId:                pulumi.Any(_var.Security_policy_id),
// 			ForwardToHttpHeader:             pulumi.Bool(true),
// 			ForwardSharedIpToHttpHeaderSiem: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecReputationProfileAnalysis struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// Whether to add value indicating that shared IPs are included in HTTP header and SIEM integration.
	ForwardSharedIpToHttpHeaderSiem pulumi.BoolOutput `pulumi:"forwardSharedIpToHttpHeaderSiem"`
	// Whether to add client reputation details to requests forwarded to origin in an HTTP header.
	ForwardToHttpHeader pulumi.BoolOutput `pulumi:"forwardToHttpHeader"`
	// The ID of the securityPolicyId to which the settings should be applied.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecReputationProfileAnalysis registers a new resource with the given unique name, arguments, and options.
func NewAppSecReputationProfileAnalysis(ctx *pulumi.Context,
	name string, args *AppSecReputationProfileAnalysisArgs, opts ...pulumi.ResourceOption) (*AppSecReputationProfileAnalysis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.ForwardSharedIpToHttpHeaderSiem == nil {
		return nil, errors.New("invalid value for required argument 'ForwardSharedIpToHttpHeaderSiem'")
	}
	if args.ForwardToHttpHeader == nil {
		return nil, errors.New("invalid value for required argument 'ForwardToHttpHeader'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	var resource AppSecReputationProfileAnalysis
	err := ctx.RegisterResource("akamai:index/appSecReputationProfileAnalysis:AppSecReputationProfileAnalysis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecReputationProfileAnalysis gets an existing AppSecReputationProfileAnalysis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecReputationProfileAnalysis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecReputationProfileAnalysisState, opts ...pulumi.ResourceOption) (*AppSecReputationProfileAnalysis, error) {
	var resource AppSecReputationProfileAnalysis
	err := ctx.ReadResource("akamai:index/appSecReputationProfileAnalysis:AppSecReputationProfileAnalysis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecReputationProfileAnalysis resources.
type appSecReputationProfileAnalysisState struct {
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// Whether to add value indicating that shared IPs are included in HTTP header and SIEM integration.
	ForwardSharedIpToHttpHeaderSiem *bool `pulumi:"forwardSharedIpToHttpHeaderSiem"`
	// Whether to add client reputation details to requests forwarded to origin in an HTTP header.
	ForwardToHttpHeader *bool `pulumi:"forwardToHttpHeader"`
	// The ID of the securityPolicyId to which the settings should be applied.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecReputationProfileAnalysisState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// Whether to add value indicating that shared IPs are included in HTTP header and SIEM integration.
	ForwardSharedIpToHttpHeaderSiem pulumi.BoolPtrInput
	// Whether to add client reputation details to requests forwarded to origin in an HTTP header.
	ForwardToHttpHeader pulumi.BoolPtrInput
	// The ID of the securityPolicyId to which the settings should be applied.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecReputationProfileAnalysisState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecReputationProfileAnalysisState)(nil)).Elem()
}

type appSecReputationProfileAnalysisArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// Whether to add value indicating that shared IPs are included in HTTP header and SIEM integration.
	ForwardSharedIpToHttpHeaderSiem bool `pulumi:"forwardSharedIpToHttpHeaderSiem"`
	// Whether to add client reputation details to requests forwarded to origin in an HTTP header.
	ForwardToHttpHeader bool `pulumi:"forwardToHttpHeader"`
	// The ID of the securityPolicyId to which the settings should be applied.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecReputationProfileAnalysis resource.
type AppSecReputationProfileAnalysisArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// Whether to add value indicating that shared IPs are included in HTTP header and SIEM integration.
	ForwardSharedIpToHttpHeaderSiem pulumi.BoolInput
	// Whether to add client reputation details to requests forwarded to origin in an HTTP header.
	ForwardToHttpHeader pulumi.BoolInput
	// The ID of the securityPolicyId to which the settings should be applied.
	SecurityPolicyId pulumi.StringInput
}

func (AppSecReputationProfileAnalysisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecReputationProfileAnalysisArgs)(nil)).Elem()
}

type AppSecReputationProfileAnalysisInput interface {
	pulumi.Input

	ToAppSecReputationProfileAnalysisOutput() AppSecReputationProfileAnalysisOutput
	ToAppSecReputationProfileAnalysisOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisOutput
}

func (*AppSecReputationProfileAnalysis) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecReputationProfileAnalysis)(nil))
}

func (i *AppSecReputationProfileAnalysis) ToAppSecReputationProfileAnalysisOutput() AppSecReputationProfileAnalysisOutput {
	return i.ToAppSecReputationProfileAnalysisOutputWithContext(context.Background())
}

func (i *AppSecReputationProfileAnalysis) ToAppSecReputationProfileAnalysisOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProfileAnalysisOutput)
}

func (i *AppSecReputationProfileAnalysis) ToAppSecReputationProfileAnalysisPtrOutput() AppSecReputationProfileAnalysisPtrOutput {
	return i.ToAppSecReputationProfileAnalysisPtrOutputWithContext(context.Background())
}

func (i *AppSecReputationProfileAnalysis) ToAppSecReputationProfileAnalysisPtrOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProfileAnalysisPtrOutput)
}

type AppSecReputationProfileAnalysisPtrInput interface {
	pulumi.Input

	ToAppSecReputationProfileAnalysisPtrOutput() AppSecReputationProfileAnalysisPtrOutput
	ToAppSecReputationProfileAnalysisPtrOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisPtrOutput
}

type appSecReputationProfileAnalysisPtrType AppSecReputationProfileAnalysisArgs

func (*appSecReputationProfileAnalysisPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecReputationProfileAnalysis)(nil))
}

func (i *appSecReputationProfileAnalysisPtrType) ToAppSecReputationProfileAnalysisPtrOutput() AppSecReputationProfileAnalysisPtrOutput {
	return i.ToAppSecReputationProfileAnalysisPtrOutputWithContext(context.Background())
}

func (i *appSecReputationProfileAnalysisPtrType) ToAppSecReputationProfileAnalysisPtrOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProfileAnalysisPtrOutput)
}

// AppSecReputationProfileAnalysisArrayInput is an input type that accepts AppSecReputationProfileAnalysisArray and AppSecReputationProfileAnalysisArrayOutput values.
// You can construct a concrete instance of `AppSecReputationProfileAnalysisArrayInput` via:
//
//          AppSecReputationProfileAnalysisArray{ AppSecReputationProfileAnalysisArgs{...} }
type AppSecReputationProfileAnalysisArrayInput interface {
	pulumi.Input

	ToAppSecReputationProfileAnalysisArrayOutput() AppSecReputationProfileAnalysisArrayOutput
	ToAppSecReputationProfileAnalysisArrayOutputWithContext(context.Context) AppSecReputationProfileAnalysisArrayOutput
}

type AppSecReputationProfileAnalysisArray []AppSecReputationProfileAnalysisInput

func (AppSecReputationProfileAnalysisArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecReputationProfileAnalysis)(nil)).Elem()
}

func (i AppSecReputationProfileAnalysisArray) ToAppSecReputationProfileAnalysisArrayOutput() AppSecReputationProfileAnalysisArrayOutput {
	return i.ToAppSecReputationProfileAnalysisArrayOutputWithContext(context.Background())
}

func (i AppSecReputationProfileAnalysisArray) ToAppSecReputationProfileAnalysisArrayOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProfileAnalysisArrayOutput)
}

// AppSecReputationProfileAnalysisMapInput is an input type that accepts AppSecReputationProfileAnalysisMap and AppSecReputationProfileAnalysisMapOutput values.
// You can construct a concrete instance of `AppSecReputationProfileAnalysisMapInput` via:
//
//          AppSecReputationProfileAnalysisMap{ "key": AppSecReputationProfileAnalysisArgs{...} }
type AppSecReputationProfileAnalysisMapInput interface {
	pulumi.Input

	ToAppSecReputationProfileAnalysisMapOutput() AppSecReputationProfileAnalysisMapOutput
	ToAppSecReputationProfileAnalysisMapOutputWithContext(context.Context) AppSecReputationProfileAnalysisMapOutput
}

type AppSecReputationProfileAnalysisMap map[string]AppSecReputationProfileAnalysisInput

func (AppSecReputationProfileAnalysisMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecReputationProfileAnalysis)(nil)).Elem()
}

func (i AppSecReputationProfileAnalysisMap) ToAppSecReputationProfileAnalysisMapOutput() AppSecReputationProfileAnalysisMapOutput {
	return i.ToAppSecReputationProfileAnalysisMapOutputWithContext(context.Background())
}

func (i AppSecReputationProfileAnalysisMap) ToAppSecReputationProfileAnalysisMapOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProfileAnalysisMapOutput)
}

type AppSecReputationProfileAnalysisOutput struct{ *pulumi.OutputState }

func (AppSecReputationProfileAnalysisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecReputationProfileAnalysis)(nil))
}

func (o AppSecReputationProfileAnalysisOutput) ToAppSecReputationProfileAnalysisOutput() AppSecReputationProfileAnalysisOutput {
	return o
}

func (o AppSecReputationProfileAnalysisOutput) ToAppSecReputationProfileAnalysisOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisOutput {
	return o
}

func (o AppSecReputationProfileAnalysisOutput) ToAppSecReputationProfileAnalysisPtrOutput() AppSecReputationProfileAnalysisPtrOutput {
	return o.ToAppSecReputationProfileAnalysisPtrOutputWithContext(context.Background())
}

func (o AppSecReputationProfileAnalysisOutput) ToAppSecReputationProfileAnalysisPtrOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSecReputationProfileAnalysis) *AppSecReputationProfileAnalysis {
		return &v
	}).(AppSecReputationProfileAnalysisPtrOutput)
}

type AppSecReputationProfileAnalysisPtrOutput struct{ *pulumi.OutputState }

func (AppSecReputationProfileAnalysisPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecReputationProfileAnalysis)(nil))
}

func (o AppSecReputationProfileAnalysisPtrOutput) ToAppSecReputationProfileAnalysisPtrOutput() AppSecReputationProfileAnalysisPtrOutput {
	return o
}

func (o AppSecReputationProfileAnalysisPtrOutput) ToAppSecReputationProfileAnalysisPtrOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisPtrOutput {
	return o
}

func (o AppSecReputationProfileAnalysisPtrOutput) Elem() AppSecReputationProfileAnalysisOutput {
	return o.ApplyT(func(v *AppSecReputationProfileAnalysis) AppSecReputationProfileAnalysis {
		if v != nil {
			return *v
		}
		var ret AppSecReputationProfileAnalysis
		return ret
	}).(AppSecReputationProfileAnalysisOutput)
}

type AppSecReputationProfileAnalysisArrayOutput struct{ *pulumi.OutputState }

func (AppSecReputationProfileAnalysisArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecReputationProfileAnalysis)(nil))
}

func (o AppSecReputationProfileAnalysisArrayOutput) ToAppSecReputationProfileAnalysisArrayOutput() AppSecReputationProfileAnalysisArrayOutput {
	return o
}

func (o AppSecReputationProfileAnalysisArrayOutput) ToAppSecReputationProfileAnalysisArrayOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisArrayOutput {
	return o
}

func (o AppSecReputationProfileAnalysisArrayOutput) Index(i pulumi.IntInput) AppSecReputationProfileAnalysisOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecReputationProfileAnalysis {
		return vs[0].([]AppSecReputationProfileAnalysis)[vs[1].(int)]
	}).(AppSecReputationProfileAnalysisOutput)
}

type AppSecReputationProfileAnalysisMapOutput struct{ *pulumi.OutputState }

func (AppSecReputationProfileAnalysisMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecReputationProfileAnalysis)(nil))
}

func (o AppSecReputationProfileAnalysisMapOutput) ToAppSecReputationProfileAnalysisMapOutput() AppSecReputationProfileAnalysisMapOutput {
	return o
}

func (o AppSecReputationProfileAnalysisMapOutput) ToAppSecReputationProfileAnalysisMapOutputWithContext(ctx context.Context) AppSecReputationProfileAnalysisMapOutput {
	return o
}

func (o AppSecReputationProfileAnalysisMapOutput) MapIndex(k pulumi.StringInput) AppSecReputationProfileAnalysisOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecReputationProfileAnalysis {
		return vs[0].(map[string]AppSecReputationProfileAnalysis)[vs[1].(string)]
	}).(AppSecReputationProfileAnalysisOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProfileAnalysisInput)(nil)).Elem(), &AppSecReputationProfileAnalysis{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProfileAnalysisPtrInput)(nil)).Elem(), &AppSecReputationProfileAnalysis{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProfileAnalysisArrayInput)(nil)).Elem(), AppSecReputationProfileAnalysisArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProfileAnalysisMapInput)(nil)).Elem(), AppSecReputationProfileAnalysisMap{})
	pulumi.RegisterOutputType(AppSecReputationProfileAnalysisOutput{})
	pulumi.RegisterOutputType(AppSecReputationProfileAnalysisPtrOutput{})
	pulumi.RegisterOutputType(AppSecReputationProfileAnalysisArrayOutput{})
	pulumi.RegisterOutputType(AppSecReputationProfileAnalysisMapOutput{})
}
