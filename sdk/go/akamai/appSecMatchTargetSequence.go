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
// Specifies the order in which match targets are applied within a security configuration. As a general rule, you should process broader and more-general match targets first, gradually working your way down to more granular and highly-specific targets.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/match-targets/sequence](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)
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
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: pulumi.StringRef("Documentation"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecMatchTargetSequence(ctx, "matchTargetSequence", &akamai.AppSecMatchTargetSequenceArgs{
// 			ConfigId:            pulumi.Int(configuration.ConfigId),
// 			MatchTargetSequence: readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/match_targets_sequence.json")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecMatchTargetSequence struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the match target sequence being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Path to a JSON file containing the processing sequence for all the match targets defined for the security configuration. You can find a sample match target sequence JSON file in the [Modify match target order](https://developer.akamai.com/api/cloud_security/application_security/v1.html#matchtargetorder) section of the Application Security API documentation.
	MatchTargetSequence pulumi.StringPtrOutput `pulumi:"matchTargetSequence"`
}

// NewAppSecMatchTargetSequence registers a new resource with the given unique name, arguments, and options.
func NewAppSecMatchTargetSequence(ctx *pulumi.Context,
	name string, args *AppSecMatchTargetSequenceArgs, opts ...pulumi.ResourceOption) (*AppSecMatchTargetSequence, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	var resource AppSecMatchTargetSequence
	err := ctx.RegisterResource("akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecMatchTargetSequence gets an existing AppSecMatchTargetSequence resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecMatchTargetSequence(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecMatchTargetSequenceState, opts ...pulumi.ResourceOption) (*AppSecMatchTargetSequence, error) {
	var resource AppSecMatchTargetSequence
	err := ctx.ReadResource("akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecMatchTargetSequence resources.
type appSecMatchTargetSequenceState struct {
	// . Unique identifier of the security configuration associated with the match target sequence being modified.
	ConfigId *int `pulumi:"configId"`
	// . Path to a JSON file containing the processing sequence for all the match targets defined for the security configuration. You can find a sample match target sequence JSON file in the [Modify match target order](https://developer.akamai.com/api/cloud_security/application_security/v1.html#matchtargetorder) section of the Application Security API documentation.
	MatchTargetSequence *string `pulumi:"matchTargetSequence"`
}

type AppSecMatchTargetSequenceState struct {
	// . Unique identifier of the security configuration associated with the match target sequence being modified.
	ConfigId pulumi.IntPtrInput
	// . Path to a JSON file containing the processing sequence for all the match targets defined for the security configuration. You can find a sample match target sequence JSON file in the [Modify match target order](https://developer.akamai.com/api/cloud_security/application_security/v1.html#matchtargetorder) section of the Application Security API documentation.
	MatchTargetSequence pulumi.StringPtrInput
}

func (AppSecMatchTargetSequenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecMatchTargetSequenceState)(nil)).Elem()
}

type appSecMatchTargetSequenceArgs struct {
	// . Unique identifier of the security configuration associated with the match target sequence being modified.
	ConfigId int `pulumi:"configId"`
	// . Path to a JSON file containing the processing sequence for all the match targets defined for the security configuration. You can find a sample match target sequence JSON file in the [Modify match target order](https://developer.akamai.com/api/cloud_security/application_security/v1.html#matchtargetorder) section of the Application Security API documentation.
	MatchTargetSequence *string `pulumi:"matchTargetSequence"`
}

// The set of arguments for constructing a AppSecMatchTargetSequence resource.
type AppSecMatchTargetSequenceArgs struct {
	// . Unique identifier of the security configuration associated with the match target sequence being modified.
	ConfigId pulumi.IntInput
	// . Path to a JSON file containing the processing sequence for all the match targets defined for the security configuration. You can find a sample match target sequence JSON file in the [Modify match target order](https://developer.akamai.com/api/cloud_security/application_security/v1.html#matchtargetorder) section of the Application Security API documentation.
	MatchTargetSequence pulumi.StringPtrInput
}

func (AppSecMatchTargetSequenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecMatchTargetSequenceArgs)(nil)).Elem()
}

type AppSecMatchTargetSequenceInput interface {
	pulumi.Input

	ToAppSecMatchTargetSequenceOutput() AppSecMatchTargetSequenceOutput
	ToAppSecMatchTargetSequenceOutputWithContext(ctx context.Context) AppSecMatchTargetSequenceOutput
}

func (*AppSecMatchTargetSequence) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecMatchTargetSequence)(nil)).Elem()
}

func (i *AppSecMatchTargetSequence) ToAppSecMatchTargetSequenceOutput() AppSecMatchTargetSequenceOutput {
	return i.ToAppSecMatchTargetSequenceOutputWithContext(context.Background())
}

func (i *AppSecMatchTargetSequence) ToAppSecMatchTargetSequenceOutputWithContext(ctx context.Context) AppSecMatchTargetSequenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecMatchTargetSequenceOutput)
}

// AppSecMatchTargetSequenceArrayInput is an input type that accepts AppSecMatchTargetSequenceArray and AppSecMatchTargetSequenceArrayOutput values.
// You can construct a concrete instance of `AppSecMatchTargetSequenceArrayInput` via:
//
//          AppSecMatchTargetSequenceArray{ AppSecMatchTargetSequenceArgs{...} }
type AppSecMatchTargetSequenceArrayInput interface {
	pulumi.Input

	ToAppSecMatchTargetSequenceArrayOutput() AppSecMatchTargetSequenceArrayOutput
	ToAppSecMatchTargetSequenceArrayOutputWithContext(context.Context) AppSecMatchTargetSequenceArrayOutput
}

