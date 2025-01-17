// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `CloudletsApplicationLoadBalancerActivation` resource to activate the Application Load Balancer Cloudlet configuration. An activation deploys the configuration version to either the Akamai staging or production network. You can activate a specific version multiple times if you need to.
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
//
//	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := akamai.NewCloudletsApplicationLoadBalancerActivation(ctx, "example", &akamai.CloudletsApplicationLoadBalancerActivationArgs{
//				OriginId: pulumi.String("alb_test_1"),
//				Network:  pulumi.String("staging"),
//				Version:  pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("status", example.Status)
//			return nil
//		})
//	}
//
// ```
type CloudletsApplicationLoadBalancerActivation struct {
	pulumi.CustomResourceState

	// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
	Network pulumi.StringOutput `pulumi:"network"`
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId pulumi.StringOutput `pulumi:"originId"`
	// The activation status for this load balancing configuration.
	Status pulumi.StringOutput `pulumi:"status"`
	// The Application Load Balancer Cloudlet configuration version you want to activate.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewCloudletsApplicationLoadBalancerActivation registers a new resource with the given unique name, arguments, and options.
func NewCloudletsApplicationLoadBalancerActivation(ctx *pulumi.Context,
	name string, args *CloudletsApplicationLoadBalancerActivationArgs, opts ...pulumi.ResourceOption) (*CloudletsApplicationLoadBalancerActivation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.OriginId == nil {
		return nil, errors.New("invalid value for required argument 'OriginId'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource CloudletsApplicationLoadBalancerActivation
	err := ctx.RegisterResource("akamai:index/cloudletsApplicationLoadBalancerActivation:CloudletsApplicationLoadBalancerActivation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudletsApplicationLoadBalancerActivation gets an existing CloudletsApplicationLoadBalancerActivation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudletsApplicationLoadBalancerActivation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudletsApplicationLoadBalancerActivationState, opts ...pulumi.ResourceOption) (*CloudletsApplicationLoadBalancerActivation, error) {
	var resource CloudletsApplicationLoadBalancerActivation
	err := ctx.ReadResource("akamai:index/cloudletsApplicationLoadBalancerActivation:CloudletsApplicationLoadBalancerActivation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudletsApplicationLoadBalancerActivation resources.
type cloudletsApplicationLoadBalancerActivationState struct {
	// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
	Network *string `pulumi:"network"`
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId *string `pulumi:"originId"`
	// The activation status for this load balancing configuration.
	Status *string `pulumi:"status"`
	// The Application Load Balancer Cloudlet configuration version you want to activate.
	Version *int `pulumi:"version"`
}

type CloudletsApplicationLoadBalancerActivationState struct {
	// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
	Network pulumi.StringPtrInput
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId pulumi.StringPtrInput
	// The activation status for this load balancing configuration.
	Status pulumi.StringPtrInput
	// The Application Load Balancer Cloudlet configuration version you want to activate.
	Version pulumi.IntPtrInput
}

func (CloudletsApplicationLoadBalancerActivationState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudletsApplicationLoadBalancerActivationState)(nil)).Elem()
}

type cloudletsApplicationLoadBalancerActivationArgs struct {
	// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
	Network string `pulumi:"network"`
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId string `pulumi:"originId"`
	// The Application Load Balancer Cloudlet configuration version you want to activate.
	Version int `pulumi:"version"`
}

// The set of arguments for constructing a CloudletsApplicationLoadBalancerActivation resource.
type CloudletsApplicationLoadBalancerActivationArgs struct {
	// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
	Network pulumi.StringInput
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId pulumi.StringInput
	// The Application Load Balancer Cloudlet configuration version you want to activate.
	Version pulumi.IntInput
}

func (CloudletsApplicationLoadBalancerActivationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudletsApplicationLoadBalancerActivationArgs)(nil)).Elem()
}

type CloudletsApplicationLoadBalancerActivationInput interface {
	pulumi.Input

	ToCloudletsApplicationLoadBalancerActivationOutput() CloudletsApplicationLoadBalancerActivationOutput
	ToCloudletsApplicationLoadBalancerActivationOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerActivationOutput
}

func (*CloudletsApplicationLoadBalancerActivation) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudletsApplicationLoadBalancerActivation)(nil)).Elem()
}

func (i *CloudletsApplicationLoadBalancerActivation) ToCloudletsApplicationLoadBalancerActivationOutput() CloudletsApplicationLoadBalancerActivationOutput {
	return i.ToCloudletsApplicationLoadBalancerActivationOutputWithContext(context.Background())
}

func (i *CloudletsApplicationLoadBalancerActivation) ToCloudletsApplicationLoadBalancerActivationOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerActivationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsApplicationLoadBalancerActivationOutput)
}

// CloudletsApplicationLoadBalancerActivationArrayInput is an input type that accepts CloudletsApplicationLoadBalancerActivationArray and CloudletsApplicationLoadBalancerActivationArrayOutput values.
// You can construct a concrete instance of `CloudletsApplicationLoadBalancerActivationArrayInput` via:
//
//	CloudletsApplicationLoadBalancerActivationArray{ CloudletsApplicationLoadBalancerActivationArgs{...} }
type CloudletsApplicationLoadBalancerActivationArrayInput interface {
	pulumi.Input

	ToCloudletsApplicationLoadBalancerActivationArrayOutput() CloudletsApplicationLoadBalancerActivationArrayOutput
	ToCloudletsApplicationLoadBalancerActivationArrayOutputWithContext(context.Context) CloudletsApplicationLoadBalancerActivationArrayOutput
}

type CloudletsApplicationLoadBalancerActivationArray []CloudletsApplicationLoadBalancerActivationInput

func (CloudletsApplicationLoadBalancerActivationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudletsApplicationLoadBalancerActivation)(nil)).Elem()
}

