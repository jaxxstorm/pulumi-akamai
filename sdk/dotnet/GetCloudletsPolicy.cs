// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetCloudletsPolicy
    {
        /// <summary>
        /// Use the `akamai.CloudletsPolicy` data source to list details about a policy with and its specified version, or latest if not specified.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the policy details based on the policy ID and optionally, a version:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetCloudletsPolicy.InvokeAsync(new Akamai.GetCloudletsPolicyArgs
        ///         {
        ///             PolicyId = 1234,
        ///             Version = 1,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `group_id` - Defines the group association for the policy. You must have edit privileges for the group.
        /// * `name` - The unique name of the policy.
        /// * `api_version` - The specific version of the Cloudlets API.
        /// * `cloudlet_id` - A unique identifier that corresponds to a Cloudlets policy type. Enter `0` for Edge Redirector, `1` for Visitor Prioritization, `3` for Forward Rewrite, `5` for API Prioritization, `7` for Phased Release, or `9` for Application Load Balancer.
        /// * `cloudlet_code` - The two- or three- character code for the type of Cloudlet, `ALB` for Application Load Balancer, `AP` for API Prioritization, `CD` for Phased Release, `ER` for Edge Redirector, `FR` for Forward Rewrite, and `VP` for Visitor Prioritization.
        /// * `revision_id` - A unique identifier given to every policy version update.
        /// * `description` - The description of this specific policy.
        /// * `version_description` - The description of this specific policy version.
        /// * `rules_locked` - Whether editing `match_rules` for the Cloudlet policy version is blocked.
        /// * `match_rules`- A JSON structure that defines the rules for this policy.
        /// * `match_rule_format` - The format of the Cloudlet-specific `match_rules`.
        /// * `warnings` - A JSON encoded list of warnings.
        /// * `activations` - A list of of current policy activation information, including:
        ///   * `api_version` - The specific version of the Cloudlets API.
        ///   * `network` - The network, either `staging` or `prod` on which a property or a Cloudlets policy has been activated.
        ///   * `policy_info` - A list of Cloudlet policy information, including:
        ///       * `policy_id` - An integer identifier that is associated with all versions of a policy.
        ///       * `name` - The name of the policy.
        ///       * `version` - The version number of the policy.
        ///       * `status` - The activation status for the policy. Values include the following: `inactive` where the policy version has not been activated. No active property versions reference this policy. `active` where the policy version is currently active (published) and its associated property version is also active. `deactivated` where the policy version was previously activated but it has been superseded by a more recent activation of another policy version. `pending` where the policy version is proceeding through the activation workflow. `failed` where the policy version activation workflow has failed.
        ///       * `status_detail` - Information about the status of an activation operation. This field is not returned when it has no value.
        ///       * `activated_by` - The name of the user who activated the policy.
        ///       * `activation_date` - The date on which the policy was activated in milliseconds since epoch.
        ///   * `property_info` A list of Cloudlet property information, including:
        ///       * `name` - The name of the property.
        ///       * `version` - The version number of the activated property.
        ///       * `group_id` - Defines the group association for the policy or property. If returns `0`, the policy is not tied to a group and in effect appears in all groups for the account. You must have edit privileges for the group.
        ///       * `status` - The activation status for the property. Values include the following: `inactive` where the policy version has not been activated. No active property versions reference this policy. `active` where the policy version is currently active (published) and its associated property version is also active. `deactivated` where the policy version was previously activated but it has been superseded by a more recent activation of another policy version. `pending` where the policy version is proceeding through the activation workflow. `failed` where the policy version activation workflow has failed.
        ///       * `activated_by` - The name of the user who activated the property.
        ///       * `activation_date` - The date on which the property was activated in milliseconds since epoch.
        /// </summary>
        public static Task<GetCloudletsPolicyResult> InvokeAsync(GetCloudletsPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudletsPolicyResult>("akamai:index/getCloudletsPolicy:getCloudletsPolicy", args ?? new GetCloudletsPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Use the `akamai.CloudletsPolicy` data source to list details about a policy with and its specified version, or latest if not specified.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the policy details based on the policy ID and optionally, a version:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetCloudletsPolicy.InvokeAsync(new Akamai.GetCloudletsPolicyArgs
        ///         {
        ///             PolicyId = 1234,
        ///             Version = 1,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `group_id` - Defines the group association for the policy. You must have edit privileges for the group.
        /// * `name` - The unique name of the policy.
        /// * `api_version` - The specific version of the Cloudlets API.
        /// * `cloudlet_id` - A unique identifier that corresponds to a Cloudlets policy type. Enter `0` for Edge Redirector, `1` for Visitor Prioritization, `3` for Forward Rewrite, `5` for API Prioritization, `7` for Phased Release, or `9` for Application Load Balancer.
        /// * `cloudlet_code` - The two- or three- character code for the type of Cloudlet, `ALB` for Application Load Balancer, `AP` for API Prioritization, `CD` for Phased Release, `ER` for Edge Redirector, `FR` for Forward Rewrite, and `VP` for Visitor Prioritization.
        /// * `revision_id` - A unique identifier given to every policy version update.
        /// * `description` - The description of this specific policy.
        /// * `version_description` - The description of this specific policy version.
        /// * `rules_locked` - Whether editing `match_rules` for the Cloudlet policy version is blocked.
        /// * `match_rules`- A JSON structure that defines the rules for this policy.
        /// * `match_rule_format` - The format of the Cloudlet-specific `match_rules`.
        /// * `warnings` - A JSON encoded list of warnings.
        /// * `activations` - A list of of current policy activation information, including:
        ///   * `api_version` - The specific version of the Cloudlets API.
        ///   * `network` - The network, either `staging` or `prod` on which a property or a Cloudlets policy has been activated.
        ///   * `policy_info` - A list of Cloudlet policy information, including:
        ///       * `policy_id` - An integer identifier that is associated with all versions of a policy.
        ///       * `name` - The name of the policy.
        ///       * `version` - The version number of the policy.
        ///       * `status` - The activation status for the policy. Values include the following: `inactive` where the policy version has not been activated. No active property versions reference this policy. `active` where the policy version is currently active (published) and its associated property version is also active. `deactivated` where the policy version was previously activated but it has been superseded by a more recent activation of another policy version. `pending` where the policy version is proceeding through the activation workflow. `failed` where the policy version activation workflow has failed.
        ///       * `status_detail` - Information about the status of an activation operation. This field is not returned when it has no value.
        ///       * `activated_by` - The name of the user who activated the policy.
        ///       * `activation_date` - The date on which the policy was activated in milliseconds since epoch.
        ///   * `property_info` A list of Cloudlet property information, including:
        ///       * `name` - The name of the property.
        ///       * `version` - The version number of the activated property.
        ///       * `group_id` - Defines the group association for the policy or property. If returns `0`, the policy is not tied to a group and in effect appears in all groups for the account. You must have edit privileges for the group.
        ///       * `status` - The activation status for the property. Values include the following: `inactive` where the policy version has not been activated. No active property versions reference this policy. `active` where the policy version is currently active (published) and its associated property version is also active. `deactivated` where the policy version was previously activated but it has been superseded by a more recent activation of another policy version. `pending` where the policy version is proceeding through the activation workflow. `failed` where the policy version activation workflow has failed.
        ///       * `activated_by` - The name of the user who activated the property.
        ///       * `activation_date` - The date on which the property was activated in milliseconds since epoch.
        /// </summary>
        public static Output<GetCloudletsPolicyResult> Invoke(GetCloudletsPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudletsPolicyResult>("akamai:index/getCloudletsPolicy:getCloudletsPolicy", args ?? new GetCloudletsPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudletsPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (Required) An integer identifier that is associated with all versions of a policy.
        /// </summary>
        [Input("policyId", required: true)]
        public int PolicyId { get; set; }

        /// <summary>
        /// - (Optional) The version number of a policy.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetCloudletsPolicyArgs()
        {
        }
    }

    public sealed class GetCloudletsPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (Required) An integer identifier that is associated with all versions of a policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<int> PolicyId { get; set; } = null!;

        /// <summary>
        /// - (Optional) The version number of a policy.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public GetCloudletsPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudletsPolicyResult
    {
        public readonly ImmutableArray<Outputs.GetCloudletsPolicyActivationResult> Activations;
        public readonly string ApiVersion;
        public readonly string CloudletCode;
        public readonly int CloudletId;
        public readonly string Description;
        public readonly int GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MatchRuleFormat;
        public readonly string MatchRules;
        public readonly string Name;
        public readonly int PolicyId;
        public readonly int RevisionId;
        public readonly bool RulesLocked;
        public readonly int? Version;
        public readonly string VersionDescription;
        public readonly string Warnings;

        [OutputConstructor]
        private GetCloudletsPolicyResult(
            ImmutableArray<Outputs.GetCloudletsPolicyActivationResult> activations,

            string apiVersion,

            string cloudletCode,

            int cloudletId,

            string description,

            int groupId,

            string id,

            string matchRuleFormat,

            string matchRules,

            string name,

            int policyId,

            int revisionId,

            bool rulesLocked,

            int? version,

            string versionDescription,

            string warnings)
        {
            Activations = activations;
            ApiVersion = apiVersion;
            CloudletCode = cloudletCode;
            CloudletId = cloudletId;
            Description = description;
            GroupId = groupId;
            Id = id;
            MatchRuleFormat = matchRuleFormat;
            MatchRules = matchRules;
            Name = name;
            PolicyId = policyId;
            RevisionId = revisionId;
            RulesLocked = rulesLocked;
            Version = version;
            VersionDescription = versionDescription;
            Warnings = warnings;
        }
    }
}
