// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AppSecSelectedHostnames struct {
	pulumi.CustomResourceState

	ConfigId  pulumi.IntOutput         `pulumi:"configId"`
	Hostnames pulumi.StringArrayOutput `pulumi:"hostnames"`
	Mode      pulumi.StringOutput      `pulumi:"mode"`
	Version   pulumi.IntOutput         `pulumi:"version"`
}

// NewAppSecSelectedHostnames registers a new resource with the given unique name, arguments, and options.
func NewAppSecSelectedHostnames(ctx *pulumi.Context,
	name string, args *AppSecSelectedHostnamesArgs, opts ...pulumi.ResourceOption) (*AppSecSelectedHostnames, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.Hostnames == nil {
		return nil, errors.New("invalid value for required argument 'Hostnames'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource AppSecSelectedHostnames
	err := ctx.RegisterResource("akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecSelectedHostnames gets an existing AppSecSelectedHostnames resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecSelectedHostnames(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecSelectedHostnamesState, opts ...pulumi.ResourceOption) (*AppSecSelectedHostnames, error) {
	var resource AppSecSelectedHostnames
	err := ctx.ReadResource("akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecSelectedHostnames resources.
type appSecSelectedHostnamesState struct {
	ConfigId  *int     `pulumi:"configId"`
	Hostnames []string `pulumi:"hostnames"`
	Mode      *string  `pulumi:"mode"`
	Version   *int     `pulumi:"version"`
}

type AppSecSelectedHostnamesState struct {
	ConfigId  pulumi.IntPtrInput
	Hostnames pulumi.StringArrayInput
	Mode      pulumi.StringPtrInput
	Version   pulumi.IntPtrInput
}

func (AppSecSelectedHostnamesState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSelectedHostnamesState)(nil)).Elem()
}

type appSecSelectedHostnamesArgs struct {
	ConfigId  int      `pulumi:"configId"`
	Hostnames []string `pulumi:"hostnames"`
	Mode      string   `pulumi:"mode"`
	Version   int      `pulumi:"version"`
}

// The set of arguments for constructing a AppSecSelectedHostnames resource.
type AppSecSelectedHostnamesArgs struct {
	ConfigId  pulumi.IntInput
	Hostnames pulumi.StringArrayInput
	Mode      pulumi.StringInput
	Version   pulumi.IntInput
}

func (AppSecSelectedHostnamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSelectedHostnamesArgs)(nil)).Elem()
}

type AppSecSelectedHostnamesInput interface {
	pulumi.Input

	ToAppSecSelectedHostnamesOutput() AppSecSelectedHostnamesOutput
	ToAppSecSelectedHostnamesOutputWithContext(ctx context.Context) AppSecSelectedHostnamesOutput
}

func (*AppSecSelectedHostnames) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecSelectedHostnames)(nil))
}

func (i *AppSecSelectedHostnames) ToAppSecSelectedHostnamesOutput() AppSecSelectedHostnamesOutput {
	return i.ToAppSecSelectedHostnamesOutputWithContext(context.Background())
}

func (i *AppSecSelectedHostnames) ToAppSecSelectedHostnamesOutputWithContext(ctx context.Context) AppSecSelectedHostnamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSelectedHostnamesOutput)
}

func (i *AppSecSelectedHostnames) ToAppSecSelectedHostnamesPtrOutput() AppSecSelectedHostnamesPtrOutput {
	return i.ToAppSecSelectedHostnamesPtrOutputWithContext(context.Background())
}

func (i *AppSecSelectedHostnames) ToAppSecSelectedHostnamesPtrOutputWithContext(ctx context.Context) AppSecSelectedHostnamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSelectedHostnamesPtrOutput)
}

type AppSecSelectedHostnamesPtrInput interface {
	pulumi.Input

	ToAppSecSelectedHostnamesPtrOutput() AppSecSelectedHostnamesPtrOutput
	ToAppSecSelectedHostnamesPtrOutputWithContext(ctx context.Context) AppSecSelectedHostnamesPtrOutput
}

type appSecSelectedHostnamesPtrType AppSecSelectedHostnamesArgs

func (*appSecSelectedHostnamesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSelectedHostnames)(nil))
}

func (i *appSecSelectedHostnamesPtrType) ToAppSecSelectedHostnamesPtrOutput() AppSecSelectedHostnamesPtrOutput {
	return i.ToAppSecSelectedHostnamesPtrOutputWithContext(context.Background())
}

func (i *appSecSelectedHostnamesPtrType) ToAppSecSelectedHostnamesPtrOutputWithContext(ctx context.Context) AppSecSelectedHostnamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSelectedHostnamesPtrOutput)
}

// AppSecSelectedHostnamesArrayInput is an input type that accepts AppSecSelectedHostnamesArray and AppSecSelectedHostnamesArrayOutput values.
// You can construct a concrete instance of `AppSecSelectedHostnamesArrayInput` via:
//
//          AppSecSelectedHostnamesArray{ AppSecSelectedHostnamesArgs{...} }
type AppSecSelectedHostnamesArrayInput interface {
	pulumi.Input

	ToAppSecSelectedHostnamesArrayOutput() AppSecSelectedHostnamesArrayOutput
	ToAppSecSelectedHostnamesArrayOutputWithContext(context.Context) AppSecSelectedHostnamesArrayOutput
}

type AppSecSelectedHostnamesArray []AppSecSelectedHostnamesInput

func (AppSecSelectedHostnamesArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecSelectedHostnames)(nil))
}

