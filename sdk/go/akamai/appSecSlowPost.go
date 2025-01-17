// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy
//
// Modifies slow POST protection settings for a security configuration and security policy. Slow POST protections help defend a site against attacks that try to tie up the site by using extremely slow requests and responses.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/slow-post](https://techdocs.akamai.com/application-security/reference/put-policy-slow-post)
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
//			_, err = akamai.NewAppSecSlowPost(ctx, "slowPost", &akamai.AppSecSlowPostArgs{
//				ConfigId:                 pulumi.Int(configuration.ConfigId),
//				SecurityPolicyId:         pulumi.String("gms1_134637"),
//				SlowRateAction:           pulumi.String("alert"),
//				SlowRateThresholdRate:    pulumi.Int(10),
//				SlowRateThresholdPeriod:  pulumi.Int(30),
//				DurationThresholdTimeout: pulumi.Int(20),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AppSecSlowPost struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the slow POST settings being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
	DurationThresholdTimeout pulumi.IntPtrOutput `pulumi:"durationThresholdTimeout"`
	// . Unique identifier of the security policy associated with the slow POST settings being modified.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
	// . Action to be taken if slow POST protection is triggered. Allowed values are:
	// - **alert**. Record the event.
	// - **abort**. Block the request.
	SlowRateAction pulumi.StringOutput `pulumi:"slowRateAction"`
	// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
	SlowRateThresholdPeriod pulumi.IntPtrOutput `pulumi:"slowRateThresholdPeriod"`
	// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
	SlowRateThresholdRate pulumi.IntPtrOutput `pulumi:"slowRateThresholdRate"`
}

// NewAppSecSlowPost registers a new resource with the given unique name, arguments, and options.
func NewAppSecSlowPost(ctx *pulumi.Context,
	name string, args *AppSecSlowPostArgs, opts ...pulumi.ResourceOption) (*AppSecSlowPost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	if args.SlowRateAction == nil {
		return nil, errors.New("invalid value for required argument 'SlowRateAction'")
	}
	var resource AppSecSlowPost
	err := ctx.RegisterResource("akamai:index/appSecSlowPost:AppSecSlowPost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecSlowPost gets an existing AppSecSlowPost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecSlowPost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecSlowPostState, opts ...pulumi.ResourceOption) (*AppSecSlowPost, error) {
	var resource AppSecSlowPost
	err := ctx.ReadResource("akamai:index/appSecSlowPost:AppSecSlowPost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecSlowPost resources.
type appSecSlowPostState struct {
	// . Unique identifier of the security configuration associated with the slow POST settings being modified.
	ConfigId *int `pulumi:"configId"`
	// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
	DurationThresholdTimeout *int `pulumi:"durationThresholdTimeout"`
	// . Unique identifier of the security policy associated with the slow POST settings being modified.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
	// . Action to be taken if slow POST protection is triggered. Allowed values are:
	// - **alert**. Record the event.
	// - **abort**. Block the request.
	SlowRateAction *string `pulumi:"slowRateAction"`
	// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
	SlowRateThresholdPeriod *int `pulumi:"slowRateThresholdPeriod"`
	// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
	SlowRateThresholdRate *int `pulumi:"slowRateThresholdRate"`
}

type AppSecSlowPostState struct {
	// . Unique identifier of the security configuration associated with the slow POST settings being modified.
	ConfigId pulumi.IntPtrInput
	// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
	DurationThresholdTimeout pulumi.IntPtrInput
	// . Unique identifier of the security policy associated with the slow POST settings being modified.
	SecurityPolicyId pulumi.StringPtrInput
	// . Action to be taken if slow POST protection is triggered. Allowed values are:
	// - **alert**. Record the event.
	// - **abort**. Block the request.
	SlowRateAction pulumi.StringPtrInput
	// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
	SlowRateThresholdPeriod pulumi.IntPtrInput
	// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
	SlowRateThresholdRate pulumi.IntPtrInput
}

func (AppSecSlowPostState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSlowPostState)(nil)).Elem()
}

type appSecSlowPostArgs struct {
	// . Unique identifier of the security configuration associated with the slow POST settings being modified.
	ConfigId int `pulumi:"configId"`
	// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
	DurationThresholdTimeout *int `pulumi:"durationThresholdTimeout"`
	// . Unique identifier of the security policy associated with the slow POST settings being modified.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	// . Action to be taken if slow POST protection is triggered. Allowed values are:
	// - **alert**. Record the event.
	// - **abort**. Block the request.
	SlowRateAction string `pulumi:"slowRateAction"`
	// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
	SlowRateThresholdPeriod *int `pulumi:"slowRateThresholdPeriod"`
	// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
	SlowRateThresholdRate *int `pulumi:"slowRateThresholdRate"`
}

// The set of arguments for constructing a AppSecSlowPost resource.
type AppSecSlowPostArgs struct {
	// . Unique identifier of the security configuration associated with the slow POST settings being modified.
	ConfigId pulumi.IntInput
	// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
	DurationThresholdTimeout pulumi.IntPtrInput
	// . Unique identifier of the security policy associated with the slow POST settings being modified.
	SecurityPolicyId pulumi.StringInput
	// . Action to be taken if slow POST protection is triggered. Allowed values are:
	// - **alert**. Record the event.
	// - **abort**. Block the request.
	SlowRateAction pulumi.StringInput
	// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
	SlowRateThresholdPeriod pulumi.IntPtrInput
	// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
	SlowRateThresholdRate pulumi.IntPtrInput
}

func (AppSecSlowPostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSlowPostArgs)(nil)).Elem()
}

