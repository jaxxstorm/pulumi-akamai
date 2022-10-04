// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `CloudletsApplicationLoadBalancer` data source to list details about the Application Load Balancer configuration with a specified policy version, or latest if not specified.
//
// ## Basic usage
//
// This example returns the load balancing configuration details based on the origin ID and optionally, a version:
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
//			_, err := akamai.LookupCloudletsApplicationLoadBalancer(ctx, &GetCloudletsApplicationLoadBalancerArgs{
//				OriginId: "alb_test_1",
//				Version:  pulumi.IntRef(1),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `description` - The description of the load balancing configuration.
// * `type` - The type of Conditional Origin. `APPLICATION_LOAD_BALANCER` is the only supported value.
// * `balancingType` - The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
// * `createdBy` - The name of the user who created this load balancing configuration.
// * `createdDate` - The date, in ISO 8601 format, when this load balancing configuration was created.
// * `deleted` - Whether the Conditional Origin version has been deleted. If `false`, you can use this version again.
// * `immutable` - Whether you can edit the load balancing version. The default setting for this member is false. It automatically becomes true when the load balancing version is activated for the first time.
// * `lastModifiedBy` - The user who last modified the load balancing configuration.
// * `lastModifiedDate` - The date, in ISO 8601 format, when the initial load balancing configuration was last modified.
// * `warnings` - A list of warnings that occurred during the activation of the load balancing configuration.
// * `dataCenters` - Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
//   - `city` - The city in which the data center is located.
//   - `cloudServerHostHeaderOverride` - Whether the cloud server host header is overridden.
//   - `cloudService` - Whether this datacenter is a cloud service.
//   - `continent` - The code of the continent on which the data center is located. See [Continent Codes](https://control.akamai.com/dl/edgescape/continentCodes.csv) for a list of valid codes.
//   - `country` - The country in which the data center is located. See [Country Codes](https://control.akamai.com/dl/edgescape/cc2continent.csv) for a list of valid codes.
//   - `hostname` - The name of the host that can be used as a Conditional Origin. This should match the `hostname` value defined for this datacenter in Property Manager.
//   - `latitude` - The latitude value for the data center. This member supports six decimal places of precision.
//   - `livenessHosts` - A list of the origin servers used to poll the data centers in an Application Load Balancer configuration. These servers support basic HTTP polling.
//   - `longitude` - The longitude value for the data center. This member supports six decimal places of precision.
//   - `originId` - The ID of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
//   - `percent` - The percent of traffic that is sent to the data center. The total for all data centers must equal 100%.
//   - `stateOrProvince` - The state, province, or region where the data center is located.
//
// * `livenessSettings` - Specifies the health of each load balanced data center defined in the data center list.
//   - `hostHeader` - The Host header for the liveness HTTP request.
//   - `additionalHeaders` - Maps additional case-insensitive HTTP header names included to the liveness testing requests.
//   - `interval` - The frequency of liveness tests. Defaults to 60 seconds, minimum is 10 seconds.
//   - `path` - The path to the test object used for liveness testing. The function of the test object is to help determine whether the data center is functioning.
//   - `peerCertificateVerification` - Whether to validate the origin certificate for an HTTPS request.
//   - `port` - The port for the test object. The default port is 80, which is standard for HTTP. Enter 443 if you are using HTTPS.
//   - `protocol` - The protocol or scheme for the database, either `HTTP` or `HTTPS`.
//   - `requestString` - The request used for TCP and TCPS tests.
//   - `responseString` - The response used for TCP and TCPS tests.
//   - `status3xxFailure` - If `true`, marks the liveness test as failed when the request returns a 3xx (redirection) status code.
//   - `status4xxFailure` - If `true`, marks the liveness test as failed when the request returns a 4xx (client error) status code.
//   - `status5xxFailure` - If `true`, marks the liveness test as failed when the request returns a 5xx (server error) status code.
//   - `timeout` - The number of seconds the system waits before failing the liveness test.
func LookupCloudletsApplicationLoadBalancer(ctx *pulumi.Context, args *LookupCloudletsApplicationLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupCloudletsApplicationLoadBalancerResult, error) {
	var rv LookupCloudletsApplicationLoadBalancerResult
	err := ctx.Invoke("akamai:index/getCloudletsApplicationLoadBalancer:getCloudletsApplicationLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudletsApplicationLoadBalancer.
type LookupCloudletsApplicationLoadBalancerArgs struct {
	// - (Required) A unique identifier for the Conditional Origin that supports the load balancing configuration. The Conditional Origin type must be set to `APPLICATION_LOAD_BALANCER` in the `origin` behavior. See property rules for more information.
	OriginId string `pulumi:"originId"`
	// - (Optional) The version number of the load balancing configuration.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getCloudletsApplicationLoadBalancer.
type LookupCloudletsApplicationLoadBalancerResult struct {
	BalancingType string                                          `pulumi:"balancingType"`
	CreatedBy     string                                          `pulumi:"createdBy"`
	CreatedDate   string                                          `pulumi:"createdDate"`
	DataCenters   []GetCloudletsApplicationLoadBalancerDataCenter `pulumi:"dataCenters"`
	Deleted       bool                                            `pulumi:"deleted"`
	Description   string                                          `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                                               `pulumi:"id"`
	Immutable        bool                                                 `pulumi:"immutable"`
	LastModifiedBy   string                                               `pulumi:"lastModifiedBy"`
	LastModifiedDate string                                               `pulumi:"lastModifiedDate"`
	LivenessSettings []GetCloudletsApplicationLoadBalancerLivenessSetting `pulumi:"livenessSettings"`
	OriginId         string                                               `pulumi:"originId"`
	Type             string                                               `pulumi:"type"`
	Version          *int                                                 `pulumi:"version"`
	Warnings         string                                               `pulumi:"warnings"`
}

func LookupCloudletsApplicationLoadBalancerOutput(ctx *pulumi.Context, args LookupCloudletsApplicationLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupCloudletsApplicationLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudletsApplicationLoadBalancerResult, error) {
			args := v.(LookupCloudletsApplicationLoadBalancerArgs)
			r, err := LookupCloudletsApplicationLoadBalancer(ctx, &args, opts...)
			var s LookupCloudletsApplicationLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudletsApplicationLoadBalancerResultOutput)
}

// A collection of arguments for invoking getCloudletsApplicationLoadBalancer.
type LookupCloudletsApplicationLoadBalancerOutputArgs struct {
	// - (Required) A unique identifier for the Conditional Origin that supports the load balancing configuration. The Conditional Origin type must be set to `APPLICATION_LOAD_BALANCER` in the `origin` behavior. See property rules for more information.
	OriginId pulumi.StringInput `pulumi:"originId"`
	// - (Optional) The version number of the load balancing configuration.
	Version pulumi.IntPtrInput `pulumi:"version"`
}

func (LookupCloudletsApplicationLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudletsApplicationLoadBalancerArgs)(nil)).Elem()
}

// A collection of values returned by getCloudletsApplicationLoadBalancer.
type LookupCloudletsApplicationLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupCloudletsApplicationLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudletsApplicationLoadBalancerResult)(nil)).Elem()
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) ToLookupCloudletsApplicationLoadBalancerResultOutput() LookupCloudletsApplicationLoadBalancerResultOutput {
	return o
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) ToLookupCloudletsApplicationLoadBalancerResultOutputWithContext(ctx context.Context) LookupCloudletsApplicationLoadBalancerResultOutput {
	return o
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) BalancingType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.BalancingType }).(pulumi.StringOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) DataCenters() GetCloudletsApplicationLoadBalancerDataCenterArrayOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) []GetCloudletsApplicationLoadBalancerDataCenter {
		return v.DataCenters
	}).(GetCloudletsApplicationLoadBalancerDataCenterArrayOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) bool { return v.Deleted }).(pulumi.BoolOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudletsApplicationLoadBalancerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) Immutable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) bool { return v.Immutable }).(pulumi.BoolOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) LastModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.LastModifiedDate }).(pulumi.StringOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) LivenessSettings() GetCloudletsApplicationLoadBalancerLivenessSettingArrayOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) []GetCloudletsApplicationLoadBalancerLivenessSetting {
		return v.LivenessSettings
	}).(GetCloudletsApplicationLoadBalancerLivenessSettingArrayOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) OriginId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.OriginId }).(pulumi.StringOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func (o LookupCloudletsApplicationLoadBalancerResultOutput) Warnings() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsApplicationLoadBalancerResult) string { return v.Warnings }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudletsApplicationLoadBalancerResultOutput{})
}
