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
// Creates or modifies a reputation profile.
// Reputation profiles grade the security risk of an IP address based on previous activities associated with that address.
// Depending on the reputation score and how your configuration has been set up, requests from a specific IP address can trigger an alert or even be blocked.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/reputation-profiles](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles)
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
// 		reputationProfile, err := akamai.NewAppSecReputationProfile(ctx, "reputationProfile", &akamai.AppSecReputationProfileArgs{
// 			ConfigId:          pulumi.Int(configuration.ConfigId),
// 			ReputationProfile: readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/reputation_profile.json")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("reputationProfileId", reputationProfile.ReputationProfileId)
// 		return nil
// 	})
// }
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `reputationProfileId`. ID of the newly-created or newly-modified reputation profile.
type AppSecReputationProfile struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the reputation profile being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Path to a JSON file containing a definition of the reputation profile. You can view a sample JSON file in the [Create a reputation profile](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles) section of the Application Security API documentation.
	ReputationProfile   pulumi.StringOutput `pulumi:"reputationProfile"`
	ReputationProfileId pulumi.IntOutput    `pulumi:"reputationProfileId"`
}

// NewAppSecReputationProfile registers a new resource with the given unique name, arguments, and options.
func NewAppSecReputationProfile(ctx *pulumi.Context,
	name string, args *AppSecReputationProfileArgs, opts ...pulumi.ResourceOption) (*AppSecReputationProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.ReputationProfile == nil {
		return nil, errors.New("invalid value for required argument 'ReputationProfile'")
	}
	var resource AppSecReputationProfile
	err := ctx.RegisterResource("akamai:index/appSecReputationProfile:AppSecReputationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecReputationProfile gets an existing AppSecReputationProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecReputationProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecReputationProfileState, opts ...pulumi.ResourceOption) (*AppSecReputationProfile, error) {
	var resource AppSecReputationProfile
	err := ctx.ReadResource("akamai:index/appSecReputationProfile:AppSecReputationProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecReputationProfile resources.
type appSecReputationProfileState struct {
	// . Unique identifier of the security configuration associated with the reputation profile being modified.
	ConfigId *int `pulumi:"configId"`
	// . Path to a JSON file containing a definition of the reputation profile. You can view a sample JSON file in the [Create a reputation profile](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles) section of the Application Security API documentation.
	ReputationProfile   *string `pulumi:"reputationProfile"`
	ReputationProfileId *int    `pulumi:"reputationProfileId"`
}

type AppSecReputationProfileState struct {
	// . Unique identifier of the security configuration associated with the reputation profile being modified.
	ConfigId pulumi.IntPtrInput
	// . Path to a JSON file containing a definition of the reputation profile. You can view a sample JSON file in the [Create a reputation profile](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles) section of the Application Security API documentation.
	ReputationProfile   pulumi.StringPtrInput
	ReputationProfileId pulumi.IntPtrInput
}

func (AppSecReputationProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecReputationProfileState)(nil)).Elem()
}

type appSecReputationProfileArgs struct {
	// . Unique identifier of the security configuration associated with the reputation profile being modified.
	ConfigId int `pulumi:"configId"`
	// . Path to a JSON file containing a definition of the reputation profile. You can view a sample JSON file in the [Create a reputation profile](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles) section of the Application Security API documentation.
	ReputationProfile string `pulumi:"reputationProfile"`
}

// The set of arguments for constructing a AppSecReputationProfile resource.
type AppSecReputationProfileArgs struct {
	// . Unique identifier of the security configuration associated with the reputation profile being modified.
	ConfigId pulumi.IntInput
	// . Path to a JSON file containing a definition of the reputation profile. You can view a sample JSON file in the [Create a reputation profile](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles) section of the Application Security API documentation.
	ReputationProfile pulumi.StringInput
}

func (AppSecReputationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecReputationProfileArgs)(nil)).Elem()
}

type AppSecReputationProfileInput interface {
	pulumi.Input

	ToAppSecReputationProfileOutput() AppSecReputationProfileOutput
	ToAppSecReputationProfileOutputWithContext(ctx context.Context) AppSecReputationProfileOutput
}

func (*AppSecReputationProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecReputationProfile)(nil)).Elem()
}

func (i *AppSecReputationProfile) ToAppSecReputationProfileOutput() AppSecReputationProfileOutput {
	return i.ToAppSecReputationProfileOutputWithContext(context.Background())
}

func (i *AppSecReputationProfile) ToAppSecReputationProfileOutputWithContext(ctx context.Context) AppSecReputationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProfileOutput)
}

