// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `CloudletsPolicyActivation` resource to activate a specific version of a Cloudlet policy. An activation deploys the version to either the Akamai staging or production network. You can activate a specific version multiple times if you need to.
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
// 		_, err := akamai.NewCloudletsPolicyActivation(ctx, "example", &akamai.CloudletsPolicyActivationArgs{
// 			AssociatedProperties: pulumi.StringArray{
// 				pulumi.String("Property_1"),
// 				pulumi.String("Property_2"),
// 				pulumi.String("Property_3"),
// 			},
// 			Network:  pulumi.String("staging"),
// 			PolicyId: pulumi.Int(1234),
// 			Version:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// If you're handling two `CloudletsPolicyActivation` resources in the same configuration file with the same `policyId`, but different `network` arguments (for example, `production` and `staging`), you need to add `dependsOn` to the production resource. See the example:
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
// 		stag, err := akamai.NewCloudletsPolicyActivation(ctx, "stag", &akamai.CloudletsPolicyActivationArgs{
// 			PolicyId: pulumi.Int(1234567),
// 			Network:  pulumi.String("staging"),
// 			Version:  pulumi.Int(1),
// 			AssociatedProperties: pulumi.StringArray{
// 				pulumi.String("Property_1"),
// 				pulumi.String("Property_2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewCloudletsPolicyActivation(ctx, "prod", &akamai.CloudletsPolicyActivationArgs{
// 			PolicyId: pulumi.Int(1234567),
// 			Network:  pulumi.String("production"),
// 			Version:  pulumi.Int(1),
// 			AssociatedProperties: pulumi.StringArray{
// 				pulumi.String("Property_1"),
// 				pulumi.String("Property_2"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			stag,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CloudletsPolicyActivation struct {
	pulumi.CustomResourceState

	// A set of property identifiers related to this Cloudlet policy. You can't activate a Cloudlet policy if it doesn't have any properties associated with it.
	AssociatedProperties pulumi.StringArrayOutput `pulumi:"associatedProperties"`
	// The network you want to activate the policy version on. For the Staging network, specify either `staging`, `stag`, or `s`. For the Production network, specify either `production`, `prod`, or `p`. All values are case insensitive.
	Network pulumi.StringOutput `pulumi:"network"`
	// An identifier for the Cloudlet policy you want to activate.
	PolicyId pulumi.IntOutput `pulumi:"policyId"`
	// The activation status for this Cloudlet policy.
	Status pulumi.StringOutput `pulumi:"status"`
	// The Cloudlet policy version you want to activate.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewCloudletsPolicyActivation registers a new resource with the given unique name, arguments, and options.
