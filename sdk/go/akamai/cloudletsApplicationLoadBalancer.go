// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `CloudletsApplicationLoadBalancer` resource to create the Application Load Balancer Cloudlet configuration. The Application Load Balancer Cloudlet provides intelligent, scalable traffic management across physical, virtual, and cloud-hosted data centers without requiring the origin to send load feedback. This Cloudlet can automatically detect load conditions and route traffic to the optimal data source while maintaining custom routing policies and consistent visitor session behavior for your visitors.
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
// 		_, err := akamai.NewCloudletsApplicationLoadBalancer(ctx, "example", &akamai.CloudletsApplicationLoadBalancerArgs{
// 			BalancingType: pulumi.String("WEIGHTED"),
// 			DataCenters: CloudletsApplicationLoadBalancerDataCenterArray{
// 				&CloudletsApplicationLoadBalancerDataCenterArgs{
// 					City:                          pulumi.String("Boston"),
// 					CloudServerHostHeaderOverride: pulumi.Bool(false),
// 					CloudService:                  pulumi.Bool(true),
// 					Continent:                     pulumi.String("NA"),
// 					Country:                       pulumi.String("US"),
// 					Hostname:                      pulumi.String("example-hostname"),
// 					Latitude:                      pulumi.Float64(102.78108),
// 					LivenessHosts: pulumi.StringArray{
// 						pulumi.String("example"),
// 					},
// 					Longitude:       -116.07064,
// 					OriginId:        pulumi.String("alb_test_1"),
// 					Percent:         pulumi.Float64(100),
// 					StateOrProvince: pulumi.String("MA"),
// 				},
// 			},
// 			Description: pulumi.String("application_load_balancer description"),
// 			LivenessSettings: &CloudletsApplicationLoadBalancerLivenessSettingsArgs{
// 				AdditionalHeaders: pulumi.StringMap{
// 					"additionalHeaders": pulumi.String("123"),
// 				},
// 				HostHeader:     pulumi.String("header"),
// 				Interval:       pulumi.Int(10),
// 				Path:           pulumi.String("/status"),
// 				Port:           pulumi.Int(1234),
// 				Protocol:       pulumi.String("HTTP"),
// 				RequestString:  pulumi.String("test_request_string"),
// 				ResponseString: pulumi.String("test_response_string"),
// 				Timeout:        pulumi.Float64(60),
// 			},
// 			OriginId: pulumi.String("alb_test_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Basic usagehcl resource "akamai_cloudlets_application_load_balancer" "example" {
//
// # (resource arguments)
//
//  } You can import your Akamai Application Load Balancer configuration using an origin ID. For example
//
// ```sh
//  $ pulumi import akamai:index/cloudletsApplicationLoadBalancer:CloudletsApplicationLoadBalancer example alb_test_1
// ```
type CloudletsApplicationLoadBalancer struct {
	pulumi.CustomResourceState

	// The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
	BalancingType pulumi.StringPtrOutput `pulumi:"balancingType"`
	// Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
	DataCenters CloudletsApplicationLoadBalancerDataCenterArrayOutput `pulumi:"dataCenters"`
	// The description of the load balancing configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the health of each load balanced data center defined in the data center list.
	LivenessSettings CloudletsApplicationLoadBalancerLivenessSettingsPtrOutput `pulumi:"livenessSettings"`
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId pulumi.StringOutput `pulumi:"originId"`
	// The version number of the load balancing configuration.
	Version pulumi.IntOutput `pulumi:"version"`
	// A list of warnings that occurred during the activation of the load balancing configuration.
	Warnings pulumi.StringOutput `pulumi:"warnings"`
}

// NewCloudletsApplicationLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewCloudletsApplicationLoadBalancer(ctx *pulumi.Context,
	name string, args *CloudletsApplicationLoadBalancerArgs, opts ...pulumi.ResourceOption) (*CloudletsApplicationLoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataCenters == nil {
		return nil, errors.New("invalid value for required argument 'DataCenters'")
	}
	if args.OriginId == nil {
		return nil, errors.New("invalid value for required argument 'OriginId'")
	}
	var resource CloudletsApplicationLoadBalancer
	err := ctx.RegisterResource("akamai:index/cloudletsApplicationLoadBalancer:CloudletsApplicationLoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudletsApplicationLoadBalancer gets an existing CloudletsApplicationLoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudletsApplicationLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudletsApplicationLoadBalancerState, opts ...pulumi.ResourceOption) (*CloudletsApplicationLoadBalancer, error) {
	var resource CloudletsApplicationLoadBalancer
	err := ctx.ReadResource("akamai:index/cloudletsApplicationLoadBalancer:CloudletsApplicationLoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudletsApplicationLoadBalancer resources.
type cloudletsApplicationLoadBalancerState struct {
	// The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
	BalancingType *string `pulumi:"balancingType"`
	// Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
	DataCenters []CloudletsApplicationLoadBalancerDataCenter `pulumi:"dataCenters"`
	// The description of the load balancing configuration.
	Description *string `pulumi:"description"`
	// Specifies the health of each load balanced data center defined in the data center list.
	LivenessSettings *CloudletsApplicationLoadBalancerLivenessSettings `pulumi:"livenessSettings"`
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId *string `pulumi:"originId"`
	// The version number of the load balancing configuration.
	Version *int `pulumi:"version"`
	// A list of warnings that occurred during the activation of the load balancing configuration.
	Warnings *string `pulumi:"warnings"`
}

type CloudletsApplicationLoadBalancerState struct {
	// The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
	BalancingType pulumi.StringPtrInput
	// Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
	DataCenters CloudletsApplicationLoadBalancerDataCenterArrayInput
	// The description of the load balancing configuration.
	Description pulumi.StringPtrInput
	// Specifies the health of each load balanced data center defined in the data center list.
	LivenessSettings CloudletsApplicationLoadBalancerLivenessSettingsPtrInput
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId pulumi.StringPtrInput
	// The version number of the load balancing configuration.
	Version pulumi.IntPtrInput
	// A list of warnings that occurred during the activation of the load balancing configuration.
	Warnings pulumi.StringPtrInput
}

func (CloudletsApplicationLoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudletsApplicationLoadBalancerState)(nil)).Elem()
}

type cloudletsApplicationLoadBalancerArgs struct {
	// The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
	BalancingType *string `pulumi:"balancingType"`
	// Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
	DataCenters []CloudletsApplicationLoadBalancerDataCenter `pulumi:"dataCenters"`
	// The description of the load balancing configuration.
	Description *string `pulumi:"description"`
	// Specifies the health of each load balanced data center defined in the data center list.
	LivenessSettings *CloudletsApplicationLoadBalancerLivenessSettings `pulumi:"livenessSettings"`
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId string `pulumi:"originId"`
}

// The set of arguments for constructing a CloudletsApplicationLoadBalancer resource.
type CloudletsApplicationLoadBalancerArgs struct {
	// The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
	BalancingType pulumi.StringPtrInput
	// Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
	DataCenters CloudletsApplicationLoadBalancerDataCenterArrayInput
	// The description of the load balancing configuration.
	Description pulumi.StringPtrInput
	// Specifies the health of each load balanced data center defined in the data center list.
	LivenessSettings CloudletsApplicationLoadBalancerLivenessSettingsPtrInput
	// The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
	OriginId pulumi.StringInput
}

func (CloudletsApplicationLoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudletsApplicationLoadBalancerArgs)(nil)).Elem()
}

type CloudletsApplicationLoadBalancerInput interface {
	pulumi.Input

	ToCloudletsApplicationLoadBalancerOutput() CloudletsApplicationLoadBalancerOutput
	ToCloudletsApplicationLoadBalancerOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerOutput
}

func (*CloudletsApplicationLoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudletsApplicationLoadBalancer)(nil)).Elem()
}

func (i *CloudletsApplicationLoadBalancer) ToCloudletsApplicationLoadBalancerOutput() CloudletsApplicationLoadBalancerOutput {
	return i.ToCloudletsApplicationLoadBalancerOutputWithContext(context.Background())
}

func (i *CloudletsApplicationLoadBalancer) ToCloudletsApplicationLoadBalancerOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsApplicationLoadBalancerOutput)
}

// CloudletsApplicationLoadBalancerArrayInput is an input type that accepts CloudletsApplicationLoadBalancerArray and CloudletsApplicationLoadBalancerArrayOutput values.
// You can construct a concrete instance of `CloudletsApplicationLoadBalancerArrayInput` via:
//
//          CloudletsApplicationLoadBalancerArray{ CloudletsApplicationLoadBalancerArgs{...} }
type CloudletsApplicationLoadBalancerArrayInput interface {
	pulumi.Input

	ToCloudletsApplicationLoadBalancerArrayOutput() CloudletsApplicationLoadBalancerArrayOutput
	ToCloudletsApplicationLoadBalancerArrayOutputWithContext(context.Context) CloudletsApplicationLoadBalancerArrayOutput
}

type CloudletsApplicationLoadBalancerArray []CloudletsApplicationLoadBalancerInput

