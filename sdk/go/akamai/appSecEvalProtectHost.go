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
// **Important**: This data source is deprecated and may be removed in a future release. You may use the `AppSecWapSelectedHostnames` resource instead.
//
// Moves hostnames being evaluated to active protection. When you move a hostname from the evaluation hostnames list that host is added to your security policy as a protected hostname and is removed from the collection of hosts being evaluated.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/protect-eval-hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putmoveevaluationhostnamestoprotection)
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
// 		evalHostnames, err := akamai.LookupAppSecEvalHostnames(ctx, &GetAppSecEvalHostnamesArgs{
// 			ConfigId: configuration.ConfigId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecEvalProtectHost(ctx, "protectHost", &akamai.AppSecEvalProtectHostArgs{
// 			ConfigId:  pulumi.Int(configuration.ConfigId),
// 			Hostnames: interface{}(evalHostnames.Hostnames),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecEvalProtectHost struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . JSON array of the hostnames to be moved from the evaluation hostname list to the protected hostname list.
	Hostnames pulumi.StringArrayOutput `pulumi:"hostnames"`
}

// NewAppSecEvalProtectHost registers a new resource with the given unique name, arguments, and options.
func NewAppSecEvalProtectHost(ctx *pulumi.Context,
	name string, args *AppSecEvalProtectHostArgs, opts ...pulumi.ResourceOption) (*AppSecEvalProtectHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.Hostnames == nil {
		return nil, errors.New("invalid value for required argument 'Hostnames'")
	}
	var resource AppSecEvalProtectHost
	err := ctx.RegisterResource("akamai:index/appSecEvalProtectHost:AppSecEvalProtectHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecEvalProtectHost gets an existing AppSecEvalProtectHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecEvalProtectHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecEvalProtectHostState, opts ...pulumi.ResourceOption) (*AppSecEvalProtectHost, error) {
	var resource AppSecEvalProtectHost
	err := ctx.ReadResource("akamai:index/appSecEvalProtectHost:AppSecEvalProtectHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecEvalProtectHost resources.
type appSecEvalProtectHostState struct {
	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId *int `pulumi:"configId"`
	// . JSON array of the hostnames to be moved from the evaluation hostname list to the protected hostname list.
	Hostnames []string `pulumi:"hostnames"`
}

type AppSecEvalProtectHostState struct {
	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId pulumi.IntPtrInput
	// . JSON array of the hostnames to be moved from the evaluation hostname list to the protected hostname list.
	Hostnames pulumi.StringArrayInput
}

func (AppSecEvalProtectHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecEvalProtectHostState)(nil)).Elem()
}

type appSecEvalProtectHostArgs struct {
	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId int `pulumi:"configId"`
	// . JSON array of the hostnames to be moved from the evaluation hostname list to the protected hostname list.
	Hostnames []string `pulumi:"hostnames"`
}

// The set of arguments for constructing a AppSecEvalProtectHost resource.
type AppSecEvalProtectHostArgs struct {
	// . Unique identifier of the security configuration in evaluation mode.
	ConfigId pulumi.IntInput
	// . JSON array of the hostnames to be moved from the evaluation hostname list to the protected hostname list.
	Hostnames pulumi.StringArrayInput
}

func (AppSecEvalProtectHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecEvalProtectHostArgs)(nil)).Elem()
}

type AppSecEvalProtectHostInput interface {
	pulumi.Input

	ToAppSecEvalProtectHostOutput() AppSecEvalProtectHostOutput
	ToAppSecEvalProtectHostOutputWithContext(ctx context.Context) AppSecEvalProtectHostOutput
}

func (*AppSecEvalProtectHost) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecEvalProtectHost)(nil))
}

func (i *AppSecEvalProtectHost) ToAppSecEvalProtectHostOutput() AppSecEvalProtectHostOutput {
	return i.ToAppSecEvalProtectHostOutputWithContext(context.Background())
}

func (i *AppSecEvalProtectHost) ToAppSecEvalProtectHostOutputWithContext(ctx context.Context) AppSecEvalProtectHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalProtectHostOutput)
}

