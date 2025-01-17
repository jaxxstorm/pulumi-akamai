// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `CloudletsPolicy` data source to list details about a policy with and its specified version, or latest if not specified.
//
// ## Basic usage
//
// This example returns the policy details based on the policy ID and optionally, a version:
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
//			_, err := akamai.LookupCloudletsPolicy(ctx, &GetCloudletsPolicyArgs{
//				PolicyId: 1234,
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
// * `groupId` - Defines the group association for the policy. You must have edit privileges for the group.
// * `name` - The unique name of the policy.
// * `apiVersion` - The specific version of the Cloudlets API.
// * `cloudletId` - A unique identifier that corresponds to a Cloudlets policy type. Enter `0` for Edge Redirector, `1` for Visitor Prioritization, `3` for Forward Rewrite, `4` for Request Control, `5` for API Prioritization, `6` for Audience Segmentation, `7` for Phased Release, `8` for Input Validation, or `9` for Application Load Balancer.
// * `cloudletCode` - The two- or three- character code for the type of Cloudlet. Enter `ALB` for Application Load Balancer, `AP` for API Prioritization, `AS` for Audience Segmentation, `CD` for Phased Release, `ER` for Edge Redirector, `FR` for Forward Rewrite, `IG` for Request Control, `IV` for Input Validation, or `VP` for Visitor Prioritization.
// * `revisionId` - A unique identifier given to every policy version update.
// * `description` - The description of this specific policy.
// * `versionDescription` - The description of this specific policy version.
// * `rulesLocked` - Whether editing `matchRules` for the Cloudlet policy version is blocked.
// * `matchRules`- A JSON structure that defines the rules for this policy.
// * `matchRuleFormat` - The format of the Cloudlet-specific `matchRules`.
// * `warnings` - A JSON encoded list of warnings.
// * `activations` - A list of of current policy activation information, including:
//   - `apiVersion` - The specific version of the Cloudlets API.
//   - `network` - The network, either `staging` or `prod` on which a property or a Cloudlets policy has been activated.
//   - `policyInfo` - A list of Cloudlet policy information, including:
//   - `policyId` - An integer identifier that is associated with all versions of a policy.
//   - `name` - The name of the policy.
//   - `version` - The version number of the policy.
//   - `status` - The activation status for the policy. Values include the following: `inactive` where the policy version has not been activated. No active property versions reference this policy. `active` where the policy version is currently active (published) and its associated property version is also active. `deactivated` where the policy version was previously activated but it has been superseded by a more recent activation of another policy version. `pending` where the policy version is proceeding through the activation workflow. `failed` where the policy version activation workflow has failed.
//   - `statusDetail` - Information about the status of an activation operation. This field is not returned when it has no value.
//   - `activatedBy` - The name of the user who activated the policy.
//   - `activationDate` - The date on which the policy was activated in milliseconds since epoch.
//   - `propertyInfo` A list of Cloudlet property information, including:
//   - `name` - The name of the property.
//   - `version` - The version number of the activated property.
//   - `groupId` - Defines the group association for the policy or property. If returns `0`, the policy is not tied to a group and in effect appears in all groups for the account. You must have edit privileges for the group.
//   - `status` - The activation status for the property. Values include the following: `inactive` where the policy version has not been activated. No active property versions reference this policy. `active` where the policy version is currently active (published) and its associated property version is also active. `deactivated` where the policy version was previously activated but it has been superseded by a more recent activation of another policy version. `pending` where the policy version is proceeding through the activation workflow. `failed` where the policy version activation workflow has failed.
//   - `activatedBy` - The name of the user who activated the property.
//   - `activationDate` - The date on which the property was activated in milliseconds since epoch.
func LookupCloudletsPolicy(ctx *pulumi.Context, args *LookupCloudletsPolicyArgs, opts ...pulumi.InvokeOption) (*LookupCloudletsPolicyResult, error) {
	var rv LookupCloudletsPolicyResult
	err := ctx.Invoke("akamai:index/getCloudletsPolicy:getCloudletsPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudletsPolicy.
type LookupCloudletsPolicyArgs struct {
	// - (Required) An integer identifier that is associated with all versions of a policy.
	PolicyId int `pulumi:"policyId"`
	// - (Optional) The version number of a policy.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getCloudletsPolicy.
type LookupCloudletsPolicyResult struct {
	Activations  []GetCloudletsPolicyActivationType `pulumi:"activations"`
	ApiVersion   string                             `pulumi:"apiVersion"`
	CloudletCode string                             `pulumi:"cloudletCode"`
	CloudletId   int                                `pulumi:"cloudletId"`
	Description  string                             `pulumi:"description"`
	GroupId      int                                `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string `pulumi:"id"`
	MatchRuleFormat    string `pulumi:"matchRuleFormat"`
	MatchRules         string `pulumi:"matchRules"`
	Name               string `pulumi:"name"`
	PolicyId           int    `pulumi:"policyId"`
	RevisionId         int    `pulumi:"revisionId"`
	RulesLocked        bool   `pulumi:"rulesLocked"`
	Version            *int   `pulumi:"version"`
	VersionDescription string `pulumi:"versionDescription"`
	Warnings           string `pulumi:"warnings"`
}

func LookupCloudletsPolicyOutput(ctx *pulumi.Context, args LookupCloudletsPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupCloudletsPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudletsPolicyResult, error) {
			args := v.(LookupCloudletsPolicyArgs)
			r, err := LookupCloudletsPolicy(ctx, &args, opts...)
			var s LookupCloudletsPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudletsPolicyResultOutput)
}

// A collection of arguments for invoking getCloudletsPolicy.
type LookupCloudletsPolicyOutputArgs struct {
	// - (Required) An integer identifier that is associated with all versions of a policy.
	PolicyId pulumi.IntInput `pulumi:"policyId"`
	// - (Optional) The version number of a policy.
	Version pulumi.IntPtrInput `pulumi:"version"`
}

func (LookupCloudletsPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudletsPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getCloudletsPolicy.
type LookupCloudletsPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupCloudletsPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudletsPolicyResult)(nil)).Elem()
}

func (o LookupCloudletsPolicyResultOutput) ToLookupCloudletsPolicyResultOutput() LookupCloudletsPolicyResultOutput {
	return o
}

func (o LookupCloudletsPolicyResultOutput) ToLookupCloudletsPolicyResultOutputWithContext(ctx context.Context) LookupCloudletsPolicyResultOutput {
	return o
}

func (o LookupCloudletsPolicyResultOutput) Activations() GetCloudletsPolicyActivationTypeArrayOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) []GetCloudletsPolicyActivationType { return v.Activations }).(GetCloudletsPolicyActivationTypeArrayOutput)
}

func (o LookupCloudletsPolicyResultOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.ApiVersion }).(pulumi.StringOutput)
}

func (o LookupCloudletsPolicyResultOutput) CloudletCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.CloudletCode }).(pulumi.StringOutput)
}

func (o LookupCloudletsPolicyResultOutput) CloudletId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) int { return v.CloudletId }).(pulumi.IntOutput)
}

func (o LookupCloudletsPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupCloudletsPolicyResultOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) int { return v.GroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudletsPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudletsPolicyResultOutput) MatchRuleFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.MatchRuleFormat }).(pulumi.StringOutput)
}

func (o LookupCloudletsPolicyResultOutput) MatchRules() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.MatchRules }).(pulumi.StringOutput)
}

func (o LookupCloudletsPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudletsPolicyResultOutput) PolicyId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) int { return v.PolicyId }).(pulumi.IntOutput)
}

func (o LookupCloudletsPolicyResultOutput) RevisionId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) int { return v.RevisionId }).(pulumi.IntOutput)
}

func (o LookupCloudletsPolicyResultOutput) RulesLocked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) bool { return v.RulesLocked }).(pulumi.BoolOutput)
}

func (o LookupCloudletsPolicyResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func (o LookupCloudletsPolicyResultOutput) VersionDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.VersionDescription }).(pulumi.StringOutput)
}

func (o LookupCloudletsPolicyResultOutput) Warnings() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudletsPolicyResult) string { return v.Warnings }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudletsPolicyResultOutput{})
}
