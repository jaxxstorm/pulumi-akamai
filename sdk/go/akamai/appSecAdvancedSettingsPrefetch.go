// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security configuration
//
// Enables inspection of internal requests (that is, requests between your origin servers and Akamai's edge servers). You can also use this resource to apply rate controls to prefetch requests.
//
// When prefetch is enabled, internal requests are inspected by your firewall the same way that external requests (requests that originate outside the firewall and outside Akamai's edge servers) are inspected.
//
// This operation applies at the security configuration level, meaning that the settings affect all the security policies in that configuration.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/prefetch](https://techdocs.akamai.com/application-security/reference/put-advanced-settings-prefetch)
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
//			_, err = akamai.NewAppSecAdvancedSettingsPrefetch(ctx, "prefetch", &akamai.AppSecAdvancedSettingsPrefetchArgs{
//				ConfigId:           pulumi.Int(configuration.ConfigId),
//				EnableAppLayer:     pulumi.Bool(false),
//				AllExtensions:      pulumi.Bool(true),
//				EnableRateControls: pulumi.Bool(false),
//				Extensions: pulumi.StringArray{
//					pulumi.String(".tiff"),
//					pulumi.String(".bmp"),
//					pulumi.String(".jpg"),
//					pulumi.String(".gif"),
//					pulumi.String(".png"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AppSecAdvancedSettingsPrefetch struct {
	pulumi.CustomResourceState

	// . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
	AllExtensions pulumi.BoolOutput `pulumi:"allExtensions"`
	// . Unique identifier of the security configuration associated with the prefetch settings being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
	EnableAppLayer pulumi.BoolOutput `pulumi:"enableAppLayer"`
	// . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
	EnableRateControls pulumi.BoolOutput `pulumi:"enableRateControls"`
	// . If `allExtensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `allExtensions` is **true**, then this argument must be set to an empty array: **[]**.
	Extensions pulumi.StringArrayOutput `pulumi:"extensions"`
}

// NewAppSecAdvancedSettingsPrefetch registers a new resource with the given unique name, arguments, and options.
func NewAppSecAdvancedSettingsPrefetch(ctx *pulumi.Context,
	name string, args *AppSecAdvancedSettingsPrefetchArgs, opts ...pulumi.ResourceOption) (*AppSecAdvancedSettingsPrefetch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllExtensions == nil {
		return nil, errors.New("invalid value for required argument 'AllExtensions'")
	}
	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.EnableAppLayer == nil {
		return nil, errors.New("invalid value for required argument 'EnableAppLayer'")
	}
	if args.EnableRateControls == nil {
		return nil, errors.New("invalid value for required argument 'EnableRateControls'")
	}
	if args.Extensions == nil {
		return nil, errors.New("invalid value for required argument 'Extensions'")
	}
	var resource AppSecAdvancedSettingsPrefetch
	err := ctx.RegisterResource("akamai:index/appSecAdvancedSettingsPrefetch:AppSecAdvancedSettingsPrefetch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecAdvancedSettingsPrefetch gets an existing AppSecAdvancedSettingsPrefetch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecAdvancedSettingsPrefetch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecAdvancedSettingsPrefetchState, opts ...pulumi.ResourceOption) (*AppSecAdvancedSettingsPrefetch, error) {
	var resource AppSecAdvancedSettingsPrefetch
	err := ctx.ReadResource("akamai:index/appSecAdvancedSettingsPrefetch:AppSecAdvancedSettingsPrefetch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecAdvancedSettingsPrefetch resources.
type appSecAdvancedSettingsPrefetchState struct {
	// . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
	AllExtensions *bool `pulumi:"allExtensions"`
	// . Unique identifier of the security configuration associated with the prefetch settings being modified.
	ConfigId *int `pulumi:"configId"`
	// . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
	EnableAppLayer *bool `pulumi:"enableAppLayer"`
	// . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
	EnableRateControls *bool `pulumi:"enableRateControls"`
	// . If `allExtensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `allExtensions` is **true**, then this argument must be set to an empty array: **[]**.
	Extensions []string `pulumi:"extensions"`
}

type AppSecAdvancedSettingsPrefetchState struct {
	// . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
	AllExtensions pulumi.BoolPtrInput
	// . Unique identifier of the security configuration associated with the prefetch settings being modified.
	ConfigId pulumi.IntPtrInput
	// . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
	EnableAppLayer pulumi.BoolPtrInput
	// . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
	EnableRateControls pulumi.BoolPtrInput
	// . If `allExtensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `allExtensions` is **true**, then this argument must be set to an empty array: **[]**.
	Extensions pulumi.StringArrayInput
}

func (AppSecAdvancedSettingsPrefetchState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecAdvancedSettingsPrefetchState)(nil)).Elem()
}

type appSecAdvancedSettingsPrefetchArgs struct {
	// . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
	AllExtensions bool `pulumi:"allExtensions"`
	// . Unique identifier of the security configuration associated with the prefetch settings being modified.
	ConfigId int `pulumi:"configId"`
	// . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
	EnableAppLayer bool `pulumi:"enableAppLayer"`
	// . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
	EnableRateControls bool `pulumi:"enableRateControls"`
	// . If `allExtensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `allExtensions` is **true**, then this argument must be set to an empty array: **[]**.
	Extensions []string `pulumi:"extensions"`
}

// The set of arguments for constructing a AppSecAdvancedSettingsPrefetch resource.
type AppSecAdvancedSettingsPrefetchArgs struct {
	// . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
	AllExtensions pulumi.BoolInput
	// . Unique identifier of the security configuration associated with the prefetch settings being modified.
	ConfigId pulumi.IntInput
	// . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
	EnableAppLayer pulumi.BoolInput
	// . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
	EnableRateControls pulumi.BoolInput
	// . If `allExtensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `allExtensions` is **true**, then this argument must be set to an empty array: **[]**.
	Extensions pulumi.StringArrayInput
}

func (AppSecAdvancedSettingsPrefetchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecAdvancedSettingsPrefetchArgs)(nil)).Elem()
}

