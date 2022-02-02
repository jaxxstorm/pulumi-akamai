// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration
//
// **Note**: This data source is deprecated and may be removed in a future release.
//
// Modifies the list of hostnames evaluated while a security configuration is in evaluation mode.
// During evaluation mode, hosts take no action of any kind when responding to traffic.
// Instead, these hosts simply maintain a record of the actions they *would* have taken if they had been responding to live traffic in your production network.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames/eval-hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevaluationhostnames)
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
// 		_, err = akamai.NewAppSecEvalHostnames(ctx, "evalHostnames", &akamai.AppSecEvalHostnamesArgs{
// 			ConfigId: pulumi.Int(configuration.ConfigId),
// 			Hostnames: pulumi.StringArray{
// 				pulumi.String("documentation.akamai.com"),
// 				pulumi.String("training.akamai.com"),
// 				pulumi.String("videos.akamai.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecEvalHostnames struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
	Hostnames pulumi.StringArrayOutput `pulumi:"hostnames"`
}

// NewAppSecEvalHostnames registers a new resource with the given unique name, arguments, and options.
func NewAppSecEvalHostnames(ctx *pulumi.Context,
	name string, args *AppSecEvalHostnamesArgs, opts ...pulumi.ResourceOption) (*AppSecEvalHostnames, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.Hostnames == nil {
		return nil, errors.New("invalid value for required argument 'Hostnames'")
	}
	var resource AppSecEvalHostnames
	err := ctx.RegisterResource("akamai:index/appSecEvalHostnames:AppSecEvalHostnames", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecEvalHostnames gets an existing AppSecEvalHostnames resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecEvalHostnames(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecEvalHostnamesState, opts ...pulumi.ResourceOption) (*AppSecEvalHostnames, error) {
	var resource AppSecEvalHostnames
	err := ctx.ReadResource("akamai:index/appSecEvalHostnames:AppSecEvalHostnames", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecEvalHostnames resources.
type appSecEvalHostnamesState struct {
	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId *int `pulumi:"configId"`
	// . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
	Hostnames []string `pulumi:"hostnames"`
}

type AppSecEvalHostnamesState struct {
	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId pulumi.IntPtrInput
	// . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
	Hostnames pulumi.StringArrayInput
}

func (AppSecEvalHostnamesState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecEvalHostnamesState)(nil)).Elem()
}

type appSecEvalHostnamesArgs struct {
	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId int `pulumi:"configId"`
	// . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
	Hostnames []string `pulumi:"hostnames"`
}

// The set of arguments for constructing a AppSecEvalHostnames resource.
type AppSecEvalHostnamesArgs struct {
	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId pulumi.IntInput
	// . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
	Hostnames pulumi.StringArrayInput
}

func (AppSecEvalHostnamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecEvalHostnamesArgs)(nil)).Elem()
}

type AppSecEvalHostnamesInput interface {
	pulumi.Input

	ToAppSecEvalHostnamesOutput() AppSecEvalHostnamesOutput
	ToAppSecEvalHostnamesOutputWithContext(ctx context.Context) AppSecEvalHostnamesOutput
}

func (*AppSecEvalHostnames) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecEvalHostnames)(nil)).Elem()
}

func (i *AppSecEvalHostnames) ToAppSecEvalHostnamesOutput() AppSecEvalHostnamesOutput {
	return i.ToAppSecEvalHostnamesOutputWithContext(context.Background())
}

func (i *AppSecEvalHostnames) ToAppSecEvalHostnamesOutputWithContext(ctx context.Context) AppSecEvalHostnamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalHostnamesOutput)
}

// AppSecEvalHostnamesArrayInput is an input type that accepts AppSecEvalHostnamesArray and AppSecEvalHostnamesArrayOutput values.
// You can construct a concrete instance of `AppSecEvalHostnamesArrayInput` via:
//
//          AppSecEvalHostnamesArray{ AppSecEvalHostnamesArgs{...} }
type AppSecEvalHostnamesArrayInput interface {
	pulumi.Input

	ToAppSecEvalHostnamesArrayOutput() AppSecEvalHostnamesArrayOutput
	ToAppSecEvalHostnamesArrayOutputWithContext(context.Context) AppSecEvalHostnamesArrayOutput
}

type AppSecEvalHostnamesArray []AppSecEvalHostnamesInput

func (AppSecEvalHostnamesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecEvalHostnames)(nil)).Elem()
}

