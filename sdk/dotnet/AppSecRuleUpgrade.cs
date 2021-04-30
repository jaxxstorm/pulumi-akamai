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
    /// TBD
    /// Use the `akamai.AppSecRuleUpgrade` resource to upgrade to the most recent version of the KRS rule set. Akamai periodically updates these rules to keep protections current. However, the rules you use in your security policies do not automatically upgrade to the latest version when using mode: KRS. These rules do update automatically when you have mode set to AAG. Before you upgrade, run Get upgrade details to see which rules have changed. If you want to test how these rules would operate with live traffic before committing to the upgrade, run them in evaluation mode. This applies to KRS rules only and does not allow you to make any changes to the rules themselves. The response is the same as the mode response.
    /// 
    /// ## Example Usage
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
    ///         var ruleUpgrade = new Akamai.AppSecRuleUpgrade("ruleUpgrade", new Akamai.AppSecRuleUpgradeArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             Version = configuration.Apply(configuration =&gt; configuration.LatestVersion),
    ///             SecurityPolicyId = @var.Security_policy_id,
    ///         });
    ///         this.RuleUpgradeCurrentRuleset = ruleUpgrade.CurrentRuleset;
    ///         this.RuleUpgradeMode = ruleUpgrade.Mode;
    ///         this.RuleUpgradeEvalStatus = ruleUpgrade.EvalStatus;
    ///     }
    /// 
    ///     [Output("ruleUpgradeCurrentRuleset")]
    ///     public Output&lt;string&gt; RuleUpgradeCurrentRuleset { get; set; }
    ///     [Output("ruleUpgradeMode")]
    ///     public Output&lt;string&gt; RuleUpgradeMode { get; set; }
    ///     [Output("ruleUpgradeEvalStatus")]
    ///     public Output&lt;string&gt; RuleUpgradeEvalStatus { get; set; }
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecRuleUpgrade:AppSecRuleUpgrade")]
    public partial class AppSecRuleUpgrade : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// A string indicating the version number and release date of the current KRS rule set.
        /// </summary>
        [Output("currentRuleset")]
        public Output<string> CurrentRuleset { get; private set; } = null!;

        /// <summary>
        /// TBD
        /// </summary>
        [Output("evalStatus")]
        public Output<string> EvalStatus { get; private set; } = null!;

        /// <summary>
        /// A string indicating the current mode, either "KRS" or "AAG".
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecRuleUpgrade resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecRuleUpgrade(string name, AppSecRuleUpgradeArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecRuleUpgrade:AppSecRuleUpgrade", name, args ?? new AppSecRuleUpgradeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecRuleUpgrade(string name, Input<string> id, AppSecRuleUpgradeState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecRuleUpgrade:AppSecRuleUpgrade", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecRuleUpgrade resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecRuleUpgrade Get(string name, Input<string> id, AppSecRuleUpgradeState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecRuleUpgrade(name, id, state, options);
        }
    }

    public sealed class AppSecRuleUpgradeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public AppSecRuleUpgradeArgs()
        {
        }
    }

    public sealed class AppSecRuleUpgradeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// A string indicating the version number and release date of the current KRS rule set.
        /// </summary>
        [Input("currentRuleset")]
        public Input<string>? CurrentRuleset { get; set; }

        /// <summary>
        /// TBD
        /// </summary>
        [Input("evalStatus")]
        public Input<string>? EvalStatus { get; set; }

        /// <summary>
        /// A string indicating the current mode, either "KRS" or "AAG".
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public AppSecRuleUpgradeState()
        {
        }
    }
}