type AppSecSlowPostInput interface {
	pulumi.Input

	ToAppSecSlowPostOutput() AppSecSlowPostOutput
	ToAppSecSlowPostOutputWithContext(ctx context.Context) AppSecSlowPostOutput
}

func (*AppSecSlowPost) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSlowPost)(nil)).Elem()
}

func (i *AppSecSlowPost) ToAppSecSlowPostOutput() AppSecSlowPostOutput {
	return i.ToAppSecSlowPostOutputWithContext(context.Background())
}

func (i *AppSecSlowPost) ToAppSecSlowPostOutputWithContext(ctx context.Context) AppSecSlowPostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSlowPostOutput)
}

// AppSecSlowPostArrayInput is an input type that accepts AppSecSlowPostArray and AppSecSlowPostArrayOutput values.
// You can construct a concrete instance of `AppSecSlowPostArrayInput` via:
//
//	AppSecSlowPostArray{ AppSecSlowPostArgs{...} }
type AppSecSlowPostArrayInput interface {
	pulumi.Input

	ToAppSecSlowPostArrayOutput() AppSecSlowPostArrayOutput
	ToAppSecSlowPostArrayOutputWithContext(context.Context) AppSecSlowPostArrayOutput
}

type AppSecSlowPostArray []AppSecSlowPostInput

func (AppSecSlowPostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecSlowPost)(nil)).Elem()
}

func (i AppSecSlowPostArray) ToAppSecSlowPostArrayOutput() AppSecSlowPostArrayOutput {
	return i.ToAppSecSlowPostArrayOutputWithContext(context.Background())
}

func (i AppSecSlowPostArray) ToAppSecSlowPostArrayOutputWithContext(ctx context.Context) AppSecSlowPostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSlowPostArrayOutput)
}

// AppSecSlowPostMapInput is an input type that accepts AppSecSlowPostMap and AppSecSlowPostMapOutput values.
// You can construct a concrete instance of `AppSecSlowPostMapInput` via:
//
//	AppSecSlowPostMap{ "key": AppSecSlowPostArgs{...} }
type AppSecSlowPostMapInput interface {
	pulumi.Input

	ToAppSecSlowPostMapOutput() AppSecSlowPostMapOutput
	ToAppSecSlowPostMapOutputWithContext(context.Context) AppSecSlowPostMapOutput
}

type AppSecSlowPostMap map[string]AppSecSlowPostInput

func (AppSecSlowPostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecSlowPost)(nil)).Elem()
}

func (i AppSecSlowPostMap) ToAppSecSlowPostMapOutput() AppSecSlowPostMapOutput {
	return i.ToAppSecSlowPostMapOutputWithContext(context.Background())
}

