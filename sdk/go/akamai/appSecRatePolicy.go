// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `resourceAkamaiAppsecRatePolicy` resource allows you to create, modify or delete rate policies for a specific security configuration.
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
// 		opt0 := _var.Security_configuration
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
type AppSecRatePolicy struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// The name of a file containing a JSON-formatted rate policy definition ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#57c65cbd)).
	RatePolicy pulumi.StringOutput `pulumi:"ratePolicy"`
	// The ID of an existing rate policy to be modified.
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
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// The name of a file containing a JSON-formatted rate policy definition ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#57c65cbd)).
	RatePolicy *string `pulumi:"ratePolicy"`
	// The ID of an existing rate policy to be modified.
	RatePolicyId *int `pulumi:"ratePolicyId"`
}

type AppSecRatePolicyState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// The name of a file containing a JSON-formatted rate policy definition ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#57c65cbd)).
	RatePolicy pulumi.StringPtrInput
	// The ID of an existing rate policy to be modified.
	RatePolicyId pulumi.IntPtrInput
}

func (AppSecRatePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecRatePolicyState)(nil)).Elem()
}

type appSecRatePolicyArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The name of a file containing a JSON-formatted rate policy definition ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#57c65cbd)).
	RatePolicy string `pulumi:"ratePolicy"`
}

// The set of arguments for constructing a AppSecRatePolicy resource.
type AppSecRatePolicyArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// The name of a file containing a JSON-formatted rate policy definition ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#57c65cbd)).
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
	return reflect.TypeOf((*AppSecRatePolicy)(nil))
}

func (i *AppSecRatePolicy) ToAppSecRatePolicyOutput() AppSecRatePolicyOutput {
	return i.ToAppSecRatePolicyOutputWithContext(context.Background())
}

func (i *AppSecRatePolicy) ToAppSecRatePolicyOutputWithContext(ctx context.Context) AppSecRatePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyOutput)
}

func (i *AppSecRatePolicy) ToAppSecRatePolicyPtrOutput() AppSecRatePolicyPtrOutput {
	return i.ToAppSecRatePolicyPtrOutputWithContext(context.Background())
}

func (i *AppSecRatePolicy) ToAppSecRatePolicyPtrOutputWithContext(ctx context.Context) AppSecRatePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyPtrOutput)
}

type AppSecRatePolicyPtrInput interface {
	pulumi.Input

	ToAppSecRatePolicyPtrOutput() AppSecRatePolicyPtrOutput
	ToAppSecRatePolicyPtrOutputWithContext(ctx context.Context) AppSecRatePolicyPtrOutput
}

type appSecRatePolicyPtrType AppSecRatePolicyArgs

func (*appSecRatePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRatePolicy)(nil))
}

func (i *appSecRatePolicyPtrType) ToAppSecRatePolicyPtrOutput() AppSecRatePolicyPtrOutput {
	return i.ToAppSecRatePolicyPtrOutputWithContext(context.Background())
}

func (i *appSecRatePolicyPtrType) ToAppSecRatePolicyPtrOutputWithContext(ctx context.Context) AppSecRatePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecRatePolicyPtrOutput)
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
	return reflect.TypeOf((*AppSecRatePolicy)(nil))
}

func (o AppSecRatePolicyOutput) ToAppSecRatePolicyOutput() AppSecRatePolicyOutput {
	return o
}

func (o AppSecRatePolicyOutput) ToAppSecRatePolicyOutputWithContext(ctx context.Context) AppSecRatePolicyOutput {
	return o
}

func (o AppSecRatePolicyOutput) ToAppSecRatePolicyPtrOutput() AppSecRatePolicyPtrOutput {
	return o.ToAppSecRatePolicyPtrOutputWithContext(context.Background())
}

func (o AppSecRatePolicyOutput) ToAppSecRatePolicyPtrOutputWithContext(ctx context.Context) AppSecRatePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSecRatePolicy) *AppSecRatePolicy {
		return &v
	}).(AppSecRatePolicyPtrOutput)
}

type AppSecRatePolicyPtrOutput struct{ *pulumi.OutputState }

func (AppSecRatePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecRatePolicy)(nil))
}

func (o AppSecRatePolicyPtrOutput) ToAppSecRatePolicyPtrOutput() AppSecRatePolicyPtrOutput {
	return o
}

func (o AppSecRatePolicyPtrOutput) ToAppSecRatePolicyPtrOutputWithContext(ctx context.Context) AppSecRatePolicyPtrOutput {
	return o
}

func (o AppSecRatePolicyPtrOutput) Elem() AppSecRatePolicyOutput {
	return o.ApplyT(func(v *AppSecRatePolicy) AppSecRatePolicy {
		if v != nil {
			return *v
		}
		var ret AppSecRatePolicy
		return ret
	}).(AppSecRatePolicyOutput)
}

type AppSecRatePolicyArrayOutput struct{ *pulumi.OutputState }

func (AppSecRatePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecRatePolicy)(nil))
}

func (o AppSecRatePolicyArrayOutput) ToAppSecRatePolicyArrayOutput() AppSecRatePolicyArrayOutput {
	return o
}

func (o AppSecRatePolicyArrayOutput) ToAppSecRatePolicyArrayOutputWithContext(ctx context.Context) AppSecRatePolicyArrayOutput {
	return o
}

func (o AppSecRatePolicyArrayOutput) Index(i pulumi.IntInput) AppSecRatePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecRatePolicy {
		return vs[0].([]AppSecRatePolicy)[vs[1].(int)]
	}).(AppSecRatePolicyOutput)
}

type AppSecRatePolicyMapOutput struct{ *pulumi.OutputState }

func (AppSecRatePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecRatePolicy)(nil))
}

func (o AppSecRatePolicyMapOutput) ToAppSecRatePolicyMapOutput() AppSecRatePolicyMapOutput {
	return o
}

func (o AppSecRatePolicyMapOutput) ToAppSecRatePolicyMapOutputWithContext(ctx context.Context) AppSecRatePolicyMapOutput {
	return o
}

func (o AppSecRatePolicyMapOutput) MapIndex(k pulumi.StringInput) AppSecRatePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecRatePolicy {
		return vs[0].(map[string]AppSecRatePolicy)[vs[1].(string)]
	}).(AppSecRatePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRatePolicyInput)(nil)).Elem(), &AppSecRatePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRatePolicyPtrInput)(nil)).Elem(), &AppSecRatePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRatePolicyArrayInput)(nil)).Elem(), AppSecRatePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecRatePolicyMapInput)(nil)).Elem(), AppSecRatePolicyMap{})
	pulumi.RegisterOutputType(AppSecRatePolicyOutput{})
	pulumi.RegisterOutputType(AppSecRatePolicyPtrOutput{})
	pulumi.RegisterOutputType(AppSecRatePolicyArrayOutput{})
	pulumi.RegisterOutputType(AppSecRatePolicyMapOutput{})
}