type AppSecAdvancedSettingsPrefetchInput interface {
	pulumi.Input

	ToAppSecAdvancedSettingsPrefetchOutput() AppSecAdvancedSettingsPrefetchOutput
	ToAppSecAdvancedSettingsPrefetchOutputWithContext(ctx context.Context) AppSecAdvancedSettingsPrefetchOutput
}

func (*AppSecAdvancedSettingsPrefetch) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecAdvancedSettingsPrefetch)(nil)).Elem()
}

func (i *AppSecAdvancedSettingsPrefetch) ToAppSecAdvancedSettingsPrefetchOutput() AppSecAdvancedSettingsPrefetchOutput {
	return i.ToAppSecAdvancedSettingsPrefetchOutputWithContext(context.Background())
}

func (i *AppSecAdvancedSettingsPrefetch) ToAppSecAdvancedSettingsPrefetchOutputWithContext(ctx context.Context) AppSecAdvancedSettingsPrefetchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecAdvancedSettingsPrefetchOutput)
}

// AppSecAdvancedSettingsPrefetchArrayInput is an input type that accepts AppSecAdvancedSettingsPrefetchArray and AppSecAdvancedSettingsPrefetchArrayOutput values.
// You can construct a concrete instance of `AppSecAdvancedSettingsPrefetchArrayInput` via:
//
//	AppSecAdvancedSettingsPrefetchArray{ AppSecAdvancedSettingsPrefetchArgs{...} }
type AppSecAdvancedSettingsPrefetchArrayInput interface {
	pulumi.Input

	ToAppSecAdvancedSettingsPrefetchArrayOutput() AppSecAdvancedSettingsPrefetchArrayOutput
	ToAppSecAdvancedSettingsPrefetchArrayOutputWithContext(context.Context) AppSecAdvancedSettingsPrefetchArrayOutput
}

type AppSecAdvancedSettingsPrefetchArray []AppSecAdvancedSettingsPrefetchInput

func (AppSecAdvancedSettingsPrefetchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecAdvancedSettingsPrefetch)(nil)).Elem()
}

func (i AppSecAdvancedSettingsPrefetchArray) ToAppSecAdvancedSettingsPrefetchArrayOutput() AppSecAdvancedSettingsPrefetchArrayOutput {
	return i.ToAppSecAdvancedSettingsPrefetchArrayOutputWithContext(context.Background())
}

func (i AppSecAdvancedSettingsPrefetchArray) ToAppSecAdvancedSettingsPrefetchArrayOutputWithContext(ctx context.Context) AppSecAdvancedSettingsPrefetchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecAdvancedSettingsPrefetchArrayOutput)
}

// AppSecAdvancedSettingsPrefetchMapInput is an input type that accepts AppSecAdvancedSettingsPrefetchMap and AppSecAdvancedSettingsPrefetchMapOutput values.
// You can construct a concrete instance of `AppSecAdvancedSettingsPrefetchMapInput` via:
//
//	AppSecAdvancedSettingsPrefetchMap{ "key": AppSecAdvancedSettingsPrefetchArgs{...} }
type AppSecAdvancedSettingsPrefetchMapInput interface {
	pulumi.Input

	ToAppSecAdvancedSettingsPrefetchMapOutput() AppSecAdvancedSettingsPrefetchMapOutput
	ToAppSecAdvancedSettingsPrefetchMapOutputWithContext(context.Context) AppSecAdvancedSettingsPrefetchMapOutput
}

type AppSecAdvancedSettingsPrefetchMap map[string]AppSecAdvancedSettingsPrefetchInput

func (AppSecAdvancedSettingsPrefetchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecAdvancedSettingsPrefetch)(nil)).Elem()
}

func (i AppSecAdvancedSettingsPrefetchMap) ToAppSecAdvancedSettingsPrefetchMapOutput() AppSecAdvancedSettingsPrefetchMapOutput {
	return i.ToAppSecAdvancedSettingsPrefetchMapOutputWithContext(context.Background())
}

