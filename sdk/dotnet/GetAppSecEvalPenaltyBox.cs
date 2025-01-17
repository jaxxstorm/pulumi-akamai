// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecEvalPenaltyBox
    {
        /// <summary>
        /// **Scopes**: Security policy
        /// 
        ///  __ASE_Beta__.:
        /// Returns the penalty box settings for a security policy in evaluation mode - evaluation penalty box.
        /// When the penalty box is enabled for a policy in evaluation mode, clients that trigger a WAF Deny action are placed in the “penalty box”.
        /// There, the action you select for the penalty box (either Alert or Deny) continues to apply to any requests from that client for the next 10 minutes.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval_penalty-box](https://techdocs.akamai.com/application-security/reference/get-policy-eval_penalty-box)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var evalPenaltyBox = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecEvalPenaltyBox.InvokeAsync(new Akamai.GetAppSecEvalPenaltyBoxArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.EvalPenaltyBoxAction = evalPenaltyBox.Apply(evalPenaltyBox =&gt; evalPenaltyBox.Action);
        ///         this.EvalPenaltyBoxEnabled = evalPenaltyBox.Apply(evalPenaltyBox =&gt; evalPenaltyBox.Enabled);
        ///         this.EvalPenaltyBoxText = evalPenaltyBox.Apply(evalPenaltyBox =&gt; evalPenaltyBox.OutputText);
        ///     }
        /// 
        ///     [Output("evalPenaltyBoxAction")]
        ///     public Output&lt;string&gt; EvalPenaltyBoxAction { get; set; }
        ///     [Output("evalPenaltyBoxEnabled")]
        ///     public Output&lt;string&gt; EvalPenaltyBoxEnabled { get; set; }
        ///     [Output("evalPenaltyBoxText")]
        ///     public Output&lt;string&gt; EvalPenaltyBoxText { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `action`. Action taken any time the penalty box is triggered. Valid values are:
        ///   - **alert**. Record the event.
        ///   - **deny**. The request is blocked.
        ///   - **deny_custom_{custom_deny_id}**. The action defined by the custom deny is taken.
        ///   - **none**. Take no action.
        /// - `enabled`. If **true**, evaluation penalty box protection is enabled. If **false**, evaluation penalty box protection is disabled.
        /// - `output_text`. Tabular report of evaluation penalty box protection settings.
        /// </summary>
        public static Task<GetAppSecEvalPenaltyBoxResult> InvokeAsync(GetAppSecEvalPenaltyBoxArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecEvalPenaltyBoxResult>("akamai:index/getAppSecEvalPenaltyBox:getAppSecEvalPenaltyBox", args ?? new GetAppSecEvalPenaltyBoxArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security policy
        /// 
        ///  __ASE_Beta__.:
        /// Returns the penalty box settings for a security policy in evaluation mode - evaluation penalty box.
        /// When the penalty box is enabled for a policy in evaluation mode, clients that trigger a WAF Deny action are placed in the “penalty box”.
        /// There, the action you select for the penalty box (either Alert or Deny) continues to apply to any requests from that client for the next 10 minutes.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval_penalty-box](https://techdocs.akamai.com/application-security/reference/get-policy-eval_penalty-box)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var evalPenaltyBox = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecEvalPenaltyBox.InvokeAsync(new Akamai.GetAppSecEvalPenaltyBoxArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.EvalPenaltyBoxAction = evalPenaltyBox.Apply(evalPenaltyBox =&gt; evalPenaltyBox.Action);
        ///         this.EvalPenaltyBoxEnabled = evalPenaltyBox.Apply(evalPenaltyBox =&gt; evalPenaltyBox.Enabled);
        ///         this.EvalPenaltyBoxText = evalPenaltyBox.Apply(evalPenaltyBox =&gt; evalPenaltyBox.OutputText);
        ///     }
        /// 
        ///     [Output("evalPenaltyBoxAction")]
        ///     public Output&lt;string&gt; EvalPenaltyBoxAction { get; set; }
        ///     [Output("evalPenaltyBoxEnabled")]
        ///     public Output&lt;string&gt; EvalPenaltyBoxEnabled { get; set; }
        ///     [Output("evalPenaltyBoxText")]
        ///     public Output&lt;string&gt; EvalPenaltyBoxText { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `action`. Action taken any time the penalty box is triggered. Valid values are:
        ///   - **alert**. Record the event.
        ///   - **deny**. The request is blocked.
        ///   - **deny_custom_{custom_deny_id}**. The action defined by the custom deny is taken.
        ///   - **none**. Take no action.
        /// - `enabled`. If **true**, evaluation penalty box protection is enabled. If **false**, evaluation penalty box protection is disabled.
        /// - `output_text`. Tabular report of evaluation penalty box protection settings.
        /// </summary>
        public static Output<GetAppSecEvalPenaltyBoxResult> Invoke(GetAppSecEvalPenaltyBoxInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecEvalPenaltyBoxResult>("akamai:index/getAppSecEvalPenaltyBox:getAppSecEvalPenaltyBox", args ?? new GetAppSecEvalPenaltyBoxInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecEvalPenaltyBoxArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the evaluation penalty box settings.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the evaluation penalty box settings.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        public GetAppSecEvalPenaltyBoxArgs()
        {
        }
    }

    public sealed class GetAppSecEvalPenaltyBoxInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the evaluation penalty box settings.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the evaluation penalty box settings.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public GetAppSecEvalPenaltyBoxInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecEvalPenaltyBoxResult
    {
        public readonly string Action;
        public readonly int ConfigId;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OutputText;
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecEvalPenaltyBoxResult(
            string action,

            int configId,

            bool enabled,

            string id,

            string outputText,

            string securityPolicyId)
        {
            Action = action;
            ConfigId = configId;
            Enabled = enabled;
            Id = id;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
