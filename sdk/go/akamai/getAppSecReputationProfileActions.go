// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getAppSecReputationProfileActions` data source to retrieve details about reputation profiles and their associated actions, or about the actions associated with a specific reputation profile.
func GetAppSecReputationProfileActions(ctx *pulumi.Context, args *GetAppSecReputationProfileActionsArgs, opts ...pulumi.InvokeOption) (*GetAppSecReputationProfileActionsResult, error) {
	var rv GetAppSecReputationProfileActionsResult
	err := ctx.Invoke("akamai:index/getAppSecReputationProfileActions:getAppSecReputationProfileActions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecReputationProfileActions.
type GetAppSecReputationProfileActionsArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of a given reputation profile. If not supplied, information about all reputation profiles is returned.
	ReputationProfileId *int `pulumi:"reputationProfileId"`
	// THe ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecReputationProfileActions.
type GetAppSecReputationProfileActionsResult struct {
	// The action that the specified reputation profile or profiles take when triggered.
	Action   string `pulumi:"action"`
	ConfigId int    `pulumi:"configId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A JSON-formatted display of the specified reputation profile action information.
	Json string `pulumi:"json"`
	// A tabular display of the specified reputation profile action information.
	OutputText          string `pulumi:"outputText"`
	ReputationProfileId *int   `pulumi:"reputationProfileId"`
	SecurityPolicyId    string `pulumi:"securityPolicyId"`
}