func (i *AppSecEvalProtectHost) ToAppSecEvalProtectHostPtrOutput() AppSecEvalProtectHostPtrOutput {
	return i.ToAppSecEvalProtectHostPtrOutputWithContext(context.Background())
}

func (i *AppSecEvalProtectHost) ToAppSecEvalProtectHostPtrOutputWithContext(ctx context.Context) AppSecEvalProtectHostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalProtectHostPtrOutput)
}

type AppSecEvalProtectHostPtrInput interface {
	pulumi.Input

	ToAppSecEvalProtectHostPtrOutput() AppSecEvalProtectHostPtrOutput
	ToAppSecEvalProtectHostPtrOutputWithContext(ctx context.Context) AppSecEvalProtectHostPtrOutput
}

type appSecEvalProtectHostPtrType AppSecEvalProtectHostArgs

func (*appSecEvalProtectHostPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecEvalProtectHost)(nil))
}

func (i *appSecEvalProtectHostPtrType) ToAppSecEvalProtectHostPtrOutput() AppSecEvalProtectHostPtrOutput {
	return i.ToAppSecEvalProtectHostPtrOutputWithContext(context.Background())
}

func (i *appSecEvalProtectHostPtrType) ToAppSecEvalProtectHostPtrOutputWithContext(ctx context.Context) AppSecEvalProtectHostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalProtectHostPtrOutput)
}

// AppSecEvalProtectHostArrayInput is an input type that accepts AppSecEvalProtectHostArray and AppSecEvalProtectHostArrayOutput values.
// You can construct a concrete instance of `AppSecEvalProtectHostArrayInput` via:
//
//          AppSecEvalProtectHostArray{ AppSecEvalProtectHostArgs{...} }
type AppSecEvalProtectHostArrayInput interface {
	pulumi.Input

	ToAppSecEvalProtectHostArrayOutput() AppSecEvalProtectHostArrayOutput
	ToAppSecEvalProtectHostArrayOutputWithContext(context.Context) AppSecEvalProtectHostArrayOutput
}

type AppSecEvalProtectHostArray []AppSecEvalProtectHostInput

func (AppSecEvalProtectHostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecEvalProtectHost)(nil)).Elem()
}

func (i AppSecEvalProtectHostArray) ToAppSecEvalProtectHostArrayOutput() AppSecEvalProtectHostArrayOutput {
	return i.ToAppSecEvalProtectHostArrayOutputWithContext(context.Background())
}

func (i AppSecEvalProtectHostArray) ToAppSecEvalProtectHostArrayOutputWithContext(ctx context.Context) AppSecEvalProtectHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalProtectHostArrayOutput)
}

// AppSecEvalProtectHostMapInput is an input type that accepts AppSecEvalProtectHostMap and AppSecEvalProtectHostMapOutput values.
// You can construct a concrete instance of `AppSecEvalProtectHostMapInput` via:
//
//          AppSecEvalProtectHostMap{ "key": AppSecEvalProtectHostArgs{...} }
type AppSecEvalProtectHostMapInput interface {
	pulumi.Input

	ToAppSecEvalProtectHostMapOutput() AppSecEvalProtectHostMapOutput
	ToAppSecEvalProtectHostMapOutputWithContext(context.Context) AppSecEvalProtectHostMapOutput
}

type AppSecEvalProtectHostMap map[string]AppSecEvalProtectHostInput

func (AppSecEvalProtectHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecEvalProtectHost)(nil)).Elem()
}

func (i AppSecEvalProtectHostMap) ToAppSecEvalProtectHostMapOutput() AppSecEvalProtectHostMapOutput {
	return i.ToAppSecEvalProtectHostMapOutputWithContext(context.Background())
}

func (i AppSecEvalProtectHostMap) ToAppSecEvalProtectHostMapOutputWithContext(ctx context.Context) AppSecEvalProtectHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalProtectHostMapOutput)
}

type AppSecEvalProtectHostOutput struct{ *pulumi.OutputState }

func (AppSecEvalProtectHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecEvalProtectHost)(nil))
}

