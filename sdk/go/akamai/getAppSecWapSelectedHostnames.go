// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecWapSelectedHostnames` data source to retrieve lists of the hostnames that are currently
// protected and currently being evaluated under a given security configuration and policy. This resource is available
// only for WAP accounts. (WAP selected hostnames is currently in beta. Please contact your Akamai representative for
// more information about this feature.)
func LookupAppSecWapSelectedHostnames(ctx *pulumi.Context, args *LookupAppSecWapSelectedHostnamesArgs, opts ...pulumi.InvokeOption) (*LookupAppSecWapSelectedHostnamesResult, error) {
	var rv LookupAppSecWapSelectedHostnamesResult
	err := ctx.Invoke("akamai:index/getAppSecWapSelectedHostnames:getAppSecWapSelectedHostnames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppSecWapSelectedHostnames.
type LookupAppSecWapSelectedHostnamesArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// A collection of values returned by getAppSecWapSelectedHostnames.
type LookupAppSecWapSelectedHostnamesResult struct {
	ConfigId       int      `pulumi:"configId"`
	EvaluatedHosts []string `pulumi:"evaluatedHosts"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	Json         string `pulumi:"json"`
	MatchTargets string `pulumi:"matchTargets"`
	// A tabular display of the protected and evaluated hostnames.
	OutputText       string   `pulumi:"outputText"`
	ProtectedHosts   []string `pulumi:"protectedHosts"`
	SecurityPolicyId string   `pulumi:"securityPolicyId"`
	SelectedHosts    []string `pulumi:"selectedHosts"`
}