func (i AppSecSlowPostMap) ToAppSecSlowPostMapOutputWithContext(ctx context.Context) AppSecSlowPostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSlowPostMapOutput)
}

type AppSecSlowPostOutput struct{ *pulumi.OutputState }

func (AppSecSlowPostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSlowPost)(nil)).Elem()
}

func (o AppSecSlowPostOutput) ToAppSecSlowPostOutput() AppSecSlowPostOutput {
	return o
}

func (o AppSecSlowPostOutput) ToAppSecSlowPostOutputWithContext(ctx context.Context) AppSecSlowPostOutput {
	return o
}

// . Unique identifier of the security configuration associated with the slow POST settings being modified.
func (o AppSecSlowPostOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecSlowPost) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// . Maximum amount of time (in seconds) that the first eight kilobytes of the POST body must be received in to avoid triggering the specified action.
func (o AppSecSlowPostOutput) DurationThresholdTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppSecSlowPost) pulumi.IntPtrOutput { return v.DurationThresholdTimeout }).(pulumi.IntPtrOutput)
}

// . Unique identifier of the security policy associated with the slow POST settings being modified.
func (o AppSecSlowPostOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecSlowPost) pulumi.StringOutput { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

// . Action to be taken if slow POST protection is triggered. Allowed values are:
// - **alert**. Record the event.
// - **abort**. Block the request.
func (o AppSecSlowPostOutput) SlowRateAction() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecSlowPost) pulumi.StringOutput { return v.SlowRateAction }).(pulumi.StringOutput)
}

// . Amount of time (in seconds) that the server should allow a request before marking the request as being too slow.
func (o AppSecSlowPostOutput) SlowRateThresholdPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppSecSlowPost) pulumi.IntPtrOutput { return v.SlowRateThresholdPeriod }).(pulumi.IntPtrOutput)
}

// . Average rate (in bytes per second over the specified time period) allowed before the specified action is triggered.
func (o AppSecSlowPostOutput) SlowRateThresholdRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppSecSlowPost) pulumi.IntPtrOutput { return v.SlowRateThresholdRate }).(pulumi.IntPtrOutput)
}

type AppSecSlowPostArrayOutput struct{ *pulumi.OutputState }

func (AppSecSlowPostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecSlowPost)(nil)).Elem()
}

func (o AppSecSlowPostArrayOutput) ToAppSecSlowPostArrayOutput() AppSecSlowPostArrayOutput {
	return o
}

func (o AppSecSlowPostArrayOutput) ToAppSecSlowPostArrayOutputWithContext(ctx context.Context) AppSecSlowPostArrayOutput {
	return o
}

func (o AppSecSlowPostArrayOutput) Index(i pulumi.IntInput) AppSecSlowPostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecSlowPost {
		return vs[0].([]*AppSecSlowPost)[vs[1].(int)]
	}).(AppSecSlowPostOutput)
}

type AppSecSlowPostMapOutput struct{ *pulumi.OutputState }

func (AppSecSlowPostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecSlowPost)(nil)).Elem()
}

func (o AppSecSlowPostMapOutput) ToAppSecSlowPostMapOutput() AppSecSlowPostMapOutput {
	return o
}

func (o AppSecSlowPostMapOutput) ToAppSecSlowPostMapOutputWithContext(ctx context.Context) AppSecSlowPostMapOutput {
	return o
}

func (o AppSecSlowPostMapOutput) MapIndex(k pulumi.StringInput) AppSecSlowPostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecSlowPost {
		return vs[0].(map[string]*AppSecSlowPost)[vs[1].(string)]
	}).(AppSecSlowPostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSlowPostInput)(nil)).Elem(), &AppSecSlowPost{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSlowPostArrayInput)(nil)).Elem(), AppSecSlowPostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSlowPostMapInput)(nil)).Elem(), AppSecSlowPostMap{})
	pulumi.RegisterOutputType(AppSecSlowPostOutput{})
	pulumi.RegisterOutputType(AppSecSlowPostArrayOutput{})
	pulumi.RegisterOutputType(AppSecSlowPostMapOutput{})
}
