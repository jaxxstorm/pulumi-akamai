// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `EdgeWorkersActivation` resource to activate a specific EdgeWorker version. An activation deploys the version to either the Akamai staging or production network.
//
// Before activating on production, activate on staging first. This way you can detect any problems in staging before your changes progress to production.
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
// 		_, err := akamai.NewEdgeWorkersActivation(ctx, "test", &akamai.EdgeWorkersActivationArgs{
// 			EdgeworkerId: pulumi.Int(1234),
// 			Network:      pulumi.String("STAGING"),
// 			Version:      pulumi.String("test1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EdgeWorkersActivation struct {
	pulumi.CustomResourceState

	// (Required) Unique identifier of the activation.
	ActivationId pulumi.IntOutput `pulumi:"activationId"`
	// A unique identifier for the EdgeWorker ID you want to activate.
	EdgeworkerId pulumi.IntOutput `pulumi:"edgeworkerId"`
	// The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network pulumi.StringOutput `pulumi:"network"`
	// The EdgeWorker version you want to activate.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewEdgeWorkersActivation registers a new resource with the given unique name, arguments, and options.
func NewEdgeWorkersActivation(ctx *pulumi.Context,
	name string, args *EdgeWorkersActivationArgs, opts ...pulumi.ResourceOption) (*EdgeWorkersActivation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeworkerId == nil {
		return nil, errors.New("invalid value for required argument 'EdgeworkerId'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource EdgeWorkersActivation
	err := ctx.RegisterResource("akamai:index/edgeWorkersActivation:EdgeWorkersActivation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeWorkersActivation gets an existing EdgeWorkersActivation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeWorkersActivation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeWorkersActivationState, opts ...pulumi.ResourceOption) (*EdgeWorkersActivation, error) {
	var resource EdgeWorkersActivation
	err := ctx.ReadResource("akamai:index/edgeWorkersActivation:EdgeWorkersActivation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeWorkersActivation resources.
type edgeWorkersActivationState struct {
	// (Required) Unique identifier of the activation.
	ActivationId *int `pulumi:"activationId"`
	// A unique identifier for the EdgeWorker ID you want to activate.
	EdgeworkerId *int `pulumi:"edgeworkerId"`
	// The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network *string `pulumi:"network"`
	// The EdgeWorker version you want to activate.
	Version *string `pulumi:"version"`
}

type EdgeWorkersActivationState struct {
	// (Required) Unique identifier of the activation.
	ActivationId pulumi.IntPtrInput
	// A unique identifier for the EdgeWorker ID you want to activate.
	EdgeworkerId pulumi.IntPtrInput
	// The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network pulumi.StringPtrInput
	// The EdgeWorker version you want to activate.
	Version pulumi.StringPtrInput
}

func (EdgeWorkersActivationState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeWorkersActivationState)(nil)).Elem()
}

type edgeWorkersActivationArgs struct {
	// A unique identifier for the EdgeWorker ID you want to activate.
	EdgeworkerId int `pulumi:"edgeworkerId"`
	// The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network string `pulumi:"network"`
	// The EdgeWorker version you want to activate.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a EdgeWorkersActivation resource.
type EdgeWorkersActivationArgs struct {
	// A unique identifier for the EdgeWorker ID you want to activate.
	EdgeworkerId pulumi.IntInput
	// The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
	Network pulumi.StringInput
	// The EdgeWorker version you want to activate.
	Version pulumi.StringInput
}

func (EdgeWorkersActivationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeWorkersActivationArgs)(nil)).Elem()
}

type EdgeWorkersActivationInput interface {
	pulumi.Input

	ToEdgeWorkersActivationOutput() EdgeWorkersActivationOutput
	ToEdgeWorkersActivationOutputWithContext(ctx context.Context) EdgeWorkersActivationOutput
}

func (*EdgeWorkersActivation) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeWorkersActivation)(nil)).Elem()
}

func (i *EdgeWorkersActivation) ToEdgeWorkersActivationOutput() EdgeWorkersActivationOutput {
	return i.ToEdgeWorkersActivationOutputWithContext(context.Background())
}

func (i *EdgeWorkersActivation) ToEdgeWorkersActivationOutputWithContext(ctx context.Context) EdgeWorkersActivationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeWorkersActivationOutput)
}

// EdgeWorkersActivationArrayInput is an input type that accepts EdgeWorkersActivationArray and EdgeWorkersActivationArrayOutput values.
// You can construct a concrete instance of `EdgeWorkersActivationArrayInput` via:
//
//          EdgeWorkersActivationArray{ EdgeWorkersActivationArgs{...} }
type EdgeWorkersActivationArrayInput interface {
	pulumi.Input

	ToEdgeWorkersActivationArrayOutput() EdgeWorkersActivationArrayOutput
	ToEdgeWorkersActivationArrayOutputWithContext(context.Context) EdgeWorkersActivationArrayOutput
}