func NewCloudletsPolicyActivation(ctx *pulumi.Context,
	name string, args *CloudletsPolicyActivationArgs, opts ...pulumi.ResourceOption) (*CloudletsPolicyActivation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssociatedProperties == nil {
		return nil, errors.New("invalid value for required argument 'AssociatedProperties'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource CloudletsPolicyActivation
	err := ctx.RegisterResource("akamai:index/cloudletsPolicyActivation:CloudletsPolicyActivation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudletsPolicyActivation gets an existing CloudletsPolicyActivation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudletsPolicyActivation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudletsPolicyActivationState, opts ...pulumi.ResourceOption) (*CloudletsPolicyActivation, error) {
	var resource CloudletsPolicyActivation
	err := ctx.ReadResource("akamai:index/cloudletsPolicyActivation:CloudletsPolicyActivation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudletsPolicyActivation resources.
type cloudletsPolicyActivationState struct {
	// A set of property identifiers related to this Cloudlet policy. You can't activate a Cloudlet policy if it doesn't have any properties associated with it.
	AssociatedProperties []string `pulumi:"associatedProperties"`
	// The network you want to activate the policy version on. For the Staging network, specify either `staging`, `stag`, or `s`. For the Production network, specify either `production`, `prod`, or `p`. All values are case insensitive.
	Network *string `pulumi:"network"`
	// An identifier for the Cloudlet policy you want to activate.
	PolicyId *int `pulumi:"policyId"`
	// The activation status for this Cloudlet policy.
	Status *string `pulumi:"status"`
	// The Cloudlet policy version you want to activate.
	Version *int `pulumi:"version"`
}

type CloudletsPolicyActivationState struct {
	// A set of property identifiers related to this Cloudlet policy. You can't activate a Cloudlet policy if it doesn't have any properties associated with it.
	AssociatedProperties pulumi.StringArrayInput
	// The network you want to activate the policy version on. For the Staging network, specify either `staging`, `stag`, or `s`. For the Production network, specify either `production`, `prod`, or `p`. All values are case insensitive.
	Network pulumi.StringPtrInput
	// An identifier for the Cloudlet policy you want to activate.
	PolicyId pulumi.IntPtrInput
	// The activation status for this Cloudlet policy.
	Status pulumi.StringPtrInput
	// The Cloudlet policy version you want to activate.
	Version pulumi.IntPtrInput
}

func (CloudletsPolicyActivationState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudletsPolicyActivationState)(nil)).Elem()
}

type cloudletsPolicyActivationArgs struct {
	// A set of property identifiers related to this Cloudlet policy. You can't activate a Cloudlet policy if it doesn't have any properties associated with it.
	AssociatedProperties []string `pulumi:"associatedProperties"`
	// The network you want to activate the policy version on. For the Staging network, specify either `staging`, `stag`, or `s`. For the Production network, specify either `production`, `prod`, or `p`. All values are case insensitive.
	Network string `pulumi:"network"`
	// An identifier for the Cloudlet policy you want to activate.
	PolicyId int `pulumi:"policyId"`
	// The Cloudlet policy version you want to activate.
	Version int `pulumi:"version"`
}

// The set of arguments for constructing a CloudletsPolicyActivation resource.
type CloudletsPolicyActivationArgs struct {
	// A set of property identifiers related to this Cloudlet policy. You can't activate a Cloudlet policy if it doesn't have any properties associated with it.
	AssociatedProperties pulumi.StringArrayInput
	// The network you want to activate the policy version on. For the Staging network, specify either `staging`, `stag`, or `s`. For the Production network, specify either `production`, `prod`, or `p`. All values are case insensitive.
	Network pulumi.StringInput
	// An identifier for the Cloudlet policy you want to activate.
	PolicyId pulumi.IntInput
	// The Cloudlet policy version you want to activate.
	Version pulumi.IntInput
}

func (CloudletsPolicyActivationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudletsPolicyActivationArgs)(nil)).Elem()
}

type CloudletsPolicyActivationInput interface {
	pulumi.Input

	ToCloudletsPolicyActivationOutput() CloudletsPolicyActivationOutput
	ToCloudletsPolicyActivationOutputWithContext(ctx context.Context) CloudletsPolicyActivationOutput
}

func (*CloudletsPolicyActivation) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudletsPolicyActivation)(nil)).Elem()
}

func (i *CloudletsPolicyActivation) ToCloudletsPolicyActivationOutput() CloudletsPolicyActivationOutput {
	return i.ToCloudletsPolicyActivationOutputWithContext(context.Background())
}

func (i *CloudletsPolicyActivation) ToCloudletsPolicyActivationOutputWithContext(ctx context.Context) CloudletsPolicyActivationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsPolicyActivationOutput)
}

// CloudletsPolicyActivationArrayInput is an input type that accepts CloudletsPolicyActivationArray and CloudletsPolicyActivationArrayOutput values.
// You can construct a concrete instance of `CloudletsPolicyActivationArrayInput` via:
//
//          CloudletsPolicyActivationArray{ CloudletsPolicyActivationArgs{...} }
type CloudletsPolicyActivationArrayInput interface {
	pulumi.Input

	ToCloudletsPolicyActivationArrayOutput() CloudletsPolicyActivationArrayOutput
	ToCloudletsPolicyActivationArrayOutputWithContext(context.Context) CloudletsPolicyActivationArrayOutput
}

type CloudletsPolicyActivationArray []CloudletsPolicyActivationInput

func (CloudletsPolicyActivationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudletsPolicyActivation)(nil)).Elem()
}

func (i CloudletsPolicyActivationArray) ToCloudletsPolicyActivationArrayOutput() CloudletsPolicyActivationArrayOutput {
	return i.ToCloudletsPolicyActivationArrayOutputWithContext(context.Background())
}