func (i AppSecAdvancedSettingsPrefetchMap) ToAppSecAdvancedSettingsPrefetchMapOutputWithContext(ctx context.Context) AppSecAdvancedSettingsPrefetchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecAdvancedSettingsPrefetchMapOutput)
}

type AppSecAdvancedSettingsPrefetchOutput struct{ *pulumi.OutputState }

func (AppSecAdvancedSettingsPrefetchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecAdvancedSettingsPrefetch)(nil)).Elem()
}

func (o AppSecAdvancedSettingsPrefetchOutput) ToAppSecAdvancedSettingsPrefetchOutput() AppSecAdvancedSettingsPrefetchOutput {
	return o
}

func (o AppSecAdvancedSettingsPrefetchOutput) ToAppSecAdvancedSettingsPrefetchOutputWithContext(ctx context.Context) AppSecAdvancedSettingsPrefetchOutput {
	return o
}

// . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
func (o AppSecAdvancedSettingsPrefetchOutput) AllExtensions() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppSecAdvancedSettingsPrefetch) pulumi.BoolOutput { return v.AllExtensions }).(pulumi.BoolOutput)
}

// . Unique identifier of the security configuration associated with the prefetch settings being modified.
func (o AppSecAdvancedSettingsPrefetchOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecAdvancedSettingsPrefetch) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
func (o AppSecAdvancedSettingsPrefetchOutput) EnableAppLayer() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppSecAdvancedSettingsPrefetch) pulumi.BoolOutput { return v.EnableAppLayer }).(pulumi.BoolOutput)
}

// . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
func (o AppSecAdvancedSettingsPrefetchOutput) EnableRateControls() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppSecAdvancedSettingsPrefetch) pulumi.BoolOutput { return v.EnableRateControls }).(pulumi.BoolOutput)
}

// . If `allExtensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `allExtensions` is **true**, then this argument must be set to an empty array: **[]**.
func (o AppSecAdvancedSettingsPrefetchOutput) Extensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppSecAdvancedSettingsPrefetch) pulumi.StringArrayOutput { return v.Extensions }).(pulumi.StringArrayOutput)
}

type AppSecAdvancedSettingsPrefetchArrayOutput struct{ *pulumi.OutputState }

func (AppSecAdvancedSettingsPrefetchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecAdvancedSettingsPrefetch)(nil)).Elem()
}

func (o AppSecAdvancedSettingsPrefetchArrayOutput) ToAppSecAdvancedSettingsPrefetchArrayOutput() AppSecAdvancedSettingsPrefetchArrayOutput {
	return o
}

func (o AppSecAdvancedSettingsPrefetchArrayOutput) ToAppSecAdvancedSettingsPrefetchArrayOutputWithContext(ctx context.Context) AppSecAdvancedSettingsPrefetchArrayOutput {
	return o
}

func (o AppSecAdvancedSettingsPrefetchArrayOutput) Index(i pulumi.IntInput) AppSecAdvancedSettingsPrefetchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecAdvancedSettingsPrefetch {
		return vs[0].([]*AppSecAdvancedSettingsPrefetch)[vs[1].(int)]
	}).(AppSecAdvancedSettingsPrefetchOutput)
}

type AppSecAdvancedSettingsPrefetchMapOutput struct{ *pulumi.OutputState }

func (AppSecAdvancedSettingsPrefetchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecAdvancedSettingsPrefetch)(nil)).Elem()
}

func (o AppSecAdvancedSettingsPrefetchMapOutput) ToAppSecAdvancedSettingsPrefetchMapOutput() AppSecAdvancedSettingsPrefetchMapOutput {
	return o
}

func (o AppSecAdvancedSettingsPrefetchMapOutput) ToAppSecAdvancedSettingsPrefetchMapOutputWithContext(ctx context.Context) AppSecAdvancedSettingsPrefetchMapOutput {
	return o
}

func (o AppSecAdvancedSettingsPrefetchMapOutput) MapIndex(k pulumi.StringInput) AppSecAdvancedSettingsPrefetchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecAdvancedSettingsPrefetch {
		return vs[0].(map[string]*AppSecAdvancedSettingsPrefetch)[vs[1].(string)]
	}).(AppSecAdvancedSettingsPrefetchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecAdvancedSettingsPrefetchInput)(nil)).Elem(), &AppSecAdvancedSettingsPrefetch{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecAdvancedSettingsPrefetchArrayInput)(nil)).Elem(), AppSecAdvancedSettingsPrefetchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecAdvancedSettingsPrefetchMapInput)(nil)).Elem(), AppSecAdvancedSettingsPrefetchMap{})
	pulumi.RegisterOutputType(AppSecAdvancedSettingsPrefetchOutput{})
	pulumi.RegisterOutputType(AppSecAdvancedSettingsPrefetchArrayOutput{})
	pulumi.RegisterOutputType(AppSecAdvancedSettingsPrefetchMapOutput{})
}
