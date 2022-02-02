// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration; rate policy
//
// Creates, modifies or deletes rate policies.
// Rate polices help you monitor and moderate the number and  rate of all the requests you receive.
// In turn, this helps you prevent your website from being overwhelmed by a dramatic and unexpected surge in traffic.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/rate-policies](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postratepolicies)
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-akamai/sdk/v2/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
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
// 		ratePolicy, err := akamai.NewAppSecRatePolicy(ctx, "ratePolicy", &akamai.AppSecRatePolicyArgs{
// 			ConfigId:   pulumi.Int(configuration.ConfigId),
// 			RatePolicy: readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/rate_policy.json")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ratePolicyId", ratePolicy.RatePolicyId)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `ratePolicyId`. ID of the modified or newly-created rate policy.
type AppSecRatePolicy struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the rate policy being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Path to a JSON file containing a rate policy definition. You can view a sample rate policy JSON file in the [RatePolicy](https://developer.akamai.com/api/cloud_security/application_security/v1.html#ratepolicy) section of the Application Security API documentation.
	RatePolicy pulumi.StringOutput `pulumi:"ratePolicy"`
	// . Unique identifier of an existing rate policy.
	RatePolicyId pulumi.IntOutput `pulumi:"ratePolicyId"`
}

// NewAppSecRatePolicy registers a new resource with the given unique name, arguments, and options.
func NewAppSecRatePolicy(ctx *pulumi.Context,
	name string, args *AppSecRatePolicyArgs, opts ...pulumi.ResourceOption) (*AppSecRatePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.RatePolicy == nil {
		return nil, errors.New("invalid value for required argument 'RatePolicy'")
	}
	var resource AppSecRatePolicy
	err := ctx.RegisterResource("akamai:index/appSecRatePolicy:AppSecRatePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecRatePolicy gets an existing AppSecRatePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecRatePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecRatePolicyState, opts ...pulumi.ResourceOption) (*AppSecRatePolicy, error) {
	var resource AppSecRatePolicy
	err := ctx.ReadResource("akamai:index/appSecRatePolicy:AppSecRatePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecRatePolicy resources.
type appSecRatePolicyState struct {
	// . Unique identifier of the security configuration associated with the rate policy being modified.
	ConfigId *int `pulumi:"configId"`
	// . Path to a JSON file containing a rate policy definition. You can view a sample rate policy JSON file in the [RatePolicy](https://developer.akamai.com/api/cloud_security/application_security/v1.html#ratepolicy) section of the Application Security API documentation.
	RatePolicy *string `pulumi:"ratePolicy"`
	// . Unique identifier of an existing rate policy.
	RatePolicyId *int `pulumi:"ratePolicyId"`
}

type AppSecRatePolicyState struct {
	// . Unique identifier of the security configuration associated with the rate policy being modified.
	ConfigId pulumi.IntPtrInput
	// . Path to a JSON file containing a rate policy definition. You can view a sample rate policy JSON file in the [RatePolicy](https://developer.akamai.com/api/cloud_security/application_security/v1.html#ratepolicy) section of the Application Security API documentation.
	RatePolicy pulumi.StringPtrInput
	// . Unique identifier of an existing rate policy.
	RatePolicyId pulumi.IntPtrInput
}

func (AppSecRatePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRatePolicyState)(nil)).Elem()
}

type appSecRatePolicyArgs struct {
	// . Unique identifier of the security configuration associated with the rate policy being modified.
	ConfigId int `pulumi:"configId"`
	// . Path to a JSON file containing a rate policy definition. You can view a sample rate policy JSON file in the [RatePolicy](https://developer.akamai.com/api/cloud_security/application_security/v1.html#ratepolicy) section of the Application Security API documentation.
	RatePolicy string `pulumi:"ratePolicy"`
}

// The set of arguments for constructing a AppSecRatePolicy resource.
type AppSecRatePolicyArgs struct {
	// . Unique identifier of the security configuration associated with the rate policy being modified.
	ConfigId pulumi.IntInput
	// . Path to a JSON file containing a rate policy definition. You can view a sample rate policy JSON file in the [RatePolicy](https://developer.akamai.com/api/cloud_security/application_security/v1.html#ratepolicy) section of the Application Security API documentation.
	RatePolicy pulumi.StringInput
}

func (AppSecRatePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRatePolicyArgs)(nil)).Elem()
}

type AppSecRatePolicyInput interface {
	pulumi.Input

	ToAppSecRatePolicyOutput() AppSecRatePolicyOutput
	ToAppSecRatePolicyOutputWithContext(ctx context.Context) AppSecRatePolicyOutput
}

func (*AppSecRatePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRatePolicy)(nil)).Elem()
}

func (i *AppSecRatePolicy) ToAppSecRatePolicyOutput() AppSecRatePolicyOutput {
	return i.ToAppSecRatePolicyOutputWithContext(context.Background())
}