func (o AppSecEvalProtectHostOutput) ToAppSecEvalProtectHostOutput() AppSecEvalProtectHostOutput {
	return o
}

func (o AppSecEvalProtectHostOutput) ToAppSecEvalProtectHostOutputWithContext(ctx context.Context) AppSecEvalProtectHostOutput {
	return o
}

func (o AppSecEvalProtectHostOutput) ToAppSecEvalProtectHostPtrOutput() AppSecEvalProtectHostPtrOutput {
	return o.ToAppSecEvalProtectHostPtrOutputWithContext(context.Background())
}

func (o AppSecEvalProtectHostOutput) ToAppSecEvalProtectHostPtrOutputWithContext(ctx context.Context) AppSecEvalProtectHostPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSecEvalProtectHost) *AppSecEvalProtectHost {
		return &v
	}).(AppSecEvalProtectHostPtrOutput)
}

type AppSecEvalProtectHostPtrOutput struct{ *pulumi.OutputState }

func (AppSecEvalProtectHostPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecEvalProtectHost)(nil))
}

func (o AppSecEvalProtectHostPtrOutput) ToAppSecEvalProtectHostPtrOutput() AppSecEvalProtectHostPtrOutput {
	return o
}

func (o AppSecEvalProtectHostPtrOutput) ToAppSecEvalProtectHostPtrOutputWithContext(ctx context.Context) AppSecEvalProtectHostPtrOutput {
	return o
}

func (o AppSecEvalProtectHostPtrOutput) Elem() AppSecEvalProtectHostOutput {
	return o.ApplyT(func(v *AppSecEvalProtectHost) AppSecEvalProtectHost {
		if v != nil {
			return *v
		}
		var ret AppSecEvalProtectHost
		return ret
	}).(AppSecEvalProtectHostOutput)
}

type AppSecEvalProtectHostArrayOutput struct{ *pulumi.OutputState }

func (AppSecEvalProtectHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecEvalProtectHost)(nil))
}

func (o AppSecEvalProtectHostArrayOutput) ToAppSecEvalProtectHostArrayOutput() AppSecEvalProtectHostArrayOutput {
	return o
}

func (o AppSecEvalProtectHostArrayOutput) ToAppSecEvalProtectHostArrayOutputWithContext(ctx context.Context) AppSecEvalProtectHostArrayOutput {
	return o
}

func (o AppSecEvalProtectHostArrayOutput) Index(i pulumi.IntInput) AppSecEvalProtectHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecEvalProtectHost {
		return vs[0].([]AppSecEvalProtectHost)[vs[1].(int)]
	}).(AppSecEvalProtectHostOutput)
}

type AppSecEvalProtectHostMapOutput struct{ *pulumi.OutputState }

func (AppSecEvalProtectHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecEvalProtectHost)(nil))
}

func (o AppSecEvalProtectHostMapOutput) ToAppSecEvalProtectHostMapOutput() AppSecEvalProtectHostMapOutput {
	return o
}

func (o AppSecEvalProtectHostMapOutput) ToAppSecEvalProtectHostMapOutputWithContext(ctx context.Context) AppSecEvalProtectHostMapOutput {
	return o
}

func (o AppSecEvalProtectHostMapOutput) MapIndex(k pulumi.StringInput) AppSecEvalProtectHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecEvalProtectHost {
		return vs[0].(map[string]AppSecEvalProtectHost)[vs[1].(string)]
	}).(AppSecEvalProtectHostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalProtectHostInput)(nil)).Elem(), &AppSecEvalProtectHost{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalProtectHostPtrInput)(nil)).Elem(), &AppSecEvalProtectHost{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalProtectHostArrayInput)(nil)).Elem(), AppSecEvalProtectHostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalProtectHostMapInput)(nil)).Elem(), AppSecEvalProtectHostMap{})
	pulumi.RegisterOutputType(AppSecEvalProtectHostOutput{})
	pulumi.RegisterOutputType(AppSecEvalProtectHostPtrOutput{})
	pulumi.RegisterOutputType(AppSecEvalProtectHostArrayOutput{})
	pulumi.RegisterOutputType(AppSecEvalProtectHostMapOutput{})
}