func (i CloudletsPolicyActivationArray) ToCloudletsPolicyActivationArrayOutputWithContext(ctx context.Context) CloudletsPolicyActivationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsPolicyActivationArrayOutput)
}

// CloudletsPolicyActivationMapInput is an input type that accepts CloudletsPolicyActivationMap and CloudletsPolicyActivationMapOutput values.
// You can construct a concrete instance of `CloudletsPolicyActivationMapInput` via:
//
//          CloudletsPolicyActivationMap{ "key": CloudletsPolicyActivationArgs{...} }
type CloudletsPolicyActivationMapInput interface {
	pulumi.Input

	ToCloudletsPolicyActivationMapOutput() CloudletsPolicyActivationMapOutput
	ToCloudletsPolicyActivationMapOutputWithContext(context.Context) CloudletsPolicyActivationMapOutput
}

type CloudletsPolicyActivationMap map[string]CloudletsPolicyActivationInput

func (CloudletsPolicyActivationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudletsPolicyActivation)(nil)).Elem()
}

func (i CloudletsPolicyActivationMap) ToCloudletsPolicyActivationMapOutput() CloudletsPolicyActivationMapOutput {
	return i.ToCloudletsPolicyActivationMapOutputWithContext(context.Background())
}

func (i CloudletsPolicyActivationMap) ToCloudletsPolicyActivationMapOutputWithContext(ctx context.Context) CloudletsPolicyActivationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsPolicyActivationMapOutput)
}

type CloudletsPolicyActivationOutput struct{ *pulumi.OutputState }

func (CloudletsPolicyActivationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudletsPolicyActivation)(nil)).Elem()
}

func (o CloudletsPolicyActivationOutput) ToCloudletsPolicyActivationOutput() CloudletsPolicyActivationOutput {
	return o
}

func (o CloudletsPolicyActivationOutput) ToCloudletsPolicyActivationOutputWithContext(ctx context.Context) CloudletsPolicyActivationOutput {
	return o
}

type CloudletsPolicyActivationArrayOutput struct{ *pulumi.OutputState }

func (CloudletsPolicyActivationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudletsPolicyActivation)(nil)).Elem()
}

func (o CloudletsPolicyActivationArrayOutput) ToCloudletsPolicyActivationArrayOutput() CloudletsPolicyActivationArrayOutput {
	return o
}

func (o CloudletsPolicyActivationArrayOutput) ToCloudletsPolicyActivationArrayOutputWithContext(ctx context.Context) CloudletsPolicyActivationArrayOutput {
	return o
}

func (o CloudletsPolicyActivationArrayOutput) Index(i pulumi.IntInput) CloudletsPolicyActivationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudletsPolicyActivation {
		return vs[0].([]*CloudletsPolicyActivation)[vs[1].(int)]
	}).(CloudletsPolicyActivationOutput)
}

type CloudletsPolicyActivationMapOutput struct{ *pulumi.OutputState }

func (CloudletsPolicyActivationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudletsPolicyActivation)(nil)).Elem()
}

func (o CloudletsPolicyActivationMapOutput) ToCloudletsPolicyActivationMapOutput() CloudletsPolicyActivationMapOutput {
	return o
}

func (o CloudletsPolicyActivationMapOutput) ToCloudletsPolicyActivationMapOutputWithContext(ctx context.Context) CloudletsPolicyActivationMapOutput {
	return o
}

func (o CloudletsPolicyActivationMapOutput) MapIndex(k pulumi.StringInput) CloudletsPolicyActivationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudletsPolicyActivation {
		return vs[0].(map[string]*CloudletsPolicyActivation)[vs[1].(string)]
	}).(CloudletsPolicyActivationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsPolicyActivationInput)(nil)).Elem(), &CloudletsPolicyActivation{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsPolicyActivationArrayInput)(nil)).Elem(), CloudletsPolicyActivationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsPolicyActivationMapInput)(nil)).Elem(), CloudletsPolicyActivationMap{})
	pulumi.RegisterOutputType(CloudletsPolicyActivationOutput{})
	pulumi.RegisterOutputType(CloudletsPolicyActivationArrayOutput{})
	pulumi.RegisterOutputType(CloudletsPolicyActivationMapOutput{})
}