func (i AppSecEvalHostnamesArray) ToAppSecEvalHostnamesArrayOutput() AppSecEvalHostnamesArrayOutput {
	return i.ToAppSecEvalHostnamesArrayOutputWithContext(context.Background())
}

func (i AppSecEvalHostnamesArray) ToAppSecEvalHostnamesArrayOutputWithContext(ctx context.Context) AppSecEvalHostnamesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalHostnamesArrayOutput)
}

// AppSecEvalHostnamesMapInput is an input type that accepts AppSecEvalHostnamesMap and AppSecEvalHostnamesMapOutput values.
// You can construct a concrete instance of `AppSecEvalHostnamesMapInput` via:
//
//          AppSecEvalHostnamesMap{ "key": AppSecEvalHostnamesArgs{...} }
type AppSecEvalHostnamesMapInput interface {
	pulumi.Input

	ToAppSecEvalHostnamesMapOutput() AppSecEvalHostnamesMapOutput
	ToAppSecEvalHostnamesMapOutputWithContext(context.Context) AppSecEvalHostnamesMapOutput
}

type AppSecEvalHostnamesMap map[string]AppSecEvalHostnamesInput

func (AppSecEvalHostnamesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecEvalHostnames)(nil)).Elem()
}

func (i AppSecEvalHostnamesMap) ToAppSecEvalHostnamesMapOutput() AppSecEvalHostnamesMapOutput {
	return i.ToAppSecEvalHostnamesMapOutputWithContext(context.Background())
}

func (i AppSecEvalHostnamesMap) ToAppSecEvalHostnamesMapOutputWithContext(ctx context.Context) AppSecEvalHostnamesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalHostnamesMapOutput)
}

type AppSecEvalHostnamesOutput struct{ *pulumi.OutputState }

func (AppSecEvalHostnamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecEvalHostnames)(nil)).Elem()
}

func (o AppSecEvalHostnamesOutput) ToAppSecEvalHostnamesOutput() AppSecEvalHostnamesOutput {
	return o
}

func (o AppSecEvalHostnamesOutput) ToAppSecEvalHostnamesOutputWithContext(ctx context.Context) AppSecEvalHostnamesOutput {
	return o
}

type AppSecEvalHostnamesArrayOutput struct{ *pulumi.OutputState }

func (AppSecEvalHostnamesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecEvalHostnames)(nil)).Elem()
}

func (o AppSecEvalHostnamesArrayOutput) ToAppSecEvalHostnamesArrayOutput() AppSecEvalHostnamesArrayOutput {
	return o
}

func (o AppSecEvalHostnamesArrayOutput) ToAppSecEvalHostnamesArrayOutputWithContext(ctx context.Context) AppSecEvalHostnamesArrayOutput {
	return o
}

func (o AppSecEvalHostnamesArrayOutput) Index(i pulumi.IntInput) AppSecEvalHostnamesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecEvalHostnames {
		return vs[0].([]*AppSecEvalHostnames)[vs[1].(int)]
	}).(AppSecEvalHostnamesOutput)
}

type AppSecEvalHostnamesMapOutput struct{ *pulumi.OutputState }

func (AppSecEvalHostnamesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecEvalHostnames)(nil)).Elem()
}

func (o AppSecEvalHostnamesMapOutput) ToAppSecEvalHostnamesMapOutput() AppSecEvalHostnamesMapOutput {
	return o
}

func (o AppSecEvalHostnamesMapOutput) ToAppSecEvalHostnamesMapOutputWithContext(ctx context.Context) AppSecEvalHostnamesMapOutput {
	return o
}

func (o AppSecEvalHostnamesMapOutput) MapIndex(k pulumi.StringInput) AppSecEvalHostnamesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecEvalHostnames {
		return vs[0].(map[string]*AppSecEvalHostnames)[vs[1].(string)]
	}).(AppSecEvalHostnamesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalHostnamesInput)(nil)).Elem(), &AppSecEvalHostnames{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalHostnamesArrayInput)(nil)).Elem(), AppSecEvalHostnamesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalHostnamesMapInput)(nil)).Elem(), AppSecEvalHostnamesMap{})
	pulumi.RegisterOutputType(AppSecEvalHostnamesOutput{})
	pulumi.RegisterOutputType(AppSecEvalHostnamesArrayOutput{})
	pulumi.RegisterOutputType(AppSecEvalHostnamesMapOutput{})
}