func (CloudletsApplicationLoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudletsApplicationLoadBalancer)(nil)).Elem()
}

func (i CloudletsApplicationLoadBalancerArray) ToCloudletsApplicationLoadBalancerArrayOutput() CloudletsApplicationLoadBalancerArrayOutput {
	return i.ToCloudletsApplicationLoadBalancerArrayOutputWithContext(context.Background())
}

func (i CloudletsApplicationLoadBalancerArray) ToCloudletsApplicationLoadBalancerArrayOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsApplicationLoadBalancerArrayOutput)
}

// CloudletsApplicationLoadBalancerMapInput is an input type that accepts CloudletsApplicationLoadBalancerMap and CloudletsApplicationLoadBalancerMapOutput values.
// You can construct a concrete instance of `CloudletsApplicationLoadBalancerMapInput` via:
//
//          CloudletsApplicationLoadBalancerMap{ "key": CloudletsApplicationLoadBalancerArgs{...} }
type CloudletsApplicationLoadBalancerMapInput interface {
	pulumi.Input

	ToCloudletsApplicationLoadBalancerMapOutput() CloudletsApplicationLoadBalancerMapOutput
	ToCloudletsApplicationLoadBalancerMapOutputWithContext(context.Context) CloudletsApplicationLoadBalancerMapOutput
}

type CloudletsApplicationLoadBalancerMap map[string]CloudletsApplicationLoadBalancerInput

func (CloudletsApplicationLoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudletsApplicationLoadBalancer)(nil)).Elem()
}

func (i CloudletsApplicationLoadBalancerMap) ToCloudletsApplicationLoadBalancerMapOutput() CloudletsApplicationLoadBalancerMapOutput {
	return i.ToCloudletsApplicationLoadBalancerMapOutputWithContext(context.Background())
}

func (i CloudletsApplicationLoadBalancerMap) ToCloudletsApplicationLoadBalancerMapOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudletsApplicationLoadBalancerMapOutput)
}

type CloudletsApplicationLoadBalancerOutput struct{ *pulumi.OutputState }

func (CloudletsApplicationLoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudletsApplicationLoadBalancer)(nil)).Elem()
}

func (o CloudletsApplicationLoadBalancerOutput) ToCloudletsApplicationLoadBalancerOutput() CloudletsApplicationLoadBalancerOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerOutput) ToCloudletsApplicationLoadBalancerOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerOutput {
	return o
}

type CloudletsApplicationLoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (CloudletsApplicationLoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudletsApplicationLoadBalancer)(nil)).Elem()
}

func (o CloudletsApplicationLoadBalancerArrayOutput) ToCloudletsApplicationLoadBalancerArrayOutput() CloudletsApplicationLoadBalancerArrayOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerArrayOutput) ToCloudletsApplicationLoadBalancerArrayOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerArrayOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerArrayOutput) Index(i pulumi.IntInput) CloudletsApplicationLoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudletsApplicationLoadBalancer {
		return vs[0].([]*CloudletsApplicationLoadBalancer)[vs[1].(int)]
	}).(CloudletsApplicationLoadBalancerOutput)
}

type CloudletsApplicationLoadBalancerMapOutput struct{ *pulumi.OutputState }

func (CloudletsApplicationLoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudletsApplicationLoadBalancer)(nil)).Elem()
}

func (o CloudletsApplicationLoadBalancerMapOutput) ToCloudletsApplicationLoadBalancerMapOutput() CloudletsApplicationLoadBalancerMapOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerMapOutput) ToCloudletsApplicationLoadBalancerMapOutputWithContext(ctx context.Context) CloudletsApplicationLoadBalancerMapOutput {
	return o
}

func (o CloudletsApplicationLoadBalancerMapOutput) MapIndex(k pulumi.StringInput) CloudletsApplicationLoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudletsApplicationLoadBalancer {
		return vs[0].(map[string]*CloudletsApplicationLoadBalancer)[vs[1].(string)]
	}).(CloudletsApplicationLoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsApplicationLoadBalancerInput)(nil)).Elem(), &CloudletsApplicationLoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsApplicationLoadBalancerArrayInput)(nil)).Elem(), CloudletsApplicationLoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudletsApplicationLoadBalancerMapInput)(nil)).Elem(), CloudletsApplicationLoadBalancerMap{})
	pulumi.RegisterOutputType(CloudletsApplicationLoadBalancerOutput{})
	pulumi.RegisterOutputType(CloudletsApplicationLoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(CloudletsApplicationLoadBalancerMapOutput{})
}
