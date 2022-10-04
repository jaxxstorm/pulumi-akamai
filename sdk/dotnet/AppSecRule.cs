// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    /// <summary>
    /// **Scopes**: Rule
    /// 
    /// Modifies a Kona Rule Set rule's action, conditions, and exceptions.
    /// 
    /// **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/{ruleId}](https://techdocs.akamai.com/application-security/reference/put-rule) *and* [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/{ruleId}/condition-exception](https://techdocs.akamai.com/application-security/reference/put-rule-condition-exception)
    /// 
    /// ## Example Usage
    /// 
    /// Basic usage:
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Akamai = Pulumi.Akamai;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
    ///         {
    ///             Name = "Documentation",
    ///         }));
    ///         var rule = new Akamai.AppSecRule("rule", new Akamai.AppSecRuleArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             SecurityPolicyId = "gms1_134637",
    ///             RuleId = 60029316,
    ///             RuleAction = "deny",
    ///             ConditionException = File.ReadAllText($"{path.Module}/condition_exception.json"),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecRule:AppSecRule")]
    public partial class AppSecRule : Pulumi.CustomResource
    {
        /// <summary>
        /// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule.
        /// </summary>
        [Output("conditionException")]
        public Output<string?> ConditionException { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// Allowed values are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        /// - **none**. Take no action. or `none` to take no action.
        /// </summary>
        [Output("ruleAction")]
        public Output<string> RuleAction { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the rule being modified.
        /// </summary>
        [Output("ruleId")]
        public Output<int> RuleId { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecRule(string name, AppSecRuleArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecRule:AppSecRule", name, args ?? new AppSecRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecRule(string name, Input<string> id, AppSecRuleState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecRule:AppSecRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppSecRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecRule Get(string name, Input<string> id, AppSecRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecRule(name, id, state, options);
        }
    }

    public sealed class AppSecRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule.
        /// </summary>
        [Input("conditionException")]
        public Input<string>? ConditionException { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// Allowed values are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        /// - **none**. Take no action. or `none` to take no action.
        /// </summary>
        [Input("ruleAction")]
        public Input<string>? RuleAction { get; set; }

        /// <summary>
        /// . Unique identifier of the rule being modified.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<int> RuleId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public AppSecRuleArgs()
        {
        }
    }

    public sealed class AppSecRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule.
        /// </summary>
        [Input("conditionException")]
        public Input<string>? ConditionException { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// Allowed values are:
        /// - **alert**. Record the event.
        /// - **deny**. Block the request.
        /// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        /// - **none**. Take no action. or `none` to take no action.
        /// </summary>
        [Input("ruleAction")]
        public Input<string>? RuleAction { get; set; }

        /// <summary>
        /// . Unique identifier of the rule being modified.
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public AppSecRuleState()
        {
        }
    }
}
