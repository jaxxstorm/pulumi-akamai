// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Akamai
{
    public static class GetAppSecWafMode
    {
        /// <summary>
        /// Use the `akamai.AppSecWafMode` data source to retrieve the mode that indicates how the WAF rules of the given security configuration and security policy will be updated.
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
        ///         var wafMode = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecWafMode.InvokeAsync(new Akamai.GetAppSecWafModeArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Policy_id,
        ///         })));
        ///         this.WafModeMode = wafMode.Apply(wafMode =&gt; wafMode.Mode);
        ///         this.WafModeCurrentRuleset = wafMode.Apply(wafMode =&gt; wafMode.CurrentRuleset);
        ///         this.WafModeEvalStatus = wafMode.Apply(wafMode =&gt; wafMode.EvalStatus);
        ///         this.WafModeEvalRuleset = wafMode.Apply(wafMode =&gt; wafMode.EvalRuleset);
        ///         this.WafModeEvalExpirationDate = wafMode.Apply(wafMode =&gt; wafMode.EvalExpirationDate);
        ///         this.WafModeText = wafMode.Apply(wafMode =&gt; wafMode.OutputText);
        ///         this.WafModeJson = wafMode.Apply(wafMode =&gt; wafMode.Json);
        ///     }
        /// 
        ///     [Output("wafModeMode")]
        ///     public Output&lt;string&gt; WafModeMode { get; set; }
        ///     [Output("wafModeCurrentRuleset")]
        ///     public Output&lt;string&gt; WafModeCurrentRuleset { get; set; }
        ///     [Output("wafModeEvalStatus")]
        ///     public Output&lt;string&gt; WafModeEvalStatus { get; set; }
        ///     [Output("wafModeEvalRuleset")]
        ///     public Output&lt;string&gt; WafModeEvalRuleset { get; set; }
        ///     [Output("wafModeEvalExpirationDate")]
        ///     public Output&lt;string&gt; WafModeEvalExpirationDate { get; set; }
        ///     [Output("wafModeText")]
        ///     public Output&lt;string&gt; WafModeText { get; set; }
        ///     [Output("wafModeJson")]
        ///     public Output&lt;string&gt; WafModeJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecWafModeResult> InvokeAsync(GetAppSecWafModeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecWafModeResult>("akamai:index/getAppSecWafMode:getAppSecWafMode", args ?? new GetAppSecWafModeArgs(), options.WithVersion());

        /// <summary>
        /// Use the `akamai.AppSecWafMode` data source to retrieve the mode that indicates how the WAF rules of the given security configuration and security policy will be updated.
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
        ///         var wafMode = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecWafMode.InvokeAsync(new Akamai.GetAppSecWafModeArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Policy_id,
        ///         })));
        ///         this.WafModeMode = wafMode.Apply(wafMode =&gt; wafMode.Mode);
        ///         this.WafModeCurrentRuleset = wafMode.Apply(wafMode =&gt; wafMode.CurrentRuleset);
        ///         this.WafModeEvalStatus = wafMode.Apply(wafMode =&gt; wafMode.EvalStatus);
        ///         this.WafModeEvalRuleset = wafMode.Apply(wafMode =&gt; wafMode.EvalRuleset);
        ///         this.WafModeEvalExpirationDate = wafMode.Apply(wafMode =&gt; wafMode.EvalExpirationDate);
        ///         this.WafModeText = wafMode.Apply(wafMode =&gt; wafMode.OutputText);
        ///         this.WafModeJson = wafMode.Apply(wafMode =&gt; wafMode.Json);
        ///     }
        /// 
        ///     [Output("wafModeMode")]
        ///     public Output&lt;string&gt; WafModeMode { get; set; }
        ///     [Output("wafModeCurrentRuleset")]
        ///     public Output&lt;string&gt; WafModeCurrentRuleset { get; set; }
        ///     [Output("wafModeEvalStatus")]
        ///     public Output&lt;string&gt; WafModeEvalStatus { get; set; }
        ///     [Output("wafModeEvalRuleset")]
        ///     public Output&lt;string&gt; WafModeEvalRuleset { get; set; }
        ///     [Output("wafModeEvalExpirationDate")]
        ///     public Output&lt;string&gt; WafModeEvalExpirationDate { get; set; }
        ///     [Output("wafModeText")]
        ///     public Output&lt;string&gt; WafModeText { get; set; }
        ///     [Output("wafModeJson")]
        ///     public Output&lt;string&gt; WafModeJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAppSecWafModeResult> Invoke(GetAppSecWafModeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecWafModeResult>("akamai:index/getAppSecWafMode:getAppSecWafMode", args ?? new GetAppSecWafModeInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAppSecWafModeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        public GetAppSecWafModeArgs()
        {
        }
    }

    public sealed class GetAppSecWafModeInvokeArgs : Pulumi.InvokeArgs
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

        public GetAppSecWafModeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecWafModeResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The current rule set version and the ISO 8601 date the rule set version was introduced; this date acts like a version number.
        /// </summary>
        public readonly string CurrentRuleset;
        /// <summary>
        /// The ISO 8601 time stamp when the evaluation is expiring. This value only appears when `eval` is set to "enabled".
        /// </summary>
        public readonly string EvalExpirationDate;
        /// <summary>
        /// The evaluation rule set version and the ISO 8601 date the evaluation starts.
        /// </summary>
        public readonly string EvalRuleset;
        /// <summary>
        /// Whether the evaluation mode is enabled or disabled."
        /// </summary>
        public readonly string EvalStatus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted list of the mode information.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// The security policy mode, either `KRS` (update manually) or `AAG` (update automatically), For Adaptive Security Engine (ASE) __BETA__, use `ASE_AUTO` for automatic updates or `ASE_MANUAL` to manually get current rules. Please contact your Akamai representative to learn more about ASE.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// A tabular display of the mode information.
        /// </summary>
        public readonly string OutputText;
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecWafModeResult(
            int configId,

            string currentRuleset,

            string evalExpirationDate,

            string evalRuleset,

            string evalStatus,

            string id,

            string json,

            string mode,

            string outputText,

            string securityPolicyId)
        {
            ConfigId = configId;
            CurrentRuleset = currentRuleset;
            EvalExpirationDate = evalExpirationDate;
            EvalRuleset = evalRuleset;
            EvalStatus = evalStatus;
            Id = id;
            Json = json;
            Mode = mode;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