type AppSecMatchTargetSequenceArray []AppSecMatchTargetSequenceInput

func (AppSecMatchTargetSequenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecMatchTargetSequence)(nil)).Elem()
}

func (i AppSecMatchTargetSequenceArray) ToAppSecMatchTargetSequenceArrayOutput() AppSecMatchTargetSequenceArrayOutput {
	return i.ToAppSecMatchTargetSequenceArrayOutputWithContext(context.Background())
}

func (i AppSecMatchTargetSequenceArray) ToAppSecMatchTargetSequenceArrayOutputWithContext(ctx context.Context) AppSecMatchTargetSequenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecMatchTargetSequenceArrayOutput)
}

// AppSecMatchTargetSequenceMapInput is an input type that accepts AppSecMatchTargetSequenceMap and AppSecMatchTargetSequenceMapOutput values.
// You can construct a concrete instance of `AppSecMatchTargetSequenceMapInput` via:
//
//          AppSecMatchTargetSequenceMap{ "key": AppSecMatchTargetSequenceArgs{...} }
type AppSecMatchTargetSequenceMapInput interface {
	pulumi.Input

	ToAppSecMatchTargetSequenceMapOutput() AppSecMatchTargetSequenceMapOutput
	ToAppSecMatchTargetSequenceMapOutputWithContext(context.Context) AppSecMatchTargetSequenceMapOutput
}

type AppSecMatchTargetSequenceMap map[string]AppSecMatchTargetSequenceInput

func (AppSecMatchTargetSequenceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecMatchTargetSequence)(nil)).Elem()
}

func (i AppSecMatchTargetSequenceMap) ToAppSecMatchTargetSequenceMapOutput() AppSecMatchTargetSequenceMapOutput {
	return i.ToAppSecMatchTargetSequenceMapOutputWithContext(context.Background())
}

func (i AppSecMatchTargetSequenceMap) ToAppSecMatchTargetSequenceMapOutputWithContext(ctx context.Context) AppSecMatchTargetSequenceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecMatchTargetSequenceMapOutput)
}

type AppSecMatchTargetSequenceOutput struct{ *pulumi.OutputState }

func (AppSecMatchTargetSequenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecMatchTargetSequence)(nil)).Elem()
}

func (o AppSecMatchTargetSequenceOutput) ToAppSecMatchTargetSequenceOutput() AppSecMatchTargetSequenceOutput {
	return o
}

func (o AppSecMatchTargetSequenceOutput) ToAppSecMatchTargetSequenceOutputWithContext(ctx context.Context) AppSecMatchTargetSequenceOutput {
	return o
}

type AppSecMatchTargetSequenceArrayOutput struct{ *pulumi.OutputState }

func (AppSecMatchTargetSequenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecMatchTargetSequence)(nil)).Elem()
}

func (o AppSecMatchTargetSequenceArrayOutput) ToAppSecMatchTargetSequenceArrayOutput() AppSecMatchTargetSequenceArrayOutput {
	return o
}

func (o AppSecMatchTargetSequenceArrayOutput) ToAppSecMatchTargetSequenceArrayOutputWithContext(ctx context.Context) AppSecMatchTargetSequenceArrayOutput {
	return o
}

func (o AppSecMatchTargetSequenceArrayOutput) Index(i pulumi.IntInput) AppSecMatchTargetSequenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecMatchTargetSequence {
		return vs[0].([]*AppSecMatchTargetSequence)[vs[1].(int)]
	}).(AppSecMatchTargetSequenceOutput)
}

type AppSecMatchTargetSequenceMapOutput struct{ *pulumi.OutputState }

func (AppSecMatchTargetSequenceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecMatchTargetSequence)(nil)).Elem()
}

func (o AppSecMatchTargetSequenceMapOutput) ToAppSecMatchTargetSequenceMapOutput() AppSecMatchTargetSequenceMapOutput {
	return o
}

func (o AppSecMatchTargetSequenceMapOutput) ToAppSecMatchTargetSequenceMapOutputWithContext(ctx context.Context) AppSecMatchTargetSequenceMapOutput {
	return o
}

func (o AppSecMatchTargetSequenceMapOutput) MapIndex(k pulumi.StringInput) AppSecMatchTargetSequenceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecMatchTargetSequence {
		return vs[0].(map[string]*AppSecMatchTargetSequence)[vs[1].(string)]
	}).(AppSecMatchTargetSequenceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecMatchTargetSequenceInput)(nil)).Elem(), &AppSecMatchTargetSequence{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecMatchTargetSequenceArrayInput)(nil)).Elem(), AppSecMatchTargetSequenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecMatchTargetSequenceMapInput)(nil)).Elem(), AppSecMatchTargetSequenceMap{})
	pulumi.RegisterOutputType(AppSecMatchTargetSequenceOutput{})
	pulumi.RegisterOutputType(AppSecMatchTargetSequenceArrayOutput{})
	pulumi.RegisterOutputType(AppSecMatchTargetSequenceMapOutput{})
}
