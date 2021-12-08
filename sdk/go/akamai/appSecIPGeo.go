// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy
//
// Modifies the method used for firewall blocking, and manages the network lists used for IP/Geo firewall blocking.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/ip-geo-firewall](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putipgeofirewall)
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
// 		ipGeoBlock, err := akamai.NewAppSecIPGeo(ctx, "ipGeoBlock", &akamai.AppSecIPGeoArgs{
// 			ConfigId:         pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId: pulumi.String("gms1_134637"),
// 			Mode:             pulumi.String("block"),
// 			GeoNetworkLists: pulumi.StringArray{
// 				pulumi.String("06038_GEO_TEST"),
// 			},
// 			IpNetworkLists: pulumi.StringArray{
// 				pulumi.String("56921_TEST"),
// 			},
// 			ExceptionIpNetworkLists: pulumi.StringArray{
// 				pulumi.String("07126_EXCEPTION_TEST"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ipGeoAllow, err := akamai.NewAppSecIPGeo(ctx, "ipGeoAllow", &akamai.AppSecIPGeoArgs{
// 			ConfigId:         pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId: pulumi.String("gms1-090334"),
// 			Mode:             pulumi.String("allow"),
// 			ExceptionIpNetworkLists: pulumi.StringArray{
// 				pulumi.String("07126_EXCEPTION_TEST"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ipGeoModeBlock", ipGeoBlock.Mode)
// 		ctx.Export("blockGeoNetworkLists", ipGeoBlock.GeoNetworkLists)
// 		ctx.Export("blockIpNetworkLists", ipGeoBlock.IpNetworkLists)
// 		ctx.Export("blockExceptionIpNetworkLists", ipGeoBlock.ExceptionIpNetworkLists)
// 		ctx.Export("ipGeoModeAllow", ipGeoAllow.Mode)
// 		ctx.Export("allowExceptionIpNetworkLists", ipGeoAllow.ExceptionIpNetworkLists)
// 		return nil
// 	})
// }
// ```
type AppSecIPGeo struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
	ExceptionIpNetworkLists pulumi.StringArrayOutput `pulumi:"exceptionIpNetworkLists"`
	// . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
	GeoNetworkLists pulumi.StringArrayOutput `pulumi:"geoNetworkLists"`
	// . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall..
	IpNetworkLists pulumi.StringArrayOutput `pulumi:"ipNetworkLists"`
	// . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// . Unique identifier of the security policy associated with the IP/Geo lists being modified.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecIPGeo registers a new resource with the given unique name, arguments, and options.
func NewAppSecIPGeo(ctx *pulumi.Context,
	name string, args *AppSecIPGeoArgs, opts ...pulumi.ResourceOption) (*AppSecIPGeo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	var resource AppSecIPGeo
	err := ctx.RegisterResource("akamai:index/appSecIPGeo:AppSecIPGeo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecIPGeo gets an existing AppSecIPGeo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecIPGeo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecIPGeoState, opts ...pulumi.ResourceOption) (*AppSecIPGeo, error) {
	var resource AppSecIPGeo
	err := ctx.ReadResource("akamai:index/appSecIPGeo:AppSecIPGeo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecIPGeo resources.
type appSecIPGeoState struct {
	// . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
	ConfigId *int `pulumi:"configId"`
	// . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
	ExceptionIpNetworkLists []string `pulumi:"exceptionIpNetworkLists"`
	// . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
	GeoNetworkLists []string `pulumi:"geoNetworkLists"`
	// . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall..
	IpNetworkLists []string `pulumi:"ipNetworkLists"`
	// . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
	Mode *string `pulumi:"mode"`
	// . Unique identifier of the security policy associated with the IP/Geo lists being modified.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecIPGeoState struct {
	// . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
	ConfigId pulumi.IntPtrInput
	// . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
	ExceptionIpNetworkLists pulumi.StringArrayInput
	// . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
	GeoNetworkLists pulumi.StringArrayInput
	// . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall..
	IpNetworkLists pulumi.StringArrayInput
	// . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
	Mode pulumi.StringPtrInput
	// . Unique identifier of the security policy associated with the IP/Geo lists being modified.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecIPGeoState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecIPGeoState)(nil)).Elem()
}

type appSecIPGeoArgs struct {
	// . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
	ConfigId int `pulumi:"configId"`
	// . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
	ExceptionIpNetworkLists []string `pulumi:"exceptionIpNetworkLists"`
	// . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
	GeoNetworkLists []string `pulumi:"geoNetworkLists"`
	// . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall..
	IpNetworkLists []string `pulumi:"ipNetworkLists"`
	// . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
	Mode string `pulumi:"mode"`
	// . Unique identifier of the security policy associated with the IP/Geo lists being modified.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecIPGeo resource.
type AppSecIPGeoArgs struct {
	// . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
	ConfigId pulumi.IntInput
	// . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
	ExceptionIpNetworkLists pulumi.StringArrayInput
	// . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
	GeoNetworkLists pulumi.StringArrayInput
	// . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall..
	IpNetworkLists pulumi.StringArrayInput
	// . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
	Mode pulumi.StringInput
	// . Unique identifier of the security policy associated with the IP/Geo lists being modified.
	SecurityPolicyId pulumi.StringInput
}

func (AppSecIPGeoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecIPGeoArgs)(nil)).Elem()
}

type AppSecIPGeoInput interface {
	pulumi.Input

	ToAppSecIPGeoOutput() AppSecIPGeoOutput
	ToAppSecIPGeoOutputWithContext(ctx context.Context) AppSecIPGeoOutput
}

func (*AppSecIPGeo) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecIPGeo)(nil))
}