type EdgeWorkersActivationArray []EdgeWorkersActivationInput

func (EdgeWorkersActivationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeWorkersActivation)(nil)).Elem()
}

func (i EdgeWorkersActivationArray) ToEdgeWorkersActivationArrayOutput() EdgeWorkersActivationArrayOutput {
	return i.ToEdgeWorkersActivationArrayOutputWithContext(context.Background())
}

func (i EdgeWorkersActivationArray) ToEdgeWorkersActivationArrayOutputWithContext(ctx context.Context) EdgeWorkersActivationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeWorkersActivationArrayOutput)
}

// EdgeWorkersActivationMapInput is an input type that accepts EdgeWorkersActivationMap and EdgeWorkersActivationMapOutput values.
// You can construct a concrete instance of `EdgeWorkersActivationMapInput` via:
//
//          EdgeWorkersActivationMap{ "key": EdgeWorkersActivationArgs{...} }
type EdgeWorkersActivationMapInput interface {
	pulumi.Input

	ToEdgeWorkersActivationMapOutput() EdgeWorkersActivationMapOutput
	ToEdgeWorkersActivationMapOutputWithContext(context.Context) EdgeWorkersActivationMapOutput
}

type EdgeWorkersActivationMap map[string]EdgeWorkersActivationInput

func (EdgeWorkersActivationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeWorkersActivation)(nil)).Elem()
}

func (i EdgeWorkersActivationMap) ToEdgeWorkersActivationMapOutput() EdgeWorkersActivationMapOutput {
	return i.ToEdgeWorkersActivationMapOutputWithContext(context.Background())
}

func (i EdgeWorkersActivationMap) ToEdgeWorkersActivationMapOutputWithContext(ctx context.Context) EdgeWorkersActivationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeWorkersActivationMapOutput)
}

type EdgeWorkersActivationOutput struct{ *pulumi.OutputState }

func (EdgeWorkersActivationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeWorkersActivation)(nil)).Elem()
}

func (o EdgeWorkersActivationOutput) ToEdgeWorkersActivationOutput() EdgeWorkersActivationOutput {
	return o
}

func (o EdgeWorkersActivationOutput) ToEdgeWorkersActivationOutputWithContext(ctx context.Context) EdgeWorkersActivationOutput {
	return o
}

type EdgeWorkersActivationArrayOutput struct{ *pulumi.OutputState }

func (EdgeWorkersActivationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeWorkersActivation)(nil)).Elem()
}

func (o EdgeWorkersActivationArrayOutput) ToEdgeWorkersActivationArrayOutput() EdgeWorkersActivationArrayOutput {
	return o
}

func (o EdgeWorkersActivationArrayOutput) ToEdgeWorkersActivationArrayOutputWithContext(ctx context.Context) EdgeWorkersActivationArrayOutput {
	return o
}

func (o EdgeWorkersActivationArrayOutput) Index(i pulumi.IntInput) EdgeWorkersActivationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EdgeWorkersActivation {
		return vs[0].([]*EdgeWorkersActivation)[vs[1].(int)]
	}).(EdgeWorkersActivationOutput)
}

type EdgeWorkersActivationMapOutput struct{ *pulumi.OutputState }

func (EdgeWorkersActivationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeWorkersActivation)(nil)).Elem()
}

func (o EdgeWorkersActivationMapOutput) ToEdgeWorkersActivationMapOutput() EdgeWorkersActivationMapOutput {
	return o
}

func (o EdgeWorkersActivationMapOutput) ToEdgeWorkersActivationMapOutputWithContext(ctx context.Context) EdgeWorkersActivationMapOutput {
	return o
}

func (o EdgeWorkersActivationMapOutput) MapIndex(k pulumi.StringInput) EdgeWorkersActivationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EdgeWorkersActivation {
		return vs[0].(map[string]*EdgeWorkersActivation)[vs[1].(string)]
	}).(EdgeWorkersActivationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeWorkersActivationInput)(nil)).Elem(), &EdgeWorkersActivation{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeWorkersActivationArrayInput)(nil)).Elem(), EdgeWorkersActivationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeWorkersActivationMapInput)(nil)).Elem(), EdgeWorkersActivationMap{})
	pulumi.RegisterOutputType(EdgeWorkersActivationOutput{})
	pulumi.RegisterOutputType(EdgeWorkersActivationArrayOutput{})
	pulumi.RegisterOutputType(EdgeWorkersActivationMapOutput{})
}