func (i AppSecSelectedHostnamesArray) ToAppSecSelectedHostnamesArrayOutput() AppSecSelectedHostnamesArrayOutput {
	return i.ToAppSecSelectedHostnamesArrayOutputWithContext(context.Background())
}

func (i AppSecSelectedHostnamesArray) ToAppSecSelectedHostnamesArrayOutputWithContext(ctx context.Context) AppSecSelectedHostnamesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSelectedHostnamesArrayOutput)
}

// AppSecSelectedHostnamesMapInput is an input type that accepts AppSecSelectedHostnamesMap and AppSecSelectedHostnamesMapOutput values.
// You can construct a concrete instance of `AppSecSelectedHostnamesMapInput` via:
//
//          AppSecSelectedHostnamesMap{ "key": AppSecSelectedHostnamesArgs{...} }
type AppSecSelectedHostnamesMapInput interface {
	pulumi.Input

	ToAppSecSelectedHostnamesMapOutput() AppSecSelectedHostnamesMapOutput
	ToAppSecSelectedHostnamesMapOutputWithContext(context.Context) AppSecSelectedHostnamesMapOutput
}

type AppSecSelectedHostnamesMap map[string]AppSecSelectedHostnamesInput

func (AppSecSelectedHostnamesMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecSelectedHostnames)(nil))
}

func (i AppSecSelectedHostnamesMap) ToAppSecSelectedHostnamesMapOutput() AppSecSelectedHostnamesMapOutput {
	return i.ToAppSecSelectedHostnamesMapOutputWithContext(context.Background())
}

func (i AppSecSelectedHostnamesMap) ToAppSecSelectedHostnamesMapOutputWithContext(ctx context.Context) AppSecSelectedHostnamesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSelectedHostnamesMapOutput)
}

type AppSecSelectedHostnamesOutput struct {
	*pulumi.OutputState
}

func (AppSecSelectedHostnamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecSelectedHostnames)(nil))
}

func (o AppSecSelectedHostnamesOutput) ToAppSecSelectedHostnamesOutput() AppSecSelectedHostnamesOutput {
	return o
}

func (o AppSecSelectedHostnamesOutput) ToAppSecSelectedHostnamesOutputWithContext(ctx context.Context) AppSecSelectedHostnamesOutput {
	return o
}

func (o AppSecSelectedHostnamesOutput) ToAppSecSelectedHostnamesPtrOutput() AppSecSelectedHostnamesPtrOutput {
	return o.ToAppSecSelectedHostnamesPtrOutputWithContext(context.Background())
}

func (o AppSecSelectedHostnamesOutput) ToAppSecSelectedHostnamesPtrOutputWithContext(ctx context.Context) AppSecSelectedHostnamesPtrOutput {
	return o.ApplyT(func(v AppSecSelectedHostnames) *AppSecSelectedHostnames {
		return &v
	}).(AppSecSelectedHostnamesPtrOutput)
}

type AppSecSelectedHostnamesPtrOutput struct {
	*pulumi.OutputState
}

func (AppSecSelectedHostnamesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSelectedHostnames)(nil))
}

func (o AppSecSelectedHostnamesPtrOutput) ToAppSecSelectedHostnamesPtrOutput() AppSecSelectedHostnamesPtrOutput {
	return o
}

func (o AppSecSelectedHostnamesPtrOutput) ToAppSecSelectedHostnamesPtrOutputWithContext(ctx context.Context) AppSecSelectedHostnamesPtrOutput {
	return o
}

type AppSecSelectedHostnamesArrayOutput struct{ *pulumi.OutputState }

func (AppSecSelectedHostnamesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecSelectedHostnames)(nil))
}

func (o AppSecSelectedHostnamesArrayOutput) ToAppSecSelectedHostnamesArrayOutput() AppSecSelectedHostnamesArrayOutput {
	return o
}

func (o AppSecSelectedHostnamesArrayOutput) ToAppSecSelectedHostnamesArrayOutputWithContext(ctx context.Context) AppSecSelectedHostnamesArrayOutput {
	return o
}

func (o AppSecSelectedHostnamesArrayOutput) Index(i pulumi.IntInput) AppSecSelectedHostnamesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecSelectedHostnames {
		return vs[0].([]AppSecSelectedHostnames)[vs[1].(int)]
	}).(AppSecSelectedHostnamesOutput)
}

type AppSecSelectedHostnamesMapOutput struct{ *pulumi.OutputState }

func (AppSecSelectedHostnamesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecSelectedHostnames)(nil))
}

func (o AppSecSelectedHostnamesMapOutput) ToAppSecSelectedHostnamesMapOutput() AppSecSelectedHostnamesMapOutput {
	return o
}

func (o AppSecSelectedHostnamesMapOutput) ToAppSecSelectedHostnamesMapOutputWithContext(ctx context.Context) AppSecSelectedHostnamesMapOutput {
	return o
}

func (o AppSecSelectedHostnamesMapOutput) MapIndex(k pulumi.StringInput) AppSecSelectedHostnamesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecSelectedHostnames {
		return vs[0].(map[string]AppSecSelectedHostnames)[vs[1].(string)]
	}).(AppSecSelectedHostnamesOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecSelectedHostnamesOutput{})
	pulumi.RegisterOutputType(AppSecSelectedHostnamesPtrOutput{})
	pulumi.RegisterOutputType(AppSecSelectedHostnamesArrayOutput{})
	pulumi.RegisterOutputType(AppSecSelectedHostnamesMapOutput{})
}
