// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecTuningRecommendations
    {
        /// <summary>
        /// Returns tuning recommendations for the specified attack group or rule (or, if both the `attack_group` and the `rule_id` arguments are not included, returns tuning recommendations for all the attack groups and rules in the specified security policy).
        /// Tuning recommendations help minimize the number of false positives triggered by a security policy. With a false positive, a client request is marked as having violated the security policy restrictions even though it actually did not.
        /// Tuning recommendations are returned as attack group or rule exceptions: if you choose, you can copy the response and use the `akamai.AppSecAttackGroup` resource to add the recommended exception to an attack group or the `akamai.AppSecRule` resource to add the recommended exception to a rule.  
        /// If the data source response is empty, that means that there are no further recommendations for tuning your security policy or attack group.
        /// If you need, you can manually merge a recommended exception for an attack group or a rule with the exception previously configured.
        /// You can find additional information in our [Application Security API v1 documentation](https://techdocs.akamai.com/application-security/reference/get-recommendations).
        /// 
        /// **Related API endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/recommendation](https://techdocs.akamai.com/application-security/reference/get-recommendations)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic usage:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = @var.Security_configuration,
        ///         }));
        ///         var policyRecommendations = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecTuningRecommendations.InvokeAsync(new Akamai.GetAppSecTuningRecommendationsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///         })));
        ///         this.PolicyRecommendationsJson = policyRecommendations.Apply(policyRecommendations =&gt; policyRecommendations.Json);
        ///         var attackGroupRecommendations = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecTuningRecommendations.InvokeAsync(new Akamai.GetAppSecTuningRecommendationsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///             RulesetType = @var.Ruleset_type,
        ///             AttackGroup = @var.Attack_group,
        ///         })));
        ///         this.AttackGroupRecommendationsJson = attackGroupRecommendations.Apply(attackGroupRecommendations =&gt; attackGroupRecommendations.Json);
        ///     }
        /// 
        ///     [Output("policyRecommendationsJson")]
        ///     public Output&lt;string&gt; PolicyRecommendationsJson { get; set; }
        ///     [Output("attackGroupRecommendationsJson")]
        ///     public Output&lt;string&gt; AttackGroupRecommendationsJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecTuningRecommendationsResult> InvokeAsync(GetAppSecTuningRecommendationsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecTuningRecommendationsResult>("akamai:index/getAppSecTuningRecommendations:getAppSecTuningRecommendations", args ?? new GetAppSecTuningRecommendationsArgs(), options.WithDefaults());

        /// <summary>
        /// Returns tuning recommendations for the specified attack group or rule (or, if both the `attack_group` and the `rule_id` arguments are not included, returns tuning recommendations for all the attack groups and rules in the specified security policy).
        /// Tuning recommendations help minimize the number of false positives triggered by a security policy. With a false positive, a client request is marked as having violated the security policy restrictions even though it actually did not.
        /// Tuning recommendations are returned as attack group or rule exceptions: if you choose, you can copy the response and use the `akamai.AppSecAttackGroup` resource to add the recommended exception to an attack group or the `akamai.AppSecRule` resource to add the recommended exception to a rule.  
        /// If the data source response is empty, that means that there are no further recommendations for tuning your security policy or attack group.
        /// If you need, you can manually merge a recommended exception for an attack group or a rule with the exception previously configured.
        /// You can find additional information in our [Application Security API v1 documentation](https://techdocs.akamai.com/application-security/reference/get-recommendations).
        /// 
        /// **Related API endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/recommendation](https://techdocs.akamai.com/application-security/reference/get-recommendations)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic usage:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = @var.Security_configuration,
        ///         }));
        ///         var policyRecommendations = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecTuningRecommendations.InvokeAsync(new Akamai.GetAppSecTuningRecommendationsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///         })));
        ///         this.PolicyRecommendationsJson = policyRecommendations.Apply(policyRecommendations =&gt; policyRecommendations.Json);
        ///         var attackGroupRecommendations = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecTuningRecommendations.InvokeAsync(new Akamai.GetAppSecTuningRecommendationsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///             RulesetType = @var.Ruleset_type,
        ///             AttackGroup = @var.Attack_group,
        ///         })));
        ///         this.AttackGroupRecommendationsJson = attackGroupRecommendations.Apply(attackGroupRecommendations =&gt; attackGroupRecommendations.Json);
        ///     }
        /// 
        ///     [Output("policyRecommendationsJson")]
        ///     public Output&lt;string&gt; PolicyRecommendationsJson { get; set; }
        ///     [Output("attackGroupRecommendationsJson")]
        ///     public Output&lt;string&gt; AttackGroupRecommendationsJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAppSecTuningRecommendationsResult> Invoke(GetAppSecTuningRecommendationsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecTuningRecommendationsResult>("akamai:index/getAppSecTuningRecommendations:getAppSecTuningRecommendations", args ?? new GetAppSecTuningRecommendationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecTuningRecommendationsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique name of the attack group you want tuning recommendations for. If both `attack_group` and `rule_id` not included, recommendations are returned for all attack groups.
        /// </summary>
        [Input("attackGroup")]
        public string? AttackGroup { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration you want tuning recommendations for.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Unique id of the rule you want tuning recommendations for. If both `attack_group` and `rule_id` not included, recommendations are returned for all attack groups.
        /// </summary>
        [Input("ruleId")]
        public int? RuleId { get; set; }

        /// <summary>
        /// . Type of ruleset used by the security configuration you want tuning recommendations for. Supported values are `active` and `evaluation`. Defaults to `active`.
        /// </summary>
        [Input("rulesetType")]
        public string? RulesetType { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy you want tuning recommendations for.
        /// </summary>
        [Input("securityPolicyId")]
        public string? SecurityPolicyId { get; set; }

        public GetAppSecTuningRecommendationsArgs()
        {
        }
    }

    public sealed class GetAppSecTuningRecommendationsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique name of the attack group you want tuning recommendations for. If both `attack_group` and `rule_id` not included, recommendations are returned for all attack groups.
        /// </summary>
        [Input("attackGroup")]
        public Input<string>? AttackGroup { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration you want tuning recommendations for.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Unique id of the rule you want tuning recommendations for. If both `attack_group` and `rule_id` not included, recommendations are returned for all attack groups.
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// . Type of ruleset used by the security configuration you want tuning recommendations for. Supported values are `active` and `evaluation`. Defaults to `active`.
        /// </summary>
        [Input("rulesetType")]
        public Input<string>? RulesetType { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy you want tuning recommendations for.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public GetAppSecTuningRecommendationsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecTuningRecommendationsResult
    {
        public readonly string? AttackGroup;
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// JSON-formatted list of the tuning recommendations for the security policy, the attack group or the rule. The exception block format in a recommendation conforms to the exception block format used in `condition_exception` element of `attack_group` or ASE rule resource.
        /// </summary>
        public readonly string Json;
        public readonly int? RuleId;
        public readonly string? RulesetType;
        public readonly string? SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecTuningRecommendationsResult(
            string? attackGroup,

            int configId,

            string id,

            string json,

            int? ruleId,

            string? rulesetType,

            string? securityPolicyId)
        {
            AttackGroup = attackGroup;
            ConfigId = configId;
            Id = id;
            Json = json;
            RuleId = ruleId;
            RulesetType = rulesetType;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