func (i *AppSecRatePolicy) ToAppSecRatePolicyOutputWithContext(ctx context.Context) AppSecRatePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyOutput)
}

// AppSecRatePolicyArrayInput is an input type that accepts AppSecRatePolicyArray and AppSecRatePolicyArrayOutput values.
// You can construct a concrete instance of `AppSecRatePolicyArrayInput` via:
//
//          AppSecRatePolicyArray{ AppSecRatePolicyArgs{...} }
type AppSecRatePolicyArrayInput interface {
	pulumi.Input

	ToAppSecRatePolicyArrayOutput() AppSecRatePolicyArrayOutput
	ToAppSecRatePolicyArrayOutputWithContext(context.Context) AppSecRatePolicyArrayOutput
}

type AppSecRatePolicyArray []AppSecRatePolicyInput

func (AppSecRatePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecRatePolicy)(nil)).Elem()
}

func (i AppSecRatePolicyArray) ToAppSecRatePolicyArrayOutput() AppSecRatePolicyArrayOutput {
	return i.ToAppSecRatePolicyArrayOutputWithContext(context.Background())
}

func (i AppSecRatePolicyArray) ToAppSecRatePolicyArrayOutputWithContext(ctx context.Context) AppSecRatePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyArrayOutput)
}

// AppSecRatePolicyMapInput is an input type that accepts AppSecRatePolicyMap and AppSecRatePolicyMapOutput values.
// You can construct a concrete instance of `AppSecRatePolicyMapInput` via:
//
//          AppSecRatePolicyMap{ "key": AppSecRatePolicyArgs{...} }
type AppSecRatePolicyMapInput interface {
	pulumi.Input

	ToAppSecRatePolicyMapOutput() AppSecRatePolicyMapOutput
	ToAppSecRatePolicyMapOutputWithContext(context.Context) AppSecRatePolicyMapOutput
}

type AppSecRatePolicyMap map[string]AppSecRatePolicyInput

func (AppSecRatePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecRatePolicy)(nil)).Elem()
}

func (i AppSecRatePolicyMap) ToAppSecRatePolicyMapOutput() AppSecRatePolicyMapOutput {
	return i.ToAppSecRatePolicyMapOutputWithContext(context.Background())
}

func (i AppSecRatePolicyMap) ToAppSecRatePolicyMapOutputWithContext(ctx context.Context) AppSecRatePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyMapOutput)
}

type AppSecRatePolicyOutput struct{ *pulumi.OutputState }

func (AppSecRatePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRatePolicy)(nil)).Elem()
}

func (o AppSecRatePolicyOutput) ToAppSecRatePolicyOutput() AppSecRatePolicyOutput {
	return o
}

func (o AppSecRatePolicyOutput) ToAppSecRatePolicyOutputWithContext(ctx context.Context) AppSecRatePolicyOutput {
	return o
}

type AppSecRatePolicyArrayOutput struct{ *pulumi.OutputState }

func (AppSecRatePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecRatePolicy)(nil)).Elem()
}

func (o AppSecRatePolicyArrayOutput) ToAppSecRatePolicyArrayOutput() AppSecRatePolicyArrayOutput {
	return o
}

func (o AppSecRatePolicyArrayOutput) ToAppSecRatePolicyArrayOutputWithContext(ctx context.Context) AppSecRatePolicyArrayOutput {
	return o
}

func (o AppSecRatePolicyArrayOutput) Index(i pulumi.IntInput) AppSecRatePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecRatePolicy {
		return vs[0].([]*AppSecRatePolicy)[vs[1].(int)]
	}).(AppSecRatePolicyOutput)
}

type AppSecRatePolicyMapOutput struct{ *pulumi.OutputState }

func (AppSecRatePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecRatePolicy)(nil)).Elem()
}

func (o AppSecRatePolicyMapOutput) ToAppSecRatePolicyMapOutput() AppSecRatePolicyMapOutput {
	return o
}

func (o AppSecRatePolicyMapOutput) ToAppSecRatePolicyMapOutputWithContext(ctx context.Context) AppSecRatePolicyMapOutput {
	return o
}

func (o AppSecRatePolicyMapOutput) MapIndex(k pulumi.StringInput) AppSecRatePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecRatePolicy {
		return vs[0].(map[string]*AppSecRatePolicy)[vs[1].(string)]
	}).(AppSecRatePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRatePolicyInput)(nil)).Elem(), &AppSecRatePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRatePolicyArrayInput)(nil)).Elem(), AppSecRatePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRatePolicyMapInput)(nil)).Elem(), AppSecRatePolicyMap{})
	pulumi.RegisterOutputType(AppSecRatePolicyOutput{})
	pulumi.RegisterOutputType(AppSecRatePolicyArrayOutput{})
	pulumi.RegisterOutputType(AppSecRatePolicyMapOutput{})
}
