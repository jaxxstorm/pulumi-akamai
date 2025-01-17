// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecSecurityPolicy
    {
        public static Task<GetAppSecSecurityPolicyResult> InvokeAsync(GetAppSecSecurityPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecSecurityPolicyResult>("akamai:index/getAppSecSecurityPolicy:getAppSecSecurityPolicy", args ?? new GetAppSecSecurityPolicyArgs(), options.WithDefaults());

        public static Output<GetAppSecSecurityPolicyResult> Invoke(GetAppSecSecurityPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecSecurityPolicyResult>("akamai:index/getAppSecSecurityPolicy:getAppSecSecurityPolicy", args ?? new GetAppSecSecurityPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecSecurityPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the security policies.
        /// - `security_policy_name`. (Optional). Name of the security policy you want to return information for (be sure to reference the policy name and not the policy ID). If not included, information is returned for all your security policies.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        [Input("securityPolicyName")]
        public string? SecurityPolicyName { get; set; }

        public GetAppSecSecurityPolicyArgs()
        {
        }
    }

    public sealed class GetAppSecSecurityPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the security policies.
        /// - `security_policy_name`. (Optional). Name of the security policy you want to return information for (be sure to reference the policy name and not the policy ID). If not included, information is returned for all your security policies.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("securityPolicyName")]
        public Input<string>? SecurityPolicyName { get; set; }

        public GetAppSecSecurityPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecSecurityPolicyResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OutputText;
        public readonly string SecurityPolicyId;
        public readonly ImmutableArray<string> SecurityPolicyIdLists;
        public readonly string? SecurityPolicyName;

        [OutputConstructor]
        private GetAppSecSecurityPolicyResult(
            int configId,

            string id,

            string outputText,

            string securityPolicyId,

            ImmutableArray<string> securityPolicyIdLists,

            string? securityPolicyName)
        {
            ConfigId = configId;
            Id = id;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
            SecurityPolicyIdLists = securityPolicyIdLists;
            SecurityPolicyName = securityPolicyName;
        }
    }
}