// AppSecReputationProfileArrayInput is an input type that accepts AppSecReputationProfileArray and AppSecReputationProfileArrayOutput values.
// You can construct a concrete instance of `AppSecReputationProfileArrayInput` via:
//
//          AppSecReputationProfileArray{ AppSecReputationProfileArgs{...} }
type AppSecReputationProfileArrayInput interface {
	pulumi.Input

	ToAppSecReputationProfileArrayOutput() AppSecReputationProfileArrayOutput
	ToAppSecReputationProfileArrayOutputWithContext(context.Context) AppSecReputationProfileArrayOutput
}

type AppSecReputationProfileArray []AppSecReputationProfileInput

func (AppSecReputationProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecReputationProfile)(nil)).Elem()
}

func (i AppSecReputationProfileArray) ToAppSecReputationProfileArrayOutput() AppSecReputationProfileArrayOutput {
	return i.ToAppSecReputationProfileArrayOutputWithContext(context.Background())
}

func (i AppSecReputationProfileArray) ToAppSecReputationProfileArrayOutputWithContext(ctx context.Context) AppSecReputationProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProfileArrayOutput)
}

// AppSecReputationProfileMapInput is an input type that accepts AppSecReputationProfileMap and AppSecReputationProfileMapOutput values.
// You can construct a concrete instance of `AppSecReputationProfileMapInput` via:
//
//          AppSecReputationProfileMap{ "key": AppSecReputationProfileArgs{...} }
type AppSecReputationProfileMapInput interface {
	pulumi.Input

	ToAppSecReputationProfileMapOutput() AppSecReputationProfileMapOutput
	ToAppSecReputationProfileMapOutputWithContext(context.Context) AppSecReputationProfileMapOutput
}

type AppSecReputationProfileMap map[string]AppSecReputationProfileInput

func (AppSecReputationProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecReputationProfile)(nil)).Elem()
}

func (i AppSecReputationProfileMap) ToAppSecReputationProfileMapOutput() AppSecReputationProfileMapOutput {
	return i.ToAppSecReputationProfileMapOutputWithContext(context.Background())
}

func (i AppSecReputationProfileMap) ToAppSecReputationProfileMapOutputWithContext(ctx context.Context) AppSecReputationProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecReputationProfileMapOutput)
}

type AppSecReputationProfileOutput struct{ *pulumi.OutputState }

func (AppSecReputationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecReputationProfile)(nil)).Elem()
}

func (o AppSecReputationProfileOutput) ToAppSecReputationProfileOutput() AppSecReputationProfileOutput {
	return o
}

func (o AppSecReputationProfileOutput) ToAppSecReputationProfileOutputWithContext(ctx context.Context) AppSecReputationProfileOutput {
	return o
}

type AppSecReputationProfileArrayOutput struct{ *pulumi.OutputState }

func (AppSecReputationProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecReputationProfile)(nil)).Elem()
}

func (o AppSecReputationProfileArrayOutput) ToAppSecReputationProfileArrayOutput() AppSecReputationProfileArrayOutput {
	return o
}

func (o AppSecReputationProfileArrayOutput) ToAppSecReputationProfileArrayOutputWithContext(ctx context.Context) AppSecReputationProfileArrayOutput {
	return o
}

func (o AppSecReputationProfileArrayOutput) Index(i pulumi.IntInput) AppSecReputationProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecReputationProfile {
		return vs[0].([]*AppSecReputationProfile)[vs[1].(int)]
	}).(AppSecReputationProfileOutput)
}

type AppSecReputationProfileMapOutput struct{ *pulumi.OutputState }

func (AppSecReputationProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecReputationProfile)(nil)).Elem()
}

func (o AppSecReputationProfileMapOutput) ToAppSecReputationProfileMapOutput() AppSecReputationProfileMapOutput {
	return o
}

func (o AppSecReputationProfileMapOutput) ToAppSecReputationProfileMapOutputWithContext(ctx context.Context) AppSecReputationProfileMapOutput {
	return o
}

func (o AppSecReputationProfileMapOutput) MapIndex(k pulumi.StringInput) AppSecReputationProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecReputationProfile {
		return vs[0].(map[string]*AppSecReputationProfile)[vs[1].(string)]
	}).(AppSecReputationProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProfileInput)(nil)).Elem(), &AppSecReputationProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProfileArrayInput)(nil)).Elem(), AppSecReputationProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecReputationProfileMapInput)(nil)).Elem(), AppSecReputationProfileMap{})
	pulumi.RegisterOutputType(AppSecReputationProfileOutput{})
	pulumi.RegisterOutputType(AppSecReputationProfileArrayOutput{})
	pulumi.RegisterOutputType(AppSecReputationProfileMapOutput{})
}