func (i *AppSecIPGeo) ToAppSecIPGeoOutput() AppSecIPGeoOutput {
	return i.ToAppSecIPGeoOutputWithContext(context.Background())
}

func (i *AppSecIPGeo) ToAppSecIPGeoOutputWithContext(ctx context.Context) AppSecIPGeoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecIPGeoOutput)
}

func (i *AppSecIPGeo) ToAppSecIPGeoPtrOutput() AppSecIPGeoPtrOutput {
	return i.ToAppSecIPGeoPtrOutputWithContext(context.Background())
}

func (i *AppSecIPGeo) ToAppSecIPGeoPtrOutputWithContext(ctx context.Context) AppSecIPGeoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecIPGeoPtrOutput)
}

type AppSecIPGeoPtrInput interface {
	pulumi.Input

	ToAppSecIPGeoPtrOutput() AppSecIPGeoPtrOutput
	ToAppSecIPGeoPtrOutputWithContext(ctx context.Context) AppSecIPGeoPtrOutput
}

type appSecIPGeoPtrType AppSecIPGeoArgs

func (*appSecIPGeoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecIPGeo)(nil))
}

func (i *appSecIPGeoPtrType) ToAppSecIPGeoPtrOutput() AppSecIPGeoPtrOutput {
	return i.ToAppSecIPGeoPtrOutputWithContext(context.Background())
}

func (i *appSecIPGeoPtrType) ToAppSecIPGeoPtrOutputWithContext(ctx context.Context) AppSecIPGeoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecIPGeoPtrOutput)
}

// AppSecIPGeoArrayInput is an input type that accepts AppSecIPGeoArray and AppSecIPGeoArrayOutput values.
// You can construct a concrete instance of `AppSecIPGeoArrayInput` via:
//
//          AppSecIPGeoArray{ AppSecIPGeoArgs{...} }
type AppSecIPGeoArrayInput interface {
	pulumi.Input

	ToAppSecIPGeoArrayOutput() AppSecIPGeoArrayOutput
	ToAppSecIPGeoArrayOutputWithContext(context.Context) AppSecIPGeoArrayOutput
}

type AppSecIPGeoArray []AppSecIPGeoInput

func (AppSecIPGeoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecIPGeo)(nil)).Elem()
}

func (i AppSecIPGeoArray) ToAppSecIPGeoArrayOutput() AppSecIPGeoArrayOutput {
	return i.ToAppSecIPGeoArrayOutputWithContext(context.Background())
}

func (i AppSecIPGeoArray) ToAppSecIPGeoArrayOutputWithContext(ctx context.Context) AppSecIPGeoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecIPGeoArrayOutput)
}

// AppSecIPGeoMapInput is an input type that accepts AppSecIPGeoMap and AppSecIPGeoMapOutput values.
// You can construct a concrete instance of `AppSecIPGeoMapInput` via:
//
//          AppSecIPGeoMap{ "key": AppSecIPGeoArgs{...} }
type AppSecIPGeoMapInput interface {
	pulumi.Input

	ToAppSecIPGeoMapOutput() AppSecIPGeoMapOutput
	ToAppSecIPGeoMapOutputWithContext(context.Context) AppSecIPGeoMapOutput
}