func (i CloudletsApplicationLoadBalancerActivationArray) ToCloudletsApplicationLoadBalancerActivationArrayOutput() CloudletsApplicationLoadBalancerActivationArrayOutput {
	return i.ToCloudletsApplicationLoadBalancerActivationArrayOutputWithContext(context.Background())
}

func (i CloudletsApplicationLoadBalancerActivationArray) ToCloudletsApplicationLoadBalancerActivationArrayOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerActivationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsApplicationLoadBalancerActivationArrayOutput)
}

// CloudletsApplicationLoadBalancerActivationMapInput is an input type that accepts CloudletsApplicationLoadBalancerActivationMap and CloudletsApplicationLoadBalancerActivationMapOutput values.
// You can construct a concrete instance of `CloudletsApplicationLoadBalancerActivationMapInput` via:
//
//	CloudletsApplicationLoadBalancerActivationMap{ "key": CloudletsApplicationLoadBalancerActivationArgs{...} }
type CloudletsApplicationLoadBalancerActivationMapInput interface {
	pulumi.Input

	ToCloudletsApplicationLoadBalancerActivationMapOutput() CloudletsApplicationLoadBalancerActivationMapOutput
	ToCloudletsApplicationLoadBalancerActivationMapOutputWithContext(context.Context) CloudletsApplicationLoadBalancerActivationMapOutput
}

type CloudletsApplicationLoadBalancerActivationMap map[string]CloudletsApplicationLoadBalancerActivationInput

func (CloudletsApplicationLoadBalancerActivationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudletsApplicationLoadBalancerActivation)(nil)).Elem()
}

func (i CloudletsApplicationLoadBalancerActivationMap) ToCloudletsApplicationLoadBalancerActivationMapOutput() CloudletsApplicationLoadBalancerActivationMapOutput {
	return i.ToCloudletsApplicationLoadBalancerActivationMapOutputWithContext(context.Background())
}

func (i CloudletsApplicationLoadBalancerActivationMap) ToCloudletsApplicationLoadBalancerActivationMapOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerActivationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsApplicationLoadBalancerActivationMapOutput)
}

type CloudletsApplicationLoadBalancerActivationOutput struct{ *pulumi.OutputState }

func (CloudletsApplicationLoadBalancerActivationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudletsApplicationLoadBalancerActivation)(nil)).Elem()
}

func (o CloudletsApplicationLoadBalancerActivationOutput) ToCloudletsApplicationLoadBalancerActivationOutput() CloudletsApplicationLoadBalancerActivationOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerActivationOutput) ToCloudletsApplicationLoadBalancerActivationOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerActivationOutput {
	return o
}

// The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
func (o CloudletsApplicationLoadBalancerActivationOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudletsApplicationLoadBalancerActivation) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
func (o CloudletsApplicationLoadBalancerActivationOutput) OriginId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudletsApplicationLoadBalancerActivation) pulumi.StringOutput { return v.OriginId }).(pulumi.StringOutput)
}

// The activation status for this load balancing configuration.
func (o CloudletsApplicationLoadBalancerActivationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudletsApplicationLoadBalancerActivation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The Application Load Balancer Cloudlet configuration version you want to activate.
func (o CloudletsApplicationLoadBalancerActivationOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudletsApplicationLoadBalancerActivation) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type CloudletsApplicationLoadBalancerActivationArrayOutput struct{ *pulumi.OutputState }

func (CloudletsApplicationLoadBalancerActivationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudletsApplicationLoadBalancerActivation)(nil)).Elem()
}

func (o CloudletsApplicationLoadBalancerActivationArrayOutput) ToCloudletsApplicationLoadBalancerActivationArrayOutput() CloudletsApplicationLoadBalancerActivationArrayOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerActivationArrayOutput) ToCloudletsApplicationLoadBalancerActivationArrayOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerActivationArrayOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerActivationArrayOutput) Index(i pulumi.IntInput) CloudletsApplicationLoadBalancerActivationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudletsApplicationLoadBalancerActivation {
		return vs[0].([]*CloudletsApplicationLoadBalancerActivation)[vs[1].(int)]
	}).(CloudletsApplicationLoadBalancerActivationOutput)
}

type CloudletsApplicationLoadBalancerActivationMapOutput struct{ *pulumi.OutputState }

func (CloudletsApplicationLoadBalancerActivationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudletsApplicationLoadBalancerActivation)(nil)).Elem()
}

func (o CloudletsApplicationLoadBalancerActivationMapOutput) ToCloudletsApplicationLoadBalancerActivationMapOutput() CloudletsApplicationLoadBalancerActivationMapOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerActivationMapOutput) ToCloudletsApplicationLoadBalancerActivationMapOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerActivationMapOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerActivationMapOutput) MapIndex(k pulumi.StringInput) CloudletsApplicationLoadBalancerActivationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudletsApplicationLoadBalancerActivation {
		return vs[0].(map[string]*CloudletsApplicationLoadBalancerActivation)[vs[1].(string)]
	}).(CloudletsApplicationLoadBalancerActivationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsApplicationLoadBalancerActivationInput)(nil)).Elem(), &CloudletsApplicationLoadBalancerActivation{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsApplicationLoadBalancerActivationArrayInput)(nil)).Elem(), CloudletsApplicationLoadBalancerActivationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsApplicationLoadBalancerActivationMapInput)(nil)).Elem(), CloudletsApplicationLoadBalancerActivationMap{})
	pulumi.RegisterOutputType(CloudletsApplicationLoadBalancerActivationOutput{})
	pulumi.RegisterOutputType(CloudletsApplicationLoadBalancerActivationArrayOutput{})
	pulumi.RegisterOutputType(CloudletsApplicationLoadBalancerActivationMapOutput{})
}
