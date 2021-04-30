// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecEvalRuleActions
    {
        /// <summary>
        /// Use the `akamai.getAppSecEvalRuleActions` data source to retrieve the rules available for evaluation and their actions, or the action for a specific rule available for evaluation.
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
        ///         var ruleActions = Output.Tuple(configuration, configuration).Apply(values =&gt;
        ///         {
        ///             var configuration = values.Item1;
        ///             var configuration1 = values.Item2;
        ///             return Output.Create(Akamai.GetAppSecEvalRuleActions.InvokeAsync(new Akamai.GetAppSecEvalRuleActionsArgs
        ///             {
        ///                 ConfigId = configuration.ConfigId,
        ///                 Version = configuration1.LatestVersion,
        ///                 SecurityPolicyId = @var.Security_policy_id,
        ///             }));
        ///         });
        ///         this.RuleActionsText = ruleActions.Apply(ruleActions =&gt; ruleActions.OutputText);
        ///         this.RuleActionsJson = ruleActions.Apply(ruleActions =&gt; ruleActions.Json);
        ///         var ruleActionAppSecEvalRuleActions = Output.Tuple(configuration, configuration).Apply(values =&gt;
        ///         {
        ///             var configuration = values.Item1;
        ///             var configuration1 = values.Item2;
        ///             return Output.Create(Akamai.GetAppSecEvalRuleActions.InvokeAsync(new Akamai.GetAppSecEvalRuleActionsArgs
        ///             {
        ///                 ConfigId = configuration.ConfigId,
        ///                 Version = configuration1.LatestVersion,
        ///                 SecurityPolicyId = @var.Security_policy_id,
        ///                 RuleId = @var.Rule_id,
        ///             }));
        ///         });
        ///         this.RuleAction = akamai_appsec_eval_rule_actions.Rule_action.Action;
        ///         this.RuleId = akamai_appsec_eval_rule_actions.Rule_action.Rule_id;
        ///     }
        /// 
        ///     [Output("ruleActionsText")]
        ///     public Output&lt;string&gt; RuleActionsText { get; set; }
        ///     [Output("ruleActionsJson")]
        ///     public Output&lt;string&gt; RuleActionsJson { get; set; }
        ///     [Output("ruleAction")]
        ///     public Output&lt;string&gt; RuleAction { get; set; }
        ///     [Output("ruleId")]
        ///     public Output&lt;string&gt; RuleId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecEvalRuleActionsResult> InvokeAsync(GetAppSecEvalRuleActionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecEvalRuleActionsResult>("akamai:index/getAppSecEvalRuleActions:getAppSecEvalRuleActions", args ?? new GetAppSecEvalRuleActionsArgs(), options.WithVersion());
    }


    public sealed class GetAppSecEvalRuleActionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of a specific rule. If not supplied, information about all eval rules will be returned.
        /// </summary>
        [Input("ruleId")]
        public int? RuleId { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version", required: true)]
        public int Version { get; set; }

        public GetAppSecEvalRuleActionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecEvalRuleActionsResult
    {
        /// <summary>
        /// The action configured for the given rule if a `rule_id` was specified.
        /// </summary>
        public readonly string Action;
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted display of the ID and action for all rules in the security policy.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display of the ID and action for all rules in the security policy.
        /// </summary>
        public readonly string OutputText;
        public readonly int? RuleId;
        public readonly string SecurityPolicyId;
        public readonly int Version;

        [OutputConstructor]
        private GetAppSecEvalRuleActionsResult(
            string action,

            int configId,

            string id,

            string json,

            string outputText,

            int? ruleId,

            string securityPolicyId,

            int version)
        {
            Action = action;
            ConfigId = configId;
            Id = id;
            Json = json;
            OutputText = outputText;
            RuleId = ruleId;
            SecurityPolicyId = securityPolicyId;
            Version = version;
        }
    }
}