type AppSecIPGeoMap map[string]AppSecIPGeoInput

func (AppSecIPGeoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecIPGeo)(nil)).Elem()
}

func (i AppSecIPGeoMap) ToAppSecIPGeoMapOutput() AppSecIPGeoMapOutput {
	return i.ToAppSecIPGeoMapOutputWithContext(context.Background())
}

func (i AppSecIPGeoMap) ToAppSecIPGeoMapOutputWithContext(ctx context.Context) AppSecIPGeoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecIPGeoMapOutput)
}

type AppSecIPGeoOutput struct{ *pulumi.OutputState }

func (AppSecIPGeoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecIPGeo)(nil))
}

func (o AppSecIPGeoOutput) ToAppSecIPGeoOutput() AppSecIPGeoOutput {
	return o
}

func (o AppSecIPGeoOutput) ToAppSecIPGeoOutputWithContext(ctx context.Context) AppSecIPGeoOutput {
	return o
}

func (o AppSecIPGeoOutput) ToAppSecIPGeoPtrOutput() AppSecIPGeoPtrOutput {
	return o.ToAppSecIPGeoPtrOutputWithContext(context.Background())
}

func (o AppSecIPGeoOutput) ToAppSecIPGeoPtrOutputWithContext(ctx context.Context) AppSecIPGeoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSecIPGeo) *AppSecIPGeo {
		return &v
	}).(AppSecIPGeoPtrOutput)
}

type AppSecIPGeoPtrOutput struct{ *pulumi.OutputState }

func (AppSecIPGeoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecIPGeo)(nil))
}

func (o AppSecIPGeoPtrOutput) ToAppSecIPGeoPtrOutput() AppSecIPGeoPtrOutput {
	return o
}

func (o AppSecIPGeoPtrOutput) ToAppSecIPGeoPtrOutputWithContext(ctx context.Context) AppSecIPGeoPtrOutput {
	return o
}

func (o AppSecIPGeoPtrOutput) Elem() AppSecIPGeoOutput {
	return o.ApplyT(func(v *AppSecIPGeo) AppSecIPGeo {
		if v != nil {
			return *v
		}
		var ret AppSecIPGeo
		return ret
	}).(AppSecIPGeoOutput)
}

type AppSecIPGeoArrayOutput struct{ *pulumi.OutputState }

func (AppSecIPGeoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecIPGeo)(nil))
}

func (o AppSecIPGeoArrayOutput) ToAppSecIPGeoArrayOutput() AppSecIPGeoArrayOutput {
	return o
}

func (o AppSecIPGeoArrayOutput) ToAppSecIPGeoArrayOutputWithContext(ctx context.Context) AppSecIPGeoArrayOutput {
	return o
}

func (o AppSecIPGeoArrayOutput) Index(i pulumi.IntInput) AppSecIPGeoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecIPGeo {
		return vs[0].([]AppSecIPGeo)[vs[1].(int)]
	}).(AppSecIPGeoOutput)
}

type AppSecIPGeoMapOutput struct{ *pulumi.OutputState }

func (AppSecIPGeoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecIPGeo)(nil))
}

func (o AppSecIPGeoMapOutput) ToAppSecIPGeoMapOutput() AppSecIPGeoMapOutput {
	return o
}

func (o AppSecIPGeoMapOutput) ToAppSecIPGeoMapOutputWithContext(ctx context.Context) AppSecIPGeoMapOutput {
	return o
}

func (o AppSecIPGeoMapOutput) MapIndex(k pulumi.StringInput) AppSecIPGeoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecIPGeo {
		return vs[0].(map[string]AppSecIPGeo)[vs[1].(string)]
	}).(AppSecIPGeoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecIPGeoInput)(nil)).Elem(), &AppSecIPGeo{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecIPGeoPtrInput)(nil)).Elem(), &AppSecIPGeo{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecIPGeoArrayInput)(nil)).Elem(), AppSecIPGeoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecIPGeoMapInput)(nil)).Elem(), AppSecIPGeoMap{})
	pulumi.RegisterOutputType(AppSecIPGeoOutput{})
	pulumi.RegisterOutputType(AppSecIPGeoPtrOutput{})
	pulumi.RegisterOutputType(AppSecIPGeoArrayOutput{})
	pulumi.RegisterOutputType(AppSecIPGeoMapOutput{})
}
