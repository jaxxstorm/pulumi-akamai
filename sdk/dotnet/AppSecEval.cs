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
    /// **Scopes**: Security policy
    /// 
    /// Issues an evaluation mode command (`Start`, `Stop`, `Restart`, `Update`, or `Complete`) to a security configuration.
    /// Evaluation mode is used for testing and fine-tuning your Kona Rule Set rules and configuration settings.
    /// In evaluation mode rules are triggered by events, but the only thing those rules do is record the actions they *would* have taken had the event occurred on the production network.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval](https://techdocs.akamai.com/application-security/reference/post-policy-eval)
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
    ///             Name = "Documentation",
    ///         }));
    ///         var evalOperation = new Akamai.AppSecEval("evalOperation", new Akamai.AppSecEvalArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             SecurityPolicyId = "gms1_134637",
    ///             EvalOperation = "START",
    ///         });
    ///         this.EvalModeEvaluatingRuleset = evalOperation.EvaluatingRuleset;
    ///         this.EvalModeExpirationDate = evalOperation.ExpirationDate;
    ///         this.EvalModeCurrentRuleset = evalOperation.CurrentRuleset;
    ///         this.EvalModeStatus = evalOperation.EvalStatus;
    ///     }
    /// 
    ///     [Output("evalModeEvaluatingRuleset")]
    ///     public Output&lt;string&gt; EvalModeEvaluatingRuleset { get; set; }
    ///     [Output("evalModeExpirationDate")]
    ///     public Output&lt;string&gt; EvalModeExpirationDate { get; set; }
    ///     [Output("evalModeCurrentRuleset")]
    ///     public Output&lt;string&gt; EvalModeCurrentRuleset { get; set; }
    ///     [Output("evalModeStatus")]
    ///     public Output&lt;string&gt; EvalModeStatus { get; set; }
    /// }
    /// ```
    /// ## Output Options
    /// 
    /// The following options can be used to determine the information returned, and how that returned information is formatted:
    /// 
    /// - `evaluation_ruleset`. Versioning information for the Kona Rule Set being evaluated.
    /// - `expiration_date`. Date when the evaluation period ends.
    /// - `current_ruleset`. Versioning information for the Kona Rule Set currently in use on the production network.
    /// - `eval_status`. If **true**, an evaluation is currently in progress; if **false**, evaluation is either paused or is not running.
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecEval:AppSecEval")]
    public partial class AppSecEval : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        [Output("currentRuleset")]
        public Output<string> CurrentRuleset { get; private set; } = null!;

        /// <summary>
        /// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        /// </summary>
        [Output("evalMode")]
        public Output<string?> EvalMode { get; private set; } = null!;

        /// <summary>
        /// . Evaluation mode operation. Allowed values are:
        /// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
        /// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
        /// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
        /// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
        /// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        /// </summary>
        [Output("evalOperation")]
        public Output<string> EvalOperation { get; private set; } = null!;

        [Output("evalStatus")]
        public Output<string> EvalStatus { get; private set; } = null!;

        [Output("evaluatingRuleset")]
        public Output<string> EvaluatingRuleset { get; private set; } = null!;

        [Output("expirationDate")]
        public Output<string> ExpirationDate { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the evaluation process.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecEval resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecEval(string name, AppSecEvalArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecEval:AppSecEval", name, args ?? new AppSecEvalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecEval(string name, Input<string> id, AppSecEvalState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecEval:AppSecEval", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecEval resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecEval Get(string name, Input<string> id, AppSecEvalState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecEval(name, id, state, options);
        }
    }

    public sealed class AppSecEvalArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        /// </summary>
        [Input("evalMode")]
        public Input<string>? EvalMode { get; set; }

        /// <summary>
        /// . Evaluation mode operation. Allowed values are:
        /// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
        /// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
        /// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
        /// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
        /// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        /// </summary>
        [Input("evalOperation", required: true)]
        public Input<string> EvalOperation { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the evaluation process.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public AppSecEvalArgs()
        {
        }
    }

    public sealed class AppSecEvalState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        [Input("currentRuleset")]
        public Input<string>? CurrentRuleset { get; set; }

        /// <summary>
        /// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        /// </summary>
        [Input("evalMode")]
        public Input<string>? EvalMode { get; set; }

        /// <summary>
        /// . Evaluation mode operation. Allowed values are:
        /// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
        /// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
        /// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
        /// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
        /// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        /// </summary>
        [Input("evalOperation")]
        public Input<string>? EvalOperation { get; set; }

        [Input("evalStatus")]
        public Input<string>? EvalStatus { get; set; }

        [Input("evaluatingRuleset")]
        public Input<string>? EvaluatingRuleset { get; set; }

        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the evaluation process.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public AppSecEvalState()
        {
        }
    }
}
